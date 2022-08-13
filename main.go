package main

import (
	"fmt"
	"poker/cardgame"
)

func main() {
	var game = cardgame.Game{}.Init(25)

	var player1 = cardgame.Player{}.Init("Bob Joe")
	var player2 = cardgame.Player{}.Init("Booty Bish")

	game.Players = append(game.Players, player1, player2)

	fmt.Println("Game:: code:", game.Code, "id:", game.Id, "n_players:", len(game.Players))

	game.NewRound()

	game.DealPlayersHoleCards()
	game.Flop()

	println("Players:")
	for _, player := range game.Players {
		fmt.Println("\tPlayer ", player.Name, ":", player.Cards[0].ValueToFace(), player.Cards[0].SuitToString(), ",", player.Cards[1].ValueToFace(), player.Cards[1].SuitToString())
	}

	println("Flop:")
	for i, card := range game.FaceUp {
		fmt.Println("\t", i, ":", card.ValueToFace(), "of", card.SuitToString())
	}
	// for card := range game.GameDeck.Cards {
	// 	elt := game.GameDeck.Cards[card]
	// 	fmt.Println(card, ":", elt.ValueToFace(), "of", elt.SuitToString())
	// }
}
