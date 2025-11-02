# Count Unguarded Cells in the Grid (LeetCode 2257) — README

## Problem Title

**2257. Count Unguarded Cells in the Grid**

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

I am given two integers `m` and `n` for an `m x n` grid. I also have two lists: `guards` and `walls`. Each `guards[i] = [row_i, col_i]` and `walls[j] = [row_j, col_j]`.
A guard can see every cell in the four cardinal directions (north, south, west, east) until blocked by a wall or another guard. A cell is *guarded* if at least one guard can see it. I need to return the number of **unoccupied** cells that are **not guarded**.

Unoccupied means the cell is not a guard and not a wall.

---

## Constraints

* `1 <= m, n <= 10^3` (but product `m*n` must be reasonable for memory usage in typical settings)
* `0 <= guards.length, walls.length <= m*n`
* Guard and wall coordinates are valid and distinct from each other (a cell won't be both guard and wall).

---

## Intuition

I thought of this as simulating each guard's line of sight. If I mark walls and guards on a grid first, then for every guard I can march outward in four directions and mark empty cells as guarded until I hit a wall or another guard. After processing all guards, any cell still empty is unguarded. This is simple and direct, and it avoids unnecessary BFS or complicated data structures.

---

## Approach

1. Create a 2D grid `grid[m][n]` to encode states:

   * `0` = empty/unoccupied
   * `1` = guard
   * `2` = wall
   * `3` = guarded (seen by at least one guard)
2. Place all walls and guards in the grid.
3. For each guard, expand in four directions (up, down, left, right):

   * Move step by step.
   * Stop if the next cell is out of bounds, a wall, or another guard.
   * If the next cell is empty (`0`) mark it as guarded (`3`) and continue.
4. After scanning all guards, count cells that remain `0`. That's the answer.

---

## Data Structures Used

* **2D array** (grid of integers) to store cell states: guards, walls, guarded, empty.
* Input arrays `guards` and `walls` are used to initialize the grid.

---

## Operations & Behavior Summary

* **Initialization:** O(m*n) to create the grid (or O(1) if using sparse structures, but here we use dense grid).
* **Placement:** O(G + W) where G = number of guards, W = number of walls.
* **Line-of-sight marking:** For each guard, we move in at most `max(m, n)` steps per direction. So work ≈ O(G * max(m, n)) in the worst case.
* **Counting:** O(m*n) to count remaining empty cells.

---

## Complexity

* **Time Complexity:** `O(m*n + G * max(m, n))`

  * `m*n` for grid setup and final count.
  * `G` is number of guards. For each guard we scan up to `max(m, n)` cells per direction.
* **Space Complexity:** `O(m*n)` for the grid.

---

## Multi-language Solutions

> The implementations all follow the same approach: encode a grid, place guards/walls, propagate guard sight in 4 directions, then count remaining empty cells.

### C++

```c++
#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    int countUnguarded(int m, int n, vector<vector<int>>& guards, vector<vector<int>>& walls) {
        // 0 = empty, 1 = guard, 2 = wall, 3 = guarded
        vector<vector<int>> grid(m, vector<int>(n, 0));
        for (auto &w : walls) grid[w[0]][w[1]] = 2;
        for (auto &g : guards) grid[g[0]][g[1]] = 1;

        const int dirs[4][2] = {{-1,0},{1,0},{0,-1},{0,1}};
        for (auto &g : guards) {
            int r = g[0], c = g[1];
            for (int d = 0; d < 4; ++d) {
                int nr = r + dirs[d][0], nc = c + dirs[d][1];
                while (nr >= 0 && nr < m && nc >= 0 && nc < n) {
                    if (grid[nr][nc] == 2 || grid[nr][nc] == 1) break;
                    if (grid[nr][nc] == 0) grid[nr][nc] = 3;
                    nr += dirs[d][0];
                    nc += dirs[d][1];
                }
            }
        }

        int ans = 0;
        for (int i = 0; i < m; ++i)
            for (int j = 0; j < n; ++j)
                if (grid[i][j] == 0) ++ans;
        return ans;
    }
};
```

---

### Java

```java
import java.util.*;

class Solution {
    public int countUnguarded(int m, int n, int[][] guards, int[][] walls) {
        // 0 = empty, 1 = guard, 2 = wall, 3 = guarded
        int[][] grid = new int[m][n];
        for (int[] w : walls) grid[w[0]][w[1]] = 2;
        for (int[] g : guards) grid[g[0]][g[1]] = 1;

        int[][] dirs = {{-1,0},{1,0},{0,-1},{0,1}};
        for (int[] g : guards) {
            int r = g[0], c = g[1];
            for (int[] dir : dirs) {
                int nr = r + dir[0], nc = c + dir[1];
                while (nr >= 0 && nr < m && nc >= 0 && nc < n) {
                    if (grid[nr][nc] == 2 || grid[nr][nc] == 1) break;
                    if (grid[nr][nc] == 0) grid[nr][nc] = 3;
                    nr += dir[0];
                    nc += dir[1];
                }
            }
        }

        int ans = 0;
        for (int i = 0; i < m; ++i)
            for (int j = 0; j < n; ++j)
                if (grid[i][j] == 0) ans++;
        return ans;
    }
}
```

---

### JavaScript

```javascript
/**
 * @param {number} m
 * @param {number} n
 * @param {number[][]} guards
 * @param {number[][]} walls
 * @return {number}
 */
var countUnguarded = function(m, n, guards, walls) {
    // 0 = empty, 1 = guard, 2 = wall, 3 = guarded
    const grid = Array.from({length: m}, () => Array(n).fill(0));
    for (const w of walls) grid[w[0]][w[1]] = 2;
    for (const g of guards) grid[g[0]][g[1]] = 1;

    const dirs = [[-1,0],[1,0],[0,-1],[0,1]];
    for (const g of guards) {
        const r = g[0], c = g[1];
        for (const d of dirs) {
            let nr = r + d[0], nc = c + d[1];
            while (nr >= 0 && nr < m && nc >= 0 && nc < n) {
                if (grid[nr][nc] === 2 || grid[nr][nc] === 1) break;
                if (grid[nr][nc] === 0) grid[nr][nc] = 3;
                nr += d[0];
                nc += d[1];
            }
        }
    }

    let ans = 0;
    for (let i = 0; i < m; ++i)
        for (let j = 0; j < n; ++j)
            if (grid[i][j] === 0) ans++;
    return ans;
};
```

---

### Python3

```python
class Solution:
    def countUnguarded(self, m: int, n: int, guards: List[List[int]], walls: List[List[int]]) -> int:
        # 0 = empty, 1 = guard, 2 = wall, 3 = guarded
        grid = [[0]*n for _ in range(m)]
        for r, c in walls:
            grid[r][c] = 2
        for r, c in guards:
            grid[r][c] = 1

        dirs = [(-1,0),(1,0),(0,-1),(0,1)]
        for r, c in guards:
            for dr, dc in dirs:
                nr, nc = r + dr, c + dc
                while 0 <= nr < m and 0 <= nc < n:
                    if grid[nr][nc] == 2 or grid[nr][nc] == 1:
                        break
                    if grid[nr][nc] == 0:
                        grid[nr][nc] = 3
                    nr += dr
                    nc += dc

        ans = 0
        for i in range(m):
            for j in range(n):
                if grid[i][j] == 0:
                    ans += 1
        return ans
```

---

### Go

```go
package main

func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
    // 0 = empty, 1 = guard, 2 = wall, 3 = guarded
    grid := make([][]int, m)
    for i := 0; i < m; i++ {
        grid[i] = make([]int, n)
    }
    for _, w := range walls {
        grid[w[0]][w[1]] = 2
    }
    for _, g := range guards {
        grid[g[0]][g[1]] = 1
    }

    dirs := [4][2]int{{-1,0},{1,0},{0,-1},{0,1}}
    for _, g := range guards {
        r, c := g[0], g[1]
        for _, d := range dirs {
            nr, nc := r + d[0], c + d[1]
            for nr >= 0 && nr < m && nc >= 0 && nc < n {
                if grid[nr][nc] == 2 || grid[nr][nc] == 1 {
                    break
                }
                if grid[nr][nc] == 0 {
                    grid[nr][nc] = 3
                }
                nr += d[0]
                nc += d[1]
            }
        }
    }

    ans := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 0 {
                ans++
            }
        }
    }
    return ans
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

I'll explain the important logic once and then call out differences per language. I explain as if I'm teaching a friend.

### 1) High-level flow (common)

* I create `grid[m][n]` and mark walls (`2`) and guards (`1`).
* I define 4 directions: up, down, left, right.
* For each guard position `(r,c)`:

  * For each direction `(dr,dc)`:

    * Start at `(r+dr, c+dc)`.
    * While inside bounds:

      * If the cell is a wall (`2`) or a guard (`1`), **stop** that direction.
      * If the cell is empty (`0`), set it to guarded (`3`) and continue.
* After all guards processed, count cells that are still `0`.

### 2) C++ notes

* `vector<vector<int>> grid(m, vector<int>(n, 0));` initializes grid fast.
* `dirs` is a small fixed array of direction offsets.
* We check `grid[nr][nc] == 2 || grid[nr][nc] == 1` to stop scanning.
* Counting is simple nested loops.

### 3) Java notes

* Use `int[][] grid = new int[m][n];` to create the grid.
* Arrays of directions used as `int[][] dirs`.
* Java loops and bounds checks mirror C++ logic exactly.

### 4) JavaScript notes

* `Array.from({length: m}, () => Array(n).fill(0))` creates the 2D array.
* Use `const` for constants and `let` / `var` for loop variables.
* Comparison `===` used.

### 5) Python3 notes

* `grid = [[0]*n for _ in range(m)]` creates grid.
* Use tuples `dirs = [(-1,0),(1,0),(0,-1),(0,1)]`.
* Python list iteration makes code compact.

### 6) Go notes

* `grid := make([][]int, m)` and then inner `make([]int, n)` per row.
* `dirs` is an array of arrays `[4][2]int`.
* For performance, avoid unnecessary allocations inside loops.

---

## Examples

Example 1:

```
Input:
m = 4, n = 6
guards = [[0,0],[1,1],[2,3]]
walls = [[0,1],[2,2],[1,4]]

Output: 7
Explanation: The example in the problem statement shows guarded (red) and unguarded (green). There are 7 unguarded cells.
```

Example 2 (small):

```
m = 1, n = 3
guards = [[0,1]]
walls = []
Result:
- Guard at (0,1) guards (0,0) and (0,2).
Return 0 unguarded cells.
```

---

## How to use / Run locally

### C++

* Create a `main()` to instantiate `Solution` and call `countUnguarded`.
* Compile with `g++ -std=c++17 solution.cpp -O2` and run.

### Java

* Put the `Solution` class in `Solution.java` and provide a `main` to test.
* Compile and run:

  ```
  javac Solution.java
  java Solution
  ```

### JavaScript

* Save function into `solution.js` and call it from Node with test inputs:

  ```
  node
  > const res = countUnguarded(m, n, guards, walls);
  ```

### Python3

* Save as `solution.py`, add a test block:

  ```python
  if __name__ == "__main__":
      s = Solution()
      print(s.countUnguarded(m, n, guards, walls))
  ```

* Run with `python3 solution.py`.

### Go

* Save into `main.go` with a `main()` wrapper calling `countUnguarded`.
* Run:

  ```
  go run main.go
  ```

---

## Notes & Optimizations

* **Memory trade-off:** I use a dense `m*n` grid. If `m*n` is too large for memory, a sparse approach (using hash maps / sets for walls & guards and scanning until hitting a wall/guard spot checked via set membership) may be used, but scanning may require repeated membership checks.
* **Early marking:** I mark guarded cells as `3` and continue scanning through them — already-guarded cells do not block sight.
* **Stopping at other guards:** The problem explicitly states that another guard blocks sight, so I stop when encountering a guard.
* **Alternative approach:** We could do four sweeps across rows and columns marking guarded cells in O(m*n) total by sorting/organizing guards and walls per row/column; that can be more optimal if G is large and grid is huge (it can reduce factor of G * max(m,n) to O(m*n)), but it's more complex to implement. The current approach is straightforward and works within typical LeetCode limits.

---

## Author

* [Md. Aarzoo Islam](https://bento.me/withaarzoo)
