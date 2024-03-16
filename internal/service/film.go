package service

import (
	"context"
	"vk/internal/repository/models"
)

type Film interface {
	AddFilm(ctx context.Context, film models.Film) (int, error)
	DeleteFilm(ctx context.Context, id int) error

	GetFilms(ctx context.Context, sortBy, orderBy string) ([]models.Film, error)
	GetFilmsByTitle(ctx context.Context, title string, sortBy, orderBy string) ([]models.Film, error)
	GetFilmsByActorName(ctx context.Context, name string, sortBy, orderBy string) ([]models.Film, error)
}

func (s service) DeleteFilm(ctx context.Context, id int) error {

	if err := s.repo.DeleteFilmByID(ctx, id); err != nil {
		return nil
	}
	return nil
}
func (s service) GetFilmsByTitle(ctx context.Context, title string, sortBy, orderBy string) ([]models.Film, error) {
	films, err := s.repo.GetFilmsByTitle(ctx, title, sortBy, orderBy)
	if err != nil {
		return nil, nil
	}
	return films, nil
}
func (s service) GetFilmsByActorName(ctx context.Context, name string, sortBy, orderBy string) ([]models.Film, error) {
	films, err := s.repo.GetFilmsByActorName(ctx, name, sortBy, orderBy)
	if err != nil {
		return nil, nil
	}
	return films, nil
}

func (s service) AddFilm(ctx context.Context, film models.Film) (int, error) {
	ctx = s.repo.Begin(ctx)
	defer s.repo.Commit(ctx)
	id, err := s.repo.AddFilm(ctx, film)
	if err != nil {
		s.repo.Rollback(ctx)
		return -1, nil
	}
	return id, nil
}

func (s service) GetFilms(ctx context.Context, sortBy, orderBy string) ([]models.Film, error) {
	films, err := s.repo.GetFilms(ctx, sortBy, orderBy)
	if err != nil {
		return nil, nil
	}
	return films, nil
}
