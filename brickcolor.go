/*
Package brickcolor includes all of the Roblox BrickColor Codes as
listed on https://developer.roblox.com/articles/BrickColor-Codes

Installation

	go get -u github.com/peterhellberg/brickcolor

Usage

	package main

	import (
		"fmt"
		"math/rand"

		"github.com/peterhellberg/brickcolor"
	)

	func main() {
		rand.Seed(123)

		println(brickcolor.Random())
		println(brickcolor.Number(1032))
	}

	func println(bc brickcolor.BrickColor) {
		fmt.Println(bc.R, bc.G, bc.B, bc.Name)
	}
*/
package brickcolor

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
