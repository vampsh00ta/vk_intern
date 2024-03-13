package vk

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"vk/config"
	"vk/internal/repository"
	"vk/internal/service"

	//"vk/internal/transport/http/v1"
	"vk/pkg/client"
	//"vk/pkg/httpserver"
	"vk/pkg/logger"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	ctx := context.Background()
	pg, err := client.NewPostgresClient(ctx, 5, cfg.PG)
	if err != nil {
		l.Fatal(fmt.Errorf("vk - Run - postgres.New: %w", err))
	}
	defer pg.Close()
	repo := repository.New(pg)

	srvc := service.New(repo)
	srvc = srvc
	//handler := gin.New()
	//v1.NewRouter(handler, l, srvc)
	//httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	//select {
	//case s := <-interrupt:
	//	l.Info("vk - Run - signal: " + s.String())
	//case err = <-httpServer.Notify():
	//	l.Error(fmt.Errorf("vk - Run - httpServer.Notify: %w", err))
	//
	//}
	//
	//// Shutdown
	//err = httpServer.Shutdown()
	//if err != nil {
	//	l.Error(fmt.Errorf("vk - Run - httpServer.Shutdown: %w", err))
	//}

}
