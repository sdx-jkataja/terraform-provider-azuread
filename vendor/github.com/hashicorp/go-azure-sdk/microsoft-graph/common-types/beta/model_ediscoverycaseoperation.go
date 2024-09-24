package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryCaseOperation interface {
	Entity
	EdiscoveryCaseOperation() BaseEdiscoveryCaseOperationImpl
}

var _ EdiscoveryCaseOperation = BaseEdiscoveryCaseOperationImpl{}

type BaseEdiscoveryCaseOperationImpl struct {
	// The type of action the operation represents. Possible values are:
	// addToReviewSet,applyTags,contentExport,convertToPdf,estimateStatistics, purgeData
	Action *EdiscoveryCaseAction `json:"action,omitempty"`

	// The date and time the operation was completed.
	CompletedDateTime nullable.Type[string] `json:"completedDateTime,omitempty"`

	// The user that created the operation.
	CreatedBy IdentitySet `json:"createdBy"`

	// The date and time the operation was created.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The progress of the operation.
	PercentProgress nullable.Type[int64] `json:"percentProgress,omitempty"`

	// Contains success and failure-specific result information.
	ResultInfo *ResultInfo `json:"resultInfo,omitempty"`

	// The status of the case operation. Possible values are: notStarted, submissionFailed, running, succeeded,
	// partiallySucceeded, failed.
	Status *EdiscoveryCaseOperationStatus `json:"status,omitempty"`

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

func (s BaseEdiscoveryCaseOperationImpl) EdiscoveryCaseOperation() BaseEdiscoveryCaseOperationImpl {
	return s
}

func (s BaseEdiscoveryCaseOperationImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ EdiscoveryCaseOperation = RawEdiscoveryCaseOperationImpl{}

// RawEdiscoveryCaseOperationImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawEdiscoveryCaseOperationImpl struct {
	ediscoveryCaseOperation BaseEdiscoveryCaseOperationImpl
	Type                    string
	Values                  map[string]interface{}
}

func (s RawEdiscoveryCaseOperationImpl) EdiscoveryCaseOperation() BaseEdiscoveryCaseOperationImpl {
	return s.ediscoveryCaseOperation
}

func (s RawEdiscoveryCaseOperationImpl) Entity() BaseEntityImpl {
	return s.ediscoveryCaseOperation.Entity()
}

var _ json.Marshaler = BaseEdiscoveryCaseOperationImpl{}

func (s BaseEdiscoveryCaseOperationImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseEdiscoveryCaseOperationImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseEdiscoveryCaseOperationImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseEdiscoveryCaseOperationImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.ediscovery.caseOperation"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseEdiscoveryCaseOperationImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseEdiscoveryCaseOperationImpl{}

func (s *BaseEdiscoveryCaseOperationImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Action            *EdiscoveryCaseAction          `json:"action,omitempty"`
		CompletedDateTime nullable.Type[string]          `json:"completedDateTime,omitempty"`
		CreatedDateTime   nullable.Type[string]          `json:"createdDateTime,omitempty"`
		PercentProgress   nullable.Type[int64]           `json:"percentProgress,omitempty"`
		ResultInfo        *ResultInfo                    `json:"resultInfo,omitempty"`
		Status            *EdiscoveryCaseOperationStatus `json:"status,omitempty"`
		Id                *string                        `json:"id,omitempty"`
		ODataId           *string                        `json:"@odata.id,omitempty"`
		ODataType         *string                        `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Action = decoded.Action
	s.CompletedDateTime = decoded.CompletedDateTime
	s.CreatedDateTime = decoded.CreatedDateTime
	s.PercentProgress = decoded.PercentProgress
	s.ResultInfo = decoded.ResultInfo
	s.Status = decoded.Status
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseEdiscoveryCaseOperationImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'BaseEdiscoveryCaseOperationImpl': %+v", err)
		}
		s.CreatedBy = impl
	}

	return nil
}

func UnmarshalEdiscoveryCaseOperationImplementation(input []byte) (EdiscoveryCaseOperation, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling EdiscoveryCaseOperation into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.ediscovery.addToReviewSetOperation") {
		var out EdiscoveryAddToReviewSetOperation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EdiscoveryAddToReviewSetOperation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.ediscovery.caseExportOperation") {
		var out EdiscoveryCaseExportOperation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EdiscoveryCaseExportOperation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.ediscovery.caseHoldOperation") {
		var out EdiscoveryCaseHoldOperation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EdiscoveryCaseHoldOperation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.ediscovery.caseIndexOperation") {
		var out EdiscoveryCaseIndexOperation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EdiscoveryCaseIndexOperation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.ediscovery.estimateStatisticsOperation") {
		var out EdiscoveryEstimateStatisticsOperation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EdiscoveryEstimateStatisticsOperation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.ediscovery.purgeDataOperation") {
		var out EdiscoveryPurgeDataOperation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EdiscoveryPurgeDataOperation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.ediscovery.tagOperation") {
		var out EdiscoveryTagOperation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EdiscoveryTagOperation: %+v", err)
		}
		return out, nil
	}

	var parent BaseEdiscoveryCaseOperationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseEdiscoveryCaseOperationImpl: %+v", err)
	}

	return RawEdiscoveryCaseOperationImpl{
		ediscoveryCaseOperation: parent,
		Type:                    value,
		Values:                  temp,
	}, nil

}
