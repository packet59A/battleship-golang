package main

import (
	"strconv"
)

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
