package imageclip

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"

	"github.com/termin/ocr-dere-result/fields"
	"go.uber.org/zap"
)

type SubImager interface {
	SubImage(r image.Rectangle) image.Image
}

func Clip(img image.Image, field fields.Field) (image.Image, error) {
	rect := field.Coordinate.Rect()
	clippedImage := img.(SubImager).SubImage(rect)

	env := os.Getenv("GO_ENV")
	if env == "dev" {
		output, err := os.Create(fmt.Sprintf("./debug/cutted_%v.jpg", field.Name))
		if err != nil {
			zap.S().Error("create debug image", err)
			return nil, err
		}
		defer output.Close()
		jpeg.Encode(output, clippedImage, &jpeg.Options{Quality: 100})
	}

	return clippedImage, nil
}
