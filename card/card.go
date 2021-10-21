package card

type CardColor uint8

const (
	RED CardColor = iota
	BLUE
	GREEN
	YELLOW
	BLACK
)

type CardValue uint8

const (
	ZERO CardValue = iota
	ONE
	TWO
	THREE
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
	SKIP
	REV
	WILD
	WILD_DRAW_FOUR
	DRAW_TWO
)

type Card struct {
	Val   CardValue
	Color CardColor
}

func (cc CardColor) String() string {
	switch cc {
	case RED:
		return "red"
	case BLUE:
		return "blue"
	case GREEN:
		return "green"
	case YELLOW:
		return "yellow"
	case BLACK:
		return "black"
	}
	panic("encountered invalid card color: " + string(cc))
}

func (cv CardValue) String() string {
	switch cv {
	case ZERO:
		return "zero"
	case ONE:
		return "one"
	case TWO:
		return "two"
	case THREE:
		return "three"
	case FOUR:
		return "four"
	case FIVE:
		return "five"
	case SIX:
		return "six"
	case SEVEN:
		return "seven"
	case EIGHT:
		return "eight"
	case NINE:
		return "nine"
	case SKIP:
		return "skip"
	case REV:
		return "reverse"
	case WILD:
		return "wild"
	case WILD_DRAW_FOUR:
		return "wild draw four"
	case DRAW_TWO:
		return "draw two"
	}
	panic("encountered invalid card value: " + string(cv))
}

func (c *Card) CanPlay(card *Card) bool {
	return c.Color == card.Color || c.Val == card.Val
}

func (c Card) String() string {
	return c.Color.String() + " " + c.Val.String()
}
