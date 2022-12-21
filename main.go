package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	fmt.Print("\033[H\033[2J") //Used to clear the console at the start of the game
	p := tea.NewProgram(model{})

	// Run returns the model as a tea.Model.
	m, err := p.Run()
	if err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}

	// Assert the final tea.Model to our local model and print the choice.
	if m, ok := m.(model); ok && m.choice != "" {
		gamemode = m.choice //assign the value of the choice to the global variable
		gameOver = false
		fmt.Print("\033[H\033[2J") //Used to clear the console at the start of the game
		fmt.Printf("Gamemode Selected: %s\n", gamemode)

		startGame() //print the initial 10x10 board without any ships
		fmt.Print("\033[H\033[2J")
		fmt.Println("Ships have been placed on both boards")

		//loop until gameOver is true

		for !gameOver {

			//player 1 turn
			player1Turn()

			if gameOver {
				break
			}
			//player 2 turn
			player2Turn()

			turn++                     //increment the turn counter
			fmt.Print("\033[H\033[2J") //clear console after each turn has finished
		}

		fmt.Printf("Congratulations %s for winning the game!\n", playerWon)
		//game over print and show option to restart or exit
	}

	p2 := tea.NewProgram(model{})
	// Run returns the model as a tea.Model.
	m2, err := p2.Run()
	if err != nil {
		fmt.Println("Oh no2:", err)
		os.Exit(1)
	}

	if m2, ok := m2.(model); ok && m2.choice != "" {
		//fmt.Print("\033[H\033[2J") //Used to clear the console at the start of the game
		if m2.choice == "Restart" {
			main()
		} else {
			fmt.Println("Quit Program")
			os.Exit(0)
		}

	}
}
