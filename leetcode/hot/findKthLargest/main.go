package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k := 2
	res := findKthLargest(nums, k)
	fmt.Println(res)
}

// 是的，我现在的性能已经很差了，所以需要换个项目做做。
func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(nums []int, l, r int, k int) int {
	if l == r {
		return nums[k]
	}
	partition := nums[l]
	i := l - 1
	j := r + 1
	for i < j {
		i++
		j--
		for nums[i] < partition {
			i++
		}
		for nums[j] > partition {
			j--
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}

	}

	// 导致区间一直向 k 逼近
	if k <= j {
		return quickSelect(nums, l, j, k)
	} else {
		return quickSelect(nums, j+1, r, k)
	}
}
