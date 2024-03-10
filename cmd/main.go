package main

import (
	"log"

	"github.com/Raadwal/boids-simulation/internal/boid"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	single_boid = boid.CreateBoid(0, 100.0, 100.0)
)

const (
	screenWidth  = 1080
	screenHeight = 720
)

type Game struct{}

func (g *Game) Update() error {
	// Write your game's logical update.
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	single_boid.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	game := &Game{}

	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Boids simulation")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
