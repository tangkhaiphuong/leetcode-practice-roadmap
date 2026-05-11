# Tree & Graph

6 problems including BFS, DFS, topological sort, Union-Find, and binary lifting.

> The README also classifies Word Ladder (058) and Word Ladder II (059) as graph problems; they live in `03-hash-table/` per the README's grouping. Practice them as graph BFS.

## Core patterns

| Pattern | When to use | Examples |
|---|---|---|
| **Border DFS / multi-source BFS** | "Reachable from boundary" grid problems | 039 |
| **Range-bound DFS** | Validate BST (open intervals on subtrees) | 080 |
| **Topological sort (BFS / Kahn's)** | Order from precedence constraints | 079 |
| **Topological sort (DFS w/ colors)** | Detect cycles + topo simultaneously | (alt for 079) |
| **Union-Find with cases** | Cycle detection, redundant edges in directed | 081 |
| **Binary lifting** | k-th ancestor, LCA in O(log n) per query | 082 |
| **BFS on state graph** | Reachability through legal operations | 046 |

## Common interview pitfalls

- **BST bounds, not parent comparison** — `node.Left.Val < node.Val` is INSUFFICIENT. Need all left descendants < node < all right descendants. Pass `(lo, hi)` ranges.
- **Topo sort + cycle detection** — if final ordering has fewer nodes than the graph, there's a cycle. Count UNIQUE letters/nodes (not edges).
- **Alien Dictionary edge case** — `["abc","ab"]` is invalid (longer is prefix of shorter); handle before looking for the differing char.
- **Redundant Connection II three cases** — (a) two parents only, (b) cycle only, (c) both. The "skip the candidate edge then check cycle" trick handles all three.
- **Binary lifting jump table** — `up[v][j] = up[up[v][j-1]][j-1]`; if any intermediate is -1, the jump is also -1.
- **Recursive DFS stack overflow** — for grids 10^6+ cells, switch to BFS or iterative DFS.
- **Water and Jug** — recognize that BFS state graph reduces to Bezout's identity (`z % gcd(x, y) == 0`).

## Suggested practice order

1. **Day 1**: 080 (Validate BST) — recursion fluency
2. **Day 2**: 079 (Alien Dictionary) — topo sort fundamentals
3. **Day 3**: 039 (Surrounded Regions) — border DFS
4. **Day 4**: 046 (Water & Jug) — state-graph BFS OR math
5. **Day 5**: 081 (Redundant Connection II) — Union-Find with twist
6. **Day 6**: 082 (Kth Ancestor) — binary lifting

## Problems in this folder

| # | Problem | Difficulty | Key idea |
|---|---|---|---|
| 039 | Surrounded Regions | Medium | Border DFS marker pass |
| 046 | Water and Jug Problem | Medium | Bezout's identity (or state BFS) |
| 079 | Alien Dictionary | Hard | Topological sort via Kahn's |
| 080 | Validate Binary Search Tree | Medium | Range-bound DFS |
| 081 | Redundant Connection II | Hard | Three-case Union-Find |
| 082 | Kth Ancestor of a Tree Node | Hard | Binary lifting jump table |
