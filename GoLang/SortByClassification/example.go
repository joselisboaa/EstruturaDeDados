package main

import (
	"fmt"
)

func main() {
	list := []int{13, 2, 3, 4, 5, 10, 8, 9}
	sortedList := SortingBySelection(list)

	fmt.Println("Sorted List: ", sortedList)
}

func removeElement(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

// SearchMinor O(n)
func SearchMinor(list []int) int {
	lowestNumber := list[0]
	lowestIndex := 0

	for i := 0; i < len(list); i++ {
		if list[i] < lowestNumber {
			lowestNumber = list[i]
			lowestIndex = i
		}
	}

	return lowestIndex
}

// SortingBySelection O(n²)
func SortingBySelection(list []int) []int {
	var auxArray []int

	for i := 0; i < len(list)+len(auxArray); i++ {
		lowest := SearchMinor(list)

		auxArray = append(auxArray, list[lowest])

		list = removeElement(list, lowest)
	}

	return auxArray
}
