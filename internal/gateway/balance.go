package gateway

import "github.com.br/devfullcycle/fc-ms-wallet/internal/entity"

type BalanceGateway interface {
	Create(balance *entity.Balance) error
}
