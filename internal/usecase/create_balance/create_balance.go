package createbalance

import (
	"context"

	"github.com.br/devfullcycle/fc-ms-wallet/internal/gateway"
	"github.com.br/devfullcycle/fc-ms-wallet/pkg/events"
	"github.com.br/devfullcycle/fc-ms-wallet/pkg/uow"
)

type CreateBalanceInputDTO struct {
	AccountID string  `json:"account_id"`
	Amount    float64 `json:"amount"`
}

type CreateBalanceOutputDTO struct {
	ID      string  `json:"id"`
	Account string  `json:"account"`
	Amount  float64 `json:"amount"`
}

type CreateBalanceUseCase struct {
	EventDispatcher events.EventDispatcherInterface
	BalanceCreated  events.EventInterface
}

func NewCreateBalanceUseCase(
	eventDispatcher events.EventDispatcherInterface,
	balanceCreated events.EventInterface,
) *CreateBalanceUseCase {
	return &CreateBalanceUseCase{
		EventDispatcher: eventDispatcher,
		BalanceCreated:  balanceCreated,
	}
}

func (uc *CreateBalanceUseCase) Execute(input CreateBalanceInputDTO) (*CreateBalanceOutputDTO, error) {
	output := &CreateBalanceOutputDTO{
		ID:      input.AccountID,
		Account: input.AccountID,
		Amount:  input.Amount,
	}
	uc.EventDispatcher.Dispatch(uc.BalanceCreated)
	return output, nil
}

type GetBalanceUseCase struct {
	Uow uow.UowInterface
}

func (uc *GetBalanceUseCase) NewGetBalanceUseCase(context context.Context) gateway.BalanceGateway {
	repo, err := uc.Uow.GetRepository(context, "BalanceDB")
	if err != nil {
		panic(err)
	}
	return repo.(gateway.BalanceGateway)
}
