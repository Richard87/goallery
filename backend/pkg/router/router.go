package router

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/Richard87/goallery/pkg/router/middleware"
	"github.com/rs/zerolog/log"
	"github.com/urfave/negroni"
)

type RouteMapper func(router *http.ServeMux)

type Router struct {
	handler http.Handler
}

func NewRouter(mappers ...RouteMapper) *Router {
	r := http.NewServeMux()
	for _, mapper := range mappers {
		mapper(r)
	}

	handler := negroni.New(
		negroni.NewRecovery(),
		middleware.CreateLoggingMiddleware(),
		negroni.Wrap(r),
	)
	return &Router{handler: handler}
}

func (r *Router) Serve(ctx context.Context, port int) error {
	s := &http.Server{
		Handler: r.handler,
		Addr:    fmt.Sprintf(":%d", port),
	}
	go func() {
		log.Ctx(ctx).Info().Msgf("Starting server on http://localhost:%d/swagger/", port)

		if err := s.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Ctx(ctx).Fatal().Msg(err.Error())
		}
	}()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.WithoutCancel(ctx), 10*time.Second)
	defer cancel()

	return s.Shutdown(shutdownCtx)
}
