## Advanced Programming 2 Assignment by Wares Islam ([Video](https://drive.google.com/file/d/1d81wOOo7UKAHNNwhTqpgKnUH1NnTvQIX/view?usp=share_link))
##### Assignment Objective:
Create a game that has been agreed with the lecturer and consists of the things we have been through in our Advanced Programming Lectures and previous modules.

The program we are making should have used some of the following things:
- Decomposition, Abstration, Pattern Recognition and Algorithms 
- Structure
- Object Oriented Programming
- Test driven development
- Code reuse
- Testing 
- Inheritance
- and more...

## Assignment Solution
##### Game Name: 
Battleship

##### Game Objective: 
Try and sink the enemy ships before they shink you ships.

##### Game Rules:
- [🚢] There is a list of ships with different sizes (2 Destroyer, 3 Submarine, 3 Cruiser, 4 Battleship, 5 Carrier)
- [&puncsp;📍&puncsp;] Ships for both sides have to be placed before shooting them
- [🔫] You can shot a location once since ships do not move
- [💘] If a shot hit a ship then it will marked on the board as a hit
- [❤️‍] If a shot missed a ship then it will marked on the board as a miss
- [&puncsp;🫠&puncsp;] If every part of a ship is hit then it will be all marked as sinked
- [🏆] The first side that sinks the entire fleet of 5 ships will win the game

##### Game Variations:
- [1 VS 1] The player will be competing against another player
- [1 VS BOT] The player will be competing against a bot

##### Game Running Screenshots:
The game was mainly tested on Windows 10 with Windows Terminal & CMD so they are expected to work perfectly with them.
The colours supported by the CMD tool on windows is very small so it's best to use the Windows Terminal to get the best experience.
**Best Case Scenario (Windows Terminal)**   
![Best Case Scenario](https://cdn.discordapp.com/attachments/683383222578839626/1055031353160237056/image.png)
**Worst Case Scenario (Windows CMD)**  
![Worst Case Scenario](https://cdn.discordapp.com/attachments/683383222578839626/1055031410395713586/image.png)

##### How to Build & Run Locally:
- Build Requirements: Go 1.19
- Install [GoReleaser](https://goreleaser.com/install/)
- Run the following command: 
    - ``goreleaser release --snapshot --rm-dist``

- To start the game run:
    - ``.\dist\BattleshipGo_windows_amd64_v1\BattleshipGo.exe``

##### How to Build & Upload:
- Build Requirements: Go 1.19
- Get your [GITHUB API TOKEN](https://github.com/settings/tokens/new) then do one of the following:
    - Read https://goreleaser.com/scm/github/?h=api#api-token
    - Edit the ``example.sh`` and add your GITHUB API TOKEN 

- Make sure your origin and local repo are both synced
- Make sure your next version is an increment of the current version:
    - To check the current version ``git describe --tags --abbrev=0``

- Install [GoReleaser](https://goreleaser.com/install/)
- Run the following commands: 
    - ``git tag -a v1.0.0 -m "Release message"``
    - ``git push origin v1.0.0``
    - ``goreleaser release --rm-dist``
    - For more info https://goreleaser.com/quick-start/#build-only-mode

- A successful build log should look like this:
![Successful build](https://cdn.discordapp.com/attachments/683383222578839626/1055059825161150494/image.png)

## Assignment Breakdown
1a. To start the assignment I started by looking at the provided resources, and when I looked at the AdaShip repository and compared it to other games that I could make I decided to create Battleship but due to the time limitations (very busy end of year work schedule) with the project I decided to create it in Go instead of C++ as I have almost 2 years of experience with it with actual large scale live projects (https://speedproxies.net/ & https://speedshare.app/) so remaking a "simple" game like Battleship shouldn't be that hard at all.

Since I have played some variations of Battleship before I know how the game works and the mechanics behind it, but implementing in code isn't that simple so splitting up the game into 3 simple sections was very helpful as I could use it as guidance later on.
- Menu
- Place ships
- Shoot ships

1b. This was the simplest way I could represent the idea of the game without going in too much detail about it 
![Flowchart](https://cdn.discordapp.com/attachments/683383222578839626/1055082972388282388/image.png)
1c. The initial plan started off with creating the basic 10x10 board with the labeled axis and move on from there to placing ships on the board and move on from there step by step as a basic "framework" was needed before moving onto more advanced features like hitscan/hitdetection and etc.

<img src="https://cdn.discordapp.com/attachments/683383222578839626/1055054834266542100/image.png" width=50% height=50%> <img src="https://cdn.discordapp.com/attachments/683383222578839626/1055055074176536626/image.png" width=15% height=15%>


After adding few of the basic game components we started moving onto to working with storing player data seperately (board, ship) to get ready for creating ship objects to be used when we are doing hitscans for both players when we are shooting at each other's board, after we moved onto improving some of the game logic as it lagged behind as it wouldn't be able to support Player vs Player without it looking broken, and after that we had to introduce the text user interface to start the game and select the gamemodes.

After  this most changes were logic improvements and try finding extreme cases to figure our why certai bugs happened and redo some of the game logic like the ship placement and shooting logic has changed quite a bit compared to the start.
At the end of the development most of the updates introduced were small fixes, visual updates and certain extra features to improve the "usability" of the game.

The game was tested a lot of times before pushing any commits as some of the bugs were very hard to find and pinpoint, and since I was using Go on Windows the [Delve debugger](https://github.com/go-delve/delve) was very helpful at times.

1d. Object oriented design ideas weren't very hard to implement in this project, one of the main ways it was implemented was the way we created ship objects for each player and the ship object contained info such the status of the ship (sunk or not), size, name and position (used for hit detection).
```go
type Ship struct {
	sunk     bool
	size     int
	name     string
	position []Position
}

type Position struct {
	x, y int
}
```

The ship object was used when creating the ship, adding it's position relative to the board, updating the sunk status and used to check who won (who sunk all the ships on one side).

**Creating the ship**
```go
func (ship *Ship) create(shipSize, index int) {
	ship.size = shipSize
	ship.sunk = false
	ship.name = boardShipName[index]
	ship.position = make([]Position, shipSize)
}
```

**Adding ship position**
```go
func (ship *Ship) addCoords(xPos, yPos, index int) {
	ship.position[index] = Position{x: xPos, y: yPos}
}
```

**Checking if the whole ship was shot**
```go
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
```

**Changing ship sunk status**
```go
if ships[s].allShot() {
    ships[s].sunk = true //change the object bool to sunk if all the whole ship was shot
}
```

**Checking if Player 2 ships have all sunk or not**
```go
go func() {
    playerWon := true

    //if the enemy ships have all been shot then this should return false and not send anything to chan bool
    for s := range player2Ships {
        if !player2Ships[s].sunk {
            playerWon = false
        }
    }
    if playerWon {
        chanPlayer <- "Player 1 (HUMAN)" //sends the value via the string channel
    } else {
        chanPlayer <- "" //sends the value via the string channel
    }
}()
```

2a. The use of the expected standards of the assignments was an important concept, however since Go it's already a very easily readable language a lot of these standards were met automatically for example:
- **Use space consistently to separate operators and delimiters**
    - This was automatically met since Go automatically lints your code everytime you save it if you use a supported editor
- **Use appropriate and consistent indentation, logical grouping and spaced blocks within your codebases; adopt tabs or a set number of spaces (ideally tabs) for indenting**
    - This was automatically met since Go automatically lints your code everytime you save it if you use a supported editor
- **Conserve system resources.**
    - Since Go is a compiled language, its doesn't take much in resources in general to run, for example the compiled program is about 2MB and the memory usage is around 8MB of RAM and at it's peak uses 0.1% CPU fa(when using the bot to randomly guess) 

The other expected standards can mostly be found implemented in the code when being reviewed.

2b. Code reusability was implemented where possible and it wasn't always easy as a lot of conditions were added for them to work together, for example some code for placing ships was easily reused when using the 1 VS 1 and 1 VS BOT mode, but when doing turns this became more complicated, so making sure the Player 2 functions required checks to not display certain things or do certain things depending on the gamemode and at times they didn't work as intended for example the switch from a only a verification statement for coordinates to a for loop and a verification statement inside in case of the coordinates not working.
![Example code reusability](https://cdn.discordapp.com/attachments/683383222578839626/1055070361936478238/image.png)
![Example2 code reusability](https://cdn.discordapp.com/attachments/683383222578839626/1055071011684499457/image.png)

3a. I believe some of the implementations of the game that could have been made better would have been the menu with more options and more functionality like a leaderboard as it feels pretty empty, however I also do believe that I made most of what was possible with such a simple game in terms of code reuse as some was even grabbed from previous projects like the use of git release, goreleaser, bash scripts, env files and restarting on multiple os environments.

3b. One of the ways I implemented "advanced programming" was with my spawnCoordinates function that spawned a ship on the board and created the object with the corresponding coordinates (to be used later in hit detection).
![Read the comments](https://cdn.discordapp.com/attachments/683383222578839626/1055072970093121536/image.png)

3c. I believe in terms of algorithms there wasn't much to improve as the bot was randomly guessing where to put a shit and where to shoot the board, however I was able to showcase the use of Go Routines to speed up the winner check by around 50% of the time since the checks are simultaneous.
![Read the comments](https://cdn.discordapp.com/attachments/683383222578839626/1055076878974722088/image.png)

3d. To advanced the project 1 step further I could introduce different bot difficulties and start try to make use of AI (https://paulvanderlaken.com/2019/01/21/beating-battleships-with-algorithms-and-ai/) as it's possible to integrate it with the bot instead of doing random guesses or even start making use of smarter guesses as shown in the following article https://www.datagenetics.com/blog/december32011/.
