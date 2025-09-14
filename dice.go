package main

import "math/rand/v2"

type Dice struct {
	Value int
}

func NewDice() *Dice {
	d := &Dice{}
	d.Roll()
	return d
}

func (d *Dice) Roll() {
	d.Value = rand.IntN(6) + 1
}
