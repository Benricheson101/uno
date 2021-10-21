package player

import "wah.wtf/uno/card"

type PlayerHand []card.Card

func (h *PlayerHand) Add(c card.Card) {
	*h = append(*h, c)
}
