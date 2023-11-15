package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{6, 10}
	index := 0
	num := 2
	arr = append(arr[:index], append([]int{num}, arr[index:]...)...)
	fmt.Println(arr)
}

type ele struct {
	fre int
	val int
}

type es []ele

func topKFrequent(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}
	var e es
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	for key, val := range m {
		el := ele{
			fre: val,
			val: key,
		}
		e = append(e, el)
	}
	sort.Slice(e, func(i, j int) bool {
		return e[i].fre > e[j].fre
	})
	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = e[i].val
	}
	return res
}
