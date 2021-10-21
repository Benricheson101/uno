package card_test

import (
	"testing"

	"wah.wtf/uno/card"
)

func TestCardCanPlaySameNumber(t *testing.T) {
	cardA := card.Card{Color: card.BLUE, Val: card.NINE}
	cardB := card.Card{Color: card.GREEN, Val: card.NINE}

	if !cardA.CanPlay(&cardB) {
		t.Fail()
	}
}

func TestCardCanPlaySameColor(t *testing.T) {
	cardA := card.Card{Color: card.BLUE, Val: card.NINE}
	cardB := card.Card{Color: card.BLUE, Val: card.EIGHT}

	if !cardA.CanPlay(&cardB) {
		t.Fail()
	}
}
