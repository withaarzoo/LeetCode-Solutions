# Twisted Mirror Path Count – LeetCode 3665 Solution

## Table of Contents

- [Problem Summary](#problem-summary)
- [Constraints](#constraints)
- [Intuition](#intuition)
- [Approach](#approach)
- [Data Structures Used](#data-structures-used)
- [Operations & Behavior Summary](#operations--behavior-summary)
- [Complexity](#complexity)
- [Multi-language Solutions](#multi-language-solutions)
  - [C++](#c)
  - [Java](#java)
  - [JavaScript](#javascript)
  - [TypeScript](#typescript)
  - [Python3](#python3)
  - [Go](#go)
- [Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)](#step-by-step-detailed-explanation-c-java-javascript-typescript-python3-go)
- [Examples](#examples)
- [How to Use / Run Locally](#how-to-use--run-locally)
- [Notes & Optimizations](#notes--optimizations)
- [Author](#author)

## Problem Summary

You are given an m by n binary grid. Each cell is either empty (0) or contains a mirror (1). A robot starts at the top-left cell (0, 0) and wants to reach the bottom-right cell (m-1, n-1). It can only try to move right or down.

When the robot tries to step onto a mirror cell, it never enters that cell. Instead it is reflected:

- Trying to move right into a mirror turns the robot downward and places it in the cell directly below the mirror.
- Trying to move down into a mirror turns the robot rightward and places it in the cell directly to the right of the mirror.

If a reflection would push the robot outside the grid, that path is invalid. Reflections can chain: if the new cell is also a mirror, the robot is reflected again using the new direction. This continues until the robot lands on an empty cell or goes out of bounds.

Return the number of unique valid paths from start to end. Because the answer can be large, return it modulo 10^9 + 7.

The start and end cells are always empty.

## Constraints

- 2 <= m, n <= 500
- grid[i][j] is either 0 or 1
- grid[0][0] == grid[m-1][n-1] == 0

## Intuition

The key observation is that every reflection always increases either the row index, the column index, or both. The robot can never move left or up. That means the movement graph is a DAG (directed acyclic graph).  

Because the graph has no cycles, I can safely pre-compute the final landing cell for every possible move from every cell. Once those “jump targets” are known, counting paths becomes a standard dynamic-programming problem on a grid: just add the number of ways while scanning cells in topological order.

## Approach

1. Create two tables that store, for every cell and every direction (right or down), the final row and column the robot will land on after following the entire chain of mirrors.
2. Fill these tables from the bottom-right corner toward the top-left. A mirror always redirects to a later cell, so the answer for the next cell is already known.
3. If a cell is empty, the robot simply stops there. If it is a mirror, look up the already-computed landing of the single cell dictated by the reflection rule. Out-of-bounds is marked invalid.
4. After the jump tables are ready, run ordinary path-counting DP. Let dp[i][j] be the number of ways to reach the empty cell (i, j).
5. Set dp[0][0] = 1. Scan the grid in row-major order. From every cell that already has ways, read the two pre-computed landing positions and add the current ways to them (modulo 10^9 + 7).
6. Because every jump lands on a later cell in the same order, all contributions are present before a cell is processed. The value left in dp[m-1][n-1] is the answer.

## Data Structures Used

- Three-dimensional arrays (or equivalent nested vectors/lists) to store the final landing row and landing column for every cell and every direction. These tables let us answer any jump query in constant time after a linear pre-computation.
- A two-dimensional DP table that stores the number of ways to reach each empty cell. Only empty cells ever receive values because the robot never lands on mirrors.

## Operations & Behavior Summary

- Pre-computation phase (reverse order): for each cell and each direction decide whether the robot stops, redirects, or goes out of bounds; store the result.
- DP phase (forward order): for each cell that has a positive number of ways, look up the right-jump landing and the down-jump landing; if the landing is valid, add the current ways to that landing cell.
- Modular arithmetic is applied on every addition to keep numbers inside 32-bit range.
- The algorithm never simulates individual paths; it only moves aggregated way-counts along the pre-computed jumps.

## Complexity

| Type              | Complexity | Explanation |
|-------------------|------------|-------------|
| Time Complexity   | O(m · n)   | Both the reverse-order pre-computation and the forward DP visit each of the m · n cells a constant number of times. |
| Space Complexity  | O(m · n)   | Two landing tables and one DP table, each of size proportional to the grid. |

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int uniquePaths(vector<vector<int>>& grid) {
        int m = grid.size(), n = grid[0].size();
        const int MOD = 1e9 + 7;
        vector<vector<array<int,2>>> landR(m, vector<array<int,2>>(n));
        vector<vector<array<int,2>>> landC(m, vector<array<int,2>>(n));
        for (int x = m - 1; x >= 0; --x) {
            for (int y = n - 1; y >= 0; --y) {
                for (int d = 0; d < 2; ++d) {
                    if (grid[x][y] == 0) {
                        landR[x][y][d] = x;
                        landC[x][y][d] = y;
                    } else {
                        int nx = (d == 0) ? x + 1 : x;
                        int ny = (d == 0) ? y : y + 1;
                        int nd = 1 - d;
                        if (nx < 0 || nx >= m || ny < 0 || ny >= n) {
                            landR[x][y][d] = -1;
                            landC[x][y][d] = -1;
                        } else {
                            landR[x][y][d] = landR[nx][ny][nd];
                            landC[x][y][d] = landC[nx][ny][nd];
                        }
                    }
                }
            }
        }
        vector<vector<int>> dp(m, vector<int>(n, 0));
        dp[0][0] = 1;
        for (int i = 0; i < m; ++i) {
            for (int j = 0; j < n; ++j) {
                if (dp[i][j] == 0) continue;
                if (j + 1 < n) {
                    int nr = landR[i][j + 1][0];
                    int nc = landC[i][j + 1][0];
                    if (nr != -1) dp[nr][nc] = (dp[nr][nc] + dp[i][j]) % MOD;
                }
                if (i + 1 < m) {
                    int nr = landR[i + 1][j][1];
                    int nc = landC[i + 1][j][1];
                    if (nr != -1) dp[nr][nc] = (dp[nr][nc] + dp[i][j]) % MOD;
                }
            }
        }
        return dp[m - 1][n - 1];
    }
};
```

### Java

```java
class Solution {
    public int uniquePaths(int[][] grid) {
        int m = grid.length, n = grid[0].length;
        final int MOD = 1000000007;
        int[][][] landR = new int[m][n][2];
        int[][][] landC = new int[m][n][2];
        for (int x = m - 1; x >= 0; --x) {
            for (int y = n - 1; y >= 0; --y) {
                for (int d = 0; d < 2; ++d) {
                    if (grid[x][y] == 0) {
                        landR[x][y][d] = x;
                        landC[x][y][d] = y;
                    } else {
                        int nx = (d == 0) ? x + 1 : x;
                        int ny = (d == 0) ? y : y + 1;
                        int nd = 1 - d;
                        if (nx < 0 || nx >= m || ny < 0 || ny >= n) {
                            landR[x][y][d] = -1;
                            landC[x][y][d] = -1;
                        } else {
                            landR[x][y][d] = landR[nx][ny][nd];
                            landC[x][y][d] = landC[nx][ny][nd];
                        }
                    }
                }
            }
        }
        int[][] dp = new int[m][n];
        dp[0][0] = 1;
        for (int i = 0; i < m; ++i) {
            for (int j = 0; j < n; ++j) {
                if (dp[i][j] == 0) continue;
                if (j + 1 < n) {
                    int nr = landR[i][j + 1][0];
                    int nc = landC[i][j + 1][0];
                    if (nr != -1) dp[nr][nc] = (dp[nr][nc] + dp[i][j]) % MOD;
                }
                if (i + 1 < m) {
                    int nr = landR[i + 1][j][1];
                    int nc = landC[i + 1][j][1];
                    if (nr != -1) dp[nr][nc] = (dp[nr][nc] + dp[i][j]) % MOD;
                }
            }
        }
        return dp[m - 1][n - 1];
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[][]} grid
 * @return {number}
 */
var uniquePaths = function(grid) {
    const m = grid.length, n = grid[0].length;
    const MOD = 1000000007;
    const landR = Array.from({length: m}, () => Array.from({length: n}, () => Array(2)));
    const landC = Array.from({length: m}, () => Array.from({length: n}, () => Array(2)));
    for (let x = m - 1; x >= 0; --x) {
        for (let y = n - 1; y >= 0; --y) {
            for (let d = 0; d < 2; ++d) {
                if (grid[x][y] === 0) {
                    landR[x][y][d] = x;
                    landC[x][y][d] = y;
                } else {
                    const nx = d === 0 ? x + 1 : x;
                    const ny = d === 0 ? y : y + 1;
                    const nd = 1 - d;
                    if (nx < 0 || nx >= m || ny < 0 || ny >= n) {
                        landR[x][y][d] = -1;
                        landC[x][y][d] = -1;
                    } else {
                        landR[x][y][d] = landR[nx][ny][nd];
                        landC[x][y][d] = landC[nx][ny][nd];
                    }
                }
            }
        }
    }
    const dp = Array.from({length: m}, () => Array(n).fill(0));
    dp[0][0] = 1;
    for (let i = 0; i < m; ++i) {
        for (let j = 0; j < n; ++j) {
            if (dp[i][j] === 0) continue;
            if (j + 1 < n) {
                const nr = landR[i][j + 1][0];
                const nc = landC[i][j + 1][0];
                if (nr !== -1) dp[nr][nc] = (dp[nr][nc] + dp[i][j]) % MOD;
            }
            if (i + 1 < m) {
                const nr = landR[i + 1][j][1];
                const nc = landC[i + 1][j][1];
                if (nr !== -1) dp[nr][nc] = (dp[nr][nc] + dp[i][j]) % MOD;
            }
        }
    }
    return dp[m - 1][n - 1];
};
```

### TypeScript

```typescript
function uniquePaths(grid: number[][]): number {
    const m = grid.length, n = grid[0].length;
    const MOD = 1000000007;
    const landR: number[][][] = Array.from({length: m}, () => Array.from({length: n}, () => Array(2)));
    const landC: number[][][] = Array.from({length: m}, () => Array.from({length: n}, () => Array(2)));
    for (let x = m - 1; x >= 0; --x) {
        for (let y = n - 1; y >= 0; --y) {
            for (let d = 0; d < 2; ++d) {
                if (grid[x][y] === 0) {
                    landR[x][y][d] = x;
                    landC[x][y][d] = y;
                } else {
                    const nx = d === 0 ? x + 1 : x;
                    const ny = d === 0 ? y : y + 1;
                    const nd = 1 - d;
                    if (nx < 0 || nx >= m || ny < 0 || ny >= n) {
                        landR[x][y][d] = -1;
                        landC[x][y][d] = -1;
                    } else {
                        landR[x][y][d] = landR[nx][ny][nd];
                        landC[x][y][d] = landC[nx][ny][nd];
                    }
                }
            }
        }
    }
    const dp: number[][] = Array.from({length: m}, () => Array(n).fill(0));
    dp[0][0] = 1;
    for (let i = 0; i < m; ++i) {
        for (let j = 0; j < n; ++j) {
            if (dp[i][j] === 0) continue;
            if (j + 1 < n) {
                const nr = landR[i][j + 1][0];
                const nc = landC[i][j + 1][0];
                if (nr !== -1) dp[nr][nc] = (dp[nr][nc] + dp[i][j]) % MOD;
            }
            if (i + 1 < m) {
                const nr = landR[i + 1][j][1];
                const nc = landC[i + 1][j][1];
                if (nr !== -1) dp[nr][nc] = (dp[nr][nc] + dp[i][j]) % MOD;
            }
        }
    }
    return dp[m - 1][n - 1];
}
```

### Python3

```python
from typing import List
class Solution:
    def uniquePaths(self, grid: List[List[int]]) -> int:
        m, n = len(grid), len(grid[0])
        MOD = 10**9 + 7
        landR = [[[0, 0] for _ in range(n)] for _ in range(m)]
        landC = [[[0, 0] for _ in range(n)] for _ in range(m)]
        for x in range(m - 1, -1, -1):
            for y in range(n - 1, -1, -1):
                for d in range(2):
                    if grid[x][y] == 0:
                        landR[x][y][d] = x
                        landC[x][y][d] = y
                    else:
                        nx = x + 1 if d == 0 else x
                        ny = y if d == 0 else y + 1
                        nd = 1 - d
                        if nx < 0 or nx >= m or ny < 0 or ny >= n:
                            landR[x][y][d] = -1
                            landC[x][y][d] = -1
                        else:
                            landR[x][y][d] = landR[nx][ny][nd]
                            landC[x][y][d] = landC[nx][ny][nd]
        dp = [[0] * n for _ in range(m)]
        dp[0][0] = 1
        for i in range(m):
            for j in range(n):
                if dp[i][j] == 0:
                    continue
                if j + 1 < n:
                    nr = landR[i][j + 1][0]
                    nc = landC[i][j + 1][0]
                    if nr != -1:
                        dp[nr][nc] = (dp[nr][nc] + dp[i][j]) % MOD
                if i + 1 < m:
                    nr = landR[i + 1][j][1]
                    nc = landC[i + 1][j][1]
                    if nr != -1:
                        dp[nr][nc] = (dp[nr][nc] + dp[i][j]) % MOD
        return dp[m - 1][n - 1]
```

### Go

```go
func uniquePaths(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    const MOD = 1000000007
    landR := make([][][2]int, m)
    landC := make([][][2]int, m)
    for i := range landR {
        landR[i] = make([][2]int, n)
        landC[i] = make([][2]int, n)
    }
    for x := m - 1; x >= 0; x-- {
        for y := n - 1; y >= 0; y-- {
            for d := 0; d < 2; d++ {
                if grid[x][y] == 0 {
                    landR[x][y][d] = x
                    landC[x][y][d] = y
                } else {
                    nx, ny, nd := x, y, 1-d
                    if d == 0 {
                        nx = x + 1
                    } else {
                        ny = y + 1
                    }
                    if nx < 0 || nx >= m || ny < 0 || ny >= n {
                        landR[x][y][d] = -1
                        landC[x][y][d] = -1
                    } else {
                        landR[x][y][d] = landR[nx][ny][nd]
                        landC[x][y][d] = landC[nx][ny][nd]
                    }
                }
            }
        }
    }
    dp := make([][]int, m)
    for i := range dp {
        dp[i] = make([]int, n)
    }
    dp[0][0] = 1
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if dp[i][j] == 0 {
                continue
            }
            if j+1 < n {
                nr := landR[i][j+1][0]
                nc := landC[i][j+1][0]
                if nr != -1 {
                    dp[nr][nc] = (dp[nr][nc] + dp[i][j]) % MOD
                }
            }
            if i+1 < m {
                nr := landR[i+1][j][1]
                nc := landC[i+1][j][1]
                if nr != -1 {
                    dp[nr][nc] = (dp[nr][nc] + dp[i][j]) % MOD
                }
            }
        }
    }
    return dp[m-1][n-1]
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

The logic is identical in every language; only syntax differs.

First the landing tables are filled from the bottom-right corner upward and leftward. For an empty cell the landing is the cell itself. For a mirror cell the algorithm computes the single next cell required by the reflection rule and copies the already-known landing of that next cell. Any out-of-bounds next cell is stored as an invalid marker.

After the tables are ready the DP table is created and the start cell is given one way. The grid is then scanned from top-left to bottom-right. Whenever a cell already contains a positive number of ways, the two pre-computed landings (right move and down move) are read. If a landing is valid the current ways are added to it under modular arithmetic.

Because every jump always lands on a cell that appears later in the same scan order, every contribution is already present when a cell is examined. When the scan finishes, the bottom-right entry of the DP table holds the final answer.

Edge cases such as immediate out-of-bounds moves, long chains of consecutive mirrors, and grids that contain no mirrors at all are all handled uniformly by the same landing tables.

## Examples

**Example 1**  
Input: grid = [[0,1,0],[0,0,1],[1,0,0]]  
Output: 5  

The five valid paths are the ones shown in the problem statement. Each time the robot would step onto a mirror it is redirected according to the rules and eventually reaches (2,2).

**Example 2**  
Input: grid = [[0,0],[0,0]]  
Output: 2  

No mirrors exist. The two ordinary paths (right-then-down and down-then-right) are both valid.

**Example 3**  
Input: grid = [[0,1,1],[1,1,0]]  
Output: 1  

Only one sequence of reflections stays inside the grid and lands on the empty end cell. The other possible first move eventually walks off the board.

## How to Use / Run Locally

1. Copy the code for the language you want into a file named according to the language (Solution.cpp, Solution.java, solution.js, etc.).
2. Make sure a recent compiler or runtime is installed:
   - C++: g++ with C++17 or later
   - Java: JDK 11 or later
   - JavaScript / TypeScript: Node.js 16 or later (for TypeScript also install ts-node or compile with tsc)
   - Python3: Python 3.8 or later
   - Go: Go 1.18 or later
3. Compile or run the file. Most online judges accept the class or function exactly as written; for local testing you can wrap a small main that builds a sample grid and prints the result.

## Notes & Optimizations

- The reverse-order pre-computation eliminates recursion and any risk of stack overflow on long mirror chains.
- Because m and n are at most 500, the O(m · n) solution runs comfortably under typical time limits and usually finishes in well under a millisecond.
- An alternative recursive memoized jump function also works, but the iterative version is simpler and slightly faster in practice.
- The algorithm never needs to store the actual paths, only the number of ways, which keeps memory usage linear in the size of the grid.

## Author

[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)
