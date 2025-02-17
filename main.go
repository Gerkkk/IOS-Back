package main

import (
	"context"
	"fmt"
	"github.com/Gerkkk/IOS-Back/handler"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func main() {
	srv := new(Server)

	handlers := handler.NewHandler()

	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		fmt.Println("failed to run server")
	}
}
