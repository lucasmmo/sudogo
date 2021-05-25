package sudogo

import (
	"fmt"
	"math/rand"
)

func Generate(dificult string) {
	// the dificults can be easy medium hard
	// default is easy
	switch dificult {
	case "easy":
		GenerateWithNumbs(50)
	case "medium":
		GenerateWithNumbs(25)
	case "hard":
		GenerateWithNumbs(12)
	default:
		GenerateWithNumbs(50)
	}
}

func GenerateWithNumbs(nums int) {
	// nums is how much numbers will print to help to solve it
	var board = [9][9]int{}
	// printBoard(board)

	for i := 0; i < nums; i++ {
		row := rand.Intn(9)
		col := rand.Intn(9)
		num := rand.Intn(9) + 1
		for !checkIsValid(board, row, col, num) {
			row = rand.Intn(9)
			col = rand.Intn(9)
			num = rand.Intn(9) + 1
		}
		board[row][col] = num
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 0 {
				fmt.Printf("o indice [%v][%v] é = a 0 \t \n ", i, j)
			}
		}
	}
	fmt.Println()
	printBoard(board)

}

func printBoard(board [9][9]int) {
	TableTB := "|--------------------------------|"
	TableMD := "|----------+----------+----------|"

	fmt.Println(TableTB)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if (i == 3 || i == 6) && j == 0 {
				// first section
				fmt.Println(TableMD)
			}
			if j == 0 || j == 3 || j == 6 {
				fmt.Printf("| ")
			}
			fmt.Printf(" %v ", board[i][j])
			if j == 8 {
				fmt.Printf("| \n")
			}
		}
	}
	fmt.Println(TableTB)
}

func checkIsValid(board [9][9]int, row int, column int, num int) bool {
	valid := true
	for i := 0; i < 9; i++ {
		// check column numbers
		if board[i][column] == num {
			valid = false
		}
	}
	for i := 0; i < 9; i++ {
		// check row numbers
		if board[row][i] == num {
			valid = false
		}
	}
	// check section number
	rowsection := row / 3
	colsection := column / 3

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[rowsection*3+i][colsection*3+j] == num {
				valid = false
			}
		}
	}
	return valid
}
