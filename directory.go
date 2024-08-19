// // directory.go

// package main

// import (
// 	"fmt"
// 	"os"
// )

// func createDir(dirName string) error {
// 	return os.Mkdir(dirName, 0755)
// }

// func listFiles(dirName string) ([]string, error) {
// 	files, err := os.ReadDir(dirName)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var fileNames []string
// 	for _, file := range files {
// 		fileNames = append(fileNames, file.Name())
// 	}
// 	return fileNames, nil
// }


// func main() {
// 	err := createDir("game_data")
// 	if err != nil {
// 		fmt.Println("Error creating directory:", err)
// 		return
// 	}

// 	files, err := listFiles("game_data")
// 	if err != nil {
// 		fmt.Println("Error listing files:", err)

// 		return
// 	}
// 	fmt.Println("Files in 'game_data':", files)
// }