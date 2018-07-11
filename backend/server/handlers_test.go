package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ComputePractice2018/videohosting/backend/data"
)

var testVideo = `{"title":"Название измененное","description":"Новое","link":"Ссылка на видео"}`

func TestCrdHandlers(t *testing.T) {
	cl := data.NewVideoList()
	router := NewRouter(cl)

	req, err := http.NewRequest("GET", "/api/videohosting/videos", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Error("Expected 200 code")
	}
	if resp.Header.Get("Content-Type") !=
		"application/json; charset=utf-8" {
		t.Error("Expected json MIME-type")
	}

	testData := strings.NewReader(testVideo)
	req, err = http.NewRequest("POST", "/api/videohosting/videos", testData)

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	resp = w.Result()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected 201 code (gotten: %d)", resp.StatusCode)
	}
	if resp.Header.Get("Location") != "/api/videohosting/videos/0" {
		t.Error("Expected another location")
	}

	if len(cl.GetVideos()) != 1 {
		t.Error("Expected new value")
	}

	req, err = http.NewRequest("DELETE", "/api/videohosting/videos/0", nil)

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	resp = w.Result()

	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("Expected 204 code (gotten: %d)", resp.StatusCode)
	}

	if len(cl.GetVideos()) != 0 {
		t.Error("Expected old value")
	}
}
