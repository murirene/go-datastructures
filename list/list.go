// Package list has methods to make and manage linked lists
package list

import (
	"errors"
	"fmt"
	"strings"
)

type node struct {
	value int
	next  *node
	prev  *node
}

type LinkedList struct {
	head *node
	tail *node
	size int
}

// Stringer interface implementation for node
func (n node) String() string {
	link := ""
	if n.next != nil {
		link = "=>"
	}
	return fmt.Sprintf("(%d)%s", n.value, link)
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

// Stringer interface implementation for Linked List
func (l LinkedList) String() string {
	if l.head == nil {
		return "Empty List"
	}
	var sb strings.Builder
	runner := l.head
	for runner != nil {
		sb.WriteString(fmt.Sprintf("%v", runner))
		runner = runner.next
	}
	return sb.String()
}

// Add to the head of the linked list
func (list *LinkedList) AddFirst(value int) {
	newNode := &node{
		value: value,
		next:  nil,
		prev:  nil,
	}

	if list.head == nil {
		list.head = newNode
		list.tail = newNode
		list.size = 1
		return
	}

	tmp := list.head
	list.head = newNode
	tmp.prev = list.head
	list.head.next = tmp
	list.size += 1
}

func (list *LinkedList) AddLast(value int) {
	newNode := &node{
		value: value,
		next:  nil,
		prev:  nil,
	}

	if list.head == nil {
		list.head = newNode
		list.tail = newNode
		list.size = 1
		return
	}
	list.tail.next = newNode
	newNode.prev = list.tail
	list.tail = newNode
	list.size += 1
}

// Remove from the Head
func (list *LinkedList) RemoveFirst() (int, error) {
	if list.head == nil {
		return 0, errors.New("Can't remove from an Empty List")
	}

	purged := list.head
	list.head = list.head.next
	list.head.prev = nil
	list.size -= 1
	if list.size == 0 {
		list.tail = nil
	}

	return purged.value, nil
}

func (list *LinkedList) RemoveLast() (int, error) {
	if list.head == nil {
		return 0, errors.New("Can't remove from an Empty List")
	}

	purged := list.tail
	list.tail = list.tail.prev
	list.tail.next = nil
	list.size -= 1
	if list.size == 0 {
		list.head = nil
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

// Peek the value of the first element
func (list LinkedList) PeekFirst() (int, bool) {
	if list.head == nil {
		return 0, true
	}

	return list.head.value, false
}

func (list LinkedList) PeekLast() (int, bool) {
	if list.tail == nil {
		return 0, true
	}

	return list.tail.value, false
}
