package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
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

	//Test 1
	//Checks if the file contains the required number
	//of characters needed for a 9x9 board. This is
	//done by finding the size of the file contents.
	b, err := ioutil.ReadFile(input)
	if err != nil {
		fmt.Println(err)
	}
	str := string(b)
	if len(str) != boardSize {
		fmt.Println("Error: The file contents are not valid. Please use a file with" +
			"81 integers. (Blanks should be represented with zeros)")
		testPassed++
	}

	//Test 2
	//Checks if the board is organized in a 9x9
	//configuration. This is done by counting
	//the number of lines that contain the file's
	//contents.
	file, err := os.Open(input)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		rowCount++
	}

	if rowCount == 9 {
		testPassed++
	} else {
		fmt.Println("Error: The file contents are not valid. Please organize file contents" +
			"into a 9x9 board with spaces. (Blanks should be represented with zeros)")
	}

	if testPassed == 2 {
		isValid = true
	}

	return isValid
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

	if validator("file1.txt") == true {
		fmt.Println("The file contains a valid board.")
	}
}
