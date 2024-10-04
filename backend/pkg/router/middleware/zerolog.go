package middleware

import (
	"net/http"

	"github.com/felixge/httpsnoop"
	"github.com/rs/zerolog/log"
	"github.com/urfave/negroni"
)

func CreateLoggingMiddleware() negroni.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request, next http.HandlerFunc) {
		metrics := httpsnoop.CaptureMetrics(next, writer, request)
		log.Info().
			Str("operation_id", request.Pattern).
			Dur("duration", metrics.Duration).
			Int("status_code", metrics.Code).
			Int64("response_size", metrics.Written).
			Msg("Handled request")
	}
}
