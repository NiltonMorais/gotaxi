package repository

import (
	"context"

	"github.com/NiltonMorais/gotaxi/internal/domain/entity"
)

type AccountRepository interface {
	Save(ctx context.Context, account *entity.AccountEntity) error
	GetById(ctx context.Context, id string) (*entity.AccountEntity, error)
	GetByEmail(ctx context.Context, email string) (*entity.AccountEntity, error)
}
