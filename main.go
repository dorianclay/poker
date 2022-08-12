package main

import (
	"fmt"
	"poker/cardgame"
)

func main() {
	var deck = cardgame.Deck{}
	deck.Init()

	var player = cardgame.Player{}.Init("Bob Joe")
	for card := range player.Cards {
		player.Cards[card] = deck.Draw()
	}

	fmt.Println(player.Name, player.Id, player.Cards)
}
