package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ComputePractice2018/videohosting/backend/data"
)

//GetVideos обрабатывает запросы на получение списка
func GetVideos(w http.ResponseWriter, r *http.Request) {
	binaryData, err := json.Marshal(data.VideoList)
	if err != nil {
		w.Header().Add("Content-Type", "plain/text; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "JSON marshaling error: %v", data.VideoList)
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(binaryData)
	if err != nil {
		log.Printf("Handler write error: %v", err)
	}
	return
}
