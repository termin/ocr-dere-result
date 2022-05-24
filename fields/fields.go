package fields

import (
	"fmt"
	"image"
	"strings"
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

type Result struct {
	*Field
	Text string
}

// 正規化した文字列を返す
func (r *Result) NormalizedText() (string, error) {
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

type Results []*Result

func NewResultByField(field Field) *Result {
	result := &Result{
		Field: &field,
	}
	return result
}

// TODO
func (r *Results) IsSuccessed() bool {
	return true
}
