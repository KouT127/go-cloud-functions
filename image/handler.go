package image

import (
	"encoding/json"
	"github.com/KouT127/go-cloud-functions/module"
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

func UploadHandler(writer http.ResponseWriter, request *http.Request) {
	file, header, err := request.FormFile("file")
	if err != nil {
		log.Printf("%s", err.Error())
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	client, err := module.NewStorageClient(option.WithAPIKey(apiKey))
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
