package data

//структура для хранения данных о видео-файле
type Video struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
}

//VideoList хранит список видео
var VideoList = []Video{Video{
	Title:       "Название",
	Description: "Описание",
	Link:        "Ссылка на видео",
}}
