package data

import (
	"fmt"
)

//Video структура для хранения данных о видео-файле
type Video struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
}

// VideoList структура для списка записей о видео-файлах
type VideoList struct {
	videos []Video
}

// Editable интерфейс для работы со списком записей
type Editable interface {
	GetVideos() []Video
	AddVideo(video Video) int
	DeleteVideo(id int) error
}

// NewVideoList конструктор списка записей
func NewVideoList() *VideoList {
	return &VideoList{}
}

// GetVideos возвращает список видео
func (cl *VideoList) GetVideos() []Video {
	return cl.videos
}

// AddVideo добавляет видео video в конец списка и возвращает id
func (cl *VideoList) AddVideo(video Video) int {
	id := len(cl.videos) + 1
	video.ID = id
	cl.videos = append(cl.videos, video)
	return id
}

// DeleteVideo удаляет видео по id
func (cl *VideoList) DeleteVideo(id int) error {
	if id < 1 || id > len(cl.videos) {
		return fmt.Errorf("Incorrect ID")
	}
	id--
	cl.videos = append(cl.videos[:id], cl.videos[id+1:]...)
	return nil
}
