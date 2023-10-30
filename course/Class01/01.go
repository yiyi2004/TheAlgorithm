package Class01

func selectionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		minIndex := i
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[minIndex], arr[i] = arr[i], arr[minIndex]
	}
}

func nearLeft(arr []int, value int) int {
	L := 0
	R := len(arr) - 1
	index := -1

	for L <= R {
		mid := (L + R) >> 1
		if arr[mid] >= value {
			index = mid
			R = mid - 1
		} else {
			L = mid + 1
		}
	}

	return index
}

func nearRight(arr []int, value int) int {
	L := 0
	R := len(arr) - 1
	index := -1
	for L <= R {
		mid := (L + R) >> 1
		if arr[mid] <= value {
			index = mid
			L = mid + 1
		} else {
			R = mid - 1
		}
	}
	return index
}

func getLessIndex(arr []int) int {
	if len(arr) == 0 || arr == nil {
		return -1
	}
	if len(arr) == 1 {
		return 0
	}

	L := 0
	R := len(arr) - 1

	if arr[0] < arr[1] {
		return 0
	}
	if arr[R] < arr[R-1] {
		return R
	}

	for L <= R {
		mid := (L + R) >> 1
		if arr[mid] > arr[mid-1] {
			R = mid - 1
		} else if arr[mid] > arr[mid+1] {
			L = mid + 1
		} else {
			return mid
		}
	}

	return -1
}
