package pkg

import "fmt"

type CardColor string

const (
	Red    CardColor = "red"
	Blue             = "blue"
	Green            = "green"
	Yellow           = "yellow"

	// Wild
	Black = "black"
)

type CardVal string

const (
	Zero  CardVal = "zero"
	One           = "one"
	Two           = "two"
	Three         = "three"
	Four          = "four"
	Five          = "five"
	Six           = "six"
	Seven         = "seven"
	Eight         = "eight"
	Nine          = "nine"

	Skip         = "skip"
	Rev          = "rev"
	DrawTwo      = "draw_two"
	Wild         = "wild"
	WildDrawFour = "wild_draw_four"
)

type Card struct {
	Color CardColor
	Val   CardVal
}

func (c Card) CanPlay(c2 Card) bool {
	return c.Color == c2.Color || c.Val == c2.Val
}

func (c Card) String() string {
	return fmt.Sprintf("%v %v", c.Color, c.Val)
}
