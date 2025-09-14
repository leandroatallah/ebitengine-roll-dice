package main

import "github.com/hajimehoshi/ebiten/v2"

const splashScreenTime = 100

type GameModeState interface {
	draw(g *Game, screen *ebiten.Image)
	update(g *Game) error
}

type GameModeInit struct{}

func (m *GameModeInit) draw(g *Game, screen *ebiten.Image) {
	poweredBy := g.Assets.UI.PoweredBy
	frameWidth, frameHeight := poweredBy.Bounds().Dx(), poweredBy.Bounds().Dy()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(frameWidth)/2, -float64(frameHeight)/2)
	op.GeoM.Translate(float64(screenWidth)/2, float64(screenHeight)/2)
	screen.DrawImage(poweredBy, op)
}

func (m *GameModeInit) update(g *Game) error {
	if g.count > splashScreenTime {
		// TODO: Make smooth transition
		g.Mode = &GameModePlay{}
	}

	return nil
}

type GameModePlay struct{}

func (m *GameModePlay) draw(g *Game, screen *ebiten.Image) {
	g.drawPressSpace(screen)
	g.drawDie(screen)
}

func (m *GameModePlay) update(g *Game) error {
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		// // TODO: Make smooth transition
		g.Mode = &GameModeRoll{}
	}

	return nil
}

type GameModeRoll struct{}

func (m *GameModeRoll) draw(g *Game, screen *ebiten.Image) {
	g.drawDie(screen)
}

func (m *GameModeRoll) update(g *Game) error {
	if g.count%frameRate == 0 {
		g.Dice.Roll()
		g.rollCount++
	}

	if g.rollCount > rollCountLimit {
		g.Mode = &GameModePlay{}
		g.rollCount = 0
	}

	return nil
}
