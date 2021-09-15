package list_test

import (
	"fmt"
	"go-data-structure/list"
	"testing"
)
/*
func TestCreateUpdateLinkedList(t *testing.T) {
	linklist := list.MakeLinkedList(0)
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

func TestCreateUpdateLinkedList2(t *testing.T) {
	linklist := list.MakeLinkedList(0)
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

func TestMiddleNode(t *testing.T) {
    l := list.MakeLinkedList(0)
    l.AddFirst(1)
    middle := list.FindMiddleNode(l.GetHead(), l.GetTail())
    
    if middle != nil &&  middle.GetValue() != 1 {
        t.Fatal("Value is not right")
    }
}

func TestMiddleNode2(t *testing.T) {
    l := list.MakeLinkedList(0)
    l.AddFirst(1)
    l.AddFirst(2)
    middle := list.FindMiddleNode(l.GetHead(), l.GetTail())
    
    if middle != nil &&  middle.GetValue() != 2 {
        t.Fatalf("Value is not right %d", middle.GetValue())
    }
}

func TestMiddleNode3(t *testing.T) {
    l := list.MakeLinkedList(0)
    l.AddFirst(1)
    l.AddFirst(2)
    l.AddFirst(3)
    middle := list.FindMiddleNode(l.GetHead(), l.GetTail())
    
    if middle != nil &&  middle.GetValue() != 2 {
        t.Fatal("Value is not right")
    }
}
*/
func TestSorting(t *testing.T) {
    l := list.MakeLinkedList(0) 
    l.AddFirst(4)
    l.AddFirst(3)
    l.AddFirst(6)
    l.AddFirst(7)
    l.AddFirst(1)
    l.Mergesort()
    
	str := fmt.Sprintf("%v\n", l)
	if str != "(1)=>(2)=>(3)=>(4)=>(5)=>(6)" {
		t.Fatalf("Stringify Failed, %v", l)
	}
}
