package main

//Global vars to be used througout the code
var (
	gameOver bool //used for checking before displaying options after a finished turn

	//board sizing (max x = 10 and max y = 10)
	player1Board [10][10]int //contains the board of 10 slice (x) of 10 numbers each (y)
	player2Board [10][10]int //contains the board of 10 slice (x) of 10 numbers each (y)

	boardAxisX = [10]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
)

func main() {

	//loop until gameOver is true (finished game with a winner)

	if !gameOver {

		printBoard() //print the initial 10x10 board without any ships
		//player 1 turn
		//player 2 turn
	}

	//game over print and show option to restart or exit
}
