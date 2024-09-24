package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecurityResponseAction = SecuritySoftDeleteResponseAction{}

type SecuritySoftDeleteResponseAction struct {
	Identifier *SecurityEmailEntityIdentifier `json:"identifier,omitempty"`

	// Fields inherited from SecurityResponseAction

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s SecuritySoftDeleteResponseAction) SecurityResponseAction() BaseSecurityResponseActionImpl {
	return BaseSecurityResponseActionImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecuritySoftDeleteResponseAction{}

func (s SecuritySoftDeleteResponseAction) MarshalJSON() ([]byte, error) {
	type wrapper SecuritySoftDeleteResponseAction
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecuritySoftDeleteResponseAction: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecuritySoftDeleteResponseAction: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.softDeleteResponseAction"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecuritySoftDeleteResponseAction: %+v", err)
	}

	return encoded, nil
}
