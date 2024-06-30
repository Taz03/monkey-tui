package theme

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// implements Theme
type RGBTheme struct {
	backgroundColor         string
	mainColor               string
	caretColor              string
	subColor                string
	subAltColor             string
	textColor               string
	errorColor              string
	errorExtraColor         string
	colorfulErrorColor      string
	colorfulErrorExtraColor string
}

var interpolatedColors []string

func (rgb *RGBTheme) Init() {
	rgb.backgroundColor = "#111"
	rgb.mainColor = "#eee"
	rgb.caretColor = "#eee"
	rgb.subColor = "#444"
	rgb.subAltColor = "#1a1a1a"
	rgb.textColor = "#eee"
	rgb.errorColor = "#eee"
	rgb.errorExtraColor = "#b3b3b3"
	rgb.colorfulErrorColor = "#eee"
	rgb.colorfulErrorExtraColor = "#b3b3b3"

	c1 := RGB{76, 174, 76}
	c2 := RGB{64, 158, 181}
	c3 := RGB{129, 52, 244}
	c4 := RGB{241, 14, 25}
	c5 := RGB{255, 197, 5}
	c6 := RGB{76, 174, 76}
	interpolatedColors = append(interpolatedColors, InterpolateColor(c1, c2, 25)...)
	interpolatedColors = append(interpolatedColors, InterpolateColor(c2, c3, 25)...)
	interpolatedColors = append(interpolatedColors, InterpolateColor(c3, c4, 25)...)
	interpolatedColors = append(interpolatedColors, InterpolateColor(c4, c5, 25)...)
	interpolatedColors = append(interpolatedColors, InterpolateColor(c5, c6, 25)...)
	interpolatedColors = append(interpolatedColors, InterpolateColor(c6, c5, 25)...)
	interpolatedColors = append(interpolatedColors, InterpolateColor(c5, c4, 25)...)
	interpolatedColors = append(interpolatedColors, InterpolateColor(c4, c3, 25)...)
	interpolatedColors = append(interpolatedColors, InterpolateColor(c3, c2, 25)...)
	interpolatedColors = append(interpolatedColors, InterpolateColor(c2, c1, 25)...)
}

func (rgb *RGBTheme) BackgroundColor() string {
	return rgb.backgroundColor
}

func (rgb *RGBTheme) MainColor() string {
	return rgb.mainColor
}

func (rgb *RGBTheme) CaretColor() string {
	return rgb.caretColor
}

func (rgb *RGBTheme) SubColor() string {
	return rgb.subColor
}

func (rgb *RGBTheme) SubAltColor() string {
	return rgb.subAltColor
}

func (rgb *RGBTheme) TextColor() string {
	return rgb.textColor
}

func (rgb *RGBTheme) ErrorColor() string {
	return rgb.errorColor
}

func (rgb *RGBTheme) ErrorExtraColor() string {
	return rgb.errorExtraColor
}

func (rgb *RGBTheme) ColorfulErrorColor() string {
	return rgb.colorfulErrorColor
}

func (rgb *RGBTheme) ColorfulErrorExtraColor() string {
	return rgb.colorfulErrorExtraColor
}

func (rgb *RGBTheme) Update(app *tea.Program) {
	for {
		for _, color := range interpolatedColors {
			rgb.mainColor = color
			rgb.textColor = color

			app.Send(color)
			time.Sleep(time.Millisecond * 25)
		}
	}
}
