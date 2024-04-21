package main

func main() {
	nextPermutation([]int{2, 3, 1})
}

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}

	// 找到一个升序对
	i, j := len(nums)-2, len(nums)-1
	for i >= 0 && j >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}
	if i == -1 {
		l, r := 0, len(nums)-1
		for l < r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
		return
	}

	// 找到第一个比 nums[i] 大的数字
	index := len(nums) - 1
	for nums[index] <= nums[i] {
		index--
	}
	// 交换
	nums[i], nums[index] = nums[index], nums[i]

	l, r := j, len(nums)-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}
