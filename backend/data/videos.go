package data

import (
	"fmt"
)

//Video структура для хранения данных о видео-файле
type Video struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
}

//videos хранит список видео
var videos []Video

//GetVideos возвращает список видео
func GetVideos() []Video {
	return videos
}

//AddVideo добавляет видео video в конец списка и возвращает id
func AddVideo(video Video) int {
	id := len(videos)
	videos = append(videos, video)
	return id
}

//DeleteVideo удаляет видео по id
func DeleteVideo(id int) error {
	if id < 0 || id >= len(videos) {
		return fmt.Errorf("Incorrect ID")
	}
	videos = append(videos[:id], videos[id+1:]...)
	return nil
}
