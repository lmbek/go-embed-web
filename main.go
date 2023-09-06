package main

import (
	"embed"
	"net/http"
)

//go:embed frontend/*
var content embed.FS

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Serve other files by prepending "/frontend" to the path.
		r.URL.Path = "/frontend" + r.URL.Path
		http.FileServer(http.FS(content)).ServeHTTP(w, r)
	})

	port := "9950"
	addr := ":" + port
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		panic(err)
	}
}
