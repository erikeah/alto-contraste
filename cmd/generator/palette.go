package main

import (
	"encoding/json"
	"fmt"
)

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
	Orange    Color `json:"orange"`
	AltOrange Color `json:"alt_orange"`
}

type Palette struct {
	Name          string
	Background    Color `json:"background"`     // Default
	Foreground    Color `json:"foreground"`     // Default
	AltBackground Color `json:"alt_background"` // Cursor
	AltForeground Color `json:"alt_foreground"` // Cursor
	*Base16
	*Extras
}

func (palette Palette) ToJSON() (string, error) {
	data, err := json.Marshal(palette)
	if err != nil {
		return "", fmt.Errorf("failed to marshal User to JSON: %w", err)
	}
	return string(data), nil
}
