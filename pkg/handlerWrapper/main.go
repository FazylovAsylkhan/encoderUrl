package handlerWrapper

import (
	"bytes"
	"io"
	"net/http"
	"strings"
)

type HandlerWrapper struct {
	http.ResponseWriter
	StatusCode int
	Headers    http.Header
	ReqBody       string
	ResBody       bytes.Buffer
	RequestHeaders map[string]string
	ResponseHeaders map[string]string
}

func (rw *HandlerWrapper) Write(b []byte) (int, error) {
	rw.ResponseHeaders = make(map[string]string)
	for key, values := range rw.Header() {
		rw.ResponseHeaders[key] = strings.Join(values, ", ")
	}
	rw.ResBody.Write(b)
	return rw.ResponseWriter.Write(b)
}

func (rw *HandlerWrapper) WriteHeader(statusCode int) {
	rw.StatusCode = statusCode
	rw.ResponseHeaders = make(map[string]string)
	for key, values := range rw.ResponseWriter.Header() {
		rw.ResponseHeaders[key] = strings.Join(values, ", ")
	}
	rw.ResponseWriter.WriteHeader(statusCode)
}

func New(w http.ResponseWriter, r *http.Request) *HandlerWrapper {
	rw := &HandlerWrapper{
		ResponseWriter: w,
		StatusCode: http.StatusOK,
		RequestHeaders:  make(map[string]string),
		ResponseHeaders: make(map[string]string),
	}

	for key, values := range r.Header {
		rw.RequestHeaders[key] = strings.Join(values, ", ")
	}

	if r.Body != nil {
		bodyBytes, _ := io.ReadAll(r.Body)
		rw.ReqBody = string(bodyBytes)
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	}


	return rw
}

func (rw *HandlerWrapper) HeadersString(headers map[string]string) string {
	var str string
	for key, value := range headers {
		switch key {
		case "User-Agent", "Content-Type", "Content-Length", "Location":
			str += key + ": " + value + ", "
		}
	}
	return str
}