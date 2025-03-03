package logger

import (
	"fmt"
	"io"
	"os"

	"github.com/sirupsen/logrus"
)


func New() *logrus.Logger {
    var log = logrus.New()
	log.SetLevel(logrus.InfoLevel)
	log.SetReportCaller(true)
	
    file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err == nil {
        log.SetOutput(io.MultiWriter(file, os.Stderr))
    } else {
        log.Info("Не удалось открыть файл логов, используется стандартный stderr")
    }

	return log
}

type GeneralFormatter struct {
	Timestamp string
	Level     logrus.Level
	Message   string
}

func (f *GeneralFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := entry.Time.Format("2006-01-02T15:04:05.999-07:00")
	return []byte(fmt.Sprintf("%s %s %s\n", 
		timestamp, 
		entry.Level, 
		entry.Message, 
	)), nil
}
