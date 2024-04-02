package repository

import (
	"context"

	"github.com/NiltonMorais/gotaxi/internal/domain/entity"
)

type RideRepository interface {
	Save(ctx context.Context, ride *entity.RideEntity) error
	GetActivieRidesByPassengerID(ctx context.Context, passengerID string) ([]*entity.RideEntity, error)
}
