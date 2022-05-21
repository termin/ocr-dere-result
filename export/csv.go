package export

import (
	"fmt"
	"io"
	"strings"

	"github.com/termin/ocr-dere-result/fields"
)

type CSVExporter struct {
	results fields.Results
	w       io.Writer
}

func NewCSVExporter(r fields.Results, w io.Writer) *CSVExporter {
	return &CSVExporter{results: r, w: w}
}

func (e *CSVExporter) Export() error {
	if !e.results.IsSuccessed() {
		return fmt.Errorf("results is incomplete")
	}

	var texts []string
	for _, result := range e.results {
		text, err := result.NormalizeText()
		if err != nil {
			return err
		}
		texts = append(texts, text)
	}
	// TODO: 真面目にCSV文字列を作る
	csvString := strings.Join(texts, ", ")

	_, err := e.w.Write([]byte(csvString))
	if err != nil {
		return err
	}

	return nil
}
