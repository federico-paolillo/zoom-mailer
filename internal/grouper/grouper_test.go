package grouper

import (
	"slices"
	"testing"
	"time"
	"zoo-mailer/internal/parser"
)

func TestGroupsByMonth(t *testing.T) {

	someAvailabilities := []*parser.Availability{
		parser.NewAvailability(2022, time.January, 01),
		parser.NewAvailability(2022, time.February, 01),
		parser.NewAvailability(2022, time.February, 03),
		parser.NewAvailability(2022, time.January, 10),
		parser.NewAvailability(2022, time.February, 13),
	}

	groups := GroupByMonth(someAvailabilities)

	groupLen := len(groups)

	if groupLen != 2 {
		t.Fatal("wrong number of groups parsed")
	}

	januaryIndex := slices.IndexFunc(groups,
		func(a *AvailabilityGroup) bool {
			return a.Month == time.January
		})

	februaryIndex := slices.IndexFunc(groups,
		func(a *AvailabilityGroup) bool {
			return a.Month == time.February
		})

	if len(groups[januaryIndex].Availabilities) != 2 {
		t.Error("not all january availabilities were grouped under january")
	}

	if len(groups[februaryIndex].Availabilities) != 3 {
		t.Error("not all february availabilities were grouped under february")
	}

	// Equality check assumes grouping keeps order of entry

	if !slices.EqualFunc(
		groups[januaryIndex].Availabilities,
		[]*parser.Availability{
			parser.NewAvailability(2022, time.January, 01),
			parser.NewAvailability(2022, time.January, 10),
		},
		func(a1, a2 *parser.Availability) bool { return *a1 == *a2 }) {

		t.Error("some january availabilites are missing")

	}

	if !slices.EqualFunc(
		groups[februaryIndex].Availabilities,
		[]*parser.Availability{
			parser.NewAvailability(2022, time.February, 01),
			parser.NewAvailability(2022, time.February, 03),
			parser.NewAvailability(2022, time.February, 13),
		},
		func(a1, a2 *parser.Availability) bool { return *a1 == *a2 }) {

		t.Error("some february availabilites are missing")

	}

}
