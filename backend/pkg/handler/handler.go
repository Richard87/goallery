package handler

import (
	"context"
	"fmt"

	"github.com/Richard87/goallery/api"
	"github.com/Richard87/goallery/pkg/inmemorydb"
	"github.com/equinor/radix-common/utils/pointers"
	"github.com/rs/zerolog/log"
)

type Handler struct {
	db *inmemorydb.InMemoryDb
}

var _ api.StrictServerInterface = Handler{}

func New(ctx context.Context, db *inmemorydb.InMemoryDb) Handler {
	return Handler{
		db: db,
	}
}

func (h Handler) GetToken(ctx context.Context, request api.GetTokenRequestObject) (api.GetTokenResponseObject, error) {
	return api.GetToken200JSONResponse{
		Token: "tokenABC",
	}, nil
}

func (h Handler) GetImages(ctx context.Context, request api.GetImagesRequestObject) (api.GetImagesResponseObject, error) {
	// TODO implement me

	images, err := h.db.ListImages(ctx)
	if err != nil {
		return api.GetImages500JSONResponse{Detail: pointers.Ptr(err.Error()), Status: 500, Title: "Internal Server Error"}, nil
	}

	return api.GetImages200JSONResponse(images), nil
}

func (h Handler) GetImageById(ctx context.Context, request api.GetImageByIdRequestObject) (api.GetImageByIdResponseObject, error) {

	image, err := h.db.GetImage(ctx, request.Id)
	if err != nil {
		return api.GetImageById200JSONResponse{}, err
	}
	return api.GetImageById200JSONResponse(image), nil
}

func (h Handler) DownloadImageById(ctx context.Context, request api.DownloadImageByIdRequestObject) (api.DownloadImageByIdResponseObject, error) {
	body, image, err := h.db.OpenImage(ctx, request.Id)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("failed to get image")
		return api.DownloadImageById200ImagegifResponse{}, err
	}

	switch image.Mime {
	case "image/jpeg":
		return api.DownloadImageById200ImagejpegResponse{body, image.Size}, nil
	case "image/png":
		return api.DownloadImageById200ImagepngResponse{body, image.Size}, nil
	case "image/gif":
		return api.DownloadImageById200ImagegifResponse{body, image.Size}, nil
	case "image/svg":
		return api.DownloadImageById200ImagesvgXmlResponse{body, image.Size}, nil
	case "image/webp":
		return api.DownloadImageById200ImagewebpResponse{body, image.Size}, nil
	}

	return api.DownloadImageById200ImagejpegResponse{}, fmt.Errorf("Not implemented image type: %s", image.Mime)
}
