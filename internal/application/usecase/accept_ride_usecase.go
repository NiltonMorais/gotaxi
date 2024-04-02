package usecase

import (
	"context"
	"errors"

	"github.com/NiltonMorais/gotaxi/internal/domain/repository"
)

type AcceptRideUseCase struct {
	accountRepository repository.AccountRepository
	rideRepository    repository.RideRepository
}

func NewAcceptRideUseCase(accountRepository repository.AccountRepository, rideRepository repository.RideRepository) *AcceptRideUseCase {
	return &AcceptRideUseCase{
		accountRepository: accountRepository,
		rideRepository:    rideRepository,
	}
}

func (a *AcceptRideUseCase) Execute(ctx context.Context, driverID, rideID string) error {
	driverAccount, err := a.accountRepository.GetById(ctx, driverID)
	if err != nil {
		return err
	}
	if !driverAccount.IsDriver() {
		return errors.New("account is not a driver")
	}
	ride, err := a.rideRepository.GetByID(ctx, rideID)
	if err != nil {
		return err
	}
	err = ride.Accept(driverID)
	if err != nil {
		return err
	}

	err = a.rideRepository.Update(ctx, ride)
	if err != nil {
		return err
	}
	return nil
}
