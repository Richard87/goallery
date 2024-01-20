package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/Richard87/goallery/generated/models"
	"github.com/Richard87/goallery/internal/pointers"
	"github.com/Richard87/goallery/pkg/handlers/auth"
	"github.com/Richard87/goallery/pkg/handlers/images"
	"github.com/Richard87/goallery/pkg/inmemorydb"
	"github.com/Richard87/goallery/pkg/restapi"
	"github.com/jessevdk/go-flags"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type AppConfig struct {
	Photos    string `long:"photos-folder" description:"Directory to photos" default:"../photos" env:"PHOTOS_FOLDER"`
	LogLevel  string `long:"log-level" description:"Log level" default:"info" env:"LOG_LEVEL"`
	LogFormat string `long:"log-format" description:"Log format ('json' or 'text')" default:"text" env:"LOG_FORMAT"`
	Port      int    `long:"port" description:"Port to listen on" default:"8000" env:"PORT"`
}

func main() {
	ctx, cancel := createContextWithGracefulShutdown(time.Second * 15)
	defer cancel()

	config := ParseConfig()

	configureLogger(config)
	ctx = log.Logger.WithContext(ctx)

	db, err := inmemorydb.New(ctx, config.Photos)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load InMemoryDb")
	}

	err = restapi.New(ctx,
		restapi.WithZerolog(log.Logger),
		restapi.WithAuthApi(auth.New(db)),
		restapi.WithImagesApi(images.New(db)),
		restapi.WithHttpPort(config.Port),
		WithTodoAuth(),
	)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to configure server")
	}

	<-ctx.Done()
}

func WithTodoAuth() restapi.OptionFunc {
	return func(config *restapi.Config) error {
		config.AuthBearer = func(token string) (*models.User, error) {
			return &models.User{
				ID:       pointers.String("1"),
				Password: pointers.String("password"),
				Token:    &token,
				Username: pointers.String("todo"),
			}, nil
		}
		return nil
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
