package process

import (
	"encoding/json"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"log"
	"os"

	"github.com/termin/ocr-dere-result/export"
	"github.com/termin/ocr-dere-result/fields"
	"github.com/termin/ocr-dere-result/imageclip"
	"github.com/termin/ocr-dere-result/ocr"
)

func LoadFields(configFilepath string) ([]fields.Field, error) {
	bytes, err := ioutil.ReadFile(configFilepath)
	if err != nil {
		log.Println("cannot load config", err)
		return nil, err
	}

	var fs []fields.Field
	err = json.Unmarshal(bytes, &fs)
	if err != nil {
		log.Println("failed to parse json", err)
		return nil, err
	}

	return fs, nil
}

// TODO: ocr clientを外から与える
func Do(fs []fields.Field, filePath string) error {
	sourceImageFile, err := os.Open(filePath)
	if err != nil {
		log.Println("open error:", err)
		return err
	}
	defer sourceImageFile.Close()

	sourceImage, _, err := image.Decode(sourceImageFile)
	if err != nil {
		log.Println("decode error:", err)
		return err
	}

	var results fields.Results
	for _, field := range fs {
		clipped, err := imageclip.Clip(sourceImage, field)
		if err != nil {
			log.Println("failed to clip", filePath, err)
			// TODO: 具体的にどこで起きたかを追記したい
			return err
		}

		text, err := ocr.Request(clipped)
		result := fields.NewResultByField(field)
		result.Text = text
		fmt.Printf("field: %v, text: %v\n", field.Name, text)
		results = append(results, result)
	}

	csvExport := export.NewCSVExporter(os.Stdout)
	csvExport.Export(results)

	return nil
}
