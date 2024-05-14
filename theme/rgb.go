package theme

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type RGB struct {
	R, G, B int
}

func rgbToHex(r, g, b int) string {
	return fmt.Sprintf("#%02X%02X%02X", r, g, b)
}

func interpolateColor(color1, color2 RGB, steps int) []string {
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

// implements Theme
type RGBTheme struct {
    backgroundColor string
    mainColor       string
    caretColor      string
    subColor        string
    subAltColor     string
    textColor       string
    errorColor      string
    errorExtraColor string
    colorfulErrorColor string
    colorfulErrorExtraColor string
}

var interpolatedColors []string

func init() {
    startColor := RGB{255, 0, 0}
    endColor := RGB{0, 0, 255}
    interpolatedColors = interpolateColor(startColor, endColor, 100)
}

func (this *RGBTheme) Init() {
    this.backgroundColor = "#111"
    this.mainColor = "#eee"
    this.caretColor = "#eee"
    this.subColor = "#444"
    this.subAltColor = "#1a1a1a"
    this.textColor = "#eee"
    this.errorColor = "#eee"
    this.errorExtraColor = "#b3b3b3"
    this.colorfulErrorColor = "#eee"
    this.colorfulErrorExtraColor = "#b3b3b3"
}

func (this *RGBTheme) BackgroundColor() string {
    return this.backgroundColor
}

func (this *RGBTheme) MainColor() string {
    return this.mainColor
}

func (this *RGBTheme) CaretColor() string {
    return this.caretColor
}

func (this *RGBTheme) SubColor() string {
    return this.subColor
}

func (this *RGBTheme) SubAltColor() string {
    return this.subAltColor
}

func (this *RGBTheme) TextColor() string {
    return this.textColor
}

func (this *RGBTheme) ErrorColor() string {
    return this.errorColor
}

func (this *RGBTheme) ErrorExtraColor() string {
    return this.errorExtraColor
}

func (this *RGBTheme) ColorfulErrorColor() string {
    return this.colorfulErrorColor
}

func (this *RGBTheme) ColorfulErrorExtraColor() string {
    return this.colorfulErrorExtraColor
}

func (this *RGBTheme) Update(app *tea.Program) {
    for {
        for _, color := range interpolatedColors {
            this.mainColor = color
            this.textColor = color

            app.Send(color)
            time.Sleep(time.Millisecond * 50)
        }
    }
}
