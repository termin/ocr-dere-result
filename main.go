package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/termin/ocr-dere-result/process"
	"go.uber.org/zap"
)

func main() {
	verbose := flag.Bool("verbose", false, "output debug log")
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "requires at least 1 argument\n")
		os.Exit(1)
	}

	if *verbose {
		fmt.Println("setenv dev")
		os.Setenv("GO_ENV", "dev")
	}

	// TODO: --verboseの場合に

	var logger *zap.Logger
	env := os.Getenv("GO_ENV")
	var err error
	if env == "dev" {
		logger, err = zap.NewDevelopment()
	} else {
		// 未定義もprod
		logger, err = zap.NewProduction()
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to initialize logger")
		os.Exit(1)
	}
	defer logger.Sync()
	zap.ReplaceGlobals(logger)

	zap.S().Debug(args)

	p, err := os.Executable()
	configFilepath := filepath.Join(filepath.Dir(p), "configs/coordinates.json")
	fs, err := process.LoadFields(configFilepath)
	if err != nil {
		zap.S().Fatal("failed to load config", err)
	}

	zap.S().Debug("Coordinates")
	for _, field := range fs {
		zap.S().Debug(field)
	}

	for _, filePath := range args {
		err := process.Do(fs, filePath)
		if err != nil {
			zap.S().Error(err)
		}
	}
}
