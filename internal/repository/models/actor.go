package models

import "time"

type Actor struct {
	Id     int    `json:"id,omitempty" db:"id"`
	Name   string `json:"name,omitempty" db:"name"`
	Gender bool   `json:"gender,omitempty" db:"gender"`
	//Middlename string    `json:"middlename,omitempty" db:"middlename"`
	//Surname    string    `json:"surname,omitempty" db:"surname"`
	BirthDate time.Time `json:"birth_date,omitempty" db:"birth_date"`
	Films     []Film    `json:"films,omitempty" gorm:"many2many:actor_films;"`
}

func (Actor) TableName() string {
	return "actor"
}
