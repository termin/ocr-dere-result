package imageclip

import (
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"

	"github.com/termin/ocr-dere-result/fields"
)

type SubImager interface {
	SubImage(r image.Rectangle) image.Image
}

func Clip(img image.Image, field fields.Field) (image.Image, error) {
	rect := field.Coordinate.Rect()
	clippedImage := img.(SubImager).SubImage(rect)

	// TODO: 検証用のファイル出力
	output, err := os.Create(fmt.Sprintf("./debug/cutted_%v.jpg", field.Name))
	if err != nil {
		log.Println("create error:", err)
		return nil, err
	}
	defer output.Close()
	jpeg.Encode(output, clippedImage, &jpeg.Options{Quality: 100})

	return clippedImage, nil
}
