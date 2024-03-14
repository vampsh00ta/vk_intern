package repository

import (
	"gorm.io/gorm"
)

type Repository interface {
	Actor
	Film
	Customer
	Tx
}
type Pg struct {
	client *gorm.DB
}

func New(client *gorm.DB) Repository {
	return &Pg{
		client,
	}
}
