package ocr

import (
	"bytes"
	"context"
	"image"
	"image/png"
	"io"

	vision "cloud.google.com/go/vision/apiv1"
)

func detectDocumentText(f io.Reader) (string, error) {
	ctx := context.Background()

	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		return "", err
	}

	image, err := vision.NewImageFromReader(f)
	if err != nil {
		return "", err
	}

	annotation, err := client.DetectDocumentText(ctx, image, nil)
	if err != nil {
		return "", err
	}

	return annotation.GetText(), nil
}

func Request(img image.Image) (string, error) {
	buf := new(bytes.Buffer)
	if err := png.Encode(buf, img); err != nil {
		return "", err
	}

	text, err := detectDocumentText(buf)
	return text, err
}
