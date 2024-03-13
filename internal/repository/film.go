package repository

import (
	"context"
	"vk/internal/repository/models"
)

type Film interface {
	AddFilm(ctx context.Context, film models.Film) error
	DeleteFilmByID(ctx context.Context, filmId int) error
	DeleteFilmByTitle(ctx context.Context, filmId int) error
	ChangeFilmByID(ctx context.Context, film models.Film) error
	ChangeFilmByFullName(ctx context.Context, filmId int) error
	AddActorsToFilms(ctx context.Context, filmId int, actor ...models.Actor) error

	GetFilmById(ctx context.Context, filmId string) (models.Film, error)
	GetFilmByTitle(ctx context.Context, filmTitle string) (models.Film, error)
	GetFilms(ctx context.Context) ([]models.Film, error)
}
