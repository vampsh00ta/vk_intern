package service

import (
	"context"
	"vk/internal/repository/models"
)

type Film interface {
	AddFilm(ctx context.Context, film models.Film) (int, error)
	GetFilms(ctx context.Context) ([]models.Film, error)
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

func (s service) GetFilms(ctx context.Context) ([]models.Film, error) {
	films, err := s.repo.GetFilms(ctx)
	if err != nil {
		return nil, nil
	}
	return films, nil
}
