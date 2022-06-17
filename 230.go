package main

func main() {

}

func kthSmallest(root *TreeNode, k int) int {
	rank := 0
	res := 0

	var traverse func(*TreeNode, int)
	traverse = func(n *TreeNode, k int) {
		if n == nil {
			return
		}

		traverse(n.Left, k)
		rank++
		if k == rank {
			res = n.Val
			return
		}
		traverse(n.Right, k)

	}
	traverse(root, k)
	return res
}
