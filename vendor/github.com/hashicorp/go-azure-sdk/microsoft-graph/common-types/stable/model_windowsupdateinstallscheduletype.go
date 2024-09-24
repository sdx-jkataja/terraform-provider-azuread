package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdateInstallScheduleType interface {
	WindowsUpdateInstallScheduleType() BaseWindowsUpdateInstallScheduleTypeImpl
}

var _ WindowsUpdateInstallScheduleType = BaseWindowsUpdateInstallScheduleTypeImpl{}

type BaseWindowsUpdateInstallScheduleTypeImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseWindowsUpdateInstallScheduleTypeImpl) WindowsUpdateInstallScheduleType() BaseWindowsUpdateInstallScheduleTypeImpl {
	return s
}

var _ WindowsUpdateInstallScheduleType = RawWindowsUpdateInstallScheduleTypeImpl{}

// RawWindowsUpdateInstallScheduleTypeImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawWindowsUpdateInstallScheduleTypeImpl struct {
	windowsUpdateInstallScheduleType BaseWindowsUpdateInstallScheduleTypeImpl
	Type                             string
	Values                           map[string]interface{}
}

func (s RawWindowsUpdateInstallScheduleTypeImpl) WindowsUpdateInstallScheduleType() BaseWindowsUpdateInstallScheduleTypeImpl {
	return s.windowsUpdateInstallScheduleType
}

func UnmarshalWindowsUpdateInstallScheduleTypeImplementation(input []byte) (WindowsUpdateInstallScheduleType, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsUpdateInstallScheduleType into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsUpdateActiveHoursInstall") {
		var out WindowsUpdateActiveHoursInstall
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUpdateActiveHoursInstall: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsUpdateScheduledInstall") {
		var out WindowsUpdateScheduledInstall
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUpdateScheduledInstall: %+v", err)
		}
		return out, nil
	}

	var parent BaseWindowsUpdateInstallScheduleTypeImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseWindowsUpdateInstallScheduleTypeImpl: %+v", err)
	}

	return RawWindowsUpdateInstallScheduleTypeImpl{
		windowsUpdateInstallScheduleType: parent,
		Type:                             value,
		Values:                           temp,
	}, nil

}
