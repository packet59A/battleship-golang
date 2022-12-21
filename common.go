package main

// Global vars to be used througout the code
var (
	gameOver bool //used for checking before displaying options after a finished turn

	//board sizing (max x = 10 and max y = 10)
	player1Board [10][10]int //contains the board of 10 slice (x) of 10 numbers each (y)
	player2Board [10][10]int //contains the board of 10 slice (x) of 10 numbers each (y)

	boardAxisY = [10]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"} //Y Axis label for the board

	boardShipSize = [5]int{2, 3, 3, 4, 5}                                                   //Ship sizes decided by the official game rules
	boardShipName = [5]string{"Destroyer", "Submarine", "Cruiser", "Battleship", "Carrier"} //Ship names decided by the official game rules

	player1Ships []Ship //used for holding the player's ship data
	player2Ships []Ship //used for holding the player's ship data

	turn = 1 //start the game at turn 1

	choices  = []string{"1 VS BOT", "1 VS 1"} //current gamemode options to show in the menu
	gamemode string                           //holds the user chosen gamemode and will be used for when deciding what functions to run
)

type Ship struct {
	sunk     bool
	size     int
	name     string
	position []Position
}

type Position struct {
	x, y int
}

type model struct {
	cursor int
	choice string
}

func (ship *Ship) addCoords(xPos, yPos, index int) {
	//fmt.Println("xP", xP, "yP", yP, "index", index)
	ship.position[index] = Position{x: xPos, y: yPos}
}

func (ship *Ship) create(shipSize, index int) {
	ship.size = shipSize
	ship.sunk = false
	ship.name = boardShipName[index]
	ship.position = make([]Position, shipSize)
}

func (ship *Ship) allShot() bool {
	//loop through all ship coordinates occupied
	for p := range ship.position {

		//if any of the ship coordinates haven't been shot then return false
		if ship.position[p].x != -1 || ship.position[p].y != -1 {
			return false
		}
	}

	//if all ship coordinates were shot then return true
	return true
}
