package blurredimage

import (
	"bytes"
	"context"
	"encoding/base64"
	"image"
	"image/png"

	"github.com/Richard87/goallery/api"
	"github.com/Richard87/goallery/pkg/inmemorydb"
	"github.com/equinor/radix-common/utils/pointers"
	"github.com/nfnt/resize"
)

func NewBlurredImageFeature() inmemorydb.AddFeatureFunc {
	buf := bytes.NewBuffer(make([]byte, 100))

	return func(_ context.Context, imageBytes []byte, i image.Image, feature *api.ImageFeature) error {
		buf.Reset()
		newImage := resize.Resize(5, 5, i, resize.Lanczos3)
		if err := png.Encode(buf, newImage); err != nil {
			return err
		}

		feature.PluginBlurryimage = pointers.Ptr("data:image/png;base64," + base64.RawStdEncoding.EncodeToString(buf.Bytes()))

		return nil
	}
}
