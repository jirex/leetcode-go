package main

func main() {

}

func diameterOfBinaryTree(root *TreeNode) int {
	ans := 0

	if root == nil {
		return ans
	}

	var depth func(*TreeNode) int
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l := depth(node.Left)
		r := depth(node.Right)

		ans = max(ans, l+r)

		return max(l, r) + 1
	}

	depth(root)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
