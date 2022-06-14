package main

func main() {
	
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	traverse(root.Left,root.Right)
	return root
}

func traverse(node1, node2 *Node) {
	if node1 == nil || node2 == nil {
		return;
	}

	node1.Next = node2

	traverse(node1.Left, node1.Right)
	traverse(node2.Left, node2.Right)
	traverse(node1.Right,node2.Left)
}
