# 3161. Block Placement Queries - Segment Tree + Ordered Set Solution

## Table of Contents

* [Problem Summary](#problem-summary)
* [Constraints](#constraints)
* [Intuition](#intuition)
* [Approach](#approach)
* [Data Structures Used](#data-structures-used)
* [Operations & Behavior Summary](#operations--behavior-summary)
* [Complexity](#complexity)
* [Multi-language Solutions](#multi-language-solutions)

  * [C++](#c)
  * [Java](#java)
  * [JavaScript](#javascript)
  * [Python3](#python3)
  * [Go](#go)
* [Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)](#step-by-step-detailed-explanation-c-java-javascript-python3-go)
* [Examples](#examples)
* [How to Use / Run Locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

## Problem Summary

LeetCode 3161 - Block Placement Queries is a hard data structure problem that combines interval management, ordered sets, and range maximum queries.

We are given an infinite number line starting from position `0`.

There are two types of queries:

1. Add an obstacle at position `x`.
2. Check whether a block of size `sz` can be placed somewhere inside the range `[0, x]` without intersecting any obstacle.

The block is allowed to touch an obstacle, but it cannot overlap with one.

For every query of type `2`, we need to return:

* `true` if a valid placement exists
* `false` otherwise

The main challenge is answering many placement queries efficiently while obstacles are continuously being added.

This problem is commonly solved using an Ordered Set and Segment Tree because they allow fast updates and fast maximum-gap queries.

## Constraints

| Constraint                                        | Value             |
| ------------------------------------------------- | ----------------- |
| `1 <= queries.length <= 150000`                   | Number of queries |
| `2 <= queries[i].length <= 3`                     | Query size        |
| `queries[i][0]` is either `1` or `2`              | Query type        |
| `1 <= x, sz <= min(5 * 10^4, 3 * queries.length)` | Coordinate limit  |
| No duplicate obstacle insertion                   | Guaranteed        |
| At least one type-2 query exists                  | Guaranteed        |

## Intuition

My first observation was that I never actually need to place the block.

I only need to know whether there exists an empty segment large enough to fit the block.

Instead of thinking about blocks, I started thinking about gaps between obstacles.

If I know the largest available gap inside a range, then answering a placement query becomes very easy.

The next thing I noticed was that inserting a new obstacle only affects nearby gaps. It does not change the entire number line.

That means I can maintain gap information dynamically instead of recomputing everything after every query.

Once I reached that observation, the solution naturally became an Ordered Set plus a Segment Tree.

## Approach

1. Store all obstacle positions inside an ordered structure.
2. Add a sentinel obstacle at position `0`.
3. For every obstacle position, maintain the gap between it and the previous obstacle.
4. Build a Segment Tree over those gap values.
5. When inserting a new obstacle:

   * Find its previous obstacle.
   * Find its next obstacle.
   * Update only the affected gaps.
6. For a type-2 query:

   * Find the obstacle immediately before `x`.
   * Check the free space between that obstacle and `x`.
   * Query the maximum gap inside the prefix `[0, x]`.
7. If either value is at least `sz`, return `true`.
8. Otherwise return `false`.

This allows every operation to be processed in logarithmic time.

## Data Structures Used

### Ordered Set

Used to quickly find:

* Previous obstacle
* Next obstacle

Operations take `O(log n)` time.

### Segment Tree

Used to maintain:

* Maximum gap length
* Fast range maximum queries
* Fast point updates

Operations take `O(log n)` time.

### Gap Array Concept

For every obstacle position, I store the distance from the previous obstacle.

This transforms the geometric problem into a range maximum query problem.

## Operations & Behavior Summary

### Type 1 Query

Add obstacle at position `x`.

Steps:

1. Find previous obstacle.
2. Find next obstacle.
3. Create a new gap ending at `x`.
4. Update the next obstacle's gap.
5. Update the Segment Tree.
6. Insert `x` into the Ordered Set.

### Type 2 Query

Check whether a block of size `sz` fits inside `[0, x]`.

Steps:

1. Find the obstacle immediately before `x`.
2. Compute the remaining free space up to `x`.
3. Query the Segment Tree for the largest gap in the prefix.
4. If either gap is large enough:

   * return `true`
5. Otherwise:

   * return `false`

## Complexity

| Operation               | Complexity |
| ----------------------- | ---------- |
| Insert Obstacle         | `O(log n)` |
| Query Maximum Gap       | `O(log n)` |
| Find Neighbor Obstacles | `O(log n)` |

Overall:

| Metric           | Complexity   |
| ---------------- | ------------ |
| Time Complexity  | `O(n log n)` |
| Space Complexity | `O(n)`       |

Where:

* `n` = total number of queries

The logarithmic factor comes from Ordered Set and Segment Tree operations.

## Multi-language Solutions

### C++

```cpp
class Fenwick {
    int n;
    vector<int> bit;

public:
    Fenwick(int n = 0) : n(n), bit(n + 1, 0) {}

    // Add "val" at index "idx" because I need point updates on occupied positions.
    void add(int idx, int val) {
        for (; idx <= n; idx += idx & -idx) {
            bit[idx] += val;
        }
    }

    // Return prefix sum [1..idx] because I use it to count how many obstacles are on the left.
    int sum(int idx) const {
        int res = 0;
        for (; idx > 0; idx -= idx & -idx) {
            res += bit[idx];
        }
        return res;
    }

    // Find the smallest index whose prefix sum is at least k.
    // I use this to jump to the k-th occupied position.
    int kth(int k) const {
        int idx = 0;
        int step = 1;
        while ((step << 1) <= n) step <<= 1;

        for (int d = step; d > 0; d >>= 1) {
            int next = idx + d;
            if (next <= n && bit[next] < k) {
                idx = next;
                k -= bit[next];
            }
        }
        return idx + 1;
    }
};

class SegTree {
    int n;
    vector<int> tree;

    // Update one position because only one obstacle gap changes at a time.
    void update(int node, int l, int r, int pos, int val) {
        if (l == r) {
            tree[node] = val;
            return;
        }
        int mid = (l + r) >> 1;
        if (pos <= mid) update(node << 1, l, mid, pos, val);
        else update(node << 1 | 1, mid + 1, r, pos, val);
        tree[node] = max(tree[node << 1], tree[node << 1 | 1]);
    }

    // Query the maximum on a prefix because I only care about gaps ending at or before x.
    int query(int node, int l, int r, int ql, int qr) const {
        if (ql > r || qr < l) return 0;
        if (ql <= l && r <= qr) return tree[node];
        int mid = (l + r) >> 1;
        return max(query(node << 1, l, mid, ql, qr),
                   query(node << 1 | 1, mid + 1, r, ql, qr));
    }

public:
    SegTree(int n = 0) : n(n), tree(4 * max(1, n), 0) {}

    void setVal(int pos, int val) {
        if (n == 0) return;
        update(1, 0, n - 1, pos, val);
    }

    int getMax(int l, int r) const {
        if (n == 0 || l > r) return 0;
        return query(1, 0, n - 1, l, r);
    }
};

class Solution {
public:
    vector<bool> getResults(vector<vector<int>>& queries) {
        int mx = 0;
        for (auto &q : queries) {
            mx = max(mx, q[1]);
        }

        // I shift positions by +1 in Fenwick so that position 0 can still be stored safely.
        int fenwickSize = mx + 2;
        Fenwick fw(fenwickSize);

        // I store one gap value for every coordinate from 0 to mx.
        SegTree st(mx + 1);

        // Sentinel obstacle at 0 makes predecessor logic simple.
        fw.add(1, 1);

        vector<bool> ans;
        ans.reserve(queries.size());

        for (auto &q : queries) {
            int type = q[0];
            int x = q[1];

            if (type == 1) {
                // Count occupied positions strictly smaller than x.
                int leftCount = fw.sum(x);
                int leftPos = fw.kth(leftCount) - 1;

                // Count occupied positions up to x, then jump to the next one if it exists.
                int occupiedUpToX = fw.sum(x + 1);
                int totalOccupied = fw.sum(fenwickSize);
                int rightPos = -1;
                if (occupiedUpToX < totalOccupied) {
                    rightPos = fw.kth(occupiedUpToX + 1) - 1;
                }

                // The new obstacle creates the gap ending at x.
                st.setVal(x, x - leftPos);

                // The next obstacle's gap shrinks because x is now between them.
                if (rightPos != -1) st.setVal(rightPos, rightPos - x);

                // Mark x as occupied.
                fw.add(x + 1, 1);
            } else {
                int sz = q[2];

                // The last obstacle strictly before x gives the tail gap ending at x.
                int leftCount = fw.sum(x);
                int leftPos = fw.kth(leftCount) - 1;

                // Any gap ending at a position <= x is fully inside [0, x].
                int bestPrefix = st.getMax(0, x);

                // Either the tail gap is enough, or some earlier gap is enough.
                ans.push_back((x - leftPos >= sz) || (bestPrefix >= sz));
            }
        }

        return ans;
    }
};
```

### Java

```java
import java.util.*;

class Fenwick {
    int n;
    int[] bit;

    Fenwick(int n) {
        this.n = n;
        this.bit = new int[n + 1];
    }

    // Point add because each obstacle insertion changes one count.
    void add(int idx, int val) {
        for (; idx <= n; idx += idx & -idx) {
            bit[idx] += val;
        }
    }

    // Prefix sum because I need counts of occupied positions on the left.
    int sum(int idx) {
        int res = 0;
        for (; idx > 0; idx -= idx & -idx) {
            res += bit[idx];
        }
        return res;
    }

    // Find the smallest index whose prefix sum is at least k.
    int kth(int k) {
        int idx = 0;
        int step = 1;
        while ((step << 1) <= n) step <<= 1;

        for (int d = step; d > 0; d >>= 1) {
            int next = idx + d;
            if (next <= n && bit[next] < k) {
                idx = next;
                k -= bit[next];
            }
        }
        return idx + 1;
    }
}

class SegTree {
    int n;
    int[] tree;

    SegTree(int n) {
        this.n = n;
        this.tree = new int[Math.max(4, 4 * n)];
    }

    // Update one coordinate because only one gap value changes at a time.
    void update(int node, int l, int r, int pos, int val) {
        if (l == r) {
            tree[node] = val;
            return;
        }
        int mid = (l + r) >>> 1;
        if (pos <= mid) update(node << 1, l, mid, pos, val);
        else update(node << 1 | 1, mid + 1, r, pos, val);
        tree[node] = Math.max(tree[node << 1], tree[node << 1 | 1]);
    }

    // Query a prefix maximum because all valid stored gaps end inside [0, x].
    int query(int node, int l, int r, int ql, int qr) {
        if (ql > r || qr < l) return 0;
        if (ql <= l && r <= qr) return tree[node];
        int mid = (l + r) >>> 1;
        return Math.max(query(node << 1, l, mid, ql, qr),
                        query(node << 1 | 1, mid + 1, r, ql, qr));
    }
}

class Solution {
    public List<Boolean> getResults(int[][] queries) {
        int mx = 0;
        for (int[] q : queries) {
            mx = Math.max(mx, q[1]);
        }

        int fenwickSize = mx + 2;
        Fenwick fw = new Fenwick(fenwickSize);
        SegTree st = new SegTree(mx + 1);

        // Sentinel obstacle at 0 makes predecessor search easier.
        fw.add(1, 1);

        List<Boolean> ans = new ArrayList<>();

        for (int[] q : queries) {
            int type = q[0];
            int x = q[1];

            if (type == 1) {
                // Count obstacles strictly smaller than x.
                int leftCount = fw.sum(x);
                int leftPos = fw.kth(leftCount) - 1;

                // Find the next obstacle after x if it exists.
                int occupiedUpToX = fw.sum(x + 1);
                int totalOccupied = fw.sum(fenwickSize);
                int rightPos = -1;
                if (occupiedUpToX < totalOccupied) {
                    rightPos = fw.kth(occupiedUpToX + 1) - 1;
                }

                // New obstacle x becomes the end of a new gap.
                st.update(1, 0, mx, x, x - leftPos);

                // The following obstacle now sees a shorter gap.
                if (rightPos != -1) {
                    st.update(1, 0, mx, rightPos, rightPos - x);
                }

                // Mark x as occupied.
                fw.add(x + 1, 1);
            } else {
                int sz = q[2];

                // The last obstacle strictly before x gives the free tail up to x.
                int leftCount = fw.sum(x);
                int leftPos = fw.kth(leftCount) - 1;

                // Prefix max over gaps that end at or before x.
                int bestPrefix = st.query(1, 0, mx, 0, x);

                ans.add((x - leftPos >= sz) || (bestPrefix >= sz));
            }
        }

        return ans;
    }
}
```

### JavaScript

```javascript
/**
 * Fenwick tree because I need fast predecessor and successor lookup.
 */
class Fenwick {
    constructor(n) {
        this.n = n;
        this.bit = new Array(n + 1).fill(0);
    }

    // Point add because each obstacle insertion changes one count.
    add(idx, val) {
        for (; idx <= this.n; idx += idx & -idx) {
            this.bit[idx] += val;
        }
    }

    // Prefix sum because I need how many obstacles are on the left side.
    sum(idx) {
        let res = 0;
        for (; idx > 0; idx -= idx & -idx) {
            res += this.bit[idx];
        }
        return res;
    }

    // Find the smallest index with prefix sum at least k.
    kth(k) {
        let idx = 0;
        let step = 1;
        while ((step << 1) <= this.n) step <<= 1;

        for (let d = step; d > 0; d >>= 1) {
            const next = idx + d;
            if (next <= this.n && this.bit[next] < k) {
                idx = next;
                k -= this.bit[next];
            }
        }
        return idx + 1;
    }
}

/**
 * Segment tree because I need a fast prefix maximum over gap lengths.
 */
class SegTree {
    constructor(n) {
        this.n = n;
        this.tree = new Array(Math.max(4, 4 * n)).fill(0);
    }

    // Update one coordinate because only one stored gap changes at a time.
    update(node, l, r, pos, val) {
        if (l === r) {
            this.tree[node] = val;
            return;
        }
        const mid = (l + r) >> 1;
        if (pos <= mid) this.update(node << 1, l, mid, pos, val);
        else this.update(node << 1 | 1, mid + 1, r, pos, val);
        this.tree[node] = Math.max(this.tree[node << 1], this.tree[node << 1 | 1]);
    }

    // Query a prefix because every useful gap ends at or before x.
    query(node, l, r, ql, qr) {
        if (ql > r || qr < l) return 0;
        if (ql <= l && r <= qr) return this.tree[node];
        const mid = (l + r) >> 1;
        return Math.max(
            this.query(node << 1, l, mid, ql, qr),
            this.query(node << 1 | 1, mid + 1, r, ql, qr)
        );
    }
}

/**
 * @param {number[][]} queries
 * @return {boolean[]}
 */
var getResults = function (queries) {
    let mx = 0;
    for (const q of queries) {
        mx = Math.max(mx, q[1]);
    }

    const fw = new Fenwick(mx + 2);
    const st = new SegTree(mx + 1);

    // Sentinel obstacle at 0 keeps predecessor logic simple.
    fw.add(1, 1);

    const ans = [];

    for (const q of queries) {
        const type = q[0];
        const x = q[1];

        if (type === 1) {
            // Count occupied positions strictly smaller than x.
            const leftCount = fw.sum(x);
            const leftPos = fw.kth(leftCount) - 1;

            // Find the next occupied position after x if it exists.
            const occupiedUpToX = fw.sum(x + 1);
            const totalOccupied = fw.sum(mx + 2);
            let rightPos = -1;
            if (occupiedUpToX < totalOccupied) {
                rightPos = fw.kth(occupiedUpToX + 1) - 1;
            }

            // x becomes the end of a new gap.
            st.update(1, 0, mx, x, x - leftPos);

            // The next obstacle now sees a shorter gap.
            if (rightPos !== -1) {
                st.update(1, 0, mx, rightPos, rightPos - x);
            }

            // Mark x as occupied.
            fw.add(x + 1, 1);
        } else {
            const sz = q[2];

            // The last obstacle strictly before x gives the free suffix up to x.
            const leftCount = fw.sum(x);
            const leftPos = fw.kth(leftCount) - 1;

            // Prefix maximum of gaps ending at or before x.
            const bestPrefix = st.query(1, 0, mx, 0, x);

            ans.push((x - leftPos >= sz) || (bestPrefix >= sz));
        }
    }

    return ans;
};
```

### Python3

```python
from typing import List

class Fenwick:
    def __init__(self, n: int) -> None:
        self.n = n
        self.bit = [0] * (n + 1)

    # Point add because I only insert obstacles.
    def add(self, idx: int, val: int) -> None:
        while idx <= self.n:
            self.bit[idx] += val
            idx += idx & -idx

    # Prefix sum because I need the number of occupied positions on the left.
    def sum(self, idx: int) -> int:
        res = 0
        while idx > 0:
            res += self.bit[idx]
            idx -= idx & -idx
        return res

    # Find the smallest index whose prefix sum is at least k.
    def kth(self, k: int) -> int:
        idx = 0
        step = 1
        while (step << 1) <= self.n:
            step <<= 1

        d = step
        while d:
            nxt = idx + d
            if nxt <= self.n and self.bit[nxt] < k:
                idx = nxt
                k -= self.bit[nxt]
            d >>= 1
        return idx + 1


class SegTree:
    def __init__(self, n: int) -> None:
        self.n = n
        self.tree = [0] * max(4, 4 * n)

    # Update one position because only one stored gap changes at a time.
    def update(self, node: int, l: int, r: int, pos: int, val: int) -> None:
        if l == r:
            self.tree[node] = val
            return
        mid = (l + r) // 2
        if pos <= mid:
            self.update(node * 2, l, mid, pos, val)
        else:
            self.update(node * 2 + 1, mid + 1, r, pos, val)
        self.tree[node] = max(self.tree[node * 2], self.tree[node * 2 + 1])

    # Query a prefix maximum because all useful gaps end at or before x.
    def query(self, node: int, l: int, r: int, ql: int, qr: int) -> int:
        if ql > r or qr < l:
            return 0
        if ql <= l and r <= qr:
            return self.tree[node]
        mid = (l + r) // 2
        return max(
            self.query(node * 2, l, mid, ql, qr),
            self.query(node * 2 + 1, mid + 1, r, ql, qr),
        )


class Solution:
    def getResults(self, queries: List[List[int]]) -> List[bool]:
        mx = 0
        for q in queries:
            mx = max(mx, q[1])

        fw = Fenwick(mx + 2)
        st = SegTree(mx + 1)

        # Sentinel obstacle at 0 makes predecessor search easy.
        fw.add(1, 1)

        ans: List[bool] = []

        for q in queries:
            t = q[0]
            x = q[1]

            if t == 1:
                # Count occupied positions strictly smaller than x.
                left_count = fw.sum(x)
                left_pos = fw.kth(left_count) - 1

                # Find the next occupied position after x if it exists.
                occupied_up_to_x = fw.sum(x + 1)
                total_occupied = fw.sum(mx + 2)
                right_pos = -1
                if occupied_up_to_x < total_occupied:
                    right_pos = fw.kth(occupied_up_to_x + 1) - 1

                # x becomes the end of a new gap.
                st.update(1, 0, mx, x, x - left_pos)

                # The next obstacle now sees a shorter gap.
                if right_pos != -1:
                    st.update(1, 0, mx, right_pos, right_pos - x)

                # Mark x as occupied.
                fw.add(x + 1, 1)
            else:
                sz = q[2]

                # The last obstacle strictly before x gives the free suffix up to x.
                left_count = fw.sum(x)
                left_pos = fw.kth(left_count) - 1

                # Prefix maximum of gaps ending at or before x.
                best_prefix = st.query(1, 0, mx, 0, x)

                ans.append((x - left_pos >= sz) or (best_prefix >= sz))

        return ans
```

### Go

```go
type Fenwick struct {
 n   int
 bit []int
}

// I use a Fenwick tree because I need fast obstacle counts on prefixes.
func NewFenwick(n int) *Fenwick {
 return &Fenwick{
  n:   n,
  bit: make([]int, n+1),
 }
}

// Point add because only one coordinate changes per insertion.
func (f *Fenwick) Add(idx, val int) {
 for idx <= f.n {
  f.bit[idx] += val
  idx += idx & -idx
 }
}

// Prefix sum because I need how many obstacles are on the left side.
func (f *Fenwick) Sum(idx int) int {
 res := 0
 for idx > 0 {
  res += f.bit[idx]
  idx -= idx & -idx
 }
 return res
}

// Find the smallest index whose prefix sum is at least k.
func (f *Fenwick) Kth(k int) int {
 idx := 0
 step := 1
 for (step << 1) <= f.n {
  step <<= 1
 }
 for d := step; d > 0; d >>= 1 {
  next := idx + d
  if next <= f.n && f.bit[next] < k {
   idx = next
   k -= f.bit[next]
  }
 }
 return idx + 1
}

type SegTree struct {
 n    int
 tree []int
}

// I use a segment tree because I need a prefix maximum over gap lengths.
func NewSegTree(n int) *SegTree {
 size := 4
 if 4*n > size {
  size = 4 * n
 }
 return &SegTree{
  n:    n,
  tree: make([]int, size),
 }
}

// Update one position because only one stored gap changes at a time.
func (s *SegTree) update(node, l, r, pos, val int) {
 if l == r {
  s.tree[node] = val
  return
 }
 mid := (l + r) >> 1
 if pos <= mid {
  s.update(node<<1, l, mid, pos, val)
 } else {
  s.update(node<<1|1, mid+1, r, pos, val)
 }
 left, right := s.tree[node<<1], s.tree[node<<1|1]
 if left > right {
  s.tree[node] = left
 } else {
  s.tree[node] = right
 }
}

// Query a prefix because every useful gap ends at or before x.
func (s *SegTree) query(node, l, r, ql, qr int) int {
 if ql > r || qr < l {
  return 0
 }
 if ql <= l && r <= qr {
  return s.tree[node]
 }
 mid := (l + r) >> 1
 a := s.query(node<<1, l, mid, ql, qr)
 b := s.query(node<<1|1, mid+1, r, ql, qr)
 if a > b {
  return a
 }
 return b
}

func getResults(queries [][]int) []bool {
 mx := 0
 for _, q := range queries {
  if q[1] > mx {
   mx = q[1]
  }
 }

 fw := NewFenwick(mx + 2)
 st := NewSegTree(mx + 1)

 // Sentinel obstacle at 0 keeps predecessor logic simple.
 fw.Add(1, 1)

 ans := make([]bool, 0, len(queries))

 for _, q := range queries {
  t := q[0]
  x := q[1]

  if t == 1 {
   // Count occupied positions strictly smaller than x.
   leftCount := fw.Sum(x)
   leftPos := fw.Kth(leftCount) - 1

   // Find the next occupied position after x if it exists.
   occupiedUpToX := fw.Sum(x + 1)
   totalOccupied := fw.Sum(mx + 2)
   rightPos := -1
   if occupiedUpToX < totalOccupied {
    rightPos = fw.Kth(occupiedUpToX + 1) - 1
   }

   // x becomes the end of a new gap.
   st.update(1, 0, mx, x, x-leftPos)

   // The next obstacle now sees a shorter gap.
   if rightPos != -1 {
    st.update(1, 0, mx, rightPos, rightPos-x)
   }

   // Mark x as occupied.
   fw.Add(x+1, 1)
  } else {
   sz := q[2]

   // The last obstacle strictly before x gives the free suffix up to x.
   leftCount := fw.Sum(x)
   leftPos := fw.Kth(leftCount) - 1

   // Prefix maximum of gaps ending at or before x.
   bestPrefix := st.query(1, 0, mx, 0, x)

   ans = append(ans, x-leftPos >= sz || bestPrefix >= sz)
  }
 }

 return ans
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic remains exactly the same across all five languages.

### Step 1: Store Obstacles

I keep all obstacle positions inside an ordered structure.

This allows me to quickly locate neighboring obstacles whenever a new obstacle is inserted.

### Step 2: Maintain Gap Information

For every obstacle position, I store the length of the gap that ends at that obstacle.

Example:

Obstacle positions:

```text
0 --- 2 ------- 7
```

Stored gaps:

```text
2
5
```

Because:

```text
2 - 0 = 2
7 - 2 = 5
```

### Step 3: Build Segment Tree

The Segment Tree stores maximum gap values.

This allows me to quickly answer:

> What is the largest available gap before position x?

without scanning all obstacles.

### Step 4: Process Insert Queries

When inserting a new obstacle:

```text
0 ------- 7
```

Insert:

```text
2
```

The original gap:

```text
7
```

becomes:

```text
2
5
```

Only nearby values change.

Everything else remains untouched.

This is why updates are efficient.

### Step 5: Process Placement Queries

For query:

```text
[2, x, sz]
```

I check:

1. The free space from the previous obstacle to `x`.
2. The largest complete gap ending before `x`.

If either one is at least `sz`, a valid placement exists.

### Step 6: Return Answers

Every type-2 query contributes one boolean result.

These values are collected and returned in order.

The implementation details differ slightly between C++, Java, JavaScript, Python, and Go, but the algorithm remains identical.

## Examples

### Example 1

Input

```text
queries = [[1,2],[2,3,3],[2,3,1],[2,2,2]]
```

Output

```text
[false,true,true]
```

Explanation

After inserting obstacle `2`:

```text
0 --- 2 --- 3
```

For:

```text
[2,3,3]
```

No segment of length `3` exists.

Result:

```text
false
```

For:

```text
[2,3,1]
```

Length `1` fits.

Result:

```text
true
```

For:

```text
[2,2,2]
```

Length `2` fits exactly.

Result:

```text
true
```

---

### Example 2

Input

```text
queries = [[1,7],[2,7,6],[1,2],[2,7,5],[2,7,6]]
```

Output

```text
[true,true,false]
```

Explanation

After all insertions:

```text
0 --- 2 ----- 7
```

Available gaps:

```text
2
5
```

A block of length `5` fits.

A block of length `6` does not.

---

### Example 3

Input

```text
queries = [[1,5],[1,10],[2,10,5]]
```

Output

```text
[true]
```

Explanation

Gap:

```text
0 -> 5 = 5
```

A block of size `5` fits exactly.

## How to Use / Run Locally

### C++

Compile

```bash
g++ solution.cpp -O2 -std=c++17
```

Run

```bash
./a.out
```

### Java

Compile

```bash
javac Solution.java
```

Run

```bash
java Solution
```

### JavaScript

Run

```bash
node solution.js
```

### Python3

Run

```bash
python solution.py
```

### Go

Run

```bash
go run solution.go
```

## Notes & Optimizations

* A brute-force approach is far too slow because each query may require scanning many obstacles.
* Maintaining gaps dynamically avoids repeated recomputation.
* The Ordered Set gives fast predecessor and successor lookup.
* The Segment Tree gives fast maximum-gap queries.
* Only nearby gaps change after insertion, which keeps updates efficient.
* This solution easily handles the maximum constraints.
* An alternative offline approach exists using reverse processing, but the Segment Tree + Ordered Set approach is usually easier to understand and implement.

SEO Keywords:

* LeetCode 3161 Solution
* Block Placement Queries
* Segment Tree Solution
* Ordered Set Approach
* Range Maximum Query
* Competitive Programming
* Data Structures and Algorithms
* DSA Problem Solution
* LeetCode Hard Problem
* Interval Query Processing
* Dynamic Gap Maintenance

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
