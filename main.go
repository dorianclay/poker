package main

import (
	"fmt"
	"poker/cardgame"
)

func main() {
	var card = cardgame.Card{Value: 1, Suit: cardgame.Spades}
	println(card.SuitToString())
	println(card.ValueToFace())

	var deck = cardgame.Deck{}.Init().Shuffle()
	for card := range deck.Cards {
		fmt.Println("count:", card, "-", deck.Cards[card].ValueToFace(), "of", deck.Cards[card].SuitToString())
	}
}
