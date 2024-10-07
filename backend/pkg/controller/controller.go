package controller

import (
	"context"
	"net/http"

	"github.com/Richard87/goallery/api"
	"github.com/Richard87/goallery/pkg/inmemorydb"
	"github.com/Richard87/goallery/pkg/router"
)

type Controller struct {
	db *inmemorydb.InMemoryDb
}

var _ api.StrictServerInterface = Controller{}

func NewController(ctx context.Context, db *inmemorydb.InMemoryDb) router.RouteMapper {
	controller := Controller{
		db: db,
	}

	si := api.NewStrictHandler(controller, []api.StrictMiddlewareFunc{})
	return func(router *http.ServeMux) {
		api.HandlerFromMuxWithBaseURL(si, router, "/api/v1")
	}
}
