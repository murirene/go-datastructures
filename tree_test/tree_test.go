package main

import (
	"fmt"
	"go-data-structure/tree"
	"testing"
)

func TestCRUBinarySearchTree(t *testing.T) {
	b := tree.MakeBinarySearchTree(4)
	if b.GetSize() != 1 {
		t.Fatalf("Tree does not have a size of 1")
	}
	b.Add(6)
	b.Add(3)
	b.Add(2)
	b.Add(9)
	b.Add(5)

	str := fmt.Sprintf("%v", b)

	if str != "(2)(3)(4)(5)(6)(9)" && b.GetSize() != 6 {
		t.Fatalf("Failed Stringer compare %v", b)
	}
	str = b.BfsString()

	if str != "(4)(3)(6)(2)(5)(9)" {
		t.Fatalf("Failed BFS %s", str)
	}

	b.Remove(6)

	str = b.BfsString()

	if str != "(4)(3)(9)(2)(5)" {
		t.Fatalf("Failed BFS %s", str)
	}

	str = fmt.Sprintf("%v", b)

	if str != "(2)(3)(4)(5)(9)" && b.GetSize() != 5 {
		t.Fatalf("Failed Stringer compare %v", b)
	}

	b.Remove(2)
	str = b.BfsString()

	if str != "(4)(3)(9)(5)" {
		t.Fatalf("Failed BFS %s", str)
	}
	b.Remove(9)

	str = b.BfsString()

	if str != "(4)(3)(5)" {
		t.Fatalf("Failed BFS %s", str)
	}

	b.Remove(5)

	str = b.BfsString()
	if str != "(4)(3)" {
		t.Fatalf("Failed BFS %s", str)
	}
}

func TestSecondLargest(t *testing.T) {
	b := tree.MakeBinarySearchTree(20)
	b.Add(10)
	b.Add(5)
	val, err := b.GetSecondLargest()

	if val != 10 && err == nil {
		t.Fatalf("%d is not 10", val)
	}
}

func TestSecondLargest2(t *testing.T) {
	b := tree.MakeBinarySearchTree(20)
	b.Add(10)
	b.Add(5)
	b.Add(15)
	b.Add(12)
	val, err := b.GetSecondLargest()

	if val != 15 && err == nil {
		t.Fatalf("%d is not 15", val)
	}
}

func TestSecondLargest3(t *testing.T) {
	b := tree.MakeBinarySearchTree(20)
	b.Add(30)
	b.Add(23)
	b.Add(21)
	b.Add(22)
	val, err := b.GetSecondLargest()

	if val != 23 && err == nil {
		t.Fatalf("%d is not 23", val)
	}
}
