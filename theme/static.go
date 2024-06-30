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

	return &StaticPreset{
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
}

func readColor(body, name string) string {
	colorRegex := regexp.MustCompile(fmt.Sprintf("--%s: (#.{6})", name))
	return colorRegex.FindStringSubmatch(body)[1]
}

func (static *StaticPreset) BackgroundColor() string {
	return static.backgroundColor
}

func (static *StaticPreset) MainColor() string {
	return static.mainColor
}

func (static *StaticPreset) CaretColor() string {
	return static.caretColor
}

func (static *StaticPreset) SubColor() string {
	return static.subColor
}

func (static *StaticPreset) SubAltColor() string {
	return static.subAltColor
}

func (static *StaticPreset) TextColor() string {
	return static.textColor
}

func (static *StaticPreset) ErrorColor() string {
	return static.errorColor
}

func (static *StaticPreset) ErrorExtraColor() string {
	return static.errorExtraColor
}

func (static *StaticPreset) ColorfulErrorColor() string {
	return static.colorfulErrorColor
}

func (static *StaticPreset) ColorfulErrorExtraColor() string {
	return static.colorfulErrorExtraColor
}

func (_ *StaticPreset) Init()                 {}
func (_ *StaticPreset) Update(_ *tea.Program) {}
