# 3651. Minimum Cost Path with Teleportations

## Table of Contents

* ## Problem Summary

* ## Constraints

* ## Intuition

* ## Approach

* ## Data Structures Used

* ## Operations & Behavior Summary

* ## Complexity

* ## Multi-language Solutions

  * ### C++

  * ### Java

  * ### JavaScript

  * ### Python3

  * ### Go

* ## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

* ## Examples

* ## How to use / Run locally

* ## Notes & Optimizations

* ## Author

---

## Problem Summary

I am given an `m x n` 2D integer array `grid` and an integer `k`. I start at the top-left cell `(0,0)` and my goal is to reach the bottom-right cell `(m-1, n-1)`.

There are two types of moves:

* **Normal move:** move right or down. Moving into a cell `(i,j)` adds `grid[i][j]` to my cost.
* **Teleportation:** I may teleport from any cell `u` to any cell `v` **for zero cost** if `grid[v] <= grid[u]`. I can teleport at most `k` times.

Return the **minimum total cost** to reach `(m-1,n-1)` from `(0,0)` using up to `k` teleports.

---

## Constraints

* `2 <= m, n <= 80`
* `0 <= grid[i][j] <= 10^4`
* `0 <= k <= 10`
* The grid size `m*n` ≤ 6400 (so `m*n*log(m*n)` is okay for k ≤ 10)

---

## Intuition

I thought about the normal DP for grid (only right/down): the cost to reach a cell is the min of coming from top or left plus the destination cell's value. Teleports change this: I can jump from any `u` to any `v` with `grid[v] <= grid[u]` for free. That suggests that if I know the best cost to reach every cell with up to `t` teleports, then when allowing one more teleport (`t+1`), I can place myself at any `v` with cost equal to the minimum `dp[u]` among all `u` with `grid[u] >= grid[v]`. So I group cells by value and scan values descending to compute that efficiently, then propagate normal moves (right/down) to compute new dp. Repeat up to `k` teleports.

---

## Approach

1. Compute base DP `dp` for 0 teleports using normal right/down propagation starting from `(0,0)` with cost `0`.
2. Create a list of all cells `(value, i, j)` and sort in **descending** order of `value`. Group equal values.
3. For each teleport step `t` from `0` to `k-1`:

   * Maintain `running_min` = minimum `dp` among all cells with value ≥ current value (scanning descending).
   * For each cell `v`, compute `start[v] = min(dp[v], running_min)` — this means either I don't teleport, or I teleport from the best suitable `u`.
   * From `start` as starting costs, propagate normal right/down moves (another DP) to get new `dp`.
4. After `k` iterations `dp[m-1][n-1]` is the answer.

This avoids enumerating all teleport pairs and is efficient: each teleport iteration is O(m*n log(m*n)) due to sorting (can be optimized with bucketing), and propagation is O(m*n).

---

## Data Structures Used

* 2D arrays/matrices (`dp`, `start`, `dp2`) sized `m x n`.
* A list/array of cell triples `(value, i, j)` sorted descending to group by value.
* Primitive types for storing large costs (`long`/`long long`/`int64` or big `INF`).

---

## Operations & Behavior Summary

* **Right/Down propagation:** standard DP from top-left to bottom-right (cost to enter neighbor adds neighbor value).
* **Teleport step computation:** group cells by value descending; compute `running_min` to represent best teleport source for any cell with value ≤ current group value.
* **Repeat teleport + propagate** up to `k` times.

---

## Complexity

* **Time Complexity:**

  * Sorting cells: `O(m*n log(m*n))` (done once or equivalently once per teleport if re-sorted).
  * For each of the `k` teleports: grouping + scanning = `O(m*n)`, propagation = `O(m*n)`.
  * Combined: `O(k * m * n + m*n log(m*n))`. Since `k ≤ 10` and `m*n ≤ 6400`, this is efficient.
* **Space Complexity:** `O(m * n)` for dp arrays and cell list.

---

## Multi-language Solutions

### C++

```c++
/* C++ solution */
#include <bits/stdc++.h>
using namespace std;
using ll = long long;
const ll INF = (ll)4e18;

class Solution {
public:
    int minCost(vector<vector<int>>& grid, int k) {
        int m = grid.size(), n = grid[0].size();
        vector<vector<ll>> dp(m, vector<ll>(n, INF));
        dp[0][0] = 0;
        // base DP (no teleport)
        for (int i = 0; i < m; ++i) {
            for (int j = 0; j < n; ++j) {
                if (i > 0) dp[i][j] = min(dp[i][j], dp[i-1][j] + grid[i][j]);
                if (j > 0) dp[i][j] = min(dp[i][j], dp[i][j-1] + grid[i][j]);
            }
        }

        // prepare cells sorted by value desc
        vector<tuple<int,int,int>> cells;
        cells.reserve(m*n);
        for (int i=0;i<m;++i) for (int j=0;j<n;++j) cells.emplace_back(grid[i][j], i, j);
        sort(cells.begin(), cells.end(), [](auto &a, auto &b){ return get<0>(a) > get<0>(b); });

        for (int step = 0; step < k; ++step) {
            vector<vector<ll>> start(m, vector<ll>(n, INF));
            ll running_min = INF;
            int idx = 0;
            while (idx < (int)cells.size()) {
                int val = get<0>(cells[idx]);
                int j = idx;
                ll min_group = INF;
                while (j < (int)cells.size() && get<0>(cells[j]) == val) {
                    int ii = get<1>(cells[j]);
                    int jj = get<2>(cells[j]);
                    min_group = min(min_group, dp[ii][jj]);
                    ++j;
                }
                running_min = min(running_min, min_group);
                for (int p = idx; p < j; ++p) {
                    int ii = get<1>(cells[p]);
                    int jj = get<2>(cells[p]);
                    start[ii][jj] = min(dp[ii][jj], running_min);
                }
                idx = j;
            }

            vector<vector<ll>> dp2(m, vector<ll>(n, INF));
            for (int i=0;i<m;++i) {
                for (int j=0;j<n;++j) {
                    dp2[i][j] = min(dp2[i][j], start[i][j]);
                    if (i+1 < m && dp2[i][j] < INF) dp2[i+1][j] = min(dp2[i+1][j], dp2[i][j] + grid[i+1][j]);
                    if (j+1 < n && dp2[i][j] < INF) dp2[i][j+1] = min(dp2[i][j+1], dp2[i][j] + grid[i][j+1]);
                }
            }
            dp.swap(dp2);
        }

        return (int)dp[m-1][n-1];
    }
};
```

---

### Java

```java
// Java solution
import java.util.*;

class Solution {
    public int minCost(int[][] grid, int k) {
        int m = grid.length, n = grid[0].length;
        final long INF = (long)4e18;
        long[][] dp = new long[m][n];
        for (int i=0;i<m;i++) Arrays.fill(dp[i], INF);

        // base dp (no teleport)
        dp[0][0] = 0;
        for (int i=0;i<m;i++){
            for (int j=0;j<n;j++){
                if (i>0) dp[i][j] = Math.min(dp[i][j], dp[i-1][j] + grid[i][j]);
                if (j>0) dp[i][j] = Math.min(dp[i][j], dp[i][j-1] + grid[i][j]);
            }
        }

        ArrayList<int[]> cells = new ArrayList<>();
        for (int i=0;i<m;i++) for (int j=0;j<n;j++) cells.add(new int[]{grid[i][j], i, j});
        cells.sort((a,b)-> Integer.compare(b[0], a[0]));

        for (int step=0; step<k; step++) {
            long[][] start = new long[m][n];
            for (int i=0;i<m;i++) Arrays.fill(start[i], INF);

            long runningMin = INF;
            int idx = 0;
            while (idx < cells.size()) {
                int val = cells.get(idx)[0];
                int j = idx;
                long minGroup = INF;
                while (j < cells.size() && cells.get(j)[0] == val) {
                    int ii = cells.get(j)[1], jj = cells.get(j)[2];
                    minGroup = Math.min(minGroup, dp[ii][jj]);
                    j++;
                }
                runningMin = Math.min(runningMin, minGroup);
                for (int p=idx; p<j; p++) {
                    int ii = cells.get(p)[1], jj = cells.get(p)[2];
                    start[ii][jj] = Math.min(dp[ii][jj], runningMin);
                }
                idx = j;
            }

            long[][] dp2 = new long[m][n];
            for (int i=0;i<m;i++) Arrays.fill(dp2[i], INF);
            for (int i=0;i<m;i++){
                for (int j=0;j<n;j++){
                    if (start[i][j] < dp2[i][j]) dp2[i][j] = start[i][j];
                    if (i+1 < m && dp2[i][j] < INF) dp2[i+1][j] = Math.min(dp2[i+1][j], dp2[i][j] + grid[i+1][j]);
                    if (j+1 < n && dp2[i][j] < INF) dp2[i][j+1] = Math.min(dp2[i][j+1], dp2[i][j] + grid[i][j+1]);
                }
            }
            dp = dp2;
        }

        return (int)dp[m-1][n-1];
    }
}
```

---

### JavaScript

```javascript
/**
 * JavaScript solution
 * @param {number[][]} grid
 * @param {number} k
 * @return {number}
 */
var minCost = function(grid, k) {
    const m = grid.length, n = grid[0].length;
    const INF = Number.MAX_SAFE_INTEGER / 4;

    // dp (no teleports)
    let dp = Array.from({length: m}, ()=> Array(n).fill(INF));
    dp[0][0] = 0;
    for (let i=0;i<m;i++){
        for (let j=0;j<n;j++){
            if (i>0) dp[i][j] = Math.min(dp[i][j], dp[i-1][j] + grid[i][j]);
            if (j>0) dp[i][j] = Math.min(dp[i][j], dp[i][j-1] + grid[i][j]);
        }
    }

    // prepare sorted cells by value desc
    const cells = [];
    for (let i=0;i<m;i++) for (let j=0;j<n;j++) cells.push([grid[i][j], i, j]);
    cells.sort((a,b) => b[0] - a[0]);

    for (let step=0; step<k; step++){
        const start = Array.from({length: m}, ()=> Array(n).fill(INF));
        let runningMin = INF;
        let idx = 0;
        while (idx < cells.length) {
            const val = cells[idx][0];
            let j = idx;
            let minGroup = INF;
            while (j < cells.length && cells[j][0] === val) {
                const [_, ii, jj] = cells[j];
                minGroup = Math.min(minGroup, dp[ii][jj]);
                j++;
            }
            runningMin = Math.min(runningMin, minGroup);
            for (let p = idx; p < j; p++) {
                const [_, ii, jj] = cells[p];
                start[ii][jj] = Math.min(dp[ii][jj], runningMin);
            }
            idx = j;
        }

        const dp2 = Array.from({length: m}, ()=> Array(n).fill(INF));
        for (let i=0;i<m;i++){
            for (let j=0;j<n;j++){
                if (start[i][j] < dp2[i][j]) dp2[i][j] = start[i][j];
                if (i+1 < m && dp2[i][j] < INF) dp2[i+1][j] = Math.min(dp2[i+1][j], dp2[i][j] + grid[i+1][j]);
                if (j+1 < n && dp2[i][j] < INF) dp2[i][j+1] = Math.min(dp2[i][j+1], dp2[i][j] + grid[i][j+1]);
            }
        }
        dp = dp2;
    }

    return dp[m-1][n-1];
};
```

---

### Python3

```python3
# Python3 solution
from typing import List

class Solution:
    def minCost(self, grid: List[List[int]], k: int) -> int:
        INF = 10**18
        m, n = len(grid), len(grid[0])
        dp = [[INF]*n for _ in range(m)]
        dp[0][0] = 0
        # base DP: normal right/down moves
        for i in range(m):
            for j in range(n):
                if i > 0:
                    dp[i][j] = min(dp[i][j], dp[i-1][j] + grid[i][j])
                if j > 0:
                    dp[i][j] = min(dp[i][j], dp[i][j-1] + grid[i][j])

        cells = []
        for i in range(m):
            for j in range(n):
                cells.append((grid[i][j], i, j))
        cells.sort(reverse=True)  # descending by value

        for _ in range(k):
            start = [[INF]*n for _ in range(m)]
            running_min = INF
            idx = 0
            L = len(cells)
            while idx < L:
                val = cells[idx][0]
                j = idx
                min_group = INF
                while j < L and cells[j][0] == val:
                    _, ii, jj = cells[j]
                    min_group = min(min_group, dp[ii][jj])
                    j += 1
                running_min = min(running_min, min_group)
                for p in range(idx, j):
                    _, ii, jj = cells[p]
                    start[ii][jj] = min(dp[ii][jj], running_min)
                idx = j

            dp2 = [[INF]*n for _ in range(m)]
            for i in range(m):
                for j in range(n):
                    dp2[i][j] = min(dp2[i][j], start[i][j])
                    if dp2[i][j] < INF:
                        if i+1 < m:
                            dp2[i+1][j] = min(dp2[i+1][j], dp2[i][j] + grid[i+1][j])
                        if j+1 < n:
                            dp2[i][j+1] = min(dp2[i][j+1], dp2[i][j] + grid[i][j+1])
            dp = dp2

        return dp[m-1][n-1]
```

---

### Go

```go
// Go solution
package main
import (
 "sort"
)

func min(a,b int64) int64 { if a<b { return a }; return b }

func minCost(grid [][]int, k int) int {
 m := len(grid)
 n := len(grid[0])
 const INF int64 = 4e18

 dp := make([][]int64, m)
 for i:=0;i<m;i++ {
  dp[i] = make([]int64, n)
  for j:=0;j<n;j++ { dp[i][j] = INF }
 }
 dp[0][0] = 0
 for i:=0;i<m;i++ {
  for j:=0;j<n;j++ {
   if i>0 {
    val := dp[i-1][j] + int64(grid[i][j])
    if val < dp[i][j] { dp[i][j] = val }
   }
   if j>0 {
    val := dp[i][j-1] + int64(grid[i][j])
    if val < dp[i][j] { dp[i][j] = val }
   }
  }
 }

 type Cell struct { v, i, j int }
 cells := make([]Cell, 0, m*n)
 for i:=0;i<m;i++ {
  for j:=0;j<n;j++ {
   cells = append(cells, Cell{grid[i][j], i, j})
  }
 }
 sort.Slice(cells, func(a,b int) bool { return cells[a].v > cells[b].v })

 for step:=0; step<k; step++ {
  start := make([][]int64, m)
  for i:=0;i<m;i++ {
   start[i] = make([]int64, n)
   for j:=0;j<n;j++ { start[i][j] = INF }
  }
  runningMin := INF
  idx := 0
  for idx < len(cells) {
   val := cells[idx].v
   j := idx
   minGroup := INF
   for j < len(cells) && cells[j].v == val {
    ii := cells[j].i; jj := cells[j].j
    if dp[ii][jj] < minGroup { minGroup = dp[ii][jj] }
    j++
   }
   if minGroup < runningMin { runningMin = minGroup }
   for p := idx; p < j; p++ {
    ii := cells[p].i; jj := cells[p].j
    if dp[ii][jj] < runningMin { start[ii][jj] = dp[ii][jj] } else { start[ii][jj] = runningMin }
   }
   idx = j
  }

  dp2 := make([][]int64, m)
  for i:=0;i<m;i++ {
   dp2[i] = make([]int64, n)
   for j:=0;j<n;j++ { dp2[i][j] = INF }
  }
  for i:=0;i<m;i++ {
   for j:=0;j<n;j++ {
    if start[i][j] < dp2[i][j] { dp2[i][j] = start[i][j] }
    if dp2[i][j] < INF {
     if i+1 < m {
      v := dp2[i][j] + int64(grid[i+1][j])
      if v < dp2[i+1][j] { dp2[i+1][j] = v }
     }
     if j+1 < n {
      v := dp2[i][j] + int64(grid[i][j+1])
      if v < dp2[i][j+1] { dp2[i][j+1] = v }
     }
    }
   }
  }
  dp = dp2
 }

 return int(dp[m-1][n-1])
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

I'll explain the algorithm in a way that's easy to follow — the code in every language follows the same logical steps.

1. **Base DP (no teleport)**

   * Initialize a `dp` matrix with very large values (`INF`).
   * Set `dp[0][0] = 0` because starting cost is 0 at top-left.
   * For every cell `(i,j)` do:

     * `dp[i][j] = min(dp[i][j], dp[i-1][j] + grid[i][j])` if `i>0`
     * `dp[i][j] = min(dp[i][j], dp[i][j-1] + grid[i][j])` if `j>0`
   * This computes the best cost to reach each cell using only normal moves.

2. **Prepare cells**

   * Build a list `cells` storing `(grid[i][j], i, j)` for every cell.
   * Sort `cells` by `value` descending. This helps to compute, for each target value, the best teleport source among cells with value ≥ target.

3. **One teleport iteration (repeat k times)**

   * `start` matrix: the cost to *be at* that cell after using one additional teleport (or not).
   * `running_min` variable: while scanning cells in descending value, it stores the smallest `dp[u]` seen so far among processed cells (i.e., among cells with value ≥ current).
   * For a group of equal-value cells:

     * Compute `min_group = min(dp[u])` for u in this group.
     * Update `running_min = min(running_min, min_group)`.
     * For every cell `v` in this value-group, set `start[v] = min(dp[v], running_min)`. This accounts for:

       * Not teleporting to `v` (keep `dp[v]`), or
       * Teleporting from some `u` with `grid[u] >= grid[v]` at zero additional cost (`running_min`).
   * After computing `start`, run right/down DP propagation starting from `start` to compute `dp2`:

     * Initialize `dp2` with `INF`.
     * For each `(i,j)`, set `dp2[i][j] = min(dp2[i][j], start[i][j])`.
     * Then from `dp2[i][j]` propagate:

       * `dp2[i+1][j] = min(dp2[i+1][j], dp2[i][j] + grid[i+1][j])`
       * `dp2[i][j+1] = min(dp2[i][j+1], dp2[i][j] + grid[i][j+1])`
   * Set `dp = dp2`. This `dp` now represents best costs using up to `t+1` teleports.

4. **Answer**

   * After repeating up to `k` teleports, return `dp[m-1][n-1]`.

**Why grouping by value helps:** teleport condition `grid[v] <= grid[u]` is monotonic by value. Scanning values descending allows one pass to compute `min_{u: grid[u] >= val} dp[u]` as `running_min`. That avoids considering all pairs `u->v`.

---

## Examples

Example 1:

```
Input: grid = [[1,3,3],[2,5,4],[4,3,5]], k = 2
Output: 7
Explanation: One optimal path uses a teleport to lower cost as described in problem examples.
```

Example 2:

```
Input: grid = [[1,2],[2,3],[3,4]], k = 1
Output: 9
```

(You can test these examples with the provided code by writing a small main/driver or adapting to your platform's testing harness.)

---

## How to use / Run locally

General notes: each snippet is a function/class that you can put into your LeetCode-style environment, or wrap with a `main`/test harness.

**C++**

* Put `Solution` class into file `solution.cpp`.
* Compile & run:
  `g++ -std=c++17 solution.cpp -O2 -o solution && ./solution`
  (If you provide a `main` with test cases.)

**Java**

* Save `Solution` class into `Solution.java` and supply a `main` to call `minCost`.
* Compile: `javac Solution.java`
* Run: `java Solution`

**JavaScript (Node)**

* Save code in `solution.js`. Add a small test invocation below the function:

  ```javascript
  console.log(minCost([[1,3,3],[2,5,4],[4,3,5]], 2));
  ```

* Run: `node solution.js`

**Python3**

* Save class in `solution.py` and append a test:

  ```python
  if __name__ == "__main__":
      grid = [[1,3,3],[2,5,4],[4,3,5]]
      print(Solution().minCost(grid, 2))
  ```

* Run: `python3 solution.py`

**Go**

* Put code in `main.go`, write a `main()` that calls `minCost` with a sample and prints the result.
* Run: `go run main.go`

---

## Notes & Optimizations

* Sorting `m*n` cells is `O(m*n log(m*n))`. Because `grid[i][j] ≤ 10^4`, we can do bucket/counting sort by value to replace sorting with `O(m*n + V)` where `V` is max value range (10^4) — this makes each teleport iteration strictly `O(m*n)`.
* If memory is tight, you can reuse `dp` arrays in place but be careful with overwrites during propagation.
* Because `k ≤ 10`, repeating the grouping/propagation `k` times is fine.
* Using 64-bit integers (`long long` / `long` / `int64`) is important to avoid overflow of accumulated costs.
* This approach avoids O((m*n)^2) teleport pair enumeration.

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
