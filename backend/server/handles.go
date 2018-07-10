package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/ComputePractice2018/videohosting/backend/data"
)

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

// DeleteVideo обрабатывает DELETE запрос
func DeleteVideo(w http.ResponseWriter, r *http.Request) {

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
