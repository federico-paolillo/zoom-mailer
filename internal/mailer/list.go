package mailer

import (
	"io"
	"strings"
	"zoo-mailer/internal/lines"
)

type Sendlist struct {
	members []string
}

func ParseSendlist(sendlistScanner lines.LineScanner) (*Sendlist, error) {

	members := make([]string, 0, 3) // We init. with a size of 3 because I know I won't have more than 3 recipients

	for {

		memberLine, err := sendlistScanner.Scan()

		if err == io.EOF {

			break

		}

		if err != nil {

			return nil, err

		}

		memberLine = strings.TrimSpace(memberLine)

		if memberLine == "" {

			continue

		}

		members = append(members, memberLine)

	}

	return &Sendlist{members}, nil

}
