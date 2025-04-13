package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("./test/wibble.txt")
	if err != nil {
		fmt.Println("There is no wibble so going to wobble")
	}

	fi, err := os.Create("./test/wobble.txt")
	if err != nil {
		fmt.Println("Creating the file failed.")
		panic(err)
	}
	_, err = fi.Write([]byte("There is no Dana, only ZUUL"))
	if err != nil {
		fmt.Println("Did not expect an issue writing to the file.")
		panic(err)
	}
	fi.Close()

	wobble, err := os.ReadFile("./test/wobble.txt")
	if err != nil {
		fmt.Println("Expected wobble.txt to exist")
		panic(err)
	}
	fmt.Printf("From the file: %s", string(wobble))
}
