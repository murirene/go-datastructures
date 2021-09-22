package heap

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
	h.minHeapify(idx)

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

func (h *Heap) minHeapify(idx int) {
	if idx == 0 {
		return
	}
	pIdx := (idx - 1) / 2

	if h.values[pIdx] > h.values[idx] {
		swap(h.values, pIdx, idx)
		h.minHeapify(pIdx)
	}
}
