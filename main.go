package main

import (
	"context"
	_ "embed"
	"encoding/json"
	"io/fs"
	"net/http"
	"os"

	"github.com/Richard87/goallery/frontend"
	_ "github.com/Richard87/goallery/frontend"
	"github.com/Richard87/goallery/pkg/db"
	"github.com/Richard87/goallery/pkg/interfaces"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type App struct {
	Db interfaces.Db
	Fs fs.FS
}

func main() {
	ctx := context.Background()
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.DurationFieldInteger = true
	fsPath := "/Users/richard/Pictures/Darktable/20210720_1" // os.Getenv("GALLERY_PATH")
	pretty := os.Getenv("PRETTY") == "1"
	if pretty {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
	ctx = log.Logger.WithContext(ctx)

	log.Info().Msg("Starting Goallery")
	db, err := db.NewInMemoryDb(ctx, fsPath)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load InMemoryDb")
		os.Exit(1)
	}

	fileSystem := os.DirFS(fsPath)
	app := App{
		Db: db,
		Fs: fileSystem,
	}

	router := mux.NewRouter().StrictSlash(true)
	initializeFrontend(router)
	router.Handle("/api", app.api(ctx))

	log.Info().Str("server", "http://localhost:3000").Msg("Listening")

	handler := http.ListenAndServe(":3000", router)
	err = handler
	if err != nil {
		log.Error().Err(err).Msg("Server closed")
		os.Exit(1)
	}
}

func (a *App) api(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Ctx(ctx).Debug().Str("path", r.URL.Path).Msg("Got request")

		imgs, err := a.Db.ListImages(ctx)
		if err != nil {
			log.Ctx(ctx).Error().Err(err).Msg("Unable to fetch images!")
		}

		data, err := json.MarshalIndent(imgs, "", "  ")
		if err != nil {
			w.WriteHeader(500)
			log.Ctx(ctx).Error().Err(err).Msg("Unable to marshall images!")
			_, _ = w.Write([]byte("Unable to marshall data"))
			return
		}

		w.WriteHeader(200)
		_, _ = w.Write(data)
	}
}

func initializeFrontend(router *mux.Router) {
	frontend := http.FileServer(http.FS(frontend.FS()))
	router.PathPrefix("/").Handler(frontend)
}
