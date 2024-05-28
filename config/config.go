package config

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/lipgloss"
	"github.com/taz03/monkeytui/theme"
)

type Model struct {
	Theme                   string   `json:"theme"`
	ThemeDark               string   `json:"themeDark"`
	ThemeLight              string   `json:"themeLight"`
	AutoSwitchTheme         bool     `json:"autoSwitchTheme"`
	CustomTheme             bool     `json:"customTheme"`
	CustomThemeColors       []string `json:"customThemeColors"`
	FavThemes               []string `json:"favThemes"`
	ShowKeyTips             bool     `json:"showKeyTips"`
	QuickRestart            string   `json:"quickRestart"`
	Punctuation             bool     `json:"punctuation"`
	Numbers                 bool     `json:"numbers"`
	Words                   int      `json:"words"`
	Time                    int      `json:"time"`
	Mode                    string   `json:"mode"`
	QuoteLength             []int    `json:"quoteLength"`
	Language                string   `json:"language"`
	FreedomMode             bool     `json:"freedomMode"`
	Difficulty              string   `json:"difficulty"`
	BlindMode               bool     `json:"blindMode"`
	QuickEnd                bool     `json:"quickEnd"`
	CaretStyle              string   `json:"caretStyle"`
	PaceCaretStyle          string   `json:"paceCaretStyle"`
	FlipTestColors          bool     `json:"flipTestColors"`
	Layout                  string   `json:"layout"`
	Funbox                  string   `json:"funbox"`
	ConfidenceMode          string   `json:"confidenceMode"`
	IndicateTypos           string   `json:"indicateTypos"`
	TimerStyle              string   `json:"timerStyle"`
	LiveSpeedStyle          string   `json:"liveSpeedStyle"`
	LiveAccStyle            string   `json:"liveAccStyle"`
	LiveBurstStyle          string   `json:"liveBurstStyle"`
	ColorfulMode            bool     `json:"colorfulMode"`
	RandomTheme             string   `json:"randomTheme"`
	TimerColor              string   `json:"timerColor"`
	StopOnError             string   `json:"stopOnError"`
	ShowAllLines            bool     `json:"showAllLines"`
	KeymapMode              string   `json:"keymapMode"`
	KeymapStyle             string   `json:"keymapStyle"`
	KeymapLegendStyle       string   `json:"keymapLegendStyle"`
	KeymapLayout            string   `json:"keymapLayout"`
	KeymapShowTopRow        string   `json:"keymapShowTopRow"`
	AlwaysShowDecimalPlaces bool     `json:"alwaysShowDecimalPlaces"`
	AlwaysShowWordsHistory  bool     `json:"alwaysShowWordsHistory"`
	SingleListCommandLine   string   `json:"singleListCommandLine"`
	CapsLockWarning         bool     `json:"capsLockWarning"`
	PlaySoundOnError        string   `json:"playSoundOnError"`
	PlaySoundOnClick        string   `json:"playSoundOnClick"`
	SoundVolume             string   `json:"soundVolume"`
	StartGraphsAtZero       bool     `json:"startGraphsAtZero"`
	ShowOutOfFocusWarning   bool     `json:"showOutOfFocusWarning"`
	PaceCaret               string   `json:"paceCaret"`
	PaceCaretCustomSpeed    int      `json:"paceCaretCustomSpeed"`
	RepeatedPace            bool     `json:"repeatedPace"`
	AccountChart            []string `json:"accountChart"`
	MinWpm                  string   `json:"minWpm"`
	MinWpmCustomSpeed       int      `json:"minWpmCustomSpeed"`
	HighlightMode           string   `json:"highlightMode"`
	TypingSpeedUnit         string   `json:"typingSpeedUnit"`
	HideExtraLetters        bool     `json:"hideExtraLetters"`
	StrictSpace             bool     `json:"strictSpace"`
	MinAcc                  string   `json:"minAcc"`
	MinAccCustom            int      `json:"minAccCustom"`
	RepeatQuotes            string   `json:"repeatQuotes"`
	OppositeShiftMode       string   `json:"oppositeShiftMode"`
	CustomBackground        string   `json:"customBackground"`
	CustomLayoutFluid       string   `json:"customLayoutFluid"`
	MinBurst                string   `json:"minBurst"`
	MinBurstCustomSpeed     int      `json:"minBurstCustomSpeed"`
	BurstHeatmap            bool     `json:"burstHeatmap"`
	BritishEnglish          bool     `json:"britishEnglish"`
	LazyMode                bool     `json:"lazyMode"`
	ShowAverage             string   `json:"showAverage"`
	TapeMode                string   `json:"tapeMode"`
	MaxLineWidth            int      `json:"maxLineWidth"`

	MonkeyTheme theme.Theme
}

func New(path string) (config *Model) {
	fileContent, _ := os.ReadFile(path)
	json.Unmarshal(fileContent, &config)

	config.MonkeyTheme = theme.GetTheme(config.Theme)
	return
}

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
