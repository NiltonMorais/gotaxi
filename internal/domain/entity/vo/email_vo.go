package vo

import (
	"errors"
	"regexp"
)

type EmailVo struct {
	value string
}

func NewEmail(email string) (*EmailVo, error) {
	if isInvalidEmail(email) {
		return nil, errors.New("invalid email")
	}

	return &EmailVo{
		value: email,
	}, nil
}

func (e *EmailVo) Value() string {
	return e.value
}

func isInvalidEmail(email string) bool {
	return !regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`).MatchString(email)
}
