package handler

import (
	"log"
	"net/http"
)

func Server() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Home)
	mux.HandleFunc("/artist/", Artist)
	log.Println("Веб-сервер запускается на http://127.0.0.1:8000")
	mux.Handle("/style/", http.StripPrefix("/style", http.FileServer(http.Dir("./style/"))))
	err := http.ListenAndServe(":8000", mux)
	log.Fatal(err)
}
