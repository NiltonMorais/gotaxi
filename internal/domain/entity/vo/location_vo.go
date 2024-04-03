package vo

import (
	"errors"
	"math"
)

type LocationVo struct {
	lat  float64
	long float64
}

func NewLocation(lat, long float64) (*LocationVo, error) {
	if lat < -90 || lat > 90 {
		return nil, errors.New("invalid latitude")
	}
	if long < -180 || long > 180 {
		return nil, errors.New("invalid longitude")
	}
	return &LocationVo{
		lat:  lat,
		long: long,
	}, nil
}

func (l *LocationVo) GetLat() float64 {
	return l.lat
}

func (l *LocationVo) GetLong() float64 {
	return l.long
}

func (from *LocationVo) DistanceTo(to *LocationVo) float64 {
	earthRadius := 6371
	degressToRadians := math.Pi / 180
	deltaLat := (to.GetLat() - from.GetLat()) * degressToRadians
	deltaLong := (to.GetLong() - from.GetLong()) * degressToRadians
	a := math.Sin(deltaLat/2)*math.Sin(deltaLat/2) +
		math.Cos(from.GetLat()*degressToRadians)*math.Cos(to.GetLat()*degressToRadians)*
			math.Sin(deltaLong/2)*math.Sin(deltaLong/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return float64(earthRadius) * c
}
