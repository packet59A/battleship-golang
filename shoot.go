package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func player1Turn() {
	fmt.Print("\033[H\033[2J")
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

		/*
			HITSCAN function explained:
			check if the shoooting coordinates are within the board limits
			check if the coordinates have been shot at already

			if any of those conditions are false then repeat the turn until a valid
			coordinate  has been shot regardless of there being a ship or not
		*/
		if !shipHitscan(&player2Board, player2Ships, Position{x: boardCoords[0], y: boardCoords[1]}) {
			player1Turn()
		} else {

			//if true is sent then don't show the ship and only show shot areas with X and missed areas with -
			playerBoard2(true)
		}
	} else {
		fmt.Println("Wrong input for shooting coordinates")
	}

}

func shipHitscan(pointerBoard *[10][10]int, ships []Ship, pos Position) bool {

	//check if the coordinates are within the board limits
	if pos.x < 0 || pos.x >= len(pointerBoard) || pos.y < 0 || pos.y >= len(pointerBoard) {
		return false
	}

	//check if the coordinates haven't been shot at before
	if pointerBoard[pos.x][pos.y] == -1 || pointerBoard[pos.x][pos.y] == -2 {
		fmt.Println("You already shot these coordinates!")
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
				fmt.Printf("Your shot just hit a ship\n")

				//since the ships aren't just 1x1 in size we will need to
				//check if we have shot all the other existing coordinates
				if ships[s].allShot() {
					ships[s].sunk = true                                                                    //change the object bool to sunk if all the whole ship was shot
					fmt.Printf("Well done! You just sunk a %s of size: %d\n", ships[s].name, ships[s].size) //print that the player has shot all the occupied ship coordinates
				}
				return true
			}
		}
	}

	pointerBoard[pos.x][pos.y] = -2 //updated board to show as empty shot location on the board when printing
	fmt.Println("You shot at nothing!")
	return true
}
