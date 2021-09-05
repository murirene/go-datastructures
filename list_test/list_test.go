package list_test

import (
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
	size = linklist.GetSize()

	if size != 3 {
		t.Fatalf("Linklist Size %d is not 3", size)
	}

    removed, err := linklist.RemoveFirst()
    if removed != 10 {
        t.Fatalf("Removed First %d", removed)
    }

    peeked, isEmpty := linklist.Peek()
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
    if err == nil  {
        t.Fatalf("Removing First on Empty Linked List did not return ")
    }

	size = linklist.GetSize()

	if size != 0 {
		t.Fatalf("Linklist Size %d is not 3", size)
	}
}
