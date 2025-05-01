package event

import (
	"context"
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/johnkgs/imersao11-consolidation/internal/usecase"
	uow "github.com/johnkgs/imersao11-consolidation/pkg"
)

type ProcessMatchUpdateResult struct{}

func (p ProcessMatchUpdateResult) Process(ctx context.Context, msg *kafka.Message, uow uow.UowInterface) error {
	var input usecase.MatchUpdateResultInput
	err := json.Unmarshal(msg.Value, &input)

	if err != nil {
		return err
	}

	updateMatchResultUsecase := usecase.NewMatchUpdateResultUseCase(uow)
	err = updateMatchResultUsecase.Execute(ctx, input)

	if err != nil {
		return err
	}

	return nil
}
