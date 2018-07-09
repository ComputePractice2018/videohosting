package main

import (
	"log"
	"net/http"

	"github.com/ComputePractice2018/videohosting/backend/server"
)

func main() {
	//name := flag.String("name", "Michael", "имя для приветствия")
	//flag.Parse()

	http.HandleFunc("/api/videohosting/videos", server.VideosHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))

	http.ListenAndServe(":8080", nil)

}
