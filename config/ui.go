package config

import (
	"strings"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/lipgloss"
)

func (m *Model) BackgroundColor() lipgloss.TerminalColor {
	if strings.Trim(m.CustomBackground, " ") != "" {
		return lipgloss.NoColor{}
	} else {
		return lipgloss.Color(m.MonkeyTheme.BackgroundColor())
	}
}

func (m *Model) Cursor() cursor.Model {
	return cursor.Model{
		Style: lipgloss.NewStyle().Foreground(lipgloss.Color(m.MonkeyTheme.CaretColor())),
	}
}

func (m *Model) StyleWrongWordUnderline(style lipgloss.Style) lipgloss.Style {
	if m.BlindMode {
		return style.UnsetUnderline()
	}

    return style.Underline(true).UnderlineSpaces(false)
}

func (m *Model) StyleUntyped(style lipgloss.Style, word string) lipgloss.Style {
	style = style.SetString(word)

	foregroundColor := m.MonkeyTheme.SubColor()
	if m.FlipTestColors && m.ColorfulMode {
		foregroundColor = m.MonkeyTheme.MainColor()
	} else if m.FlipTestColors {
		foregroundColor = m.MonkeyTheme.TextColor()
	}
	return style.Foreground(lipgloss.Color(foregroundColor))
}

func (m *Model) StyleCorrect(style lipgloss.Style, word string) lipgloss.Style {
	foregroundColor := m.MonkeyTheme.TextColor()
	if m.FlipTestColors {
		foregroundColor = m.MonkeyTheme.SubColor()
	} else if m.ColorfulMode {
		foregroundColor = m.MonkeyTheme.MainColor()
	}
	return style.Foreground(lipgloss.Color(foregroundColor)).SetString(word)
}

func (m *Model) StyleError(style lipgloss.Style, typed, word string) lipgloss.Style {
	if m.BlindMode {
		style = m.StyleCorrect(style, word)
		return style
	}

	foregroundColor := m.MonkeyTheme.ErrorColor()
	if m.ColorfulMode {
		foregroundColor = m.MonkeyTheme.ColorfulErrorColor()
	}
	style = style.Foreground(lipgloss.Color(foregroundColor))

	if m.IndicateTypos == "replace" {
        return style.SetString(typed)
	} else {
        return style.SetString(word)
	}
}

func (m *Model) StyleErrorExtra(style lipgloss.Style, char string) lipgloss.Style {
	if m.HideExtraLetters || m.BlindMode {
		style = style.SetString("")
		return style
	}

	foregroundColor := m.MonkeyTheme.ErrorExtraColor()
	if m.ColorfulMode {
		foregroundColor = m.MonkeyTheme.ColorfulErrorExtraColor()
	}
	return style.Foreground(lipgloss.Color(foregroundColor)).SetString(char)
}

func (m *Model) LiveStatsColor() string {
    switch m.TimerColor {
    case "black":
        return "0"
    case "sub":
        return m.MonkeyTheme.SubColor()
    case "text":
        return m.MonkeyTheme.TextColor()
    case "main":
        return m.MonkeyTheme.MainColor()
    default:
        return ""
    }
}
