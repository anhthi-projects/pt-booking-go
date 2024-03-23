package dataaccess

import (
	"anhthi-projects/pt-booking-go/schemas/models"
	"context"

	"github.com/uptrace/bun"
)

/*
 * Definition
 */

type ITrainerDA interface {
	List(ctx context.Context) ([]models.TrainerModel, error)
	Get(ctx context.Context, ID int) (models.TrainerModel, error)
	Create(ctx context.Context, trainer models.TrainerModel) (models.TrainerModel, error)
	Update(ctx context.Context, trainer models.TrainerModel, ID int) (models.TrainerModel, error)
	Delete(ctx context.Context, ID int) error 
}

type TrainerDA struct {
	dbc *bun.DB
}

/*
 * Implement
 */

func (da *TrainerDA) List (ctx context.Context) ([]models.TrainerModel, error) {
	var trainers  []models.TrainerModel

	err := da.dbc.
		NewSelect().
		Model(&trainers).
		Scan(ctx)

	if err != nil {
		return trainers, err
	}

	return trainers, nil
}

func (da *TrainerDA) Get (ctx context.Context, ID int) (models.TrainerModel, error) {
	trainer := models.TrainerModel {
		ID: ID,
	}

	err := da.dbc.NewSelect().
		Model(&trainer).
		WherePK().
		Scan(ctx)

	if err != nil {
		return models.TrainerModel{}, err
	}

	return trainer, nil
}

func (da *TrainerDA) Create(ctx context.Context, trainer models.TrainerModel) (models.TrainerModel, error) {
	_, err := da.dbc.
		NewInsert().
		Model(&trainer).
		On("CONFLICT (id) DO UPDATE").
		Returning("*").
		Exec(ctx)

	if err != nil {
		return models.TrainerModel{}, err
	}

	return trainer, nil
}

func (da *TrainerDA) Update(ctx context.Context, trainer models.TrainerModel, ID int) (models.TrainerModel, error) {
	_, err := da.dbc.
		NewUpdate().
		Model(&trainer).
		Where("id = ?", ID).
		Returning("*").
		Exec(ctx)
		
	if err != nil {
		return models.TrainerModel{}, err
	}

	return trainer, nil
}

func (da *TrainerDA) Delete(ctx context.Context, ID int) error {
	trainer := models.TrainerModel{
		ID: ID,
	}

	_, err := da.dbc.
		NewDelete().
		Model(&trainer).
		WherePK().
		Exec(ctx)

	if err != nil {
		return err
	}

	return nil
}

func CreateTrainerDA(dbc *bun.DB) *TrainerDA {
	return &TrainerDA{dbc}
}