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

func (n node) GetValue() int {
	return n.value
}

// Make a linked list
func MakeLinkedList(size int) LinkedList {
	list := LinkedList{
		head: nil,
		tail: nil,
		size: 0,
	}

	for i := 0; i < size; i++ {
		list.AddFirst(0)
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

	purged := list.head.value
	list.head = list.head.next
	if list.head != nil {
		list.head.prev = nil
	}
	list.size -= 1
	if list.size == 0 {
		list.tail = nil
	}

	return purged, nil
}

func (list *LinkedList) RemoveLast() (int, error) {
	if list.head == nil {
		return 0, errors.New("Can't remove from an Empty List")
	}

	purged := list.tail.value
	list.tail = list.tail.prev
	if list.tail != nil {
		list.tail.next = nil
	}
	list.size -= 1
	if list.size == 0 {
		list.head = nil
	}

	return purged, nil
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

func (list LinkedList) GetHead() *node {
	return list.head
}

func (list LinkedList) GetTail() *node {
	return list.tail
}

func FindMiddleNode(head, tail *node) *node {
	if head == tail {
		return head
	}
	middleRunner := head
	fastRunner := head.next

	for fastRunner != tail && fastRunner != nil {
		middleRunner = middleRunner.next
		fastRunner = fastRunner.next
		if fastRunner != nil {
			fastRunner = fastRunner.next
		}
	}
	return middleRunner
}

func (l *LinkedList) Mergesort() {
	if l.head == nil || l.tail == nil {
		return
	}

	buffer := MakeLinkedList(l.GetSize())
	mergesort(l.head, l.tail, buffer)
}

func merge(head, mid, tail *node, buffer *LinkedList) {
    fmt.Printf("merge(head=%d, mid=%d, tail=%d)\n", head.value, mid.value, tail.value)
    left := head
	right := mid.next
	sorted := buffer.head

	for left != mid.next || right != tail.next {
		if right == tail.next || (left != mid.next && left.value < right.value) {
			sorted.value = left.value
			left = left.next
		} else {
			sorted.value = right.value
			right = right.next
		}
		sorted = sorted.next
	}

	sorted = buffer.head
	for runner := head; runner != tail.next; runner = runner.next {
		runner.value = sorted.value
		sorted = sorted.next
		runner = runner.next
	}
}

func mergesort(head, tail *node, buffer LinkedList) {
    fmt.Printf("mergesort(head=%d, tail=%d)\n", head.value, tail.value)
	if head == tail {
		return
	}

	if head == nil || tail == nil {
		return
	}

	mid := FindMiddleNode(head, tail)

	mergesort(head, mid, buffer)
	mergesort(mid.next, tail, buffer)
	merge(head, mid, tail, &buffer)
}
