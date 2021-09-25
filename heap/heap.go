package heap

import "errors"

type Heap struct {
	values []int
	size   int
}

func MakeMaxHeap(input []int) *Heap {
	values := make([]int, len(input))
	h := &Heap{
		values: values,
		size:   0,
	}

	for _, v := range input {
		h.MaxInsert(v)
	}

	return h
}

func (h *Heap) MaxInsert(value int) *Heap {
	if h.size == len(h.values) {
		h.values = append(h.values, value)
	}
	h.values[h.size] = value
	h.size += 1

	return h.maxInsertHeapify(h.size - 1)
}

func (h *Heap) maxInsertHeapify(idx int) *Heap {
	if idx == 0 {
		return h
	}

	pIdx := (idx - 1) / 2
	if h.values[pIdx] < h.values[idx] {
		swap(h.values, pIdx, idx)
		h.maxInsertHeapify(pIdx)
	}

	return h
}

func swap(values []int, i, j int) {
	tmp := values[i]
	values[i] = values[j]
	values[j] = tmp
}

func HeapSort(values []int) []int {
	h := MakeMaxHeap(values)

	for idx, v := range h.Values() {
		values[idx] = v
	}

	for i := len(values); i > 0; i-- {
		swap(values, 0, i-1)
		maxHeapify(0, values, i-1)
	}

	return values
}

func maxHeapify(idx int, values []int, size int) {
	leftIdx := 2*idx + 1
	rightIdx := 2*idx + 2
	largestIdx := idx

	if leftIdx < size && values[idx] < values[leftIdx] {
		largestIdx = leftIdx
	}

	if rightIdx < size && values[largestIdx] < values[rightIdx] {
		largestIdx = rightIdx
	}

	if largestIdx != idx {
		swap(values, idx, largestIdx)
		maxHeapify(largestIdx, values, size)
	}
}

func MakeMinHeap(size int) Heap {
	values := make([]int, size)
	return Heap{
		values: values,
		size:   size,
	}
}

func (h *Heap) Insert(value int) *Heap {
	if h.size == len(h.values) {
		h.values = append(h.values, value)
	}
	h.size += 1
	return h.minInsertHeapify(h.size - 1)
}

func (h Heap) Values() []int {
	return h.values[:h.size]
}

func (h *Heap) minInsertHeapify(idx int) *Heap {
	if idx == 0 {
		return h
	}

	pIdx := (idx - 1) / 2

	if h.values[pIdx] > h.values[idx] {
		swap(h.values, pIdx, idx)
		h.minInsertHeapify(pIdx)
	}

	return h
}

func (h *Heap) minHeapify(idx int) {
	leftIdx := 2*idx + 1
	rightIdx := 2*idx + 2
	smallestIdx := idx

	if leftIdx < h.size && h.values[idx] > h.values[leftIdx] {
		smallestIdx = leftIdx
	}

	if rightIdx < h.size && h.values[smallestIdx] > h.values[rightIdx] {
		smallestIdx = rightIdx
	}

	if smallestIdx != idx {
		swap(h.values, idx, smallestIdx)
		h.minHeapify(smallestIdx)
	}
}

func (h Heap) findIndex(value int) (int, error) {
	idx := -1
	for i := 0; i < h.size; i++ {
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
	if h.size == 0 {
		return h
	}

	idx, err := h.findIndex(value)

	if err != nil {
		return h
	}

	h.values[idx] = h.values[h.size-1]
	h.size -= 1
	h.minHeapify(idx)

	return h
}
