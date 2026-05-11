# Greedy Algorithms

10 problems requiring a greedy proof. The discipline: always justify why the local choice is globally optimal (exchange argument).

## Core patterns

| Pattern | When to use | Problems |
|---|---|---|
| **Custom comparator on strings** | Largest / smallest concatenation | 009, 078 |
| **Monotonic stack greedy** | Smallest result via "remove if blocked" | 016 |
| **Heap-based greedy** | "Pick the best available now" | 038 |
| **Range tracking (lo/hi)** | When backtracking would explode | 015 |
| **Strata / bulk levels** | Diminishing values, arithmetic series | 083 |
| **Sort + sliding window** | "Most you can fix with k changes" | 054 |
| **BFS layer = greedy range** | Min jumps to reach end | (cross: 034, see DP folder) |
| **Track min and max running** | When sign flips help (negative * negative) | 035 |
| **Multi-pass length-regime branching** | Edit cost with bounds | 024 |
| **Meet-in-the-middle** | Subset-sum with `n ≤ 30` | 030 |

## Common interview pitfalls

- **No proof = no points** — interviewers will press: "why does this work?" Have a one-line exchange argument ready.
- **Tie-breaking** — when greedy chooses by `compare(a,b)`, define behavior on equal keys.
- **Largest Number trap** — comparator must use `a+b > b+a` (concatenation lex), not numeric compare. And strip leading zeros at the end ("0000" → "0").
- **Modular sums on bulk operations** — for 083, compute total in regular int64 BEFORE mod (division by 2 needs modular inverse otherwise).
- **Heap by max vs min** — Go's `container/heap` is min-heap. For max, negate values or implement custom `Less`.
- **Valid Parenthesis String** — range tracking is cleaner than enumerating all `*` substitutions; track `[lo, hi]` and clamp `lo` at 0.
- **Create Maximum Number** — `O(k * (m+n)²)` is the expected complexity; don't try to make it cleaner.

## Suggested practice order

1. **Day 1**: 009 (Largest Number) — clean comparator-based greedy
2. **Day 2**: 016 — monotonic stack greedy
3. **Day 3**: 015 — range-tracking greedy
4. **Day 4**: 035 — track min and max
5. **Day 5**: 038, 054 — heap-based and window-based
6. **Day 6**: 024, 030 — multi-regime / meet-in-the-middle
7. **Day 7**: 078 — combine monotonic stack + greedy merge
8. **Day 8**: 083 — strata-based with arithmetic series

## Problems in this folder

| # | Problem | Difficulty | Key idea |
|---|---|---|---|
| 009 | Largest Number | Medium | Comparator on `a+b` vs `b+a` |
| 015 | Valid Parenthesis String | Medium | Track lo/hi range of open count |
| 016 | Remove K Digits | Medium | Monotonic increasing stack |
| 024 | Strong Password Checker | Hard | Length-regime case analysis |
| 030 | Split Array With Same Average | Hard | Meet in the middle |
| 035 | Maximum Product Subarray | Medium | Track min AND max |
| 038 | Maximum Number of Events | Medium | Day-sweep + min-heap by end |
| 054 | Frequency of the Most Frequent Element | Medium | Sort + sliding window |
| 078 | Create Maximum Number | Hard | Per-split max-subseq + best merge |
| 083 | Sell Diminishing-Valued Colored Balls | Medium | Sort + bulk levels |
