package main

import (
	"context"
	"os/signal"

	"github.com/Richard87/goallery/pkg/blurredimage"
	"github.com/Richard87/goallery/pkg/config"
	"github.com/Richard87/goallery/pkg/controller"
	"github.com/Richard87/goallery/pkg/facescanner"
	"github.com/Richard87/goallery/pkg/inmemorydb"
	"github.com/Richard87/goallery/pkg/router"
	"github.com/Richard87/goallery/pkg/swagger"
	"github.com/Richard87/goallery/pkg/templates"
	"github.com/rs/zerolog/log"
	"golang.org/x/sys/unix"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), unix.SIGTERM, unix.SIGINT)
	defer cancel()

	cfg := config.ParseConfig()
	config.ConfigureLogger(cfg)

	db := inmemorydb.New(ctx, cfg.Photos,
		blurredimage.NewBlurredImageFeature(),
		facescanner.NewFaceScannerFeature(ctx, cfg.FaceScannerModelFile),
	)

	api := controller.NewController(db)
	frontend := templates.NewController()
	swagger := swagger.NewController()

	err := router.NewRouter(api, frontend, swagger).Serve(ctx, cfg.Port)
	log.Ctx(ctx).Err(err).Msg("Router closed")
}
