package handler

import (
	"net/http"

	"github.com/adithya/pockethealth/pkg/storage"
)

// Accept and store an uploaded DICOM file
// POST - localhost:8080/dicom/
func UploadDicomFileHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	file, _, err := r.FormFile("dicomFile")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileUUID, err := storage.Put(file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte(*fileUUID))
	w.WriteHeader(http.StatusOK)
	return
}
