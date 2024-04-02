package entity

import (
	"errors"
	"regexp"

	"github.com/google/uuid"
)

type AccountEntity struct {
	id          string
	name        string
	email       string
	document    string
	carPlate    string
	isPassenger bool
	isDriver    bool
}

func NewAccountEntity(name, email, document, carPlate string, isPassenger, isDriver bool) (*AccountEntity, error) {
	if isInvalidEmail(email) {
		return nil, errors.New("invalid email")
	}
	if isInvalidName(name) {
		return nil, errors.New("invalid name")
	}
	if isDriver && isInvalidCarPlate(carPlate) {
		return nil, errors.New("invalid car plate")
	}
	uuid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	return &AccountEntity{
		id:          uuid.String(),
		name:        name,
		email:       email,
		document:    document,
		carPlate:    carPlate,
		isPassenger: isPassenger,
		isDriver:    isDriver,
	}, nil
}

func RestoreAccountEntity(id, name, email, document, carPlate string, isPassenger, isDriver bool) *AccountEntity {
	return &AccountEntity{
		id:          id,
		name:        name,
		email:       email,
		document:    document,
		carPlate:    carPlate,
		isPassenger: isPassenger,
		isDriver:    isDriver,
	}
}

func (a *AccountEntity) GetID() string {
	return a.id
}

func (a *AccountEntity) GetName() string {
	return a.name
}

func (a *AccountEntity) GetEmail() string {
	return a.email
}

func (a *AccountEntity) GetDocument() string {
	return a.document
}

func (a *AccountEntity) GetCarPlate() string {
	return a.carPlate
}

func (a *AccountEntity) IsPassenger() bool {
	return a.isPassenger
}

func (a *AccountEntity) IsDriver() bool {
	return a.isDriver
}

func isInvalidEmail(email string) bool {
	return !regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`).MatchString(email)
}

func isInvalidName(name string) bool {
	return !regexp.MustCompile(`[a-zA-Z] [a-zA-Z]+`).MatchString(name)
}

func isInvalidCarPlate(carPlate string) bool {
	return !regexp.MustCompile(`[A-Z]{3}[0-9]{4}`).MatchString(carPlate)
}
