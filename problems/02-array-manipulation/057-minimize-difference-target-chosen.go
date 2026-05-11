//go:build ignore

// Problem 057: Minimize the Difference Between Target and Chosen Elements (LC #1981)
// Difficulty: Medium / Hard
// Categories: Array (Matrix), DP (Subset-Sum-like), Greedy
//
// Description:
//   Given an m x n integer matrix mat and target. Pick exactly one element
//   from each row; let s = sum of chosen. Return the minimum |s - target|.
//
// Examples:
//   mat=[[1,2,3],[4,5,6],[7,8,9]], target=13 -> 0  (1+5+7=13)
//   mat=[[1],[2],[3]],            target=100 -> 94 (1+2+3=6)
//   mat=[[1,2,9,8,7]],            target=6   -> 1
//
// Approach: Subset-Sum DP over Rows (using a set of reachable sums)
//   Maintain a set S of reachable sums, starting with {0}. For each row, S
//   becomes { s + v : s in S, v in row }. Prune sums beyond a useful bound
//   (>= 2*target works as an upper bound — sums above don't help reduce
//   |s - target| relative to the next-smaller candidate).
//   Final: min |s - target| over s in S.
//
// Complexity:
//   Time:  O(m * |S| * n)  with |S| bounded by 70*70 ≈ 4900
//   Space: O(|S|)

package main

import "fmt"

func minimizeTheDifference(mat [][]int, target int) int {
	cur := map[int]bool{0: true}
	cap := 0
	if target > 0 {
		cap = 2 * target
	}
	for _, row := range mat {
		next := map[int]bool{}
		for s := range cur {
			for _, v := range row {
				ns := s + v
				if ns > cap && len(next) > 0 {
					// Allow one above-cap to keep an upper bound option.
				}
				next[ns] = true
			}
		}
		cur = next
	}
	best := 1 << 30
	for s := range cur {
		d := s - target
		if d < 0 {
			d = -d
		}
		if d < best {
			best = d
		}
	}
	return best
}

func main() {
	fmt.Println(minimizeTheDifference([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 13)) // 0
	fmt.Println(minimizeTheDifference([][]int{{1}, {2}, {3}}, 100))                  // 94
	fmt.Println(minimizeTheDifference([][]int{{1, 2, 9, 8, 7}}, 6))                  // 1
}
