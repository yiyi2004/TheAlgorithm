package main

// https://leetcode.cn/problems/binary-tree-level-order-traversal/?envType=study-plan-v2&envId=top-100-liked
// 层序遍历如何标记层数
// BFS ---> 层序遍历和最短路径
// 你需要更为深刻的理解这些算法

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type queue []*TreeNode

// this is an easy way to implement the queue.
func (q *queue) push(val *TreeNode) {
	*q = append(*q, val)
}

func (q *queue) pop() *TreeNode {
	val := (*q)[0]
	*q = (*q)[1:]
	return val
}

func (q *queue) isEmpty() bool {
	return len(*q) == 0
}

func (q *queue) size() int {
	return len(*q)
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	var q queue
	q.push(root)
	for !q.isEmpty() {
		n := q.size()
		var level []int
		for i := 0; i < n; i++ {
			tmp := q.pop()
			if tmp.Left != nil {
				q.push(tmp.Left)
			}
			if tmp.Right != nil {
				q.push(tmp.Right)
			}
			level = append(level, tmp.Val)
		}
		res = append(res, level)
	}
	return res
}

func levelOrderBasic(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	var q queue
	q.push(root)
	for !q.isEmpty() {
		tmp := q.pop()
		if tmp.Left != nil {
			q.push(tmp.Left)
		}
		if tmp.Right != nil {
			q.push(tmp.Right)
		}
		res = append(res, tmp.Val)
	}
	return res
}
