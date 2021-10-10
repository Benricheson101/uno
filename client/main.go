package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	"github.com/benricheson101/uno/pkg"
	pb "github.com/benricheson101/uno/pkg/proto"
	"github.com/google/uuid"
)

const ADDR = "localhost:9000"

func main() {
  conn, err := grpc.Dial(ADDR, grpc.WithInsecure())

  if err != nil {
    log.Fatalf("failed to dial add=%v: %v", ADDR, err)
  }
  defer conn.Close()

  c := pb.NewUnoClient(conn)

  ctx, cancel := context.WithTimeout(context.Background(), time.Second)
  defer cancel()

  game, err := c.NewGame(ctx, &pb.NewGameRequest{})
  if err != nil {
    log.Fatalf("failed to create game: %v\n", err)
  }

  gameId, _ := uuid.FromBytes(game.Id)

  log.Printf("created game id=%v\n", gameId.String())

  player, _ := c.NewPlayer(ctx, &pb.NewPlayerRequest{GameId: game.Id, Name: "ben"})

  playerId, _ := uuid.FromBytes(player.Id)

  log.Printf("created user name=%v id=%v\n", player.Name, playerId.String())

  card := pkg.Card{Color: pkg.Blue, Val: pkg.DrawTwo}
  cardProto := card.Proto()

  newStack, _ := c.PlayCard(ctx, &pb.PlayCardRequest{Player: player, Card: &cardProto})

  log.Printf("newStack=%+v\n", newStack.String())
}

// TODO:
//  - Cards
//  - Deck
//  - Hand
//  - Stack

//  - server
//  - clients
//    - commect woth gRPC or WebSocket
