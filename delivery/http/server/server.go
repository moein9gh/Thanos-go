package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/thanos-go/log"
)

func NewServer(router *echo.Echo, port int, gracefulShutdown time.Duration) *Server {
	return &Server{
		addr:             fmt.Sprintf("%s:%v", os.Getenv("HOST_IP"), port),
		router:           router,
		gracefulShutdown: gracefulShutdown,
	}
}

type Server struct {
	addr             string
	router           *echo.Echo
	gracefulShutdown time.Duration
}

func (s *Server) StartListening() {

	go func() {
		if err := s.router.Start(s.addr); err != nil && err != http.ErrServerClosed {
			log.Fatal("server start", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Info("server shutting down in %s...", s.gracefulShutdown)
	c, cancel := context.WithTimeout(context.Background(), s.gracefulShutdown)
	defer cancel()
	if err := s.router.Shutdown(c); err != nil {
		log.Fatal("server shutdown", err)
	}

	<-c.Done()
	log.Info("Good Luck!")
}
