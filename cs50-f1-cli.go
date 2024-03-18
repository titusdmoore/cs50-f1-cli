package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbletea"
)

type model struct {
	active_page int
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func initialModel() model {
	return model{
		active_page: 0,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			m.active_page++

		case "down", "j":
			m.active_page--

		}
	}

	return m, nil
}

func (m model) View() string {
	// The header
	s := "Welcome to the Formula 1 CLI\n\n"

	s += fmt.Sprintf("Active page: %d\n", m.active_page)

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}
