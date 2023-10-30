package Class04

import (
	"fmt"
	"testing"
)

func Test_mergeSort1(t *testing.T) {
	arr := []int{2, 1, 22, 12, 14, 5, 6}
	mergeSort1(arr)
	fmt.Println(arr)
}

func Test_smallSum(t *testing.T) {
	arr := []int{2, 1, 22, 12, 14, 5, 6}
	sum := smallSum(arr)
	fmt.Println(sum)
	fmt.Println(arr)
}

func Test_reversePair(t *testing.T) {
	//arr := []int{1, 2, 22, 12, 14, 5, 6}
	arr := []int{1, 2, 3, 4, 5, 6}
	rp := reversePair(arr)
	fmt.Println(rp)
	fmt.Println(arr)

	// golang 里面是不允许歧义的操作的
}
