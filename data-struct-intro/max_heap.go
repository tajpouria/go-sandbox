package main

import (
	"fmt"
	"math"
)

type maxHeap struct {
	array []int
}

func (h *maxHeap) Insert(v int) {
	h.array = append(h.array, v)
	h.MaxHeapifyUp(len(h.array) - 1)
}

func (h *maxHeap) Extract() int {
	li := len(h.array) - 1
	if li == 0 {
		return -1
	}

	ex := h.array[0]
	h.array[0] = h.array[li]
	h.array = h.array[:li]

	h.MaxHeapifyDown(0)

	return ex
}

func (h *maxHeap) MaxHeapifyUp(i int) {
	for h.array[i] > h.array[parent(i)] {
		h.Swap(i, parent(i))
		i = parent(i)
	}
}

func (h maxHeap) MaxHeapifyDown(i int) {
	lai := len(h.array) - 1
	li, ri := left(i), right(i)

	for li <= lai {
		var toCompareIdx int
		if li == lai || h.array[li] > h.array[ri] {
			toCompareIdx = li
		} else {
			toCompareIdx = ri
		}

		if h.array[toCompareIdx] > h.array[i] {
			h.Swap(toCompareIdx, i)
			i = toCompareIdx
			li, ri = left(i), right(i)
		} else {
			return
		}
	}
}

func (h *maxHeap) Swap(i, ri int) {
	h.array[i], h.array[ri] = h.array[ri], h.array[i]
}

func (h *maxHeap) Print() {
	lastIdx := len(h.array) - 1
	if lastIdx <= 0 {
		return
	}
	fmt.Println(h.array[0])

	numRows := int(math.Floor(math.Log2(float64(len(h.array)))))
	for r := 1; r <= numRows; r++ {
		capacity := int(math.Pow(2.0, float64(r)))
		for i := 0; i < capacity; i++ {
			if i*2+r > lastIdx {
				break
			}
			fmt.Printf("%v ", h.array[i*2+r])
		}
		for i := 0; i < capacity; i++ {
			if i*2+r > lastIdx {
				break
			}
		}
		fmt.Print("\n")
	}
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}
