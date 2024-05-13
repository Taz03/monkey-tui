package theme

import (
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

var (
    staticPresets = []string{
        "terra", "arch",
    }
)

func GetTheme(name string) Theme {
    if slices.Contains(staticPresets, strings.ToLower(name)) {
        return GetStaticTheme(name)
    }

    return nil
}
