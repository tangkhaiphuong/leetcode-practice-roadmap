# Linked List

5 problems on rotation, sorted insertion, design implementations, and DLL-of-buckets.

## Core patterns

| Pattern | When to use | Problems |
|---|---|---|
| **Sentinel / dummy head** | Avoid special-case for empty list and head deletion | 088 |
| **Two pointers (slow/fast)** | Detect cycles, find middle, k-th from end | 086 |
| **Circular trick** | Rotation: connect tail to head, walk to new tail | 086 |
| **Doubly-linked list of buckets** | O(1) increment/decrement on counts | 060 |
| **Walk with "did we wrap?" check** | Insert into circular sorted list | 087 |
| **Merge-by-priority across user lists** | Twitter feed | 085 |

## Common interview pitfalls

- **Modifying `head` in-place** — pass `*ListNode` and remember Go is pass-by-value for the pointer; reassigning `head` inside a function won't affect the caller. Return the new head.
- **Sentinel saves bugs** — `dummy := &ListNode{}; dummy.Next = head` lets you uniformly handle "delete the first node" without branching.
- **Length first, then walk** — for k-rotation, compute length once, then `k %= n` to handle k > n.
- **Circular list infinite loops** — when walking, always have a termination guarantee (e.g., "stop when we return to head" or after `n` steps).
- **Disconnect after rotation** — forgetting to set `newTail.Next = nil` leaves a cycle.
- **All O`one Data Structure** — when a key's count changes, you may need to *create* a new bucket OR move to an *adjacent* bucket; handle both. Empty buckets must be unlinked.

## Suggested practice order

1. **Day 1**: 088 (Design Linked List) — sentinel-headed implementation discipline
2. **Day 2**: 086 (Rotate List) — two-pointer + circular trick
3. **Day 3**: 087 (Insert Circular) — three-case branch logic
4. **Day 4**: 060 (All O`one) — bucketing for O(1) ops
5. **Day 5**: 085 (Design Twitter) — heap-merge across user lists

## Problems in this folder

| # | Problem | Difficulty | Key idea |
|---|---|---|---|
| 060 | All O`one Data Structure | Hard | DLL of count-buckets |
| 085 | Design Twitter | Medium | Per-user list + max-heap merge |
| 086 | Rotate List | Medium | Make circular, find new tail |
| 087 | Insert into a Sorted Circular Linked List | Medium | Walk with three insertion cases |
| 088 | Design Linked List | Medium | Sentinel-headed singly LL |
