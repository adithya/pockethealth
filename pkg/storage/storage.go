package storage

import (
	"io"
	"mime/multipart"
	"os"

	"github.com/adithya/pockethealth/pkg/storage/storageutils"
)

type storageOperations interface {
	Put(multipart.File) error
	Get(fileId string) (os.File, error)
}

func Put(file multipart.File) (*string, error) {
	defer file.Close()

	fileName := storageutils.GenerateBlobUUID()
	dst, err := os.Create("bucket/" + fileName)
	if err != nil {
		return nil, err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		return nil, err
	}

	return &fileName, nil
}

func Get(fileId string) string {
	return "bucket/" + fileId
}
