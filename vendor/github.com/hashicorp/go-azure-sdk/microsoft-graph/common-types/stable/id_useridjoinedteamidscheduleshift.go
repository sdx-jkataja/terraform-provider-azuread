package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdJoinedTeamIdScheduleShiftId{}

// UserIdJoinedTeamIdScheduleShiftId is a struct representing the Resource ID for a User Id Joined Team Id Schedule Shift
type UserIdJoinedTeamIdScheduleShiftId struct {
	UserId  string
	TeamId  string
	ShiftId string
}

// NewUserIdJoinedTeamIdScheduleShiftID returns a new UserIdJoinedTeamIdScheduleShiftId struct
func NewUserIdJoinedTeamIdScheduleShiftID(userId string, teamId string, shiftId string) UserIdJoinedTeamIdScheduleShiftId {
	return UserIdJoinedTeamIdScheduleShiftId{
		UserId:  userId,
		TeamId:  teamId,
		ShiftId: shiftId,
	}
}

// ParseUserIdJoinedTeamIdScheduleShiftID parses 'input' into a UserIdJoinedTeamIdScheduleShiftId
func ParseUserIdJoinedTeamIdScheduleShiftID(input string) (*UserIdJoinedTeamIdScheduleShiftId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdScheduleShiftId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdScheduleShiftId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdJoinedTeamIdScheduleShiftIDInsensitively parses 'input' case-insensitively into a UserIdJoinedTeamIdScheduleShiftId
// note: this method should only be used for API response data and not user input
func ParseUserIdJoinedTeamIdScheduleShiftIDInsensitively(input string) (*UserIdJoinedTeamIdScheduleShiftId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdScheduleShiftId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdScheduleShiftId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdJoinedTeamIdScheduleShiftId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.TeamId, ok = input.Parsed["teamId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamId", input)
	}

	if id.ShiftId, ok = input.Parsed["shiftId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "shiftId", input)
	}

	return nil
}

// ValidateUserIdJoinedTeamIdScheduleShiftID checks that 'input' can be parsed as a User Id Joined Team Id Schedule Shift ID
func ValidateUserIdJoinedTeamIdScheduleShiftID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdJoinedTeamIdScheduleShiftID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Joined Team Id Schedule Shift ID
func (id UserIdJoinedTeamIdScheduleShiftId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/schedule/shifts/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.ShiftId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Joined Team Id Schedule Shift ID
func (id UserIdJoinedTeamIdScheduleShiftId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("schedule", "schedule", "schedule"),
		resourceids.StaticSegment("shifts", "shifts", "shifts"),
		resourceids.UserSpecifiedSegment("shiftId", "shiftId"),
	}
}

// String returns a human-readable description of this User Id Joined Team Id Schedule Shift ID
func (id UserIdJoinedTeamIdScheduleShiftId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Shift: %q", id.ShiftId),
	}
	return fmt.Sprintf("User Id Joined Team Id Schedule Shift (%s)", strings.Join(components, "\n"))
}
