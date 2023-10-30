package Class04

import (
	"fmt"
	"testing"
)

func Test_mergesort1(t *testing.T) {
	arr := []int{2, 4, 11, 2, 24, 3, 114514, 111, 4, 56, 11111, 4}
	mergesort1(arr)
	fmt.Println(arr)
}

func Test_mergesort2(t *testing.T) {
	arr := []int{2, 4, 11, 2, 24, 3, 114514, 111, 4, 56, 11111, 4}
	mergesort2(arr)
	fmt.Println(arr)
}
