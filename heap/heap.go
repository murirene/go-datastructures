package heap

import "errors"

type Heap struct {
	values []int
}

func MakeMinHeap() Heap {
	values := make([]int, 0)
	return Heap{
		values: values,
	}
}

func (h *Heap) Insert(value int) *Heap {
	idx := len(h.values)
	h.values = append(h.values, value)
	h.minInsertHeapify(idx)

	return h
}

func swap(values []int, i, j int) {
	tmp := values[i]
	values[i] = values[j]
	values[j] = tmp
}

func (h Heap) Values() []int {
	return h.values
}

func (h *Heap) minInsertHeapify(idx int) {
	if idx == 0 {
		return
	}
	pIdx := (idx - 1) / 2

	if h.values[pIdx] > h.values[idx] {
		swap(h.values, pIdx, idx)
		h.minInsertHeapify(pIdx)
	}
}

func (h *Heap) minHeapify(idx int) {
	leftIdx := 2*idx + 1
	rightIdx := 2*idx + 2
	smallestIdx := idx

	if leftIdx < len(h.values) && h.values[idx] > h.values[leftIdx] {
		smallestIdx = leftIdx
	}

	if rightIdx < len(h.values) && h.values[smallestIdx] > h.values[rightIdx] {
		smallestIdx = rightIdx
	}

	if smallestIdx != idx {
		swap(h.values, idx, smallestIdx)
		h.minHeapify(smallestIdx)
	}
}

func (h Heap) findIndex(value int) (int, error) {
	idx := -1
	for i := 0; i < len(h.values); i++ {
		if value == h.values[i] {
			idx = i
		}
	}

	if idx == -1 {
		return 0, errors.New("Not found")
	}

	return idx, nil
}

func (h *Heap) Delete(value int) *Heap {
	if len(h.values) == 0 {
		return h
	}

	idx, err := h.findIndex(value)

	if err != nil {
		return h
	}

	h.values[idx] = h.values[len(h.values)-1]
	h.values = h.values[:len(h.values)-1]
	h.minHeapify(idx)

	return h
}
