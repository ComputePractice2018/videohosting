package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/ComputePractice2018/videohosting/backend/data"
	"github.com/ComputePractice2018/videohosting/backend/server"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	connection := flag.String("connection", "videohosting:SuperSecretPassword@tcp(db:3306)/videohosting", "mysql connection string")
	flag.Parse()

	videoList, err := data.NewDBVideoList(*connection)
	defer videoList.Close()

	if err != nil {
		log.Fatalf("DB error: %+v", err)
	}

	router := server.NewRouter(videoList)

	log.Fatal(http.ListenAndServe(":8080", router))
}
