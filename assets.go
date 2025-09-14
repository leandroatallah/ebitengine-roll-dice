package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Assets struct {
	Dice *DiceAssets
	UI   *UIAssets
}

type DiceAssets struct {
	images map[int]*ebiten.Image
}

func loadDiceAssets() (*DiceAssets, error) {
	assets := &DiceAssets{
		images: make(map[int]*ebiten.Image),
	}
	for i := 1; i <= 6; i++ {
		d, _, err := ebitenutil.NewImageFromFile(fmt.Sprintf("assets/dice-%d.png", i))
		if err != nil {
			return nil, err
		}
		assets.images[i] = d
	}
	return assets, nil
}

type UIAssets struct {
	PoweredBy  *ebiten.Image
	PressSpace *ebiten.Image
}

func loadUIAssets() (*UIAssets, error) {
	ui := &UIAssets{}

	var err error
	ui.PoweredBy, _, err = ebitenutil.NewImageFromFile("assets/powered-by.png")
	if err != nil {
		return nil, err
	}
	ui.PressSpace, _, err = ebitenutil.NewImageFromFile("assets/press-space.png")
	if err != nil {
		return nil, err
	}

	return ui, nil
}

func NewAssets() (*Assets, error) {
	diceAssets, err := loadDiceAssets()
	if err != nil {
		return nil, err
	}

	uiAssets, err := loadUIAssets()
	if err != nil {
		return nil, err
	}

	return &Assets{
		Dice: diceAssets,
		UI:   uiAssets,
	}, nil
}
