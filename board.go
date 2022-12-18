package main

import "fmt"

func printBoard() {
	//make sure the values are initialized at 0 for each board
	//loop over the 10 slices
	for x := range player1Board {
		//loop over the 10 values inside the slice
		for y := 0; y <= (len(player1Board[x]) - 1); y++ {
			player1Board[x][y] = 0
			player2Board[x][y] = 0
		}
	}
	playerBoard1()
}

func playerBoard1() {
	fmt.Print("Player 1 Board:\n ")

	//iterate the X axis and only grab the key to use as the position/number
	for x := range boardAxisX {
		if x == 0 {
			//have double spaces at the first iteration for cleaner looking printing
			fmt.Printf("  %d ", x)
		} else {
			//else print without the double spaces as its not needed after the first iteration
			fmt.Printf("%d ", x)
		}
	}
	fmt.Println() //Go to a new line

	//iterate over the Y axis and use the key to get the value to print position/letter
	for x := range player1Board {
		//use the key to as the letter position and print it
		fmt.Print(boardAxisX[x] + " ")
		for y := range player1Board {
			if player1Board[x][y] == 0 {
				fmt.Print(" 0")
			} else if player1Board[x][y] == 1 {
				fmt.Print(" S")
			}
		}
		fmt.Println() //Go to a new line for each time a Y axis is finished printing
	}
}

//iterate the X axis and only grab the key to use as the position/number
/*
	for i, v := range boardAxisXTest {
		if i == 0 {
			//have double spaces at the first iteration for cleaner looking printing
			fmt.Printf("  %s ", string(v))
		} else {
			//else print without the double spaces as its not needed after the first iteration
			fmt.Printf("%s ", string(v))
		}
	}
*/
