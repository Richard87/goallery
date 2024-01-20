package handlers

import (
	"github.com/Richard87/goallery/generated/restapi/gaollery"
	"github.com/Richard87/goallery/pkg/db"
)

type BaseHandlers struct {
	db  *db.InMemoryDb
	api *gaollery.GoalleryAPI
}

func ConfigureHandlers(db *db.InMemoryDb, api *gaollery.GoalleryAPI) *BaseHandlers {

	h := &BaseHandlers{
		db:  db,
		api: api,
	}
	h.MapImageHandlers()

	return h
}
