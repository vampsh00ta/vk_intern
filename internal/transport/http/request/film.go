package request

import (
	"time"
	"vk/internal/repository/models"
)

type AddFilm struct {
	Title       string         `json:"title,omitempty" db:"title"`
	Description string         `json:"description,omitempty" db:"description"`
	Rating      int            `json:"rating,omitempty" db:"rating"`
	ReleaseDate time.Time      `json:"release_date,omitempty" db:"release_date"`
	Actors      []models.Actor `json:"actors,omitempty" gorm:"many2many:actor_films;"`
	//models.Film
}

type DeleteFilm struct {
	Id int `json:"id"`
}
