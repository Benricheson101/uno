package pkg

import (
	"math/rand"
	"strings"
	"time"
)

type Deck []Card

func NewDeck() Deck {
	var cards = []Card{}

	colors := []CardColor{Red, Blue, Green, Yellow}
	vals := []CardVal{
		One,
		Two,
		Three,
		Four,
		Five,
		Six,
		Seven,
		Eight,
		Nine,
		Skip,
		Rev,
		DrawTwo,
	}

	specialCards := []Card{
		{Val: Wild, Color: Black},
		{Val: WildDrawFour, Color: Black},
	}

	for _, color := range colors {
		cards = append(cards, Card{Val: Zero, Color: color})

		for _, val := range vals {
			for i := 0; i < 2; i++ {
				cards = append(cards, Card{Val: val, Color: color})
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
