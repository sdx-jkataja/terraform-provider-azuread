package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MultiTenantOrganizationState string

const (
	MultiTenantOrganizationState_Active   MultiTenantOrganizationState = "active"
	MultiTenantOrganizationState_Inactive MultiTenantOrganizationState = "inactive"
)

func PossibleValuesForMultiTenantOrganizationState() []string {
	return []string{
		string(MultiTenantOrganizationState_Active),
		string(MultiTenantOrganizationState_Inactive),
	}
}

func (s *MultiTenantOrganizationState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMultiTenantOrganizationState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMultiTenantOrganizationState(input string) (*MultiTenantOrganizationState, error) {
	vals := map[string]MultiTenantOrganizationState{
		"active":   MultiTenantOrganizationState_Active,
		"inactive": MultiTenantOrganizationState_Inactive,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MultiTenantOrganizationState(input)
	return &out, nil
}
