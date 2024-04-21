package main

func main() {

}

func combinationSum3(k int, n int) [][]int {
	res := make([][]int, 0)
	var dfs func(index int, preArr []int, preVal int)
	dfs = func(index int, preArr []int, preVal int) {
		if len(preArr) == k && preVal == n {
			tmp := make([]int, len(preArr))
			copy(tmp, preArr)
			res = append(res, tmp)
			return
		}

		for i := index + 1; i <= 9; i++ {
			preVal += i
			preArr = append(preArr, i)
			dfs(i, preArr, preVal)
			preArr = preArr[:len(preArr)-1]
			preVal -= i
		}
	}
	dfs(0, []int{}, 0)
	return res
}
