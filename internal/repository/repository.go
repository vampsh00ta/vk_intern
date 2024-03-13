package repository

import "vk/pkg/client"

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
