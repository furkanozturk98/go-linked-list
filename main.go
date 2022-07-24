package main

import (
	list "linkedList/linkedList"
)

func main() {
	list := list.List{}
	list.AddToTail(5)
	list.AddToTail(4)
	list.AddAfter(1, 6)

	list.DeleteByIndex(1)

	list.Display()

	/* list.ShowBackwards() */

	/* list.Reverse() */

	/* fmt.Println("\n==============================\n")
	fmt.Printf("Head: %v\n", link.head.key)
	fmt.Printf("Tail: %v\n", link.tail.key)
	link.Display()
	fmt.Println("\n==============================\n")
	fmt.Printf("head: %v\n", link.head.key)
	fmt.Printf("tail: %v\n", link.tail.key)
	link.Reverse()
	fmt.Println("\n==============================\n") */
}
