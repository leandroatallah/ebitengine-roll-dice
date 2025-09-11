package main

import (
	"image/color"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
)

type Mode int

const (
	ModeInit Mode = iota
	ModePlay
	ModeRolling
)

const (
	frameRate      = 8
	rollCountLimit = 8
)

type Game struct {
	count     int
	mode      Mode
	diceValue int
	rollCount int
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
			// // TODO: Make smooth transition
			g.mode = ModeRolling
		}
	case ModeRolling:
		if g.count%frameRate == 0 {
			g.RollDice()
			g.rollCount++
		}

		if g.rollCount > rollCountLimit {
			g.mode = ModePlay
			g.rollCount = 0
			return nil
		}
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
	case ModePlay, ModeRolling:
		if g.diceValue == 0 {
			g.RollDice()
		}

		if g.mode == ModePlay {
			pressSpaceWidth := pressSpace.Bounds().Dx()
			spaceOp := &ebiten.DrawImageOptions{}
			spaceOp.GeoM.Translate(-float64(pressSpaceWidth)/2, 0)
			spaceOp.GeoM.Translate(float64(screenWidth)/2, 312)
			screen.DrawImage(pressSpace, spaceOp)
		}

		diceMap := map[int]*ebiten.Image{
			1: dice1,
			2: dice2,
			3: dice3,
			4: dice4,
			5: dice5,
			6: dice6,
		}
		dice := diceMap[g.diceValue]

		diceWidth := dice.Bounds().Dx()
		diceOp := &ebiten.DrawImageOptions{}
		diceOp.GeoM.Translate(-float64(diceWidth)/2, 0)
		diceOp.GeoM.Translate(float64(screenWidth)/2, 148)
		screen.DrawImage(dice, diceOp)
	}
}

func (g *Game) RollDice() {
	g.diceValue = rand.IntN(6) + 1
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
