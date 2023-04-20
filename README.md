# PocketHealth Back-End Challange

## Overview

As described in the challenge instruction, this project exposes an API that allows users to perform three actions:

- Upload DICOM Files
- Retrieve DICOM Header Attributes based on DICOM tags
- Convert DICOM files into a PNG

## Development

### Required Dependencies

- [Golang](https://go.dev/doc/install) (>1.13)

### Getting Started

1. Clone the project

```bash
git clone https://github.com/adithya/pockethealth.git
```

2. Run the project
```bash
go run main.go
```

3. Verify API is running

**Endpoint**  `GET /GetVersion`

**Request:**
```bash
curl --request GET \
  --url http://localhost:8080/GetVersion
```

**Response:**
```json
0.1
```

## API Docs

### Upload a DICOM File

**Endpoint** `POST /dicom`

**Request:**
```bash
curl --request POST \
  --header 'content-type: multipart/form-data' \
  -F dicomFile=@0001.dcm \
  --url http://localhost:8080/dicom
```

**Response:**
```bash
0c4be1b0-4f0b-40ca-8d9c-5a141860b652
```

### Get a PNG for a DICOM File

**Endpoint** `GET /dicom/{id}/image?format=png`

**Request:**
```bash
curl --request GET \
  --url "http://localhost:8080/dicom/0c4be1b0-4f0b-40ca-8d9c-5a141860b652/image?format=png" \
  --output file.png
```

**Response:**
PNG IMAGE 

### Get a DICOM Header Attribute for DICOM file based on header tag

**Endpoint** `GET /dicom/{id}/headerAttribute?tag=(0010,0010)`

**Request:**
```bash
curl --request GET \
  --url "http://localhost:8080/dicom/0c4be1b0-4f0b-40ca-8d9c-5a141860b652/headerAttribute?tag=(0010,0010)"
```

**Response:**
```bash
[NAYYAR^HARSH]
```

## Next Steps