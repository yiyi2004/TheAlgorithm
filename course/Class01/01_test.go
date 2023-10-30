package Class01

import (
	"TheAlgorithm/utils"
	"fmt"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	testNum := 5
	maxSize := 10
	maxValue := 100

	for i := 0; i < testNum; i++ {
		arr := utils.GenerateRandomArray(maxSize, maxValue)
		selectionSort(arr)
		fmt.Println(arr)
	}
}

func Test_nearLeft(t *testing.T) {
	arr := []int{2, 6, 7, 8, 10, 11, 20}
	left := nearLeft(arr, 12)
	fmt.Println(arr[left])
}

func Test_nearRight(t *testing.T) {
	arr := []int{2, 6, 7, 8, 10, 11, 20}
	right := nearRight(arr, 12)
	fmt.Println(arr[right])
}

func Test_getLessIndex(t *testing.T) {
	arr := []int{100, 10, 7, 6, 8, 10, 11, 20}
	localMin := getLessIndex(arr)
	fmt.Println(arr[localMin])
}
