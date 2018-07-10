package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/ComputePractice2018/videohosting/backend/data"
)

//GetVideos обрабатывает запросы на получение списка
func GetVideos(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(data.GetVideos())
	if err != nil {
		message := fmt.Sprintf("Unable to encode POST data: %v", err)
		http.Error(w, message, http.StatusInternalServerError)
		log.Println(message)
		return
	}
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
	id := data.AddVideo(video)
	w.Header().Add("Location", r.URL.String()+"/"+strconv.Itoa(id))
	w.WriteHeader(http.StatusCreated)
}

// DeleteVideo обрабатывает DELETE запрос
func DeleteVideo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		message := fmt.Sprintf("Incorrect ID: %v", err)
		http.Error(w, message, http.StatusBadRequest)
		log.Println(message)
		return
	}

	err = data.DeleteVideo(id)
	if err != nil {
		message := fmt.Sprintf("Incorrect ID: %v", err)
		http.Error(w, message, http.StatusBadRequest)
		log.Println(message)
		return
	}
	w.Header().Add("Location", r.URL.String())
	w.WriteHeader(http.StatusNoContent)
}

//UploadVideo обрабатывает POST запрос на загрузку файла
func UploadVideo(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()
	fmt.Fprintf(w, "%v", handler.Header)

	f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()
	io.Copy(f, file)
}
