package main

type minHeap struct {
	array []int
}

func (h *minHeap) Insert(v int) {
	h.array = append(h.array, v)
	h.MinHeapifyUP(len(h.array) - 1)
}

func (h *minHeap) Extract() int {
	lastIdx := len(h.array) - 1
	if lastIdx < 0 {
		return -1
	}

	ex := h.array[0]

	h.array[0] = h.array[lastIdx]
	h.array = h.array[:lastIdx]
	h.MinHeapifyDown(0)

	return ex
}

func (h *minHeap) MinHeapifyUP(i int) {
	pi := parent(i)
	for h.array[i] < h.array[pi] {
		h.Swap(i, pi)
		i = pi
		pi = parent(i)
	}
}

func (h *minHeap) MinHeapifyDown(i int) {
	li, ri := left(i), right(i)
	lastIdx := len(h.array) - 1

	var idxToCompare int
	for li <= lastIdx {

		if li == lastIdx || h.array[li] < h.array[ri] {
			idxToCompare = li
		} else {
			idxToCompare = ri
		}

		if h.array[idxToCompare] < h.array[i] {
			h.Swap(idxToCompare, i)
			i = idxToCompare
			li, ri = left(i), right(i)
		} else {
			return
		}
	}

}

func (h *minHeap) Swap(i, ir int) {
	h.array[i], h.array[ir] = h.array[ir], h.array[i]
}
