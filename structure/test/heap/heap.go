package main

func main() {

}

// Heap 实现大根堆
// 实现 heapInsert 和 heapify 操作之后，实现堆的其他操作都比较简单了。
type Heap struct {
	limit    int
	heap     []int
	heapSize int
}

func NewHeap(limit int) *Heap {
	return &Heap{
		limit:    limit,
		heap:     make([]int, limit),
		heapSize: 0,
	}
}

func (h *Heap) IsEmpty() bool {
	return h.heapSize == 0
}

func (h *Heap) IsFull() bool {
	return h.heapSize == h.limit
}

func (h *Heap) Push(v int) {
}

func (h *Heap) Pop() int {
}

func (h *Heap) heapInsert(index int) {
	// up
	for h.heap[index] > h.heap[(index-1)/2] {
		h.swap(index, (index-1)/2)
		index = (index - 1) / 2
	}
}

func (h *Heap) heapify(index int) {
	// down
	left := index*2 + 1
	for left < h.heapSize { // 如果有左孩子，有没有右孩子，可能有可能没有！
		// 把较大孩子的下标，给largest
		var largest int
		if left+1 < h.heapSize && h.heap[left+1] > h.heap[left] {
			largest = left + 1
		} else {
			largest = left
		}
		if !(h.heap[largest] > h.heap[index]) {
			largest = index
		}
		if largest == index {
			break
		}
		// index和较大孩子，要互换
		h.swap(largest, index)
		index = largest
		left = index*2 + 1
	}
}

func (h *Heap) swap(i, j int) {
	h.heap[i], h.heap[j] = h.heap[j], h.heap[i]
}
