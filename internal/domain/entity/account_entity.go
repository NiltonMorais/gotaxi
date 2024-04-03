package entity

import (
	"github.com/NiltonMorais/gotaxi/internal/domain/vo"
	"github.com/google/uuid"
)

type AccountEntity struct {
	id          string
	name        *vo.NameVo
	email       *vo.EmailVo
	document    *vo.DocumentVo
	carPlate    *vo.PlateVo
	isPassenger bool
	isDriver    bool
}

func NewAccountEntity(name, email, document, carPlate string, isPassenger, isDriver bool) (*AccountEntity, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	return RestoreAccountEntity(uuid.String(), name, email, document, carPlate, isPassenger, isDriver)
}

func RestoreAccountEntity(id, name, email, document, carPlate string, isPassenger, isDriver bool) (*AccountEntity, error) {
	nameVo, err := vo.NewName(name)
	if err != nil {
		return nil, err
	}

	emailVo, err := vo.NewEmail(email)
	if err != nil {
		return nil, err
	}

	documentVo, err := vo.NewDocument(document)
	if err != nil {
		return nil, err
	}

	var carPlateVo *vo.PlateVo
	if isDriver {
		carPlateVo, err = vo.NewPlate(carPlate)
		if err != nil {
			return nil, err
		}
	}

	return &AccountEntity{
		id:          id,
		name:        nameVo,
		email:       emailVo,
		document:    documentVo,
		carPlate:    carPlateVo,
		isPassenger: isPassenger,
		isDriver:    isDriver,
	}, nil
}

func (a *AccountEntity) GetID() string {
	return a.id
}

func (a *AccountEntity) GetName() string {
	return a.name.Value()
}

func (a *AccountEntity) GetEmail() string {
	return a.email.Value()
}

func (a *AccountEntity) GetDocument() string {
	return a.document.Value()
}

func (a *AccountEntity) GetCarPlate() string {
	return a.carPlate.Value()
}

func (a *AccountEntity) IsPassenger() bool {
	return a.isPassenger
}

func (a *AccountEntity) IsDriver() bool {
	return a.isDriver
}
