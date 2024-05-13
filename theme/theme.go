package theme

import (
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

func GetTheme(name string) Theme {
    switch strings.ToLower(name) {
    case "terra":
        return Terra{}
    }

    return nil
}
