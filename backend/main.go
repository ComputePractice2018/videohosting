package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/ComputePractice2018/videohosting/backend/server"
)

func main() {
	//name := flag.String("name", "Michael", "имя для приветствия")
	//flag.Parse()

	router := mux.NewRouter()
	router.HandleFunc("/api/videohosting/videos", server.GetVideos).Methods("GET")
	router.HandleFunc("/api/videohosting/videos", server.AddVideo).Methods("POST")
	router.HandleFunc("/api/videohosting/videos/{id}", server.DeleteVideo).Methods("DELETE")

	router.HandleFunc("/api/videohosting/upload", server.UploadVideo).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
