package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"os"
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

type Base16 struct {
	Black      Color `json:"black"`
	Blue       Color `json:"blue"`
	Cyan       Color `json:"cyan"`
	Green      Color `json:"green"`
	Magenta    Color `json:"magenta"`
	Red        Color `json:"red"`
	White      Color `json:"white"`
	Yellow     Color `json:"yellow"`
	AltBlack   Color `json:"alt_black"`
	AltBlue    Color `json:"alt_blue"`
	AltCyan    Color `json:"alt_cyan"`
	AltGreen   Color `json:"alt_green"`
	AltMagenta Color `json:"alt_magenta"`
	AltRed     Color `json:"alt_red"`
	AltWhite   Color `json:"alt_white"`
	AltYellow  Color `json:"alt_yellow"`
}

type Extras struct {
}

type Palette struct {
	Background    Color `json:"background"`
	Foreground    Color `json:"foreground"`
	AltBackground Color `json:"alt_background"` // Cursor
	AltForeground Color `json:"alt_foreground"` // Cursor
	*Base16
	*Extras
}

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	theme, err := os.Open(os.Args[1])
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	decoder := json.NewDecoder(theme) // TODO: Read from file provided from param
	palette := &Palette{}
	if err := decoder.Decode(palette); err != nil {
		println(err.Error())
		os.Exit(1)
	}
	funcMap := template.FuncMap{
		"toHexAlpha":   toHexAlpha,
		"toHexNoAlpha": toHexNoAlpha,
	}
	templateBytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		println(err.Error())
		os.Exit(2)

	}
	templateContent := string(templateBytes)
	tmpl, err := template.New("").Funcs(funcMap).Parse(templateContent)
	if err != nil {
		println(err.Error())
		os.Exit(2)
	}
	err = tmpl.Execute(os.Stdout, palette)
	if err != nil {
		println(err.Error())
		os.Exit(2)
	}
}
