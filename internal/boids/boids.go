package boids

import (
	"math/rand"
	"sync"

	"github.com/Raadwal/boids-simulation/internal/vector"
	"github.com/hajimehoshi/ebiten/v2"
)

type Boids struct {
	quantity int
	boids    []*Boid
}

var (
	boids  Boids
	rwLock = sync.RWMutex{}
)

func CreateBoids(quantity int, minX float64, maxX float64, minY float64, maxY float64) *Boids {
	boids = Boids{
		quantity: quantity,
		boids:    make([]*Boid, quantity),
	}

	rwLock.Lock()
	for i := 0; i < quantity; i++ {
		boids.boids[i] = CreateBoid(
			i,
			minX+(rand.Float64()*(maxX-minX)),
			minY+(rand.Float64()*(maxY-minY)),
		)

		//fmt.Printf("%v, %v, %v\n", boids.boids[i].id, boids.boids[i].position.X, boids.boids[i].position.Y)
	}
	rwLock.Unlock()

	return &boids
}

func (boids *Boids) Draw(screen *ebiten.Image) {
	for i := 0; i < boids.quantity; i++ {
		boids.boids[i].Draw(screen)
		//fmt.Printf("%v, %v, %v\n", boids.boids[i].id, boids.boids[i].position.X, boids.boids[i].position.Y)
	}
}

func (boids *Boids) calculateAcceleration(boid *Boid) vector.Vector {
	//separation := vector.Vector{0, 0} //steer to avoid crowding local flockmates
	//allignment := vector.Vector{0, 0} //steer towards the average heading of local flockmates
	//cohesion := vector.Vector{0, 0} //steer to move towards the average position (center of mass) of local flockmates
	avgPosition := vector.Vector{0, 0}

	rwLock.RLock()
	for i := 0; i < boids.quantity; i++ {
		if boid.id == i {
			continue
		}

		//fmt.Println(boids.boids[i].position.X)
		avgPosition.Add(boids.boids[i].position)

	}
	rwLock.RUnlock()

	acceleration := vector.Vector{0, 0}

	return acceleration
}
