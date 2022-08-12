package cardgame

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
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

func popCard(cards []Card) (Card, []Card) {
	popped := cards[len(cards)-1]
	cards = cards[:len(cards)-1]
	return popped, cards
}

func pushCard(cards []Card, card Card) []Card {
	cards = append(cards, card)
	return cards
}

type Deck struct {
	Cards []Card
	Drawn []Card
}

func (d *Deck) Init() {
	for i := 1; i < 14; i++ {
		d.Cards = append(d.Cards, Card{Value: i, Suit: Spades})
	}

	for i := 1; i < 14; i++ {
		d.Cards = append(d.Cards, Card{Value: i, Suit: Clubs})
	}

	for i := 1; i < 14; i++ {
		d.Cards = append(d.Cards, Card{Value: i, Suit: Hearts})
	}

	for i := 1; i < 14; i++ {
		d.Cards = append(d.Cards, Card{Value: i, Suit: Diamonds})
	}

	d.Shuffle()
}

func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
}

func (d *Deck) Reshuffle() {
	for len(d.Drawn) > 0 {
		var card Card
		card, d.Drawn = popCard(d.Drawn)
		d.Cards = pushCard(d.Cards, card)
	}

	d.Shuffle()
}

func (d *Deck) Draw() Card {
	var card Card
	card, d.Cards = popCard(d.Cards)
	d.Drawn = pushCard(d.Drawn, card)
	return card
}

type Player struct {
	Name  string
	Id    uuid.UUID
	Cards [2]Card
}

func (p Player) Init(name string) Player {
	p.Name = name
	p.Id = uuid.New()
	return p
}
