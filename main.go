package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"syscall"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	fmt.Print("\033[H\033[2J") //Used to clear the console at the start of the game
	p := tea.NewProgram(model{})

	// Run returns the model as a tea.Model.
	m, err := p.Run()
	if err != nil {
		fmt.Println("Cause of Crash:", err)
		os.Exit(1)
	}

	// Assert the final tea.Model to our local model and print the choice.
	if m, ok := m.(model); ok && m.choice != "" {
		gamemode = m.choice //assign the value of the choice to the global variable
		gameOver = false
		fmt.Print("\033[H\033[2J") //Used to clear the console at the start of the game
		fmt.Printf("Gamemode Selected: %s\n", gamemode)

		startGame() //start the game by placing the ships on the board

		fmt.Print("\033[H\033[2J") //Used to clear the console at the start of the game
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

		fmt.Printf("\n\nCongratulations %s for winning the game!\n\n", playerWon)
		//game over print and show option to restart or exit
	}

	p2 := tea.NewProgram(model{})
	// Run returns the model as a tea.Model.
	m2, err := p2.Run()
	if err != nil {
		fmt.Println("Cause of Crash:", err)
		os.Exit(1)
	}

	if m2, ok := m2.(model); ok && m2.choice != "" {
		fmt.Print("\033[H\033[2J") //Used to clear the console at the start of the game
		if m2.choice == "Restart" {

			//we could also restart by calling main() again but here we are showcasing a different way to restart the simple game
			err = restartUniversal()
			if err != nil {
				fmt.Println("Cause of Crash:", err)
				os.Exit(0)
			}
		} else {
			fmt.Println("Quit Program")
			os.Exit(0)
		}

	}
}

/* works on ubuntu and windows 10*/
func restartUniversal() error {
	self, err := os.Executable()
	if err != nil {
		return err
	}
	args := os.Args
	env := os.Environ()
	// Windows does not support exec syscall.
	if runtime.GOOS == "windows" {
		cmd := exec.Command(self, args[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		cmd.Env = env
		err := cmd.Run()
		if err == nil {
			os.Exit(0)
		}
		return err
	}
	return syscall.Exec(self, args, env)
}

/* only works on windows 10
func restart() error {
	path, err := os.Executable()
	if err != nil {
		return err
	}
	args := os.Args
	env := os.Environ()
	cmd := exec.Command(path, args[1:]...)
	cmd.Env = env
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err == nil {
		os.Exit(0)
	}
	return err
}
*/
