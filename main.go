package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/adithya/pockethealth/handler"
	"github.com/adithya/pockethealth/pkg/fileutils"
	"github.com/adithya/pockethealth/pkg/storage"
	"github.com/gorilla/mux"
	"github.com/suyashkumar/dicom/pkg/tag"
)

var WD string

// Accept and store an uploaded DICOM file
// POST - localhost:8080/dicom/
func uploadDicomFileHandler(w http.ResponseWriter, r *http.Request) {
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

// Return any header attributed based on a DICOM Tag as query paratmeter
// GET - localhost:8080/dicom/:id/headerAttribute?tag=(dicom tag)
// Convert the file into a PNG for browser-based viewing
// GET - localhost:8080/dicom/:id/image?format=png
func getDICOMHandler(w http.ResponseWriter, r *http.Request) {
	// Get DICOM ID
	params := mux.Vars(r)
	dicomUUID, isFound := params["id"]
	if !isFound {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get file path
	dicomFilePath := storage.Get(dicomUUID)

	// Validate if file exists
	err := fileutils.ValidateFilePath(dicomFilePath)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
	}

	// if /headerAttribute route - middleware loading dicomfile object
	if params["resource"] == "headerAttribute" {
		tagQueryParameter := r.URL.Query().Get("tag")
		group, err := strconv.ParseInt(tagQueryParameter[1:5], 16, 0)
		element, err := strconv.ParseInt(tagQueryParameter[6:10], 16, 0)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		tag := tag.Tag{Group: uint16(group), Element: uint16(element)}
		attribute, err := handler.GetHeaderAttributeHandler(dicomUUID, tag)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(attribute))
		// if /image route - middleware loading dicomefile object
	} else if params["resource"] == "image" {
		imageFormat := r.URL.Query().Get("format")
		fmt.Println(imageFormat)

		buffer, err := handler.GetPNGImageHandler(dicomUUID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if _, err := w.Write(buffer.Bytes()); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "image/jpeg")
		w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/GetVersion", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("0.1\n"))
	})

	// Accept and store DICOM file: POST - /dicom/
	r.HandleFunc("/dicom", uploadDicomFileHandler)
	// Retrieve DICOM File: GET - /dicom/{id}
	r.HandleFunc("/dicom/{id}/{resource}", getDICOMHandler)

	s := http.Server{
		Addr:         ":8080",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      r,
	}

	fmt.Printf("Starting server at :8080")
	err := s.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}
}
