package vo

import (
	"errors"
	"regexp"
)

type PlateVo struct {
	value string
}

func NewPlate(plate string) (*PlateVo, error) {
	if isInvalidPlate(plate) {
		return nil, errors.New("invalid plate")
	}

	return &PlateVo{
		value: plate,
	}, nil
}

func (p *PlateVo) Value() string {
	return p.value
}

func isInvalidPlate(plate string) bool {
	return !regexp.MustCompile(`[A-Z]{3}[0-9]{4}`).MatchString(plate)
}
