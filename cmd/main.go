package main

import (
	"log"

	"github.com/Raadwal/boids-simulation/internal/boids"
	"github.com/Raadwal/boids-simulation/internal/config"
	"github.com/hajimehoshi/ebiten/v2"
)

func init() {
	err := config.LoadConfig("config/config.json", true)
	if err != nil {
		log.Fatal("Error when loading config:", err)
	}

	boidsCollection = boids.CreateBoids()
}

var (
	boidsCollection *boids.Boids
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	boidsCollection.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return config.Window.Width, config.Window.Height
}

func main() {
	game := &Game{}

	ebiten.SetWindowSize(config.Window.Width, config.Window.Height)
	ebiten.SetWindowTitle("Boids simulation")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
