package Class04

func mergeSort1(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	process1(arr, 0, len(arr)-1)
}

func process1(arr []int, L, R int) {
	if L == R {
		return
	}
	mid := (L + R) >> 1
	process1(arr, L, mid)
	process1(arr, mid+1, R)
	merge1(arr, L, mid, R)
}

func merge1(arr []int, L, mid, R int) {
	help := make([]int, R-L+1)
	i := L
	j := mid + 1
	index := 0
	for i <= mid && j <= R {
		if arr[i] < arr[j] {
			help[index] = arr[i]
			index++
			i++
		} else {
			help[index] = arr[j]
			index++
			j++
		}
	}
	for i <= mid {
		help[index] = arr[i]
		index++
		i++
	}
	for j <= R {
		help[index] = arr[j]
		index++
		j++
	}
	index = 0
	for i := L; i <= R; i++ {
		arr[i] = help[index]
		index++
	}
}

// 非递归没有实现捏
func mergeSort2(arr []int) {

}

func smallSum(arr []int) int {
	if arr == nil || len(arr) < 2 {
		return 0
	}
	return processSmallSum(arr, 0, len(arr)-1)
}

func processSmallSum(arr []int, L, R int) int {
	if R == L {
		return 0
	}
	mid := (R + L) >> 1
	return processSmallSum(arr, L, mid) + processSmallSum(arr, mid+1, R) + mergeSmallSum(arr, L, mid, R)
}

func mergeSmallSum(arr []int, L, mid, R int) int {
	i := L
	j := mid + 1
	index := 0
	result := 0
	help := make([]int, R-L+1)

	for i <= mid && j <= R {
		if arr[i] < arr[j] {
			result += arr[i] * (R - j + 1)
			help[index] = arr[i]
			index++
			i++
		} else {
			help[index] = arr[j]
			index++
			j++
		}
	}

	for i <= mid {
		help[index] = arr[i]
		index++
		i++
	}
	for j <= R {
		help[index] = arr[j]
		index++
		j++
	}
	index = 0
	for k := L; k <= R; k++ {
		arr[k] = help[index]
		index++
	}

	return result
}

func reversePair(arr []int) int {
	if arr == nil || len(arr) < 2 {
		return 0
	}
	return processReversePair(arr, 0, len(arr)-1)
}

func processReversePair(arr []int, L, R int) int {
	if L == R {
		return 0
	}
	mid := (L + R) >> 1
	return processReversePair(arr, L, mid) +
		processReversePair(arr, mid+1, R) +
		mergeReversePair(arr, L, mid, R)
}

func mergeReversePair(arr []int, L, mid, R int) int {
	i := L
	j := mid + 1
	index := 0
	result := 0
	help := make([]int, R-L+1)

	for i <= mid && j <= R {
		if arr[i] > arr[j] {
			result += mid - i + 1
			help[index] = arr[j]
			index++
			j++
		} else {
			help[index] = arr[i]
			index++
			i++
		}
	}

	for i <= mid {
		help[index] = arr[i]
		index++
		i++
	}
	for j <= R {
		help[index] = arr[j]
		index++
		j++
	}

	index = 0
	for k := L; k <= R; k++ {
		arr[k] = help[index]
		index++
	}

	return result
}

func biggerThanRightTwice(arr []int) int {
	if arr == nil || len(arr) < 2 {
		return 0
	}
	return processBiggerThan(arr, 0, len(arr)-1)
}

func processBiggerThan(arr []int, L, R int) int {
	if L == R {
		return 0
	}
	mid := (L + R) >> 1
	return processBiggerThan(arr, L, mid) +
		processBiggerThan(arr, mid+1, R) +
		mergeBiggerThan(arr, L, mid, R)
}

// 暂时不想写了，明天再写吧
func mergeBiggerThan(arr []int, L, mid, R int) int {
	return 0
}
