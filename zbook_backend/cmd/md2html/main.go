package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/zizdlp/zbook/markdown/convert"
)

func main() {
	// Set up logging to output to standard error with console formatting
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// Check if the source and target paths are provided as command-line arguments
	if len(os.Args) != 3 {
		log.Error().Msg("Please provide the source and target paths")
		return
	}
	srcDir := os.Args[1]
	destDir := os.Args[2]
	convert.ConvertFolder(srcDir, destDir)
}
