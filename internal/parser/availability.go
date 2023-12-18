package parser

import (
	"fmt"
	"time"
)

type Availability struct {
	Year  int
	Month time.Month
	Day   int
}

func NewAvailability(year int, month time.Month, day int) *Availability {
	return &Availability{year, month, day}
}

func (availability *Availability) String() string {
	return fmt.Sprintf("%04d-%02d-%02d", availability.Year, availability.Month, availability.Day)
}
