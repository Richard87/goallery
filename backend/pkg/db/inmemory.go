package db

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/Richard87/goallery/pkg/interfaces"
	"github.com/Richard87/goallery/pkg/model"
	"github.com/gabriel-vasile/mimetype"
	"github.com/nfnt/resize"
	"github.com/rs/zerolog/log"
)

type InMemoryDb struct {
	rootDir string
	dir     fs.FS
	images  map[string]model.Image
}

var ErrImageNotFound = fmt.Errorf("image not found")

func (db *InMemoryDb) GetImage(ctx context.Context, id string) (*model.Image, error) {
	i, ok := db.images[id]
	if !ok {
		return nil, ErrImageNotFound
	}

	return &i, nil
}

func (db *InMemoryDb) ListImages(ctx context.Context) ([]*model.Image, error) {
	var res []*model.Image

	for _, v := range db.images {
		res = append(res, &v)
	}

	return res, nil
}

func (db *InMemoryDb) StoreImage(ctx context.Context, image fs.File) (*model.Image, error) {
	// TODO implement me
	panic("implement me")
}

func NewInMemoryDb(ctx context.Context, rootDir string) (interfaces.Db, error) {
	db := &InMemoryDb{
		rootDir: rootDir,
		dir:     os.DirFS(rootDir),
		images:  make(map[string]model.Image),
	}

	start := time.Now()
	log.Ctx(ctx).Info().Str("path", rootDir).Msg("Scanning directory...")
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		filepath := filepath.Join(rootDir, info.Name())
		mtype, err := mimetype.DetectFile(filepath)
		if err != nil {
			log.Ctx(ctx).Warn().Str("filepath", filepath).Err(err).Msg("Unable to get Mime from file")
			return nil
		}
		if !strings.HasPrefix(mtype.String(), "image/") {
			log.Ctx(ctx).Debug().Str("filepath", path).Str("mime", mtype.String()).Int64("size", info.Size()).Msg("Skipping file")

			return nil
		}
		log.Ctx(ctx).Debug().Str("filepath", path).Str("mime", mtype.String()).Int64("size", info.Size()).Msg("Found image")

		f, err := db.dir.Open(info.Name())
		if err != nil {
			log.Ctx(ctx).Warn().Str("filepath", filepath).Err(err).Msg("Unable to load file")
			return nil
		}
		defer f.Close()

		i, _, err := image.Decode(f)
		if err != nil {
			log.Ctx(ctx).Warn().
				Str("filepath", filepath).
				Str("mime", mtype.String()).
				Err(err).
				Msg("Unable to decode file")
			return nil
		}

		var buf bytes.Buffer
		newImage := resize.Resize(5, 5, i, resize.Lanczos3)
		err = png.Encode(&buf, newImage)

		id := strconv.Itoa(len(db.images) + 1)
		db.images[id] = model.Image{
			Id:               id,
			OriginalFilename: info.Name(),
			Path:             filepath,
			Mime:             mtype.String(),
			Rect:             i.Bounds().Max,
			Size:             info.Size(),
			Placeholder:      "data:image/png;base64," + base64.RawStdEncoding.EncodeToString(buf.Bytes()),
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to walk directory: %w", err)
	}

	log.Ctx(ctx).Info().
		Int("images", len(db.images)).
		Dur("ellapsed", time.Now().Sub(start)).
		Str("path", rootDir).
		Msg("InMemoryDb loaded")

	if len(db.images) == 0 {
		log.Ctx(ctx).Warn().Msg("No images found")
	}
	return db, nil
}

var _ interfaces.Db = &InMemoryDb{}
