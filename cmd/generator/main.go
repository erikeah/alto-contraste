package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"text/template"
)

func toHex(color Color, alpha bool) string {
	var hexString string
	for i, channel := range color {
		if !alpha && i == 3 {
			break
		}
		// 1. Clamp the float to the [0.0, 1.0] range just in case.
		if channel < 0.0 {
			channel = 0.0
		} else if channel > 1.0 {
			channel = 1.0
		}
		scaled := math.Round(channel * 255.0)
		component := uint8(scaled)
		hexString += fmt.Sprintf("%02X", component)
	}
	return hexString
}

func toHexAlpha(color Color) string {
	return toHex(color, true)
}

func toHexNoAlpha(color Color) string {
	return toHex(color, false)
}

// [ r g b a ]
type Color [4]float64

type Palette struct {
	Background    Color `json:"background"`
	Foreground    Color `json:"foreground"`
	Black         Color `json:"black"`
	Blue          Color `json:"blue"`
	Cyan          Color `json:"cyan"`
	Green         Color `json:"green"`
	Magenta       Color `json:"magenta"`
	Red           Color `json:"red"`
	White         Color `json:"white"`
	Yellow        Color `json:"yellow"`
	AltBackground Color `json:"alt_background"`
	AltBlack      Color `json:"alt_black"`
	AltBlue       Color `json:"alt_blue"`
	AltCyan       Color `json:"alt_cyan"`
	AltForeground Color `json:"alt_foreground"`
	AltGreen      Color `json:"alt_green"`
	AltMagenta    Color `json:"alt_magenta"`
	AltRed        Color `json:"alt_red"`
	AltWhite      Color `json:"alt_white"`
	AltYellow     Color `json:"alt_yellow"`
}

func main() {
	decoder := json.NewDecoder(os.Stdin)
	palette := &Palette{}
	if err := decoder.Decode(palette); err != nil {
		println(err.Error())
		os.Exit(1)
	}
	funcMap := template.FuncMap{
		"toHexAlpha":   toHexAlpha,
		"toHexNoAlpha": toHexNoAlpha,
	}
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	templatePath := os.Args[1]
	tmplSet, err := template.New("").Funcs(funcMap).ParseFiles(templatePath)
	if err != nil {
		println(err.Error())
		os.Exit(2)
	}
	templateName := filepath.Base(templatePath)
	specificTmpl := tmplSet.Lookup(templateName)
	err = specificTmpl.Execute(os.Stdout, palette)
	if err != nil {
		println(err.Error())
		os.Exit(2)
	}
}
