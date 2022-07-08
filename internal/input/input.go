package input

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	string string
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyRunes, tea.KeySpace:
			m.string += msg.String()
		case tea.KeyBackspace:
			if len(m.string) > 0 {
				m.string = m.string[:len(m.string)-1]
			}
		case tea.KeyEnter:
			m.string = ""
		}
	}
	return m, nil
}

func (m Model) View() string {
	return m.string
}

func (m Model) GetText() string {
	return m.string
}

func Init() Model {
	return Model{
		string: "",
	}
}
