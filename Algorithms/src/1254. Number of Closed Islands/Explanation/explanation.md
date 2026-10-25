# Number of Closed Islands - LeetCode 1254 Solution

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

You are given a 2D grid that contains only 0s and 1s. Here 0 means land and 1 means water.

An island is any group of land cells (0s) that are connected to each other in four directions (up, down, left, right).

A closed island is a special kind of island that is completely surrounded by water on all four sides. In other words, none of its land cells touch the border of the grid.

Your task is to return the total number of closed islands present in the given grid.

This is a classic grid-based DFS / flood-fill problem that appears frequently in coding interviews and competitive programming contests under the name Number of Closed Islands (LeetCode 1254).

## Constraints

- 1 <= grid.length, grid[0].length <= 100
- 0 <= grid[i][j] <= 1

## Intuition

The key observation is simple: any land that touches the border can never be a closed island. So the first thing I do is remove every piece of land that is connected to the edge of the grid.

After I flood all the border-connected land and turn it into water, whatever land is still left must be completely trapped inside water. Counting those remaining connected groups gives me the answer directly.

This approach avoids any complicated checks for “is this island fully surrounded?” because the border-cleaning step already guarantees that the remaining islands are closed.

## Approach

1. Walk along all four borders of the grid.
2. Whenever I find a land cell (0) on the border, I run a DFS (or BFS) from that cell and turn the entire connected component into water (1).
3. After the borders are cleaned, I scan the inner part of the grid.
4. Every time I still see a 0, I know it belongs to a closed island. I increase the answer by one and flood that whole island so I do not count it again.
5. When the scan finishes, the answer holds the number of closed islands.

The same DFS function is reused for both the border-cleaning phase and the counting phase, which keeps the code short and clean.

## Data Structures Used

- 2D grid (given as input) – modified in place to mark visited land cells.
- Recursion stack – used by the DFS to explore connected components. No extra visited matrix is required because we change the grid itself.

## Operations & Behavior Summary

- Border flood phase: for every cell on the left, right, top and bottom edges, if it is land, start DFS and paint the whole connected component as water.
- Counting phase: iterate over every cell that is not on the border. When a land cell is found, increment the answer and flood that component so it is marked as visited.
- DFS helper: if the current cell is out of bounds or already water, return. Otherwise paint it water and recurse on its four neighbors.

This sequence guarantees that only fully enclosed islands are counted.

## Complexity

| Type              | Complexity | Explanation                                      |
|-------------------|------------|--------------------------------------------------|
| Time Complexity   | O(m × n)   | Every cell is visited a constant number of times |
| Space Complexity  | O(m × n)   | Worst-case depth of the recursion stack          |

Where m = number of rows and n = number of columns.

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int closedIsland(vector<vector<int>>& grid) {
        int m = grid.size(), n = grid[0].size();
        auto dfs = [&](auto&& self, int i, int j) -> void {
            if (i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 1) return;
            grid[i][j] = 1;
            self(self, i + 1, j);
            self(self, i - 1, j);
            self(self, i, j + 1);
            self(self, i, j - 1);
        };
        for (int i = 0; i < m; ++i) {
            if (grid[i][0] == 0) dfs(dfs, i, 0);
            if (grid[i][n - 1] == 0) dfs(dfs, i, n - 1);
        }
        for (int j = 0; j < n; ++j) {
            if (grid[0][j] == 0) dfs(dfs, 0, j);
            if (grid[m - 1][j] == 0) dfs(dfs, m - 1, j);
        }
        int ans = 0;
        for (int i = 1; i < m - 1; ++i) {
            for (int j = 1; j < n - 1; ++j) {
                if (grid[i][j] == 0) {
                    ++ans;
                    dfs(dfs, i, j);
                }
            }
        }
        return ans;
    }
};
```

### Java

```java
class Solution {
    public int closedIsland(int[][] grid) {
        int m = grid.length, n = grid[0].length;
        for (int i = 0; i < m; i++) {
            if (grid[i][0] == 0) dfs(grid, i, 0);
            if (grid[i][n - 1] == 0) dfs(grid, i, n - 1);
        }
        for (int j = 0; j < n; j++) {
            if (grid[0][j] == 0) dfs(grid, 0, j);
            if (grid[m - 1][j] == 0) dfs(grid, m - 1, j);
        }
        int ans = 0;
        for (int i = 1; i < m - 1; i++) {
            for (int j = 1; j < n - 1; j++) {
                if (grid[i][j] == 0) {
                    ans++;
                    dfs(grid, i, j);
                }
            }
        }
        return ans;
    }
    private void dfs(int[][] grid, int i, int j) {
        if (i < 0 || i >= grid.length || j < 0 || j >= grid[0].length || grid[i][j] == 1) return;
        grid[i][j] = 1;
        dfs(grid, i + 1, j);
        dfs(grid, i - 1, j);
        dfs(grid, i, j + 1);
        dfs(grid, i, j - 1);
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[][]} grid
 * @return {number}
 */
var closedIsland = function(grid) {
    const m = grid.length, n = grid[0].length;
    const dfs = (i, j) => {
        if (i < 0 || i >= m || j < 0 || j >= n || grid[i][j] === 1) return;
        grid[i][j] = 1;
        dfs(i + 1, j);
        dfs(i - 1, j);
        dfs(i, j + 1);
        dfs(i, j - 1);
    };
    for (let i = 0; i < m; i++) {
        if (grid[i][0] === 0) dfs(i, 0);
        if (grid[i][n - 1] === 0) dfs(i, n - 1);
    }
    for (let j = 0; j < n; j++) {
        if (grid[0][j] === 0) dfs(0, j);
        if (grid[m - 1][j] === 0) dfs(m - 1, j);
    }
    let ans = 0;
    for (let i = 1; i < m - 1; i++) {
        for (let j = 1; j < n - 1; j++) {
            if (grid[i][j] === 0) {
                ans++;
                dfs(i, j);
            }
        }
    }
    return ans;
};
```

### TypeScript

```typescript
function closedIsland(grid: number[][]): number {
    const m = grid.length, n = grid[0].length;
    const dfs = (i: number, j: number): void => {
        if (i < 0 || i >= m || j < 0 || j >= n || grid[i][j] === 1) return;
        grid[i][j] = 1;
        dfs(i + 1, j);
        dfs(i - 1, j);
        dfs(i, j + 1);
        dfs(i, j - 1);
    };
    for (let i = 0; i < m; i++) {
        if (grid[i][0] === 0) dfs(i, 0);
        if (grid[i][n - 1] === 0) dfs(i, n - 1);
    }
    for (let j = 0; j < n; j++) {
        if (grid[0][j] === 0) dfs(0, j);
        if (grid[m - 1][j] === 0) dfs(m - 1, j);
    }
    let ans = 0;
    for (let i = 1; i < m - 1; i++) {
        for (let j = 1; j < n - 1; j++) {
            if (grid[i][j] === 0) {
                ans++;
                dfs(i, j);
            }
        }
    }
    return ans;
};
```

### Python3

```python
class Solution:
    def closedIsland(self, grid: list[list[int]]) -> int:
        m, n = len(grid), len(grid[0])
        def dfs(i, j):
            if i < 0 or i >= m or j < 0 or j >= n or grid[i][j] == 1:
                return
            grid[i][j] = 1
            dfs(i + 1, j)
            dfs(i - 1, j)
            dfs(i, j + 1)
            dfs(i, j - 1)
        for i in range(m):
            if grid[i][0] == 0:
                dfs(i, 0)
            if grid[i][n - 1] == 0:
                dfs(i, n - 1)
        for j in range(n):
            if grid[0][j] == 0:
                dfs(0, j)
            if grid[m - 1][j] == 0:
                dfs(m - 1, j)
        ans = 0
        for i in range(1, m - 1):
            for j in range(1, n - 1):
                if grid[i][j] == 0:
                    ans += 1
                    dfs(i, j)
        return ans
```

### Go

```go
func closedIsland(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    var dfs func(i, j int)
    dfs = func(i, j int) {
        if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 1 {
            return
        }
        grid[i][j] = 1
        dfs(i+1, j)
        dfs(i-1, j)
        dfs(i, j+1)
        dfs(i, j-1)
    }
    for i := 0; i < m; i++ {
        if grid[i][0] == 0 {
            dfs(i, 0)
        }
        if grid[i][n-1] == 0 {
            dfs(i, n-1)
        }
    }
    for j := 0; j < n; j++ {
        if grid[0][j] == 0 {
            dfs(0, j)
        }
        if grid[m-1][j] == 0 {
            dfs(m-1, j)
        }
    }
    ans := 0
    for i := 1; i < m-1; i++ {
        for j := 1; j < n-1; j++ {
            if grid[i][j] == 0 {
                ans++
                dfs(i, j)
            }
        }
    }
    return ans
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

The logic is identical across all six languages, so the same reasoning applies everywhere.

First I store the number of rows and columns so I do not keep calling length functions inside the loops.

I define a DFS function that takes a cell coordinate. It immediately returns if the cell is outside the grid or if it is already water. Otherwise it marks the cell as water and calls itself on the four adjacent cells. This paints an entire connected island in one go.

Next I run the DFS on every border cell that still contains land. I loop over the leftmost and rightmost columns, then over the topmost and bottommost rows. After this pass, every island that touched the edge has been turned into water and will never be counted later.

Then I scan only the inner cells (rows 1 to m-2 and columns 1 to n-2). Whenever I find a remaining 0, I know it is the start of a closed island. I increase the answer and immediately flood that island so the same group is not counted again.

Edge cases such as a single-row or single-column grid are handled automatically because the inner-scan loops simply do not run, which is the correct behavior (no closed island can exist when there is no “inside”).

The same pattern works in C++, Java, JavaScript, TypeScript, Python and Go; only the syntax for defining the recursive function and accessing the grid changes.

## Examples

**Example 1**

Input:

```
[[1,1,1,1,1,1,1,0],
 [1,0,0,0,0,1,1,0],
 [1,0,1,0,1,1,1,0],
 [1,0,0,0,0,1,0,1],
 [1,1,1,1,1,1,1,0]]
```

Output: 2

Trace: The two small groups of 0s that sit completely inside the water are closed islands. The land that touches the right border is flooded away in the first phase and is not counted.

**Example 2**

Input:

```
[[0,0,1,0,0],
 [0,1,0,1,0],
 [0,1,1,1,0]]
```

Output: 1

Trace: The middle land cells form one closed island. Everything that touches the left or bottom border is removed first, leaving only the inner component.

**Example 3**

Input:

```
[[1,1,1,1,1,1,1],
 [1,0,0,0,0,0,1],
 [1,0,1,1,1,0,1],
 [1,0,1,0,1,0,1],
 [1,0,1,1,1,0,1],
 [1,0,0,0,0,0,1],
 [1,1,1,1,1,1,1]]
```

Output: 2

Trace: After the border flood, two separate inner islands remain and both are counted.

## How to Use / Run Locally

**C++**

```bash
g++ -std=c++17 solution.cpp -o solution
./solution
```

**Java**

```bash
javac Solution.java
java Solution
```

**JavaScript**

```bash
node solution.js
```

**TypeScript**

```bash
tsc solution.ts
node solution.js
```

**Python3**

```bash
python3 solution.py
```

**Go**

```bash
go run solution.go
```

Replace the empty function body with the code from the corresponding section above, then add a small main / driver that creates a test grid and prints the result.

## Notes & Optimizations

- The solution modifies the input grid in place. If the original grid must be preserved, make a copy first or use a separate visited matrix (extra O(m × n) space).
- For grids larger than 100 × 100 an iterative BFS or DFS with an explicit stack would be safer to avoid recursion-depth limits, but the given constraints make recursion perfectly fine.
- Union-Find can also solve the problem, but the DFS approach is shorter and easier to understand for most beginners.
- Single-cell grids, all-water grids, and grids that are only one row or one column all return 0, which is handled correctly by the border-first strategy.

## Author

[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)
