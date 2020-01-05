package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	path, _ := os.Hostname()

	fmt.Println(path)
	file, err := os.Open("/Users/mac/Downloads/README.md")

	if err != nil {
		fmt.Println(err)

	}
	defer file.Close()
	line := bufio.NewReader(file)
	readline, _, _ := line.ReadLine()
	fmt.Printf(string(readline))

}
