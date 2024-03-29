package images

import (
	"context"

	"github.com/Richard87/goallery/generated/models"
	"github.com/Richard87/goallery/generated/restapi"
	"github.com/Richard87/goallery/generated/restapi/gaollery/images"
	"github.com/Richard87/goallery/internal/pointers"
	"github.com/Richard87/goallery/pkg/inmemorydb"
	"github.com/go-openapi/runtime/middleware"
	"github.com/rs/zerolog"
)

type Handler struct {
	db     *inmemorydb.InMemoryDb
	logger zerolog.Logger
}

var _ restapi.ImagesAPI = (*Handler)(nil)

func New(db *inmemorydb.InMemoryDb, logger zerolog.Logger) *Handler {
	return &Handler{
		db:     db,
		logger: logger.With().Str("pkg", "images").Logger(),
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

	image.Src = pointers.String("/api/v1/images/" + *image.ID + "/download")
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

	for key, image := range list {
		list[key].Src = pointers.String("/api/v1/images/" + *image.ID + "/download")
	}

	return images.NewGetImagesOK().WithPayload(list)
}

func (h *Handler) DownloadImageByID(ctx context.Context, params images.DownloadImageByIDParams) middleware.Responder {
	f, err := h.db.OpenImage(params.ID)
	if err != nil {
		h.logger.Error().Err(err).Msg("failed to list images")
		images.NewGetImagesInternalServerError().WithPayload(&models.ProblemDetails{
			Detail: err.Error(),
			Status: pointers.Int32(500),
			Title:  pointers.String("Internal Server Error"),
		})
	}

	return images.NewDownloadImageByIDOK().WithPayload(f)
}
