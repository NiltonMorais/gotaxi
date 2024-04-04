package repository

import (
	"context"
	"errors"
	"sync"

	"github.com/NiltonMorais/gotaxi/internal/domain/entity"
)

// AccountMemoryRepository é uma implementação de AccountRepository que armazena os dados em memória.
type AccountMemoryRepository struct {
	mu       sync.RWMutex
	accounts map[string]*entity.AccountEntity
}

// NewAccountMemoryRepository cria uma nova instância de AccountMemoryRepository.
func NewAccountMemoryRepository() *AccountMemoryRepository {
	return &AccountMemoryRepository{
		accounts: make(map[string]*entity.AccountEntity),
	}
}

// Save salva uma conta na memória.
func (repo *AccountMemoryRepository) Save(ctx context.Context, account *entity.AccountEntity) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.accounts[account.GetID()] = account
	return nil
}

// GetById recupera uma conta pelo ID da memória.
func (repo *AccountMemoryRepository) GetById(ctx context.Context, id string) (*entity.AccountEntity, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()
	account, ok := repo.accounts[id]
	if !ok {
		return nil, errors.New("conta não encontrada")
	}
	return account, nil
}

// GetByEmail recupera uma conta pelo e-mail da memória.
func (repo *AccountMemoryRepository) GetByEmail(ctx context.Context, email string) (*entity.AccountEntity, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()
	for _, account := range repo.accounts {
		if account.GetEmail() == email {
			return account, nil
		}
	}
	return nil, errors.New("conta não encontrada")
}
