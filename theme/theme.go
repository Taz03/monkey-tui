package theme

import (
	"encoding/json"
	"os"
	"slices"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Theme interface {
    BackgroundColor() string
    MainColor() string
    CaretColor() string
    SubColor() string
    SubAltColor() string
    TextColor() string
    ErrorColor() string
    ErrorExtraColor() string
    ColorfulErrorColor() string
    ColorfulErrorExtraColor() string

    Update(*tea.Program)
}

var staticPresets []string

func init() {
    var model struct {
        Names []string `json:"names"`
    }

    fileContent, _ := os.ReadFile("static_presets.json")
    json.Unmarshal(fileContent, &model)

    staticPresets = model.Names
}

func GetTheme(name string) Theme {
    if slices.Contains(staticPresets, strings.ToLower(name)) {
        return GetStaticTheme(name)
    }

    switch strings.ToLower(name) {
    case "rgb":
        rgb := &RGBTheme{}
        rgb.Init()
        return rgb
    }

    return nil
}
