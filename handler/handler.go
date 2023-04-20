package handler

import (
	"bytes"

	"github.com/adithya/pockethealth/pkg/dicomfile"
	"github.com/adithya/pockethealth/pkg/fileutils"
	"github.com/suyashkumar/dicom/pkg/tag"
)

func GetHeaderAttributeHandler(fileID string, expectedTag tag.Tag) (string, error) {
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

func GetPNGImageHandler(fileID string) (*bytes.Buffer, error) {
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
