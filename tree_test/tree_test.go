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
}
