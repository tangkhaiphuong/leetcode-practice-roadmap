//go:build ignore

// Problem 028: Making File Names Unique (LeetCode #1487)
// Difficulty: Medium
// Categories: String, Array, Hash Table
//
// Description:
//   Given a list of names processed in order, assign each request a unique
//   name. If name already used, append the smallest k>=1 such that
//   "<name>(k)" hasn't been used. Return the assigned names.
//
// Examples:
//   ["pes","fifa","gta","pes(2019)"] -> same (none collide)
//   ["gta","gta(1)","gta","avalon"]  -> ["gta","gta(1)","gta(2)","avalon"]
//   ["onepiece_01(1)","onepiece_02","onepiece_03","onepiece_04","onepiece_01"]
//     -> ["onepiece_01(1)","onepiece_02","onepiece_03","onepiece_04","onepiece_01"]
//   ["wano","wano","wano","wano"]    -> ["wano","wano(1)","wano(2)","wano(3)"]
//
// Approach: Map of next-suffix-to-try
//   used: name -> next k to try.
//   For each name n: if !used[n], take n; set used[n]=1.
//   Else start k = used[n]; while used contains "n(k)" increment k. Take
//   "n(k)"; set used[n]=k+1; set used[that name]=1.
//
// Complexity:
//   Time:  amortized O(n) total (each k++ corresponds to a previously used slot)
//   Space: O(n)

package main

import (
	"fmt"
	"strconv"
)

func getFolderNames(names []string) []string {
	used := map[string]int{}
	out := make([]string, len(names))
	for i, n := range names {
		if _, ok := used[n]; !ok {
			out[i] = n
			used[n] = 1
			continue
		}
		k := used[n]
		for {
			cand := n + "(" + strconv.Itoa(k) + ")"
			if _, ok := used[cand]; !ok {
				out[i] = cand
				used[cand] = 1
				used[n] = k + 1
				break
			}
			k++
		}
	}
	return out
}

func main() {
	fmt.Println(getFolderNames([]string{"pes", "fifa", "gta", "pes(2019)"}))
	fmt.Println(getFolderNames([]string{"gta", "gta(1)", "gta", "avalon"}))
	fmt.Println(getFolderNames([]string{"wano", "wano", "wano", "wano"}))
}
