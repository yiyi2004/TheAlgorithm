package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	info := process(root)
	return maxV(info.maxHeadVal, info.maxNoneHeadVal)
}

type info struct {
	maxHeadVal     int
	maxNoneHeadVal int
}

func process(root *TreeNode) info {
	if root == nil {
		return info{
			maxHeadVal:     0,
			maxNoneHeadVal: 0,
		}
	}
	leftInfo := process(root.Left)
	rightInfo := process(root.Right)
	res := info{}
	// 不放 head
	res.maxNoneHeadVal = maxV(leftInfo.maxHeadVal, leftInfo.maxNoneHeadVal) + maxV(rightInfo.maxHeadVal, rightInfo.maxNoneHeadVal)
	// 放 head
	res.maxHeadVal = leftInfo.maxNoneHeadVal + root.Val + rightInfo.maxNoneHeadVal
	return res
}

func maxV(arr ...int) int {
	res := arr[0]
	for _, v := range arr {
		if v > res {
			res = v
		}
	}
	return res
}
