package controller

import (
	"context"

	"github.com/Richard87/goallery/api"
)

func (h Controller) GetToken(ctx context.Context, request api.GetTokenRequestObject) (api.GetTokenResponseObject, error) {
	return api.GetToken200JSONResponse{
		Token: "tokenABC",
	}, nil
}
