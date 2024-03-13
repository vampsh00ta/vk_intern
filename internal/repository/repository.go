package repository

import "vk_intern/pkg/client"

type Repository interface {
}
type Pg struct {
	client client.Client
}

func New(client client.Client) Repository {
	return &Pg{
		client,
	}
}
