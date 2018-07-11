package server

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
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
		message := fmt.Sprintf("Incorrect file: %v", err)
		http.Error(w, message, http.StatusBadRequest)
		log.Println(message)
		return
	}
	file.Close()

	err = os.MkdirAll("./mp4", 0777)
	if err != nil {
		message := fmt.Sprintf("Unable to create directory: %v", err)
		http.Error(w, message, http.StatusInternalServerError)
		log.Println(message)
		return
	}

	f, err := os.OpenFile("./mp4/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		message := fmt.Sprintf("Unable to opent file: %v", err)
		http.Error(w, message, http.StatusBadRequest)
		log.Println(message)
		return
	}

	defer f.Close()
	io.Copy(f, file)

	newpath := r.URL.String()[:len(r.URL.String())-7] + "/mp4/" + handler.Filename

	w.Header().Add("Location", newpath)
	w.WriteHeader(http.StatusCreated)

	fmt.Println(newpath)
}

//GetVideo возвращает файл видео
func GetVideo(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	name := vars["name"]

	w.Header().Add("Content-Type", "video/mp4; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	binvid, err := ioutil.ReadFile("./mp4/" + name)

	if err != nil {
		message := fmt.Sprintf("Unable to read file: %v", err)
		http.Error(w, message, http.StatusInternalServerError)
		log.Println(message)
		return
	}

	fmt.Println(len(binvid))

	_, err = w.Write(binvid)

	if err != nil {
		message := fmt.Sprintf("Unable to write file: %v", err)
		http.Error(w, message, http.StatusInternalServerError)
		log.Println(message)
		return
	}

}
