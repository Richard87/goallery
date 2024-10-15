package config

import (
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type AppConfig struct {
	Photos               string `long:"photos-folder" description:"Directory to photos" default:"../photos" env:"PHOTOS_FOLDER"`
	LogLevel             string `long:"log-level" description:"Log level" default:"info" env:"LOG_LEVEL"`
	LogFormat            string `long:"log-format" description:"Log format ('json' or 'text')" default:"text" env:"LOG_FORMAT"`
	Port                 int    `long:"port" description:"Port to listen on" default:"8000" env:"PORT"`
	FaceScannerModelFile string `long:"facenet-model" description:"Path to facenet model file" default:"facenet.pb" env:"FACENET_MODEL"`
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
func ConfigureLogger(config *AppConfig) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.DurationFieldInteger = false
	if config.LogFormat == "text" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	log.Info().Msg("Setting log level to " + config.LogLevel + ", format: " + config.LogFormat)
	level, err := zerolog.ParseLevel(config.LogLevel)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to parse log-level")
	}
	zerolog.SetGlobalLevel(level)
	zerolog.DefaultContextLogger = &log.Logger
}
