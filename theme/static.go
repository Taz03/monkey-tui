package theme

import (
	"fmt"
	"io"
	"net/http"
	"regexp"

	tea "github.com/charmbracelet/bubbletea"
)

// implements Theme
type StaticPreset struct {
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

func GetStaticTheme(themeName string) Theme {
	response, err := http.Get("https://monkeytype.com/themes/" + themeName + ".css")
	if err != nil || response.StatusCode != 200 {
		panic("Failed to fetch theme")
	}

	defer response.Body.Close()

    bodySlice, _ := io.ReadAll(response.Body)
	body := string(bodySlice)

    val := StaticPreset{
		backgroundColor:         readColor(body, "bg-color"),
		mainColor:               readColor(body, "main-color"),
		caretColor:              readColor(body, "caret-color"),
		subColor:                readColor(body, "sub-color"),
		subAltColor:             readColor(body, "sub-alt-color"),
		textColor:               readColor(body, "text-color"),
		errorColor:              readColor(body, "error-color"),
		errorExtraColor:         readColor(body, "error-extra-color"),
		colorfulErrorColor:      readColor(body, "colorful-error-color"),
		colorfulErrorExtraColor: readColor(body, "colorful-error-extra-color"),
	}

    return val
}

func readColor(body, name string) string {
    colorRegex := regexp.MustCompile(fmt.Sprintf("--%s: (#.{6})", name))
	return colorRegex.FindStringSubmatch(body)[1]
}

func (this StaticPreset) BackgroundColor() string {
	return this.backgroundColor
}

func (this StaticPreset) MainColor() string {
	return this.mainColor
}

func (this StaticPreset) CaretColor() string {
	return this.caretColor
}

func (this StaticPreset) SubColor() string {
	return this.subColor
}

func (this StaticPreset) SubAltColor() string {
	return this.subAltColor
}

func (this StaticPreset) TextColor() string {
	return this.textColor
}

func (this StaticPreset) ErrorColor() string {
	return this.errorColor
}

func (this StaticPreset) ErrorExtraColor() string {
	return this.errorExtraColor
}

func (this StaticPreset) ColorfulErrorColor() string {
	return this.colorfulErrorColor
}

func (this StaticPreset) ColorfulErrorExtraColor() string {
	return this.colorfulErrorExtraColor
}

func (_ StaticPreset) Update(_ *tea.Program) {}
