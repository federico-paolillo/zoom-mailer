package zoo

import (
	"html/template"
	"log/slog"
	"zoo-mailer/internal/grouper"
	"zoo-mailer/internal/mailer"
	"zoo-mailer/internal/parser"
)

type StatusCode = int

const (
	Ok    StatusCode = 0
	NotOk StatusCode = 1
)

func Run(zooConfig *ZooAppConfig) StatusCode {

	stdoutMailer := mailer.NewStdoutMailer()

	sendlist, sendlistError := loadSendlistFromFile(zooConfig.emailTemplateFilePath)

	if sendlistError != nil {

		slog.Error("failed to process sendlist", slog.String("file", zooConfig.sendlistFilePath), slog.Any("err", sendlistError))

		return NotOk

	}

	candidateFilePaths, gatherError := gatherAvailabilityFiles(zooConfig.zooDirPath)

	slog.Info("processing zoo dir", slog.String("dir", zooConfig.zooDirPath))

	if gatherError != nil {

		slog.Error("failed to process zoo dir", slog.String("dir", zooConfig.zooDirPath), slog.Any("err", gatherError))

		return NotOk

	}

	emailTemplate, templateError := template.ParseFiles(zooConfig.emailTemplateFilePath)

	if templateError != nil {

		slog.Error("failed to process email template", slog.String("template", zooConfig.emailTemplateFilePath))

		return NotOk

	}

	everyAvailability := make([]*parser.Availability, 0)

	for _, candidateFilePath := range *candidateFilePaths {

		availabilities, processError := processAvailabilityFile(candidateFilePath)

		if processError != nil {

			slog.Warn("failed to process file", slog.String("file", candidateFilePath), slog.Any("err", processError))

			continue

		}

		everyAvailability = append(everyAvailability, availabilities...)

		slog.Info("processed file", slog.String("file", candidateFilePath))

	}

	slog.Info("processed all availabilites", slog.Int("count", len(everyAvailability)))

	availabilityGroups := grouper.GroupByMonth(everyAvailability)

	emailError := emailAvailabilities(stdoutMailer, sendlist, emailTemplate, availabilityGroups)

	if emailError != nil {

		slog.Warn("failed to send email", slog.Any("err", emailError))

	}

	slog.Info("processed zoo dir", slog.String("dir", zooConfig.zooDirPath))

	return Ok

}
