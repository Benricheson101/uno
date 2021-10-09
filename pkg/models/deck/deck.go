package deck

import (
	"math/rand"
	"strings"
	"time"

  "github.com/benricheson101/uno/pkg/models/card"
)

type Deck []card.Card

func NewDeck() Deck {
	var cards = []card.Card{}

	colors := []card.CardColor{card.Red, card.Blue, card.Green, card.Yellow}
	vals := []card.CardVal{
		card.One,
		card.Two,
		card.Three,
		card.Four,
		card.Five,
		card.Six,
		card.Seven,
		card.Eight,
		card.Nine,
		card.Skip,
		card.Rev,
		card.DrawTwo,
	}

	specialCards := []card.Card{
		{Val: card.Wild, Color: card.Black},
		{Val: card.WildDrawFour, Color: card.Black},
	}

	for _, color := range colors {
		cards = append(cards, card.Card{Val: card.Zero, Color: color})

		for _, val := range vals {
			for i := 0; i < 2; i++ {
				cards = append(cards, card.Card{Val: val, Color: color})
			}
		}
	}

	for _, card := range specialCards {
		for i := 0; i < 4; i++ {
			cards = append(cards, card)
		}
	}

	return cards
}

func (d Deck) IsEmpty() bool {
	return len(d) == 0
}

func (d Deck) String() string {
	cards := []string{}

	for _, card := range d {
		cards = append(cards, card.String())
	}

	return strings.Join(cards, "\n")
}

func (d Deck) Len() int {
	return len(d)
}

func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())

	if !d.IsEmpty() {
		rand.Shuffle(
			len(*d),
			func(i, j int) {
				(*d)[i], (*d)[j] = (*d)[j], (*d)[i]
			},
    )
	}
}
