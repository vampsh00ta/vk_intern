package service

import "vk_intern/internal/repository"

type Service interface {
}
type service struct {
	repo repository.Repository
}

func New(repo repository.Repository) Service {
	return &service{
		repo: repo,
	}
}
