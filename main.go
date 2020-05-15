package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

//Constants
const boardSize = 171

//Validating the input file.
//This function uses multiple tests to check if
//the provided file contents represent a
//solvable Sudoku board.
func validator(input string) bool {

	//Variables
	rowCount := 0
	testPassed := 0
	isValid := false
	var line []string
	lineCount := 0
	test3 := 0

	//Test 1
	//Checks if the file contains the required number
	//of characters needed for a 9x9 board. This is
	//done by finding the size of the file contents.

	//Opens file and checks for errors.
	b, err := ioutil.ReadFile(input)
	if err != nil {
		fmt.Println(err)
	}

	//Converts file contents from bytes to a string.
	str := string(b)

	//Prints error message if the string containing
	//the file contents is not equal to 171.
	if len(str) != boardSize {
		fmt.Println("Error: The file contents are not valid. Please use a file with" +
			"81 integers. (Blanks should be represented with zeros)")

	} else {
		testPassed++
	}

	//Test 2
	//Checks if the board is organized in a 9x9
	//configuration. This is done by counting
	//the number of lines that contain the file's
	//contents.

	//Opens file and checks for errors.
	file, err := os.Open(input)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	//For loop reads file line by line.
	for scanner.Scan() {
		//Adds a line from the file to a string array.
		line = append(line, scanner.Text())

		//Test 3
		//Each row is also analyzed to
		//ensure that the board only contains integers
		//0 through 9.
		str := strings.ReplaceAll(line[lineCount], " ", "")
		//Checks if the given string contains an integer.
		if _, err := strconv.Atoi(str); err == nil {
			test3++
		}

		lineCount++
		rowCount++
	}

	//Prints error message if the board does not have
	//nine rows.
	if rowCount == 9 {
		testPassed++
	} else {
		fmt.Println("Error: The file contents are not valid. Please organize file contents" +
			"into a 9x9 board with spaces. (Blanks should be represented with zeros)")
	}

	//Prints error message if the board contains
	//a non-integer.
	if test3 == 9 {
		testPassed++
	} else {
		fmt.Println("Error: The file contents are not valid. Please submit a file that only contains" +
			"the integers 0 through 9.")
	}

	//The board is determined to be solvable if
	//it passes all the tests.
	if testPassed == 3 {
		isValid = true
	}

	return isValid
}

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

	board := [9][9]int{}

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {

			str := string([]rune(newlines[row])[col])
			i1, _ := strconv.Atoi(str)

			board[row][col] = i1
		}
	}

	return board
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

func main() {

	//Variables
	var fileName string

	//Initial Output
	//Prompts user for file name.
	fmt.Printf("****************Marco's Sudoku Machine****************\n")
	fmt.Println("Enter name of Sudoku file: ")
	fmt.Scanln(&fileName)
	fmt.Println("Reading", fileName)

	//Reading File
	//Reads file with given name &
	//checks if file was opened w/o issues.
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(data))

	//Validating File
	//The validator function is used to
	//check if the file contents contain
	//a solvable Sudoku board.
	fmt.Println("Checking if file contains a valid board...")

	if validator(fileName) == true {
		fmt.Println("The file contains a valid board.")

		board := parseInput(fileName)

		k := 0

		if backtrack(&board) {
			fmt.Println("The Sudoku was solved successfully:")
			for row := 0; row < 9; row++ {
				for col := 0; col < 9; col++ {

					k++
					fmt.Print(board[row][col], " ")
					if k == 9 {
						k = 0
						fmt.Println("")
					}
				}
			}
		} else {
			fmt.Printf("The Sudoku can't be solved.")
		}
	}
}
