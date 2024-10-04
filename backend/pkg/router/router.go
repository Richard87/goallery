package router

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/Richard87/goallery/api"
	"github.com/Richard87/goallery/pkg/config"
	"github.com/Richard87/goallery/pkg/router/middleware"
	"github.com/Richard87/goallery/pkg/swagger"
	"github.com/rs/zerolog/log"
	"github.com/urfave/negroni"
)

func NewRouter(server api.StrictServerInterface) http.Handler {
	r := http.NewServeMux()
	r.Handle("/swagger/", http.StripPrefix("/swagger", swagger.Handler()))

	si := api.NewStrictHandler(server, []api.StrictMiddlewareFunc{})
	api.HandlerFromMuxWithBaseURL(si, r, "/api/v1")

	n := negroni.New(
		negroni.NewRecovery(),
		middleware.CreateLoggingMiddleware(),
	)
	n.UseHandler(r)

	return n
}

func RunServer(ctx context.Context, handler http.Handler, cfg *config.AppConfig) error {
	addr := fmt.Sprintf(":%d", cfg.Port)
	// And we serve HTTP until the world ends.
	s := &http.Server{
		Handler: handler,
		Addr:    addr,
	}
	go func() {
		log.Ctx(ctx).Info().Msgf("Starting server on http://localhost:%d/swagger/", cfg.Port)
		log.Ctx(ctx).Info().Msgf("Starting server on http://0.0.0.0:%d/swagger/", cfg.Port)

		err := s.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Ctx(ctx).Fatal().Msg(err.Error())
		}
	}()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.WithoutCancel(ctx), 10*time.Second)
	defer cancel()

	return s.Shutdown(shutdownCtx)
}
