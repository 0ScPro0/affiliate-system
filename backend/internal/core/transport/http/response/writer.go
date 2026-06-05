package core_http_response

import "net/http"

var (
	StatusCodeUnitialized = -1
)

type ResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{
		ResponseWriter: w,
		statusCode: StatusCodeUnitialized,
	}
}

func (rw *ResponseWriter) WriteHeader(statusCode int) {
	rw.ResponseWriter.WriteHeader(statusCode)
	rw.statusCode = statusCode
}

func (rw *ResponseWriter) Write(data []byte) (int, error) {
	if rw.statusCode == StatusCodeUnitialized {
		rw.WriteHeader(http.StatusOK)
	}
	return rw.ResponseWriter.Write(data)
}

func (rw *ResponseWriter) StatusCode() int {
	if rw.statusCode == StatusCodeUnitialized {
		return http.StatusOK
	}
	return rw.statusCode
}