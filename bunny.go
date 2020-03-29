package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"golang.org/x/image/math/f64"
)

var bunnyImage *ebiten.Image
var bunnyWidth float64
var bunnyHeight float64

const speedMultiplier = 50

// Bunny represents a bunny on the screen
type Bunny struct {
	Velocity f64.Vec2
	Position f64.Vec2
}

// Update the position of the bunny
func (b *Bunny) Update() {
	// check if bunny is out of bounds on the X axis
	if b.Position[0] < 0 || b.Position[0] > windowWidth-bunnyWidth {
		newX := 0.0
		if b.Position[0] > windowWidth-bunnyWidth {
			newX = windowWidth - bunnyWidth
		}

		b.Position[0] = newX
		b.Velocity[0] *= -1
	}

	// check if bunny is out of bounds on the Y axis
	if b.Position[1] < 0 || b.Position[1] > windowHeight-bunnyHeight {
		newY := 0.0
		if b.Position[1] > windowHeight-bunnyHeight {
			newY = windowHeight - bunnyHeight
		}

		b.Position[1] = newY
		b.Velocity[1] *= -1
	}

	// apply velocity
	b.Position[0] += b.Velocity[0] * delta
	b.Position[1] += b.Velocity[1] * delta
}

// Draw the bunny on a given screen
func (b *Bunny) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(b.Position[0], b.Position[1])
	_ = screen.DrawImage(bunnyImage, op)
}

// NewBunny creates a new bunny with random speeds
func NewBunny(x, y float64) *Bunny {
	return &Bunny{
		Velocity: [2]float64{
			speedMultiplier * random(-5, 5),
			speedMultiplier * random(-7.5, 2.5),
		},
		Position: [2]float64{x, y},
	}
}

func random(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func init() {
	rand.Seed(time.Now().UnixNano())

	image, _, err := ebitenutil.NewImageFromFile("./assets/wabbit_alpha.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	bunnyImage = image

	w, h := bunnyImage.Size()
	bunnyWidth = float64(w)
	bunnyHeight = float64(h)
}
