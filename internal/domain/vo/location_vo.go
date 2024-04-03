package vo

import "errors"

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

func (l *LocationVo) Latitude() float64 {
	return l.lat
}

func (l *LocationVo) Longitude() float64 {
	return l.long
}
