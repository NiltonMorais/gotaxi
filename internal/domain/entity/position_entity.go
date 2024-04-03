package entity

import (
	"time"

	"github.com/NiltonMorais/gotaxi/internal/domain/entity/vo"
	"github.com/google/uuid"
)

type PositionEntity struct {
	id       string
	rideID   string
	location *vo.LocationVo
	date     time.Time
}

func NewPositionEntity(rideID string, lat, long float64) (*PositionEntity, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	return RestorePositionEntity(uuid.String(), rideID, lat, long, time.Now())
}

func RestorePositionEntity(id, rideID string, lat, long float64, date time.Time) (*PositionEntity, error) {
	location, err := vo.NewLocation(lat, long)
	if err != nil {
		return nil, err
	}

	return &PositionEntity{
		id:       id,
		rideID:   rideID,
		location: location,
		date:     date,
	}, nil
}

func (p *PositionEntity) GetID() string {
	return p.id
}

func (p *PositionEntity) GetRideID() string {
	return p.rideID
}

func (p *PositionEntity) GetLat() float64 {
	return p.location.GetLat()
}

func (p *PositionEntity) GetLong() float64 {
	return p.location.GetLong()
}

func (p *PositionEntity) GetDate() time.Time {
	return p.date
}
