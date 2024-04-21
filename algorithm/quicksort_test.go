package algorithm

import "testing"

func partition(list []int, low, high int) int {
	pivot := list[low]
	for low < high {
		for low < high && list[high] >= pivot {
			high--
		}
		list[low] = list[high]
		for low < high && list[low] <= pivot {
			low++
		}
		list[high] = list[low]
	}
	list[low] = pivot
	return low
}

func TestQuickSort(t *testing.T) {
	list := []int{2, 44, 4, 8, 33, 1, 22, -11, 6, 34, 55, 54, 9}
	QuickSort(list, 0, len(list)-1)
	t.Log(list)
}

func QuickSort(list []int, low, high int) {
	if low < high {
		pivot := partition(list, low, high)
		QuickSort(list, low, pivot-1)
		QuickSort(list, low+1, len(list)-1)
	}
}
