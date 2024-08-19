// // save_game.go
// package main

// import (
// 	"os"
// )

// func saveGameState(filename, data string) error {
//     file, err := os.Create(filename)
//     if err != nil {
//         return err
//     }
//     defer file.Close()

//     _, err = file.WriteString(data)
//     return err
// }

// func loadGameState(filename string) (string, error) {
//     file, err := os.Open(filename)
//     if err != nil {
//         return "", err
//     }
//     defer file.Close()

//     data := make([]byte, 100)
//     n, err := file.Read(data)
//     if err != nil {
//         return "", err
//     }
//     return string(data[:n]), nil
// }
