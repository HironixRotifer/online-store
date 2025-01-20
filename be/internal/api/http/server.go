package http

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
)

const (
	shutDownTimeout = 10 * time.Second
)

type ServerHTTP struct {
	Handler *gin.Engine

	port   int
	server *http.Server
}

func NewServerHTTP(port int) *ServerHTTP {
	handler := gin.Default()

	addr := fmt.Sprintf(":%v", port)
	server := &http.Server{Addr: addr, Handler: handler}

	return &ServerHTTP{
		port:   port,
		server: server,

		Handler: handler,
	}
}

func (s *ServerHTTP) Serve() {
	go func() {
		if err := s.server.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				log.Fatal().Err(err).Msg("failed to start http server")
			}
		}
	}()

	log.Info().Msg("http server is started!")
}

func (s *ServerHTTP) GracefullShutdown(wg *sync.WaitGroup) {
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), shutDownTimeout)
		defer cancel()
		defer wg.Done()

		if err := s.server.Shutdown(ctx); err != nil {
			log.Error().Err(err).Msg("Server forced to shutdown")
		} else {
			log.Info().Msg("HTTP stopped")
		}
	}()
}
