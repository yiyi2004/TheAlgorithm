package main

func main() {

}

// 二分查找

// 滑动窗口
func slidingWindow(nums []int) {
	for l, r := 0, 0; r < len(nums); r++ {
		for l <= r && check() {
			l++
		}
	}

}

// 单调栈

type stack struct {
}

func (s *stack) Top() int   { return 0 }
func (s *stack) Pop() int   { return 0 }
func (s *stack) Push(v int) {}
func singleStack(nums []int) {
	s := new(stack)
	for i := 0; i < len(nums); i++ {
		for check() && nums[s.Top()] < nums[i] {
			bse
			top := s.Pop()
			_ = top
			// nums[top] 的下一个更大的元素是 num[i]
		}
		s.Push(i)
	}
}

func check() bool { return true }

// 二分朝朝模板

// left

func bsec2Left(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) >> 1
		if checkMid(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}

func bsec2Right(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r + 1) >> 1
		if checkMid(mid) {
			l = mid
		} else {
			r = mid - 1
		}

	}
	return r
}

func checkMid(mid int) bool { return true }

// 最短路算法
// 动态规划 BFS(迪杰斯特拉算法)
// 不成环的

// 任意两个节点之间的权重是一样的
type access struct {
	node string
	w    int
}

func bfs1() {
	//queue := make([]access, 0)
	//queue = append(queue)
	// 1. 将起始点放进队列
	// 2. for  queue != empty
	// 3. 拿出一个节点，如果这个节点是目标节点，返回
	// 4. 如果不是，遍历他的孩子节点
	// 5. 如果孩子节点没被访问过，加入队列，设置为被访问过，步长是 w+1
	// 附加，如果带有权重，那么设置优先级队列
}
