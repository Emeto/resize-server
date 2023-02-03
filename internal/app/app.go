package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"resize-server/config"
	v1 "resize-server/internal/controller/http/v1"
	"resize-server/internal/usecase"
	"resize-server/pkg/http"
	"resize-server/pkg/logger"
	"syscall"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	scalingUseCase := usecase.New(cfg.Presets, cfg.App.InterFunc)

	handler := gin.New()
	v1.NewRouter(handler, l, scalingUseCase)
	httpServer := http.New(handler, http.Port(cfg.HTTP.Port))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err := <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	err := httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
