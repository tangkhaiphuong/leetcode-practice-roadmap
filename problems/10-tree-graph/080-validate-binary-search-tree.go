//go:build ignore

// Problem 080: Validate Binary Search Tree (LeetCode #98)
// Difficulty: Medium
// Categories: Tree, DFS, Inorder
//
// Description:
//   Given the root of a binary tree, determine if it is a valid BST: for
//   every node, all left descendants < node.val < all right descendants.
//
// Examples:
//      2          [Valid]      5
//     / \                     / \
//    1   3                   1   4   [Invalid: 3 < 5 in right subtree]
//                                / \
//                               3   6
//
// Approach: Range-Bound DFS
//   Pass (low, high) bounds. Each node must satisfy low < val < high. Recurse
//   with (low, val) for left and (val, high) for right.
//
//   Alternative: in-order traversal must produce strictly increasing values.
//
// Complexity:
//   Time:  O(n)
//   Space: O(h)

package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var dfs func(*TreeNode, int, int) bool
	dfs = func(n *TreeNode, lo, hi int) bool {
		if n == nil {
			return true
		}
		if n.Val <= lo || n.Val >= hi {
			return false
		}
		return dfs(n.Left, lo, n.Val) && dfs(n.Right, n.Val, hi)
	}
	return dfs(root, math.MinInt64, math.MaxInt64)
}

func main() {
	tree := &TreeNode{2, &TreeNode{Val: 1}, &TreeNode{Val: 3}}
	fmt.Println(isValidBST(tree)) // true
	bad := &TreeNode{5, &TreeNode{Val: 1}, &TreeNode{4, &TreeNode{Val: 3}, &TreeNode{Val: 6}}}
	fmt.Println(isValidBST(bad)) // false
}
