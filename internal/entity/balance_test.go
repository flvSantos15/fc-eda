package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateBalance(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j")
	account := NewAccount(client)
	balance := NewBalance(account)
	assert.NotNil(t, balance)
	assert.Equal(t, account.ID, balance.Account.ID)
}

func TestCreateBalanceWithNilAccount(t *testing.T) {
	account := NewAccount(nil)
	assert.Nil(t, account)
}
