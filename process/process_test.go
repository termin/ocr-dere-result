package process

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/termin/ocr-dere-result/fields"
)

func TestLoadFields(t *testing.T) {
	p, _ := os.Getwd()
	configFilepath := filepath.Join(p, "../configs/coordinates.json")
	fs, err := LoadFields(configFilepath)
	if err != nil {
		t.Error(err)
	}

	if len(fs) == 0 {
		t.Errorf("empty return value")
	}
}

func TestDo(t *testing.T) {
	var fs []fields.Field
	c := fields.Coordinate{StartX: 1005, StartY: 340, EndX: 1145, EndY: 410}
	fs = append(fs, fields.Field{Name: fields.Lv, Coordinate: c})
	c = fields.Coordinate{StartX: 1005, StartY: 340, EndX: 1145, EndY: 410}
	fs = append(fs, fields.Field{Name: fields.Title, Coordinate: c})

	p, _ := os.Getwd()
	imageFilepath := filepath.Join(p, "../examples/demolish.png")
	err := Do(fs, imageFilepath)
	if err != nil {
		t.Error(err)
	}
}
