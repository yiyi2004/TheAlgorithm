package main

import "fmt"

func main() {
	arr := []int{0, 0}
	consecutive := longestConsecutive(arr)
	fmt.Println(consecutive)
}

// AC
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	hm := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		hm[nums[i]] = true
	}
	maxLong := 0
	for num := range hm {
		if !hm[num-1] {
			maxTmpLong := 1
			curr := num + 1
			for hm[curr] {
				maxTmpLong++
				curr++
			}
			if maxTmpLong > maxLong {
				maxLong = maxTmpLong
			}
		}
	}
	return maxLong
}

func process(m map[int]int, val int) int {
	if _, ok := m[val]; !ok {
		return 0
	} else {
		return process(m, val-1) + 1
	}
}

//// 69 / 73 个通过的测试用例
//func longestConsecutive(nums []int) int {
//	if len(nums) == 0 {
//		return 0
//	}
//	hm := make(map[int]int)
//	for i := 0; i < len(nums); i++ {
//		hm[nums[i]] = i
//	}
//
//	maxLong := 0
//	for k, _ := range hm {
//		tmp := process(hm, k)
//		if tmp > maxLong {
//			maxLong = tmp
//		}
//	}
//	return maxLong
//}
//
//func process(m map[int]int, val int) int {
//	if _, ok := m[val]; !ok {
//		return 0
//	} else {
//		return process(m, val-1) + 1
//	}
//}

// 超出时间限制
//func longestConsecutive(nums []int) int {
//	if len(nums) == 0 {
//		return 0
//	}
//	maxValue := nums[0]
//	minValue := nums[0]
//	hm := make(map[int]int)
//	for i := 0; i < len(nums); i++ {
//		if nums[i] > maxValue {
//			maxValue = nums[i]
//		}
//		if nums[i] < minValue {
//			minValue = nums[i]
//		}
//		hm[nums[i]] = i
//	}
//
//	maxLong := 0
//	tmp := 0
//	for i := minValue; i <= maxValue; i++ {
//		if _, ok := hm[i]; ok {
//			tmp++
//		} else {
//			if tmp > maxLong {
//				maxLong = tmp
//			}
//			tmp = 0
//		}
//		if i == maxValue {
//			if tmp > maxLong {
//				maxLong = tmp
//			}
//		}
//	}
//
//	return maxLong
//}
