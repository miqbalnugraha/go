package main

import "fmt"

type Gamer struct {
	Name  string
	Games []string
}

func (gamer *Gamer) AddGame(game string) {
	gamer.Games = append(gamer.Games, game)
}

func main() {
	gamer := Gamer{Name: "Mobile Legends"}

	gamer.AddGame("Borderlands")
	gamer.AddGame("Fifa 2022")
	gamer.AddGame("Point Blank")

	for _, game := range gamer.Games {
		fmt.Println(game)
	}
}
