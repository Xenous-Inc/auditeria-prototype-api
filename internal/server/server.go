package server

import (
	"auditeria-prototype/internal/utils"
	"fmt"
	"net/http"
	"time"
)

var port = 8080

type Server struct {
	port int
	mapFiles map[int]*utils.Files
}

func NewServer() *http.Server {

	mapFiles,err := utils.MustLoadFiles()
	if err != nil {
		return nil
	}

	NewServer := &Server{
		port: port,
		mapFiles: mapFiles, 
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
