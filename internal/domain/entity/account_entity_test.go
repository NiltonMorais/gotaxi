package entity

import "testing"

func TestNewAccountEntity(t *testing.T) {
	t.Run("should return an account entity with a valid id", func(t *testing.T) {
		accountEntity, err := NewAccountEntity("Nilton", "nilton@gmail.com", "123456789", "ABC1234", true, true)
		if err != nil {
			t.Errorf("NewAccountEntity() error = %v; want nil", err)
			return
		}

		if accountEntity.GetID() == "" {
			t.Errorf("NewAccountEntity() got = %v; want a valid id", accountEntity.GetID())
		}
	})
}

func TestNewAccountEntity_InvalidEmail(t *testing.T) {
	t.Run("should return an error when email is invalid", func(t *testing.T) {
		_, err := NewAccountEntity("Nilton", "nilton", "123456789", "ABC1234", true, true)
		if err == nil {
			t.Error("NewAccountEntity() error = nil; want an error")
		}
	})
}
