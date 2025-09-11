package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

var (
	dice       []*ebiten.Image
	poweredBy  *ebiten.Image
	pressSpace *ebiten.Image
)

func init() {
	for i := 1; i <= 6; i++ {
		d, _, err := ebitenutil.NewImageFromFile(fmt.Sprintf("assets/dice-%d.png", i))
		if err != nil {
			log.Fatal(err)
		}
		dice = append(dice, d)
	}

	var err error
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
