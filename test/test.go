package test

import (
	"time"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/progress"
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
	config *config.Model

	ProgressBar progress.Model
	Statistics  string

	Width int

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

	m := &Model{
		config:     config,
		words:      GenerateWords(config, wordsController),
		addWord:    wordsController,
		typedWords: []string{""},
	}

	if config.TimerStyle == "bar" {
		m.ProgressBar = progress.New(
			progress.WithSolidFill(config.LiveStatsColor()),
			progress.WithoutPercentage(),
		)
		m.ProgressBar.Full = '▀'
		m.ProgressBar.Empty = ' '

		m.ProgressBar.EmptyColor = m.config.MonkeyTheme.BackgroundColor()

		if config.Mode == "time" {
			m.ProgressBar.SetPercent(1)
		}
	}

	return m
}

func (m *Model) Init() tea.Cmd {
	return tickCmd()
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Millisecond*50, func(t time.Time) tea.Msg {
		return t
	})
}

func (m *Model) calculateTestWidth(width int) int {
	if m.config.MaxLineWidth == 0 {
		return width - 10
	}

	if m.config.MaxLineWidth > width {
		return width
	}

	return m.config.MaxLineWidth
}

func (m *Model) calculateProgressPercentage() float64 {
	if !m.started {
		return 0
	}

	if m.config.Mode == "time" {
		return (float64(m.config.Time) - float64(time.Now().Sub(m.startTime).Seconds())) / float64(m.config.Time)
	}

	return float64(len(m.typedWords)) / float64(len(*m.words))
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case time.Time:
		return m, tea.Batch(tickCmd(), m.ProgressBar.SetPercent(m.calculateProgressPercentage()))

	case tea.KeyMsg:
		switch msg.String() {
		case tea.KeySpace.String():
			m.typedWords = append(m.typedWords, "")
			if !(m.config.Mode == "words" && m.config.Words != 0) {
				m.addWord <- true
			}
			m.pos[0]++
			m.pos[1] = 0

		case tea.KeyBackspace.String():
			if m.pos[1]--; m.pos[1] < 0 {
				if m.pos[0] > 0 {
					m.pos[0]--
					m.typedWords = m.typedWords[:len(m.typedWords)-1]
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

			m.typedWords[len(m.typedWords)-1] += msg.String()
			m.pos[1]++
		}

	case progress.FrameMsg:
		progressModel, cmd := m.ProgressBar.Update(msg)
		m.ProgressBar = progressModel.(progress.Model)
		return m, cmd
	}

	return m, nil
}
