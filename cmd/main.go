package main

import (
	"github.com/KouT127/go-cloud-functions/health"
	"github.com/KouT127/go-cloud-functions/image"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", health.Check)
	http.HandleFunc("/images/upload", image.UploadHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
