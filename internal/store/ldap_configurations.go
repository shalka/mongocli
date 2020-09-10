// Copyright 2020 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package store

import (
	"context"
	"fmt"

	"github.com/mongodb/mongocli/internal/config"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

//go:generate mockgen -destination=../mocks/mock_ldap_configurations.go -package=mocks github.com/mongodb/mongocli/internal/store LDAPConfigurationVerifier

type LDAPConfigurationVerifier interface {
	VerifyLDAPConfiguration(string, *atlas.LDAP) (*atlas.LDAPConfiguration, error)
}

// VerifyLDAPConfiguration encapsulates the logic to manage different cloud providers
func (s *Store) VerifyLDAPConfiguration(projectID string, ldap *atlas.LDAP) (*atlas.LDAPConfiguration, error) {
	switch s.service {
	case config.CloudService:
		resp, _, err := s.client.(*atlas.Client).LDAPConfigurations.Verify(context.Background(), projectID, ldap)
		return resp, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}