/*

	Package brickcolor includes all of the Roblox BrickColor Codes as
	listed on https://developer.roblox.com/articles/BrickColor-Codes

	Installation

			go get -u github.com/peterhellberg/brickcolor

*/
package brickcolor

//go:generate go run -tags=generator _generator/main.go

import (
	"image/color"
	"math/rand"
)

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
