package pkg

import (
	"fmt"

	pb "github.com/benricheson101/uno/pkg/proto"
)

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

// func CardColorFromUint(i int) CardColor {
//   switch i {
//   case 1: return Red
//   case 2: return Blue
//   case 3: return Green
//   case 4: return Yellow
//   case 5: return Black
//   default: log.Fatalf("unknown card color %v", i)
//   }
// }

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

func (c Card) Proto() pb.Card {
  return pb.Card{Color: pb.CardColor(c.Color), Val: pb.CardVal(c.Val)}
}

func CardFromProto(proto *pb.Card) Card {
  color := CardColor(uint8(proto.Color.Number()))
  val:=CardVal(uint8(proto.Val.Number()))

  return Card{Color: color, Val: val}
}
