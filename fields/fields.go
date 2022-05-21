package fields

import (
	"fmt"
	"image"
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
	// TODO: FieldNameに存在しない文字列の場合はどうなる？
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
func (r *Result) NormalizeText() (string, error) {
	// TODO
	// switch FieldTypeByName(r.Field.Name) {
	// case condition:
	// }
	return r.Text, nil
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
