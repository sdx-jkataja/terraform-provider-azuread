package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDriveIdRootListItemVersionId{}

// UserIdDriveIdRootListItemVersionId is a struct representing the Resource ID for a User Id Drive Id Root List Item Version
type UserIdDriveIdRootListItemVersionId struct {
	UserId            string
	DriveId           string
	ListItemVersionId string
}

// NewUserIdDriveIdRootListItemVersionID returns a new UserIdDriveIdRootListItemVersionId struct
func NewUserIdDriveIdRootListItemVersionID(userId string, driveId string, listItemVersionId string) UserIdDriveIdRootListItemVersionId {
	return UserIdDriveIdRootListItemVersionId{
		UserId:            userId,
		DriveId:           driveId,
		ListItemVersionId: listItemVersionId,
	}
}

// ParseUserIdDriveIdRootListItemVersionID parses 'input' into a UserIdDriveIdRootListItemVersionId
func ParseUserIdDriveIdRootListItemVersionID(input string) (*UserIdDriveIdRootListItemVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdRootListItemVersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdRootListItemVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdDriveIdRootListItemVersionIDInsensitively parses 'input' case-insensitively into a UserIdDriveIdRootListItemVersionId
// note: this method should only be used for API response data and not user input
func ParseUserIdDriveIdRootListItemVersionIDInsensitively(input string) (*UserIdDriveIdRootListItemVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdRootListItemVersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdRootListItemVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdDriveIdRootListItemVersionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.ListItemVersionId, ok = input.Parsed["listItemVersionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "listItemVersionId", input)
	}

	return nil
}

// ValidateUserIdDriveIdRootListItemVersionID checks that 'input' can be parsed as a User Id Drive Id Root List Item Version ID
func ValidateUserIdDriveIdRootListItemVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdDriveIdRootListItemVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Drive Id Root List Item Version ID
func (id UserIdDriveIdRootListItemVersionId) ID() string {
	fmtString := "/users/%s/drives/%s/root/listItem/versions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DriveId, id.ListItemVersionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Drive Id Root List Item Version ID
func (id UserIdDriveIdRootListItemVersionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("root", "root", "root"),
		resourceids.StaticSegment("listItem", "listItem", "listItem"),
		resourceids.StaticSegment("versions", "versions", "versions"),
		resourceids.UserSpecifiedSegment("listItemVersionId", "listItemVersionId"),
	}
}

// String returns a human-readable description of this User Id Drive Id Root List Item Version ID
func (id UserIdDriveIdRootListItemVersionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("List Item Version: %q", id.ListItemVersionId),
	}
	return fmt.Sprintf("User Id Drive Id Root List Item Version (%s)", strings.Join(components, "\n"))
}
