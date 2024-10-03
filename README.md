# Boids simulation
 Boids simulation utilizing multithreading created with Go.

### Simulation
<img src="recording\Simulation.gif" alt="Boids simulation" width="1080">

The settings of the simulation shown above:
```json
{
    "window": {
        "width": 1080,
        "height": 720
    },
    "boids": {
        "quantity": 1000,
        "protectedRange": 5,
        "visualRange": 35,
        "avoidFactor": 0.15,
        "matchingFactor": 0.05,
        "centeringfactor": 0.0005,
        "minSpeed": 2,
        "maxSpeed": 5,
        "screenMargin": 100,
        "turnFactor": 0.2
    }
}
```