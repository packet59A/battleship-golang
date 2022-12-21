package main

import (
	"fmt"
	"time"
)

func timeRemaining(t time.Time) int {
	currentTime := time.Now()
	difference := int(currentTime.Sub(t).Seconds())

	return difference
}

func turnSwitch(i int) {
	fmt.Print("\033[H\033[2J") //Used to clear the console

	if gamemode == "1 VS 1" {
		if i == 1 { //Used after player 1 finished their turn
			fmt.Println("[GAME] Please stop looking at the screen and pass over to player 2")
			fmt.Println("[GAME] Player switchover in 5 seconds.")
			timerStart := time.Now()

			//loop every 1 second to print the remaining time left
			for range time.Tick(1 * time.Second) {
				timeRemaining := timeRemaining(timerStart)

				if timeRemaining >= 5 {
					fmt.Println("[GAME] Player Switchover 2 -> 1")
					break
				}

				fmt.Printf("[GAME] %d seconds until Player 1 starts Turn %d\n", timeRemaining, turn+1)
			}
		} else if i == 2 { //Used after player 2 finished their turn
			fmt.Println("[GAME] Please stop looking at the screen and pass over to player 1")
			fmt.Println("[GAME] Player switchover in 5 seconds.")
			timerStart := time.Now()

			//loop every 1 second
			for range time.Tick(1 * time.Second) {
				timeRemaining := timeRemaining(timerStart)

				if timeRemaining >= 5 {
					fmt.Println("[GAME] Player Switchover 1 -> 2")
					break
				}

				fmt.Printf("[GAME] %d seconds until Player 2 starts Turn %d\n", timeRemaining, turn)
			}
		} else if i == 0 { //Used for ship placing delay on player 1->2
			fmt.Println("[GAME] Please stop looking at the screen and pass over to player 2")
			fmt.Println("[GAME] Player switchover in 5 seconds.")
			timerStart := time.Now()

			//loop every 1 second to print the remaining time left
			for range time.Tick(1 * time.Second) {
				timeRemaining := timeRemaining(timerStart)

				if timeRemaining >= 5 {
					fmt.Println("[GAME] Player Switchover 1 -> 2")
					break
				}

				fmt.Printf("[GAME] %d seconds until Player 2 starts placing ships\n", timeRemaining)
			}
		}
	}

	fmt.Print("\033[H\033[2J") //Used to clear the console
}
