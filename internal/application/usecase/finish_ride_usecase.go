package usecase

import (
	"context"

	"github.com/NiltonMorais/gotaxi/internal/domain/repository"
)

type FinishRideUseCase struct {
	rideRepository repository.RideRepository
}

func NewFinishRideUseCase(rideRepository repository.RideRepository) *FinishRideUseCase {
	return &FinishRideUseCase{
		rideRepository: rideRepository,
	}
}

func (f *FinishRideUseCase) Execute(ctx context.Context, rideID string) error {
	ride, err := f.rideRepository.GetByID(ctx, rideID)
	if err != nil {
		return err
	}
	err = ride.Finish()
	if err != nil {
		return err
	}
	err = f.rideRepository.Update(ctx, ride)
	if err != nil {
		return err
	}
	return nil
}
