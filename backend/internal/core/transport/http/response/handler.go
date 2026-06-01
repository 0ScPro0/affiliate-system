package core_http_response

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	core_error "github.com/0ScPro0/affiliate-system/internal/core/errors"
	core_errors "github.com/0ScPro0/affiliate-system/internal/core/errors"
	"github.com/0ScPro0/affiliate-system/internal/core/logger"
	"go.uber.org/zap"
)

type HTTPResponseHandler struct {
	log *logger.Logger
	rw http.ResponseWriter
}

func NewHTTPResponseHandler(
	log *logger.Logger,
	rw http.ResponseWriter,
) *HTTPResponseHandler {
	return &HTTPResponseHandler{
		log: log,
		rw: rw,
	}
}

func (h *HTTPResponseHandler) JSONResponse(response any, statusCode int) {
	h.rw.WriteHeader(statusCode)

	if err := json.NewEncoder(h.rw).Encode(response); err != nil {
		h.log.Error("Write HTTP response: %w", zap.Error(err))
	}
}

func (h *HTTPResponseHandler) ErrorResponse(err error, msg string) {
	statusCode, logFunc := h.compareStatusError(err)

	logFunc(msg, zap.Error(err))
	h.errorResponse(statusCode, err, msg)
}

func (h *HTTPResponseHandler) PanicResponse(p any, msg string) {
	statusCode := http.StatusInternalServerError
	err := fmt.Errorf("Unexpected panic: %v", p)

	h.log.Error(msg, zap.Error(err))
	
	h.errorResponse(statusCode, err, msg)
}

func (h *HTTPResponseHandler) errorResponse(
	statusCode int,
	err error,
	msg string,
) {
	h.rw.WriteHeader(statusCode)

	response := map[string]string {
		"message": msg,
		"error": err.Error(),
	}

	h.JSONResponse(
		response,
		statusCode,
	)
}

func (h *HTTPResponseHandler) compareStatusError(err error) (int, func(string, ...zap.Field)) {
	var statusCode int
	var logFunc func(string, ...zap.Field)

	switch {
	case errors.Is(err, core_errors.ErrNotFound):
		statusCode = http.StatusNotFound
		logFunc = h.log.Debug
	
	case errors.Is(err, core_error.ErrInvalidArgument):
		statusCode = http.StatusBadRequest
		logFunc = h.log.Warn

	case errors.Is(err, core_error.ErrConflict):
		statusCode = http.StatusConflict
		logFunc = h.log.Warn
	
	default:
		statusCode = http.StatusInternalServerError
		logFunc = h.log.Error
	}

	return statusCode, logFunc
}