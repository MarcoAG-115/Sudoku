package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput(input string) [9][9]int {
	file, err := os.Open(input)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		lines = append(lines, scanner.Text())

	}

	var newlines []string
	for i := 0; i < 9; i++ {
		newlines = append(newlines, strings.ReplaceAll(lines[i], " ", ""))
	}
	//fmt.Println(newlines)

	board := [9][9]int{}

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {

			str := string([]rune(newlines[row])[col])
			i1, _ := strconv.Atoi(str)

			board[row][col] = i1
		}
	}
	//fmt.Println(board)
	return board
}

func printBoard(board [9][9]int) {
	fmt.Println("+-------+-------+-------+")
	for row := 0; row < 9; row++ {
		fmt.Print("| ")
		for col := 0; col < 9; col++ {
			if col == 3 || col == 6 {
				fmt.Print("| ")
			}
			fmt.Printf("%d ", board[row][col])
			if col == 8 {
				fmt.Print("|")
			}
		}
		if row == 2 || row == 5 || row == 8 {
			fmt.Println("\n+-------+-------+-------+")
		} else {
			fmt.Println()
		}
	}
}

func main() {
	// b, err := ioutil.ReadFile("file1.txt") // just pass the file name
	// if err != nil {
	// 	fmt.Print(err)
	// }

	// //fmt.Println(b)   // print the content as 'bytes'
	// str := string(b) // convert content to a 'string'
	// x := strings.ReplaceAll(str, " ", "")
	// y := strings.ReplaceAll(x, "\n", "")

	// fmt.Println(x)
	// //fmt.Println(len(str))

	// //v := "a"

	// if _, err := strconv.Atoi(y); err == nil {
	// 	fmt.Printf("%q looks like a number.\n", y)
	// }

	board := parseInput("file1.txt")

	printBoard(board)

	if backtrack(&board) {
		fmt.Println("The Sudoku was solved successfully:")
		printBoard(board)
	} else {
		fmt.Printf("The Sudoku can't be solved.")
	}
}

func backtrack(board *[9][9]int) bool {
	if !hasEmptyCell(board) {
		return true
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				for candidate := 9; candidate >= 1; candidate-- {
					board[i][j] = candidate
					if isBoardValid(board) {
						if backtrack(board) {
							return true
						}
						board[i][j] = 0
					} else {
						board[i][j] = 0
					}
				}
				return false
			}
		}
	}
	return false
}

func hasEmptyCell(board *[9][9]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				return true
			}
		}
	}
	return false
}

func isBoardValid(board *[9][9]int) bool {

	//check duplicates by row
	for row := 0; row < 9; row++ {
		counter := [10]int{}
		for col := 0; col < 9; col++ {
			counter[board[row][col]]++
		}
		if hasDuplicates(counter) {
			return false
		}
	}

	//check duplicates by column
	for row := 0; row < 9; row++ {
		counter := [10]int{}
		for col := 0; col < 9; col++ {
			counter[board[col][row]]++
		}
		if hasDuplicates(counter) {
			return false
		}
	}

	//check 3x3 section
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			counter := [10]int{}
			for row := i; row < i+3; row++ {
				for col := j; col < j+3; col++ {
					counter[board[row][col]]++
				}
				if hasDuplicates(counter) {
					return false
				}
			}
		}
	}

	return true
}

func hasDuplicates(counter [10]int) bool {
	for i, count := range counter {
		if i == 0 {
			continue
		}
		if count > 1 {
			return true
		}
	}
	return false
}
