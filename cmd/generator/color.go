package main

import (
	"fmt"
	"math"
)

// [ r g b a ]
type Color [4]float64

func (color Color) ToHex(channels int) string {
	var hexString string
	for i, value := range color {
		channel := 1 << i
		if channels&channel != channel {
			continue
		}
		if value < 0.0 {
			value = 0.0
		} else if value > 1.0 {
			value = 1.0
		}
		scaled := math.Round(value * 255.0)
		component := uint8(scaled)
		hexString += fmt.Sprintf("%02X", component)
	}
	return hexString
}

func (color Color) ToHex3() string {
	return color.ToHex(7) // 0111
}

func (color Color) ToHex4() string {
	return color.ToHex(15) // 1111
}
