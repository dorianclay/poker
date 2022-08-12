package cardgame

import (
	"math/rand"
	"time"
)

type CardSuit int

const (
	Clubs CardSuit = iota
	Diamonds
	Hearts
	Spades
)

type Card struct {
	Value int
	Suit  CardSuit
}

func (c Card) ValueToFace() string {
	switch c.Value {
	case 1:
		return "Ace"
	case 2:
		return "2"
	case 3:
		return "3"
	case 4:
		return "4"
	case 5:
		return "5"
	case 6:
		return "6"
	case 7:
		return "7"
	case 8:
		return "8"
	case 9:
		return "9"
	case 10:
		return "10"
	case 11:
		return "Jack"
	case 12:
		return "Queen"
	case 13:
		return "King"
	default:
		return ""
	}
}

func (c Card) SuitToString() string {
	switch c.Suit {
	case Clubs:
		return "Clubs"
	case Diamonds:
		return "Diamonds"
	case Hearts:
		return "Hearts"
	case Spades:
		return "Spades"
	default:
		return ""
	}
}

type Deck struct {
	Cards [52]Card
}

func (d Deck) Init() Deck {
	var count = 0
	var i = 1

	for ; i < 14; i++ {
		d.Cards[count] = Card{Value: i, Suit: Clubs}
		count += 1
	}

	i = 1
	for ; i < 14; i++ {
		d.Cards[count] = Card{Value: i, Suit: Diamonds}
		count += 1
	}

	i = 1
	for ; i < 14; i++ {
		d.Cards[count] = Card{Value: i, Suit: Hearts}
		count += 1
	}

	i = 1
	for ; i < 14; i++ {
		d.Cards[count] = Card{Value: i, Suit: Spades}
		count += 1
	}

	d.Shuffle()

	return d
}

func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(52, func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
}
