package Class04

func mergesort1(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	process(arr, 0, len(arr)-1)
}

func process(arr []int, l int, r int) {
	if l == r {
		return
	}
	m := (l + r) / 2
	process(arr, l, m)
	process(arr, m+1, r)
	merge(arr, l, m, r)
}

func merge(arr []int, l, m, r int) {
	if l == r {
		return
	}
	tmp := make([]int, r-l+1)
	i, j := l, m+1
	index := 0
	for i <= m && j <= r {
		if arr[i] < arr[j] {
			tmp[index] = arr[i]
			i++
		} else {
			tmp[index] = arr[j]
			j++
		}
		index++
	}
	for i <= m {
		tmp[index] = arr[i]
		index++
		i++
	}
	for j <= r {
		tmp[index] = arr[j]
		index++
		j++
	}
	index = 0
	for k := l; k <= r; k++ {
		arr[k] = tmp[index]
		index++
	}
}

// 非递归实现
func mergesort2(arr []int) {
	// 大概是用步长来实现的
}
