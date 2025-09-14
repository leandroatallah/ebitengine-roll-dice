package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	frameRate      = 6
	rollCountLimit = 10
	screenWidth    = 640
	screenHeight   = 480
)

type Game struct {
	count     int
	rollCount int
	Mode      GameModeState
	Assets    *Assets
	Dice      *Dice
}

func NewGame() (*Game, error) {
	assets, err := NewAssets()
	if err != nil {
		return nil, err
	}

	return &Game{
		Assets: assets,
		Mode:   &GameModeInit{},
		Dice:   NewDice(),
	}, nil
}

func (g *Game) Update() error {
	g.count++

	return g.Mode.update(g)
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x21, 0x23, 0x43, 0xff})

	g.Mode.draw(g, screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
