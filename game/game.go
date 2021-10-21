package game

import (
	"wah.wtf/uno/card"
	"wah.wtf/uno/deck"
	"wah.wtf/uno/player"
	"wah.wtf/uno/stack"
)

type UnoGame struct {
	Players    []player.Player
	Stack      stack.CardStack
	Deck       deck.Deck
	LastCard   *card.Card
	LastPlayer *player.Player
	Reverse    bool
}

func NewUnoGame(players []player.Player) UnoGame {
	deck := deck.NewDeck()
	deck.Shuffle()

	firstCard := *deck.Take()

	cardStack := stack.CardStack{}
	cardStack.Push(firstCard)

	game := UnoGame{
		Players:    players,
		LastCard:   &firstCard,
		Stack:      cardStack,
		Deck:       deck,
		LastPlayer: nil,
		Reverse:    false,
	}

	return game
}

func (g *UnoGame) Play() {
	// make game do thing
}
