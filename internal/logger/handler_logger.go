package logger

import (
	"fmt"
	"net/http"
	"time"

	"github.com/FazylovAsylkhan/encoderUrl/pkg/handlerWrapper"
	"github.com/sirupsen/logrus"
)

type HandlerFormatter struct {}

func MiddlewareHandler(log *logrus.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rw := handlerWrapper.New(w, r)
		next.ServeHTTP(rw, r)

		request := fmt.Sprintf("Headers: %v Body: %v", rw.HeadersString(rw.RequestHeaders), rw.ReqBody)
		response := fmt.Sprintf("Headers: %v Body: %v", rw.HeadersString(rw.ResponseHeaders), rw.ResBody.String())
		duration := time.Since(start)

		if rw.StatusCode >= 400 {
			log.WithFields(logrus.Fields{
				"service":  "encoderUrl",
				"method":   r.Method,
				"statusCode":   rw.StatusCode,
				"uri":      r.RequestURI,
				"duration": duration,
				"request":  request,
				"response": response,
	
			}).Error("Request failed")
		} else {
			log.WithFields(logrus.Fields{
				"service":  "encoderUrl",
				"method":   r.Method,
				"statusCode":   rw.StatusCode,
				"uri":      r.RequestURI,
				"duration": duration,
				"request":  request,
				"response": response,
	
			}).Info("Request handled")
		}
	})
}

func (f *HandlerFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := entry.Time.Format("2006-01-02T15:04:05.999-07:00")
	return []byte(fmt.Sprintf("%s %s %s %s %d uri=%s request=%s response=%s duration=%v %s\n", 
	timestamp, 
	entry.Level, 
	entry.Data["service"], 
	entry.Data["method"], 
	entry.Data["statusCode"], 
	entry.Data["uri"], 
	entry.Data["request"],
	entry.Data["response"], 
	entry.Data["duration"],
	entry.Message, 
	)), nil
}