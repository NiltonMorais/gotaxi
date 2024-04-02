package usecase

import (
	"context"

	"github.com/NiltonMorais/gotaxi/internal/domain/repository"
)

type StartRideUseCase struct {
	rideRepository repository.RideRepository
}

func NewStartRideUseCase(rideRepository repository.RideRepository) *StartRideUseCase {
	return &StartRideUseCase{
		rideRepository: rideRepository,
	}
}

func (a *StartRideUseCase) Execute(ctx context.Context, driverID, rideID string) error {
	ride, err := a.rideRepository.GetByID(ctx, rideID)
	if err != nil {
		return err
	}
	err = ride.Start()
	if err != nil {
		return err
	}

	err = a.rideRepository.Update(ctx, ride)
	if err != nil {
		return err
	}
	return nil
}
