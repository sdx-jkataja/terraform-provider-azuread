package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RelatedContact struct {
	// Indicates whether the user has been consented to access student data.
	AccessConsent nullable.Type[bool] `json:"accessConsent,omitempty"`

	// Name of the contact. Required.
	DisplayName string `json:"displayName"`

	// Primary email address of the contact. Required.
	EmailAddress string `json:"emailAddress"`

	// Mobile phone number of the contact.
	MobilePhone nullable.Type[string] `json:"mobilePhone,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	Relationship *ContactRelationship `json:"relationship,omitempty"`
}
