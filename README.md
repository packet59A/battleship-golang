## Project Name: battleship-golang

#### Game Name: Battleship
#### Game Objective: Try and sink the enemy ships before they shink you ships.

##### Game Rules:
- [🚢] There is a list of ships with different sizes (2 Destroyer, 3 Submarine, 3 Cruiser, 4 Battleship, 5 Carrier)
- [&puncsp;📍&puncsp;] Ships for both sides have to be placed before shooting them
- [🔫] You can shot a location once since ships do not move
- [💘] If a shot hit a ship then it will marked on the board as a hit
- [❤️‍] If a shot missed a ship then it will marked on the board as a miss
- [&puncsp;🫠&puncsp;] If every part of a ship is hit then it will be all marked as sinked
- [🏆] The first side that sinks the entire fleet of 5 ships will win the game

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
![Successful build](https://cdn.discordapp.com/attachments/683383222578839626/1054981881386508308/image.png)

