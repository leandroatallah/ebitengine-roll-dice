package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Mode int

const (
	ModeInit Mode = iota
	ModePlay
	ModeRolling
)

type Game struct {
	count int
	mode  Mode
}

func (g *Game) Update() error {
	g.count++

	switch g.mode {
	case ModeInit:
		if g.count > 100 {
			// TODO: Make smooth transition
			g.mode = ModePlay
		}
	case ModePlay:
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			// TODO: Make smooth transition
			g.mode = ModeRolling
			fmt.Println("SPACE")
		}
	case ModeRolling:
		fmt.Println("ROLLING")
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x21, 0x23, 0x43, 0xff})

	switch g.mode {
	case ModeInit:
		frameWidth, frameHeight := poweredBy.Bounds().Dx(), poweredBy.Bounds().Dy()
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(-float64(frameWidth)/2, -float64(frameHeight)/2)
		op.GeoM.Translate(float64(screenWidth)/2, float64(screenHeight)/2)
		screen.DrawImage(poweredBy, op)
	case ModePlay:
		pressSpaceWidth := pressSpace.Bounds().Dx()
		spaceOp := &ebiten.DrawImageOptions{}
		spaceOp.GeoM.Translate(-float64(pressSpaceWidth)/2, 0)
		spaceOp.GeoM.Translate(float64(screenWidth)/2, 312)
		screen.DrawImage(pressSpace, spaceOp)

		dice1Width := dice1.Bounds().Dx()
		diceOp := &ebiten.DrawImageOptions{}
		diceOp.GeoM.Translate(-float64(dice1Width)/2, 0)
		diceOp.GeoM.Translate(float64(screenWidth)/2, 148)
		screen.DrawImage(dice1, diceOp)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
