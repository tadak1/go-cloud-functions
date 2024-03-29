package image

import (
	"encoding/json"
	"google.golang.org/api/option"
	"log"
	"net/http"
	"os"
)

var (
	apiKey string
)

func init() {
	apiKey = os.Getenv("API_KEY")
}

func Upload(writer http.ResponseWriter, request *http.Request) {
	file, header, err := request.FormFile("file")
	if err != nil {
		log.Printf("%s", err.Error())
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	client, err := NewStorageClient(option.WithAPIKey(apiKey))
	if err != nil {
		log.Printf("%s", err.Error())
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	url, err := client.Upload(file, header)
	if err != nil {
		log.Printf("%s", err.Error())
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(map[string]string{
		"url": url,
	})
	if err != nil {
		log.Printf("%s", err.Error())
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
}
