package export

import (
	"bytes"
	"testing"

	"github.com/termin/ocr-dere-result/fields"
)

func fixture() fields.Results {
	var results fields.Results
	f0 := &fields.Field{
		Name:       fields.Lv,
		Coordinate: fields.Coordinate{StartX: 0, StartY: 0, EndX: 0, EndY: 0},
	}
	r0 := fields.Result{Field: f0, Text: "31"}
	results = append(results, &r0)

	f1 := &fields.Field{
		Name:       fields.Title,
		Coordinate: fields.Coordinate{StartX: 0, StartY: 0, EndX: 0, EndY: 0},
	}
	r1 := &fields.Result{Field: f1, Text: "title"}
	results = append(results, r1)

	f2 := &fields.Field{
		Name:       fields.Difficulty,
		Coordinate: fields.Coordinate{StartX: 0, StartY: 0, EndX: 0, EndY: 0},
	}
	r2 := &fields.Result{Field: f2, Text: "MASTER+"}
	results = append(results, r2)

	f3 := &fields.Field{
		Name:       fields.Perfect,
		Coordinate: fields.Coordinate{StartX: 0, StartY: 0, EndX: 0, EndY: 0},
	}
	r3 := &fields.Result{Field: f3, Text: "900"}
	results = append(results, r3)

	f4 := &fields.Field{
		Name:       fields.Great,
		Coordinate: fields.Coordinate{StartX: 0, StartY: 0, EndX: 0, EndY: 0},
	}
	r4 := &fields.Result{Field: f4, Text: "5"}
	results = append(results, r4)

	f5 := &fields.Field{
		Name:       fields.Nice,
		Coordinate: fields.Coordinate{StartX: 0, StartY: 0, EndX: 0, EndY: 0},
	}
	r5 := &fields.Result{Field: f5, Text: "4"}
	results = append(results, r5)

	f6 := &fields.Field{
		Name:       fields.Bad,
		Coordinate: fields.Coordinate{StartX: 0, StartY: 0, EndX: 0, EndY: 0},
	}
	r6 := &fields.Result{Field: f6, Text: "3"}
	results = append(results, r6)

	f7 := &fields.Field{
		Name:       fields.Miss,
		Coordinate: fields.Coordinate{StartX: 0, StartY: 0, EndX: 0, EndY: 0},
	}
	r7 := &fields.Result{Field: f7, Text: "2"}
	results = append(results, r7)

	f8 := &fields.Field{
		Name:       fields.Combo,
		Coordinate: fields.Coordinate{StartX: 0, StartY: 0, EndX: 0, EndY: 0},
	}
	r8 := &fields.Result{Field: f8, Text: "600"}
	results = append(results, r8)

	f9 := &fields.Field{
		Name:       fields.Score,
		Coordinate: fields.Coordinate{StartX: 0, StartY: 0, EndX: 0, EndY: 0},
	}
	r9 := &fields.Result{Field: f9, Text: "1000000"}
	results = append(results, r9)

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
