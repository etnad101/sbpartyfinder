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

type DungeonParty struct {
	BaseParty
}

type DianaParty struct {
	BaseParty
}

func NewDungeonParty() *DungeonParty {
	return &DungeonParty{
		BaseParty: BaseParty{
			Capacity: 4,
			Count:    0,
			Players:  make([]*player.Player, 4),
		},
	}
}

func NewDianaParty(capacity int) *DianaParty {
	return &DianaParty{
		BaseParty: BaseParty{
			Capacity: capacity,
			Count:    0,
			Players:  make([]*player.Player, capacity),
		},
	}
}

func (bp *BaseParty) CanAdd() bool {
	return bp.Count < bp.Capacity
}

func (bp *BaseParty) Add(name string) error {
	len := len(bp.Players)
	if len >= bp.Capacity {
		return errors.New("Party Is Full")
	}
	player := player.NewPlayer(name, 0, 0, 0)
	bp.Players[bp.Count] = player
	bp.Count += 1
	return nil
}

func (p *DungeonParty) AddPlayer(name string) error {
	return p.Add(name)
}

func (p *DianaParty) AddPlayer(name string) error {
	return p.Add(name)
}
