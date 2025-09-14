package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) drawDie(screen *ebiten.Image) {
	img := g.Assets.Dice.images[g.Dice.Value]
	yPos := 148.0
	drawCentered(screen, img, yPos)
}

func (g *Game) drawPressSpace(screen *ebiten.Image) {
	pressSpace := g.Assets.UI.PressSpace
	yPos := 312.0
	drawCentered(screen, pressSpace, yPos)
}

func drawCentered(screen, img *ebiten.Image, yPos float64) {
	width := img.Bounds().Dx()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(width)/2, 0)
	op.GeoM.Translate(float64(screenWidth)/2, yPos)
	screen.DrawImage(img, op)
}
