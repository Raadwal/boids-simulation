package boids

import (
	"image/color"
	"math"
	"math/rand"
	"time"

	"github.com/Raadwal/boids-simulation/internal/vector"
	"github.com/hajimehoshi/ebiten/v2"
)

type Boid struct {
	id            int
	position      vector.Vector
	velocity      vector.Vector
	width         int
	height        int
	color         color.RGBA
	stepSleep     time.Duration
	whiteImage    *ebiten.Image
	triangleImage *ebiten.Image
}

func CreateBoid(id int, posX float64, posY float64) *Boid {
	boid := Boid{
		id:       id,
		position: vector.Vector{X: posX, Y: posY},
		velocity: vector.Vector{
			X: (rand.Float64() * 2) - 1.0,
			Y: (rand.Float64() * 2) - 1.0,
		},
		width:         10,
		height:        20,
		color:         color.RGBA{10, 255, 50, 255},
		stepSleep:     5,
		triangleImage: ebiten.NewImage(10, 20),
		whiteImage:    ebiten.NewImage(10, 20),
	}
	boid.whiteImage.Fill(color.White)
	boid.triangleImage.Fill(color.Transparent)

	vertices := []ebiten.Vertex{
		{
			DstX:   float32(boid.width / 2),
			DstY:   0.0,
			ColorR: float32(boid.color.R), ColorG: float32(boid.color.G), ColorB: float32(boid.color.B),
			ColorA: float32(boid.color.A),
		},
		{
			DstX:   0.0,
			DstY:   float32(boid.height),
			ColorR: float32(boid.color.R), ColorG: float32(boid.color.G), ColorB: float32(boid.color.B),
			ColorA: float32(boid.color.A),
		},
		{
			DstX:   float32(boid.width),
			DstY:   float32(boid.height),
			ColorR: float32(boid.color.R), ColorG: float32(boid.color.G), ColorB: float32(boid.color.B),
			ColorA: float32(boid.color.A),
		},
	}

	indices := []uint16{0, 1, 2}
	opt := ebiten.DrawTrianglesOptions{}
	opt.AntiAlias = true

	boid.triangleImage.DrawTriangles(vertices, indices, boid.whiteImage, &opt)

	go boid.start()
	return &boid
}

func (boid *Boid) Draw(screen *ebiten.Image) {
	options := ebiten.DrawImageOptions{}

	centerX, centerY := float64(boid.width/2), float64(boid.height/2)
	options.GeoM.Translate(-centerX, -centerY)
	options.GeoM.Rotate(boid.calculateRotation())
	options.GeoM.Translate(boid.position.X, boid.position.Y)

	screen.DrawImage(boid.triangleImage, &options)
}

func (boid *Boid) step() {
	boid.velocity = boid.velocity.Add(boids.calculateAcceleration(boid))
	boid.position = boid.position.Add(boid.velocity)
	boid.bounceIfNeeded()
}

func (boid *Boid) start() {
	for {
		boid.step()
		time.Sleep(boid.stepSleep * time.Millisecond)
	}
}

func (boid *Boid) calculateRotation() float64 {
	angleRad := math.Atan2(boid.velocity.Y, boid.velocity.X) + math.Pi/2
	return angleRad
}

func (boid *Boid) bounceIfNeeded() {
	if boid.position.X < 0 || boid.position.X > 1080 {
		boid.velocity.X = -boid.velocity.X
	}
	if boid.position.Y < 0 || boid.position.Y > 720 {
		boid.velocity.Y = -boid.velocity.Y
	}
}
