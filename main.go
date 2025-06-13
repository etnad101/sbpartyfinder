package main

import (
	"fmt"

	"github.com/etnad101/sbpartyfinder/internal/party"
	"github.com/etnad101/sbpartyfinder/internal/player"
)

func main() {
	p := party.NewDungeonParty(true, "Chill Dungeon Party")
	archer := player.NewPlayer("4ante", 230, 30, 30, player.Archer)
	bers := player.NewPlayer("4ante", 230, 30, 30, player.Berserk)
	healer := player.NewPlayer("4ante", 230, 30, 30, player.Healer)
	mage := player.NewPlayer("4ante", 230, 30, 30, player.Mage)
	tank := player.NewPlayer("4ante", 230, 30, 30, player.Tank)

	players := []*player.Player{archer, bers, healer, mage, tank}

	for i := range 5 {
		err := p.AddPlayer(players[i])
		if err != nil {
			fmt.Println(err)
		}
	}
}
