package inmemorydb

import (
	"bytes"
	"context"
	"fmt"
	"image"
	_ "image/jpeg"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/Richard87/goallery/api"
	"github.com/gabriel-vasile/mimetype"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type InMemoryDb struct {
	rootDir string
	fs      fs.FS
	images  map[string]Image
	logger  zerolog.Logger
}
type Image struct {
	api.Image
	path string
}

var (
	ErrInMemoryDb     = fmt.Errorf("inmemorydb error")
	ErrImageNotFound  = fmt.Errorf("%w: image not found", ErrInMemoryDb)
	ErrNotImplemented = fmt.Errorf("%w: not implemented", ErrInMemoryDb)
)

type AddFeatureFunc func(context.Context, []byte, image.Image, *api.ImageFeature) error

func New(ctx context.Context, rootDir string, features ...AddFeatureFunc) *InMemoryDb {
	cwd, _ := os.Getwd()
	db := &InMemoryDb{
		rootDir: path.Join(cwd, rootDir),
		fs:      os.DirFS(rootDir),
		images:  make(map[string]Image),
		logger:  log.Ctx(ctx).With().Str("pkg", "inmemorydb").Logger(),
	}

	go db.ScanPhotos(ctx, features...)

	return db
}

func (db *InMemoryDb) GetImage(ctx context.Context, id string) (api.Image, error) {

	i, ok := db.images[id]
	if !ok {
		return api.Image{}, ErrImageNotFound
	}

	return i.Image, nil
}

func (db *InMemoryDb) ListImages(_ context.Context) ([]api.Image, error) {
	res := make([]api.Image, len(db.images))

	x := 0
	for _, v := range db.images {
		v2 := v
		res[x] = v2.Image
		x++
	}

	return res, nil
}

func (db *InMemoryDb) StoreImage(ctx context.Context, image fs.File) (api.Image, error) {
	panic("implement me")
	return api.Image{}, ErrNotImplemented
}

func (db *InMemoryDb) ScanPhotos(ctx context.Context, featureFuncs ...AddFeatureFunc) {

	start := time.Now()

	db.logger.Info().Str("path", db.rootDir).Msg("Scanning directory...")

	err := filepath.Walk(db.rootDir, func(path string, info os.FileInfo, err error) error {
		if ctx.Err() != nil {
			return ctx.Err()
		}

		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		localPath := strings.TrimPrefix(strings.Replace(path, db.rootDir, "", 1), "/")
		logger := log.Ctx(ctx).With().Str("path", localPath).Logger()

		mtype, err := mimetype.DetectFile(path)
		if err != nil {
			logger.Warn().Err(err).Msg("Unable to get Mime from file")
			return nil
		}
		mime := mtype.String()
		fileSize := info.Size()
		logger = logger.With().Str("mime", mtype.String()).Int64("size", fileSize).Logger()
		if !strings.HasPrefix(mime, "image/") {
			logger.Debug().Msg("Skipping file")
			return nil
		}
		logger.Debug().Msg("Found image")

		f, err := db.fs.Open(localPath)
		if err != nil {
			logger.Warn().Err(err).Msg("Unable to open file")
			return nil
		}
		defer f.Close()

		imageBytes, err := io.ReadAll(f)
		if err != nil {
			logger.Warn().Err(err).Msg("Unable to read file")
			return nil
		}
		r := bytes.NewReader(imageBytes)

		i, _, err := image.Decode(r)
		if err != nil {
			logger.Warn().Err(err).Msg("Unable to decode file")
			return nil
		}

		id := strconv.Itoa(len(db.images) + 1)
		image := Image{
			Image: api.Image{
				Id:       id,
				Filename: info.Name(),
				Mime:     mime,
				Width:    int64(i.Bounds().Max.X),
				Height:   int64(i.Bounds().Max.Y),
				Size:     fileSize,
				Features: api.ImageFeature{},
			},
			path: localPath,
		}
		for _, fn := range featureFuncs {
			featureCtx := logger.WithContext(ctx)
			err := fn(featureCtx, imageBytes, i, &image.Features)
			if err != nil {
				logger.Warn().Err(err).Msg("Unable to add feature")
				return err
			}
		}
		db.images[id] = image
		return nil
	})
	if err != nil {
		db.logger.Error().Err(err).Msg("Failed to walk directory")
		return
	}

	db.logger.Info().
		Int("images", len(db.images)).
		Dur("ellapsed", time.Since(start)).
		Str("path", db.rootDir).
		Msg("InMemoryDb loaded")

	if len(db.images) == 0 {
		db.logger.Warn().Msg("No images found")
	}
}

func (db *InMemoryDb) OpenImage(ctx context.Context, id string) (fs.File, api.Image, error) {
	i := db.images[id]

	image, err := db.GetImage(ctx, id)
	if err != nil {
		return nil, api.Image{}, err
	}

	open, err := db.fs.Open(i.path)
	return open, image, err
}
