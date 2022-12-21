package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func player1Turn() {
	fmt.Print("\033[H\033[2J") //clear console
	fmt.Println("You need to destroy all the enemy ships first to win!")
	fmt.Printf("[P1] Turn %d Starting \n", turn)

	playerBoard1(false) //if true is sent then don't show the ship and only show shot areas with X and missed areas with -
	playerBoard2(true)  //if true is sent then don't show the ship and only show shot areas with X and missed areas with -
	fmt.Print("\n")

	//keep looping until a valid coordinate has been shot at
	for {
		//Create a reader with a default size buffer to temporarily store any text input data
		readerBuffer := bufio.NewReader(os.Stdin)
		fmt.Print("[P1] Enter coordinates to shoot: ")
		coordinates, _ := readerBuffer.ReadString('\n')
		coordinates = strings.ToUpper(strings.TrimSpace(coordinates)) //remove whitespace and tabspace from the string that was just read and make everything uppercase
		if len(coordinates) < 2 {                                     //Checking for the minimum required length before using the value as 2 slices
			fmt.Println("\n[P1] Wrong input for shooting coordinates")
			continue //start the loop from the top again
		}
		//Only grab the 1st 2 chars of the coordinates and verify they are formatted correctly
		coordinates, verified := verifyCoordsFormat(coordinates[0:1], coordinates[1:2])
		if verified {
			boardCoords := transformCoordinates(coordinates) //Transform the coordinates to plottable coords

			/*
				HITSCAN function explained:
				check if the shoooting coordinates are within the board limits
				check if the coordinates have been shot at already

				if any of those conditions are false then repeat the turn until a valid
				coordinate  has been shot regardless of there being a ship or not
			*/
			if shipHitscan(&player2Board, player2Ships, Position{x: boardCoords[0], y: boardCoords[1]}, false) {
				//time.Sleep(time.Millisecond * 1200) //wait for 1.2s before moving on so the message from the hitscan can be read
				break //start the loop from the top again
			} else {
				continue
			}
		} else {
			fmt.Println("\n[P1] Wrong input for shooting coordinates")
			continue //start the loop from the top again
		}
	}

	//Show turn ending board
	//fmt.Print("\033[H\033[2J")
	fmt.Printf("[P1] Turn %d Ending\n", turn)
	playerBoard2(true)          //if true is sent then don't show the ship and only show shot areas with X and missed areas with -
	time.Sleep(time.Second * 2) //delay before finishing the function so the user can read the updated board
}

func player2Turn() {
	fmt.Print("\033[H\033[2J") //clear console
	if gamemode == "1 VS BOT" {
		//fmt.Printf("[P2 BOT] Turn %d Ending\n", turn)
		//Player 2 (BOT) shoots the ships

		//keep looping until a valid coordinate has been shot at
		for {
			rand.Seed(time.Now().UnixNano())                       //Generate time based seed
			boardCoords := transformCoordinates(generatedCoords()) //Transform the randomly generated coordinates to plottable coords

			//check if the coordinates haven't been shot at before
			if player1Board[boardCoords[0]][boardCoords[1]] == -1 || player1Board[boardCoords[0]][boardCoords[1]] == -2 {
				continue //start the loop from the top again
			} else {
				shipHitscan(&player1Board, player1Ships, Position{x: boardCoords[0], y: boardCoords[1]}, true)
				break //get out of the for loop
			}
		}
		//fmt.Printf("[P2 BOT] Turn %d Ending\n", turn)
		//fmt.Print("\033[H\033[2J")
	} else if gamemode == "1 VS 1" {
		//Player 2 (HUMAN) shoots their ships

		//fmt.Print("\033[H\033[2J")
		fmt.Println("[P2] You need to destroy all the enemy ships first to win!")
		fmt.Printf("[P2] Turn %d Starting \n", turn)

		playerBoard2(false) //if true is sent then don't show the ship and only show shot areas with X and missed areas with -
		playerBoard1(true)  //if true is sent then don't show the ship and only show shot areas with X and missed areas with -
		fmt.Print("\n")

		//keep looping until a valid coordinate has been shot at
		for {
			//Create a reader with a default size buffer to temporarily store any text input data
			readerBuffer := bufio.NewReader(os.Stdin)
			fmt.Print("[P2] Enter coordinates to shoot: ")
			coordinates, _ := readerBuffer.ReadString('\n')
			coordinates = strings.ToUpper(strings.TrimSpace(coordinates)) //remove whitespace and tabspace from the string that was just read and make everything uppercase
			if len(coordinates) < 2 {                                     //Checking for the minimum required length before using the value as 2 slices
				fmt.Println("\n[P2] Wrong input for shooting coordinates")
				continue //start the loop from the top again
			}
			//Only grab the 1st 2 chars of the coordinates and verify they are formatted correctly
			coordinates, verified := verifyCoordsFormat(coordinates[0:1], coordinates[1:2])
			if verified {
				boardCoords := transformCoordinates(coordinates) //Transform the coordinates to plottable coords

				/*
					HITSCAN function explained:
					check if the shoooting coordinates are within the board limits
					check if the coordinates have been shot at already

					if any of those conditions are false then repeat the turn until a valid
					coordinate  has been shot regardless of there being a ship or not
				*/
				if shipHitscan(&player1Board, player1Ships, Position{x: boardCoords[0], y: boardCoords[1]}, false) {
					//time.Sleep(time.Millisecond * 1200) //wait for 1.2s before moving on so the message from the hitscan can be read
					break //start the loop from the top again
				} else {
					continue
				}
			} else {
				fmt.Println("\n[P2] Wrong input for shooting coordinates")
				continue //start the loop from the top again
			}
		}

		//Show turn ending board
		//fmt.Print("\033[H\033[2J")
		fmt.Printf("[P2] Turn %d Ending\n", turn)
		playerBoard2(true)          //if true is sent then don't show the ship and only show shot areas with X and missed areas with -
		time.Sleep(time.Second * 2) //delay before finishing the function so the user can read the updated board
	}

	//Show turn ending board
	//fmt.Print("\033[H\033[2J")
}

// The hide bool in the hitscan function is used for when using a bot to hide the outputs
func shipHitscan(pointerBoard *[10][10]int, ships []Ship, pos Position, hide bool) bool {

	//check if the coordinates are within the board limits
	if pos.x < 0 || pos.y < 0 || pos.x >= len(pointerBoard) || pos.y >= len(pointerBoard) {
		return false
	}

	//check if the coordinates haven't been shot at before
	if pointerBoard[pos.x][pos.y] == -1 || pointerBoard[pos.x][pos.y] == -2 {
		if !hide {
			fmt.Println("Coordinates already shot before!")
		}
		return false
	}

	//check if any ship is present where we are shooting
	for s := range ships {
		for p := range ships[s].position {

			//if a ship is present at the coordinates we are shooting then change value to shot on the board
			if ships[s].position[p].x == pos.x && ships[s].position[p].y == pos.y {
				pointerBoard[pos.x][pos.y] = -1 //updated board to show as ship shot location on the board when printing
				ships[s].position[p].x = -1     //remove the coordinates from the ship object as this coordinate was hit
				ships[s].position[p].y = -1     //remove the coordinates from the ship object as this coordinate was hit
				if !hide {
					fmt.Printf("Shot hit a ship!\n")
				}

				//since the ships aren't just 1x1 in size we will need to
				//check if we have shot all the other existing coordinates
				if ships[s].allShot() {
					ships[s].sunk = true //change the object bool to sunk if all the whole ship was shot
					winnerCheck()
					if !hide {
						fmt.Printf("Well done! %s ship of size %d just sunk!\n", ships[s].name, ships[s].size) //print that the player has shot all the occupied ship coordinates
					}
				}
				return true
			}
		}
	}

	pointerBoard[pos.x][pos.y] = -2 //updated board to show as empty shot location on the board when printing
	if !hide {
		fmt.Println("Shot at nothing!")
	}
	return true
}

func winnerCheck() {
	chanPlayer := make(chan string) //create once channel since there can be only one winner per game

	//using goroutines to check who won simultaneously and this won't be an issue as there can only be one winner per game
	go func() {
		playerWon := true

		//if the enemy ships have all been shot then this should return false and not send anything to chan bool
		for s := range player2Ships {
			if !player2Ships[s].sunk {
				playerWon = false
			}
		}
		if playerWon {
			chanPlayer <- "Player 1 (HUMAN)" //sends the value via the string channel
		} else {
			chanPlayer <- "" //sends the value via the string channel
		}
	}()

	//using goroutines to check who won simultaneously and this won't be an issue as there can only be one winner per game
	go func() {
		playerWon := true

		///if the enemy ships have all been shot then this should return false and not send anything to chan bool
		for s := range player1Ships {
			if !player1Ships[s].sunk {
				playerWon = false
			}
		}

		if playerWon && gamemode == "1 VS 1" {
			chanPlayer <- "Player 2 (HUMAN)" //sends the value via the string channel
		} else if playerWon && gamemode == "1 VS BOT" {
			chanPlayer <- "Player 2 (BOT)" //sends the value via the string channel
		} else {
			chanPlayer <- "" //sends the value via the string channel
		}
	}()

	playerWon = <-chanPlayer //wait for any of the goroutine to finish 1st before moving down the code

	//if a player won then announce it
	if playerWon != "" {
		gameOver = true
	}
}
