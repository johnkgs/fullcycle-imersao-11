package event

import (
	"context"
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/johnkgs/imersao11-consolidation/internal/usecase"
	uow "github.com/johnkgs/imersao11-consolidation/pkg"
)

type ProcessNewMatch struct{}

func (process ProcessNewMatch) Process(ctx context.Context, msg *kafka.Message, uow uow.UowInterface) error {
	var input usecase.MatchInput
	err := json.Unmarshal(msg.Value, &input)

	if err != nil {
		return err
	}

	newMatchUseCase := usecase.NewMatchUseCase(uow)
	err = newMatchUseCase.Execute(ctx, input)

	if err != nil {
		return err
	}

	return nil

}
