package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	res := GetNextHigh([]int{73, 74, 75, 71, 69, 72, 76, 73})
	fmt.Println(res)
}

func GetNextHigh(temperatures []int) []int {
	if len(temperatures) == 0 {
		return nil
	}
	stack := make([]int, 0, len(temperatures))
	res := make([]int, len(temperatures))

	stack = append(stack, 0)
	index := 0
	for index < len(temperatures) {
		if len(stack) == 0 {
			stack = append(stack, index)
			index++
			continue
		}
		if temperatures[index] > temperatures[stack[len(stack)-1]] {
			for temperatures[index] > temperatures[stack[len(stack)-1]] {
				tmpIndex := stack[len(stack)-1]
				res[tmpIndex] = index - tmpIndex
				stack = stack[:len(stack)-1]
				if len(stack) <= 0 {
					break
				}
			}
			stack = append(stack, index)
		} else {
			stack = append(stack, index)
		}
		index++
	}
	return res
}
