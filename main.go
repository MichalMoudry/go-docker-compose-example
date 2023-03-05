package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var userInput string
	for strings.ToLower(userInput) != "exit" {
		for key, variable := range os.Environ() {
			fmt.Println(key, "=>", variable)
		}
		fmt.Scan(&userInput)
	}
}
