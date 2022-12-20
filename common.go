package main

// Global vars to be used througout the code
var (
	gameOver bool //used for checking before displaying options after a finished turn

	//board sizing (max x = 10 and max y = 10)
	player1Board [10][10]int //contains the board of 10 slice (x) of 10 numbers each (y)
	player2Board [10][10]int //contains the board of 10 slice (x) of 10 numbers each (y)

	boardAxisY = [10]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}

	boardShipSize = [5]int{2, 3, 3, 4, 5}                                                   //Ship sizes decided by the official game rules
	boardShipName = [5]string{"Destroyer", "Submarine", "Cruiser", "Battleship", "Carrier"} //Ship names decided by the official game rules

	player1Ships []Ship
	player2Ships []Ship
)

type Ship struct {
	sunk     bool
	size     int
	position []Position
}

type Position struct {
	x, y int
}

func (ship *Ship) addCoords(xPos, yPos, index int) {
	//fmt.Println("xP", xP, "yP", yP, "index", index)
	ship.position[index] = Position{x: xPos, y: yPos}
}

func (ship *Ship) create(shipSize int) {
	ship.size = shipSize
	ship.sunk = false
	ship.position = make([]Position, shipSize)
}

var choices = []string{"1 VS BOT", "1 VS 1"}

type model struct {
	cursor int
	choice string
}