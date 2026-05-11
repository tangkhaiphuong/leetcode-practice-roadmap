# Hash Table & Hashing

13 problems on fast lookup, group-by-key, prefix-sum residues, BFS over implicit graphs, and slope hashing.

## Core patterns

| Pattern | When to use | Problems |
|---|---|---|
| **Last-seen / first-seen index map** | Sliding window starts, prefix conditions | 004, 064 |
| **Two-sum-style lookup** | Pair completion (sum, diff, ratio) | 063, 065 |
| **Group by key + pairwise** | Cross-group constraints | 062 |
| **Bucket modulo** | Numeric proximity / nearby duplicate | 061 |
| **BFS on implicit word graph** | One-step transformations | 058, 059 |
| **Doubly-linked-list buckets** | O(1) max/min on counts | 060 |
| **Window of words (multiset)** | Substring concatenation match | 023 |
| **Remainder map (long division)** | Detect repeating decimals | 025 |
| **Counter map + suffix retry** | Name disambiguation | 028 |
| **Slope hashing** (gcd-reduced pair as key) | Collinear-point counting | 040 |
| **Mirror set** | Geometric symmetry | 065 |

## Common interview pitfalls

- **Map zero value vs absent key** — `m[k]` returns the zero value if absent. Use `v, ok := m[k]` when you need to distinguish.
- **Iteration order** — Go map iteration is RANDOMIZED. Never rely on order; collect keys to a slice and sort if needed.
- **Composite keys** — for tuples use `[2]int` arrays (comparable), not slices (`[]int` is not). Or build a `string` key.
- **Mod-key collisions in prefix sum** — when subtracting then mod, normalize: `((a - b) % p + p) % p` for non-negative residue.
- **BFS pattern bucketing** — for word-ladder style, build adjacency via wildcard buckets ONCE, not per node, or you'll TLE.
- **Word Ladder II termination** — keep BFS going to FINISH the level when end is found, so all parents in the shortest layer are recorded.
- **Slope hashing precision** — never use float slopes. Reduce `dy/dx` by GCD and normalize sign.

## Suggested practice order

1. **Two-sum-style** (Day 1): 063, 064, 065
2. **Sliding hash + bucket** (Day 2): 004, 061
3. **Group-by-key & remainder maps** (Day 3): 025, 028, 062
4. **BFS on implicit graphs** (Day 4): 058, 059
5. **Design with O(1) ops** (Day 5): 060
6. **Hash for geometry / multiset window** (Day 6): 023, 040

## Problems in this folder

| # | Problem | Difficulty | Key idea |
|---|---|---|---|
| 004 | Longest Substring Without Repeating Characters | Medium | Sliding window + last-seen map |
| 023 | Substring with Concatenation of All Words | Hard | Sliding window of words (multiset) |
| 025 | Fraction to Recurring Decimal | Medium | Long division + remainder map |
| 028 | Making File Names Unique | Medium | Counter map + suffix retry |
| 040 | Max Points on a Line | Hard | GCD-reduced slope hashing |
| 058 | Word Ladder | Hard | BFS w/ wildcard pattern buckets |
| 059 | Word Ladder II | Hard | BFS parent graph + DFS reconstruct |
| 060 | All O`one Data Structure | Hard | DLL of count-buckets |
| 061 | Contains Duplicate III | Hard | Bucketing by value range |
| 062 | Invalid Transactions | Medium | Group-by-name pairwise |
| 063 | Count Good Meals | Medium | Two-sum for each power-of-2 |
| 064 | Make Sum Divisible by P | Medium | Prefix sum mod p + first-index map |
| 065 | Line Reflection | Medium | Mirror set across (minX+maxX)/2 |
