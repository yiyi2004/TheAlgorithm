package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates([]int{1, 2, 2}))
}

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	loc := -1
	curNum := nums[0]
	count := 0
	for i := 0; i < len(nums); i++ {
		if curNum == nums[i] {
			count++
			if count == 3 {
				loc = i
				break
			}
		} else {
			curNum = nums[i]
			count = 1
		}
	}
	if loc == -1 {
		return len(nums)
	}
	for i := loc + 1; i < len(nums); i++ {
		if nums[i] == curNum {
			count++
			if count <= 2 {
				nums[loc] = curNum
				loc++
			}
		} else {
			curNum = nums[i]
			count = 1
			if loc > 0 {
				nums[loc] = curNum
				loc++
			}
		}
	}

	return loc
}
