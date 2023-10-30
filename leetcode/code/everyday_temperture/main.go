package main

import (
	"fmt"
)

type Stack []int

func (s *Stack) Push(val int) {
	*s = append(*s, val)
}

func (s *Stack) Pop() int {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}

func (s *Stack) Peek() int {
	return (*s)[len(*s)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func main() {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	res := dailyTemperatures(temperatures)
	fmt.Println(res)
}

func dailyTemperatures(temperatures []int) []int {
	if temperatures == nil || len(temperatures) == 0 {
		return nil
	}
	res := make([]int, len(temperatures))

	var s Stack
	for i := 0; i < len(temperatures); i++ {
		if s.IsEmpty() {
			s.Push(i)
			continue
		}
		if temperatures[i] > temperatures[s.Peek()] {
			for !s.IsEmpty() && temperatures[i] > temperatures[s.Peek()] {
				indexPop := s.Pop()
				res[indexPop] = i - indexPop
			}
			s.Push(i)
		} else {
			s.Push(i)
		}
	}
	return res
}

// 总是一遍又一遍的调试，会有各种各样的麻烦捏。

// 每日温度问题
