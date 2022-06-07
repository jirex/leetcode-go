package main

func main() {

}

func preorderTraversal(root *TreeNode) []int {
	re := []int{}
	var pre func(*TreeNode)
	pre = func(node *TreeNode) {
		if node == nil {
			return
		}
		re = append(re, node.Val)
		pre(node.Left)
		pre(node.Right)
	}

	pre(root)
	return re
}
