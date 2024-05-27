package test

import (
	"time"

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

type Model struct {
    Width int

    config     *config.Model
    words      []string
    typedWords []string
    pos        [2]int
    started    bool
    startTime  time.Time
}

func New(config *config.Model) *Model {
    style = lipgloss.NewStyle().Background(config.BackgroundColor())
    caret = config.Cursor()
    space = lipgloss.NewStyle().Background(config.BackgroundColor()).Render(" ")

    return &Model{
        config:     config,
        words:      GenerateWords(config),
        typedWords: []string{""},
    }
}

func (this *Model) Init() tea.Cmd {
    return nil
}

func (this *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case tea.KeySpace.String():
            this.typedWords = append(this.typedWords, "")
            this.pos[0]++
            this.pos[1] = 0

        case tea.KeyBackspace.String():
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
            if !this.started {
                this.startTime = time.Now()
                this.started = true
            }

            this.typedWords[len(this.typedWords) - 1] += msg.String()
            this.pos[1]++
        }
    }

    return this, nil
}
