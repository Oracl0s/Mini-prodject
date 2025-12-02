package p

import (
	"fmt"
	"time"
)

type Game struct {
	Name   string
	Age    int
	Number string
	point  int
}

func Gamemode() {
	var castomString Game = Game{"Ivan", 23, "3423", 14}
	fmt.Println(castomString)
	castomString.Number = "123"
	fmt.Println(castomString)

	Jon := Game{
		Name:   "Jon",
		Age:    12,
		Number: time.Now().String(),
	}
	fmt.Println(Jon.Number)
}
