package boids

import (
	"math/rand"

	"github.com/Raadwal/boids-simulation/internal/vector"
	"github.com/hajimehoshi/ebiten/v2"
)

type Boids struct {
	quantity int
	boids    []*Boid
}

var (
	boids Boids
)

func CreateBoids(quantity int, minX float64, maxX float64, minY float64, maxY float64) *Boids {
	boids = Boids{
		quantity: quantity,
		boids:    make([]*Boid, quantity),
	}

	for i := 0; i < quantity; i++ {
		boids.boids[i] = CreateBoid(
			i,
			minX+(rand.Float64()*(maxX-minX)),
			minY+(rand.Float64()*(maxY-minY)),
		)

		//fmt.Printf("%v, %v, %v\n", boids.boids[i].id, boids.boids[i].position.X, boids.boids[i].position.Y)
	}

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
	cohesion := vector.Vector{0, 0} //steer to move towards the average position (center of mass) of local flockmates

	for i := 0; i < boids.quantity; i++ {
		if boid.id == i {
			continue
		}

		//cohesion = cohesion.Add(boids.boids[i].position)
	}

	//result := vector.Vector{cohesion.X / float64(boids.quantity-1), cohesion.Y / float64(boids.quantity-1)}
	//result = result.Subtract(boid.position)
	//fmt.Println(result.X, result.Y)

	return cohesion
}
