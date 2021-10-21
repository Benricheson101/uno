package deck_test

import (
	"testing"

	"wah.wtf/uno/deck"
)

func TestDeckRemaining(t *testing.T) {
	deck := deck.NewDeck()

	for i := 0; i < 50; i++ {
		deck.Take()
	}

	remaining := deck.Remaining()

	if remaining != 58 {
		t.Errorf("expected: %v, got: %v", 58, remaining)
	}
}

func TestDeckShuffle(t *testing.T) {
	unshuffledDeck := deck.NewDeck()
	shuffledDeck := deck.NewDeck()
	shuffledDeck.Shuffle()

	if unshuffledDeck.Remaining() != shuffledDeck.Remaining() {
		t.Errorf("deck changed size. expected: %v, got: %v", unshuffledDeck.Remaining(), shuffledDeck.Remaining())
	}

	same := 0

	for i, c := range unshuffledDeck.Cards {
		if c == shuffledDeck.Cards[i] {
			same++
		}
	}

	if same == shuffledDeck.Remaining() {
		t.Error("shuffled and unshuffled decks are the same")
	}
}
