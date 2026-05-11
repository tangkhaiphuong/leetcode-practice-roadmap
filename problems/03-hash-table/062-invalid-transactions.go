//go:build ignore

// Problem 062: Invalid Transactions (LeetCode #1169)
// Difficulty: Medium
// Categories: Hash Table, Sorting
//
// Description:
//   Each transaction is "name,time,amount,city". A transaction is invalid if:
//     - amount > 1000, OR
//     - there exists another tx with same name in a different city within
//       60 minutes (|t1 - t2| <= 60).
//   Return all invalid transactions.
//
// Approach: Group by Name, Then Pairwise Check
//   Parse each tx. Group indices by name. For each group, compare every pair
//   for the cross-city + within-60 condition; mark both invalid. Also flag
//   any tx with amount > 1000.
//
// Complexity:
//   Time:  O(N + sum(group_size^2)). Acceptable for the constraints.
//   Space: O(N)

package main

import (
	"fmt"
	"strconv"
	"strings"
)

type tx struct {
	name, city string
	time, amt  int
}

func invalidTransactions(transactions []string) []string {
	parsed := make([]tx, len(transactions))
	for i, t := range transactions {
		parts := strings.Split(t, ",")
		time, _ := strconv.Atoi(parts[1])
		amt, _ := strconv.Atoi(parts[2])
		parsed[i] = tx{parts[0], parts[3], time, amt}
	}
	bad := make([]bool, len(transactions))
	for i := range parsed {
		if parsed[i].amt > 1000 {
			bad[i] = true
		}
		for j := i + 1; j < len(parsed); j++ {
			if parsed[i].name != parsed[j].name {
				continue
			}
			if parsed[i].city != parsed[j].city &&
				abs(parsed[i].time-parsed[j].time) <= 60 {
				bad[i] = true
				bad[j] = true
			}
		}
	}
	var res []string
	for i, b := range bad {
		if b {
			res = append(res, transactions[i])
		}
	}
	return res
}

func abs(x int) int { if x < 0 { return -x }; return x }

func main() {
	fmt.Println(invalidTransactions([]string{"alice,20,800,mtv", "alice,50,100,beijing"}))
	fmt.Println(invalidTransactions([]string{"alice,20,800,mtv", "alice,50,1200,mtv"}))
	fmt.Println(invalidTransactions([]string{"alice,20,800,mtv", "bob,50,1200,mtv"}))
}
