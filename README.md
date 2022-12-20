## Project Name: battleship-golang

#### Game Name: Battleship
#### Game Objective: Try and sink the enemy ships before they shink you ships.

##### Game Rules:
- [🚢] There is a list of ships with different sizes (2 Destroyer, 3 Submarine, 3 Cruiser, 4 Battleship, 5 Carrier)
- [&puncsp;📍&puncsp;] Ships for both sides have to be placed before shooting them
- [🔫] You can shot a location once since ships do not move
- [💘] If a shot hit a ship then it will marked on the board as a hit
- [❤️‍] If a shot missed a ship then it will marked on the board as a miss
- [🫠] If every part of a ship is hit then it will be all marked as sinked
- [🏆] The first side that sinks the entire fleet of 5 ships will win the game

##### How to Build & Run:
- Build Requirements: Go 1.19
- Run the following command ``go build -o battleship.exe``
- To start the game run ``./battleship.exe``

