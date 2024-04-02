package repository

import (
	"context"

	"github.com/NiltonMorais/gotaxi/internal/domain/entity"
)

type RideRepository interface {
	GetActivieRidesByPassengerID(ctx context.Context, passengerID string) ([]*entity.RideEntity, error)
	GetByID(ctx context.Context, id string) (*entity.RideEntity, error)
	Save(ctx context.Context, ride *entity.RideEntity) error
	Update(ctx context.Context, ride *entity.RideEntity) error
}
