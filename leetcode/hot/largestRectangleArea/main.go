package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 1, 5, 6, 2, 3}
	res := largestRectangleArea(nums)
	fmt.Println(res)
}

// 通过单调栈解决问题，找左边第一个最小的 + 右边第一个最小的。
// 最后得到一个总体最大的。

// 实现了，但是性能好差

type stack struct {
	stack []int
}

func newStack() *stack {
	return &stack{
		stack: make([]int, 0),
	}
}

func (s *stack) push(ele int) {
	s.stack = append(s.stack, ele)
}
func (s *stack) peek() int {
	return s.stack[len(s.stack)-1]
}
func (s *stack) pop() int {
	v := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return v
}
func (s *stack) isEmpty() bool {
	return len(s.stack) == 0
}

func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	resMap := make(map[int]int)
	s := newStack()
	heights = append([]int{0}, append(heights, 0)...)
	for i := 0; i < len(heights); i++ {
		if s.isEmpty() {
			s.push(i)
			continue
		}
		if heights[i] >= heights[s.peek()] {
			s.push(i)
		} else {
			for heights[i] < heights[s.peek()] {
				right := i
				midIndex := s.pop()
				left := s.peek()
				for left == midIndex {
					left = s.pop()
				}
				resMap[midIndex] = (right - left - 1) * heights[midIndex]
			}
			s.push(i)
		}
	}
	max := 0
	for _, v := range resMap {
		if v > max {
			max = v
		}
	}
	return max
}

//
//func largestRectangleArea(heights []int) int {
//	heights = append([]int{0}, append(heights, 0)...)
//	stk := make([]int, 0)
//
//}
//
//func findMax(arr []int) int {
//	mx := arr[0]
//	for i := 1; i < len(arr); i++ {
//		if arr[i] > mx {
//			mx = arr[i]
//		}
//
//	}
//	return mx
//}
