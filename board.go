package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func printBoard() {
	//make sure the values are initialized at 0 for each board

	//loop over the 10 slices
	for x := range player1Board {
		//iterate over Y axis slice
		for y := 0; y <= (len(player1Board[x]) - 1); y++ {
			player1Board[x][y] = 0 //reset the board values to zero, mainly to be used for when restarting the game
			player2Board[x][y] = 0 //both boards are equal in size (10x10) so lets reset the values for the second board as well
		}
	}

	//Player 1 places their ships
	shipPlacementPlayer1()

	//Player 2 places their ships
	//shipPlacementPlayer2()
}

func shipPlacementPlayer1() {
	//Show the starting board
	playerBoard1()

	//Store the various ship data for player 1
	player1Ships = make([]Ship, len(boardShipSize))

	//Create a reader with a default size buffer to temporarily store any text input data
	readerBuffer := bufio.NewReader(os.Stdin)

	//iterate over the number of ships
	for i := 0; i < len(boardShipSize); {
		fmt.Print("\n\n")
		fmt.Printf("Ship Name: %s\n", boardShipName[i])
		fmt.Printf("Ship Size: %s\n", strconv.Itoa(boardShipSize[i]))
		fmt.Printf("Direction (H/V): ")
		direction, _ := readerBuffer.ReadString('\n') //store data in the var and use a new line as the delimiter
		fmt.Println(strings.TrimSpace(direction))

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
