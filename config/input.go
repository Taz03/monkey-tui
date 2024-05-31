package config

import tea "github.com/charmbracelet/bubbletea"

func (m *Model) RestartKey() string {
    if m.QuickRestart == "off" {
        return tea.KeyShiftTab.String()
    }

    return m.QuickRestart
}

