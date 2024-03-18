package service

import (
	"context"
	"filmoteca/internal/models"
	"github.com/go-pg/pg/v10"
)

type MovieService struct {
	DB *pg.DB
}

func NewMovieService(db *pg.DB) *MovieService {
	return &MovieService{
		DB: db,
	}
}

func (s *MovieService) GetAllMovies(ctx context.Context) ([]models.Movie, error) {
	var movies []models.Movie
	err := s.DB.Model(&movies).Context(ctx).Select()
	if err != nil {
		return nil, err
	}
	return movies, nil
}

func (s *MovieService) CreateMovie(ctx context.Context, movie *models.Movie) error {
	_, err := s.DB.Model(movie).Context(ctx).Insert()
	return err
}

func (s *MovieService) UpdateMovie(ctx context.Context, id int, movie *models.Movie) error {
	movie.ID = id // Убедитесь, что ID устанавливается корректно
	_, err := s.DB.Model(movie).WherePK().Context(ctx).Update()
	return err
}

func (s *MovieService) DeleteMovie(ctx context.Context, id int) error {
	_, err := s.DB.Model(&models.Movie{}).Where("id = ?", id).Context(ctx).Delete()
	return err
}
