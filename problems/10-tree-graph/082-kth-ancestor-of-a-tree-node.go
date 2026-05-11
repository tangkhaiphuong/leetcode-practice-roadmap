//go:build ignore

// Problem 082: Kth Ancestor of a Tree Node (LeetCode #1483)
// Difficulty: Hard
// Categories: Tree, Binary Lifting, Design
//
// Description:
//   Tree of n nodes (0..n-1) with parent[]; parent[0] = -1. Implement
//   getKthAncestor(node, k): the k-th ancestor of node, or -1.
//
// Approach: Binary Lifting
//   Precompute up[node][j] = the 2^j-th ancestor.
//     up[node][0] = parent[node]
//     up[node][j] = up[ up[node][j-1] ][j-1]
//   For a query (node, k), iterate over bits of k; for each set bit j, jump
//   node = up[node][j]. If node becomes -1, return -1.
//
// Complexity:
//   Build: O(n log n)
//   Query: O(log k)

package main

import "fmt"

type TreeAncestor struct {
	up    [][]int
	logK  int
}

func ConstructorTA(n int, parent []int) TreeAncestor {
	logK := 1
	for (1 << logK) < n {
		logK++
	}
	up := make([][]int, n)
	for i := range up {
		up[i] = make([]int, logK+1)
		up[i][0] = parent[i]
	}
	for j := 1; j <= logK; j++ {
		for i := 0; i < n; i++ {
			if up[i][j-1] == -1 {
				up[i][j] = -1
			} else {
				up[i][j] = up[up[i][j-1]][j-1]
			}
		}
	}
	return TreeAncestor{up: up, logK: logK}
}

func (t *TreeAncestor) GetKthAncestor(node, k int) int {
	for j := 0; j <= t.logK && node != -1; j++ {
		if k>>j&1 == 1 {
			node = t.up[node][j]
		}
	}
	return node
}

func main() {
	t := ConstructorTA(7, []int{-1, 0, 0, 1, 1, 2, 2})
	fmt.Println(t.GetKthAncestor(3, 1)) // 1
	fmt.Println(t.GetKthAncestor(5, 2)) // 0
	fmt.Println(t.GetKthAncestor(6, 3)) // -1
}
