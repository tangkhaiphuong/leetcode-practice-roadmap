# Stack & Monotonic Stack

7 problems on plain stacks and monotonic stacks — a high-frequency FAANG pattern especially in array/string problems.

## Core patterns

| Pattern | When to use | Problems |
|---|---|---|
| **Plain stack** | Balanced parens, expression eval | 010, 015 |
| **Stack of indices** | Need distance to previous/next event | 010, 037 |
| **Monotonic increasing stack** | "Smallest result", "min so far" | 016 |
| **Monotonic decreasing stack (right-to-left)** | Next greater / 132 pattern | 037 |
| **Monotonic decreasing stack + greedy merge** | Lexicographically largest concatenation | 078 |
| **Two-pass stack from each end** | Out-of-order bounds | 042 |
| **Range simulation with stack** | Decoded string virtualization | 029 |
| **Range-bound counters as a stack** | Valid parenthesis range tracking | 015 |

## Common interview pitfalls

- **When to use `[]int` vs a custom struct** — for index-only stacks, plain `[]int` is fastest. For value-with-meta, a struct slice keeps things readable.
- **Sentinel pushes** — pushing `-1` (or `n`) at the start lets you uniformly compute window lengths without special-casing the empty stack.
- **Pop-while loops** — check the loop guard each iteration; don't capture `len(stk)` once.
- **Direction matters** — for "next greater", you can scan right-to-left with a decreasing stack OR left-to-right popping when current breaks monotonicity. Both work; pick one and stick with it.
- **`132 Pattern` subtle invariant** — `k` (the popped, candidate "2") must be tracked across the whole right-to-left walk, not reset inside the loop.
- **Create Maximum Number** — the trick is per-split brute force: try every `i` for length-i from `nums1` and length-(k-i) from `nums2`, then greedy-merge.

## Suggested practice order

1. **Day 1**: 010 (Longest Valid Parens) — sentinel + stack of indices
2. **Day 2**: 016 (Remove K Digits) — monotonic increasing stack
3. **Day 3**: 037 (132 Pattern) — right-to-left monotonic decreasing
4. **Day 4**: 042 — two-pass stack from each end
5. **Day 5**: 015 — range tracking (alternative to stack)
6. **Day 6**: 029 — decoded string virtualization
7. **Day 7**: 078 (Create Max Number) — combines monotonic stack with greedy merge

## Problems in this folder

| # | Problem | Difficulty | Key idea |
|---|---|---|---|
| 010 | Longest Valid Parentheses | Hard | Stack of indices with sentinel |
| 015 | Valid Parenthesis String | Medium | Lo/hi range tracking |
| 016 | Remove K Digits | Medium | Monotonic increasing stack |
| 029 | Decoded String at Index | Medium | Reverse-walk virtualization |
| 037 | 132 Pattern | Medium | Right-to-left monotonic stack |
| 042 | Shortest Unsorted Continuous Subarray | Medium | Two-pass min/max bounds |
| 078 | Create Maximum Number | Hard | Per-split max-subseq + best merge |
