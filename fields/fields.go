package fields

import (
	"fmt"
	"image"
	"log"
	"os"
	"strings"
	"time"

	"github.com/rwcarlsen/goexif/exif"
)

type Coordinate struct {
	StartX int `json:"startx"`
	StartY int `json:"starty"`
	EndX   int `json:"endx"`
	EndY   int `json:"endy"`
}

func (c Coordinate) Rect() image.Rectangle {
	return image.Rect(c.StartX, c.StartY, c.EndX, c.EndY)
}

type Field struct {
	Name       FieldName  `json:"name"`
	Coordinate Coordinate `json:"coordinate"`
}

func (f Field) String() string {
	fieldType, _ := FieldTypeByName(FieldName(f.Name))
	return fmt.Sprintf(
		"name: %v, type: %v, Coordinate: {startx: %v, starty: %v, endx: %v, endy: %v}",
		f.Name,
		fieldType,
		f.Coordinate.StartX,
		f.Coordinate.StartY,
		f.Coordinate.EndX,
		f.Coordinate.EndY,
	)
}

type ResultField struct {
	*Field
	Text string
}

// 正規化した文字列を返す
func (r *ResultField) NormalizedText() (string, error) {
	var normalized string
	fieldType, err := FieldTypeByName(r.Field.Name)
	if err != nil {
		return "", err
	}

	switch fieldType {
	case FieldTypeString:
		normalized = r.Text
		normalized = strings.ReplaceAll(r.Text, "\n", "")
	case FieldTypeDigits:
		for _, n := range r.Text {
			// 数字以外除外
			if strings.ContainsAny(string(n), "0123456789") {
				normalized += string(n)
			}
		}
		if strings.Count(normalized, "0") == len(normalized) {
			// 0000 => 0
			normalized = "0"
		} else {
			normalized = strings.TrimLeft(normalized, "0")
		}
	default:
		return "", fmt.Errorf("unknown FieldType")
	}
	return normalized, nil
}

func NewResultFieldByField(field Field) *ResultField {
	result := &ResultField{
		Field: &field,
	}
	return result
}

type Result struct {
	SourceImageFile *os.File
	Fields          []*ResultField
}

func (r *Result) AddResultField(field *ResultField) {
	r.Fields = append(r.Fields, field)
}

// TODO
func (r *Result) IsSuccessful() bool {
	return true
}

// 日付を返す. Exif, FileInfo.ModTimeの順にフォールバック
func (r *Result) DateTime() (time.Time, error) {
	t, err := r.DateTimeFromExif()
	if err == nil {
		return t, nil
	}

	t, err = r.DateTimeFromFileInfo()
	if err == nil {
		return t, nil
	}

	return time.Time{}, err
}

func (r *Result) DateTimeFromFileInfo() (time.Time, error) {
	info, err := r.SourceImageFile.Stat()
	if err != nil {
		return time.Time{}, err
	}

	return info.ModTime(), nil
}

func (r *Result) DateTimeFromExif() (time.Time, error) {
	e, err := exif.Decode(r.SourceImageFile)
	if err != nil {
		log.Println("failed to decode")
		return time.Time{}, err
	}

	origTime, err := e.DateTime()
	if err != nil {
		log.Println("failed to get DateTimeOriginal")
		return time.Time{}, err
	}

	return origTime, nil
}
