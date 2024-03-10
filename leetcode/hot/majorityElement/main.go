package main

func main() {

}

func majorityElement(nums []int) (ans int) {
	cnt := 0
	for _, v := range nums {
		if v == ans {
			cnt++
		} else if cnt == 0 {
			ans = v
		} else {
			cnt--
		}
	}

	return
}
