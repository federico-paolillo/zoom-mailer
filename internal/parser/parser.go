package parser

import (
	"io"
	"strconv"
	"time"
	"unicode/utf8"
	"zoo-mailer/internal/lines"
)

type AvailabilityScanner interface {
	Scan() (string, error)
}

const expectedLineLen = 4 + 2 + 2 // YYYYMMDD

func parseAvailabilityLine(lineToParse string) (*Availability, error) {

	lineLen := utf8.RuneCountInString(lineToParse)

	if lineLen == 0 {
		return nil, NewParseError(lineToParse, ErrWrongLineLen)
	}

	lineRunes := []rune(lineToParse)

	if len(lineRunes) != expectedLineLen {
		return nil, NewParseError(lineToParse, ErrWrongLineLen)
	}

	yearPartString := string(lineRunes[0:4])
	monthPartString := string(lineRunes[4:6])
	dayPartString := string(lineRunes[6:8])

	year, err := strconv.Atoi(yearPartString)

	if err != nil {
		return nil, NewParseError(lineToParse, ErrWrongYearFormat)
	}

	month, err := strconv.Atoi(monthPartString)

	if err != nil {
		return nil, NewParseError(lineToParse, ErrWrongMonthFormat)
	}

	day, err := strconv.Atoi(dayPartString)

	if err != nil {
		return nil, NewParseError(lineToParse, ErrWrongDayFormat)
	}

	return NewAvailability(year, time.Month(month), day), nil

}

func ParseAvailabilityLines(availabilityScanner lines.LineScanner) ([]*Availability, error) {

	// The longest month is 31 days. We start from that and then grow

	result := make([]*Availability, 0, 31)

	for {

		lineToParse, err := availabilityScanner.Scan()

		if err == io.EOF {

			break

		}

		if err != nil {

			return nil, err

		}

		availabilityParsed, err := parseAvailabilityLine(lineToParse)

		if err != nil {

			return nil, err

		}

		result = append(result, availabilityParsed)

	}

	return result, nil

}
