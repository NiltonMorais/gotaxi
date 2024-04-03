package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAccountEntity(t *testing.T) {
	t.Run("should return an account entity with a valid id", func(t *testing.T) {
		accountEntity, err := NewAccountEntity("Nilton Morais", "nilton@gmail.com", "123.456.789-10", "ABC1234", true, true)
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
		_, err := NewAccountEntity("Nilton Morais", "nilton", "123.456.789-10", "ABC1234", true, true)
		if err == nil {
			t.Error("NewAccountEntity() error = nil; want an error")
		}
		if err.Error() != "invalid email" {
			t.Errorf("NewAccountEntity() error = %v; want invalid email", err)
		}
	})
}

func TestNewAccountEntity_InvalidName(t *testing.T) {
	t.Run("should return an error when name is invalid", func(t *testing.T) {
		t.Run("should return an error when name is invalid", func(t *testing.T) {
			_, err := NewAccountEntity("Nilton", "nilton@gmail.com", "123.456.789-10", "ABC1234", true, true)
			if err == nil {
				t.Error("NewAccountEntity() error = nil; want an error")
			}
			if err.Error() != "invalid name" {
				t.Errorf("NewAccountEntity() error = %v; want invalid name", err)
			}
		})
	})
}

func TestNewAccountEntity_InvalidDocument(t *testing.T) {
	t.Run("should return an error when document is invalid", func(t *testing.T) {
		t.Run("should return an error when document is invalid", func(t *testing.T) {
			_, err := NewAccountEntity("Nilton Morais", "nilton@gmail.com", "123-invalid", "ABC1234", true, true)
			if err == nil {
				t.Error("NewAccountEntity() error = nil; want an error")
			}
			if err.Error() != "invalid document" {
				t.Errorf("NewAccountEntity() error = %v; want invalid document", err)
			}
		})
	})
}

func TestNewAccountEntity_InvalidCarPlate(t *testing.T) {
	t.Run("should return an error when car plate is invalid", func(t *testing.T) {
		_, err := NewAccountEntity("Nilton Morais", "nilton@gmail.com", "123.456.789-10", "A1", true, true)
		if err == nil {
			t.Error("NewAccountEntity() error = nil; want an error")
		}
		if err.Error() != "invalid plate" {
			t.Errorf("NewAccountEntity() error = %v; want invalid car plate", err)
		}
	})
}

func TestRestoreAccountEntity(t *testing.T) {
	t.Run("should return an account entity with a valid id", func(t *testing.T) {
		accountEntity, err := RestoreAccountEntity("123", "Nilton Morais", "nilton@gmail.com", "123.456.789-10", "ABC1234", true, true)
		assert.Nil(t, err)
		assert.Equal(t, accountEntity.GetID(), "123")
		assert.Equal(t, accountEntity.GetName(), "Nilton Morais")
		assert.Equal(t, accountEntity.GetEmail(), "nilton@gmail.com")
		assert.Equal(t, accountEntity.GetDocument(), "123.456.789-10")
		assert.Equal(t, accountEntity.GetCarPlate(), "ABC1234")
		assert.Equal(t, accountEntity.IsPassenger(), true)
		assert.Equal(t, accountEntity.IsDriver(), true)
	})
}
