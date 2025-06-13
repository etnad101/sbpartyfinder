package party

import (
	"errors"

	"github.com/etnad101/sbpartyfinder/internal/player"
)

type DungeonParty struct {
	BaseParty
	NoDupe      bool
	Description string
}

func NewDungeonParty(noDupe bool, description string) *DungeonParty {
	return &DungeonParty{
		BaseParty{
			Capacity: 5,
			Count:    0,
			Players:  make([]*player.Player, 5),
		},
		noDupe,
		description,
	}
}

func (p *DungeonParty) AddPlayer(newPlayer *player.Player) error {
	if p.NoDupe {
		for i := range p.Count {
			player := p.Players[i]
			if newPlayer.SelectedClass == player.SelectedClass {
				return errors.New("duplicated class, change class or turn off noDupe")
			}
		}
	}
	return p.Add(newPlayer)
}
