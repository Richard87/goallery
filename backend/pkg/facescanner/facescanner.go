package facescanner

import (
	"context"
	"image"

	"github.com/Richard87/goallery/api"
	"github.com/Richard87/goallery/pkg/inmemorydb"
	"github.com/jdeng/goface"
	"github.com/rs/zerolog/log"
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
)

type FaceScanner struct {
	fn *goface.Facenet
}

func NewFaceScannerFeature(ctx context.Context, modelFile string) inmemorydb.AddFeatureFunc {
	fn, err := goface.NewFacenet(modelFile)
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot create face scanner feature")
	}

	scanner := FaceScanner{fn: fn}

	go func() {
		<-ctx.Done()
		fn.Close()
	}()

	return scanner.Scan
}

func (f *FaceScanner) Scan(ctx context.Context, imageBytes []byte, _ image.Image, feature *api.ImageFeature) error {

	// detection
	img, err := goface.TensorFromJpeg(imageBytes)
	if err != nil {
		return err
	}

	det, err := goface.NewMtcnnDetector("mtcnn.pb")
	if err != nil {
		return err
	}
	bbox, err := det.DetectFaces(img) // [][]float32, i.e., [x1,y1,x2,y2],...
	if err != nil {
		return err
	}

	var faces []api.ImageFeatureFace
	for _, box := range bbox {
		faces = append(faces, api.ImageFeatureFace{
			X1: int(box[0]),
			Y1: int(box[1]),
			X2: int(box[2]),
			Y2: int(box[3]),
		})
	}

	feature.PluginFaces = &faces

	var cropSize int32 = 160
	ximgs, err := goface.CropResizeImage(img, bbox, []int32{cropSize, cropSize})
	if err != nil {
		return err
	}
	imgs := ximgs.Value().([][][][]float32)
	for _, img := range imgs {
		mean, std := goface.MeanStd(img)

		timg, err := tf.NewTensor([][][][]float32{img})
		if err != nil {
			log.Ctx(ctx).Warn().Err(err).Msg("Cannot create tensor")
			continue
		}

		wimg, err := goface.PrewhitenImage(timg, mean, std)
		if err != nil {
			log.Ctx(ctx).Warn().Err(err).Msg("Could not prewhiten image")
			continue
		}

		emb, err := f.fn.Embedding(wimg)
		if err != nil {
			log.Ctx(ctx).Warn().Err(err).Msg("Failed to get Embeddings")
			continue
		}

		log.Info().Int("embbeddings", len(emb)).Msg("Found Embeddings")
	}

	return nil
}
