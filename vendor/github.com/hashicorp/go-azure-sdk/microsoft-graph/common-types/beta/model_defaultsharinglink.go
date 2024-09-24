package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultSharingLink struct {
	// Indicates whether the default link setting for this object is a direct URL rather than a sharing link.
	DefaultToExistingAccess nullable.Type[bool] `json:"defaultToExistingAccess,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The default sharing link role. The possible values are: none, view, edit, manageList, review, restrictedView,
	// submitOnly, unknownFutureValue.
	Role *SharingRole `json:"role,omitempty"`

	// The default sharing link scope. The possible values are: anyone, organization, specificPeople, anonymous, users,
	// unknownFutureValue.
	Scope *SharingScope `json:"scope,omitempty"`
}
