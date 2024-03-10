package main

func main() {
	root := TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param pRoot TreeNode类
 * @return int整型二维数组
 */
func Print(pRoot *TreeNode) [][]int {
	// write code here
	var res [][]int
	height := 1
	queue := []*TreeNode{pRoot}

	for len(queue) != 0 {
		if height%2 == 1 {
			// height 为奇数
			tmp := make([]int, 0)
			for i := 0; i < len(queue); i++ {
				tmp = append(tmp, queue[i].Val)
			}
			res = append(res, tmp)
			nNum := len(queue)

			for i := 0; i < len(queue); i++ {
				if queue[i].Left != nil {
					queue = append(queue, queue[i].Left)
				}
				if queue[i].Right != nil {
					queue = append(queue, queue[i].Right)
				}
			}
			queue = queue[nNum-1:]
		} else {
			// height 为偶数
			tmp := make([]int, 0)
			for i := len(queue) - 1; i >= 0; i-- {
				tmp = append(tmp, queue[i].Val)
			}
			res = append(res, tmp)
			nNum := len(queue)

			for i := 0; i < len(queue); i++ {
				if queue[i].Left != nil {
					queue = append(queue, queue[i].Left)
				}
				if queue[i].Right != nil {
					queue = append(queue, queue[i].Right)
				}
			}
			queue = queue[nNum-1:]
		}
		height++
	}
	return res
}
