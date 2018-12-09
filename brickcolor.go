package brickcolor

import (
	"image/color"
	"math/rand"
)

//go:generate go run -tags=generator _generator/main.go

// BrickColor contains the color, name, number and hex value for a Roblox part color
type BrickColor struct {
	color.RGBA
	Name   string
	Number int
	Hex    string
}

// Random BrickColor (remember to seed the random number generator using rand.Seed)
func Random() BrickColor {
	return All[rand.Intn(len(All))]
}

// Number returns a BrickColor from its numerical index.
func Number(n int) BrickColor {
	return byNumber[n]
}
