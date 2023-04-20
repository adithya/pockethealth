package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/adithya/pockethealth/handler"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/GetVersion", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("0.1\n"))
	})

	// Accept and store DICOM file: POST - /dicom/
	r.HandleFunc("/dicom", handler.UploadDicomFileHandler)
	// Retrieve DICOM File: GET - /dicom/{id}
	r.HandleFunc("/dicom/{id}/{resource}", handler.GetDICOMHandler)

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
