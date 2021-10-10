package pkg

import "github.com/google/uuid"

type Player struct {
  Id uuid.UUID
}

func NewPlayer() Player {
  return Player{Id: uuid.New()}
}
