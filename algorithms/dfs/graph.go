package main

// Binary Tree Implementation
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
BinaryTreeFromNestedSlices returns the root of a binary
tree which is created assuming each index in the provided
slice is equal to the level of the node. From left to right
the nodes are assigned parents

E.g: [[1], [2, 3], [4, 5, 6]] is equivalent to...

                  1
                /   \
               2     3
             /  \   /
            4   5  6

And the root TreeNode (1) is returned.
*/
func BinaryTreeFromNestedSlices(ns [][]int) *TreeNode {
	levelMap := buildLevelMap(ns)

	for lvl := 1; lvl < len(ns); lvl++ {
		parents := levelMap[lvl-1]
		children := levelMap[lvl]

		for i := 0; i < len(children); i += 2 {
			parent := parents[i/2]
			parent.Left = children[i]
			if (i + 1) < len(children) {
				parent.Right = children[i+1]
			}
		}
	}

	return levelMap[0][0]
}

func buildLevelMap(s [][]int) map[int][]*TreeNode {
	levelMap := make(map[int][]*TreeNode, len(s))
	for lvl, vals := range s {
		levelNodes := make([]*TreeNode, len(vals))
		for i, v := range vals {
			levelNodes[i] = &TreeNode{
				Val: v,
			}
		}
		levelMap[lvl] = levelNodes
	}

	return levelMap
}
