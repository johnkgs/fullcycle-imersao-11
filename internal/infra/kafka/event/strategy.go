package event

import (
	"context"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	uow "github.com/johnkgs/imersao11-consolidation/pkg"
)

type ProcessEventStrategy interface {
	Process(ctx context.Context, mdg *kafka.Message, uou uow.UowInterface) error
}
