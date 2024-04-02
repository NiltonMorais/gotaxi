package usecase

import (
	"context"
	"errors"

	"github.com/NiltonMorais/gotaxi/internal/domain/entity"
	"github.com/NiltonMorais/gotaxi/internal/domain/repository"
)

type RequestRideUseCase struct {
	accountRepository repository.AccountRepository
	rideRepository    repository.RideRepository
}

func NewRequestRideUseCase(accountRepository repository.AccountRepository, rideRepository repository.RideRepository) *RequestRideUseCase {
	return &RequestRideUseCase{
		accountRepository: accountRepository,
		rideRepository:    rideRepository,
	}
}

func (r *RequestRideUseCase) Execute(ctx context.Context, passengerId string, fromLat, fromLng, toLat, toLng float64) (string, error) {
	account, err := r.accountRepository.GetById(ctx, passengerId)
	if err != nil {
		return "", err
	}
	if !account.IsPassenger() {
		return "", errors.New("account is not a passenger")
	}
	activeRide, _ := r.rideRepository.GetActivieRidesByPassengerID(ctx, passengerId)
	if len(activeRide) > 0 {
		return "", errors.New("passenger has an active ride")
	}
	ride, err := entity.NewRideEntity(passengerId, fromLat, fromLng, toLat, toLng)
	if err != nil {
		return "", err
	}
	err = r.rideRepository.Save(ctx, ride)
	if err != nil {
		return "", err
	}
	return ride.GetID(), nil
}
