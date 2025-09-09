package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

var (
	dice1      *ebiten.Image
	dice2      *ebiten.Image
	dice3      *ebiten.Image
	dice4      *ebiten.Image
	dice5      *ebiten.Image
	dice6      *ebiten.Image
	poweredBy  *ebiten.Image
	pressSpace *ebiten.Image
)

func init() {
	var err error
	dice1, _, err = ebitenutil.NewImageFromFile("assets/dice-1.png")
	if err != nil {
		log.Fatal(err)
	}
	dice2, _, err = ebitenutil.NewImageFromFile("assets/dice-2.png")
	if err != nil {
		log.Fatal(err)
	}
	dice3, _, err = ebitenutil.NewImageFromFile("assets/dice-3.png")
	if err != nil {
		log.Fatal(err)
	}
	dice4, _, err = ebitenutil.NewImageFromFile("assets/dice-4.png")
	if err != nil {
		log.Fatal(err)
	}
	dice5, _, err = ebitenutil.NewImageFromFile("assets/dice-5.png")
	if err != nil {
		log.Fatal(err)
	}
	dice6, _, err = ebitenutil.NewImageFromFile("assets/dice-6.png")
	if err != nil {
		log.Fatal(err)
	}
	poweredBy, _, err = ebitenutil.NewImageFromFile("assets/powered-by.png")
	if err != nil {
		log.Fatal(err)
	}
	pressSpace, _, err = ebitenutil.NewImageFromFile("assets/press-space.png")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Roll Dice")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
