package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/ComputePractice2018/videohosting/backend/data"
	"github.com/ComputePractice2018/videohosting/backend/server"
)

func main() {
	//name := flag.String("name", "Michael", "имя для приветствия")
	//flag.Parse()

	videolist := data.NewVideoList()

	router := mux.NewRouter()
	router.HandleFunc("/api/videohosting/videos", server.GetVideos(videolist)).Methods("GET")
	router.HandleFunc("/api/videohosting/videos", server.AddVideo(videolist)).Methods("POST")
	router.HandleFunc("/api/videohosting/videos/{id}", server.DeleteVideo(videolist)).Methods("DELETE")
	router.HandleFunc("/api/videohosting/upload", server.UploadVideo).Methods("POST")
	router.HandleFunc("/api/videohosting/videos/mp4/{name}", server.GetVideo).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
