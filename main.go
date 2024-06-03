package main

import (
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/taz03/monkeytui/config"
	"github.com/taz03/monkeytui/test"
)

type model struct {
	Test   *test.Model
    Config *config.Model
}

var width, height int

func main() {
    userConfig := config.New("config.json")

	app := tea.NewProgram(model{
		Test: test.New(userConfig),
        Config: userConfig,
	}, tea.WithAltScreen())
    go userConfig.MonkeyTheme.Update(app)

	app.Run()
}

func (m model) calculateTestWidth() int {
    if m.Config.MaxLineWidth == 0 {
        return width - 10
    }
    
    if m.Config.MaxLineWidth > width {
        return width
    }

    return m.Config.MaxLineWidth
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		width, height = msg.Width, msg.Height

    case progress.FrameMsg:
        _, cmd := m.Test.Update(msg)
        return m, cmd

	case tea.KeyMsg:
        if msg.String() == m.Config.RestartKey() {
            m.Test = test.New(m.Config)
            m.Test.Width = m.calculateTestWidth()
            break
        }

		switch msg.String() {
		case tea.KeyCtrlC.String():
			return m, tea.Quit
		default:
            _, cmd := m.Test.Update(msg)
            return m, cmd
		}
	}

	return m, nil
}

func (m model) View() string {
    m.Test.Width = m.calculateTestWidth()
    m.Test.ProgressBar.Width = width

    return lipgloss.JoinVertical(
        lipgloss.Left,
        m.Test.ProgressBar.View(),
        lipgloss.Place(
            width,
            height - 1,
            lipgloss.Center,
            lipgloss.Center,
            m.Test.View(),
            lipgloss.WithWhitespaceBackground(m.Config.BackgroundColor()),
        ),
    )
}
