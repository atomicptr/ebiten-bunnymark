package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const windowWidth = 800
const windowHeight = 600
const defaultNumBunnies = 128
const ticksPerSecond = ebiten.DefaultTPS
const delta = float64(1.0) / ticksPerSecond

var bunnyPool []*Bunny

var numBunnies = defaultNumBunnies
var spaceJustPressed = false

func init() {
	flag.IntVar(&numBunnies, "num-bunnies", defaultNumBunnies, "number of bunnies on the screen")
	flag.Parse()

	createBunnies()
}

func main() {
	ebiten.SetMaxTPS(ticksPerSecond)
	ebiten.SetRunnableInBackground(true)

	if err := ebiten.Run(update, windowWidth, windowHeight, 1, "Bunny Mark"); err != nil {
		log.Fatal(err)
	}
}

func createBunnies() {
	missing := numBunnies - len(bunnyPool)

	for i := 0; i < missing; i++ {
		bunnyPool = append(bunnyPool, NewBunny(windowWidth*0.5, windowHeight*0.5))
	}
}

func update(screen *ebiten.Image) error {
	if spaceJustPressed && !ebiten.IsKeyPressed(ebiten.KeySpace) {
		spaceJustPressed = false
	}

	if !spaceJustPressed && ebiten.IsKeyPressed(ebiten.KeySpace) {
		numBunnies *= 2

		createBunnies()
		spaceJustPressed = true
	}

	// update bunny positions
	for _, bunny := range bunnyPool {
		bunny.Update()
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	RenderBackground(screen)

	// draw bunnies
	for _, bunny := range bunnyPool {
		bunny.Draw(screen)
	}

	printDebugLines(screen, []string {
		fmt.Sprintf("FPS: %.2f", ebiten.CurrentFPS()),
		fmt.Sprintf("TPS: %.2f", ebiten.CurrentTPS()),
		fmt.Sprintf("Number of Bunnies: %d", numBunnies),
		"Press SPACE to double the number of bunnies!",
	})

	return nil
}

func printDebugLines(screen *ebiten.Image, lines []string) {
	y := 10

	for _, line := range lines {
		ebitenutil.DebugPrintAt(screen, line, 10, y)
		y += 20
	}
}