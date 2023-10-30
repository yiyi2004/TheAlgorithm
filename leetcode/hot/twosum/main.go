package main

import "fmt"

func main() {
	nums := []int{1, 2, 4, 5, 6, 7, 3}
	target := 3
	res := twoSum(nums, target)
	fmt.Println(res)
}

func twoSum(nums []int, target int) []int {
	hashTable := make(map[int]int)

	for index, value := range nums {
		if v, ok := hashTable[target-value]; ok {
			return []int{v, index}
		}
		hashTable[value] = index
	}

	return []int{}
}
