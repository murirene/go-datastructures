package heap_test

import (
	"go-data-structure/heap"
	"testing"
)

func sliceIsEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

func TestHeapInsert(t *testing.T) {
	h := heap.MakeMinHeap()
	h.Insert(3).Insert(1).Insert(6).Insert(5).Insert(2).Insert(4)
	results := []int{1, 2, 4, 5, 3, 6}
	values := h.Values()

	if sliceIsEqual(values, results) == false {
		t.Fatalf("%v is not %v", values, results)
	}
}

func TestHeapDelete(t *testing.T) {
	results := []int{2, 3, 6, 5}

	h := heap.MakeMinHeap()
	h.Insert(3).Insert(1).Insert(6).Insert(5).Insert(2).Insert(4)
	h.Delete(1).Delete(4)

	values := h.Values()
	if sliceIsEqual(values, results) == false {
		t.Fatalf("%v is not %v", values, results)
	}
}
