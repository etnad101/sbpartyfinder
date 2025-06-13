package party

import (
	"errors"

	"github.com/etnad101/sbpartyfinder/internal/player"
)

type BaseParty struct {
	Capacity int
	Count    int
	Players  []*player.Player
}

type Party interface {
	AddPlayer(name string) error
	RemovePlayer(name string) error
}

func (bp *BaseParty) CanAdd() bool {
	return bp.Count < bp.Capacity
}

func (bp *BaseParty) Add(player *player.Player) error {
	if bp.Count >= bp.Capacity {
		return errors.New("Party Is Full")
	}
	bp.Players[bp.Count] = player
	bp.Count += 1
	return nil
}
