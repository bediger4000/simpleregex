package main

import (
	"fmt"
	"os"
)

func main() {
	regex := os.Args[1]
	str := os.Args[2]

	if match([]rune(regex), []rune(str)) {
		fmt.Println("match")
		return
	}
	fmt.Println("no match")
}
