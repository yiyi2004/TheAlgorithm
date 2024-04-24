package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	// var a int
	// fmt.Scan(&a)
	fmt.Printf("%s", "hello world")
	res := getMaxFromA(23121, []int{2, 4, 8})
	fmt.Println(res)
}

func getMaxFromA(n int, arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	sort.Ints(arr)
	s := strconv.Itoa(n)
	res := 0
	flag := false
	for i := 0; i < len(s); i++ {
		num, _ := strconv.Atoi(string(s[i]))
		if flag {
			res = res*10 + arr[len(arr)-1]
			continue
		}
		if num < arr[0] {
			continue
		} else {
			index := find(arr, num)
			if arr[index] != num {
				flag = true
			}
			res = res*10 + arr[index]
		}
	}
	return res
}

func find(arr []int, target int) int {
	l, r := 0, len(arr)-1
	mid := (l + r) >> 1
	for l < r {
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}
