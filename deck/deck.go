package deck

import (
	"encoding/json"
	"errors"
	"math/rand"
	"time"

	. "wah.wtf/uno/card"
	"wah.wtf/uno/player"
)

type Deck struct {
	Cards []Card
}

// Creates a new deck containing the correct number of cards
func NewDeck() Deck {
	var cards = []Card{}

	colors := []CardColor{RED, BLUE, GREEN, YELLOW}
	vals := []CardValue{
		ONE,
		TWO,
		THREE,
		FOUR,
		FIVE,
		SIX,
		SEVEN,
		EIGHT,
		NINE,
		SKIP,
		REV,
		DRAW_TWO,
	}

	specialCards := []Card{
		{Val: WILD, Color: BLACK},
		{Val: WILD_DRAW_FOUR, Color: BLACK},
	}

	for _, color := range colors {
		cards = append(cards, Card{Val: ZERO, Color: color})

		for _, val := range vals {
			c := Card{Val: val, Color: color}
			cards = append(cards, c, c)
		}
	}

	for _, card := range specialCards {
		cards = append(cards, card, card, card, card)
	}

	return Deck{Cards: cards}
}

// The number of cards left in the deck
func (d *Deck) Remaining() int {
	return len(d.Cards)
}

// Take the top card off the deck and return it
func (d *Deck) Take() *Card {
	var top Card
	top, d.Cards = d.Cards[0], d.Cards[1:]

	return &top
}

// Shuffle the deck in-place
func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())

	rand.Shuffle(len(d.Cards), func(i, j int) { d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i] })
}

func (d *Deck) Deal(players *[]player.Player) error {
	if len(*players) > 10 {
		return errors.New("too many players")
	}

	for i := 0; i < 7; i++ {
		for j := 0; j < len(*players); j++ {
			(*players)[j].Cards = append((*players)[j].Cards, *d.Take())
		}
	}

	return nil
}

func (d Deck) String() string {
	data, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "[]"
	}

	return string(data)
}
