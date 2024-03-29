package main

import (
	"fmt"
	"sort"
)

// https://leetcode.cn/problems/merge-intervals/?envType=study-plan-v2&envId=top-100-liked

func main() {
	intervals := [][]int{{1, 4}, {2, 3}}
	res := merge(intervals)
	fmt.Println(res)
}

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}
	if len(intervals) == 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var res [][]int
	left, right := intervals[0][0], intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		if i == len(intervals)-1 {
			if right < intervals[i][0] {
				res = append(res, []int{left, right})
				res = append(res, []int{intervals[i][0], intervals[i][1]})
			} else if right >= intervals[i][0] && right <= intervals[i][1] {
				res = append(res, []int{left, intervals[i][1]})
			} else {
				res = append(res, []int{left, right})
			}
			return res
		}

		if right < intervals[i][0] {
			res = append(res, []int{left, right})
			left, right = intervals[i][0], intervals[i][1]
		} else if right >= intervals[i][0] && right <= intervals[i][1] {
			right = intervals[i][1]
		}
	}

	return res
}

//func merge(intervals [][]int) [][]int {
//	if len(intervals) == 0 {
//		return nil
//	}
//	sort.Slice(intervals, func(i, j int) bool {
//		return intervals[i][0] < intervals[j][0]
//	})
//	res := make([][]int, 0)
//	begin, end := intervals[0][0], intervals[0][1]
//	for i := 0; i < len(intervals); i++ {
//		if i == len(intervals)-1 {
//			res = append(res, []int{begin, end})
//			return res
//		}
//
//		if end >= intervals[i+1][0] && end <= intervals[i+1][1] {
//			end = intervals[i+1][1]
//		} else if end > intervals[i+1][1] {
//			continue
//		} else {
//			res = append(res, []int{begin, end})
//			begin = intervals[i+1][0]
//			end = intervals[i+1][1]
//		}
//	}
//
//	return res
//}
