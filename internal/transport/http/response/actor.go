package response

import "vk/internal/repository/models"

type AddActor struct {
	Id int `json:"id" `
}

type GetActors struct {
	Actors []models.Actor `json:"actors"`
}
