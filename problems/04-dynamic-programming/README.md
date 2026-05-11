# Dynamic Programming

22 problems on state machines, range DP, matching DP, subset DP, and matrix DP — DP is the single most over-tested topic at FAANG.

## Core patterns

| Pattern | When to use | Problems |
|---|---|---|
| **1D linear DP** | "Ways to ..." / "longest ending here" | 011, 012, 018, 035 |
| **2D matching DP** | Two strings, regex / wildcard | 008, 011, 014, 017 |
| **DP on subarrays / intervals** | Range with two ends | 017, 067 |
| **State machines** (buy/sell, k-trans) | Stock-style discrete states | 066 |
| **DP on grids** | Path constraints, min cost / max product | 050, 053 |
| **Subset-sum DP** | Pick subset to hit target | 030, 057 |
| **DP dimension flip** | When the obvious dp blows up; flip what's the answer vs. parameter | 068 |
| **Greedy on bit patterns** | Halve / +1 / -1 transformations | 069 |
| **Digit DP / place-value math** | Count digit occurrences in [0, n] | 071 |
| **Greedy by sum mod k** | Largest divisible | 070 |
| **Cycle-detection DP** | Repetition counting | 031 |

## Common interview pitfalls

- **Initialize -inf carefully** — `math.MinInt64` minus something can overflow. Use `-1 << 30` or branch.
- **Rolling 1D vs 2D** — for `dp[i][j]` that only reads `dp[i-1][...]`, you can collapse to 1D, but watch iteration direction.
- **State explosion** — when `k >= n/2` in Stock IV, short-circuit to unlimited-transaction case.
- **Egg drop trick** — the natural `dp[k][n]` is slow; the reformulation `dp[m][k] = max floors with m moves` makes it efficient.
- **Wildcard `*` semantics** — in #014 (Wildcard), `*` matches any sequence; in #008 (Regex), `*` means "zero or more of preceding". Don't confuse the two recurrences.
- **Decode Ways traps** — `"06"` is 0 decodings; `"10"` is 1; `"100"` is 0. Always check the two-digit number is in [10, 26], not just [1, 26].
- **Matrix DP base cases** — initialize the first row/column carefully when transition reads both `dp[i-1][j]` and `dp[i][j-1]`.

## Suggested practice order

1. **Linear DP** (Day 1): 012, 018, 034
2. **Matching DP** (Day 2): 008, 011, 014
3. **Subarray DP** (Day 3): 035, 017
4. **State machine** (Day 4): 066, 015
5. **Matrix DP** (Day 5): 050, 053, 057
6. **Advanced re-derivation** (Day 6): 067, 068
7. **Math-flavored DP** (Day 7): 071, 070, 069
8. **Niche** (Day 8): 001, 010, 030, 031

## Problems in this folder

| # | Problem | Difficulty | Key idea |
|---|---|---|---|
| 001 | Longest Palindromic Substring | Medium | Expand-around-center (or O(n²) DP) |
| 008 | Regular Expression Matching | Hard | 2D DP, `*` consumes 0 or 1+ |
| 010 | Longest Valid Parentheses | Hard | Stack of indices OR O(1) two-counter pass |
| 011 | Interleaving String | Medium | 2D DP, rolling 1D row |
| 012 | Decode Ways | Medium | Linear DP, two rolling vars |
| 014 | Wildcard Matching | Hard | 2D DP, `*` = any seq |
| 015 | Valid Parenthesis String | Medium | Lo/hi range tracking |
| 017 | Palindrome Partitioning II | Hard | Palindrome table + cuts DP |
| 018 | Decode Ways II | Hard | DP w/ case enum on `*` |
| 030 | Split Array With Same Average | Hard | Meet in the middle (subset sums) |
| 031 | Count The Repetitions | Hard | Cycle detection on s2-pointer |
| 034 | Jump Game II | Medium | BFS-style greedy range |
| 035 | Maximum Product Subarray | Medium | Track min AND max |
| 050 | Dungeon Game | Hard | Bottom-up DP from destination |
| 053 | Maximum Non Negative Product in a Matrix | Medium | Track min and max products |
| 057 | Minimize Difference Target | Hard | Subset-sum DP across rows |
| 066 | Best Time to Buy and Sell Stock IV | Hard | buy[k]/sell[k] state machine |
| 067 | Maximum Score from Multiplication | Hard | DP on (op, leftTaken) |
| 068 | Super Egg Drop | Hard | dp[m][k] = max floors with m moves |
| 069 | Integer Replacement | Medium | Greedy on trailing bits |
| 070 | Largest Multiple of Three | Hard | Greedy drop by sum mod 3 |
| 071 | Number of Digit One | Hard | Place-value digit DP |
