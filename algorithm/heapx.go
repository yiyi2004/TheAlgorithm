package algorithm

func main() {

}

type IntHeap []int

func (heap *IntHeap) Len() int {
	return len(*heap)
}

func (heap *IntHeap) Less(i, j int) bool {
	return (*heap)[i] < (*heap)[j]
}

func (heap *IntHeap) Swap(i, j int) {
	(*heap)[i], (*heap)[j] = (*heap)[j], (*heap)[i]
}

func (heap *IntHeap) Push(x any) {
	xv, ok := x.(int)
	if ok {
		*heap = append(*heap, xv)
		return
	}
	return
}

func (heap *IntHeap) Pop() any {
	res := (*heap)[heap.Len()-1]
	*heap = (*heap)[:heap.Len()-1]
	return res
}
