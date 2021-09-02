// Package list has methods to make and manage linked lists
package list

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
	if list.head == nil {
		list.head = &node{
			value: value,
			next:  nil,
		}
		list.tail = list.head
		list.size = 0
	}
}

func (list LinkedList) GetSize() (int, bool) {
	return list.size, list.head == nil
}

func (list LinkedList) HeadPeek() (int, bool) {
	return list.size, list.head == nil
}
