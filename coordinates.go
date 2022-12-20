package main

import (
	"strconv"
)

// by not defaulting the variables to player1 we can reuse this function for player2 when placing ships
func spawnShipCoordinates(playerBoard int, tempPlayerBoard [10][10]int, playerShips []Ship, boardCoords [2]int, direction string, index int) bool {

	//look at the direction and do different calculations accordingly
	if direction == "V" {

		//check if it is possible to have a ship in this direction without being out of bounds of the board and without overlapping another ship
		/*
			LOGIC EXPLAINED:
			X is where the ship starts
			check that the value for X is larger or equal to 0
			check that the value for X and adding the ship size means that it will still fit in the board (<=10)
			check that the value for Y is larger or equal to 0
			check that the value for Y is not going over the board boundaries whilst taking account for the ship width of 1 unit as its vertically placed
			check that the ship being placed will not overlap with any other ship by checking through each ship object
		*/
		if boardCoords[0] >= 0 && boardCoords[1] >= 0 && boardCoords[1] < len(tempPlayerBoard) && boardCoords[0]+boardShipSize[index] <= len(tempPlayerBoard) &&
			shipOverlapCheck(boardCoords[0], boardCoords[1], boardShipSize[index], direction, playerShips) {

			//create a new ship with the specified coordinates
			playerShips[index].create(boardShipSize[index])
			for i := 0; i < boardShipSize[index]; i++ {
				//update the tempPlayerBoard values to indicate there is a ship (1 = ship, 0 = no ship)
				tempPlayerBoard[boardCoords[0]+i][boardCoords[1]] = 1

				//add the coordinates of the ship to the newly created object
				playerShips[index].addCoords(boardCoords[0]+i, boardCoords[1], i)
			}
		} else {
			//return false if its not possible to place a ship at the specified coordinates/direction
			return false
		}
	} else if direction == "H" {

		//check if it is possible to have a ship in this direction without being out of bounds of the board and without overlapping another ship
		/*
			LOGIC EXPLAINED:
			Y is where the ship starts
			check that the value for Y is larger or equal to 0
			check that the value for Y and adding the ship size means that it will still fit in the board (<=10)
			check that the value for X is larger or equal to 0
			check that the value for X is not going over the board boundaries whilst taking account for the ship width of 1 unit as its vertically placed
			check that the ship being placed will not overlap with any other ship by checking through each ship object
		*/
		if boardCoords[1] >= 0 && boardCoords[0] >= 0 && boardCoords[0] < len(tempPlayerBoard) && boardCoords[1]+boardShipSize[index] <= len(tempPlayerBoard) &&
			shipOverlapCheck(boardCoords[0], boardCoords[1], boardShipSize[index], direction, playerShips) {

			//create a new ship with the specified coordinates
			playerShips[index].create(boardShipSize[index])
			for i := 0; i < boardShipSize[index]; i++ {

				//update the tempPlayerBoard values to indicate there is a ship (1 = ship, 0 = no ship)
				tempPlayerBoard[boardCoords[0]][boardCoords[1]+i] = 1

				//add the coordinates of the ship to the newly created object
				playerShips[index].addCoords(boardCoords[0], boardCoords[1]+i, i)
			}
		} else {
			//return false if its not possible to place a ship at the specified coordinates/direction
			return false
		}
	}

	//save the local variable tempPlayerBoard data to the actual board that will be printed in the console
	if playerBoard == 1 {
		player1Board = tempPlayerBoard
	} else {
		player2Board = tempPlayerBoard
	}

	return true
}

func shipOverlapCheck(x, y, size int, direction string, ships []Ship) bool {

	//check for direction as the axis we are checking will be depending on it (x, y)
	if direction == "V" {

		//go through each ship
		for i := range ships {

			//go through each of its position values (x, y)
			for p := range ships[i].position {

				//since ships aren't 1x1, loop by its possible size  (2, 3, 3, 4, 5)
				for s := 0; s < size; s++ {

					//check for the coordinates they can cover on both axis
					if x+s == ships[i].position[p].x && y == ships[i].position[p].y {
						return false
					}
				}
			}
		}
	} else if direction == "H" {

		//go through each ship
		for i := range ships {

			//go through each of its position values (x, y)
			for p := range ships[i].position {

				//since ships aren't 1x1, loop by its possible size  (2, 3, 3, 4, 5)
				for s := 0; s < size; s++ {

					//check for the coordinates they can cover on both axis
					if y+s == ships[i].position[p].y && x == ships[i].position[p].x {
						return false
					}
				}
			}
		}
	}
	return true
}

func transformCoordinates(str string) [2]int {
	var boardCoords [2]int //create a slice to hold the board coordinates

	//Using a switch statement reduced the code clutter by repeating elif statements,
	//even if elif statements were to be used performance wouldn't be affected since
	//the go compiler will automatically optimize the code on compile
	switch str[0:1] {
	case "A":
		boardCoords[0] = 0 //save the x board coordinates
	case "B":
		boardCoords[0] = 1 //save the x board coordinates
	case "C":
		boardCoords[0] = 2 //save the x board coordinates
	case "D":
		boardCoords[0] = 3 //save the x board coordinates
	case "E":
		boardCoords[0] = 4 //save the x board coordinates
	case "F":
		boardCoords[0] = 5 //save the x board coordinates
	case "G":
		boardCoords[0] = 6 //save the x board coordinates
	case "H":
		boardCoords[0] = 7 //save the x board coordinates
	case "I":
		boardCoords[0] = 8 //save the x board coordinates
	case "J":
		boardCoords[0] = 9 //save the x board coordinates
	}

	yCoords, _ := strconv.Atoi(str[1:2]) //convert the string to an int
	boardCoords[1] = yCoords             //save the y board coordinates

	return boardCoords //return the slice of the board coordinates to be used when plotting the ships
}

func verifyCoordsFormat(xAxis, yAxis string) (string, bool) {

	//get char1 rune
	for _, r := range xAxis {
		//verifies if rune value is corresponding to "ABCDEFGHIJ"
		if !(r > 64 && r < 75) {
			return "", false
		}
	}

	//get char2 rune
	for _, r := range yAxis {
		//verifies if rune value is corresponding to "012345789"
		if !(r > 47 && r < 58) {
			return "", false
		}
	}

	return xAxis + yAxis, true
}
