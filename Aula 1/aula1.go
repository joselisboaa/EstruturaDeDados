package main

import "fmt"

func main() {
	descRecursiveCounter(4)
	fmt.Println("\n")
	ascRecursiveCounter(0)
	fmt.Println("\n")
	ascAndDescRecursiveCounter(0)
}

func descRecursiveCounter(initialNumber int) {
	fmt.Println(initialNumber)

	if initialNumber > 0 {
		descRecursiveCounter(initialNumber - 1)
	}
}

func ascRecursiveCounter(initialNum int) {
	fmt.Println(initialNum)

	if initialNum < 4 {
		ascRecursiveCounter(initialNum + 1)
	}
}

func ascAndDescRecursiveCounter(initialNum int) {
	fmt.Println(initialNum)

	if initialNum < 10 {
		ascAndDescRecursiveCounter(initialNum + 1)
	}

	fmt.Println(initialNum)
}
