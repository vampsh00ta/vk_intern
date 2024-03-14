package models

import (
	"time"
)

type Actor struct {
	Id         int       `json:"id,omitempty" db:"id"`
	Name       string    `json:"name,omitempty" db:"name"`
	Middlename string    `json:"middlename,omitempty" db:"middlename"`
	Surname    string    `json:"surname,omitempty" db:"surname"`
	BirthDate  time.Time `json:"birth_date,omitempty" db:"birth_date"`
	Films      []Film    `json:"films,omitempty" gorm:"many2many:actor_films;"`
}

func (Actor) TableName() string {
	return "actor"
}

type Film struct {
	Id          int       `json:"id,omitempty" db:"id"`
	Title       string    `json:"title,omitempty" db:"title"`
	Description string    `json:"description,omitempty" db:"description"`
	Rating      int       `json:"rating,omitempty" db:"rating"`
	ReleaseDate time.Time `json:"release_date,omitempty" db:"release_date"`
	Actors      []Actor   `json:"actors,omitempty" gorm:"many2many:actor_films;"`
}

func (Film) TableName() string {
	return "film"
}
