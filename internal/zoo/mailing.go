package zoo

import (
	"bytes"
	"html/template"
	"zoo-mailer/internal/grouper"
	"zoo-mailer/internal/mailer"
)

func emailAvailabilities(mailer mailer.Mailer, sendlist *mailer.Sendlist, template *template.Template, availabilities []*grouper.AvailabilityGroup) error {

	templateBytes := new(bytes.Buffer)

	templateExecutionError := template.Execute(templateBytes, availabilities)

	if templateExecutionError != nil {

		return templateExecutionError

	}

	emailBody := templateBytes.String()

	sendError := mailer.Send(emailBody, sendlist)

	if sendError != nil {

		return sendError

	}

	return nil

}
