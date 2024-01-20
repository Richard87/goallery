package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/Richard87/goallery/generated/restapi"
	"github.com/Richard87/goallery/generated/restapi/gaollery"
	"github.com/Richard87/goallery/pkg/db"
	"github.com/Richard87/goallery/pkg/handlers"
	"github.com/go-openapi/loads"
	"github.com/jessevdk/go-flags"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type AppConfig struct {
	Photos    string `long:"photos-folder" description:"Directory to photos" default:"../photos" env:"PHOTOS_FOLDER"`
	LogLevel  string `long:"log-level" description:"Log level" default:"info" env:"LOG_LEVEL"`
	LogFormat string `long:"log-format" description:"Log format ('json' or 'text')" default:"text" env:"LOG_FORMAT"`
}

func main() {
	ctx, cancel := createContextWithGracefulShutdown(time.Second * 15)
	defer cancel()

	config, server, api := ParseConfig()
	defer server.Shutdown()

	configureLogger(config, api)
	ctx = log.Logger.WithContext(ctx)

	db, err := db.NewInMemoryDb(ctx, config.Photos)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load InMemoryDb")
	}

	server.ConfigureAPI()
	handlers.ConfigureHandlers(db, api)
	if err = api.Validate(); err != nil {
		log.Fatal().Err(err).Msg("Some handlers have not been implemented")
	}

	if err := server.Serve(); err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}

}

func createContextWithGracefulShutdown(timeout time.Duration) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		cancel()
		time.Sleep(timeout)
		os.Exit(2)
	}()

	cancelFunc := func() {
		cancel()
		signal.Stop(c)
	}

	return ctx, cancelFunc
}
func ParseConfig() (*AppConfig, *restapi.Server, *gaollery.GoalleryAPI) {
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load swagger spec")
	}
	config := &AppConfig{}

	api := gaollery.NewGoalleryAPI(swaggerSpec)
	server := restapi.NewServer(api)

	parser := flags.NewParser(config, flags.Default)
	parser.ShortDescription = "Goallery"

	_, err = parser.AddGroup("Server", "Server", server)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to parse server config")
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	return config, server, api
}
func configureLogger(config *AppConfig, api *gaollery.GoalleryAPI) {
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

	api.Logger = func(f string, args ...interface{}) {
		log.Info().Msgf(f, args...)
	}
}
