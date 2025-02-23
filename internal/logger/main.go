package logger

import (
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
