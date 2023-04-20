package fileutils

import (
	"bytes"
	"errors"
	"image"
	"image/png"
	"os"
)

func ValidateFilePath(filePath string) error {
	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		return err
	}

	return nil
}

func EncodeAndSaveImagesToPNG(images []image.Image) {
}

func EncodeAndSaveImageToPNG(img image.Image) (*bytes.Buffer, error) {
	buffer := new(bytes.Buffer)
	if err := png.Encode(buffer, img); err != nil {
		return nil, err
	}

	return buffer, nil
}
