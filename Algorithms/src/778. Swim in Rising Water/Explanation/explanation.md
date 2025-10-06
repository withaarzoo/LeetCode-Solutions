# 778. Swim in Rising Water

---

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
* [How to use / Run locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

---

## Problem Summary

Given an `n x n` grid where `grid[i][j]` represents the elevation of point `(i, j)`. Water level starts at `t = 0` and rises over time. At time `t`, you can move onto any cell with elevation `<= t`. You start at `(0, 0)` and want to reach `(n-1, n-1)`. You can move in four directions. Return the minimum time `t` so that it's possible to reach the bottom-right cell.

Put simply: find the minimum possible maximum elevation along any path from the top-left to bottom-right. That minimum maximum elevation equals the earliest time I can reach the end.

---

## Constraints

* `n == grid.length`
* `n == grid[i].length`
* `1 <= n <= 50`
* `0 <= grid[i][j] < n * n`
* All values of `grid[i][j]` are unique.

---

## Intuition

I thought about the problem like a race against rising water. To go from start to end I must wait until the water reaches the maximum elevation encountered on my path. So, the problem reduces to: **find a path whose maximum elevation is as small as possible**. This is the same as minimizing the maximum weight on a path, so I chose a Dijkstra-style approach where the "distance" to a cell is the maximum elevation required to reach it.

---

## Approach

1. I track the earliest time `t` I can stand on each cell — that equals the maximum elevation encountered along the path used to reach that cell.
2. I use a min-heap (priority queue) keyed by `t`. Start with `(grid[0][0], 0, 0)`.
3. Repeatedly pop the cell with smallest `t`.

   * If it's the destination, return `t`.
   * Otherwise, try the 4 neighbors: the time to enter a neighbor equals `max(t, grid[neighbor])`. Push that into heap.
4. Use a `visited` set so that once we pop a cell (first pop is optimal for that cell), we don't process it again.
5. This is like Dijkstra where path cost = maximum node value along path.

---

## Data Structures Used

* **Min-Heap / Priority Queue** — store tuples `(time, row, col)` and always expand the smallest `time`.
* **Visited boolean grid** — to mark cells that are finalized (first time popped).
* **Grid** — input 2D integer array.

---

## Operations & Behavior Summary

* Push start cell with time `grid[0][0]`.
* Pop lowest-time cell, finalize it.
* For each adjacent not-yet-finalized neighbor, compute `new_time = max(curr_time, grid[nr][nc])` and push to heap.
* Terminate when we finalize `(n-1, n-1)`.

This ensures that when a cell is popped for the first time, we have the minimal possible `max-elevation` path to that cell.

---

## Complexity

* **Time Complexity:** `O(n^2 log n^2)` = `O(n^2 log n)` where `n` is grid side length. There are `n^2` cells; each may be pushed/popped into/from the heap; heap operations cost `O(log(n^2))`.
* **Space Complexity:** `O(n^2)` for visited array and the heap in the worst case.

---

## Multi-language Solutions

### C++

```c++
#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    int swimInWater(vector<vector<int>>& grid) {
        int n = grid.size();
        vector<vector<bool>> vis(n, vector<bool>(n, false));
        using T = tuple<int,int,int>; // (time, r, c)
        priority_queue<T, vector<T>, greater<T>> pq;
        pq.emplace(grid[0][0], 0, 0);
        int dirs[4][2] = {{1,0},{-1,0},{0,1},{0,-1}};
        
        while (!pq.empty()) {
            auto [t, r, c] = pq.top(); pq.pop();
            if (vis[r][c]) continue;
            vis[r][c] = true;
            if (r == n-1 && c == n-1) return t;
            for (auto &d : dirs) {
                int nr = r + d[0], nc = c + d[1];
                if (nr >= 0 && nr < n && nc >= 0 && nc < n && !vis[nr][nc]) {
                    int nt = max(t, grid[nr][nc]);
                    pq.emplace(nt, nr, nc);
                }
            }
        }
        return -1; // theoretically unreachable
    }
};
```

---

### Java

```java
import java.util.PriorityQueue;

class Solution {
    public int swimInWater(int[][] grid) {
        int n = grid.length;
        boolean[][] vis = new boolean[n][n];
        PriorityQueue<int[]> pq = new PriorityQueue<>((a,b) -> Integer.compare(a[0], b[0]));
        pq.offer(new int[]{grid[0][0], 0, 0});
        int[][] dirs = {{1,0},{-1,0},{0,1},{0,-1}};
        
        while (!pq.isEmpty()) {
            int[] cur = pq.poll();
            int t = cur[0], r = cur[1], c = cur[2];
            if (vis[r][c]) continue;
            vis[r][c] = true;
            if (r == n-1 && c == n-1) return t;
            for (int[] d : dirs) {
                int nr = r + d[0], nc = c + d[1];
                if (nr >= 0 && nr < n && nc >= 0 && nc < n && !vis[nr][nc]) {
                    int nt = Math.max(t, grid[nr][nc]);
                    pq.offer(new int[]{nt, nr, nc});
                }
            }
        }
        return -1;
    }
}
```

---

### JavaScript

> **Note:** In online judges (LeetCode) the function alone is enough. If you run multiple times in the same Node process, declaring helper classes at top-level may cause redeclaration errors. I keep the `MinHeap` inside the function to avoid that.

```javascript
/**
 * @param {number[][]} grid
 * @return {number}
 */
var swimInWater = function(grid) {
    const n = grid.length;
    const visited = Array.from({length: n}, () => Array(n).fill(false));

    // MinHeap inside function to avoid redeclaration in reused JS environment
    class MinHeap {
        constructor() { this.heap = []; }
        size() { return this.heap.length; }
        push(val) {
            this.heap.push(val);
            this._bubbleUp(this.heap.length - 1);
        }
        pop() {
            if (this.heap.length === 0) return undefined;
            const top = this.heap[0];
            const last = this.heap.pop();
            if (this.heap.length > 0) {
                this.heap[0] = last;
                this._bubbleDown(0);
            }
            return top;
        }
        _bubbleUp(i) {
            const h = this.heap;
            while (i > 0) {
                const p = (i - 1) >> 1;
                if (h[p][0] <= h[i][0]) break;
                [h[p], h[i]] = [h[i], h[p]];
                i = p;
            }
        }
        _bubbleDown(i) {
            const h = this.heap;
            const len = h.length;
            while (true) {
                let smallest = i;
                const l = 2 * i + 1, r = 2 * i + 2;
                if (l < len && h[l][0] < h[smallest][0]) smallest = l;
                if (r < len && h[r][0] < h[smallest][0]) smallest = r;
                if (smallest === i) break;
                [h[smallest], h[i]] = [h[i], h[smallest]];
                i = smallest;
            }
        }
    }

    const heap = new MinHeap();
    heap.push([grid[0][0], 0, 0]);
    const dirs = [[1,0],[-1,0],[0,1],[0,-1]];

    while (heap.size() > 0) {
        const [t, r, c] = heap.pop();
        if (visited[r][c]) continue;
        visited[r][c] = true;
        if (r === n-1 && c === n-1) return t;
        for (const [dr, dc] of dirs) {
            const nr = r + dr, nc = c + dc;
            if (nr >= 0 && nr < n && nc >= 0 && nc < n && !visited[nr][nc]) {
                heap.push([Math.max(t, grid[nr][nc]), nr, nc]);
            }
        }
    }
    return -1;
};
```

---

### Python3

```python
import heapq
from typing import List

class Solution:
    def swimInWater(self, grid: List[List[int]]) -> int:
        n = len(grid)
        visited = [[False]*n for _ in range(n)]
        heap = [(grid[0][0], 0, 0)]  # (time, r, c)
        dirs = [(1,0),(-1,0),(0,1),(0,-1)]
        
        while heap:
            t, r, c = heapq.heappop(heap)
            if visited[r][c]:
                continue
            visited[r][c] = True
            if r == n-1 and c == n-1:
                return t
            for dr, dc in dirs:
                nr, nc = r + dr, c + dc
                if 0 <= nr < n and 0 <= nc < n and not visited[nr][nc]:
                    nt = max(t, grid[nr][nc])
                    heapq.heappush(heap, (nt, nr, nc))
        return -1
```

---

### Go

```go
package main

import (
    "container/heap"
)

type Cell struct { t, r, c int }

type CellHeap []Cell
func (h CellHeap) Len() int { return len(h) }
func (h CellHeap) Less(i, j int) bool { return h[i].t < h[j].t }
func (h CellHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *CellHeap) Push(x interface{}) { *h = append(*h, x.(Cell)) }
func (h *CellHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

func swimInWater(grid [][]int) int {
    n := len(grid)
    visited := make([][]bool, n)
    for i := 0; i < n; i++ { visited[i] = make([]bool, n) }
    dirs := [4][2]int{{1,0},{-1,0},{0,1},{0,-1}}
    h := &CellHeap{}
    heap.Init(h)
    heap.Push(h, Cell{grid[0][0], 0, 0})

    for h.Len() > 0 {
        cell := heap.Pop(h).(Cell)
        t, r, c := cell.t, cell.r, cell.c
        if visited[r][c] { continue }
        visited[r][c] = true
        if r == n-1 && c == n-1 { return t }
        for _, d := range dirs {
            nr, nc := r + d[0], c + d[1]
            if nr >= 0 && nr < n && nc >= 0 && nc < n && !visited[nr][nc] {
                nt := t
                if grid[nr][nc] > nt { nt = grid[nr][nc] }
                heap.Push(h, Cell{nt, nr, nc})
            }
        }
    }
    return -1
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

Below I explain the algorithm steps and then walk through key lines in each language. The logic is identical across languages; only syntax/data-structure differences exist.

### Common algorithm recap (my voice)

1. I start by saying: I will use a priority queue that always expands the position reachable with the smallest `required_time` so far.
2. The `required_time` for a path equals the highest elevation along that path.
3. Start at `(0,0)` with `required_time = grid[0][0]`.
4. For each popped cell `(r,c)` with time `t`, I check neighbors `(nr,nc)`. The neighbor can be reached at time `max(t, grid[nr][nc])`.
5. I push that into the heap. The first time I pop the target, I return its `t`.

---

### C++ — key lines explained

* `priority_queue<T, vector<T>, greater<T>> pq;`
  I use `tuple<int,int,int>` and a min-heap by using the `greater` comparator so smallest `t` pops first.
* `pq.emplace(grid[0][0], 0, 0);`
  Start with time equals elevation at start.
* `auto [t, r, c] = pq.top(); pq.pop();`
  Pop the cell with smallest required time. This is greedy: we always finalize the smallest time cell first.
* `int nt = max(t, grid[nr][nc]); pq.emplace(nt, nr, nc);`
  The neighbor's required time is the maximum between current path's time and neighbor elevation.

---

### Java — key lines explained

* `PriorityQueue<int[]> pq = new PriorityQueue<>((a,b) -> Integer.compare(a[0], b[0]));`
  Each element is `{time, r, c}`; comparator sorts by `time`.
* `pq.offer(new int[]{grid[0][0], 0, 0});`
  Initialize.
* `int[] cur = pq.poll();` and `if (vis[r][c]) continue;`
  We skip already-finalized cells.
* `int nt = Math.max(t, grid[nr][nc]); pq.offer(new int[]{nt, nr, nc});`
  Enqueue neighbor with updated required time.

---

### JavaScript — key lines explained

* `class MinHeap { ... }`
  A simple binary heap of arrays `[time, r, c]`. I implement `_bubbleUp` and `_bubbleDown`.
* `heap.push([grid[0][0], 0, 0]);`
  Push start.
* `const [t, r, c] = heap.pop();`
  Pop minimal time element.
* `heap.push([Math.max(t, grid[nr][nc]), nr, nc]);`
  Push neighbor with new required time.

**Implementation detail:** I placed `MinHeap` inside `swimInWater` to avoid "Identifier already declared" errors when Node reuses the environment.

---

### Python — key lines explained

* `heap = [(grid[0][0], 0, 0)]` and `heapq.heappush(heap, (nt, nr, nc))`
  Python's `heapq` is a min-heap on the first tuple element, `time`.
* `t, r, c = heapq.heappop(heap)`
  Pop the smallest time entry.
* `nt = max(t, grid[nr][nc])`
  Neighbor reachable time.

---

### Go — key lines explained

* Implement `CellHeap` with required `heap.Interface` methods. We use `cell.t` as key.
* `heap.Push(h, Cell{grid[0][0], 0, 0})` to initialize.
* Use `visited` 2D slice to mark finalized nodes.
* `nt := t; if grid[nr][nc] > nt { nt = grid[nr][nc] }` to compute neighbor's required time.

---

## Examples

1. **Example 1**

   * Input: `[[0,2],[1,3]]`
   * Output: `3`
   * Explanation: I must wait until time `3` to be able to reach `(1,1)`.

2. **Example 2**

   * Input:

     ```
     [[0,1,2,3,4],
      [24,23,22,21,5],
      [12,13,14,15,16],
      [11,17,18,19,20],
      [10,9,8,7,6]]
     ```
   * Output: `16`
   * Explanation: The path requires waiting until time `16` to connect start and finish.

---

## How to use / Run locally

These are simple instructions to run sample test cases locally; copy the language snippet into a file and run the commands.

### C++

Save as `main.cpp` — include a small `main()` wrapper to test:

```c++
#include <bits/stdc++.h>
using namespace std;
// (paste the Solution class here)

int main() {
    Solution s;
    vector<vector<int>> grid = {{0,2},{1,3}};
    cout << s.swimInWater(grid) << endl; // prints 3
    return 0;
}
```

Compile & run:

```bash
g++ -std=c++17 main.cpp -O2 -o swim && ./swim
```

### Java

Save in `Main.java`:

```java
// paste Solution class above or inline here
public class Main {
    public static void main(String[] args) {
        Solution sol = new Solution();
        int[][] grid = {{0,2},{1,3}};
        System.out.println(sol.swimInWater(grid)); // 3
    }
}
```

Compile & run:

```bash
javac Main.java && java Main
```

### JavaScript (Node)

Save as `swim.js`:

```javascript
// paste swimInWater function here
const grid = [[0,2],[1,3]];
console.log(swimInWater(grid)); // 3
```

Run:

```bash
node swim.js
```

### Python3

Save as `swim.py`:

```python
# paste Solution class here
if __name__ == "__main__":
    grid = [[0,2],[1,3]]
    sol = Solution()
    print(sol.swimInWater(grid))  # 3
```

Run:

```bash
python3 swim.py
```

### Go

Save as `main.go`:

```go
package main
import "fmt"

// paste cell heap and swimInWater code here

func main() {
    grid := [][]int{{0,2},{1,3}}
    fmt.Println(swimInWater(grid)) // 3
}
```

Run:

```bash
go run main.go
```

---

## Notes & Optimizations

* Alternative approach: **Binary search + BFS**. Binary-search time `T` and check connectivity using cells `<= T`. Complexity `O(n^2 log V)` where `V` is value domain (or `n^2`). That can sometimes be simpler to implement.
* Dijkstra-style solution used here avoids repeated full grid traversals — it tends to be efficient.
* Minor micro-optimizations:

  * Early return when start equals target (`n==1`).
  * Use bit-packed `visited` or single-array indexing if micro memory/perf is needed.
* In JavaScript, avoid top-level class declarations if running multiple tests in the same Node process — define helper classes inside the function.

---

## Author

[Aarzoo Islam](https://bento.me/withaarzoo)
