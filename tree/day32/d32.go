package main

import "math"

func main() {}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	finalAns := math.Max(float64(maxDepth(root.Left)), float64(maxDepth(root.Right))) + 1
	return int(finalAns)
}
