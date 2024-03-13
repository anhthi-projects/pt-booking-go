package services

import (
	"anhthi-projects/pt-booking-go/db/dtos"
	"anhthi-projects/pt-booking-go/db/models"
	dataaccess "anhthi-projects/pt-booking-go/pkg/data-access"
	"context"
)

/*
 * Definition
 */

type ITrainerService interface {
	Get(context.Context, int) (dtos.TrainerDTO, error)
}

type TrainerService struct {
	da dataaccess.ITrainerDA
}

/*
 * Implement
 */

 func (s *TrainerService) Get(ctx context.Context, id int) (dtos.TrainerDTO, error) {
	trainer, err := s.da.Get(ctx, id)

	if err != nil {
		return dtos.TrainerDTO{}, err
	}

	return modelToDto(trainer), nil
 }

func modelToDto(m models.TrainerModel) dtos.TrainerDTO {
	d := dtos.TrainerDTO {
		ID: m.ID,
		FirstName: m.FirstName,
		LastName: m.LastName,
		DateOfBirth: m.DateOfBirth,
		Level: m.Level,
		YearOfExp: m.YearOfExp,
	}

	return d
}

// func dtoToModel(d dtos.TrainerDTO) models.TrainerModel {
// 	m := models.TrainerModel {
// 		ID: d.ID,
// 		FirstName: d.FirstName,
// 		LastName: d.LastName,
// 		DateOfBirth: d.DateOfBirth,
// 		Level: d.Level,
// 		YearOfExp: d.YearOfExp,
// 	}

// 	return m;
// }

func CreateTrainerService(da dataaccess.ITrainerDA) *TrainerService {
	return &TrainerService{da}
}