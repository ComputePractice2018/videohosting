package main

import (
	"log"
	"net/http"

	"github.com/ComputePractice2018/videohosting/backend/data"
	"github.com/ComputePractice2018/videohosting/backend/server"
)

func main() {
	//name := flag.String("name", "Michael", "имя для приветствия")
	//flag.Parse()

	videolist := data.NewVideoList()
	router := server.NewRouter(videolist)

	log.Fatal(http.ListenAndServe(":8080", router))
}
