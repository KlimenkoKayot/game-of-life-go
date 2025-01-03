package service

import "github.com/klimenkokayot/game-of-life-go/pkg/life"

type LifeService struct {
	World *life.World
}

func New(height, width int) *LifeService {
	world := life.NewWorld(height, width)
	service := LifeService{
		World: world,
	}
	return &service
}
