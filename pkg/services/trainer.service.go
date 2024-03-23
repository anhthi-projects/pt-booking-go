package services

import (
	dataaccess "anhthi-projects/pt-booking-go/pkg/data-access"
	"anhthi-projects/pt-booking-go/schemas/dtos"
	"anhthi-projects/pt-booking-go/schemas/models"
	"context"
)

/*
 * Definition
 */

type ITrainerService interface {
	List(ctx context.Context) ([]dtos.TrainerDTO, error)
	Get(context.Context, int) (dtos.TrainerDTO, error)
	Create(ctx context.Context, trainer dtos.TrainerDTO) (dtos.TrainerDTO, error)
	Update(ctx context.Context, trainer dtos.TrainerDTO, ID int) (dtos.TrainerDTO, error)
	Delete(ctx context.Context, ID int) error
}

type TrainerService struct {
	da dataaccess.ITrainerDA
}

/*
 * Implement
 */

func (s *TrainerService) List(ctx context.Context) ([]dtos.TrainerDTO, error) {
	trainers, err := s.da.List(ctx)

	if err != nil {
		return []dtos.TrainerDTO{}, err
	}

	trainerDTOs := make([]dtos.TrainerDTO, len(trainers))

	for key, trainer := range trainers {
		trainerDTOs[key] = modelToDto(trainer)
	}

	return trainerDTOs, nil
}

func (s *TrainerService) Get(ctx context.Context, id int) (dtos.TrainerDTO, error) {
	trainer, err := s.da.Get(ctx, id)

	if err != nil {
		return dtos.TrainerDTO{}, err
	}

	return modelToDto(trainer), nil
}

func (s *TrainerService) Create(ctx context.Context, trainer dtos.TrainerDTO) (dtos.TrainerDTO, error) {
	createdTrainer, err := s.da.Create(ctx, dtoToModel(trainer))

	if err != nil {
		return dtos.TrainerDTO{}, err
	}

	return modelToDto(createdTrainer), nil
} 

func (s *TrainerService) Update(ctx context.Context, trainer dtos.TrainerDTO, ID int) (dtos.TrainerDTO, error) {
	updatedTrainer, err := s.da.Update(ctx, dtoToModel(trainer), ID)

	if err != nil {
		return dtos.TrainerDTO{}, err
	}

	return modelToDto(updatedTrainer), nil
} 

func (s *TrainerService) Delete(ctx context.Context, ID int) error {
	err := s.da.Delete(ctx, ID)

	if err != nil {
		return err
	}

	return nil
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

func dtoToModel(d dtos.TrainerDTO) models.TrainerModel {
	m := models.TrainerModel {
		ID: d.ID,
		FirstName: d.FirstName,
		LastName: d.LastName,
		DateOfBirth: d.DateOfBirth,
		Level: d.Level,
		YearOfExp: d.YearOfExp,
	}

	return m;
}

func CreateTrainerService(da dataaccess.ITrainerDA) *TrainerService {
	return &TrainerService{da}
}