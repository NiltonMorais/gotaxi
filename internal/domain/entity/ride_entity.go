package entity

import (
	"errors"
	"time"

	"github.com/NiltonMorais/gotaxi/internal/domain/vo"
	"github.com/google/uuid"
)

const (
	RideStatusRequested = "requested"
	RideStatusAccepted  = "accepted"
	RideStatusStarted   = "started"
)

type RideEntity struct {
	id           string
	passengerID  string
	driverID     string
	fromLocation *vo.LocationVo
	toLocation   *vo.LocationVo
	status       string
	date         time.Time
}

func NewRideEntity(passengerID string, fromLat, fromLng, toLat, toLng float64) (*RideEntity, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return RestoreRideEntity(uuid.String(), passengerID, "", fromLat, fromLng, toLat, toLng, RideStatusRequested, time.Now())
}

func RestoreRideEntity(id, passengerID, driverID string, fromLat, fromLng, toLat, toLng float64, status string, date time.Time) (*RideEntity, error) {
	fromLocation, err := vo.NewLocation(fromLat, fromLng)
	if err != nil {
		return nil, err
	}

	toLocation, err := vo.NewLocation(toLat, toLng)
	if err != nil {
		return nil, err
	}

	return &RideEntity{
		id:           id,
		passengerID:  passengerID,
		driverID:     driverID,
		fromLocation: fromLocation,
		toLocation:   toLocation,
		status:       status,
		date:         date,
	}, nil
}

func (r *RideEntity) Accept(driverID string) error {
	if r.status != RideStatusRequested {
		return errors.New("ride is not requested")
	}
	r.driverID = driverID
	r.status = RideStatusAccepted
	return nil
}

func (r *RideEntity) Start() error {
	if r.status != RideStatusAccepted {
		return errors.New("ride is not accepted")
	}
	r.status = RideStatusStarted
	return nil
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
	return r.fromLocation.Latitude()
}

func (r *RideEntity) GetFromLng() float64 {
	return r.fromLocation.Longitude()
}

func (r *RideEntity) GetToLat() float64 {
	return r.toLocation.Latitude()
}

func (r *RideEntity) GetToLng() float64 {
	return r.toLocation.Longitude()
}

func (r *RideEntity) GetStatus() string {
	return r.status
}

func (r *RideEntity) GetDate() time.Time {
	return r.date
}
