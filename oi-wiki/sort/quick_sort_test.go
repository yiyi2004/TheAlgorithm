package main

import (
	"fmt"
	"testing"
)

func Test_quickSort(t *testing.T) {
	arr := []int{12, 4, 5, 6, 7, 3, 1, 15}
	sortedArr := basicQsort(arr)
	fmt.Println(sortedArr)
}
