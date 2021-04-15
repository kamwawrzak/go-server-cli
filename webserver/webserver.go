package webserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Struct storing filepath
type FilesHandler struct {
	FilePath string
}

// function handling GET request and serving passed file
func (f *FilesHandler) HomePageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/html")
		http.ServeFile(w, r, f.FilePath)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// function starting server
func StartServer(file string) {
	newFilesHandler := &FilesHandler{FilePath: file}
	m := mux.NewRouter()
	m.HandleFunc("/", newFilesHandler.HomePageHandler).Methods("GET")
	http.Handle("/", m)
	port := ":5000"
	address := "http//:localhost" + port
	fmt.Printf("You can visit the page on: %s", address)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
