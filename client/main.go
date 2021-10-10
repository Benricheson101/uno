package main

import (
	"fmt"

	"github.com/benricheson101/uno/pkg"
)

func main() {
	d := pkg.NewDeck()

	fmt.Println(d)
	fmt.Println(d.Len())

	d.Shuffle()

	fmt.Println(d)
	fmt.Println(d.Len())
}

// TODO:
//  - Cards
//  - Deck
//  - Hand
//  - Stack

//  - server
//  - clients
//    - commect woth gRPC or WebSocket
