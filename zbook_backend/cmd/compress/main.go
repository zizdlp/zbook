package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/zizdlp/zbook/util"
)

func main() {
	// Set up logging to output to standard error with console formatting
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// Check if the source and target image paths are provided as command-line arguments
	if len(os.Args) != 3 {
		log.Error().Msg("Please provide the source and target image paths")
		return
	}
	srcImg := os.Args[1]
	destImg := os.Args[2]

	// Read the source image into a byte array
	imageBytes, err := util.ReadImageBytes(srcImg)
	if err != nil {
		log.Error().Err(err).Msg("Failed to read image bytes")
		return
	}

	// Compress the image
	compressedImage, err := util.CompressImage(imageBytes)
	if err != nil {
		log.Error().Err(err).Msg("Failed to compress image")
		return
	}

	// Save the compressed image to the target path
	err = os.WriteFile(destImg, compressedImage, 0644)
	if err != nil {
		log.Error().Err(err).Msg("Failed to save compressed image to file")
		return
	}

	// Log information about the source and compressed image sizes
	log.Info().Msgf("Image compressed and saved to %s. Source size: %d KB, Compressed size: %d KB", destImg, len(imageBytes)/1024, len(compressedImage)/1024)
}
