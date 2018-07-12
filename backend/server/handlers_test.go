package server

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
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
func TestFileHandlers(t *testing.T) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	// Замечание: необходим файл example.mp4 в директории server
	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", "example.mp4")
	if err != nil {
		t.Error("Error writing to buffer")
	}

	fh, err := os.Open("./example.mp4")
	if err != nil {
		t.Error("Error opening file")
	}
	defer fh.Close()

	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		t.Error("Error copying file")
	}

	bodyWriter.Close()

	cl := data.NewVideoList()
	router := NewRouter(cl)

	req, err := http.NewRequest("POST", "/api/videohosting/upload", bodyBuf)
	req.Header.Set("Content-Type", bodyWriter.FormDataContentType())

	if err != nil {
		t.Error("Error creating request")
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	resp := w.Result()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected 201 code (gotten: %d)", resp.StatusCode)
	}
	if resp.Header.Get("Location") !=
		req.URL.String()[:len(req.URL.String())-7]+"/mp4/example.mp4" {
		t.Error("Expected another location")
	}

	req, err = http.NewRequest("GET", "/api/videohosting/videos/mp4/example.mp4", nil)
	if err != nil {
		t.Fatal(err)
	}

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	resp = w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Error("Expected 200 code")
	}
	if resp.Header.Get("Content-Type") !=
		"video/mp4; charset=utf-8" {
		t.Error("Expected mp4 MIME-type")
	}

	os.Remove("./mp4/example.mp4")
	os.Remove("./mp4")
}
