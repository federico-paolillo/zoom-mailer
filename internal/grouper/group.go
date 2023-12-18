package grouper

import (
	"time"
	"zoo-mailer/internal/parser"
)

type AvailabilityGroup struct {
	Month          time.Month
	Availabilities []*parser.Availability
}

func NewAvailabilityGroup(month time.Month) *AvailabilityGroup {
	return &AvailabilityGroup{month, nil}
}

func (group *AvailabilityGroup) Add(availability *parser.Availability) {
	group.Availabilities = append(group.Availabilities, availability)
}
