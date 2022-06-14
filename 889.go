package main

func main() {

}

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {

	idxMap := make(map[int]int, len(postorder))

	for k, v := range postorder {
		idxMap[v] = k
	}
	return build(0, 0, len(preorder), preorder, idxMap)
}

func build(preStart int, postStart int, N int, preOrder []int, postOrderIndex map[int]int) *TreeNode {
	if N == 0 {
		return nil
	}

	root := &TreeNode{Val: preOrder[preStart]}
	if N == 1 {
		return root
	}

	leftNode := preOrder[preStart+1]

	L := postOrderIndex[leftNode] - postStart + 1
	root.Left = build(preStart+1, postStart, L, preOrder, postOrderIndex)
	root.Right = build(preStart+L+1, postStart+L, N-1-L, preOrder, postOrderIndex)
	return root
}
