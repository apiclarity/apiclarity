package backend

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"

	"github.com/apiclarity/apiclarity/backend/pkg/database"
	"github.com/apiclarity/apiclarity/backend/pkg/modules"
)

type backendAccessor struct {
	dbHandler *database.Handler
	clientset kubernetes.Interface
}

func (b backendAccessor) K8SClient() kubernetes.Interface {
	return b.clientset
}

func (b backendAccessor) GetAPIInfo(ctx context.Context, apiID uint) (*database.APIInfo, error) {
	apiInfo := &database.APIInfo{}
	if err := b.dbHandler.APIInventoryTable().First(apiInfo, apiID); err != nil {
		log.Errorf("Failed to retreive API info for apiID=%v: %v", apiID, err)
		return nil, fmt.Errorf("failed to retreive API info for apiID=%v: %v", apiID, err)
	}
	return apiInfo, nil
}

func (b backendAccessor) GetAPIEvents(ctx context.Context, filter database.GetAPIEventsQuery) ([]*database.APIEvent, error) {
	return b.dbHandler.APIEventsTable().GetAPIEvents(ctx, filter)
}

func (b backendAccessor) GetAPIEventAnnotation(ctx context.Context, modName string, eventID uint, name string) (*modules.Annotation, error) {
	ann, err := b.dbHandler.APIEventsAnnotationsTable().Get(ctx, modName, eventID, name)
	if err != nil {
		return nil, err
	}
	return &modules.Annotation{Name: ann.Name, Annotation: ann.Annotation}, nil
}

func (b backendAccessor) ListAPIEventAnnotations(ctx context.Context, modName string, eventID uint) (anns []*modules.Annotation, err error) {
	dbAnnotations, err := b.dbHandler.APIEventsAnnotationsTable().List(ctx, modName, eventID)
	if err != nil {
		return nil, err
	}
	for _, ann := range dbAnnotations {
		anns = append(anns, &modules.Annotation{Name: ann.Name, Annotation: ann.Annotation})
	}
	return
}

func (b backendAccessor) CreateAPIEventAnnotations(ctx context.Context, modName string, eventID uint, annotations ...modules.Annotation) error {
	var dbAnns []database.APIEventAnnotation

	for _, a := range annotations {
		dbAnns = append(dbAnns, database.APIEventAnnotation{
			EventID:    eventID,
			ModuleName: modName,
			Name:       a.Name,
			Annotation: a.Annotation,
		})
	}

	return b.dbHandler.APIEventsAnnotationsTable().Create(ctx, dbAnns...)
}

func (b backendAccessor) GetAPIInfoAnnotation(ctx context.Context, modName string, apiID uint, name string) (*modules.Annotation, error) {
	ann, err := b.dbHandler.APIInfoAnnotationsTable().Get(ctx, modName, apiID, name)
	if err != nil {
		return nil, err
	}
	return &modules.Annotation{Name: ann.Name, Annotation: ann.Annotation}, nil
}

func (b backendAccessor) ListAPIInfoAnnotations(ctx context.Context, modName string, apiID uint) (annotations []*modules.Annotation, err error) {
	anns, err := b.dbHandler.APIInfoAnnotationsTable().List(ctx, modName, apiID)
	if err != nil {
		return nil, err
	}
	for _, ann := range anns {
		annotations = append(annotations, &modules.Annotation{
			Name:       ann.ModuleName,
			Annotation: ann.Annotation,
		})
	}
	return
}

func (b backendAccessor) StoreAPIInfoAnnotations(ctx context.Context, modName string, apiID uint, annotations ...modules.Annotation) error {
	return b.dbHandler.APIInfoAnnotationsTable().UpdateOrCreate(ctx)
}

func (b backendAccessor) DeleteAPIInfoAnnotations(ctx context.Context, modName string, apiID uint, name ...string) error {
	return b.dbHandler.APIInfoAnnotationsTable().Delete(ctx, modName, apiID, name...)
}
