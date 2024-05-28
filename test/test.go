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

    config *config.Model

    words   *[]string
    addWord chan bool

    typedWords []string
    pos        [2]int

    started   bool
    startTime time.Time
}

func New(config *config.Model) *Model {
    style = lipgloss.NewStyle().Background(config.BackgroundColor())
    caret = config.Cursor()
    space = lipgloss.NewStyle().Background(config.BackgroundColor()).Render(" ")

    wordsController := make(chan bool)

    return &Model{
        config:     config,
        words:      GenerateWords(config, wordsController),
        addWord:    wordsController,
        typedWords: []string{""},
    }
}

func (m *Model) Init() tea.Cmd {
    return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case tea.KeySpace.String():
            m.typedWords = append(m.typedWords, "")
            m.addWord <- true
            m.pos[0]++
            m.pos[1] = 0

        case tea.KeyBackspace.String():
            if m.pos[1]--; m.pos[1] < 0 {
                if m.pos[0] > 0 {
                    m.pos[0]--
                    m.typedWords = m.typedWords[:len(m.typedWords) - 1]
                }
                m.pos[1] = len(m.typedWords[m.pos[0]])
            } else {
                m.typedWords[m.pos[0]] = m.typedWords[m.pos[0]][:m.pos[1]]
            }

        default:
            if !m.started {
                m.startTime = time.Now()
                m.started = true
            }

            m.typedWords[len(m.typedWords) - 1] += msg.String()
            m.pos[1]++
        }
    }

    return m, nil
}
