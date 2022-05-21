package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/termin/ocr-dere-result/process"
)

func main() {
	flag.Parse()
	args := flag.Args()
	fmt.Println(args)
	// TODO: オプションの追加

	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "requires at least 1 argument\n")
		os.Exit(1)
	}

	p, err := os.Executable()
	configFilepath := filepath.Join(filepath.Dir(p), "configs/coordinates.json")
	fs, err := process.LoadFields(configFilepath)
	if err != nil {
		log.Println("failed to load config", err)
		os.Exit(1)
	}

	for _, field := range fs {
		fmt.Println(field)
		fmt.Println()
	}

	for _, filePath := range args {
		err := process.Do(fs, filePath)
		if err != nil {
			log.Println(err)
		}
	}
}
