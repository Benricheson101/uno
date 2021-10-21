package main

// import "fmt"

// func main() {
// 	fmt.Println("Hello, world!")
// }

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"wah.wtf/uno/card"
	"wah.wtf/uno/game"
	"wah.wtf/uno/player"
)

func main() {
	game := game.NewUnoGame()
	game.AddPlayers(
		[]player.Player{
			player.NewPlayer(),
			player.NewPlayer(),
			player.NewPlayer(),
			player.NewPlayer(),
		},
	)

	game.Init()

	stdio := bufio.NewReader(os.Stdin)

	for !game.Ended {
		turn(&game, stdio)
		winner := game.CheckWin()
		if winner != nil {
			fmt.Printf("player %v wins!", winner.Id)
			break
		}
	}
}

func turn(g *game.UnoGame, reader *bufio.Reader) {
	i, p := g.NextPlayer()

	fmt.Println("last card:", g.LastCard)
	fmt.Printf("your cards (%v): %v\n", len(p.Cards), p.Cards)

	fmt.Printf("[%v] => ", i)

	ln, _ := reader.ReadString('\n')
	ln = strings.TrimSpace(strings.ToLower(ln))

	if ln == "" {
		return
	}

	if ln == "draw" {
		g.DrawCard(p)
		return
	}

	line := strings.Split(ln, " ")

	cardColor := line[0]
	cardVal := strings.Join(line[1:], " ")

	var c *card.Card
	for _, ca := range p.Cards {
		color := strings.ToLower(ca.Color.String())
		val := strings.ToLower(ca.Val.String())

		cColor := strings.ToLower(cardColor)
		cVal := strings.ToLower(cardVal)

		if val == cVal && (val == "wild" || val == "wild draw four") {
			var col card.CardColor

			switch cColor {
			case "red":
				col = card.RED
			case "blue":
				col = card.BLUE
			case "green":
				col = card.GREEN
			case "yellow":
				col = card.YELLOW
			default:
				fmt.Println("invalid card color:", cColor)
				return
			}

			ca.Color = col
			c = &ca
			break
		} else if color == cColor && val == cVal {
			c = &ca
			break
		}
	}

	if c == nil {
		fmt.Printf("you do not have a card %v %v. select another card\n", cardColor, cardVal)
		return
	}

	fmt.Printf("PlayCard player_id=%v card_color=%v card_val=%v\n", p.Id, c.Color, c.Val)

	if err := g.PlayCard(p, *c); err != nil {
		fmt.Println("Error playing card:", err)
	}

	return
}
