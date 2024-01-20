package inmemorydb

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
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/Richard87/goallery/generated/models"
	"github.com/Richard87/goallery/internal/pointers"
	"github.com/gabriel-vasile/mimetype"
	"github.com/nfnt/resize"
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
	models.Image
	path string
}

var (
	ErrInMemoryDb     = fmt.Errorf("inmemorydb error")
	ErrImageNotFound  = fmt.Errorf("%w: image not found", ErrInMemoryDb)
	ErrNotImplemented = fmt.Errorf("%w: not implemented", ErrInMemoryDb)
)

func New(ctx context.Context, rootDir string) (*InMemoryDb, error) {
	cwd, _ := os.Getwd()
	db := &InMemoryDb{
		rootDir: path.Join(cwd, rootDir),
		fs:      os.DirFS(rootDir),
		images:  make(map[string]Image),
		logger:  log.Ctx(ctx).With().Str("pkg", "inmemorydb").Logger(),
	}

	go db.ScanPhotos(ctx)

	return db, nil
}

func (db *InMemoryDb) GetImage(ctx context.Context, id string) (*models.Image, error) {

	i, ok := db.images[id]
	if !ok {
		return nil, ErrImageNotFound
	}

	return &i.Image, nil
}

func (db *InMemoryDb) ListImages(_ context.Context) ([]*models.Image, error) {
	res := make([]*models.Image, len(db.images))

	x := 0
	for _, v := range db.images {
		res[x] = &v.Image
		x++
	}

	return res, nil
}

func (db *InMemoryDb) StoreImage(ctx context.Context, image fs.File) (*models.Image, error) {
	// TODO implement me
	return nil, ErrNotImplemented
}

func (db *InMemoryDb) ScanPhotos(ctx context.Context) {

	start := time.Now()

	db.logger.Info().Str("path", db.rootDir).Msg("Scanning directory...")

	err := filepath.Walk(db.rootDir, func(path string, info os.FileInfo, err error) error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		localPath := strings.TrimPrefix(strings.Replace(path, db.rootDir, "", 1), "/")

		mtype, err := mimetype.DetectFile(path)
		if err != nil {
			log.Ctx(ctx).Warn().Str("filepath", localPath).Err(err).Msg("Unable to get Mime from file")
			return nil
		}
		mime := mtype.String()
		fileSize := info.Size()
		if !strings.HasPrefix(mime, "image/") {
			log.Ctx(ctx).Debug().Str("filepath", localPath).Str("mime", mime).Int64("size", fileSize).Msg("Skipping file")

			return nil
		}
		log.Ctx(ctx).Debug().Str("filepath", localPath).Str("mime", mime).Int64("size", fileSize).Msg("Found image")

		f, err := db.fs.Open(localPath)
		if err != nil {
			log.Ctx(ctx).Warn().Str("filepath", localPath).Err(err).Msg("Unable to load file")
			return nil
		}
		defer f.Close()

		i, _, err := image.Decode(f)
		if err != nil {
			log.Ctx(ctx).Warn().
				Str("filepath", localPath).
				Str("mime", mime).
				Err(err).
				Msg("Unable to decode file")
			return nil
		}

		var buf bytes.Buffer
		newImage := resize.Resize(5, 5, i, resize.Lanczos3)
		err = png.Encode(&buf, newImage)

		id := strconv.Itoa(len(db.images) + 1)
		db.images[id] = Image{
			Image: models.Image{
				ID:       &id,
				Filename: pointers.String(info.Name()),
				Mime:     &mime,
				Width:    pointers.Int64(int64(i.Bounds().Max.X)),
				Height:   pointers.Int64(int64(i.Bounds().Max.Y)),
				Size:     &fileSize,
				Features: &models.ImageFeature{
					PluginBlurryimage: "data:image/png;base64," + base64.RawStdEncoding.EncodeToString(buf.Bytes()),
				},
			},
			path: localPath,
		}
		return nil
	})
	if err != nil {
		db.logger.Error().Err(err).Msg("Failed to walk directory")
		return
	}

	db.logger.Info().
		Int("images", len(db.images)).
		Dur("ellapsed", time.Now().Sub(start)).
		Str("path", db.rootDir).
		Msg("InMemoryDb loaded")

	if len(db.images) == 0 {
		db.logger.Warn().Msg("No images found")
	}
}
