package main

import "fmt"

func main() {
	arr := []int{73, 74, 75, 71, 69, 72, 76, 73}
	res := dailyTemperatures(arr)
	fmt.Println(res)
}

//
//type Stack []int
//
//func (s *Stack) Push(val int) {
//	*s = append(*s, val)
//}
//
//func (s *Stack) Pop() int {
//	res := (*s)[len(*s)-1]
//	*s = (*s)[:len(*s)-1]
//	return res
//}
//
//func (s *Stack) Peek() int {
//	return (*s)[len(*s)-1]
//}
//
//func (s *Stack) IsEmpty() bool {
//	return len(*s) == 0
//}
//
//func dailyTemperatures(temperatures []int) []int {
//	if temperatures == nil || len(temperatures) == 0 {
//		return nil
//	}
//	res := make([]int, len(temperatures))
//
//	var s Stack
//	for i := 0; i < len(temperatures); i++ {
//		if s.IsEmpty() {
//			s.Push(i)
//			continue
//		}
//		if temperatures[i] > temperatures[s.Peek()] {
//			for !s.IsEmpty() && temperatures[i] > temperatures[s.Peek()] {
//				indexPop := s.Pop()
//				res[indexPop] = i - indexPop
//			}
//			s.Push(i)
//		} else {
//			s.Push(i)
//		}
//	}
//	return res
//}

func dailyTemperatures(temperatures []int) []int {
	stk := make([]int, 0)
	res := make([]int, len(temperatures))
	stk = append(stk, 0)
	index := 1
	for index < len(temperatures) {
		if temperatures[index] == temperatures[stk[len(stk)-1]] {
			stk = append(stk, index)
		} else if temperatures[index] < temperatures[stk[len(stk)-1]] {
			stk = append(stk, index)
		} else if temperatures[index] > temperatures[stk[len(stk)-1]] {
			for len(stk) > 0 && temperatures[index] > temperatures[stk[len(stk)-1]] {
				res[stk[len(stk)-1]] = index - stk[len(stk)-1]
				stk = stk[:len(stk)-1]
			}
			stk = append(stk, index)
		}
		index++
	}
	return res
}
