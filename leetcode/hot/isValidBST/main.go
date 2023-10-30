package main

func main() {}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 左边的最小值要比 val 大
type info struct {
	minVal int
	maxVal int
	isBST  bool
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return false
	}
	return process(root).isBST
}

func process(root *TreeNode) *info {
	if root == nil {
		return nil
	}

	leftInfo := process(root.Left)
	rightInfo := process(root.Right)

	var isBST bool
	res := new(info)
	if leftInfo != nil {
		res.minVal = leftInfo.minVal
	} else {
		res.minVal = root.Val
	}
	if rightInfo != nil {
		res.maxVal = rightInfo.maxVal
	} else {
		res.maxVal = root.Val
	}

	//
	if leftInfo == nil && rightInfo == nil {
		isBST = true
	} else if leftInfo == nil && rightInfo != nil && rightInfo.isBST && rightInfo.minVal > root.Val {
		isBST = true
	} else if leftInfo != nil && rightInfo == nil && leftInfo.isBST && leftInfo.maxVal < root.Val {
		isBST = true
	} else if leftInfo != nil && rightInfo != nil && leftInfo.isBST && rightInfo.isBST && leftInfo.maxVal < root.Val && rightInfo.minVal > root.Val {
		isBST = true
	}
	res.isBST = isBST
	return res
}
