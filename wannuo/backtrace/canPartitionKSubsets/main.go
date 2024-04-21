package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(canPartitionKSubsets([]int{4, 3, 2, 3, 5, 2, 1}, 4))
}

func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	maxVal := -1
	for _, v := range nums {
		sum += v
		if v > maxVal {
			maxVal = v
		}
	}
	if sum%k != 0 {
		return false
	}
	sum = sum / k
	if maxVal > sum {
		return false
	}
	subs := make([]int, k)

	// 考虑排序是否会对剪枝造成影响
	// 考虑同层去重
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	var dfs func(index int, subs []int) bool
	dfs = func(index int, subs []int) bool {
		if index == len(nums) {
			for i := 0; i < len(subs); i++ {
				if subs[i] != sum {
					return false
				}
			}
			return true
		}
		help := make(map[int]struct{})
		for i := 0; i < k; i++ {
			if subs[i]+nums[index] > sum {
				continue
			}
			if _, ok := help[subs[i]]; ok {
				continue
			}
			help[subs[i]] = struct{}{}
			subs[i] += nums[index]
			if dfs(index+1, subs) {
				return true
			}
			subs[i] -= nums[index]
		}
		return false
	}
	return dfs(0, subs)
}
