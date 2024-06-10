package photos

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type ListPhotosResponse struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Data    []string `json:"data"`
}

type GetPhotoResponse struct {
	Status int    `json:"status"`
	Data   string `json:"data"`
}

func ListAlbums(subDir string) ListPhotosResponse {
	var photos []string
	f, err := os.Open(subDir)
	if err != nil {
		return ListPhotosResponse{
			Status:  500,
			Message: "error opening photos",
			Data:    nil,
		}
	}

	files, err := f.ReadDir(0)
	if err != nil {
		return ListPhotosResponse{
			Status:  500,
			Message: "error opening photos",
			Data:    nil,
		}
	}

	for _, v := range files {
		if v.IsDir() {
			photos = append(photos, v.Name())
		}
	}

	return ListPhotosResponse{
		Status:  200,
		Message: "success",
		Data:    photos,
	}
}

func ListPhotos(subDir string) ListPhotosResponse {
	var photos []string
	f, err := os.Open(fmt.Sprintf("test/%s", subDir))
	if err != nil {
		return ListPhotosResponse{
			Status:  500,
			Message: "error opening photos",
			Data:    nil,
		}
	}

	files, err := f.ReadDir(0)
	if err != nil {
		return ListPhotosResponse{
			Status:  500,
			Message: "error opening photos",
			Data:    nil,
		}
	}

	for _, v := range files {
		if !v.IsDir() {
			photos = append(photos, v.Name())
		}
	}

	return ListPhotosResponse{
		Status:  200,
		Message: "success",
		Data:    photos,
	}
}

func GetPhoto(album string, photo string) GetPhotoResponse {
	var photoB64 string
	dir := fmt.Sprintf("test/%s/%s", album, photo)

	bytes, err := ioutil.ReadFile(dir)
	if err != nil {
		return GetPhotoResponse{
			Status: 200,
			Data:   err.Error(),
		}
	}

	photoB64 = encodeToBase64(bytes)
	fmt.Printf("made base64: %s", photoB64)

	return GetPhotoResponse{
		Status: 200,
		Data:   photoB64,
	}
}

func encodeToBase64(data []byte) string {
	var result string

	mimeType := http.DetectContentType(data)
	fmt.Printf("detected mimeType of %s", mimeType)

	switch mimeType {
	case "image/jpeg":
		result += "data:image/jpeg;base64,"
	case "image/png":
		result += "data:image/png;base64,"
	}

	photoB64 := base64.StdEncoding.EncodeToString(data)

	// writer := bytes.NewBufferString(result)
	// encoder := base64.NewEncoder(base64.StdEncoding, writer)
	// encoder.Write(data)

	return photoB64
}
