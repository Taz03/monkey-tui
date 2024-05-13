package theme

import (
	"encoding/json"
	"os"
	"slices"
	"strings"
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

    Update()
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

    return nil
}
