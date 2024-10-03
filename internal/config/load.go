package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type BoidsConfig struct {
	Quantity       int `json:"quantity"`
	ProtectedRange int `json:"protectedRange"`
	VisualRange    int `json:"visualRange"`

	AvoidFactor     float64 `json:"avoidFactor"`
	MatchingFactor  float64 `json:"matchingFactor"`
	CenteringFactor float64 `json:"centeringfactor"`
	TurnFactor      float64 `json:"turnFactor"`
	ScreenMargin    float64 `json:"screenMargin"`
	MinSpeed        float64 `json:"minSpeed"`
	MaxSpeed        float64 `json:"maxSpeed"`
}

type WindowConfig struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type AppConfig struct {
	Boids  BoidsConfig  `json:"boids"`
	Window WindowConfig `json:"window"`
}

var (
	Boids  BoidsConfig
	Window WindowConfig

	config AppConfig
)

func LoadConfig(pathToConfig string, verbose bool) error {
	content, err := os.ReadFile(pathToConfig)
	if err != nil {
		return fmt.Errorf("error when opening file: %w", err)
	}

	err = json.Unmarshal(content, &config)
	if err != nil {
		return fmt.Errorf("error during Unmarshal: %w", err)
	}

	if verbose {
		printConfig()
	}

	Boids = config.Boids
	Window = config.Window

	return nil
}

func printConfig() {
	width := 25

	fmt.Println(strings.Repeat("=", 10), "Window Config", strings.Repeat("=", 10))
	fmt.Printf("%-*s (%v, %v)\n", width, "Window size:", config.Window.Width, config.Window.Height)

	fmt.Println(strings.Repeat("=", 10), "Boids Config", strings.Repeat("=", 10))
	fmt.Printf("%-*s %v\n", width, "Quantity:", config.Boids.Quantity)
	fmt.Printf("%-*s %v\n", width, "Protected range:", config.Boids.ProtectedRange)
	fmt.Printf("%-*s %v\n", width, "Visual range:", config.Boids.VisualRange)
	fmt.Printf("%-*s <%v, %v>\n", width, "Speed range:", config.Boids.MinSpeed, config.Boids.MaxSpeed)
	fmt.Printf("%-*s %v\n", width, "Avoid factor:", config.Boids.AvoidFactor)
	fmt.Printf("%-*s %v\n", width, "Matching factor:", config.Boids.MatchingFactor)
	fmt.Printf("%-*s %v\n", width, "Centering factor:", config.Boids.CenteringFactor)
	fmt.Printf("%-*s %v\n", width, "Turn factor:", config.Boids.TurnFactor)
	fmt.Printf("%-*s %v\n", width, "Screen margin:", config.Boids.ScreenMargin)
}
