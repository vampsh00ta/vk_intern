package vk

import (
	"context"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlog "gorm.io/gorm/logger"
	"os"
	"os/signal"
	"syscall"
	"vk/config"
	"vk/internal/repository"
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
	ctx := context.Background()
	//pg, err := client.NewPostgresClient(ctx, 5, cfg.PG)
	if err != nil {
		//log.Fatal(fmt.Errorf("vk - Run - postgres.New: %w", err))
	}
	//defer pg.Close()
	repo := repository.New(db)
	ctx = repo.Begin(ctx)
	//defer repo.Commit(ctx)
	//f, err := repo.AddFilm(ctx, models.Film{Description: "asd", Title: "asdasd", Rating: 5, ReleaseDate: time.Time{}})
	//fmt.Println(f, err)
	//f1, err := repo.AddActor(ctx, models.Actor{Middlename: "asd"})
	//fmt.Println(f1, err)
	//f2, err := repo.GetFilms(ctx)
	//fmt.Println(f2[0].Actors, err)
	//a1, err := repo.GetActors(ctx)
	//fmt.Println(a1, err)
	customer, err := repo.GetCustomerById(ctx, 1)
	fmt.Println(customer)
	repo.Commit(ctx)
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
