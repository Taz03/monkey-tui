package theme

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

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
            time.Sleep(time.Millisecond * 25)
        }
    }
}
