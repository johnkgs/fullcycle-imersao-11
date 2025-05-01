package entity

import "github.com/johnkgs/imersao11-consolidation/pkg/utils"

type Team struct {
	ID      string
	Name    string
	Players []*Player
}

func NewTeam(id string, name string) *Team {
	return &Team{
		ID:   id,
		Name: name,
	}
}

func (t *Team) AddPlayer(player *Player) {
	t.Players = append(t.Players, player)
}

func (t *Team) RemovePlayer(player *Player) {
	for i, p := range t.Players {
		if p.ID == player.ID {
			t.Players = utils.Slice(t.Players, i, i+1)
			return
		}
	}
}
