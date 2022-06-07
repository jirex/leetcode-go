package main

func main() {

}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

// max ...
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

