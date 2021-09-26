package main

import (
	"log"
	"os"
)

// var (
// 	newFile *os.File
// 	err     error
// )

// func main() {
// 	newFile, err = os.Create("test.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Println(newFile)
// 	newFile.Close()
// }

func main() {
	oriPath := "*.exe"
	newPath := "/exe/*.exe"
	err := os.Rename(oriPath, newPath)
	if err != nil {
		log.Fatal(err)
	}
}
