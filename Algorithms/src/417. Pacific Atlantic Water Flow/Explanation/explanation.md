# 417. Pacific Atlantic Water Flow

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
* [Step-by-step Detailed Explanation (C++, Java, JS, Python3, Go)](#step-by-step-detailed-explanation-c-java-js-python3-go)
* [Examples](#examples)
* [How to use / Run locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

---

## Problem Summary

We are given an `m x n` grid `heights` where `heights[r][c]` is the altitude of the cell `(r, c)`. Rainwater at a cell can flow to its 4-directional neighboring cell (north, south, east, west) **only** if the neighbor's height is **less than or equal** to the current cell's height.

The island's top and left edges border the **Pacific Ocean**. The island's bottom and right edges border the **Atlantic Ocean**. Return a list of coordinates of all cells where water can flow to **both** oceans.

---

## Constraints

* `m == heights.length`
* `n == heights[r].length`
* `1 <= m, n <= 200`
* `0 <= heights[r][c] <= 10^5`

---

## Intuition

I thought about the direction of water flow. Water flows from higher/equal height to lower/equal height. Instead of trying to check from every cell whether it can reach both oceans (which would require searching outward from each cell), I reversed the idea: **start from the oceans** and move *uphill* (to neighbors that are higher or equal). If I can reach a cell from the Pacific border when walking uphill and also reach it from the Atlantic border when walking uphill, then that cell can let water flow to both oceans.

So I perform two flood-fills (BFS/DFS):

* One from all Pacific-border cells.
* One from all Atlantic-border cells.
  The intersection of visited cells is the answer.

---

## Approach

1. If the grid is empty, return `[]`.
2. Create two boolean grids `pacific` and `atlantic` the same size as `heights`. These mark cells reachable from each ocean when moving *uphill* (i.e., to neighbors with `height >= current`).
3. Start BFS (or iterative DFS) from all Pacific border cells (top row and left column). Mark visited cells in `pacific`. Movement condition: neighbor height must be `>=` current height.
4. Start BFS from all Atlantic border cells (bottom row and right column). Mark visited cells in `atlantic`.
5. The answer is all coordinates `(i, j)` such that `pacific[i][j]` and `atlantic[i][j]` are both `true`.

I prefer BFS from borders because it's iterative (no recursion depth issue) and easy to reason about.

---

## Data Structures Used

* `queue` (BFS) — to perform level-order expansion from ocean borders.
* `boolean` 2D arrays — `pacific` and `atlantic` to mark visited cells for each ocean.
* `directions` array — for iterating 4-neighbors.

---

## Operations & Behavior Summary

* Start BFS from all Pacific border cells simultaneously: these are `(0, j)` for all `j` and `(i, 0)` for all `i`.
* Start BFS from all Atlantic border cells simultaneously: these are `(m-1, j)` for all `j` and `(i, n-1)` for all `i`.
* In BFS expansion, move from `(r, c)` to `(nr, nc)` only if:

  * `(nr, nc)` is inside the grid,
  * it has not been visited for the current ocean,
  * and `heights[nr][nc] >= heights[r][c]` (we move "uphill" when working from ocean-to-land).
* Cells reachable from both BFS runs are added to the final list.

---

## Complexity

* **Time Complexity:** `O(m * n)` where `m` is number of rows and `n` is number of columns.
  *Reason:* Each BFS (Pacific & Atlantic) visits each cell at most once, so total visits are `<= 2*m*n`. All neighbor checks are constant time.

* **Space Complexity:** `O(m * n)` for the two boolean `visited` arrays and the BFS queue which in the worst case can hold `O(m*n)` items.

---

## Multi-language Solutions

### C++

```c++
#include <vector>
#include <queue>
using namespace std;

class Solution {
public:
    vector<vector<int>> pacificAtlantic(vector<vector<int>>& heights) {
        if (heights.empty() || heights[0].empty()) return {};
        int m = heights.size(), n = heights[0].size();

        vector<vector<bool>> pac(m, vector<bool>(n, false));
        vector<vector<bool>> atl(m, vector<bool>(n, false));

        queue<pair<int,int>> q;

        // Pacific: top row and left column
        for (int j = 0; j < n; ++j) { q.push({0, j}); pac[0][j] = true; }
        for (int i = 1; i < m; ++i) { q.push({i, 0}); pac[i][0] = true; }
        bfs(heights, q, pac);

        // Atlantic: bottom row and right column
        queue<pair<int,int>> q2;
        for (int j = 0; j < n; ++j) { q2.push({m-1, j}); atl[m-1][j] = true; }
        for (int i = 0; i < m-1; ++i) { q2.push({i, n-1}); atl[i][n-1] = true; }
        bfs(heights, q2, atl);

        vector<vector<int>> res;
        for (int i = 0; i < m; ++i)
            for (int j = 0; j < n; ++j)
                if (pac[i][j] && atl[i][j]) res.push_back({i, j});
        return res;
    }

private:
    void bfs(const vector<vector<int>>& heights, queue<pair<int,int>>& q, vector<vector<bool>>& visited) {
        int m = heights.size(), n = heights[0].size();
        const int dirs[4][2] = {{1,0},{-1,0},{0,1},{0,-1}};
        while (!q.empty()) {
            auto cur = q.front(); q.pop();
            int r = cur.first, c = cur.second;
            for (int k = 0; k < 4; ++k) {
                int nr = r + dirs[k][0], nc = c + dirs[k][1];
                if (nr < 0 || nr >= m || nc < 0 || nc >= n) continue;
                if (visited[nr][nc]) continue;
                if (heights[nr][nc] < heights[r][c]) continue;
                visited[nr][nc] = true;
                q.push({nr, nc});
            }
        }
    }
};
```

---

### Java

```java
import java.util.*;

class Solution {
    public List<List<Integer>> pacificAtlantic(int[][] heights) {
        List<List<Integer>> res = new ArrayList<>();
        if (heights == null || heights.length == 0 || heights[0].length == 0) return res;
        int m = heights.length, n = heights[0].length;

        boolean[][] pac = new boolean[m][n];
        boolean[][] atl = new boolean[m][n];

        Queue<int[]> q = new LinkedList<>();
        for (int j = 0; j < n; ++j) { q.offer(new int[]{0, j}); pac[0][j] = true; }
        for (int i = 1; i < m; ++i) { q.offer(new int[]{i, 0}); pac[i][0] = true; }
        bfs(heights, q, pac);

        Queue<int[]> q2 = new LinkedList<>();
        for (int j = 0; j < n; ++j) { q2.offer(new int[]{m - 1, j}); atl[m - 1][j] = true; }
        for (int i = 0; i < m - 1; ++i) { q2.offer(new int[]{i, n - 1}); atl[i][n - 1] = true; }
        bfs(heights, q2, atl);

        for (int i = 0; i < m; ++i)
            for (int j = 0; j < n; ++j)
                if (pac[i][j] && atl[i][j]) res.add(Arrays.asList(i, j));
        return res;
    }

    private void bfs(int[][] heights, Queue<int[]> q, boolean[][] visited) {
        int m = heights.length, n = heights[0].length;
        int[][] dirs = {{1,0},{-1,0},{0,1},{0,-1}};
        while (!q.isEmpty()) {
            int[] cur = q.poll();
            int r = cur[0], c = cur[1];
            for (int[] d : dirs) {
                int nr = r + d[0], nc = c + d[1];
                if (nr < 0 || nr >= m || nc < 0 || nc >= n) continue;
                if (visited[nr][nc]) continue;
                if (heights[nr][nc] < heights[r][c]) continue;
                visited[nr][nc] = true;
                q.offer(new int[]{nr, nc});
            }
        }
    }
}
```

---

### JavaScript

```javascript
/**
 * @param {number[][]} heights
 * @return {number[][]}
 */
var pacificAtlantic = function(heights) {
    if (!heights || heights.length === 0) return [];
    const m = heights.length, n = heights[0].length;
    const pac = Array.from({length: m}, () => Array(n).fill(false));
    const atl = Array.from({length: m}, () => Array(n).fill(false));

    const bfs = (queue, visited) => {
        let head = 0;
        const dirs = [[1,0],[-1,0],[0,1],[0,-1]];
        while (head < queue.length) {
            const [r,c] = queue[head++];
            for (const [dr,dc] of dirs) {
                const nr = r + dr, nc = c + dc;
                if (nr < 0 || nr >= m || nc < 0 || nc >= n) continue;
                if (visited[nr][nc]) continue;
                if (heights[nr][nc] < heights[r][c]) continue;
                visited[nr][nc] = true;
                queue.push([nr,nc]);
            }
        }
    };

    const q1 = [];
    for (let j = 0; j < n; ++j) { pac[0][j] = true; q1.push([0,j]); }
    for (let i = 1; i < m; ++i) { pac[i][0] = true; q1.push([i,0]); }
    bfs(q1, pac);

    const q2 = [];
    for (let j = 0; j < n; ++j) { atl[m-1][j] = true; q2.push([m-1,j]); }
    for (let i = 0; i < m-1; ++i) { atl[i][n-1] = true; q2.push([i,n-1]); }
    bfs(q2, atl);

    const res = [];
    for (let i = 0; i < m; ++i)
        for (let j = 0; j < n; ++j)
            if (pac[i][j] && atl[i][j]) res.push([i,j]);
    return res;
};
```

---

### Python3

```python
from collections import deque
from typing import List

class Solution:
    def pacificAtlantic(self, heights: List[List[int]]) -> List[List[int]]:
        if not heights or not heights[0]:
            return []
        m, n = len(heights), len(heights[0])

        pac = [[False]*n for _ in range(m)]
        atl = [[False]*n for _ in range(m)]

        def bfs(starts, visited):
            q = deque(starts)
            while q:
                r, c = q.popleft()
                for dr, dc in ((1,0),(-1,0),(0,1),(0,-1)):
                    nr, nc = r + dr, c + dc
                    if nr < 0 or nr >= m or nc < 0 or nc >= n:
                        continue
                    if visited[nr][nc]:
                        continue
                    if heights[nr][nc] < heights[r][c]:
                        continue
                    visited[nr][nc] = True
                    q.append((nr, nc))

        pac_starts = [(0, j) for j in range(n)] + [(i, 0) for i in range(1, m)]
        for r, c in pac_starts: pac[r][c] = True
        bfs(pac_starts, pac)

        atl_starts = [(m - 1, j) for j in range(n)] + [(i, n - 1) for i in range(0, m - 1)]
        for r, c in atl_starts: atl[r][c] = True
        bfs(atl_starts, atl)

        ans = []
        for i in range(m):
            for j in range(n):
                if pac[i][j] and atl[i][j]:
                    ans.append([i, j])
        return ans
```

---

### Go

```go
package main

func pacificAtlantic(heights [][]int) [][]int {
    if len(heights) == 0 || len(heights[0]) == 0 {
        return [][]int{}
    }
    m, n := len(heights), len(heights[0])
    pac := make([][]bool, m)
    atl := make([][]bool, m)
    for i := 0; i < m; i++ {
        pac[i] = make([]bool, n)
        atl[i] = make([]bool, n)
    }

    bfs := func(starts [][2]int, visited [][]bool) {
        head := 0
        queue := make([][2]int, len(starts))
        copy(queue, starts)
        for head < len(queue) {
            cur := queue[head]
            head++
            r, c := cur[0], cur[1]
            dirs := [][2]int{{1,0},{-1,0},{0,1},{0,-1}}
            for _, d := range dirs {
                nr, nc := r + d[0], c + d[1]
                if nr < 0 || nr >= m || nc < 0 || nc >= n {
                    continue
                }
                if visited[nr][nc] {
                    continue
                }
                if heights[nr][nc] < heights[r][c] {
                    continue
                }
                visited[nr][nc] = true
                queue = append(queue, [2]int{nr, nc})
            }
        }
    }

    var pacStarts [][2]int
    for j := 0; j < n; j++ { pacStarts = append(pacStarts, [2]int{0, j}); pac[0][j] = true }
    for i := 1; i < m; i++ { pacStarts = append(pacStarts, [2]int{i, 0}); pac[i][0] = true }
    bfs(pacStarts, pac)

    var atlStarts [][2]int
    for j := 0; j < n; j++ { atlStarts = append(atlStarts, [2]int{m - 1, j}); atl[m-1][j] = true }
    for i := 0; i < m-1; i++ { atlStarts = append(atlStarts, [2]int{i, n - 1}); atl[i][n-1] = true }
    bfs(atlStarts, atl)

    var res [][]int
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if pac[i][j] && atl[i][j] {
                res = append(res, []int{i, j})
            }
        }
    }
    return res
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

Below I explain the core idea once (line-by-line for Python as the compact example) and then explain how the other languages map to the same steps.

### Python (line-by-line, teaching a friend)

```python
from collections import deque
from typing import List

class Solution:
    def pacificAtlantic(self, heights: List[List[int]]) -> List[List[int]]:
        if not heights or not heights[0]:
            return []
        m, n = len(heights), len(heights[0])
```

* Import `deque` for efficient BFS. If the input is empty return empty list. `m` and `n` are size dims.

```python
        pac = [[False]*n for _ in range(m)]
        atl = [[False]*n for _ in range(m)]
```

* `pac` and `atl` mark cells reachable from Pacific and Atlantic respectively.

```python
        def bfs(starts, visited):
            q = deque(starts)
            while q:
                r, c = q.popleft()
                for dr, dc in ((1,0),(-1,0),(0,1),(0,-1)):
                    nr, nc = r + dr, c + dc
                    if nr < 0 or nr >= m or nc < 0 or nc >= n:
                        continue
                    if visited[nr][nc]:
                        continue
                    if heights[nr][nc] < heights[r][c]:
                        continue
                    visited[nr][nc] = True
                    q.append((nr, nc))
```

* `bfs` expands from start cells. For each neighbor:

  * Skip if outside bounds.
  * Skip if already visited.
  * Only move to neighbor if `heights[nr][nc] >= heights[r][c]` (this respects the reversed idea of moving "uphill" from the ocean).
  * Mark visited and append to queue.

```python
        pac_starts = [(0, j) for j in range(n)] + [(i, 0) for i in range(1, m)]
        for r, c in pac_starts: pac[r][c] = True
        bfs(pac_starts, pac)
```

* Prepare Pacific starting border (top row and left column), mark them visited and BFS.

```python
        atl_starts = [(m - 1, j) for j in range(n)] + [(i, n - 1) for i in range(0, m - 1)]
        for r, c in atl_starts: atl[r][c] = True
        bfs(atl_starts, atl)
```

* Do the same for Atlantic (bottom row and right column).

```python
        ans = []
        for i in range(m):
            for j in range(n):
                if pac[i][j] and atl[i][j]:
                    ans.append([i, j])
        return ans
```

* Collect cells reachable by both BFS runs.

### How other languages map to the same flow

* **C++**: uses `queue<pair<int,int>>`, `vector<vector<bool>>` and the same BFS logic inside a `bfs` helper. The conditions and neighbor checks are identical; only syntax differs.
* **Java**: uses `Queue<int[]>`, `boolean[][]`, `LinkedList` for BFS. The rest is the same flow logic.
* **JavaScript**: uses arrays as queues (push + head index) or can use `shift()` but I used head pointer to keep it O(1). Visited arrays are 2D arrays of booleans.
* **Go**: uses slices of `[2]int` as queue and `[][]bool` as visited arrays. The BFS loop and checks are otherwise identical.
* All implementations: the important comparison is `heights[nr][nc] >= heights[r][c]` when expanding from the ocean outward.

---

## Examples

**Example 1**

```
Input:
heights = [
 [1,2,2,3,5],
 [3,2,3,4,4],
 [2,4,5,3,1],
 [6,7,1,4,5],
 [5,1,1,2,4]
]

Output (order may vary):
[[0,4],[1,3],[1,4],[2,2],[3,0],[3,1],[4,0]]
```

Explanation: cells listed can reach both Pacific (top/left) and Atlantic (bottom/right) by following non-increasing-height paths.

**Example 2**

```
Input:
heights = [[1]]
Output:
[[0,0]]
```

---

## How to use / Run locally

The code blocks above are LeetCode-style solutions (function/class that the platform calls). If you want to run locally, use the suggestions below.

### Python

Create `main.py` with the `Solution` class (paste the Python code) and add a test snippet:

```python
if __name__ == "__main__":
    heights = [[1,2,2,3,5],[3,2,3,4,4],[2,4,5,3,1],[6,7,1,4,5],[5,1,1,2,4]]
    print(Solution().pacificAtlantic(heights))
```

Run:

```bash
python3 main.py
```

### JavaScript (Node)

Create `pacific.js` with the `pacificAtlantic` function and add a test harness:

```javascript
// paste pacificAtlantic function here

const heights = [[1,2,2,3,5],[3,2,3,4,4],[2,4,5,3,1],[6,7,1,4,5],[5,1,1,2,4]];
console.log(pacificAtlantic(heights));
```

Run:

```bash
node pacific.js
```

### C++

Create `main.cpp` where you paste the `Solution` code and add a `main()` that constructs `heights` and calls `Solution().pacificAtlantic(heights)` and prints result. Compile & run:

```bash
g++ -std=c++17 main.cpp -O2 -o run
./run
```

### Java

Create `Solution.java` with the `Solution` class code. To run locally you must add a `main` method in another class or inside `Solution` that constructs a test grid and calls the method. Compile and run:

```bash
javac Solution.java
java Solution
```

### Go

Create `main.go`, paste the Go function and add a `main()` to call it and print the result. Run:

```bash
go run main.go
```

---

## Notes & Optimizations

* BFS and DFS both work. BFS is iterative and avoids recursion depth issues on large grids.
* We mark border cells visited before starting BFS to avoid duplicate queue entries.
* Space can be slightly reduced by using bit masks instead of two separate boolean grids, but clarity is reduced and the asymptotic complexity stays the same.
* If memory is tight, you can reuse one visited grid by running BFS for one ocean and then reusing that memory to compute the other, but keep a copy of the first result (or store intersection progressively).
* The key trick is **reverse expansion** (from ocean into higher neighbors). This allows a single pass per ocean and avoids path-explosion.

---

## Author

[Md. Aarzoo Islam](https://bento.me/withaarzoo)
