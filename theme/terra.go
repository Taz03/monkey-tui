package theme

// implements Theme
type Terra struct {}

func (_ Terra) BackgroundColor() string {
    return "#0c100e"
}

func (_ Terra) MainColor() string {
    return "#89c559"
}

func (_ Terra) CaretColor() string {
    return "#89c559"
}

func (_ Terra) SubColor() string {
    return "#436029"
}

func (_ Terra) SubAltColor() string {
    return "#0f1d18"
}

func (_ Terra) TextColor() string {
    return "#f0edd1"
}

func (_ Terra) ErrorColor() string {
    return "#d3ca78"
}

func (_ Terra) ErrorExtraColor() string {
    return "#89844d"
}

func (_ Terra) ColorfulErrorColor() string {
    return "#d3ca78"
}

func (_ Terra) ColorfulErrorExtraColor() string {
    return "#89844d"
}

func (_ Terra) Update() {}
