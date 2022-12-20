package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func player1Turn() {
	//Create a reader with a default size buffer to temporarily store any text input data
	readerBuffer := bufio.NewReader(os.Stdin)
	fmt.Print("\n")
	fmt.Print("Enter coordinates to shoot: ")
	coordinates, _ := readerBuffer.ReadString('\n')
	coordinates = strings.ToUpper(strings.TrimSpace(coordinates)) //remove whitespace and tabspace from the string that was just read and make everything uppercase

	//Only grab the 1st 2 chars of the coordinates and verify they are formatted correctly
	coordinates, verified := verifyCoordsFormat(coordinates[0:1], coordinates[1:2])
	if verified {
		boardCoords := transformCoordinates(coordinates) //Transform the coordinates to plottable coords
		fmt.Println("Coordinates3:", boardCoords)
	} else {
		fmt.Println("Wrong input for shooting coordinates")
	}

}
