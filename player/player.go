package player

import "github.com/google/uuid"

type Player struct {
	Id    uuid.UUID
	Cards PlayerHand
}

func NewPlayer() Player {
	return Player{
		Id: uuid.New(),
	}
}
