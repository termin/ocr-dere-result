package export

import "github.com/termin/ocr-dere-result/fields"

type Exporter interface {
	Export(fields.Results) error
}
