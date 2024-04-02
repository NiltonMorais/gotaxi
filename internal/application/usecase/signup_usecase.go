package usecase

import (
	"context"
	"errors"

	"github.com/NiltonMorais/gotaxi/internal/domain/entity"
	"github.com/NiltonMorais/gotaxi/internal/domain/gateway"
	"github.com/NiltonMorais/gotaxi/internal/domain/repository"
)

type SignupUseCase struct {
	accountRepository repository.AccountRepository
	emailGateway      gateway.EmailGateway
}

func NewSignupUseCase(accountRepository repository.AccountRepository, emailGateway gateway.EmailGateway) *SignupUseCase {
	return &SignupUseCase{
		accountRepository: accountRepository,
		emailGateway:      emailGateway,
	}
}

func (s *SignupUseCase) Execute(ctx context.Context, name, email, document, carPlate string, isPassenger, isDriver bool) error {
	existingAccount, _ := s.accountRepository.GetByEmail(ctx, email)
	if existingAccount != nil {
		return errors.New("account already exists")
	}

	account, err := entity.NewAccountEntity(name, email, document, carPlate, isPassenger, isDriver)
	if err != nil {
		return err
	}

	err = s.accountRepository.Save(ctx, account)
	if err != nil {
		return err
	}
	err = s.emailGateway.SendEmail(email, "Welcome to GoTaxi", "Use this link to confirm your account: {link}")
	if err != nil {
		return err
	}

	return nil
}
