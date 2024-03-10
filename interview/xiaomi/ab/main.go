package main

import "fmt"

func main() {
	arr := []int{1, 2, 4, 5, 3, 6, 10, 11, -1, 0}
	res := foo(arr)
	fmt.Println(res)
}

func foo(arr []int) []int {
	res := mergeSort(arr, 0, len(arr)-1)
	for i := 0; i < len(res); i = i + 2 {
		res[i], res[i+1] = res[i+1], res[i]
	}
	return res
}

func mergeSort(arr []int, left, right int) []int {
	if right-left < 1 {
		return []int{arr[left]}
	}
	mid := (left + right) >> 1
	leftArr := mergeSort(arr, left, mid)
	rightArr := mergeSort(arr, mid+1, right)
	return merge(leftArr, rightArr)
}

func merge(leftArr, rightArr []int) []int {
	res := make([]int, 0, len(leftArr)+len(rightArr))
	l, r := 0, 0
	for l < len(leftArr) && r < len(rightArr) {
		if leftArr[l] < rightArr[r] {
			res = append(res, leftArr[l])
			l++
		} else {
			res = append(res, rightArr[r])
			r++
		}
	}

	for l < len(leftArr) {
		res = append(res, leftArr[l])
		l++
	}
	for r < len(rightArr) {
		res = append(res, rightArr[r])
		r++
	}
	return res
}
