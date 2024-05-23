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

func (this *Model) BackgroundColor() lipgloss.TerminalColor {
	if strings.Trim(this.CustomBackground, " ") != "" {
		return lipgloss.NoColor{}
	} else {
		return lipgloss.Color(this.MonkeyTheme.BackgroundColor())
	}
}

func (this *Model) Cursor() cursor.Model {
	return cursor.Model{
		Style: lipgloss.NewStyle().Foreground(lipgloss.Color(this.MonkeyTheme.CaretColor())).Background(lipgloss.Color(this.MonkeyTheme.SubColor())),
	}
}

func (this *Model) StyleWrongWordUnderline(style lipgloss.Style) lipgloss.Style {
	if this.BlindMode {
		return style.UnsetUnderline()
	}

    return style.Underline(true).UnderlineSpaces(false)
}

func (this *Model) StyleUntyped(style lipgloss.Style, word string) lipgloss.Style {
	style = style.SetString(word)

	foregroundColor := this.MonkeyTheme.SubColor()
	if this.FlipTestColors && this.ColorfulMode {
		foregroundColor = this.MonkeyTheme.MainColor()
	} else if this.FlipTestColors {
		foregroundColor = this.MonkeyTheme.TextColor()
	}
	style = style.Foreground(lipgloss.Color(foregroundColor))

	return style
}

func (this *Model) StyleCorrect(style lipgloss.Style, word string) lipgloss.Style {
	style = style.SetString(word)

	foregroundColor := this.MonkeyTheme.TextColor()
	if this.FlipTestColors {
		foregroundColor = this.MonkeyTheme.SubColor()
	} else if this.ColorfulMode {
		foregroundColor = this.MonkeyTheme.MainColor()
	}
	style = style.Foreground(lipgloss.Color(foregroundColor))

	return style
}

func (this *Model) StyleError(style lipgloss.Style, typed, word string) lipgloss.Style {
	if this.BlindMode {
		style = this.StyleCorrect(style, word)
		return style
	}

	foregroundColor := this.MonkeyTheme.ErrorColor()
	if this.ColorfulMode {
		foregroundColor = this.MonkeyTheme.ColorfulErrorColor()
	}
	style.Foreground(lipgloss.Color(foregroundColor))

	if this.IndicateTypos == "replace" {
		style = style.SetString(typed)
		return style
	} else {
		style = style.SetString(word)
		return style
	}
}

func (this *Model) StyleErrorExtra(style lipgloss.Style, char string) lipgloss.Style {
	if this.HideExtraLetters || this.BlindMode {
		style = style.SetString("")
		return style
	}

	style = style.SetString(char)

	foregroundColor := this.MonkeyTheme.ErrorExtraColor()
	if this.ColorfulMode {
		foregroundColor = this.MonkeyTheme.ColorfulErrorExtraColor()
	}
	style.Foreground(lipgloss.Color(foregroundColor))

	return style
}
