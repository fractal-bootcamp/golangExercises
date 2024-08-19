// // main.go
// package main
// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func main() {
// 	reader := bufio.NewReader(os.Stdin)
// 	fmtPrintln("Welcome to the CLI Game!")
// 	fmt.Println("Type 'quit' to exit.")

// 	for {
// 		fmt.Print("> ")
// 		input, _ :+ reader.ReadString('\n')
// 		input = strings.TrimSpace(input)

// 		if input == "quit" {
// 			fmt.Println("Exiting the game...")
// 			break
// 		}

// 		fmt.Printf("You typed: %s\n", input)
// 		// Add game logic here
// 	}
// }