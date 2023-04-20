// Package dicomfile provides a wrapper around external
// libraries to perform common operations on dicom files.
package dicomfile

import (
	"errors"
	"image"
	"io"

	"github.com/suyashkumar/dicom"
	"github.com/suyashkumar/dicom/pkg/tag"
)

type file interface {
	Parse(file io.Reader) (*dicom.Dataset, error)
	GetHeaderAttribute(tag string) (*dicom.Element, error)
	GetImage() ([]image.Image, error)
}

type Dicomfile struct {
	Dataset *dicom.Dataset
}

func (d *Dicomfile) Parse(filePath string) (*dicom.Dataset, error) {
	dataset, err := dicom.ParseFile(filePath, nil)
	if err != nil {
		return nil, err
	}

	d.Dataset = &dataset
	return &dataset, nil
}

func (d *Dicomfile) GetHeaderAttribute(tag tag.Tag) (*dicom.Element, error) {
	if d.Dataset == nil {
		return nil, errors.New("No dataset available for dicom file. Try parsing a dicom file first.")
	}

	element, err := d.Dataset.FindElementByTag(tag)
	if err != nil {
		return nil, err
	}

	return element, nil
}

func (d *Dicomfile) GetImage() ([]image.Image, error) {
	if d.Dataset == nil {
		return nil, errors.New("No dataset available for dicom file. Try parsing a dicom file first.")
	}

	pixelDataElement, err := d.Dataset.FindElementByTag(tag.PixelData)
	if err != nil {
		return nil, err
	}

	pixelDataInfo := dicom.MustGetPixelDataInfo(pixelDataElement.Value)

	images := make([]image.Image, len(pixelDataInfo.Frames))
	for i, fr := range pixelDataInfo.Frames {
		img, err := fr.GetImage()
		if err != nil {
			return nil, err
		}
		images[i] = img
	}

	return images, nil

}
