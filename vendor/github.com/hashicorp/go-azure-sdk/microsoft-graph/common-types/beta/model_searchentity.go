package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = SearchEntity{}

type SearchEntity struct {
	// Administrative answer in Microsoft Search results to define common acronyms in an organization.
	Acronyms *[]SearchAcronym `json:"acronyms,omitempty"`

	// Administrative answer in Microsoft Search results for common search queries in an organization.
	Bookmarks *[]SearchBookmark `json:"bookmarks,omitempty"`

	// Administrative answer in Microsoft Search results that provide answers for specific search keywords in an
	// organization.
	Qnas *[]SearchQna `json:"qnas,omitempty"`

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

func (s SearchEntity) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SearchEntity{}

func (s SearchEntity) MarshalJSON() ([]byte, error) {
	type wrapper SearchEntity
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SearchEntity: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SearchEntity: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.searchEntity"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SearchEntity: %+v", err)
	}

	return encoded, nil
}