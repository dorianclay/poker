package card

type Suit int

const (
	Clubs Suit = iota
	Diamonds
	Hearts
	Spades
)

type Card struct {
	value int
	suit  Suit
}

func (c Card) ValueToFace() string {
	switch c.value {
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
	switch c.suit {
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
