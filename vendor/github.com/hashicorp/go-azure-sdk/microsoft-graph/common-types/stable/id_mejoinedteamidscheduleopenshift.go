package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeJoinedTeamIdScheduleOpenShiftId{}

// MeJoinedTeamIdScheduleOpenShiftId is a struct representing the Resource ID for a Me Joined Team Id Schedule Open Shift
type MeJoinedTeamIdScheduleOpenShiftId struct {
	TeamId      string
	OpenShiftId string
}

// NewMeJoinedTeamIdScheduleOpenShiftID returns a new MeJoinedTeamIdScheduleOpenShiftId struct
func NewMeJoinedTeamIdScheduleOpenShiftID(teamId string, openShiftId string) MeJoinedTeamIdScheduleOpenShiftId {
	return MeJoinedTeamIdScheduleOpenShiftId{
		TeamId:      teamId,
		OpenShiftId: openShiftId,
	}
}

// ParseMeJoinedTeamIdScheduleOpenShiftID parses 'input' into a MeJoinedTeamIdScheduleOpenShiftId
func ParseMeJoinedTeamIdScheduleOpenShiftID(input string) (*MeJoinedTeamIdScheduleOpenShiftId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdScheduleOpenShiftId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdScheduleOpenShiftId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeJoinedTeamIdScheduleOpenShiftIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamIdScheduleOpenShiftId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamIdScheduleOpenShiftIDInsensitively(input string) (*MeJoinedTeamIdScheduleOpenShiftId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdScheduleOpenShiftId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdScheduleOpenShiftId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeJoinedTeamIdScheduleOpenShiftId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.TeamId, ok = input.Parsed["teamId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamId", input)
	}

	if id.OpenShiftId, ok = input.Parsed["openShiftId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "openShiftId", input)
	}

	return nil
}

// ValidateMeJoinedTeamIdScheduleOpenShiftID checks that 'input' can be parsed as a Me Joined Team Id Schedule Open Shift ID
func ValidateMeJoinedTeamIdScheduleOpenShiftID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamIdScheduleOpenShiftID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Id Schedule Open Shift ID
func (id MeJoinedTeamIdScheduleOpenShiftId) ID() string {
	fmtString := "/me/joinedTeams/%s/schedule/openShifts/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.OpenShiftId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Id Schedule Open Shift ID
func (id MeJoinedTeamIdScheduleOpenShiftId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("schedule", "schedule", "schedule"),
		resourceids.StaticSegment("openShifts", "openShifts", "openShifts"),
		resourceids.UserSpecifiedSegment("openShiftId", "openShiftId"),
	}
}

// String returns a human-readable description of this Me Joined Team Id Schedule Open Shift ID
func (id MeJoinedTeamIdScheduleOpenShiftId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Open Shift: %q", id.OpenShiftId),
	}
	return fmt.Sprintf("Me Joined Team Id Schedule Open Shift (%s)", strings.Join(components, "\n"))
}
