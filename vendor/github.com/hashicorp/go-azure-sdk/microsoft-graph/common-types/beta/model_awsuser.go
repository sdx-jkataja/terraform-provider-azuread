package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AwsIdentity = AwsUser{}

type AwsUser struct {
	// Roles assumed by the user.
	AssumableRoles *[]AwsRole `json:"assumableRoles,omitempty"`

	// Fields inherited from AuthorizationSystemIdentity

	// Navigation to the authorizationSystem object
	AuthorizationSystem *AuthorizationSystem `json:"authorizationSystem,omitempty"`

	// The name of the identity. Read-only. Supports $filter and (eq,contains).
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Unique ID of the identity within the external system. Read-only.
	ExternalId *string `json:"externalId,omitempty"`

	// Represents details of the source of the identity.
	Source AuthorizationSystemIdentitySource `json:"source"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s AwsUser) AwsIdentity() BaseAwsIdentityImpl {
	return BaseAwsIdentityImpl{
		AuthorizationSystem: s.AuthorizationSystem,
		DisplayName:         s.DisplayName,
		ExternalId:          s.ExternalId,
		Source:              s.Source,
		Id:                  s.Id,
		ODataId:             s.ODataId,
		ODataType:           s.ODataType,
	}
}

func (s AwsUser) AuthorizationSystemIdentity() BaseAuthorizationSystemIdentityImpl {
	return BaseAuthorizationSystemIdentityImpl{
		AuthorizationSystem: s.AuthorizationSystem,
		DisplayName:         s.DisplayName,
		ExternalId:          s.ExternalId,
		Source:              s.Source,
		Id:                  s.Id,
		ODataId:             s.ODataId,
		ODataType:           s.ODataType,
	}
}

func (s AwsUser) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AwsUser{}

func (s AwsUser) MarshalJSON() ([]byte, error) {
	type wrapper AwsUser
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AwsUser: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AwsUser: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.awsUser"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AwsUser: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &AwsUser{}

func (s *AwsUser) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AssumableRoles *[]AwsRole            `json:"assumableRoles,omitempty"`
		DisplayName    nullable.Type[string] `json:"displayName,omitempty"`
		ExternalId     *string               `json:"externalId,omitempty"`
		Id             *string               `json:"id,omitempty"`
		ODataId        *string               `json:"@odata.id,omitempty"`
		ODataType      *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AssumableRoles = decoded.AssumableRoles
	s.DisplayName = decoded.DisplayName
	s.ExternalId = decoded.ExternalId
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AwsUser into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["authorizationSystem"]; ok {
		impl, err := UnmarshalAuthorizationSystemImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'AuthorizationSystem' for 'AwsUser': %+v", err)
		}
		s.AuthorizationSystem = &impl
	}

	if v, ok := temp["source"]; ok {
		impl, err := UnmarshalAuthorizationSystemIdentitySourceImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Source' for 'AwsUser': %+v", err)
		}
		s.Source = impl
	}

	return nil
}
