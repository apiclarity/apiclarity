// Copyright © 2022 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package traceanalyzer

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/openclarity/apiclarity/api/server/models"
	oapicommon "github.com/openclarity/apiclarity/api3/common"
	"github.com/openclarity/apiclarity/api3/notifications"
	"github.com/openclarity/apiclarity/backend/pkg/config"
	"github.com/openclarity/apiclarity/backend/pkg/database"
	"github.com/openclarity/apiclarity/backend/pkg/modules/internal/core"

	"github.com/openclarity/apiclarity/backend/pkg/modules/internal/traceanalyzer/guessableid"
	"github.com/openclarity/apiclarity/backend/pkg/modules/internal/traceanalyzer/nlid"
	"github.com/openclarity/apiclarity/backend/pkg/modules/internal/traceanalyzer/restapi"
	"github.com/openclarity/apiclarity/backend/pkg/modules/internal/traceanalyzer/sensitive"
	"github.com/openclarity/apiclarity/backend/pkg/modules/internal/traceanalyzer/utils"
	"github.com/openclarity/apiclarity/backend/pkg/modules/internal/traceanalyzer/weakbasicauth"
	"github.com/openclarity/apiclarity/backend/pkg/modules/internal/traceanalyzer/weakjwt"
)

const (
	dictFilenamesEnvVar  = "TRACE_ANALYZER_DICT_FILENAMES"
	dictFilenamesDefault = ""

	rulesFilenamesEnvVar  = "TRACE_ANALYZER_RULES_FILENAMES"
	rulesFilenamesDefault = ""

	sensitiveKeywordsFilenamesEnvVar  = "TRACE_ANALYZER_SENSITIVE_KEYWORDS_FILENAMES"
	sensitiveKeywordsFilenamesDefault = ""

	ignoreFindingsEnvVar  = "TRACE_ANALYZER_IGNORE_FINDINGS"
	ignoreFindingsDefault = ""
)

type ParameterFinding struct {
	Location string      `json:"location"`
	Method   string      `json:"method"`
	Name     string      `json:"name"`
	Value    string      `json:"value"`
	Reason   interface{} `json:"reason"`
}

type traceAnalyzerConfig struct {
	dictFilenames              []string `yaml:"dictFilenames"`
	rulesFilenames             []string `yaml:"rulesFilenames"`
	sensitiveKeywordsFilenames []string `yaml:"keywordsFilenames"`
	ignoreFindings             []string `yaml:"ignoreFindings"`
}

type traceAnalyzer struct {
	httpHandler http.Handler

	config traceAnalyzerConfig

	ignoreFindings map[string]bool

	guessableID   *guessableid.GuessableAnalyzer
	nlid          *nlid.NLID
	weakBasicAuth *weakbasicauth.WeakBasicAuth
	weakJWT       *weakjwt.WeakJWT
	sensitive     *sensitive.Sensitive

	aggregator *APIsFindingsRepo

	accessor core.BackendAccessor
}

func newTraceAnalyzer(ctx context.Context, accessor core.BackendAccessor) (core.Module, error) {
	var err error

	p := traceAnalyzer{}
	h := restapi.HandlerWithOptions(&httpHandler{ta: &p}, restapi.ChiServerOptions{BaseURL: core.BaseHTTPPath + "/" + utils.ModuleName})
	p.httpHandler = h
	p.ignoreFindings = map[string]bool{}
	p.accessor = accessor

	p.config = loadConfig()
	log.Debugf("TraceAnalyzer Configuration: %+v", p.config)

	for _, ifinding := range p.config.ignoreFindings {
		p.ignoreFindings[ifinding] = true
	}

	passwordList, err := utils.ReadDictionaryFiles(p.config.dictFilenames)
	if err != nil {
		return nil, fmt.Errorf("unable to read password files: %w", err)
	}
	weakKeyList, err := utils.ReadDictionaryFiles(p.config.dictFilenames)
	if err != nil {
		return nil, fmt.Errorf("unable to read list of weak keys: %w", err)
	}
	sensitiveKeywords, err := utils.ReadDictionaryFiles(p.config.sensitiveKeywordsFilenames)
	if err != nil {
		return nil, fmt.Errorf("unable to read list of sensitive keywords: %w", err)
	}

	p.guessableID = guessableid.NewGuessableAnalyzer(guessableid.MaxParamHistory)
	p.nlid = nlid.NewNLID(nlid.NLIDRingBufferSize)
	p.weakBasicAuth = weakbasicauth.NewWeakBasicAuth(passwordList)
	p.weakJWT = weakjwt.NewWeakJWT(weakKeyList, sensitiveKeywords)
	if p.sensitive, err = sensitive.NewSensitive(p.config.rulesFilenames); err != nil {
		return nil, fmt.Errorf("unable to initialize Trace Analyzer Regexp Rules: %w", err)
	}

	p.aggregator = NewAPIsFindingsRepo(p.accessor)

	return &p, nil
}

func parseFilenamesFromEnv(filenames string) []string {
	if filenames == "" {
		return []string{}
	}
	fns := strings.Split(filenames, ":")
	for i := range fns {
		fns[i] = strings.TrimSpace(fns[i])
	}

	return fns
}

func loadConfig() traceAnalyzerConfig {
	viper.SetDefault(dictFilenamesEnvVar, dictFilenamesDefault)
	viper.SetDefault(rulesFilenamesEnvVar, rulesFilenamesDefault)
	viper.SetDefault(sensitiveKeywordsFilenamesEnvVar, sensitiveKeywordsFilenamesDefault)
	viper.SetDefault(ignoreFindingsEnvVar, ignoreFindingsDefault)

	dictFilenames := parseFilenamesFromEnv(viper.GetString(dictFilenamesEnvVar))
	rulesFilenames := parseFilenamesFromEnv(viper.GetString(rulesFilenamesEnvVar))
	keywordsFilenames := parseFilenamesFromEnv(viper.GetString(sensitiveKeywordsFilenamesEnvVar))
	ignoreFindings := viper.GetStringSlice(ignoreFindingsEnvVar)
	modulesAssets := viper.GetString(config.ModulesAssetsEnvVar)

	var err error
	if modulesAssets != "" {
		if len(dictFilenames) == 0 {
			dictFilenames, err = utils.WalkFiles(filepath.Join(modulesAssets, utils.ModuleName, "dictionaries"))
			if err != nil {
				log.Warnf("There was problem while reading the Trace Analyzer assets directory 'dictionaries': %s", err)
			}
		}
		if len(rulesFilenames) == 0 {
			rulesFilenames, err = utils.WalkFiles(filepath.Join(modulesAssets, utils.ModuleName, "sensitive_rules"))
			if err != nil {
				log.Warnf("There was problem while reading the Trace Analyzer assets directory 'sensitive_rules': %s", err)
			}
		}
		if len(keywordsFilenames) == 0 {
			keywordsFilenames, err = utils.WalkFiles(filepath.Join(modulesAssets, utils.ModuleName, "sensitive_keywords"))
			if err != nil {
				log.Warnf("There was problem while reading the Trace Analyzer assets directory 'sensitive_keywords': %s", err)
			}
		}
	}

	c := traceAnalyzerConfig{
		dictFilenames:              dictFilenames,
		rulesFilenames:             rulesFilenames,
		sensitiveKeywordsFilenames: keywordsFilenames,
		ignoreFindings:             ignoreFindings,
	}
	return c
}

func (p *traceAnalyzer) Name() string {
	return utils.ModuleName
}

func (p *traceAnalyzer) HTTPHandler() http.Handler {
	return p.httpHandler
}

func (p *traceAnalyzer) EventNotify(ctx context.Context, e *core.Event) {
	event, trace := e.APIEvent, e.Telemetry
	log.Debugf("[TraceAnalyzer] received a new trace for API(%v) EventID(%v)", event.APIInfoID, event.ID)
	eventAnns := []utils.TraceAnalyzerAnnotation{}
	apiAnns := []utils.TraceAnalyzerAPIAnnotation{}

	wbaEventAnns, wbaAPIAnns := p.weakBasicAuth.Analyze(trace)
	eventAnns = append(eventAnns, wbaEventAnns...)
	apiAnns = append(apiAnns, wbaAPIAnns...)

	wjtEventAnns, wjtAPIAnns := p.weakJWT.Analyze(trace)
	eventAnns = append(eventAnns, wjtEventAnns...)
	apiAnns = append(apiAnns, wjtAPIAnns...)

	sensEventAnns, sensAPIAnns := p.sensitive.Analyze(trace)
	eventAnns = append(eventAnns, sensEventAnns...)
	apiAnns = append(apiAnns, sensAPIAnns...)

	// If the status code starts with 2, it means that the request has been
	// accepted, hence, the parameters were accepted as well. So, we can look at
	// the parameters to see if they are very similar with the one in previous
	// accepted queries.
	// FIXME: Performance KILLER. For each request, this function is called, which calls the database a deserializes data to get the specinfo
	// FIXME: We MUST create a memory cache map[apiid]specinfo to avoid that
	specPath, pathParams, _, _, _, err := p.getParams(ctx, event)
	// if specPath == "" {
	// 	specPath = trace.Request.Path
	// }
	if err == nil && strings.HasPrefix(trace.Response.StatusCode, "2") {
		// Check for guessable IDs
		eventGuessable, _ := p.guessableID.Analyze(specPath, string(event.Method), pathParams, trace)
		eventAnns = append(eventAnns, eventGuessable...)

		// Check for NLIDS
		eventNLIDAnns, _ := p.nlid.Analyze(specPath, string(event.Method), pathParams, trace)
		eventAnns = append(eventAnns, eventNLIDAnns...)
	}

	// Filter ignored findings for event annotations
	filteredEventAnns := []utils.TraceAnalyzerAnnotation{}
	for _, a := range eventAnns {
		if !p.ignoreFindings[a.Name()] {
			filteredEventAnns = append(filteredEventAnns, a)
		}
	}

	if len(filteredEventAnns) > 0 {
		coreEventAnnotations := p.toCoreEventAnnotations(filteredEventAnns, false)
		if err := p.accessor.CreateAPIEventAnnotations(ctx, p.Name(), event.ID, coreEventAnnotations...); err != nil {
			log.Error(err)
		}
		p.setAlertSeverity(ctx, event.ID, filteredEventAnns)
	}

	if len(filteredEventAnns) > 0 {
		updated := p.aggregator.Aggregate(uint64(event.APIInfoID), specPath, trace.Request.Method, filteredEventAnns...)
		if updated {
			// Filter ignored findings for API annotations
			filteredAPIAnns := []utils.TraceAnalyzerAPIAnnotation{}
			for _, a := range p.aggregator.GetAPIFindings(uint64(event.APIInfoID)) {
				if !p.ignoreFindings[a.Name()] {
					filteredAPIAnns = append(filteredAPIAnns, a)
				}
			}
			if len(filteredAPIAnns) > 0 {
				coreAPIAnnotations := p.toCoreAPIAnnotations(filteredAPIAnns, false)
				if err := p.accessor.StoreAPIInfoAnnotations(ctx, p.Name(), event.APIInfoID, coreAPIAnnotations...); err != nil {
					log.Error(err)
				}
				err := p.sendAPIFindingsNotification(ctx, event.APIInfoID, filteredAPIAnns)
				if err != nil {
					log.Error(err)
				}
			}
		}
	}

	return
}

func (p *traceAnalyzer) toCoreEventAnnotations(eventAnns []utils.TraceAnalyzerAnnotation, redacted bool) (coreAnnotations []core.Annotation) {
	for _, a := range eventAnns {
		if redacted {
			a = a.Redacted()
		}
		annotation, err := a.Serialize()
		if err != nil {
			log.Errorf("unable to serialize annotation: %s", err)
		}
		coreAnnotations = append(coreAnnotations, core.Annotation{Name: a.Name(), Annotation: annotation})
	}
	return coreAnnotations
}

func fromCoreEventAnnotation(coreAnn *core.Annotation) (ann utils.TraceAnalyzerAnnotation, err error) {
	var a utils.TraceAnalyzerAnnotation
	switch coreAnn.Name {
	case weakbasicauth.KindKnownPassword:
		a = &weakbasicauth.AnnotationKnownPassword{}
	case weakbasicauth.KindShortPassword:
		a = &weakbasicauth.AnnotationShortPassword{}
	case weakbasicauth.KindSamePassword:
		a = &weakbasicauth.AnnotationSamePassword{}

	case weakjwt.JWTNoAlgField:
		a = &weakjwt.AnnotationNoAlgField{}
	case weakjwt.JWTAlgFieldNone:
		a = &weakjwt.AnnotationAlgFieldNone{}
	case weakjwt.JWTNotRecommendedAlg:
		a = &weakjwt.AnnotationNotRecommendedAlg{}
	case weakjwt.JWTNoExpireClaim:
		a = &weakjwt.AnnotationNoExpireClaim{}
	case weakjwt.JWTExpTooFar:
		a = &weakjwt.AnnotationExpTooFar{}
	case weakjwt.JWTWeakSymetricSecret:
		a = &weakjwt.AnnotationWeakSymetricSecret{}
	case weakjwt.JWTSensitiveContentInHeaders:
		a = &weakjwt.AnnotationSensitiveContentInHeaders{}
	case weakjwt.JWTSensitiveContentInClaims:
		a = &weakjwt.AnnotationSensitiveContentInClaims{}

	case sensitive.RegexpMatchingType:
		a = &sensitive.AnnotationRegexpMatching{}

	case nlid.NLIDType:
		a = &nlid.AnnotationNLID{}

	case guessableid.GuessableType:
		a = &guessableid.AnnotationGuessableID{}

	default:
		return nil, fmt.Errorf("unknown annotation '%s'", coreAnn.Name)
	}

	err = a.Deserialize(coreAnn.Annotation)
	return a, err
}

func fromCoreEventAnnotations(coreAnns []*core.Annotation) (anns []utils.TraceAnalyzerAnnotation) {
	for _, coreAnn := range coreAnns {
		taAnn, err := fromCoreEventAnnotation(coreAnn)
		if err != nil {
			log.Errorf("Unable to understand annotation: %v", err)
		} else {
			anns = append(anns, taAnn)
		}
	}

	return anns
}

func (p *traceAnalyzer) toCoreAPIAnnotations(anns []utils.TraceAnalyzerAPIAnnotation, redacted bool) (coreAnnotations []core.Annotation) {
	for _, a := range anns {
		if redacted {
			a = a.Redacted()
		}
		annotation, err := a.Serialize()
		if err != nil {
			log.Errorf("unable to serialize annotation: %s", err)
		}
		coreAnnotations = append(coreAnnotations, core.Annotation{Name: a.Name(), Annotation: annotation})
	}
	return coreAnnotations
}

func fromCoreAPIAnnotation(coreAnn *core.Annotation) (ann utils.TraceAnalyzerAPIAnnotation, err error) {
	var a utils.TraceAnalyzerAPIAnnotation
	switch coreAnn.Name {
	case weakbasicauth.KindShortPassword:
		a = &weakbasicauth.APIAnnotationShortPassword{}
	case weakbasicauth.KindKnownPassword:
		a = &weakbasicauth.APIAnnotationKnownPassword{}
	case weakbasicauth.KindSamePassword:
		a = &weakbasicauth.APIAnnotationSamePassword{}

	case sensitive.RegexpMatchingType:
		a = &sensitive.APIAnnotationRegexpMatching{}

	case nlid.NLIDType:
		a = &nlid.APIAnnotationNLID{}
	case guessableid.GuessableType:
		a = &guessableid.APIAnnotationGuessableID{}


	default:
		return nil, fmt.Errorf("unknown annotation '%s'", coreAnn.Name)
	}

	err = a.Deserialize(coreAnn.Annotation)
	return a, err
}

func fromCoreAPIAnnotations(coreAnns []*core.Annotation) (anns []utils.TraceAnalyzerAPIAnnotation) {
	for _, coreAnn := range coreAnns {
		taAnn, err := fromCoreAPIAnnotation(coreAnn)
		if err != nil {
			log.Errorf("Unable to understand annotation: %v", err)
		} else {
			anns = append(anns, taAnn)
		}
	}

	return anns
}

func (p *traceAnalyzer) sendAPIFindingsNotification(ctx context.Context, apiID uint, apiFindings []utils.TraceAnalyzerAPIAnnotation) error {
	apiN := notifications.ApiFindingsNotification{
		NotificationType: "ApiFindingsNotification",
		Items:            &[]oapicommon.APIFinding{},
	}

	for _, finding := range apiFindings {
		*(apiN.Items) = append(*(apiN.Items), finding.ToAPIFinding())
	}

	n := notifications.APIClarityNotification{}
	n.FromApiFindingsNotification(apiN)

	err := p.accessor.Notify(ctx, utils.ModuleName, apiID, n)

	return err
}

func (p *traceAnalyzer) EventAnnotationNotify(modName string, eventID uint, ann core.Annotation) error {
	return nil
}

func (p *traceAnalyzer) APIAnnotationNotify(modName string, apiID uint, annotation *core.Annotation) error {
	return nil
}

func (p *traceAnalyzer) setAlertSeverity(ctx context.Context, eventID uint, anns []utils.TraceAnalyzerAnnotation) {
	maxAlert := core.AlertInfo
	for _, a := range anns {
		alert := utils.SeverityToAlert(a.Severity())
		if alert > maxAlert {
			maxAlert = alert
		}
		// We reach the maximum alert level, not need to go further
		if maxAlert == core.AlertCritical {
			break
		}
	}

	var alertAnn core.Annotation
	switch maxAlert {
	case core.AlertInfo:
		alertAnn = core.AlertInfoAnn
	case core.AlertWarn:
		alertAnn = core.AlertWarnAnn
	case core.AlertCritical:
		alertAnn = core.AlertCriticalAnn
	}

	if err := p.accessor.CreateAPIEventAnnotations(ctx, p.Name(), eventID, alertAnn); err != nil {
		log.Error(err)
	}
}

// XXX There are too many parameters to this function. It needs refactoring.
func (p *traceAnalyzer) getParams(ctx context.Context, event *database.APIEvent) (specPath string, pathParams map[string]string, queryParams map[string]string, headerParams map[string]string, bodyParams map[string]string, err error) {
	apiInfo, err := p.accessor.GetAPIInfo(ctx, event.APIInfoID)
	if err != nil {
		return "", nil, nil, nil, nil, err
	}

	// Prefer Provided specification if available
	var serializedSpecInfo *string
	var eventPathID string
	if apiInfo.HasProvidedSpec && apiInfo.ProvidedSpecInfo != "" {
		serializedSpecInfo = &apiInfo.ProvidedSpecInfo
		eventPathID = event.ProvidedPathID
	} else if apiInfo.HasReconstructedSpec && apiInfo.ReconstructedSpecInfo != "" {
		serializedSpecInfo = &apiInfo.ReconstructedSpecInfo
		eventPathID = event.ReconstructedPathID
	} else {
		return specPath, pathParams, queryParams, headerParams, bodyParams, nil
	}

	var specInfo models.SpecInfo
	if err := json.Unmarshal([]byte(*serializedSpecInfo), &specInfo); err != nil {
		return specPath, pathParams, queryParams, headerParams, bodyParams, fmt.Errorf("failed to unmarshal spec info for api=%d: %v", event.APIInfoID, err)
	}

	pathParams = make(map[string]string)
	queryParams = make(map[string]string)
	headerParams = make(map[string]string)
	bodyParams = make(map[string]string)

	for _, t := range specInfo.Tags {
		for _, path := range t.MethodAndPathList {
			if path.PathID.String() == eventPathID && path.Method == event.Method {
				specPath = path.Path
				pathParams = utils.GetPathParams(path.Path, event.Path)
				// XXX Need to get other parameters
				break
			}
		}
	}

	return specPath, pathParams, queryParams, headerParams, bodyParams, nil
}

func (p *traceAnalyzer) getAPIFindings(ctx context.Context, apiID uint, sensitive bool) (apiFindings []oapicommon.APIFinding, err error) {
	dbAnns, err := p.accessor.ListAPIInfoAnnotations(ctx, utils.ModuleName, uint(apiID))
	if err != nil {
		return apiFindings, err
	}

	anns := fromCoreAPIAnnotations(dbAnns)
	for _, ann := range anns {
		var f oapicommon.APIFinding
		if sensitive {
			f = ann.ToAPIFinding()
		} else {
			f = ann.Redacted().ToAPIFinding()
		}
		apiFindings = append(apiFindings, f)
	}

	return apiFindings, nil
}

type httpHandler struct {
	ta *traceAnalyzer
}

func (h httpHandler) GetEventAnnotations(w http.ResponseWriter, r *http.Request, eventID int64, params restapi.GetEventAnnotationsParams) {
	dbAnns, err := h.ta.accessor.ListAPIEventAnnotations(r.Context(), utils.ModuleName, uint(eventID))
	if err != nil {
		return
	}
	annList := []restapi.Annotation{}

	taAnns := fromCoreEventAnnotations(dbAnns)
	for _, a := range taAnns {
		if params.Redacted != nil && *params.Redacted {
			a = a.Redacted()
		}
		f := a.ToFinding()
		annList = append(annList, restapi.Annotation{
			Annotation: f.DetailedDesc,
			Name:       f.ShortDesc,
			Severity:   f.Severity,
			Kind:       a.Name(),
		})
	}

	result := restapi.Annotations{
		Items: &annList,
		Total: len(annList),
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h httpHandler) GetAPIAnnotations(w http.ResponseWriter, r *http.Request, apiID int64, params restapi.GetAPIAnnotationsParams) {
	dbAnns, err := h.ta.accessor.ListAPIInfoAnnotations(r.Context(), utils.ModuleName, uint(apiID))
	if err != nil {
		return
	}
	annList := []restapi.Annotation{}

	taAnns := fromCoreAPIAnnotations(dbAnns)
	for _, a := range taAnns {
		if params.Redacted != nil && *params.Redacted {
			a = a.Redacted()
		}
		f := a.ToFinding()
		annList = append(annList, restapi.Annotation{
			Annotation: f.DetailedDesc,
			Name:       f.ShortDesc,
			Severity:   f.Severity,
			Kind:       a.Name(),
		})
	}
	result := restapi.Annotations{
		Items: &annList,
		Total: len(annList),
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h httpHandler) DeleteAPIAnnotations(w http.ResponseWriter, r *http.Request, apiID int64, params restapi.DeleteAPIAnnotationsParams) {
	err := h.ta.accessor.DeleteAPIInfoAnnotations(r.Context(), utils.ModuleName, uint(apiID), params.Name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h httpHandler) GetApiFindings(w http.ResponseWriter, r *http.Request, apiID oapicommon.ApiID, params restapi.GetApiFindingsParams) {
	// If sensitive parameter is not set, default to false (ie: do not include sensitive data)
	sensitive := params.Sensitive != nil && *params.Sensitive
	apiFindings, err := h.ta.getAPIFindings(r.Context(), uint(apiID), sensitive)
	if err != nil {
		err := oapicommon.ApiResponse{Message: "Internal error, could not read data from database"}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	apiFindingsObject := oapicommon.APIFindings{
		Items: &apiFindings,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(apiFindingsObject)
}

//nolint:gochecknoinits
func init() {
	core.RegisterModule(newTraceAnalyzer)
}
