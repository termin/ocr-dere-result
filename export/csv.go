package export

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/termin/ocr-dere-result/fields"
)

type CSVExporter struct {
	w io.Writer
}

func NewCSVExporter(w io.Writer) *CSVExporter {
	return &CSVExporter{w: w}
}

func (e *CSVExporter) Export(results fields.Results) error {
	if !results.IsSuccessed() {
		return fmt.Errorf("results is incomplete")
	}

	mapped := make(map[fields.FieldName]string)
	for _, result := range results {
		text, _ := result.NormalizedText()
		mapped[result.Name] = text
	}

	// 曲名, 日付, ALBUM, Lv, 全ノーツ数, PERFECT, GREAT, NICE, BAD, MISS, NICE以下, COMBO
	var texts []string
	texts = append(texts, mapped[fields.Title])
	texts = append(texts, dateString())
	texts = append(texts, "") // ALBUM
	texts = append(texts, mapped[fields.Lv])
	texts = append(texts, "") // ノーツ数
	texts = append(
		texts,
		mapped[fields.Perfect],
		mapped[fields.Great],
		mapped[fields.Nice],
		mapped[fields.Bad],
		mapped[fields.Miss],
	)
	texts = append(texts, "") // NICE以下
	texts = append(texts, mapped[fields.Combo])
	csvString := strings.Join(texts, "\t")
	csvString += "\n"

	// TODO: 順序制御を切り離してまともにする
	// TODO: 真面目にCSV文字列を作る

	_, err := e.w.Write([]byte(csvString))
	if err != nil {
		return err
	}

	return nil
}

func dateString() string {
	day := time.Now()
	const layout = "2006/01/02"
	return day.Format(layout)
}
