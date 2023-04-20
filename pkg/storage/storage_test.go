package storage_test

import (
	"mime/multipart"
	"os"
	"testing"

	"github.com/adithya/pockethealth/pkg/storage"
	"github.com/stretchr/testify/assert"
)

const (
	filePath = "../../testdata/MRI/PA000001/ST000001/SE000001/IM000001"
)

func TestPut(t *testing.T) {
	file, err := os.Open(filePath)
	assert.NoError(t, err)
	defer file.Close()

	type args struct {
		file multipart.File
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "dicom file should upload",
			args:    args{file: file},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fileName, err := storage.Put(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("Put() error = %v, wantErr %v", err, tt.wantErr)
			}

			os.Remove("bucket/" + *fileName)
		})

	}
}

func TestGet(t *testing.T) {
	fileId := "testfile"

	type args struct {
		fileId string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "should return correct path",
			args:    args{fileId: fileId},
			want:    "bucket/" + fileId,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := storage.Get(tt.args.fileId)
			if got != tt.want {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
