package images

import (
	"context"

	"github.com/Richard87/goallery/generated/models"
	"github.com/Richard87/goallery/generated/restapi"
	"github.com/Richard87/goallery/generated/restapi/gaollery/images"
	"github.com/Richard87/goallery/internal/pointers"
	"github.com/Richard87/goallery/pkg/inmemorydb"
	"github.com/go-openapi/runtime/middleware"
)

type Handler struct {
	db *inmemorydb.InMemoryDb
}

var _ restapi.ImagesAPI = (*Handler)(nil)

func New(db *inmemorydb.InMemoryDb) *Handler {
	return &Handler{
		db: db,
	}
}

func (h *Handler) GetImageByID(ctx context.Context, params images.GetImageByIDParams) middleware.Responder {
	image, err := h.db.GetImage(ctx, params.ID)
	if err != nil {
		images.NewGetImagesInternalServerError().WithPayload(&models.ProblemDetails{
			Detail: err.Error(),
			Status: pointers.Int32(500),
			Title:  pointers.String("Internal Server Error"),
		})
	}

	return images.NewGetImageByIDOK().WithPayload(image)

}

func (h *Handler) GetImages(ctx context.Context, params images.GetImagesParams) middleware.Responder {

	list, err := h.db.ListImages(ctx)
	if err != nil {
		images.NewGetImagesInternalServerError().WithPayload(&models.ProblemDetails{
			Detail: err.Error(),
			Status: pointers.Int32(500),
			Title:  pointers.String("Internal Server Error"),
		})
	}

	return images.NewGetImagesOK().WithPayload(list)

}
