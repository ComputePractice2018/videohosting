package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ComputePractice2018/videohosting/backend/data"
)

//VideosHandler обрабатывает все запросы к /api/videohosting/videos
func VideosHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		GetVideos(w, r)
		return
	}
	if r.Method == "POST" {
		AddVideo(w, r)
		return
	}

	message := fmt.Sprintf("Method %s is not allowed", r.Method)
	http.Error(w, message, http.StatusInternalServerError)
	log.Println(message)
}

//GetVideos обрабатывает запросы на получение списка
func GetVideos(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request method: %s", r.Method)
	binaryData, err := json.Marshal(data.VideoList)
	if err != nil {
		message := fmt.Sprintf("JSON marshaling error: %v", err)
		http.Error(w, message, http.StatusInternalServerError)
		log.Println(message)
		return
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(binaryData)
	if err != nil {
		message := fmt.Sprintf("Handler write error: %v", err)
		http.Error(w, message, http.StatusInternalServerError)
		log.Println(message)
	}
	return
}

// AddVideo обрабатывает POST запрос
func AddVideo(w http.ResponseWriter, r *http.Request) {
	var video data.Video
	err := json.NewDecoder(r.Body).Decode(&video)
	if err != nil {
		message := fmt.Sprintf("Unable to decode POST data: %v", err)
		http.Error(w, message, http.StatusUnsupportedMediaType)
		log.Println(message)
		return
	}
	log.Printf("%+v", video)
	w.WriteHeader(http.StatusCreated)
}
