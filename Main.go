package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// SudokuBoard represents a Sudoku puzzle board
type SudokuBoard struct {
	Board     [][]int `json:"board"`
	Solution  [][]int `json:"solution"`
	Completed bool    `json:"completed"`
}

var currentBoard SudokuBoard

// Generate a new Sudoku puzzle
func generateSudoku(w http.ResponseWriter, r *http.Request) {
	currentBoard = generateNewSudoku()
	json.NewEncoder(w).Encode(currentBoard)
}

// Check if the provided solution is correct
func checkSolution(w http.ResponseWriter, r *http.Request) {
	var userSolution SudokuBoard
	_ = json.NewDecoder(r.Body).Decode(&userSolution)

	if compareBoards(currentBoard.Solution, userSolution.Board) {
		json.NewEncoder(w).Encode(true)
	} else {
		json.NewEncoder(w).Encode(false)
	}
}

// Helper function to generate a new Sudoku puzzle
func generateNewSudoku() SudokuBoard {
	// Implement your Sudoku generator logic here
	// For simplicity, a mock implementation is shown
	board := [][]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}

	solution := [][]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
	}

	return SudokuBoard{
		Board:     board,
		Solution:  solution,
		Completed: false,
	}
}

// Helper function to compare two Sudoku boards
func compareBoards(board1, board2 [][]int) bool {
	// Implement board comparison logic here
	// Compare corresponding cells of two boards
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board1[i][j] != board2[i][j] {
				return false
			}
		}
	}
	return true
}

func main() {
	rand.Seed(time.Now().UnixNano())

	router := mux.NewRouter()

	// Serve static files from the static directory
	router.PathPrefix("/game/").Handler(http.StripPrefix("/game/", http.FileServer(http.Dir("./static/"))))

	// Define API endpoints
	router.HandleFunc("/api/generate", generateSudoku).Methods("GET")
	router.HandleFunc("/api/check", checkSolution).Methods("POST")

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", router))
}

