package handler

import (
	"net/http"
	"time"
)

func Run(port string, handler http.Handler) error {
	httpServer := &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return httpServer.ListenAndServe()
}
