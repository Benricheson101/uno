package pkg

// Implementation of a Stack structure used to represent a pile of cards
type CardStack []Card

// Returns if the stack is empty.
func (s *CardStack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new card onto the stack
func (s *CardStack) Push(c Card) {
	*s = append(*s, c)
}

// Remove a card from the stack. If no cards are found, a blank card will be returned.
func (s *CardStack) Pop() (Card, bool) {
	if s.IsEmpty() {
		return Card{}, false
	} else {
		topIdx := len(*s) - 1
		top := (*s)[topIdx]
		*s = (*s)[:topIdx]

		return top, true
	}
}

// Similar to CardStack.Pop(), this function returns the top card on the stack, but doesn't remove it.
func (s *CardStack) Peek() (Card, bool) {
	if s.IsEmpty() {
		return Card{}, false
	} else {
		topIdx := len(*s) - 1
		top := (*s)[topIdx]

		return top, true
	}
}
