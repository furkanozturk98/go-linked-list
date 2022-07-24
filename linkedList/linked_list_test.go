package list

import (
	"testing"
)

func TestAddToHead(t *testing.T) {
	list := List{}
	list.AddToHead(1)
	list.AddToHead(2)

	expected := 2
	actual := list.GetHead()

	if actual != expected {
		t.Errorf("Got %v expected %v", actual, expected)
	}
}

func TestAddToTail(t *testing.T) {
	list := List{}
	list.AddToTail(1)
	list.AddToTail(2)

	expected := 2
	actual := list.GetTail()

	if actual != expected {
		t.Errorf("Got %v expected %v", actual, expected)
	}
}

func TestAddAfter(t *testing.T) {
	list := List{}
	list.AddToHead(1)
	list.AddToTail(2)
	list.AddToTail(3)

	// Add 5 to index 1
	list.AddAfter(1, 5)

	expected := 5
	actual := list.GetByIndex(1)

	if actual != expected {
		t.Errorf("Got %v expected %v", actual, expected)
	}
}

func TestDeleteByData(t *testing.T) {
	list := List{}
	list.AddToHead(1)
	list.AddToTail(2)
	list.AddToTail(3)

	list.DeleteByData(2)

	expected := 3
	actual := list.GetByIndex(1)

	if actual != expected {
		t.Errorf("Got %v expected %v", actual, expected)
	}
}

func TestDeleteByIndex(t *testing.T) {
	list := List{}
	list.AddToHead(1)
	list.AddToTail(2)
	list.AddToTail(3)

	list.DeleteByIndex(1)

	expected := 3
	actual := list.GetByIndex(1)

	if actual != expected {
		t.Errorf("Got %v expected %v", actual, expected)
	}
}

func TestReverse(t *testing.T) {
	list := List{}
	list.AddToHead(1)
	list.AddToTail(2)
	list.AddToTail(3)

	list.Reverse()

	expected := 3
	actual := list.GetHead()

	if actual != expected {
		t.Errorf("Got %v expected %v", actual, expected)
	}
}
