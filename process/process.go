package process

import (
	"encoding/json"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"os"

	"github.com/termin/ocr-dere-result/export"
	"github.com/termin/ocr-dere-result/fields"
	"github.com/termin/ocr-dere-result/imageclip"
	"github.com/termin/ocr-dere-result/ocr"
	"go.uber.org/zap"
)

func LoadFields(configFilepath string) ([]fields.Field, error) {
	bytes, err := ioutil.ReadFile(configFilepath)
	if err != nil {
		return nil, err
	}

	var fs []fields.Field
	err = json.Unmarshal(bytes, &fs)
	if err != nil {
		return nil, err
	}

	return fs, nil
}

// TODO: ocr clientを外から与える
func Do(fs []fields.Field, filePath string) error {
	sourceImageFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer sourceImageFile.Close()

	sourceImage, _, err := image.Decode(sourceImageFile)
	if err != nil {
		return err
	}

	result := &fields.Result{SourceImageFile: sourceImageFile}

	for _, field := range fs {
		clipped, err := imageclip.Clip(sourceImage, field)
		if err != nil {
			return err
		}

		text, err := ocr.Request(clipped)
		resultField := fields.NewResultFieldByField(field)
		resultField.Text = text
		zap.S().Debugf("field: %v, text: %v\n", field.Name, text)

		result.AddResultField(resultField)
	}

	csvExport := export.NewCSVExporter(os.Stdout)
	csvExport.Export(result)

	return nil
}
