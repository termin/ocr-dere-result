package export

import (
	"encoding/csv"
	"fmt"
	"io"
	"time"

	"github.com/termin/ocr-dere-result/fields"
	"go.uber.org/zap"
)

type CSVExporter struct {
	w io.Writer
}

func NewCSVExporter(w io.Writer) *CSVExporter {
	return &CSVExporter{w: w}
}

func (e *CSVExporter) Export(result *fields.Result) error {
	if !result.IsSuccessful() {
		return fmt.Errorf("results is incomplete")
	}

	mapped := make(map[fields.FieldName]string)
	for _, result := range result.Fields {
		text, _ := result.NormalizedText()
		mapped[result.Name] = text
	}

	// 曲名, 日付, ALBUM, Lv, 全ノーツ数, PERFECT, GREAT, NICE, BAD, MISS, NICE以下, COMBO
	var texts []string

	origDateTime, err := result.DateTime()
	if err != nil {
		zap.S().Debug(err)
		origDateTime = time.Now()
	}

	// TODO: 順序制御を切り離してまともにする
	texts = []string{
		mapped[fields.Title],
		dateFormat(origDateTime),
		"", // ALBUM
		mapped[fields.Lv],
		"", // ノーツ数
		mapped[fields.Perfect],
		mapped[fields.Great],
		mapped[fields.Nice],
		mapped[fields.Bad],
		mapped[fields.Miss],
		"", // NICE以下
		mapped[fields.Combo],
	}

	w := csv.NewWriter(e.w)
	w.Comma = '\t'

	err = w.Write(texts)
	if err != nil {
		return err
	}

	w.Flush()
	if w.Error() != nil {
		return err
	}

	return nil
}

func dateFormat(date time.Time) string {
	const layout = "2006/01/02"
	return date.Format(layout)
}
