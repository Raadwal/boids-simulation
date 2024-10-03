package boids

import (
	"math/rand"
	"sync"

	"github.com/Raadwal/boids-simulation/internal/config"
	"github.com/Raadwal/boids-simulation/internal/vector"
	"github.com/hajimehoshi/ebiten/v2"
)

type Boids struct {
	quantity       int
	protectedRange int
	visualRange    int
	minSpeed       float64
	maxSpeed       float64

	avoidFactor     float64
	matchingFactor  float64
	centeringfactor float64
	turnFactor      float64

	boid     []*Boid
	position [][][]int
}

var (
	boids  Boids
	rwLock = sync.RWMutex{}
)

func CreateBoids() *Boids {
	boids = Boids{
		quantity:       config.Boids.Quantity,
		protectedRange: config.Boids.ProtectedRange,
		visualRange:    config.Boids.VisualRange,
		minSpeed:       config.Boids.MinSpeed,
		maxSpeed:       config.Boids.MaxSpeed,

		avoidFactor:     config.Boids.AvoidFactor,
		matchingFactor:  config.Boids.MatchingFactor,
		centeringfactor: config.Boids.CenteringFactor,
		turnFactor:      config.Boids.TurnFactor,

		boid: make([]*Boid, config.Boids.Quantity),
	}

	minX, maxX := 0, config.Window.Width+1
	minY, maxY := 0, config.Window.Height+1

	boids.initializePositionsArray(maxX, maxY)

	rwLock.Lock()
	for i := 0; i < boids.quantity; i++ {
		boids.boid[i] = CreateBoid(
			i,
			float64(minX+rand.Intn(maxX-minX)),
			float64(minY+rand.Intn(maxY-minY)),
		)

		row := int(boids.boid[i].position.X)
		col := int(boids.boid[i].position.Y)

		boids.position[row][col] = append(boids.position[row][col], i)
	}

	rwLock.Unlock()

	return &boids
}

func (boids *Boids) initializePositionsArray(rows int, cols int) {
	boids.position = make([][][]int, rows)
	for i := 0; i < rows; i++ {
		boids.position[i] = make([][]int, cols)
		for j := 0; j < cols; j++ {
			boids.position[i][j] = make([]int, 0)
		}
	}
}

func (boids *Boids) Draw(screen *ebiten.Image) {
	for i := 0; i < boids.quantity; i++ {
		boids.boid[i].Draw(screen)
	}
}

func (boids *Boids) calculateAcceleration(boid *Boid) vector.Vector {
	separation := vector.Vector{X: 0, Y: 0} //steer to avoid crowding local flockmates
	alligment := vector.Vector{X: 0, Y: 0}  //steer towards the average heading of local flockmates
	cohesion := vector.Vector{X: 0, Y: 0}   //steer to move towards the average position (center of mass) of local flockmates

	rowStart := max(int(boid.position.X)-boids.visualRange, 0)
	rowEnd := min(int(boid.position.X)+boids.visualRange, 1080)

	colStart := max(int(boid.position.Y)-boids.visualRange, 0)
	colEnd := min(int(boid.position.Y)+boids.visualRange, 720)

	neighboringBoids := 0

	rwLock.RLock()
	for i := rowStart; i <= rowEnd; i++ {
		for j := colStart; j <= colEnd; j++ {
			for element := 0; element < len(boids.position[i][j]); element++ {
				id := boids.position[i][j][element]
				if id == boid.id {
					continue
				}

				dist := boid.position.Distance(boids.boid[id].position)
				if dist < float64(boids.protectedRange) {
					close := boid.position
					close.Subtract(boids.boid[id].position)

					separation.Add(close)
				}

				if dist < float64(boids.visualRange) {
					alligment.Add(boids.boid[id].velocity)
					cohesion.Add(boids.boid[id].position)
					neighboringBoids += 1
				}

			}
		}
	}
	rwLock.RUnlock()

	if neighboringBoids > 0 {
		alligment.DivideByScalar(float64(neighboringBoids))
		alligment.Subtract(boid.velocity)

		cohesion.DivideByScalar(float64(neighboringBoids))
		cohesion.Subtract(boid.position)
	}

	separation.MultiplyByScalar(boids.avoidFactor)
	alligment.MultiplyByScalar(boids.matchingFactor)
	cohesion.MultiplyByScalar(boids.centeringfactor)

	acceleration := vector.Vector{X: 0, Y: 0}
	acceleration.Add(separation)
	acceleration.Add(alligment)
	acceleration.Add(cohesion)
	return acceleration
}
