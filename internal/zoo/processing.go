package zoo

import (
	"os"
	"zoo-mailer/internal/lines"
	"zoo-mailer/internal/parser"
)

func processAvailabilityFile(filePath string) ([]*parser.Availability, error) {

	fileToProcessHandle, openFileError := os.Open(filePath)

	if openFileError != nil {

		return nil, openFileError

	}

	defer fileToProcessHandle.Close()

	availabilityFileScanner := lines.NewFileLineScanner(fileToProcessHandle)

	availabilities, parseError := parser.ParseAvailabilityLines(availabilityFileScanner)

	if parseError != nil {

		return nil, parseError

	}

	return availabilities, nil

}
