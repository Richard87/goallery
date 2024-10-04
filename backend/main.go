package main

import (
	"context"
	"os/signal"

	"github.com/Richard87/goallery/pkg/config"
	"github.com/Richard87/goallery/pkg/controller"
	"github.com/Richard87/goallery/pkg/inmemorydb"
	"github.com/Richard87/goallery/pkg/router"
	"github.com/rs/zerolog/log"
	"golang.org/x/sys/unix"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), unix.SIGTERM, unix.SIGINT)
	defer cancel()

	cfg := config.ParseConfig()
	config.ConfigureLogger(cfg)

	db := inmemorydb.New(ctx, cfg.Photos)
	controller := controller.NewController(ctx, db)
	r := router.NewRouter(controller)

	err := router.RunServer(ctx, r, cfg)
	log.Ctx(ctx).Err(err).Msg("Server closed")
}
