package main

import "fmt"

func main() {
	descRecursiveCounter(4)
	fmt.Println("\n")
	ascRecursiveCounter(0, 10)
	fmt.Println("\n")
	ascAndDescRecursiveCounter(0)
}

func descRecursiveCounter(number int) {
	initialNum := 0

	fmt.Println(number)

	if initialNum != number {
		descRecursiveCounter(number - 1)
	}
}

func ascRecursiveCounter(initialNum int, targetNum int) {
	fmt.Println(initialNum)

	if targetNum != initialNum {
		ascRecursiveCounter(initialNum+1, targetNum)
	}
}

func ascAndDescRecursiveCounter(initialNum int) {
	fmt.Println(initialNum)

	if initialNum < 10 {
		ascAndDescRecursiveCounter(initialNum+1)
	} 
	
	fmt.Println(initialNum)
}