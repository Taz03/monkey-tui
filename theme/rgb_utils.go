package theme

import "fmt"

type RGB struct {
	R, G, B int
}

func rgbToHex(r, g, b int) string {
	return fmt.Sprintf("#%02X%02X%02X", r, g, b)
}

func InterpolateColor(color1, color2 RGB, steps int) []string {
	var interpolatedColors []string

	deltaR := float64(color2.R-color1.R) / float64(steps)
	deltaG := float64(color2.G-color1.G) / float64(steps)
	deltaB := float64(color2.B-color1.B) / float64(steps)

	for i := 0; i < steps; i++ {
		r := color1.R + int(deltaR*float64(i))
		g := color1.G + int(deltaG*float64(i))
		b := color1.B + int(deltaB*float64(i))
		interpolatedColors = append(interpolatedColors, rgbToHex(r, g, b))
	}

	return interpolatedColors
}
