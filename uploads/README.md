# Number-Guessing-game


## Overview
This is a Command-Line Interface (CLI) based number guessing game implemented in Go. The game allows users to guess a randomly selected number based on different difficulty levels. The game progress and high scores are stored in a JSON file.

[.](https://roadmap.sh/projects/number-guessing-game)

## Features
- Interactive CLI-based gameplay.
- Random number generation between 1 and 100.
- Three difficulty levels:
  - **Easy**: More attempts to guess the number.
  - **Medium**: Moderate attempts.
  - **Hard**: Limited attempts.
- User input validation for guessing the number.
- Hints provided if the guess is incorrect.
- Tracks number of attempts and displays results.
- Saves game history and high scores to a JSON file.
- Option to play multiple rounds.

## How to Play
1. Run the game from the command line.
2. The game will display a welcome message and rules.
3. Select a difficulty level: Easy, Medium, or Hard.
4. Enter your guess for the randomly generated number.
5. The game will guide you by indicating whether the correct number is higher or lower.
6. Continue guessing until you either find the correct number or run out of attempts.
7. The game will display your results and ask if you want to play again.
8. Your attempts and high scores will be saved in a JSON file.

## Installation & Setup
### Prerequisites
- Go installed on your system ([Download Go](https://go.dev/dl/)).

### Steps to Run
1. Clone the repository:
   ```sh
   git clone <repository-url>
   cd number-guessing-game
   ```
2. Build and run the game:
   ```sh
   go run main.go
   ```
3. To compile and create an executable:
   ```sh
   go build -o guess_game
   ./guess_game
   ```

## JSON File Storage
- The game stores high scores and previous results in a `result.json` file.
- The file maintains a record of user attempts and best performances.

## Example JSON Structure
```json
"Radhi": [
  {
   "success": true,
   "attempt": 1,
   "level": 3
  }
]
```



