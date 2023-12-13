package main

import (
	"container/heap"
	"fmt"
)

func main() {
	arr := []int{6, 10}
	index := 0
	num := 2
	arr = append(arr[:index], append([]int{num}, arr[index:]...)...)
	fmt.Println(arr)
}

//type ele struct {
//	fre int
//	val int
//}
//
//type es []ele
//
//func topKFrequent(nums []int, k int) []int {
//	if len(nums) == 0 {
//		return nil
//	}
//	var e es
//	m := make(map[int]int)
//	for i := 0; i < len(nums); i++ {
//		m[nums[i]]++
//	}
//	for key, val := range m {
//		el := ele{
//			fre: val,
//			val: key,
//		}
//		e = append(e, el)
//	}
//	sort.Slice(e, func(i, j int) bool {
//		return e[i].fre > e[j].fre
//	})
//	res := make([]int, k)
//	for i := 0; i < k; i++ {
//		res[i] = e[i].val
//	}
//	return res
//}

func topKFrequent(nums []int, k int) []int {
	occurrences := map[int]int{}
	for _, num := range nums {
		occurrences[num]++
	}
	h := &IHeap{}
	heap.Init(h)
	for key, value := range occurrences {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ret := make([]int, k)
	for i := 0; i < k; i++ {
		ret[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return ret
}

type IHeap [][2]int

func (I *IHeap) Len() int {
	return len(*I)
}

func (I *IHeap) Less(i, j int) bool {
	return (*I)[i][1] < (*I)[j][1]
}

func (I *IHeap) Swap(i, j int) {
	(*I)[i], (*I)[j] = (*I)[j], (*I)[i]
}

func (I *IHeap) Push(x any) {
	(*I) = append(*I, x.([2]int))
}

func (I *IHeap) Pop() any {
	v := (*I)[len(*I)-1]
	*I = (*I)[:len(*I)-1]
	return v
}

//type IHeap [][2]int
//
//func (h IHeap) Len() int           { return len(h) }
//func (h IHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
//func (h IHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
//
//func (h *IHeap) Push(x interface{}) {
//	*h = append(*h, x.([2]int))
//}
//
//func (h *IHeap) Pop() interface{} {
//	old := *h
//	n := len(old)
//	x := old[n-1]
//	*h = old[0 : n-1]
//	return x
//}
