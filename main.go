package main

import (
	"fmt"
	"io/ioutil"
)

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
	//Reads file with given name, checks
	//if file was opened w/o issues
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(data))
}
