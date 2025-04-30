package handler

import (
	"fmt"
	"sync"

	"github.com.br/devfullcycle/fc-ms-wallet/pkg/events"
	"github.com.br/devfullcycle/fc-ms-wallet/pkg/kafka"
)

type CreateBalanceKafkaHandler struct {
	Kafka *kafka.Producer
}

func NewCreateBalanceKafkaHandler(kafka *kafka.Producer) *CreateBalanceKafkaHandler {
	return &CreateBalanceKafkaHandler{
		Kafka: kafka,
	}
}

func (h *CreateBalanceKafkaHandler) Handle(message events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	h.Kafka.Publish(message, nil, "balances")
	fmt.Println("CreatedBalanceKafkaHandler called")
}
