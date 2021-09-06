package list_test

import (
	"fmt"
	"go-data-structure/list"
	"testing"
)

func CreateUpdateLinkedListTest(t *testing.T) {
	linklist := list.MakeLinkedList()
	size := linklist.GetSize()

	if size != 0 {
		t.Fatal("New LinkedList size is not 0 and is not empty")
	}

	linklist.AddFirst(9)
	linklist.AddFirst(7)
	linklist.AddFirst(10)
	str := fmt.Sprintf("%v", linklist)

	if str != "(10)=>(7)=>(9)" {
		t.Fatalf("Stringify Failed, %v", linklist)
	}

	size = linklist.GetSize()

	if size != 3 {
		t.Fatalf("Linklist Size %d is not 3", size)
	}

	removed, err := linklist.RemoveFirst()
	if removed != 10 {
		t.Fatalf("Removed First %d", removed)
	}

	peeked, isEmpty := linklist.PeekFirst()
	if peeked != 7 && isEmpty == false {
		t.Fatalf("Failed Peek returned %d", peeked)
	}

	removed, err = linklist.RemoveFirst()
	if removed != 7 {
		t.Fatalf("Removed First %d", removed)
	}

	size = linklist.GetSize()

	if size != 1 {
		t.Fatalf("Linklist Size %d is not 3", size)
	}

	removed, err = linklist.RemoveFirst()

	if removed != 9 {
		t.Fatalf("Removed First %d", removed)
	}

	size = linklist.GetSize()

	if size != 0 {
		t.Fatalf("Linklist Size %d is not 3", size)
	}

	removed, err = linklist.RemoveFirst()
	if err == nil {
		t.Fatalf("Removing First on Empty Linked List did not return ")
	}

	size = linklist.GetSize()

	if size != 0 {
		t.Fatalf("Linklist Size %d is not 3", size)
	}
}

func CreateUpdateLinkedListTest2(t *testing.T) {
	linklist := list.MakeLinkedList()
	size := linklist.GetSize()

	if size != 0 {
		t.Fatal("New LinkedList size is not 0 and is not empty")
	}

	linklist.AddFirst(9)
	linklist.AddLast(7)
	linklist.AddLast(10)
	str := fmt.Sprintf("%v", linklist)

	if str != "(9)=>(7)=>(10)" {
		t.Fatalf("Stringify Failed, %v", linklist)
	}

	size = linklist.GetSize()

	if size != 3 {
		t.Fatalf("Linklist Size %d is not 3", size)
	}

	removed, err := linklist.RemoveLast()
	if removed != 10 {
		t.Fatalf("Removed Last %d", removed)
	}

	peeked, isEmpty := linklist.PeekLast()
	if peeked != 7 && isEmpty == false {
		t.Fatalf("Failed Peek Last returned %d", peeked)
	}

	removed, err = linklist.RemoveFirst()
	if removed != 9 {
		t.Fatalf("Removed First %d", removed)
	}

	size = linklist.GetSize()

	if size != 1 {
		t.Fatalf("Linklist Size %d is not 3", size)
	}

	removed, err = linklist.RemoveLast()

	if removed != 7 {
		t.Fatalf("Removed First %d", removed)
	}

	size = linklist.GetSize()

	if size != 0 {
		t.Fatalf("Linklist Size %d is not 3", size)
	}

	removed, err = linklist.RemoveLast()
	if err == nil {
		t.Fatalf("Removing Last on Empty Linked List did not return ")
	}

	size = linklist.GetSize()

	if size != 0 {
		t.Fatalf("Linklist Size %d is not 3", size)
	}
}
