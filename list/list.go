// Package list has methods to make and manage linked lists
package list

import "errors"

type node struct {
	value int
	next  *node
}

type LinkedList struct {
	head *node
	tail *node
	size int
}

// Make a linked list
func MakeLinkedList() LinkedList {
	list := LinkedList{
		head: nil,
		tail: nil,
		size: 0,
	}

	return list
}

// Add to the head of the linked list
func (list *LinkedList) AddFirst(value int) {
    newNode := &node{
		value: value,
		next:  nil,
	}

	if list.head == nil {
		list.head = newNode
		list.tail = newNode
		list.size = 1
		return
	}

    tmp := list.head
	list.head = newNode
	list.head.next = tmp
	list.size += 1
}

// Remove from the Head
func (list *LinkedList) RemoveFirst() (int, error) {
    if list.head == nil {
        return 0, errors.New("Can't remove from an Empty List")
    }
    
    purged := list.head
    list.head = list.head.next
    list.size -= 1
    if list.size == 0 {
        list.tail = nil
    }

    return purged.value, nil
}

// Get the Linked List Size
func (list LinkedList) GetSize() int {
    if list.head == nil {
        return 0
    }

	return list.size
}

func (list LinkedList) Peek() (int, bool) {
	return list.size, list.head == nil
}
