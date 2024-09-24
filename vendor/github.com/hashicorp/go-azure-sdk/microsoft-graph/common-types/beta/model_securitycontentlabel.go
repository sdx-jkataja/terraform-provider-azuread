package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityContentLabel struct {
	AssignmentMethod *SecurityAssignmentMethod `json:"assignmentMethod,omitempty"`
	CreatedDateTime  nullable.Type[string]     `json:"createdDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	SensitivityLabelId nullable.Type[string] `json:"sensitivityLabelId,omitempty"`
}