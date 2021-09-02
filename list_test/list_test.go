package list_test

import (
	"go-data-structure/list"
	"testing"
)

func CreateUpdateLinkedListTest(t *testing.T) {
	linklist := list.MakeLinkedList()
	size, isEmpty := linklist.GetSize()

	if size != 0 && !isEmpty {
		t.Fatal("New LinkedList size is not 0 and is not empty")
	}

	linklist.AddFirst(9)
	linklist.AddFirst(7)
	linklist.AddFirst(10)

	size, isEmpty = linklist.GetSize()

	if size != 3 {
		t.Fatalf("Linklist Size %d is not 3", size)
	}

	if isEmpty == true {
		t.Fatalf("Linked List of size %d is returning Empty Status", size)
	}
}
