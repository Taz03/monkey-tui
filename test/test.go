package test

import (
	"github.com/charmbracelet/bubbles/cursor"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/taz03/monkeytui/config"
)

var (
    style lipgloss.Style
    caret cursor.Model
    space string
)

type Test struct {
    Words  []string
    Config config.Config

    Width int

    typedWords []string
    pos        [2]int
}

func (this *Test) Init() tea.Cmd {
    this.typedWords = []string{""}

    style = lipgloss.NewStyle().Background(this.Config.BackgroundColor()).Bold(true)
    caret = this.Config.Cursor()
    space = lipgloss.NewStyle().Background(this.Config.BackgroundColor()).Render(" ")

    return nil
}

func (this *Test) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case " ":
            this.typedWords = append(this.typedWords, "")
            this.pos[0]++
            this.pos[1] = 0

        case "backspace":
            if this.pos[1]--; this.pos[1] < 0 {
                if this.pos[0] > 0 {
                    this.pos[0]--
                    this.typedWords = this.typedWords[:len(this.typedWords) - 1]
                }
                this.pos[1] = len(this.typedWords[this.pos[0]])
            } else {
                this.typedWords[this.pos[0]] = this.typedWords[this.pos[0]][:this.pos[1]]
            }

        default:
            this.typedWords[len(this.typedWords) - 1] += msg.String()
            this.pos[1]++
        }
    }

    return this, nil
}
