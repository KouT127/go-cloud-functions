package common

import "net/http"

func Handle(writer http.ResponseWriter, ) {
	writer.WriteHeader(http.StatusBadRequest)
	return
}

func HandleError(writer http.ResponseWriter) {
	writer.WriteHeader(http.StatusBadRequest)
	return
}
