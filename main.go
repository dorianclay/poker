package main

func main() {
	var card = Card{value: 1, suit: Spades}
	// println(card)
	println(card.SuitToString())
	println(card.ValueToFace())
}
