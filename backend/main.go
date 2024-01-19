package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/Richard87/goallery/generated/models"
	"github.com/Richard87/goallery/generated/restapi"
	"github.com/Richard87/goallery/generated/restapi/gaollery"
	"github.com/Richard87/goallery/generated/restapi/gaollery/images"
	"github.com/Richard87/goallery/internal/pointers"
	"github.com/Richard87/goallery/pkg/db"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/jessevdk/go-flags"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type AppConfig struct {
	Photos   string `long:"photos-folder" description:"Directory to photos" default:"../photos" env:"PHOTOS_FOLDER"`
	LogLevel string `long:"log-level" description:"Log level" default:"info" env:"LOG_LEVEL"`
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	defer func() {
		signal.Stop(c)
		cancel()
	}()
	go func() {
		select {
		case <-c:
			cancel()
		case <-ctx.Done():
		}
	}()

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.DurationFieldInteger = true
	zerolog.LevelFieldName = "severity"
	pretty := os.Getenv("PRETTY") == "1"
	if pretty {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
	ctx = log.Logger.WithContext(ctx)

	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load swagger spec")
	}
	config := &AppConfig{}

	api := gaollery.NewGoalleryAPI(swaggerSpec)
	api.Logger = log.Printf
	server := restapi.NewServer(api)
	defer server.Shutdown()

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

	log.Info().Msg("Setting log level to " + config.LogLevel)
	level, err := zerolog.ParseLevel(config.LogLevel)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to parse log-level")
	}
	zerolog.SetGlobalLevel(level)

	db, err := db.NewInMemoryDb(ctx, config.Photos)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load InMemoryDb")
	}

	api.ImagesGetImagesHandler = images.GetImagesHandlerFunc(func(request images.GetImagesParams) middleware.Responder {

		list, err := db.ListImages(request.HTTPRequest.Context())
		if err != nil {
			images.NewGetImagesInternalServerError().WithPayload(&models.ProblemDetails{
				Detail: err.Error(),
				Status: pointers.Int32(500),
				Title:  pointers.String("Internal Server Error"),
			})
		}

		return images.NewGetImagesOK().WithPayload(list)
	})
	api.ImagesGetImageByIDHandler = images.GetImageByIDHandlerFunc(func(request images.GetImageByIDParams) middleware.Responder {
		image, err := db.GetImage(request.HTTPRequest.Context(), request.ID)
		if err != nil {
			images.NewGetImagesInternalServerError().WithPayload(&models.ProblemDetails{
				Detail: err.Error(),
				Status: pointers.Int32(500),
				Title:  pointers.String("Internal Server Error"),
			})
		}

		return images.NewGetImageByIDOK().WithPayload(image)
	})

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}

}
