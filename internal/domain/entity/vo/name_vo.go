package vo

import (
	"errors"
	"regexp"
)

type NameVo struct {
	value string
}

func NewName(value string) (*NameVo, error) {
	if value == "" {
		return nil, errors.New("name cannot be empty")
	}

	if isInvalidName(value) {
		return nil, errors.New("invalid name")
	}

	return &NameVo{
		value: value,
	}, nil
}

func (n *NameVo) Value() string {
	return n.value
}

func isInvalidName(name string) bool {
	return !regexp.MustCompile(`[a-zA-Z] [a-zA-Z]+`).MatchString(name)
}
