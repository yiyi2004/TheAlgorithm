package Class02

// 输入一定能够保证，数组中所有的数都出现了M次，只有一种数出现了K次
// 1 <= K < M
// 返回这种数
func km(arr []int, k, m int) int {
	var tmp [32]int
	var ans int

	for _, value := range arr {
		for i := 0; i < 32; i++ {
			if (value>>i)&1 != 0 {
				tmp[i] += 1
			}
		}
	}

	for i := 0; i < 32; i++ {
		if tmp[i]%k != 0 {
			ans |= 1 << i
		}
	}

	return ans
}
