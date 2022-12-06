package event

import (
	"context"
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/johnkgs/imersao11-consolidation/internal/usecase"
	uow "github.com/johnkgs/imersao11-consolidation/pkg"
)

type ProcessChooseTeam struct{}

func (p ProcessChooseTeam) Process(ctx context.Context, msg *kafka.Message, uow uow.UowInterface) error {
	var input usecase.MyTeamChoosePlayersInput
	err := json.Unmarshal(msg.Value, &input)

	if err != nil {
		return err
	}

	addNewMyTeamUseCase := usecase.NewMyTeamChoosePlayersUseCase(uow)
	err = addNewMyTeamUseCase.Execute(ctx, input)

	if err != nil {
		return err
	}

	return nil
}
