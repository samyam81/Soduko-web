# Sudoku Web Application

Welcome to the Sudoku web application repository! This project provides a simple yet interactive Sudoku game implemented using Go for the backend and HTML, CSS, and JavaScript for the frontend.

## Author

- Original Author: [Samyam](https://github.com/samyam81)

## Features

- **Sudoku Generation**: Automatically generates a new Sudoku puzzle on demand.
- **Solution Checking**: Allows users to check their solutions against the generated puzzle.
- **Responsive Design**: Supports responsive design for optimal viewing on different screen sizes.

## Technologies Used

- **Backend**: Go (Golang) with `net/http` for handling HTTP requests and `github.com/gorilla/mux` for routing.
- **Frontend**: HTML, CSS (styles defined in `styles.css`), and JavaScript (`sudoku.js`).
- **Data Exchange**: JSON for communicating between the frontend and backend.

## Setup Instructions

To run the Sudoku web application locally, follow these steps:

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/samyam81/Soduko-web.git
   cd Soduko-web
   ```

2. **Install Dependencies**:
   ```bash
   go mod tidy
   ```

3. **Run the Application**:
   ```bash
   go run main.go
   ```

   The server will start on `http://localhost:8080`.

4. **Access the Sudoku Game**:
   Open your web browser and navigate to `http://localhost:8080/game/sudoku.html` to play the Sudoku game.

## File Structure

The repository structure is as follows:

```
Soduko-web/
│
├── main.go          # Backend server implementation
└── static/
    ├── sudoku.html  # HTML file for the Sudoku game interface
    ├── sudoku.js    # JavaScript file for client-side logic
    └── styles.css   # CSS file for styling the Sudoku board and buttons
```

## Usage

- **Generating a New Puzzle**:
  Click on the `Generate New Puzzle` button to fetch and display a new Sudoku puzzle on the board.

- **Checking Solution**:
  Fill in your solution in the Sudoku cells and click `Check Solution` to verify if it matches the generated puzzle's solution.

## Deployment

To deploy the Sudoku web application:
- Ensure that your Go environment is set up on your deployment server.
- Copy the repository files to your server.
- Build and run the Go application (`go build main.go` and `./main`).
- Access the Sudoku game from your server's IP or domain.

## Future Enhancements

- **User Authentication**: Implement user accounts to track progress and scores.
- **Multiple Difficulty Levels**: Add options for generating puzzles of varying difficulty.
- **Leaderboard**: Include a leaderboard to display top scores and completion times.

Feel free to contribute to this repository by adding features, fixing bugs, or improving the documentation. Enjoy playing Sudoku!
