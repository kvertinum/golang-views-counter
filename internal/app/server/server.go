package server

import (
	"fmt"
	"net/http"

	"github.com/Kvertinum01/views-counter/internal/app/store"
)

type Server struct {
	rst    *store.RedisStore
	imgDir string
}

func SetupServer(config *Config) error {
	rst := store.NewRedisStore(config.RedisConfig)
	if err := rst.ConfigureStore(); err != nil {
		return err
	}

	s := &Server{rst: rst, imgDir: "./images"}

	s.configureRoutes()

	fmt.Printf("Server started at %s", config.ServerAddr)

	return http.ListenAndServe(config.ServerAddr, nil)
}

func (s *Server) configureRoutes() {
	http.Handle("/counter/", http.HandlerFunc(s.handleRandomImage))
}
