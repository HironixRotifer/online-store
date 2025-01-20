package main

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	httpserver "github.com/HironixRotifer/online-store/internal/api/http"
	"github.com/HironixRotifer/online-store/internal/config"
	"github.com/rs/zerolog/log"
)

var wg sync.WaitGroup

func main() {
	cfg := config.MustLoadConfig("")

	server := httpserver.NewServerHTTP(cfg.PortHTTP)
	server.Serve()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT)

	<-signalChan
	log.Info().Msg("Service gracefull stop started")

	wg.Add(1)
	server.GracefullShutdown(&wg)

	wg.Wait()
}
