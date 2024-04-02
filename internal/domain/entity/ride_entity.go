package entity

import (
	"time"

	"github.com/google/uuid"
)

type RideEntity struct {
	id          string
	passengerID string
	driverID    string
	fromLat     float64
	fromLng     float64
	toLat       float64
	toLng       float64
	status      string
	date        time.Time
}

func NewRideEntity(passengerID string, fromLat, fromLng, toLat, toLng float64) (*RideEntity, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return &RideEntity{
		id:          uuid.String(),
		passengerID: passengerID,
		driverID:    "",
		fromLat:     fromLat,
		fromLng:     fromLng,
		toLat:       toLat,
		toLng:       toLng,
		status:      "requested",
		date:        time.Now(),
	}, nil
}

func RestoreRideEntity(id, passengerID, driverID string, fromLat, fromLng, toLat, toLng float64, status string, date time.Time) *RideEntity {
	return &RideEntity{
		id:          id,
		passengerID: passengerID,
		driverID:    driverID,
		fromLat:     fromLat,
		fromLng:     fromLng,
		toLat:       toLat,
		toLng:       toLng,
		status:      status,
		date:        date,
	}
}

func (r *RideEntity) GetID() string {
	return r.id
}

func (r *RideEntity) GetPassengerID() string {
	return r.passengerID
}

func (r *RideEntity) GetDriverID() string {
	return r.driverID
}

func (r *RideEntity) GetFromLat() float64 {
	return r.fromLat
}

func (r *RideEntity) GetFromLng() float64 {
	return r.fromLng
}

func (r *RideEntity) GetToLat() float64 {
	return r.toLat
}

func (r *RideEntity) GetToLng() float64 {
	return r.toLng
}

func (r *RideEntity) GetStatus() string {
	return r.status
}

func (r *RideEntity) GetDate() time.Time {
	return r.date
}

func (r *RideEntity) SetDriverID(driverID string) {
	r.driverID = driverID
}
