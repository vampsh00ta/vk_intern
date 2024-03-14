package models

import (
	"time"
)

type Actor struct {
	Id         int       `json:"id" db:"id"`
	Name       string    `json:"name" db:"name"`
	Middlename string    `json:"middlename" db:"middlename"`
	Surname    string    `json:"surname" db:"surname"`
	BirthDate  time.Time `json:"birth_date" db:"birth_date"`
	Films      []Film    `gorm:"many2many:actor_films;"`
}

func (Actor) TableName() string {
	return "actor"
}

type Film struct {
	Id          int       `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Rating      int       `json:"rating" db:"rating"`
	ReleaseDate time.Time `json:"release_date" db:"release_date"`
	Actors      []Actor   `gorm:"many2many:actor_films;"`
}

func (Film) TableName() string {
	return "film"
}
