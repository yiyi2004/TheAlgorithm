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
	res := process(root)
	return maxv(res.isHeadMaxVal, res.noneHeadMaxVal)
}

type info struct {
	isHeadMaxVal   int
	noneHeadMaxVal int
}

func process(root *TreeNode) info {
	if root == nil {
		return info{0, 0}
	}
	res := info{}
	leftInfo := process(root.Left)
	rightInfo := process(root.Right)
	res.isHeadMaxVal = root.Val + leftInfo.noneHeadMaxVal + rightInfo.noneHeadMaxVal
	res.noneHeadMaxVal = maxv(leftInfo.isHeadMaxVal, leftInfo.noneHeadMaxVal) + maxv(rightInfo.isHeadMaxVal, rightInfo.noneHeadMaxVal)
	return res
}

func maxv(v1, v2 int) int {
	if v1 < v2 {
		return v2
	}
	return v1
}
