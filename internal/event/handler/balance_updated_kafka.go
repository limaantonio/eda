package handler

import (
	"fmt"
	"sync"

	"github.com.br/limaantonio/ms-wallet/pkg/events"
	"github.com.br/limaantonio/ms-wallet/pkg/kafka"
)

type UpdatedBalanceKafkaHandler struct {
	Kafka *kafka.Producer
}

func NewUpdatedBalanceKafkaHandler(kafka *kafka.Producer) *UpdatedBalanceKafkaHandler {
	return &UpdatedBalanceKafkaHandler{
		Kafka: kafka,
	}
}

func (h *UpdatedBalanceKafkaHandler) Handle(message events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	h.Kafka.Publish(message, nil, "balances")
	fmt.Println("UpdatedBalanceKafkaHandler called ")
}
