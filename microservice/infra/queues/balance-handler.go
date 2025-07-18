package queues

import (
	"encoding/json"
	"fmt"
	"microservice/application"
	"microservice/domain/dto"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type BalanceHandler struct {
	useCases *application.BalanceUseCases
}

func NewBalanceHandler(useCases *application.BalanceUseCases) *BalanceHandler {
	return &BalanceHandler{
		useCases: useCases,
	}
}

func (h *BalanceHandler) Handle(msg *ckafka.Message) error {
	var balanceEvent dto.BalanceEvent
	if err := json.Unmarshal(msg.Value, &balanceEvent); err != nil {
		return fmt.Errorf("error unmarshaling message: %v", err)
	}

	createBalanceInput := dto.CreateBalanceInput{
		AccountID: balanceEvent.Payload.AccountIDTo,
		Balance:   float64(balanceEvent.Payload.BalanceAccountIDTo),
	}

	_, err := h.useCases.CreateBalance.Execute(&createBalanceInput)
	if err != nil {
		return fmt.Errorf("error creating balance: %v", err)
	}
	return nil
}
