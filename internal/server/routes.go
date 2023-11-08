package server

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/healthCheck", s.HealthCheck)
	r.Get("/chapter", s.GetChapterMP3)

	return r
}

func (s *Server) HealthCheck(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["health-check"] = "OK"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	w.Write(jsonResp)
}

func (s *Server) GetChapterMP3(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("number")
	number, err := strconv.Atoi(q)
	if err != nil {
		log.Fatal(err)
	}

	files := s.mapFiles[number]

	audioFile := filepath.Join(files.Audio)
	textFile := filepath.Join(files.Text)
	fmt.Println(audioFile)
	fmt.Println(textFile)
	fAduio, err := os.Open(audioFile)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Unable to open file ")
		return
	}
	defer fAduio.Close()
	fText, err := os.Open(textFile)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Unable to open file ")
		return
	}
	defer fText.Close()

	audio, err := ioutil.ReadFile(audioFile)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	encodeAudio := base64.StdEncoding.EncodeToString([]byte(audio))
	text, err := ioutil.ReadFile(textFile)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	encodeText := base64.StdEncoding.EncodeToString([]byte(text))
	var resp = make(map[string]string, 0)
	resp["audio"] = encodeAudio
	resp["text"] = encodeText
	w.Header().Set("Content-Type", "application/json")
	b, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), 422)
		return
	}
	dec, err := base64.StdEncoding.DecodeString(encodeText)
    if err != nil {
        panic(err)
    }
    f, err := os.Create("text.pdf")
    if err != nil {
        panic(err)
    }
    defer f.Close()
    if _, err := f.Write(dec); err != nil {
        panic(err)
    }
    if err := f.Sync(); err != nil {
        panic(err)
    }
	w.Write(b)
}
