package health

import (
	"encoding/json"
	"github.com/KouT127/go-cloud-functions/common"
	"net/http"
)

func Check(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	err := json.NewEncoder(writer).Encode(map[string]string{
		"message": "ok",
	})
	if err != nil {
		common.HandleError(writer)
		return
	}
}
