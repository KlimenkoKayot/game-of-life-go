package main

import (
	"fmt"
	"time"

	"github.com/kayotklimenko/game-of-life-go/pkg/life"
)

func main() {
	world := life.NewWorld(16, 16)
	world.Seed()
	for {
		world.Print()
		time.Sleep(100 * time.Millisecond)
		world.NextState()
		fmt.Print("\033[H\033[2J")
	}
}
