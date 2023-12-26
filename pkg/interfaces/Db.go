package interfaces

import (
	"context"
	"io/fs"

	"github.com/Richard87/goallery/pkg/model"
)

type Db interface {
	ListImages(ctx context.Context) ([]*model.Image, error)
	StoreImage(ctx context.Context, image fs.File) (*model.Image, error)
	GetImage(ctx context.Context, id string) (*model.Image, error)
}
