package party

import "github.com/etnad101/sbpartyfinder/internal/player"

type DianaParty struct {
	BaseParty
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

func (p *DianaParty) AddPlayer(player *player.Player) error {
	return p.Add(player)
}
