package usecase

import (
	"context"

	"github.com/johnkgs/imersao11-consolidation/internal/domain/entity"
	"github.com/johnkgs/imersao11-consolidation/internal/domain/repository"
	uow "github.com/johnkgs/imersao11-consolidation/pkg"
)

type AddPlayerInput struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	InitialPrice float64 `json:"initial_price"`
}

type AddPlayerUseCase struct {
	Uow uow.UowInterface
}

func NewAddPlayerUseCase(uow uow.UowInterface) *AddPlayerUseCase {
	return &AddPlayerUseCase{
		Uow: uow,
	}
}

func (a *AddPlayerUseCase) Execute(ctx context.Context, input AddPlayerInput) error {
	playerRepository := a.getPlayerRepository(ctx)
	player := entity.NewPlayer(input.ID, input.Name, input.InitialPrice)
	err := playerRepository.Create(ctx, player)

	if err != nil {
		return err
	}

	return a.Uow.CommitOrRollback()
}

func (a *AddPlayerUseCase) getPlayerRepository(ctx context.Context) repository.PlayerRepositoryInterface {
	playerRepository, err := a.Uow.GetRepository(ctx, "PlayerRepository")

	if err != nil {
		panic(err)
	}

	return playerRepository.(repository.PlayerRepositoryInterface)
}
