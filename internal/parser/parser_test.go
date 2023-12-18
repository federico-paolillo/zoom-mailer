package parser

import (
	"os"
	"slices"
	"testing"
	"zoo-mailer/internal/lines"
)

func TestParsesFile(t *testing.T) {

	fileToParse, _ := os.Open("testdata/example_1")

	fileAvailabilityScanner := lines.NewFileLineScanner(fileToParse)

	availabilityLines, _ := ParseAvailabilityLines(fileAvailabilityScanner)

	availabilitiesExpected := []*Availability{
		NewAvailability(2023, 01, 01),
		NewAvailability(2023, 01, 02),
		NewAvailability(2023, 01, 03),
		NewAvailability(2023, 01, 04),
		NewAvailability(2023, 01, 05),
	}

	if len(availabilityLines) != 5 {
		t.Fatal("some lines were not parsed")
	}

	// for i := range availabilityLines {
	// 	if *availabilityLines[i] != *availabilitiesExpected[i] {
	// 		t.Fatalf("Availability at index %d was %v instead of %v", i, availabilityLines[i], availabilitiesExpected[i])
	// 	}
	// }

	if !slices.EqualFunc(
		availabilityLines,
		availabilitiesExpected,
		func(a1, a2 *Availability) bool { return *a1 == *a2 },
	) {
		t.Fatal("unexpected availabilites parsed")
	}

}

func TestParseAvailabilityLineRejectsLinesTooShort(t *testing.T) {

	_, err := parseAvailabilityLine("2023")

	if err == nil {
		t.Fatal("did not reject line too short")
	}

	if parseErr, ok := err.(AvailabilityParserError); ok {

		if parseErr.Problem != ErrWrongLineLen {
			t.Fatal("returned wrong problem for line too short")
		}

	} else {
		t.Fatal("did not return correct error type for line too short")
	}
}

func TestParseAvailabilityLineDoesNotRejectLineOfCorrectLength(t *testing.T) {

	_, err := parseAvailabilityLine("20230101")

	if err != nil {
		t.Fatal("did reject line with correct length")
	}

}

func TestParseAvailabilityLineRejectsLineTooLong(t *testing.T) {

	_, err := parseAvailabilityLine("20230101010101")

	if err == nil {
		t.Fatal("did not reject line too long")
	}

	if parseErr, ok := err.(AvailabilityParserError); ok {

		if parseErr.Problem != ErrWrongLineLen {
			t.Fatal("returned wrong problem for line too long")
		}

	} else {
		t.Fatal("did not return correct error type for line too long")
	}

}

func TestParseAvailabilityLineRejectsLineWithInvalidYear(t *testing.T) {

	_, err := parseAvailabilityLine("a0231122")

	if err == nil {
		t.Fatal("did not reject line with invalid year")
	}

	if parseError, ok := err.(AvailabilityParserError); ok {

		if parseError.Problem != ErrWrongYearFormat {
			t.Fatal("returned wrong problem for line with invalid year")
		}

	} else {
		t.Fatal("did not return correct error type for line with invalid year")
	}

}

func TestParseAvailabilityLineRejectsLineWithInvalidMonth(t *testing.T) {

	_, err := parseAvailabilityLine("2023a122")

	if err == nil {
		t.Fatal("did not reject line with invalid month")
	}

	if parseError, ok := err.(AvailabilityParserError); ok {

		if parseError.Problem != ErrWrongMonthFormat {
			t.Fatal("returned wrong problem for line with invalid month")
		}

	} else {
		t.Fatal("did not return correct error type for line with invalid month")
	}

}

func TestParseAvailabilityLineRejectsLineWithInvalidDay(t *testing.T) {

	_, err := parseAvailabilityLine("202311a2")

	if err == nil {
		t.Fatal("did not reject line with invalid day")
	}

	if parseError, ok := err.(AvailabilityParserError); ok {

		if parseError.Problem != ErrWrongDayFormat {
			t.Fatal("returned wrong problem for line with invalid day")
		}

	} else {
		t.Fatal("did not return correct error type for line with invalid day")
	}

}

func TestParseAvailabilityLineParses(t *testing.T) {

	availabilityParsed, err := parseAvailabilityLine("20230110")

	if err != nil {
		t.Fatal("rejected perfectly fine line")
	}

	availabilityExpected := NewAvailability(2023, 01, 10)

	if *availabilityParsed != *availabilityExpected {
		t.Fatalf("availability returned is wrong. expected %s got %s", availabilityExpected, availabilityParsed)
	}

}
