package controller

import (
	"context"
	"errors"
	"fmt"

	"github.com/Richard87/goallery/api"
	"github.com/rs/zerolog/log"
)

func (h Controller) GetImages(ctx context.Context, request api.GetImagesRequestObject) (api.GetImagesResponseObject, error) {
	// TODO implement me

	images, err := h.db.ListImages(ctx)
	if err != nil {
		return nil, errors.New("not implemented")
	}

	return api.GetImages200JSONResponse(images), nil
}

func (h Controller) GetImageById(ctx context.Context, request api.GetImageByIdRequestObject) (api.GetImageByIdResponseObject, error) {

	image, err := h.db.GetImage(ctx, request.Id)
	if err != nil {
		return api.GetImageById200JSONResponse{}, err
	}
	return api.GetImageById200JSONResponse(image), nil
}

func (h Controller) DownloadImageById(ctx context.Context, request api.DownloadImageByIdRequestObject) (api.DownloadImageByIdResponseObject, error) {
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
