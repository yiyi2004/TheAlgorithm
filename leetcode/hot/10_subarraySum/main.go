package main

import "fmt"

func main() {
	nums := []int{-1, 1, 0}
	k := 0
	sum := subarraySum(nums, k)
	fmt.Println(sum)
}

func subarraySum(nums []int, k int) int {
	var res int
	m := make(map[int]int)
	pre := 0
	m[0]++

	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if _, ok := m[pre-k]; ok {
			res += m[pre-k]
		}
		m[pre]++
	}

	return res
}

// 前缀和 + 哈希表 | 公式的转换 pre[i] - pre[j] == k ---> pre[j] == pre[i] - k
// 是的，我认为你已经明白这个题如何解决了
//func subarraySum(nums []int, k int) int {
//	var res int
//
//	for i := 0; i < len(nums); i++ {
//		tmp := nums[i]
//		index := i + 1
//		for index < len(nums) {
//			if tmp != k {
//				tmp += nums[index]
//				index++
//			}
//			if tmp == k {
//				res++
//			}
//		}
//
//	}
//
//	return res
//}
