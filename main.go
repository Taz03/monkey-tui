package main

import (
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/taz03/monkeytui/config"
	"github.com/taz03/monkeytui/test"
)

type model struct {
	Test *test.Test
}

var (
	width  int
	height int
)

func main() {
	app := tea.NewProgram(model{
		Test: &test.Test{
			Words:  []string{"The", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"},
			Config: config.UserConfig,
		},
	}, tea.WithAltScreen())
	app.Run()
}

func (m model) Init() tea.Cmd {
	m.Test.Init()
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		width, height = msg.Width, msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		default:
			m.Test.Update(msg)
			return m, nil
		}
	}

	return m, nil
}

func (m model) View() string {
	return lipgloss.Place(width, height, lipgloss.Center, lipgloss.Center, m.Test.View(), lipgloss.WithWhitespaceBackground(config.UserConfig.BackgroundColor()))
}
