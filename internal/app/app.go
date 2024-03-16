package vk

import (
	"context"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlog "gorm.io/gorm/logger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"vk/config"
	"vk/internal/repository"
	"vk/internal/transport/http/v1"

	"vk/internal/service"
	//"vk/internal/transport/http/v1"
	//"vk/pkg/client"
	//"vk/pkg/httpserver"
	//"vk/pkg/logger"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	//l := logger.New(cfg.Log.Level)
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Moscow",
		cfg.PG.Host,
		cfg.PG.Username,
		cfg.PG.Password,
		cfg.PG.Name,
		cfg.PG.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormlog.Default.LogMode(gormlog.Info),
	})
	//l := log.Default()
	ctx := context.Background()
	//pg, err := client.NewPostgresClient(ctx, 5, cfg.PG)
	if err != nil {
		//log.Fatal(fmt.Errorf("vk - Run - postgres.New: %w", err))
	}
	repo := repository.New(db)
	ctx = repo.Begin(ctx)
	srvc := service.New(repo)

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL)
	t := v1.NewTransport(srvc)

	log.Print("Listening...")
	http.ListenAndServe(":8000", t)
	select {
	case <-interrupt:
		panic("exit")

	}
	//
	//// Shutdown
	//err = http.Shutdown()
	//if err != nil {
	//	l.Error(fmt.Errorf("vk - Run - httpServer.Shutdown: %w", err))
	//}

}
