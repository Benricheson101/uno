package stack_test

import (
	"testing"

	"wah.wtf/uno/deck"
	"wah.wtf/uno/stack"
)

func TestStack(t *testing.T) {
	stack := stack.CardStack{}
	deck := deck.NewDeck()

	if size := stack.Size(); size != 0 {
		t.Errorf("expected stack size: 0, got: %v", size)
	}

	stack.Push(*deck.Take())
	stack.Push(*deck.Take())
	stack.Push(*deck.Take())

	if size := stack.Size(); size != 3 {
		t.Errorf("expected stack size: 3, got: %v", size)
	}
}
