package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func startGame(gamemode string) {
	//make sure the values are initialized at 0 for each board

	//loop over the 10 slices
	for x := range player1Board {
		//iterate over Y axis slice
		for y := 0; y <= (len(player1Board[x]) - 1); y++ {
			player1Board[x][y] = 0 //reset the board values to zero, mainly to be used for when restarting the game
			player2Board[x][y] = 0 //both boards are equal in size (10x10) so lets reset the values for the second board as well
		}
	}

	//Player (HUMAN) places their ships
	shipPlacementPlayer1()

	if gamemode == "1 VS BOT" {
		//Player 2 (BOT) places their ships
		shipPlacementBot()
	} else if gamemode == "1 VS 1" {
		//Player 2 (HUMAN) places their ships
		//shipPlacementPlayer2()
		shipPlacementBot()
	}
}

func shipPlacementBot() {
	fmt.Print("\033[H\033[2J")
	//Store the various ship data for player 2
	player2Ships = make([]Ship, len(boardShipSize))

	//iterate over the number of ships (should run 5 times since there are 5 ships only)
	for i := 0; i < len(boardShipSize); {
		rand.Seed(time.Now().UnixNano()) //Generate time based seed
		direction := []string{           //Create a string slice and randomly choose a value from within the slice
			"H",
			"V",
		}[rand.Intn(2)]
		boardCoords := transformCoordinates(generatedCoords())                                     //Transform the randomly generated coordinates to plottable coords
		shipPlaced := spawnShipCoordinates(&player2Board, player2Ships, boardCoords, direction, i) //making use of pointers since its not a local variable and the value inside will be changed if a ship is placed
		//since the coordinates are randomly generated the bot will automatically retry ship placement if it can't be placed
		if shipPlaced {
			i++ //increase the counter for the loop so it stops after all ships have been placed successfully on the board
		}

	}

	fmt.Println("Enemy ships placed")
	//if true is sent then don't show the ship and only show shot areas with X and missed areas with -
	playerBoard2(true)
}

func generatedCoords() string {
	rand.Seed(time.Now().UnixNano())                     //Generate time based seed
	randomCharacter := string('A' + rune(rand.Intn(10))) //Generate random character from A-J
	randomNumber := strconv.Itoa(rand.Intn(10))          //Generate random character from 0-10

	return (randomCharacter + randomNumber)
}

func shipPlacementPlayer1() {
	//Show the starting board
	playerBoard1()

	//Store the various ship data for player 1
	player1Ships = make([]Ship, len(boardShipSize))

	//Create a reader with a default size buffer to temporarily store any text input data
	readerBuffer := bufio.NewReader(os.Stdin)

	//iterate over the number of ships (should run 5 times since there are 5 ships only)
	for i := 0; i < len(boardShipSize); {
		fmt.Print("\n\n")
		fmt.Printf("Ship Name: %s\n", boardShipName[i])               //print ship name to the player
		fmt.Printf("Ship Size: %s\n", strconv.Itoa(boardShipSize[i])) //print ship size to the player
		fmt.Print("Direction (H/V): ")                                //ask for ship direction for placing it on the board
		direction, _ := readerBuffer.ReadString('\n')                 //store data in the var and use a new line as the delimiter
		direction = strings.ToUpper(strings.TrimSpace(direction))     //remove whitespace and tabspace from the string that was just read and make everything uppercase
		if (direction == "H") || (direction == "V") {
			fmt.Print("Ship Coordinates: ")
			coordinates, _ := readerBuffer.ReadString('\n')
			coordinates = strings.ToUpper(strings.TrimSpace(coordinates)) //remove whitespace and tabspace from the string that was just read and make everything uppercase

			//Only grab the 1st 2 chars of the coordinates and verify they are formatted correctly
			coordinates, verified := verifyCoordsFormat(coordinates[0:1], coordinates[1:2])
			if verified {
				boardCoords := transformCoordinates(coordinates)                                           //Transform the coordinates to plottable coords
				shipPlaced := spawnShipCoordinates(&player1Board, player1Ships, boardCoords, direction, i) //making use of pointers since its not a local variable and the value inside will be changed if a ship is placed
				if shipPlaced {
					i++                        //increase the counter for the loop so it stops after all ships have been placed successfully on the board
					fmt.Print("\033[H\033[2J") //clear console ahead of next ship placement
					playerBoard1()             //print updated playerboard
				} else {
					fmt.Println("Unable to place ship at these coordinates")
				}
			} else {
				fmt.Println("Wrong input for ship coordinates")
			}

		} else {
			//if input is wrong then keep looping over until we get the correct direction and print the wrong direction message
			fmt.Println("Wrong input for ship direction")
		}
	}
}

func playerBoard1() {
	fmt.Print("\n\n")
	fmt.Print("Player 1 Board:\n")

	//iterate the X axis and only grab the key to use as the position/number
	for y := range boardAxisY {
		if y == 0 {
			//have double spaces at the first iteration for cleaner looking printing
			fmt.Printf("  %d ", y)
		} else {
			//else print without the double spaces as its not needed after the first iteration
			fmt.Printf("%d ", y)
		}
	}
	fmt.Println() //Go to a new line

	//iterate over the Y axis and use the key to get the value to print position/letter
	for x := range player1Board {
		//use the key to as the letter position and print it
		fmt.Print(boardAxisY[x])
		//iterate over Y axis slice regardless of its value being 0 or 1
		for y := 0; y <= (len(player1Board[x]) - 1); y++ {
			if player1Board[x][y] == 0 {
				//0 means there is a no ship there so print the empty marker
				if y == 0 {
					fmt.Print(" 0 ")
				} else if y == 9 {
					fmt.Print("0  ")
				} else {
					fmt.Print("0 ")
				}
			} else if player1Board[x][y] == 1 {
				//1 means there is a ship there so print the ship marker
				if y == 0 {
					fmt.Print(" S ")
				} else if y == 9 {
					fmt.Print("S  ")
				} else {
					fmt.Print("S ")
				}
			} else if player1Board[x][y] == -1 {
				//-1 means there is a ship there so print the ship marker
				if y == 0 {
					fmt.Print(" X ")
				} else if y == 9 {
					fmt.Print("X  ")
				} else {
					fmt.Print("X ")
				}
			} else {
				if y == 0 {
					fmt.Print(" - ")
				} else if y == 9 {
					fmt.Print("-  ")
				} else {
					fmt.Print("- ")
				}
			}
		}
		fmt.Println() //Go to a new line for each time a Y axis is finished printing
	}
}

func playerBoard2(bot bool) {
	fmt.Print("\n\n")
	fmt.Print("Player 2 Board:\n")

	//iterate the X axis and only grab the key to use as the position/number
	for y := range boardAxisY {
		if y == 0 {
			//have double spaces at the first iteration for cleaner looking printing
			fmt.Printf("  %d ", y)
		} else {
			//else print without the double spaces as its not needed after the first iteration
			fmt.Printf("%d ", y)
		}
	}
	fmt.Println() //Go to a new line

	//iterate over the Y axis and use the key to get the value to print position/letter
	for x := range player2Board {
		//use the key to as the letter position and print it
		fmt.Print(boardAxisY[x])
		//iterate over Y axis slice regardless of its value being 0 or 1
		for y := 0; y <= (len(player2Board[x]) - 1); y++ {
			if player2Board[x][y] == 0 && !bot {
				//0 means there is a no ship there so print the empty marker
				if y == 0 {
					fmt.Print(" 0 ")
				} else if y == 9 {
					fmt.Print("0  ")
				} else {
					fmt.Print("0 ")
				}
			} else if player2Board[x][y] == 1 && !bot {
				//1 means there is a ship there so print the ship marker
				if y == 0 {
					fmt.Print(" S ")
				} else if y == 9 {
					fmt.Print("S  ")
				} else {
					fmt.Print("S ")
				}
			} else if (player2Board[x][y] == 0 || player2Board[x][y] == 1) && bot {
				//0 means there is a no ship there so print the empty marker
				if y == 0 {
					fmt.Print(" 0 ")
				} else if y == 9 {
					fmt.Print("0  ")
				} else {
					fmt.Print("0 ")
				}
			} else if player2Board[x][y] == -1 {
				//-1 means there is a shot ship there
				if y == 0 {
					fmt.Print(" X ")
				} else if y == 9 {
					fmt.Print("X  ")
				} else {
					fmt.Print("X ")
				}
			} else {
				//if its anything else then it means an empty area was shot
				if y == 0 {
					fmt.Print(" - ")
				} else if y == 9 {
					fmt.Print("-  ")
				} else {
					fmt.Print("- ")
				}
			}
		}
		fmt.Println() //Go to a new line for each time a Y axis is finished printing
	}
}

//iterate the X axis and only grab the key to use as the position/number
/*
	for i, v := range boardAxisYTest {
		if i == 0 {
			//have double spaces at the first iteration for cleaner looking printing
			fmt.Printf("  %s ", string(v))
		} else {
			//else print without the double spaces as its not needed after the first iteration
			fmt.Printf("%s ", string(v))
		}
	}
*/
