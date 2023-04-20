package handler

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"

	"github.com/adithya/pockethealth/pkg/dicomfile"
	"github.com/adithya/pockethealth/pkg/fileutils"
	"github.com/adithya/pockethealth/pkg/storage"
	"github.com/gorilla/mux"
	"github.com/suyashkumar/dicom/pkg/tag"
)

// Return any header attributed based on a DICOM Tag as query paratmeter
// GET - localhost:8080/dicom/:id/headerAttribute?tag=(dicom tag)
// Convert the file into a PNG for browser-based viewing
// GET - localhost:8080/dicom/:id/image?format=png
func GetDICOMHandler(w http.ResponseWriter, r *http.Request) {
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
		attribute, err := getHeaderAttributeHandler(dicomUUID, tag)
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

		buffer, err := getPNGImageHandler(dicomUUID)
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

func getHeaderAttributeHandler(fileID string, expectedTag tag.Tag) (string, error) {
	tag, err := tag.Find(expectedTag)
	if err != nil {
		return "", err
	}

	// Load dicomfile
	dicomFile := dicomfile.Dicomfile{}
	dicomFile.Parse("bucket/" + fileID)

	// Get Header Attribute
	element, err := dicomFile.GetHeaderAttribute(tag.Tag)

	return element.Value.String(), nil
}

func getPNGImageHandler(fileID string) (*bytes.Buffer, error) {
	// Load dicomfile
	dicomFile := dicomfile.Dicomfile{}
	dicomFile.Parse("bucket/" + fileID)

	// Get image for dicom file
	images, err := dicomFile.GetImage()
	if err != nil {
		return nil, err
	}
	image := images[0]

	// Get PNG buffer
	buffer, err := fileutils.EncodeAndSaveImageToPNG(image)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}
