package auth

import (
	"context"

	"github.com/Richard87/goallery/generated/models"
	"github.com/Richard87/goallery/generated/restapi"
	"github.com/Richard87/goallery/generated/restapi/gaollery/auth"
	"github.com/Richard87/goallery/internal/pointers"
	"github.com/Richard87/goallery/pkg/inmemorydb"
	"github.com/go-openapi/runtime/middleware"
)

type Handler struct {
	db *inmemorydb.InMemoryDb
}

var _ restapi.AuthAPI = (*Handler)(nil)

func New(db *inmemorydb.InMemoryDb) *Handler {
	return &Handler{
		db: db,
	}
}

func (h *Handler) GetToken(ctx context.Context, params auth.GetTokenParams) middleware.Responder {
	return auth.NewGetTokenOK().WithPayload(&models.AuthResponse{Token: pointers.String("tokenXYZ")})
}
