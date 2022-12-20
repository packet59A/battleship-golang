package main

import (
	"fmt"
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
		case "ctrl+c", "q", "esc":
			return m, tea.Quit

		case "enter":
			// Send the choice on the channel and exit.
			m.choice = choices[m.cursor]
			return m, tea.Quit

		case "j":
			m.cursor++
			if m.cursor >= len(choices) {
				m.cursor = 0
			}

		case "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(choices) - 1
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	s := strings.Builder{}
	s.WriteString("Battleship Golang by packet_sent\n\n")

	s.WriteString("Please choose a gamemode:\n")
	for i := 0; i < len(choices); i++ {
		if m.cursor == i {
			s.WriteString("(•) ")
		} else {
			s.WriteString("( ) ")
		}
		s.WriteString(choices[i])
		s.WriteString("\n")
	}
	s.WriteString("\nup/down: j/k • select: enter • esc/ctrl+c: quit\n")

	return s.String()
}

func main() {
	p := tea.NewProgram(model{})

	// Run returns the model as a tea.Model.
	m, err := p.Run()
	if err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}

	// Assert the final tea.Model to our local model and print the choice.
	if m, ok := m.(model); ok && m.choice != "" {
		fmt.Print("\033[H\033[2J") //Used to clear the console at the start of the game
		fmt.Print("Battleship Golang by packet_sent\n\n")
		fmt.Printf("Starting %s Gamemode\n", m.choice)

		//loop until gameOver is true (finished game with a winner)

		if !gameOver {

			startGame(m.choice) //print the initial 10x10 board without any ships
			//player 1 turn
			//player 2 turn
		}

		//game over print and show option to restart or exit
	}
}
