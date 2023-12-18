package parser

import "fmt"

type AvailabilityParseProblem int

const (
	ErrWrongLineLen AvailabilityParseProblem = iota
	ErrWrongYearFormat
	ErrWrongMonthFormat
	ErrWrongDayFormat
)

func (problem AvailabilityParseProblem) String() string {

	switch problem {
	case ErrWrongLineLen:
		return "WrongLineLength"
	case ErrWrongYearFormat:
		return "WrongYearFormat"
	case ErrWrongMonthFormat:
		return "WrongMonthFormat"
	case ErrWrongDayFormat:
		return "WrongDayFormat"
	default:
		return "Unknown!"
	}

}

type AvailabilityParserError struct {
	LineWithError string
	Problem       AvailabilityParseProblem
}

func (err AvailabilityParserError) Error() string {
	return fmt.Sprintf("Line '%s' has problem: '%s'", err.LineWithError, err.Problem)
}

func NewParseError(lineWithError string, problem AvailabilityParseProblem) AvailabilityParserError {
	return AvailabilityParserError{lineWithError, problem}
}
