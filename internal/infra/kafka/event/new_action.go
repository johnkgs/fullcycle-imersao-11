package event

import (
	"context"
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/johnkgs/imersao11-consolidation/internal/domain/entity"
	"github.com/johnkgs/imersao11-consolidation/internal/usecase"
	uow "github.com/johnkgs/imersao11-consolidation/pkg"
)

type ProcessNewAction struct{}

func (p ProcessNewAction) Process(ctx context.Context, msg *kafka.Message, uow uow.UowInterface) error {
	var input usecase.ActionAddInput
	err := json.Unmarshal(msg.Value, &input)
	if err != nil {
		return err
	}
	actionTable := entity.ActionTable{}
	actionTable.Init()
	addNewActionUseCase := usecase.NewActionAddUseCase(uow, &actionTable)
	err = addNewActionUseCase.Execute(ctx, input)
	if err != nil {
		return err
	}
	return nil
}
