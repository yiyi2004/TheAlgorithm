package main

func main() {

}

func combinationSum3(k int, n int) [][]int {
	res := make([][]int, 0)
	var dfs func(index int, sum int, path []int)
	dfs = func(index int, sum int, path []int) {
		if len(path) >= k && sum >= n {
			if len(path) > k || sum > n {
				return
			}
			res = append(res, append([]int{}, path...))
			return
		}
		for i := index; i <= 9; i++ {
			if sum+i > n {
				return
			}
			path = append(path, i)
			dfs(i+1, sum+i, path)
			path = path[:len(path)-1]
		}
	}
	dfs(1, 0, []int{})
	return res
}
