async function generateNewPuzzle() {
  const response = await fetch("/api/generate");
  const data = await response.json();
  console.log(data); // For debugging

  // Populate the Sudoku board with the generated puzzle
  const board = data.board;
  const sudokuBoard = document.getElementById("sudoku-board");
  sudokuBoard.innerHTML = ""; // Clear previous board

  for (let i = 0; i < 9; i++) {
    for (let j = 0; j < 9; j++) {
      const cell = document.createElement("input");
      cell.type = "text";
      cell.className = "sudoku-cell";
      cell.value = board[i][j] !== 0 ? board[i][j] : "";
      sudokuBoard.appendChild(cell);
    }
  }
}

async function checkSolution() {
  const board = getBoardState();
  const response = await fetch("/api/check", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ board }),
  });

  const isCorrect = await response.json();
  alert(isCorrect ? "Solution is correct!" : "Solution is incorrect!");
}

function getBoardState() {
  const board = [];
  const cells = document.getElementsByClassName("sudoku-cell");

  for (let i = 0; i < 9; i++) {
    const row = [];
    for (let j = 0; j < 9; j++) {
      const value = cells[i * 9 + j].value.trim();
      row.push(value === "" ? 0 : parseInt(value));
    }
    board.push(row);
  }

  return board;
}
