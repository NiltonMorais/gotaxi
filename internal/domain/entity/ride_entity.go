package entity

import (
	"errors"
	"time"

	"github.com/NiltonMorais/gotaxi/internal/domain/entity/vo"
	"github.com/NiltonMorais/gotaxi/internal/domain/service"
	"github.com/google/uuid"
)

const (
	RideStatusRequested = "requested"
	RideStatusAccepted  = "accepted"
	RideStatusStarted   = "started"
	RideStatusCompleted = "completed"
)

type RideEntity struct {
	id           string
	passengerID  string
	driverID     string
	fromLocation *vo.LocationVo
	toLocation   *vo.LocationVo
	status       string
	date         time.Time
	lastPosition *vo.LocationVo
	distance     float64
	price        float64
}

func NewRideEntity(passengerID string, fromLat, fromLng, toLat, toLng float64) (*RideEntity, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return RestoreRideEntity(uuid.String(), passengerID, "", fromLat, fromLng, toLat, toLng, RideStatusRequested, time.Now(), fromLat, fromLng, 0, 0)
}

func RestoreRideEntity(id, passengerID, driverID string, fromLat, fromLng, toLat, toLng float64, status string, date time.Time, lastLat, lastLng, distance, price float64) (*RideEntity, error) {
	fromLocation, err := vo.NewLocation(fromLat, fromLng)
	if err != nil {
		return nil, err
	}

	toLocation, err := vo.NewLocation(toLat, toLng)
	if err != nil {
		return nil, err
	}

	lastPosition, err := vo.NewLocation(lastLat, lastLng)
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
		lastPosition: lastPosition,
		distance:     distance,
		price:        price,
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

func (r *RideEntity) UpdatePosition(lat, long float64) error {
	if r.status != RideStatusStarted {
		return errors.New("ride is not started")
	}

	if lat == r.lastPosition.GetLat() && long == r.lastPosition.GetLong() {
		return errors.New("this old last position is same the new last position")
	}

	newLastPosition, err := vo.NewLocation(lat, long)
	if err != nil {
		return err
	}
	r.distance += r.lastPosition.DistanceTo(newLastPosition)
	r.lastPosition = newLastPosition
	r.calculatePrice()
	return nil
}

func (r *RideEntity) Finish() error {
	if r.status != RideStatusStarted {
		return errors.New("ride is not started")
	}
	r.status = RideStatusCompleted
	r.calculatePrice()
	return nil
}

func (r *RideEntity) calculatePrice() {
	r.price = service.NewPriceCalculatorServiceFactory().NewPriceCalculatorService(r.date).Calculate(r.distance)
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
	return r.fromLocation.GetLat()
}

func (r *RideEntity) GetFromLng() float64 {
	return r.fromLocation.GetLong()
}

func (r *RideEntity) GetToLat() float64 {
	return r.toLocation.GetLat()
}

func (r *RideEntity) GetToLng() float64 {
	return r.toLocation.GetLong()
}

func (r *RideEntity) GetStatus() string {
	return r.status
}

func (r *RideEntity) GetDate() time.Time {
	return r.date
}

func (r *RideEntity) GetLastLat() float64 {
	return r.lastPosition.GetLat()
}

func (r *RideEntity) GetLastLng() float64 {
	return r.lastPosition.GetLong()
}

func (r *RideEntity) GetDistance() float64 {
	return r.distance
}

func (r *RideEntity) GetPrice() float64 {
	return r.price
}
