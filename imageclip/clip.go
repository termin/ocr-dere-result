package imageclip

import (
	"image"

	"github.com/termin/ocr-dere-result/fields"
)

type SubImager interface {
	SubImage(r image.Rectangle) image.Image
}

func Clip(img image.Image, field fields.Field) (image.Image, error) {
	rect := field.Coordinate.Rect()
	clippedImage := img.(SubImager).SubImage(rect)

	// TODO: 検証用のファイル出力
	// output, err := os.Create("cutted.jpg")
	// if err != nil {
	// 	log.Println("create error:", err)
	// 	return
	// }
	// defer output.Close()
	// jpeg.Encode(output, clippedImage, &jpeg.Options{Quality: 100})
	//
	return clippedImage, nil
}
