package main

import (
	"fmt"
)

// dfs currently prints dfs traversal of a structure
// -- currently just implemented to traverse binary trees
func dfs(root *TreeNode) {
	stack := []*TreeNode{root}

	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node == nil {
			continue
		}

		stack = append(stack, node.Right)
		stack = append(stack, node.Left)
		fmt.Printf("Visitng node! Val:\n%v\n", node.Val)
	}

	fmt.Println("Traversal complete!")
	return
}
