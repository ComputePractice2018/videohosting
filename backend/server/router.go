package server

import (
	"github.com/ComputePractice2018/videohosting/backend/data"
	"github.com/gorilla/mux"
)

func NewRouter(videolist data.Editable) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/videohosting/videos", GetVideos(videolist)).Methods("GET")
	router.HandleFunc("/api/videohosting/videos", AddVideo(videolist)).Methods("POST")
	router.HandleFunc("/api/videohosting/videos/{id}", DeleteVideo(videolist)).Methods("DELETE")
	router.HandleFunc("/api/videohosting/upload", UploadVideo).Methods("POST")
	router.HandleFunc("/api/videohosting/videos/mp4/{name}", GetVideo).Methods("GET")

	return router
}
