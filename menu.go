package main

import (
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q", "Q":
			os.Exit(0)

		case "enter":
			if !gameOver {
				m.choice = gameChoices[m.cursor]
			} else {
				m.choice = menuChoices[m.cursor]
			}
			return m, tea.Quit

		case "j", "J":
			m.cursor++
			if m.cursor >= len(gameChoices) && !gameOver {
				m.cursor = 0
			} else if m.cursor >= len(gameChoices) && gameOver {
				m.cursor = 0
			}

		case "k", "K":
			m.cursor--
			if m.cursor < 0 && !gameOver {
				m.cursor = len(gameChoices) - 1
			} else if m.cursor < 0 && gameOver {
				m.cursor = len(menuChoices) - 1
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	s := strings.Builder{}

	//if game is finished then show another screen
	if gameOver {
		s.WriteString("Please choose what to do next:\n")
		for i := 0; i < len(menuChoices); i++ {
			if m.cursor == i {
				s.WriteString("(•) ")
			} else {
				s.WriteString("( ) ")
			}
			s.WriteString(menuChoices[i])
			s.WriteString("\n")
		}
	} else {
		s.WriteString("Battleship Golang by packet_sent\n\n")
		s.WriteString("Please choose a gamemode:\n")
		for i := 0; i < len(gameChoices); i++ {
			if m.cursor == i {
				s.WriteString("(•) ")
			} else {
				s.WriteString("( ) ")
			}
			s.WriteString(gameChoices[i])
			s.WriteString("\n")
		}
	}
	s.WriteString("\nup/down: j/k • select: enter\n")

	return s.String()
}
