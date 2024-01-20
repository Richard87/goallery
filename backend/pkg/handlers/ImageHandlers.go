package handlers

import (
	"github.com/Richard87/goallery/generated/models"
	"github.com/Richard87/goallery/generated/restapi/gaollery/images"
	"github.com/Richard87/goallery/internal/pointers"
	"github.com/go-openapi/runtime/middleware"
)

func (h *BaseHandlers) MapImageHandlers() {
	// h.api.ImagesGetImagesHandler = images.GetImagesHandlerFunc(func(request images.GetImagesParams) middleware.Responder {
	//
	// 	list, err := h.db.ListImages(request.HTTPRequest.Context())
	// 	if err != nil {
	// 		images.NewGetImagesInternalServerError().WithPayload(&models.ProblemDetails{
	// 			Detail: err.Error(),
	// 			Status: pointers.Int32(500),
	// 			Title:  pointers.String("Internal Server Error"),
	// 		})
	// 	}
	//
	// 	return images.NewGetImagesOK().WithPayload(list)
	// })

	h.api.ImagesGetImageByIDHandler = images.GetImageByIDHandlerFunc(func(request images.GetImageByIDParams, _ interface{}) middleware.Responder {
		image, err := h.db.GetImage(request.HTTPRequest.Context(), request.ID)
		if err != nil {
			images.NewGetImagesInternalServerError().WithPayload(&models.ProblemDetails{
				Detail: err.Error(),
				Status: pointers.Int32(500),
				Title:  pointers.String("Internal Server Error"),
			})
		}

		return images.NewGetImageByIDOK().WithPayload(image)
	})
}
