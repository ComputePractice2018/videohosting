package data

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type DBVideoList struct {
	db *gorm.DB
}

func NewDBVideoList(connection string) (*DBVideoList, error) {
	db, err := gorm.Open("mysql", connection)
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&Video{})
	return &DBVideoList{db: db}, db.Error
}

func (cl *DBVideoList) Close() {
	cl.db.Close()
}

func (cl *DBVideoList) GetVideos() []Video {
	var videos []Video
	cl.db.Find(&videos)
	return videos
}

func (cl *DBVideoList) AddVideo(video Video) int {
	cl.db.Create(&video)
	return video.ID
}

func (cl *DBVideoList) DeleteVideo(id int) error {
	var videos []Video
	cl.db.Where("id = ?", id).Find(&videos)
	if len(videos) == 0 {
		return fmt.Errorf("can't find record #%d", id)
	}
	cl.db.Delete(&videos[0])
	return cl.db.Error
}
