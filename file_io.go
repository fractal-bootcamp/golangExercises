// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {

// 	inputFile, err := os.Open("input.txt")
// 	if err != nil {
// 		fmt.Println("Error opening file:", err)
// 		return
// 	}

// 	defer inputFile.Close()

// 	outputFile, err := os.Create("output.txt")
// 	if err != nil {
// 		fmt.Println("Error creating file:", err)
// 		return
// 	}
// defer outputFile.Close()

// scanner := bufio.NewScanner(inputFile)
// for scanner.Scan() {
// 	_, err := outputFile.WriteString(scanner.Text() + "\n")
// 	if err != nil {
// 		fmt.Println("error writing to file:", err)
// 		return
// 	}

// 	if err := scanner.Err(); err !=nil {
// 		fmt.Println("Error reading file:", err)
// 	}
// }
// }