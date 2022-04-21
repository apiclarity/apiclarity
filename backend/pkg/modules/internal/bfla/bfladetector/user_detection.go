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

package bfladetector

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

func GetUserID(headers map[string]string) (*DetectedUser, error) {
	if xcustomerID, ok := headers["x-customer-id"]; ok {
		return &DetectedUser{Source: DetectedUserSourceXConsumerIDHeader, ID: xcustomerID}, nil
	}
	authz, ok := headers["authorization"]
	if !ok {
		return nil, nil
	}
	if strings.HasPrefix(authz, "Basic ") {
		basic := strings.TrimPrefix(authz, "Basic ")
		usernameAndPassword, err := base64.StdEncoding.DecodeString(basic)
		if err != nil {
			return nil, fmt.Errorf("cannot decode basic authz header: %w", err)
		}
		usernameAndPasswordParts := strings.Split(string(usernameAndPassword), ":")

		// nolint:gomnd
		if len(usernameAndPasswordParts) < 2 {
			return nil, errors.New("broken basic auth header")
		}
		return &DetectedUser{Source: DetectedUserSourceBasic, ID: usernameAndPasswordParts[0]}, nil
	}
	if strings.HasPrefix(authz, "Bearer ") {
		bearer := strings.TrimPrefix(authz, "Bearer ")
		bearerParts := strings.Split(bearer, ".")

		// nolint:gomnd
		if len(bearerParts) == 3 { // is JWT
			s, err := base64.URLEncoding.DecodeString(bearerParts[1])
			if err != nil {
				return nil, fmt.Errorf("unable to decode bearer token: %w", err)
			}
			data := struct {
				Subject string `json:"sub"`
			}{}
			if err := json.Unmarshal(s, &data); err != nil {
				return nil, fmt.Errorf("unable to unmarshal json jwt body: %w", err)
			}
			return &DetectedUser{Source: DetectedUserSourceJWT, ID: data.Subject}, nil
		}

		return nil, ErrUnsupportedAuthScheme
	}
	return nil, ErrUnsupportedAuthScheme
}
