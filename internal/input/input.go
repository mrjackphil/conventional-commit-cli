package input

import (
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	string string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyRunes, tea.KeySpace:
			m.string += msg.String()
		}
	}
	return m, nil
}

func (m model) View() string {
	return m.string
}

func Init() tea.Model {
	return model{
		string: "Some input field",
	}
}
