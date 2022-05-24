package export

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/termin/ocr-dere-result/fields"
)

func fixture() (*fields.Result, error) {
	p, _ := os.Getwd()
	sourceImageFilepath := filepath.Join(p, "../examples/demolish.png")
	f, err := os.Open(sourceImageFilepath)
	if err != nil {
		return nil, err
	}

	result := &fields.Result{SourceImageFile: f}
	source := map[fields.FieldName]string{
		fields.Lv:         "31",
		fields.Title:      "title",
		fields.Difficulty: "MASTER+",
		fields.Perfect:    "900",
		fields.Great:      "5",
		fields.Nice:       "4",
		fields.Bad:        "3",
		fields.Miss:       "2",
		fields.Combo:      "600",
		fields.Score:      "1000000",
	}

	for name, text := range source {
		f := &fields.Field{
			Name:       name,
			Coordinate: fields.Coordinate{StartX: 0, StartY: 0, EndX: 0, EndY: 0},
		}
		r := &fields.ResultField{Field: f, Text: text}
		result.AddResultField(r)
	}

	return result, nil
}

func TestExport(t *testing.T) {
	result, err := fixture()
	if err != nil {
		t.FailNow()
	}

	buf := new(bytes.Buffer)
	e := NewCSVExporter(buf)
	err = e.Export(result)

	t.Log(buf.String())
	t.Log(result)

	if err != nil {
		t.Error(err)
	}

	expect := "title	2022/05/25		31		900	5	4	3	2		600\n"
	actual := buf.String()
	if actual != expect {
		t.Errorf("expect: %s\nactual: %s\n", expect, actual)
	}
}
