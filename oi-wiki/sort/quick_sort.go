package main

//func quickSort(arr []int) []int {
//	if len(arr) <= 1 {
//		return arr
//	}
//
//	pivot := arr[0] // 选择第一个元素作为基准
//	var left, right []int
//
//	for _, v := range arr[1:] {
//		if v <= pivot {
//			left = append(left, v)
//		} else {
//			right = append(right, v)
//		}
//	}
//
//	left = quickSort(left)
//	right = quickSort(right)
//
//	return append(append(left, pivot), right...)
//}

func basicQsort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[0]
	var left, right []int
	num := 0
	for _, v := range arr[1:] {
		if v < pivot {
			left = append(left, v)
		} else if v > pivot {
			right = append(right, v)
		} else {
			num++
		}
	}

	left = basicQsort(left)
	right = basicQsort(right)
	return append(append(left, pivot), right...)
}

// 朴素优化思路
// 1. begin, end, mid 的中位数
// 2. 序列较短的情况下，使用插入排序，效率更高
// 3. 每趟排序后，将与分界元素相等的元素聚集在分界元素周围
// 三路快速排序 = 快速排序 + 基数排序
