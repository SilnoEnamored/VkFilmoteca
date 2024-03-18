package models

import (
	"time"
)

type Actor struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Gender    string    `json:"gender,omitempty"`
	BirthDate time.Time `json:"birth_date"`
}

type Movie struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	ReleaseDate time.Time `json:"release_date"`
	Rating      float64   `json:"rating"`
	Actors      []Actor   `json:"actors,omitempty" pg:"many2many:movie_actors"`
}

type MovieActor struct {
	MovieID int
	ActorID int
}
