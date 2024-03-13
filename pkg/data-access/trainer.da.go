package dataaccess

import (
	"anhthi-projects/pt-booking-go/db/models"
	"context"

	"github.com/uptrace/bun"
)

/*
 * Definition
 */

type ITrainerDA interface {
	Get(context.Context, int) (models.TrainerModel, error)
}

type TrainerDA struct {
	dbc *bun.DB
}

/*
 * Implement
 */

func (da *TrainerDA) Get (c context.Context, ID int) (models.TrainerModel, error) {
	trainer := models.TrainerModel {
		ID: ID,
	}

	err := da.dbc.NewSelect().Model(trainer).WherePK().Scan(c)

	if err != nil {
		return models.TrainerModel{}, err
	}

	return trainer, nil
}

func CreateTrainerDA(dbc *bun.DB) *TrainerDA {
	return &TrainerDA{dbc}
}