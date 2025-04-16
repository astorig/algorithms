package main

import "fmt"

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	child := &TreeNode{Value: 9}
	child3 := &TreeNode{Value: 15}
	child4 := &TreeNode{Value: 7}
	child1 := &TreeNode{Value: 20, Left: child3, Right: child4}
	root := &TreeNode{Value: 3, Left: child, Right: child1}

	result := maxDepth(root)

	fmt.Println(result)
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	depth := 1

	if root.Left != nil {
		depth = max(depth, maxDepth(root.Left)+1)
	}

	if root.Right != nil {
		depth = max(depth, maxDepth(root.Right)+1)
	}

	return depth
}
