//go:build ignore

// Problem 075: Minimum Space Wasted From Packaging (LeetCode #1889)
// Difficulty: Hard
// Categories: Binary Search, Sorting, Prefix Sum
//
// Description:
//   You have packages[] of sizes. Suppliers each provide an array of box
//   sizes (one size per box, infinite supply). Each package goes in one box
//   that fits (box size >= package size). Wasted space = sum(box - package).
//   Choose ONE supplier minimizing total wasted space. If impossible (the
//   supplier's max box < some package), skip them. Return min waste mod 1e9+7,
//   or -1 if no supplier works.
//
// Approach: Sort + Prefix Sum + Binary Search per Box
//   Sort packages. For each supplier:
//     - Sort boxes.
//     - If maxBox < maxPackage, skip.
//     - For each box b in sorted order, find via binary search the count of
//       packages still un-boxed that are <= b. Subtract their sum (prefix
//       sum) from b * count to get waste contribution. Advance pointer.
//
// Complexity:
//   Time:  O(N log N + sum(M_s log N)) where N=#packages, M_s=#boxes per supplier
//   Space: O(N)

package main

import (
	"fmt"
	"sort"
)

const mod75 = 1_000_000_007

func minWastedSpace(packages []int, boxes [][]int) int {
	sort.Ints(packages)
	n := len(packages)
	prefix := make([]int64, n+1)
	for i, p := range packages {
		prefix[i+1] = prefix[i] + int64(p)
	}
	totalPackage := prefix[n]
	const INF = int64(-1)
	best := INF
	for _, bs := range boxes {
		sort.Ints(bs)
		if bs[len(bs)-1] < packages[n-1] {
			continue
		}
		var totalBox int64
		prev := 0
		for _, b := range bs {
			j := sort.SearchInts(packages, b+1) // first index with package > b
			if j > prev {
				totalBox += int64(b) * int64(j-prev)
				prev = j
			}
			if prev == n {
				break
			}
		}
		waste := totalBox - totalPackage
		if best == INF || waste < best {
			best = waste
		}
	}
	if best == INF {
		return -1
	}
	return int(best % mod75)
}

func main() {
	fmt.Println(minWastedSpace([]int{2, 3, 5}, [][]int{{4, 8}, {2, 8}})) // 6
	fmt.Println(minWastedSpace([]int{2, 3, 5}, [][]int{{1, 4}, {2, 3}})) // -1
	fmt.Println(minWastedSpace([]int{3, 5, 8, 10, 11, 12}, [][]int{{12}, {11, 9}, {10, 5, 14}})) // 9
}
