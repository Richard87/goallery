package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Richard87/goallery/api"
	"github.com/Richard87/goallery/pkg/handler"
	"github.com/Richard87/goallery/pkg/inmemorydb"
	gincommon "github.com/equinor/radix-common/pkg/gin"
	"github.com/gin-gonic/gin"
	"github.com/jessevdk/go-flags"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/sys/unix"
)

type AppConfig struct {
	Photos    string `long:"photos-folder" description:"Directory to photos" default:"../photos" env:"PHOTOS_FOLDER"`
	LogLevel  string `long:"log-level" description:"Log level" default:"info" env:"LOG_LEVEL"`
	LogFormat string `long:"log-format" description:"Log format ('json' or 'text')" default:"text" env:"LOG_FORMAT"`
	Port      int    `long:"port" description:"Port to listen on" default:"8000" env:"PORT"`
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), unix.SIGTERM, unix.SIGINT)
	defer cancel()

	config := ParseConfig()

	configureLogger(config)
	ctx = log.Logger.WithContext(ctx)

	db := inmemorydb.New(ctx, config.Photos)

	handler := handler.New(ctx, db)
	router := initializeGin(ctx, handler)

	go runServer(ctx, router)

	<-ctx.Done()
	log.Info().Msg("Exited.")
}

func initializeGin(ctx context.Context, server api.StrictServerInterface) *gin.Engine {
	gin.SetMode("release")
	gin.DefaultWriter = log.Ctx(ctx)
	gin.DefaultErrorWriter = log.Ctx(ctx)
	router := gin.Default()
	router.Use(
		gincommon.SetZerologLogger(gincommon.ZerologLoggerWithRequestId),
		gincommon.ZerologRequestLogger(),
		gin.Recovery(),
	)

	handler := api.NewStrictHandler(server, []api.StrictMiddlewareFunc{})
	api.RegisterHandlers(router, handler)
	return router
}

func ParseConfig() *AppConfig {
	config := &AppConfig{}

	parser := flags.NewParser(config, flags.Default)
	parser.ShortDescription = "Goallery"

	if _, err := parser.Parse(); err != nil {
		os.Exit(1)
	}

	return config
}
func configureLogger(config *AppConfig) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.DurationFieldInteger = true
	if config.LogFormat == "text" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	log.Info().Msg("Setting log level to " + config.LogLevel + ", format: " + config.LogFormat)
	level, err := zerolog.ParseLevel(config.LogLevel)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to parse log-level")
	}
	zerolog.SetGlobalLevel(level)
}

func runServer(ctx context.Context, handler http.Handler) {

	// And we serve HTTP until the world ends.
	s := &http.Server{
		Handler: handler,
		Addr:    "0.0.0.0:8080",
	}
	go func() {
		log.Ctx(ctx).Info().Msg("Starting server on http://localhost:8080")

		err := s.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Ctx(ctx).Fatal().Msg(err.Error())
		}
	}()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.WithoutCancel(ctx), 10*time.Second)
	defer cancel()

	err := s.Shutdown(shutdownCtx)
	if err != nil {
		log.Ctx(ctx).Error().Err(err).Msg("Failed to close gracefully")
		return
	}

	log.Ctx(ctx).Info().Msg("Server closed")
}
