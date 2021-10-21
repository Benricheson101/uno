package stack

import "wah.wtf/uno/card"

type CardStack []card.Card

func (s *CardStack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *CardStack) Push(c card.Card) {
	*s = append(*s, c)
}

func (s *CardStack) Pop() (card.Card, bool) {
	if s.IsEmpty() {
		return card.Card{}, false
	} else {
		topidx := len(*s) - 1
		top := (*s)[topidx]
		*s = (*s)[:topidx]

		return top, true
	}
}

func (s *CardStack) Peek() (card.Card, bool) {
	if s.IsEmpty() {
		return card.Card{}, false
	} else {
		topidx := len(*s) - 1
		top := (*s)[topidx]

		return top, true
	}
}

func (s *CardStack) Size() int {
	return len(*s)
}
