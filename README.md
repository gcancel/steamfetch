[![ci](https://github.com/gcancel/steamfetch/actions/workflows/ci.yml/badge.svg)](https://github.com/gcancel/steamfetch/actions/workflows/ci.yml)
# steamfetch
**An easy to use CLI utility to fetch and display your steam information.**

![](https://github.com/gcancel/steamfetch/blob/main/img/steam_output.png)
## How to Run:
1. Install [Go](https://go.dev/dl/)
2. Clone repository.
3. Run `go mod tidy` from the cloned repository's directory.
4. Build the executable and run by executing:
    #### Linux:  
    ```bash
        sudo go build -o /usr/bin steamfetch && steamfetch update
    ```
    #### Windows:
    ```powershell
        go build -o <directory_path>/steamfetch && steamfetch udpate
    ```
    ***(you can also create `$PATH` environmental variables if you want to run directly from terminal)***
6. You will be prompted to enter your steamid and your [Steam Web API Key](https://steamcommunity.com/dev/apikey).
   - you can find your steamid under: `Settings -> Account Details` in steam.
7. An initial update should run to pull down your steam game data and you are ready to go!

## Features:
- [x] Display total steam games. 
- [x] Display steam game time across all games.
- [x] Display total steam game time in two week period.
- [x] Display most played games (top five)

### Coming soon:
- [ ] Display total number of installed games.
- [ ] Display installed game disk usage.

## How to use:
- Run `./steamfetch` locally from the repository directory or `steamfetch` from the terminal if you setup your `$PATH` environmental variables, or on Linux, moved the executable to /usr/bin.
- `--help` displays the help message.
- `update` parameter pulls down your steam data and updates the local database. You can use the `-f` or `--force` flag to clear the database table and pull the data again.
- `--mostplayed n` flag can be used to set the number of returned games for the "Most played games" section
- `backlog` parameter will return a comma delimited list of all games you have not played (0 minutes). In the case you want to pipe this to a text file.
