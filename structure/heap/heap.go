package heap

import (
	"TheAlgorithm/constraints"
	"errors"
)

// Heap heap implementation using generic.
type Heap[T any] struct {
	heaps    []T
	lessFunc func(a, b T) bool
}

// New gives a new default Min-heap
func New[T constraints.Ordered]() *Heap[T] {
	less := func(a, b T) bool {
		return a < b
	}
	h, _ := NewAny[T](less)
	return h
}

// NewAny gives a new heap object. element can be anything, but must provide less function.
func NewAny[T any](less func(a, b T) bool) (*Heap[T], error) {
	if less == nil {
		return nil, errors.New("less func is necessary")
	}
	return &Heap[T]{
		lessFunc: less,
	}, nil
}

func (heap *Heap[T]) Push(t T) {
	heap.heaps = append(heap.heaps, t)
	heap.heapInsert(len(heap.heaps) - 1)
}

func (heap *Heap[T]) Top() T {
	return heap.heaps[0]
}

func (heap *Heap[T]) Pop() {
	if heap.Size() == 1 {
		heap.heaps = nil
		return
	}

	heap.swap(0, heap.Size()-1)
	heap.heaps = heap.heaps[:len(heap.heaps)-1]
	heap.heapify(0)
}

func (heap Heap[T]) Empty() bool {
	return len(heap.heaps) == 0
}

func (heap *Heap[T]) Size() int {
	return len(heap.heaps)
}

// Min-heap
func (heap *Heap[T]) heapInsert(child int) {
	if child <= 0 {
		return
	}
	parent := (child - 1) >> 1
	if heap.lessFunc(heap.heaps[child], heap.heaps[parent]) {
		heap.swap(child, parent)
		heap.heapInsert(parent)
	}
}

func (heap *Heap[T]) Resign(index int) {
	if index < 0 || index > heap.Size() {
		return
	}
	heap.heapInsert(index)
	heap.heapify(index)
}

func (heap *Heap[T]) heapify(parent int) {
	lessIndex := parent
	lChild, rChild := (parent<<1)+1, (parent<<1)+2
	if lChild < len(heap.heaps) && heap.lessFunc(heap.heaps[lChild], heap.heaps[lessIndex]) {
		lessIndex = lChild
	}
	if rChild < len(heap.heaps) && heap.lessFunc(heap.heaps[rChild], heap.heaps[lessIndex]) {
		lessIndex = rChild
	}
	if lessIndex == parent {
		return
	}
	heap.swap(parent, lessIndex)
	heap.heapify(lessIndex)
}

func (heap *Heap[T]) swap(i, j int) {
	heap.heaps[i], heap.heaps[j] = heap.heaps[j], heap.heaps[i]
}
