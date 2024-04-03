package usecase

import (
	"context"

	"github.com/NiltonMorais/gotaxi/internal/domain/entity"
	"github.com/NiltonMorais/gotaxi/internal/domain/repository"
)

type UpdatePositionUseCase struct {
	rideRepository     repository.RideRepository
	positionRepository repository.PositionRepository
}

func NewUpdatePositionUseCase(rideRepository repository.RideRepository, positionRepository repository.PositionRepository) *UpdatePositionUseCase {
	return &UpdatePositionUseCase{
		rideRepository:     rideRepository,
		positionRepository: positionRepository,
	}
}

func (u *UpdatePositionUseCase) Execute(ctx context.Context, rideID string, lat, long float64) error {
	ride, err := u.rideRepository.GetByID(ctx, rideID)
	if err != nil {
		return err
	}
	err = ride.UpdatePosition(lat, long)
	if err != nil {
		return err
	}
	err = u.rideRepository.Update(ctx, ride)
	if err != nil {
		return err
	}
	position, err := entity.NewPositionEntity(rideID, lat, long)
	if err != nil {
		return err
	}
	err = u.positionRepository.Save(ctx, position)
	if err != nil {
		return err
	}
	return nil
}
