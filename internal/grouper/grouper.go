package grouper

import (
	"time"
	"zoo-mailer/internal/parser"

	"golang.org/x/exp/maps"
)

func GroupByMonth(someAvailabilities []*parser.Availability) []*AvailabilityGroup {

	buckets := make(map[time.Month]*AvailabilityGroup)

	for _, availability := range someAvailabilities {

		if bucket, bucketExists := buckets[availability.Month]; bucketExists {

			bucket.Add(availability)

		} else {

			bucket := NewAvailabilityGroup(availability.Month)

			bucket.Add(availability)

			buckets[availability.Month] = bucket

		}

	}

	allAvailabilities := maps.Values(buckets)

	return allAvailabilities

}
