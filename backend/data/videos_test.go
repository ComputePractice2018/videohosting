package data

import "testing"

var testVideos = []Video{
	{
		ID:          1,
		Title:       "Видео 1",
		Description: "Описание 1",
		Link:        "Ссылка 1",
	},
	{
		ID:          2,
		Title:       "Видео 2",
		Description: "Описание 2",
		Link:        "Ссылка 2",
	},
}

func TestAddVideo(t *testing.T) {
	cl := NewVideoList()

	cl.AddVideo(testVideos[0])

	if cl.GetVideos()[0] != testVideos[0] {
		t.Errorf("AddVideo is not working")
	}
}

func TestDeleteVideo(t *testing.T) {
	cl := NewVideoList()
	cl.AddVideo(testVideos[0])
	cl.AddVideo(testVideos[1])

	err := cl.DeleteVideo(1)

	if cl.GetVideos()[0] != testVideos[1] {
		t.Errorf("DeleteContact is not working")
	}
	if err != nil {
		t.Errorf("Unexpected DeleteContact error")
	}

	err = cl.DeleteVideo(-1)
	if err == nil {
		t.Errorf("Nothandled out of range error")
	}
}
