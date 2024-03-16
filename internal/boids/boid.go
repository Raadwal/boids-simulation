package boids

import (
	"image/color"
	"math"
	"math/rand"
	"time"

	"github.com/Raadwal/boids-simulation/internal/config"
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
	prev_position := boid.position

	acceleration := boids.calculateAcceleration(boid)
	boid.velocity.Add(acceleration)
	boid.applySpeedLimits()
	boid.turnAroundIfNeeded()

	boid.position.Add(boid.velocity)
	boid.bounceIfNeeded()
	boid.updateArray(prev_position)
}

func (boid *Boid) updateArray(prev_position vector.Vector) {
	prevRow := int(prev_position.X)
	prevCol := int(prev_position.Y)

	cur_row := int(boid.position.X)
	cur_col := int(boid.position.Y)

	//fmt.Println(cur_row, cur_col, "--", boid.position.X, boid.position.Y)

	index := -1

	rwLock.Lock()
	for i := 0; i < len(boids.position[prevRow][prevCol]); i++ {
		if boids.position[prevRow][prevCol][i] == boid.id {
			//println(boids.positions[prevRow][prevCol][i], boid.id, i)
			index = i
			break
		}
	}
	//fmt.Println(index)
	//rwLock.RUnlock()

	//rwLock.Lock()
	if index != -1 {
		//fmt.Println(index, len(boids.positions[prevRow][prevCol]), boids.positions[prevRow][prevCol][index])
		boids.position[prevRow][prevCol] = append(boids.position[prevRow][prevCol][:index], boids.position[prevRow][prevCol][index+1:]...)
	}

	boids.position[cur_row][cur_col] = append(boids.position[cur_row][cur_col], boid.id)
	rwLock.Unlock()
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
	if boid.position.X < 0 || boid.position.X > float64(config.Window.Width) {
		boid.velocity.X = -boid.velocity.X
		boid.position.X = math.Max(0, math.Min(float64(config.Window.Width), boid.position.X))
	}
	if boid.position.Y < 0 || boid.position.Y > float64(config.Window.Height) {
		boid.velocity.Y = -boid.velocity.Y
		boid.position.Y = math.Max(0, math.Min(float64(config.Window.Height), boid.position.Y))
	}
}

func (boid *Boid) turnAroundIfNeeded() {
	rwLock.Lock()
	if boid.position.X < config.Boids.ScreenMargin {
		boid.velocity.X = boid.velocity.X + boids.turnFactor
	} else if boid.position.X > float64(config.Window.Width)-config.Boids.ScreenMargin {
		boid.velocity.X = boid.velocity.X - boids.turnFactor
	}

	if boid.position.Y < config.Boids.ScreenMargin {
		boid.velocity.Y = boid.velocity.Y + boids.turnFactor
	} else if boid.position.Y > float64(config.Window.Height)-config.Boids.ScreenMargin {
		boid.velocity.Y = boid.velocity.Y - boids.turnFactor
	}
	rwLock.Unlock()
}

func (boid *Boid) applySpeedLimits() {
	speed := math.Sqrt(math.Pow(boid.velocity.X, 2) + math.Pow(boid.velocity.Y, 2))

	if speed > float64(boids.maxSpeed) {
		boid.velocity.DivideByScalar(speed)
		boid.velocity.MultiplyByScalar(float64(boids.maxSpeed))
	}

	if speed < float64(boids.minSpeed) {
		boid.velocity.DivideByScalar(speed)
		boid.velocity.MultiplyByScalar(float64(boids.maxSpeed))
	}
}
