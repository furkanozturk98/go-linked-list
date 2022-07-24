package list

import "fmt"

type Node struct {
	prev *Node
	next *Node
	data int
}

type List struct {
	head *Node
	tail *Node
}

func (L *List) AddToHead(data int) {
	node := &Node{
		next: L.head,
		data: data,
	}

	if L.head != nil {
		L.head.prev = node
	}

	L.head = node

	iterator := L.head

	for iterator.next != nil {
		iterator = iterator.next
	}

	L.tail = iterator
}

func (L *List) AddToTail(data int) {
	node := &Node{
		data: data,
	}

	if L.head == nil {
		L.head = node
		L.tail = node

		return
	}

	iterator := L.head

	for iterator.next != nil {
		iterator = iterator.next
	}

	node.next = iterator.next
	iterator.next = node

	L.tail = node
}

func (L *List) AddAfter(index int, data int) {

	if index <= 0 {
		return
	}

	iterator := L.head

	for i := 0; i < index-1; i++ {
		iterator = iterator.next

		if iterator == nil {
			fmt.Printf("There is less than %v elements\n", index)
			return
		}
	}

	node := &Node{
		data: data,
	}

	node.next = iterator.next
	iterator.next = node

	L.setTail()
}

func (L *List) DeleteByData(data int) {

	head := L.head

	if head.data == data {
		L.head = head.next

		return
	}

	iterator := L.head

	for iterator.next != nil && iterator.next.data != data {
		iterator = iterator.next
	}

	if iterator.next == nil {
		fmt.Printf("Could not find the data %v \n", data)

		return
	}

	iterator.next = iterator.next.next
}

func (L *List) DeleteByIndex(index int) {

	head := L.head

	if index == 0 {
		L.head = head.next

		return
	}

	counter := 1
	iterator := L.head

	for iterator.next != nil && counter != index {
		iterator = iterator.next

		counter++
	}

	if iterator.next == nil {
		fmt.Printf("Could not find the index %v \n", index)

		return
	}

	iterator.next = iterator.next.next
}

func (l *List) Display() {
	iterator := l.head

	for iterator != nil {
		fmt.Printf("%v -> ", iterator.data)
		iterator = iterator.next
	}

	fmt.Println()
}

func (l *List) ShowBackwards() {
	iterator := l.tail

	for iterator != nil {
		fmt.Printf("%v <- ", iterator.data)
		iterator = iterator.prev
	}

	fmt.Println()
}

func (l *List) Reverse() {
	iterator := l.head
	var prev *Node
	l.tail = l.head

	for iterator != nil {
		next := iterator.next
		iterator.next = prev
		prev = iterator
		iterator = next
	}

	l.head = prev

	/* Display(l.head) */
}

func (l *List) setTail() {
	iterator := l.head

	for iterator.next != nil {
		iterator = iterator.next
	}

	l.tail = iterator
}

func (l *List) GetHead() int {
	return l.head.data
}

func (l *List) GetTail() int {
	return l.tail.data
}

func (l *List) GetByIndex(index int) int {
	counter := 1
	iterator := l.head

	for iterator.next != nil && counter != index {
		iterator = iterator.next

		counter++
	}

	return iterator.next.data
}
