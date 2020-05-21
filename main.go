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
const boardSize = 81
const numOfLineSpaces = 8

//File reader.
//This function uses the user's inputted file
//name to read a file using ioutil.ReadFile.
func fileReader(fileName string) []byte {

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}

	return data
}

//Opening file.
//This function opens the provided file
//with the user's input and os.Open.
//The file's contents are returned.
func fileOpener(fileName string) *os.File {

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	return file
}

//Prints file contents.
//This function reads the file line by line
//and prints each line with spaces between
//each character.
func printer(file *os.File) {

	var lines []string
	lineCount := 0
	count := 0
	//Resets scanner position to the beginning of the file.
	file.Seek(0, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())

		//Creating and formatting board with file contents.
		str := strings.ReplaceAll(lines[lineCount], "", "   ")
		if count == 0 || count == 3 || count == 6 {

			fmt.Println(" * - - - - - * - - - - - * - - - - - *")
		} else {

			fmt.Println(" |           |           |           |")
		}

		x := []rune(str)
		r := '|'
		x[1] = r
		x[13] = r
		x[25] = r
		x[37] = r
		fmt.Println(string(x))

		lineCount++
		count++
	}
	if count == 9 {
		fmt.Println(" * - - - - - * - - - - - * - - - - - *")
	}
}

//Prints solution.
//This function prints the solution
//by generating and formatting a board.
func printBoard(board [9][9]int) {

	count := 0
	k := 0
	fmt.Println(" * - - - - - * - - - - - * - - - - - *")
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {

			k++
			if col == 8 {
				fmt.Print(board[row][col], " |")
			} else if col == 0 {
				fmt.Print(" | ", board[row][col], "   ")
			} else if col == 2 || col == 5 {
				fmt.Print(board[row][col], " | ")
			} else {
				fmt.Print(board[row][col], "   ")
			}
			if k == 9 {
				k = 0
				fmt.Println(" ")
				count++
				if count != 9 {
					if count == 3 || count == 6 {

						fmt.Println(" * - - - - - * - - - - - * - - - - - *")
					} else {

						fmt.Println(" |           |           |           |")
					}
				}
			}
		}
	}
	fmt.Println(" * - - - - - * - - - - - * - - - - - *")
}

//Validating the input file.
//This function uses multiple tests to check if
//the provided file contents represent a
//solvable Sudoku board.
func validator(input string, data []byte, file *os.File) bool {

	//Variables
	rowCount := 0
	testPassed := 0
	isValid := false
	var line []string
	lineCount := 0
	test1 := false
	test3 := 0
	check := true

	//Test 1
	//Checks if the file contains the required number
	//of characters needed for a 9x9 board. This is
	//done by finding the size of the file contents.

	//Converts file contents from bytes to a string.
	str := string(data)

	//Prints error message if the string containing
	//the file contents is not equal to 171.
	if len(str)-numOfLineSpaces != boardSize {
		test1 = true

	} else {
		testPassed++
	}

	//Test 2
	//Checks if the board is organized in a 9x9
	//configuration. This is done by counting
	//the number of lines that contain the file's
	//contents.

	//Resets scanner position to the beginning of the file.
	file.Seek(0, 0)
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

	//Prints an error message if all the tests fail.
	if test1 == true && rowCount != 9 && test3 != 9 {
		fmt.Println("Error: The file contents are not valid. Please organize file contents" +
			" into a 9x9 board with no spaces. (Blanks should be represented with zeros)")
		check = false
	}

	//Prints error message if the board does not have
	//nine rows.
	if rowCount == 9 {
		testPassed++
	}

	//Prints error message if the board contains
	//a non-integer.
	if test3 == 9 {
		testPassed++
	} else if check == true && test3 != 9 {
		fmt.Println("Error: The file contents are not valid. Please submit a file that only contains" +
			" the integers 0 through 9.")
	}

	//Prints error message if the board does not
	//contain 81 integers.
	if test1 == true && check == true {
		fmt.Println("Error: The file contents are not valid. Please use a file with" +
			" 81 integers. (Blanks should be represented with zeros)")
	}

	//The board is determined to be solvable if
	//it passes all the tests.
	if testPassed == 3 {
		isValid = true
	}

	return isValid
}

func parseInput(input string, file *os.File) [9][9]int {

	var lines []string
	//Resets scanner position to the beginning of the file.
	file.Seek(0, 0)
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
	fmt.Println("\n************************************************")
	fmt.Println("*            Marco's Sudoku Machine            *")
	fmt.Println("************************************************")
	fmt.Printf("\nEnter name of Sudoku file: ")
	fmt.Scanln(&fileName)
	fmt.Println("\nReading", fileName, "...")
	fmt.Println("")

	//Reading File
	//Reads file with given name &
	//checks if file was opened w/o issues.
	data := fileReader(fileName)
	file := fileOpener(fileName)

	//Prints formatted file contents.
	printer(file)

	//Validating File
	//The validator function is used to
	//check if the file contents contain
	//a solvable Sudoku board.

	if validator(fileName, data, file) == true {
		fmt.Println("\nChecking if file contains a valid board...")
		fmt.Println("The file contains a valid board.")

		board := parseInput(fileName, file)

		if backtrack(&board) {
			fmt.Println("The Sudoku was solved successfully:")
			fmt.Println("")
			printBoard(board)

		} else {
			fmt.Printf("The Sudoku can't be solved.")
		}
	}

	//Tests
	//Uncomment to run tests.
	//testing()
}

//Testing multiple scenarios.
//This function performs tests for the
//Sudoku solver by comparing the outputs
//generated by the different test files.
func testing() {

	fmt.Println("\n********************************\n*     Running Testing Mode     *\n********************************")

	/**Test 1**/
	//Checks if program returns an error if a file
	//has spaces between each character.

	//Case 1
	fmt.Println("\n***Test 1***\nCase 1: File contains more or less than 81 characters.")
	if false == validator("test1.txt", fileReader("test1.txt"), fileOpener("test1.txt")) {
		fmt.Println("TEST PASSED")
	}

	//Case 2
	fmt.Println("\nCase 2: File contains 81 characters.")
	if true == validator("file1.txt", fileReader("file1.txt"), fileOpener("file1.txt")) {
		fmt.Println("TEST PASSED")
	}

	/**Test 2**/
	//Checks if program returns an error if a
	//non-existing file is called.

	//Case 1
	fmt.Println("\n***Test 2***\nCase 1: The provided file does not exist.")
	if string(fileReader("nofile.txt")) == "" {
		fmt.Println("TEST PASSED")
	}

	//Case 2
	fmt.Println("\nCase 2: The provided file exists.")
	if string(fileReader("file1.txt")) != "" {
		fmt.Println("TEST PASSED")
	}

	/**Test 3**/
	//Checks if program returns an error if a file's
	//contents are not divided into 9 lines.

	//Case 1
	fmt.Println("\n***Test 3***\nCase 1: File contains less than 9 lines.")
	if false == validator("test2.txt", fileReader("test2.txt"), fileOpener("test2.txt")) {
		fmt.Println("TEST PASSED")
	}

	//Case 2
	fmt.Println("\nCase 2: File contains 9 lines.")
	if true == validator("file1.txt", fileReader("file1.txt"), fileOpener("file1.txt")) {
		fmt.Println("TEST PASSED")
	}

	/**Test 4**/
	//Checks if program returns an error if the
	//file provided contains non-integers.

	//Case 1
	fmt.Println("\n***Test 4***\nCase 1: File contains non-integers.")
	if false == validator("test3.txt", fileReader("test3.txt"), fileOpener("test3.txt")) {
		fmt.Println("TEST PASSED")
	}

	//Case 2
	fmt.Println("\nCase 2: File contains only integers.")
	if true == validator("file1.txt", fileReader("file1.txt"), fileOpener("file1.txt")) {
		fmt.Println("TEST PASSED")
	}

	/**Test 5**/
	//The program is given an easy Sudoku puzzle
	//to solve.

	//Easy Difficulty
	fmt.Println("\n***Test 5***\nEasy Difficulty:")
	data := fileReader("easyTest.txt")
	file := fileOpener("easyTest.txt")

	if validator("easyTest.txt", data, file) == true {

		board := parseInput("easyTest.txt", file)

		if backtrack(&board) {
			fmt.Println("\nGenerated Solution:")
			fmt.Println("")
			printBoard(board)

		} else {
			fmt.Printf("The Sudoku can't be solved.")
		}
	}
	fmt.Println("\nSolved board from solution file:")
	fmt.Println("")
	file2 := fileOpener("easySolution.txt")
	printer(file2)

	/**Test 6**/
	//The program is given an average Sudoku puzzle
	//to solve.

	//Medium Difficulty
	fmt.Println("\n***Test 6***\nMedium Difficulty:")
	data2 := fileReader("mediumTest.txt")
	file3 := fileOpener("mediumTest.txt")

	if validator("mediumTest.txt", data2, file3) == true {

		board := parseInput("mediumTest.txt", file3)

		if backtrack(&board) {
			fmt.Println("\nGenerated Solution:")
			fmt.Println("")
			printBoard(board)

		} else {
			fmt.Printf("The Sudoku can't be solved.")
		}
	}
	fmt.Println("\nSolved board from solution file:")
	fmt.Println("")
	file4 := fileOpener("mediumSolution.txt")
	printer(file4)

	/**Test 7**/
	//The program is given a difficult Sudoku puzzle
	//to solve.

	//Hard Difficulty
	fmt.Println("\n***Test 7***\nHard Difficulty:")
	data3 := fileReader("hardTest.txt")
	file5 := fileOpener("hardTest.txt")

	if validator("hardTest.txt", data3, file5) == true {

		board := parseInput("hardTest.txt", file5)

		if backtrack(&board) {
			fmt.Println("\nGenerated Solution:")
			fmt.Println("")
			printBoard(board)

		} else {
			fmt.Printf("The Sudoku can't be solved.")
		}
	}
	fmt.Println("\nSolved board from solution file:")
	fmt.Println("")
	file6 := fileOpener("hardSolution.txt")
	printer(file6)
}
