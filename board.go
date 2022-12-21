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

func startGame() {
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
	turnSwitch(0)

	if gamemode == "1 VS BOT" {
		//Player 2 (BOT) places their ships
		shipPlacementBot()
	} else if gamemode == "1 VS 1" {
		//Player 2 (HUMAN) places their ships
		shipPlacementPlayer2()
	}

	fmt.Print("\033[H\033[2J") //Used to clear the console at the start of the game
}

func shipPlacementBot() {
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
}

func shipPlacementPlayer1() {
	//Show the starting board

	fmt.Printf("[P1] Please place your %d ships on the board\n", len(boardShipSize))
	fmt.Printf("[P1] Ships left to be placed: %d\n", len(boardShipSize))
	playerBoard1(false) //if true is sent then don't show the ship and only show shot areas with X and missed areas with -

	//Store the various ship data for player 1
	player1Ships = make([]Ship, len(boardShipSize))

	//Create a reader with a default size buffer to temporarily store any text input data
	readerBuffer := bufio.NewReader(os.Stdin)

	//iterate over the number of ships (should run 5 times since there are 5 ships only)
	for i := 0; i < len(boardShipSize); {
		fmt.Print("\n\n")
		fmt.Printf("[P1] Ship Name: %s\n", boardShipName[i])               //print ship name to the player
		fmt.Printf("[P1] Ship Size: %s\n", strconv.Itoa(boardShipSize[i])) //print ship size to the player
		fmt.Print("[P1] Direction (H/V): ")                                //ask for ship direction for placing it on the board
		direction, _ := readerBuffer.ReadString('\n')                      //store data in the var and use a new line as the delimiter
		direction = strings.ToUpper(strings.TrimSpace(direction))          //remove whitespace and tabspace from the string that was just read and make everything uppercase
		if (direction == "H") || (direction == "V") {
			fmt.Print("[P1] Ship Coordinates: ")
			coordinates, _ := readerBuffer.ReadString('\n')
			coordinates = strings.ToUpper(strings.TrimSpace(coordinates)) //remove whitespace and tabspace from the string that was just read and make everything uppercase
			if len(coordinates) < 2 {                                     //Checking for the minimum required length before using the value as 2 slices
				fmt.Println("\n[P1] Wrong input for ship coordinates")
				continue //start the loop from the top again
			}
			//Only grab the 1st 2 chars of the coordinates and verify they are formatted correctly
			coordinates, verified := verifyCoordsFormat(coordinates[0:1], coordinates[1:2])
			if verified {
				boardCoords := transformCoordinates(coordinates)                                           //Transform the coordinates to plottable coords
				shipPlaced := spawnShipCoordinates(&player1Board, player1Ships, boardCoords, direction, i) //making use of pointers since its not a local variable and the value inside will be changed if a ship is placed
				if shipPlaced {
					i++ //increase the counter for the loop so it stops after all ships have been placed successfully on the board
					fmt.Print("\033[H\033[2J")
					fmt.Printf("[P1] Please place your %d ships on the board\n", len(boardShipSize))
					fmt.Printf("[P1] Ships left to be placed: %d\n", len(boardShipSize)-i) //show how many ships left to be placed (-1 each iteration)

					playerBoard1(false) //if true is sent then don't show the ship and only show shot areas with X and missed areas with -
				} else {
					fmt.Println("\n[P1] Unable to place ship at these coordinates")
				}
			} else {
				fmt.Println("\n[P1] Wrong input for ship coordinates")
			}

		} else {
			//if input is wrong then keep looping over until we get the correct direction and print the wrong direction message
			fmt.Println("\n[P1] Wrong input for ship direction")
		}
	}
}

func shipPlacementPlayer2() {
	//Show the starting board

	fmt.Printf("[P2] Please place your %d ships on the board\n", len(boardShipSize))
	fmt.Printf("[P2] Ships left to be placed: %d\n", len(boardShipSize))
	playerBoard2(false) //if true is sent then don't show the ship and only show shot areas with X and missed areas with -

	//Store the various ship data for player 1
	player2Ships = make([]Ship, len(boardShipSize))

	//Create a reader with a default size buffer to temporarily store any text input data
	readerBuffer := bufio.NewReader(os.Stdin)

	//iterate over the number of ships (should run 5 times since there are 5 ships only)
	for i := 0; i < len(boardShipSize); {
		fmt.Print("\n\n")
		fmt.Printf("[P2] Ship Name: %s\n", boardShipName[i])               //print ship name to the player
		fmt.Printf("[P2] Ship Size: %s\n", strconv.Itoa(boardShipSize[i])) //print ship size to the player
		fmt.Print("[P2] Direction (H/V): ")                                //ask for ship direction for placing it on the board
		direction, _ := readerBuffer.ReadString('\n')                      //store data in the var and use a new line as the delimiter
		direction = strings.ToUpper(strings.TrimSpace(direction))          //remove whitespace and tabspace from the string that was just read and make everything uppercase
		if (direction == "H") || (direction == "V") {
			fmt.Print("[P2] Ship Coordinates: ")
			coordinates, _ := readerBuffer.ReadString('\n')
			coordinates = strings.ToUpper(strings.TrimSpace(coordinates)) //remove whitespace and tabspace from the string that was just read and make everything uppercase
			if len(coordinates) < 2 {                                     //Checking for the minimum required length before using the value as 2 slices
				fmt.Println("\n[P2] Wrong input for ship coordinates")
				continue //start the loop from the top again
			}
			//Only grab the 1st 2 chars of the coordinates and verify they are formatted correctly
			coordinates, verified := verifyCoordsFormat(coordinates[0:1], coordinates[1:2])
			if verified {
				boardCoords := transformCoordinates(coordinates)                                           //Transform the coordinates to plottable coords
				shipPlaced := spawnShipCoordinates(&player2Board, player2Ships, boardCoords, direction, i) //making use of pointers since its not a local variable and the value inside will be changed if a ship is placed
				if shipPlaced {
					i++ //increase the counter for the loop so it stops after all ships have been placed successfully on the board
					fmt.Print("\033[H\033[2J")
					fmt.Printf("[P2] Please place your %d ships on the board\n", len(boardShipSize))
					fmt.Printf("[P2] Ships left to be placed: %d\n", len(boardShipSize)-i) //show how many ships left to be placed (-1 each iteration)

					playerBoard2(false) //print updated playerboard
				} else {
					fmt.Println("\n[P2] Unable to place ship at these coordinates")
				}
			} else {
				fmt.Println("\n[P2] Wrong input for ship coordinates")
			}

		} else {
			//if input is wrong then keep looping over until we get the correct direction and print the wrong direction message
			fmt.Println("\n[P2] Wrong input for ship direction")
		}
	}
}

func playerBoard1(hide bool) {
	fmt.Print("\n")
	fmt.Print("[P1] Player 1 (HUMAN) Board:\n")

	//iterate the X axis and only grab the key to use as the position/number
	for y := range boardAxisY {
		if y == 0 {
			boardColor.Printf("    %d ", y) //have double spaces at the first iteration for cleaner looking printing
		} else {
			//else print without the double spaces as its not needed after the first iteration
			boardColor.Printf("%d ", y)
		}
	}
	fmt.Println() //Go to a new line

	//iterate over the Y axis and use the key to get the value to print position/letter
	for x := range player1Board {
		//use the key to as the letter position and print it
		boardColor.Print(" " + boardAxisY[x] + " ")
		//iterate over Y axis slice regardless of its value being 0 or 1
		for y := 0; y <= (len(player1Board[x]) - 1); y++ {
			if player1Board[x][y] == 0 && !hide {
				//0 means there is a no ship there so print the empty marker
				if y == 0 {
					waterColor.Print(" 0 ")
				} else if y == 9 {
					waterColor.Print("0 ")
				} else {
					waterColor.Print("0 ")
				}
			} else if player1Board[x][y] == 1 && !hide {
				//1 means there is a ship there so print the ship marker
				if y == 0 {
					shipColor.Print(" S ")
				} else if y == 9 {
					shipColor.Print("S ")
				} else {
					shipColor.Print("S ")
				}
			} else if (player1Board[x][y] == 0 || player1Board[x][y] == 1) && hide {
				//0 means there is a no ship there so print the empty marker
				if y == 0 {
					waterColor.Print(" 0 ")
				} else if y == 9 {
					waterColor.Print("0 ")
				} else {
					waterColor.Print("0 ")
				}
			} else if player1Board[x][y] == -1 {
				//-1 means there is a shot ship there
				if y == 0 {
					shipShotColor.Print(" ")
					shipShotColor.Print(" X ")
				} else if y == 9 {
					shipShotColor.Print("X ")
				} else {
					shipShotColor.Print("X ")
				}
			} else {
				//if its anything else then it means an empty area was shot
				if y == 0 {
					fmt.Print(" - ")
				} else if y == 9 {
					fmt.Print("- ")
				} else {
					fmt.Print("- ")
				}
			}
		}
		fmt.Println() //Go to a new line for each time a Y axis is finished printing
	}
}

func playerBoard2(hide bool) {
	fmt.Print("\n")
	if gamemode == "1 VS 1" {
		fmt.Print("[P2] Player 2 (HUMAN) Board:\n")
	} else {
		fmt.Print("[P2] Player 2 (BOT) Board:\n")
	}

	//iterate the X axis and only grab the key to use as the position/number
	for y := range boardAxisY {
		if y == 0 {
			boardColor.Printf("    %d ", y) //have double spaces at the first iteration for cleaner looking printing
		} else {
			//else print without the double spaces as its not needed after the first iteration
			boardColor.Printf("%d ", y)
		}
	}
	fmt.Println() //Go to a new line

	//iterate over the Y axis and use the key to get the value to print position/letter
	for x := range player2Board {
		//use the key to as the letter position and print it
		boardColor.Print(" " + boardAxisY[x] + " ")
		//iterate over Y axis slice regardless of its value being 0 or 1
		for y := 0; y <= (len(player2Board[x]) - 1); y++ {
			if player2Board[x][y] == 0 && !hide {
				//0 means there is a no ship there so print the empty marker
				if y == 0 {
					waterColor.Print(" 0 ")
				} else if y == 9 {
					waterColor.Print("0 ")
				} else {
					waterColor.Print("0 ")
				}
			} else if player2Board[x][y] == 1 && !hide {
				//1 means there is a ship there so print the ship marker
				if y == 0 {
					shipColor.Print(" S ")
				} else if y == 9 {
					shipColor.Print("S ")
				} else {
					shipColor.Print("S ")
				}
			} else if (player2Board[x][y] == 0 || player2Board[x][y] == 1) && hide {
				//0 means there is a no ship there so print the empty marker
				if y == 0 {
					waterColor.Print(" 0 ")
				} else if y == 9 {
					waterColor.Print("0 ")
				} else {
					waterColor.Print("0 ")
				}
			} else if player2Board[x][y] == -1 {
				//-1 means there is a shot ship there
				if y == 0 {
					shipShotColor.Print(" ")
					shipShotColor.Print(" X ")
				} else if y == 9 {
					shipShotColor.Print("X ")
				} else {
					shipShotColor.Print("X ")
				}
			} else {
				//if its anything else then it means an empty area was shot
				if y == 0 {
					fmt.Print(" - ")
				} else if y == 9 {
					fmt.Print("- ")
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
