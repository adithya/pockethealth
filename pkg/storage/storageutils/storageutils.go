package storageutils

import "github.com/google/uuid"

// Using DICOM UIDs to generate blob names may not be secure
// A better option is to use system-generated identifiers in order to avoid inadvertent PII leaks
// See here: https://github.com/microsoft/dicom-server/discussions/1561%5D
func GenerateBlobUUID() string {
	return uuid.New().String()
}
