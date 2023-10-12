package main

import "fmt"

func main() {
	numbers := []int{1, 3, 5, 7, 9}
	arrayAddress, numberOfAttempts := binarySearch(numbers, 10)

	fmt.Print("Array index: ", arrayAddress, "\nNumber of attempts: ", numberOfAttempts)
}

//Retorna a posição do elemento alvo ou nulo
func binarySearch(array []int, target int) (int, int) {
	smallestArrayAddress := 0
	numberOfAttempts := 0
	biggerArrayAddress := len(array) - 1

	for smallestArrayAddress <= biggerArrayAddress {
		midArrayAddress := (smallestArrayAddress + biggerArrayAddress) / 2
		guess := array[midArrayAddress]
		numberOfAttempts += 1

		if guess == target {
			return midArrayAddress, numberOfAttempts
		}

		if target > guess {
			smallestArrayAddress = midArrayAddress + 1
		}

		if target < guess {
			biggerArrayAddress = midArrayAddress - 1
		}
	}
	return -1, numberOfAttempts
}
