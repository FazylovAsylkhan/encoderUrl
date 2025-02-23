package logger

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type ServerFormatter struct {
	Timestamp string
	Level     logrus.Level
	address    string
	baseUrl      string
	Message   string
}

func (f *ServerFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := entry.Time.Format("2006-01-02T15:04:05.999-07:00")
	return []byte(fmt.Sprintf("%s %s address=%s baseUrl=%s %s\n", 
		timestamp, 
		entry.Level, 
		entry.Data["address"], 
		entry.Data["baseUrl"], 
		entry.Message, 
	)), nil
}

func StartingServer(log *logrus.Logger, address string, baseUrl string) {
	log.WithFields(logrus.Fields{
		"address": address,
		"baseUrl":   baseUrl,
	}).Info("Server started")
}