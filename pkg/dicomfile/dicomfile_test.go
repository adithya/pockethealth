package dicomfile_test

import (
	"testing"

	"github.com/adithya/pockethealth/pkg/dicomfile"
	"github.com/stretchr/testify/assert"
	"github.com/suyashkumar/dicom/pkg/tag"
)

const (
	filePath = "../../testdata/MRI/PA000001/ST000001/SE000001/IM000001"
)

func TestParse(t *testing.T) {
	dicomFile := dicomfile.Dicomfile{}
	dataset, err := dicomFile.Parse(filePath)
	assert.NoError(t, err)
	assert.NotNil(t, dataset)
}

func TestGetHeaderAttribute(t *testing.T) {
	dicomFile := dicomfile.Dicomfile{}
	dataset, err := dicomFile.Parse(filePath)
	assert.NoError(t, err)
	assert.NotNil(t, dataset)

	patientBirthNameElement, err := dicomFile.GetHeaderAttribute(tag.PatientName)
	assert.NoError(t, err)
	assert.NotNil(t, patientBirthNameElement)
	assert.Equal(t, patientBirthNameElement.Value.String(), "[NAYYAR^HARSH]")
}

func TestGetImage(t *testing.T) {
	dicomFile := dicomfile.Dicomfile{}
	dataset, err := dicomFile.Parse(filePath)
	assert.NoError(t, err)
	assert.NotNil(t, dataset)

	images, err := dicomFile.GetImage()
	assert.NoError(t, err)
	assert.NotEmpty(t, images)

	for _, image := range images {
		assert.NotNil(t, image)
	}
}
