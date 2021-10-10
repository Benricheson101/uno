package main

import (
	"context"
	"log"
	"net"

	"github.com/benricheson101/uno/pkg"
	pb "github.com/benricheson101/uno/pkg/proto"
	"github.com/google/uuid"

	"google.golang.org/grpc"
)

const PORT = ":9000"

type UnoServer struct {
  pb.UnimplementedUnoServer
}

func (s *UnoServer) PlayCard(ctx context.Context, req *pb.PlayCardRequest) (*pb.CardStack, error) {
  c := pkg.CardFromProto(req.Card)

  playerUuid, err := uuid.FromBytes(req.Player.Id)
  if err != nil {
    log.Fatalf("failed to convert bytes to uuid: %v", err)
  }

  log.Printf("[PlayCard] player=%v | card=%v\n", playerUuid, c)

  return &pb.CardStack{}, nil
}

func (s *UnoServer) NewGame(ctx context.Context, req *pb.NewGameRequest) (*pb.Game, error) {
  game := uuid.New()
  gameId, _ := game.MarshalBinary()

  log.Printf("[NewGame] created game with id=%v\n", game.String())

  return &pb.Game{
    Id: gameId,
    Players: nil,
    CardStack: nil,
  },
  nil
}

func (s *UnoServer) NewPlayer(ctx context.Context, req *pb.NewPlayerRequest) (*pb.Player, error) {
  gameId := req.GameId
  name := req.Name
  player := uuid.New()
  playerId, _ := player.MarshalBinary()

  gameUuid, _ := uuid.FromBytes(req.GameId)
  log.Printf("[NewPlayer] created player name=%v id=%v | gameId=%v\n", name, player.String(), gameUuid.String())

  return &pb.Player{Id: playerId, GameId: gameId, Name: name}, nil
}

func main() {
  lis, err := net.Listen("tcp", PORT)
  if err != nil {
    log.Fatalf("failed to listen on PORT=%v: %v", PORT, err)
  }

  s := grpc.NewServer()
  pb.RegisterUnoServer(s, &UnoServer{})

  if err := s.Serve(lis); err != nil {
    log.Fatalf("failed to serve: %v", err)
  }
}
