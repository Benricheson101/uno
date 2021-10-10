package pkg

import "fmt"

type CardColor uint8

const (
	Red CardColor = iota
	Blue
	Green
	Yellow

	// Wild
	Black
)

func (c CardColor) String() string {
	switch c {
	case Red:
		return "red"
	case Blue:
		return "blue"
	case Green:
		return "green"
	case Yellow:
		return "yellow"
	case Black:
		return "black"
	default:
		return ""
	}
}

type CardVal uint8

const (
	Zero  CardVal = iota
	One
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine

	Skip
	Rev
	DrawTwo
	Wild
	WildDrawFour
)

func (v CardVal) String() string {
	switch v {
	case Zero:
		return "zero"
	case One:
		return "one"
	case Two:
		return "two"
	case Three:
		return "three"
	case Four:
		return "four"
	case Five:
		return "five"
	case Six:
		return "six"
	case Seven:
		return "seven"
	case Eight:
		return "eight"
	case Nine:
		return "nine"

	case Skip:
		return "skip"
	case Rev:
		return "rev"
	case DrawTwo:
		return "draw_two"
	case Wild:
		return "wild"
	case WildDrawFour:
		return "wild_draw_four"
	default:
		return ""
	}
}

type Card struct {
	Color CardColor
	Val   CardVal
}

func (c Card) CanPlay(c2 Card) bool {
	return c.Color == c2.Color || c.Val == c2.Val
}

func (c Card) String() string {
	return fmt.Sprintf("%v %v", c.Color.String(), c.Val.String())
}
