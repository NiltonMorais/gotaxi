package repository

import (
	"context"

	"github.com/NiltonMorais/gotaxi/internal/domain/entity"
)

type PositionRepository interface {
	Save(ctx context.Context, position *entity.PositionEntity) error
	ListByRideID(ctx context.Context, rideID string) ([]*entity.PositionEntity, error)
}
