package export

import (
	"bytes"
	"testing"

	"github.com/termin/ocr-dere-result/fields"
)

func fixture() fields.Results {
	var results fields.Results
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
		r := &fields.Result{Field: f, Text: text}
		results = append(results, r)
	}

	return results
}

func TestExport(t *testing.T) {
	results := fixture()
	buf := new(bytes.Buffer)
	e := NewCSVExporter(buf)
	err := e.Export(results)

	t.Log(buf.String())
	t.Log(results)

	if err != nil {
		t.Error(err)
	}

	expect := "title	2022/05/25		31		900	5	4	3	2		600\n"
	actual := buf.String()
	if actual != expect {
		t.Errorf("expect: %s\nactual: %s\n", expect, actual)
	}
}
