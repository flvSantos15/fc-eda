package entity

import (
	"time"

	"github.com/google/uuid"
)

type Balance struct {
	ID        string
	Account   *Account
	AccountID string
	Balance   float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewBalance(account *Account) *Balance {
	if account == nil {
		return nil
	}
	balance := &Balance{
		ID:        uuid.New().String(),
		Account:   account,
		Balance:   0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return balance
}
