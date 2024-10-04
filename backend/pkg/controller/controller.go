package controller

import (
	"context"

	"github.com/Richard87/goallery/api"
	"github.com/Richard87/goallery/pkg/inmemorydb"
)

type Controller struct {
	db *inmemorydb.InMemoryDb
}

var _ api.StrictServerInterface = Controller{}

func NewController(ctx context.Context, db *inmemorydb.InMemoryDb) Controller {
	return Controller{
		db: db,
	}
}
