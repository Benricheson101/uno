package game_test

import (
	"testing"

	"wah.wtf/uno/game"
	"wah.wtf/uno/player"
)

func TestGamePlayerIncrementer(t *testing.T) {
	player0 := player.NewPlayer()
	player1 := player.NewPlayer()
	player2 := player.NewPlayer()
	player3 := player.NewPlayer()

	game := game.NewUnoGame()
	game.AddPlayers(
		[]player.Player{
			player0,
			player1,
			player2,
			player3,
		},
	)

	if _, np := game.NextPlayer(); np.Id != player0.Id {
		t.Errorf("expected: %v, got: %v", player0.Id, np.Id)
	} else {
		game.LastPlayer = np
	}

	if _, np := game.NextPlayer(); np.Id != player1.Id {
		t.Errorf("expected: %v, got: %v", player1.Id, np.Id)
	} else {
		game.LastPlayer = np
	}

	if _, np := game.NextPlayer(); np.Id != player2.Id {
		t.Errorf("expected: %v, got: %v", player2.Id, np.Id)
	} else {
		game.LastPlayer = np
	}

	if _, np := game.NextPlayer(); np.Id != player3.Id {
		t.Errorf("expected: %v, got: %v", player3.Id, np.Id)
	} else {
		game.LastPlayer = np
	}

	if _, np := game.NextPlayer(); np.Id != player0.Id {
		t.Errorf("expected: %v, got: %v", player0.Id, np.Id)
	} else {
		game.LastPlayer = np
	}

	game.Reverse = true

	if _, np := game.NextPlayer(); np.Id != player3.Id {
		t.Errorf("expected: %v, got: %v", player3.Id, np.Id)
	} else {
		game.LastPlayer = np
	}
}
