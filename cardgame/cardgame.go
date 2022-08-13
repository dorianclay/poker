package cardgame

import (
	"math/rand"
	"poker/codegenerator"
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

/*
 *    CARD
 */

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

/*
 *   DECK
 */
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

type RoleButton int

const (
	Dealer RoleButton = iota
	SmallBlind
	BigBlind
	NoRole
)

/*
 *   PLAYER
 */
type Player struct {
	Name   string
	Id     uuid.UUID
	Cards  []Card
	Chips  int
	Folded bool
	Role   RoleButton
	Bet    int
}

func (p Player) Init(name string) Player {
	p.Name = name
	p.Id = uuid.New()
	p.Cards = make([]Card, 0, 2)
	p.Chips = 0
	p.Folded = false
	p.Role = NoRole
	p.Bet = 0
	return p
}

func (p *Player) PlaceBet(amount int) bool {
	if amount < p.Chips {
		p.Bet += amount
		return true
	}
	return false
}

/*
 *   GAME
 */
type Game struct {
	Players  []Player
	GameDeck Deck
	FaceUp   []Card
	Pot      int
	Code     string
	Id       uuid.UUID
	Blind    int
}

func (g Game) Init(blind int) Game {
	g.GameDeck.Init()
	g.Pot = 0
	g.FaceUp = make([]Card, 0, 5)
	g.Code = codegenerator.Code()
	g.Id = uuid.New()
	g.Blind = blind
	return g
}

func (g *Game) DealPlayersHoleCards() {
	for i, _ := range g.Players {
		for j := 0; j < 2; j++ {
			g.Players[i].Cards = append(g.Players[i].Cards, g.GameDeck.Draw())
		}
	}
}

func (g *Game) DealTable() {
	g.FaceUp = append(g.FaceUp, g.GameDeck.Draw())
}

func (g *Game) Flip() {
	_ = g.GameDeck.Draw()
	g.DealTable()
}

func (g *Game) Flop() {
	_ = g.GameDeck.Draw()
	for i := 0; i < 3; i++ {
		g.DealTable()
	}
}

func (g *Game) NewRound() {
	g.FaceUp = g.FaceUp[:0]
	g.GameDeck.Reshuffle()

	for _, player := range g.Players {
		player.Folded = false
		player.Cards = player.Cards[:0]
		player.Bet = 0
	}
}
