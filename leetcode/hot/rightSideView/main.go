package main

func main() {

}

// 刷题果然会让人上瘾，只是一开始会比较困难而已。
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 为什么这样写就是可以的了？？？
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	// 层序遍历中每一层最后一个元素
	var dfs func(node *TreeNode, depth int)
	res := make([]int, 0)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if depth == len(res) {
			res = append(res, node.Val)
		}
		dfs(node.Right, depth+1)
		dfs(node.Left, depth+1)
	}
	dfs(root, 0)
	return res
}

// 层序遍历的思路可以使用，但是写起来比较困难
//func levelOrder(root *TreeNode) []int {
//	return nil
//}

//func rightSideView(root *TreeNode) []int {
//	if root == nil {
//		return nil
//	}
//	res := make([]int, 0)
//	process(root, &res)
//	// 层序遍历中的最后一个元素
//	return res
//}
//
//func process(root *TreeNode, res *[]int) {
//	if root == nil {
//		return
//	}
//	*res = append(*res, root.Val)
//	if root.Right != nil {
//		process(root.Right, res)
//	} else {
//		process(root.Left, res)
//	}
//}
