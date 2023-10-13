package main

import "fmt"

type Node struct {
	info interface{}
	next *Node
}

type List struct {
	head *Node
}

func main() {
	sl := List{}
	for i := 0; i < 5; i++ {
		sl.Insert(1 + (i * 5))
	}

	sl.DeleteFirst()
	sl.DeleteLast()
	Show(&sl)
}

func (listPointer *List) Insert(d interface{}) {
	list := &Node{info: d, next: nil}

	if listPointer.head == nil {
		listPointer.head = list
	} else {
		pointer := listPointer.head

		for pointer.next != nil {
			pointer = pointer.next
		}

		pointer.next = list
	}
}

func Show(listPointer *List) {
	pointer := listPointer.head

	for pointer != nil {
		fmt.Printf("-> %v ", pointer.info)
		pointer = pointer.next
	}
}

// DeleteFirst O(1)
func (listPointer *List) DeleteFirst() {
	listPointer.head = listPointer.head.next
}

// DeleteLast O(n)
func (listPointer *List) DeleteLast() {
	pointer := listPointer.head

	for pointer != nil {
		if pointer.next == nil {
			pointer.next = nil

			return
		}

		pointer = pointer.next
	}
}

// Deletion O(n)
func (listPointer *List) Deletion(target int) {
	pointer := listPointer.head
	list := &Node{info: pointer.info, next: nil}

	for pointer != nil {
		fmt.Println(pointer.info)
		if pointer.next.info == target {
			listPointer.head = listPointer.head.next.next
		}

		if pointer.info == target {
			listPointer.head = list
		}

		pointer = pointer.next
	}
}
