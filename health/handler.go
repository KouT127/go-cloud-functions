package health

import (
	"encoding/json"
	"net/http"
)

func CheckHandler(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	err := json.NewEncoder(writer).Encode(map[string]string{
		"message": "ok",
	})
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
}
