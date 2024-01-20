package restapi

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Richard87/goallery/generated/restapi"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var ErrRestApi = fmt.Errorf("restapi err")
var ErrFailedToInitialiseRestApi = fmt.Errorf("%w: failed to initialise restapi", ErrRestApi)

type Config struct {
	restapi.Config
	Port int
}

type OptionFunc func(config *Config) error

func WithAuthApi(api restapi.AuthAPI) OptionFunc {
	return func(config *Config) error {
		config.AuthAPI = api
		return nil
	}
}
func WithImagesApi(api restapi.ImagesAPI) OptionFunc {
	return func(config *Config) error {
		config.ImagesAPI = api
		return nil
	}
}
func WithZerolog(log zerolog.Logger) OptionFunc {
	return func(config *Config) error {
		config.Logger = func(f string, args ...interface{}) {
			log.Info().Str("pkg", "restapi").Msgf(f, args...)
		}
		return nil
	}
}
func WithHttpPort(port int) OptionFunc {
	return func(config *Config) error {
		config.Port = port
		return nil
	}
}

func New(ctx context.Context, options ...OptionFunc) error {
	config := Config{
		Config: restapi.Config{},
		Port:   8000,
	}

	for _, option := range options {
		err := option(&config)
		if err != nil {
			return fmt.Errorf("%w: %w", ErrFailedToInitialiseRestApi, err)
		}
	}

	h, api, err := restapi.HandlerAPI(config.Config)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrFailedToInitialiseRestApi, err)
	}

	api.UseSwaggerUI()
	h = api.Serve(config.Config.InnerMiddleware)

	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(config.Port),
		Handler: h,
	}

	go func() {
		log.Info().Str("pkg", "restapi").Msgf("Starting to serve, access server on http://localhost:%d", config.Port)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Str("pkg", "restapi").Err(err).Msg("Failed to listen")
		}
	}()

	go func() {
		<-ctx.Done()

		ctx, _ = context.WithTimeout(ctx, 15*time.Second)
		err = srv.Shutdown(ctx)
		log.Error().Str("pkg", "restapi").Err(err).Msg("Failed to shutdown gracefully")
	}()

	return nil
}
