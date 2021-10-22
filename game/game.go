package game

import (
	"errors"

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
	Ended      bool
	Running    bool
}

func NewUnoGame() UnoGame {
	deck := deck.NewDeck()
	deck.Shuffle()

	hasFirstCard := false
	var firstCard card.Card
	for !hasFirstCard {
		c := *deck.Take()
		if c.Color != card.BLACK &&
			c.Val != card.DRAW_TWO &&
			c.Val != card.SKIP &&
			c.Val != card.REV {
			firstCard = c
			hasFirstCard = true
		} else {
			deck.Cards = append(deck.Cards, c)
		}
	}

	cardStack := stack.CardStack{}
	cardStack.Push(firstCard)

	game := UnoGame{
		LastCard:   &firstCard,
		Stack:      cardStack,
		Deck:       deck,
		LastPlayer: nil,
		Reverse:    false,
		Ended:      false,
		Running:    false,
	}

	return game
}

func (g *UnoGame) AddPlayer(player player.Player) {
	g.Players = append(g.Players, player)
}

func (g *UnoGame) AddPlayers(players []player.Player) {
	for _, p := range players {
		g.AddPlayer(p)
	}
}

func (g *UnoGame) Init() error {
	if !g.Running && !g.Ended {
		g.Deck.Shuffle()
		g.Deck.Deal(&g.Players)
		g.Running = true
	} else if g.Ended {
		return errors.New("game already finished")
	} else if g.Running {
		return errors.New("game already running")
	}

	return nil
}

func (g *UnoGame) PlayCard(p *player.Player, c card.Card) error {
	if g.Ended {
		return errors.New("game already ended")
	}

	_, np := g.NextPlayer()
	if np.Id != p.Id {
		return errors.New("wrong player")
	}

	hasCard := false
	cardIndex := 0

	for i, ca := range p.Cards {
		if (ca.Val == card.WILD || ca.Val == card.WILD_DRAW_FOUR) && ca.Val == c.Val {
			hasCard = true
			cardIndex = i
			break
		}

		if ca.Color == c.Color && ca.Val == c.Val {
			hasCard = true
			cardIndex = i
			break
		}
	}

	if !hasCard {
		return errors.New("player does not have card")
	}

	g.LastPlayer = p

	if g.LastCard.CanPlay(&c) {
		switch c.Val {
		case card.REV:
			g.Reverse = !g.Reverse

		case card.DRAW_TWO:
			_, next := g.NextPlayer()
			next.Cards.Add(*g.Deck.Take())
			next.Cards.Add(*g.Deck.Take())
			g.LastPlayer = next

		case card.SKIP:
			_, skipped := g.NextPlayer()
			g.LastPlayer = skipped

		case card.WILD_DRAW_FOUR:
			_, next := g.NextPlayer()
			for i := 0; i < 4; i++ {
				next.Cards.Add(*g.Deck.Take())
			}
			g.LastPlayer = next

		case card.WILD:
			if c.Color == card.BLACK {
				return errors.New("wild color cannot be black")
			}

		}

		g.Stack.Push(c)
		g.LastCard = &c

		p.Cards = append(p.Cards[:cardIndex], p.Cards[cardIndex+1:]...)
		return nil
	}

	return errors.New("cannot play that card")
}

func (g *UnoGame) DrawCard(p *player.Player) {
	p.Cards.Add(*g.Deck.Take())
	g.LastPlayer = p
}

func (g *UnoGame) CheckWin() *player.Player {
	for _, player := range g.Players {
		if len(player.Cards) == 0 {
			return &player
		}
	}

	return nil
}

func (g *UnoGame) NextPlayer() (int, *player.Player) {
	last := g.LastPlayer
	if last == nil {
		return 0, &g.Players[0]
	}

	var index int

	for idx, player := range g.Players {
		if player.Id == g.LastPlayer.Id {
			index = idx
			break
		}
	}

	var nextIndex int

	if !g.Reverse {
		nextIndex = index + 1
		if nextIndex >= len(g.Players) {
			nextIndex = 0
		}
	} else {
		nextIndex = index - 1
		if nextIndex < 0 {
			nextIndex = len(g.Players) - 1
		}
	}

	return nextIndex, &g.Players[nextIndex]
}
