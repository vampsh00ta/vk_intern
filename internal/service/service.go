package service

import (
	"vk/config"
	"vk/internal/repository"
)

type Service interface {
	Auth
	Film
}
type service struct {
	repo repository.Repository
	cfg  config.Jwt
}

func New(repo repository.Repository) Service {
	return &service{
		repo: repo,
	}
}
