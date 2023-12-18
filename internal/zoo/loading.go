package zoo

import (
	"os"
	"zoo-mailer/internal/lines"
	"zoo-mailer/internal/mailer"
)

func loadSendlistFromFile(sendlistFilePath string) (*mailer.Sendlist, error) {

	sendlistFileHandle, sendlistOpenErr := os.Open(sendlistFilePath)

	if sendlistOpenErr != nil {

		return nil, sendlistOpenErr

	}

	defer sendlistFileHandle.Close()

	sendlistFileScanner := lines.NewFileLineScanner(sendlistFileHandle)

	sendlist, sendlistParseError := mailer.ParseSendlist(sendlistFileScanner)

	if sendlistParseError != nil {

		return nil, sendlistParseError

	}

	return sendlist, nil

}
