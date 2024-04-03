package vo

import (
	"errors"
	"regexp"
)

type CarPlateVo struct {
	value string
}

func NewCarPlate(plate string) (*CarPlateVo, error) {
	if isInvalidCarPlate(plate) {
		return nil, errors.New("invalid car plate")
	}

	return &CarPlateVo{
		value: plate,
	}, nil
}

func (p *CarPlateVo) Value() string {
	return p.value
}

func isInvalidCarPlate(plate string) bool {
	return !regexp.MustCompile(`[A-Z]{3}[0-9]{4}`).MatchString(plate)
}
