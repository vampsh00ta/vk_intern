package response

import "vk/internal/repository/models"

type AddFilm struct {
	Id int `json:"id"`
}
type GetFilms struct {
	Films []models.Film `json:"films"`
}
