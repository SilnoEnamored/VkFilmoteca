package service

import (
	"context"
	"filmoteca/internal/models"
	"github.com/go-pg/pg/v10"
)

type ActorService struct {
	DB *pg.DB
}

func NewActorService(db *pg.DB) *ActorService {
	return &ActorService{
		DB: db,
	}
}

func (s *ActorService) GetAllActors(ctx context.Context) ([]models.Actor, error) {
	var actors []models.Actor
	err := s.DB.Model(&actors).Context(ctx).Select()
	if err != nil {
		return nil, err
	}
	return actors, nil
}

func (s *ActorService) CreateActor(ctx context.Context, actor *models.Actor) error {
	_, err := s.DB.Model(actor).Context(ctx).Insert()
	return err
}

func (s *ActorService) UpdateActor(ctx context.Context, id int, actor *models.Actor) error {
	actor.ID = id
	_, err := s.DB.Model(actor).WherePK().Context(ctx).Update()
	return err
}

func (s *ActorService) DeleteActor(ctx context.Context, id int) error {
	_, err := s.DB.Model(&models.Actor{}).Where("id = ?", id).Context(ctx).Delete()
	return err
}
