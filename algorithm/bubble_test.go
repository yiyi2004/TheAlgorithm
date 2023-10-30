package algorithm

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	intComparator := func(a, b int) bool {
		return a > b
	}

	arr := []int{2, 4, 5, 11, 4, 5, 1}
	fmt.Println(arr)
	BubbleSort(arr, intComparator)
	fmt.Println(arr)

	float32Comparator := func(a, b float32) bool {
		return a > b
	}

	float32Arr := []float32{10e-3, 5.1, 11.2, 4.3, 5.1, 1.2}
	BubbleSort(float32Arr, float32Comparator)
	fmt.Println(float32Arr)

}
