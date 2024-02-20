package main

import (
	"fmt"
	"math/rand"
)

func main() {
	msgNum := randInt(1, 6)
	fmt.Printf("Generated a %d\n", msgNum)
	printMessage(msgNum)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func printMessage(msgNum int) {
	if msgNum >= 1 && msgNum <= 5 {
		fmt.Printf("Printing this for a %d\n", msgNum)
	} else {
		fmt.Println("Whoops, didn't expect this")
	}
}
