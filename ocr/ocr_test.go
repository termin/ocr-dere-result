package ocr

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
	"testing"
)

func TestReuqest(t *testing.T) {
	t.SkipNow()

	p, _ := os.Getwd()
	imageFilepath := filepath.Join(p, "../examples/demolish_cutted_title.jpg")
	f, err := os.Open(imageFilepath)
	if err != nil {
		t.Error(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		t.Error(err)
	}

	text, err := Request(img)
	if err != nil {
		t.Error(err)
	}
	t.Logf("text: %v", text)
	if len(text) == 0 {
		t.Errorf("detected text is empty")
	}
}
