# 1559. Detect Cycles in 2D Grid

## Table of Contents

* [Problem Summary](#problem-summary)
* [Constraints](#constraints)
* [Intuition](#intuition)
* [Approach](#approach)
* [Data Structures Used](#data-structures-used)
* [Operations & Behavior Summary](#operations--behavior-summary)
* [Complexity](#complexity)
* [Multi-language Solutions](#multi-language-solutions)

  * [C++](#c++)
  * [Java](#java)
  * [JavaScript](#javascript)
  * [Python3](#python3)
  * [Go](#go)
* [Step-by-step Detailed Explanation](#step-by-step-detailed-explanation-c-java-javascript-python3-go)
* [Examples](#examples)
* [How to use / Run locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

## Problem Summary

I am given a 2D grid of lowercase English letters. I need to determine whether there exists a cycle made of the same character.

A cycle is a path of length 4 or more that:

* starts and ends at the same cell,
* moves only in 4 directions: up, down, left, right,
* visits only cells with the same character,
* does not immediately go back to the cell visited in the previous move.

The task is to return `true` if at least one valid cycle exists. Otherwise, return `false`.

## Constraints

* `m == grid.length`
* `n == grid[i].length`
* `1 <= m, n <= 500`
* `grid` contains only lowercase English letters

## Intuition

I thought of this problem as a graph problem.

Every cell is a node.
If two adjacent cells have the same character, then they are connected.
So the grid becomes an undirected graph.

Now the problem becomes:
"Does this graph contain a cycle?"

For an undirected graph, DFS is a very natural way to detect cycles.
While doing DFS, I keep track of the parent cell.
If I ever reach a visited cell that is not the parent, then I know a cycle exists.

## Approach

1. I create a `visited` matrix of the same size as the grid.
2. I scan every cell in the grid.
3. If a cell is not visited, I start DFS from that cell.
4. During DFS, I only move to neighboring cells that:

   * are inside the grid,
   * have the same character,
   * are not the parent cell.
5. If I find a neighbor that is already visited and is not the parent, then a cycle exists.
6. If I finish checking all cells without finding such a case, I return `false`.

I use iterative DFS in the code to avoid recursion depth problems for large grids.

## Data Structures Used

* **2D boolean array / matrix**: to mark visited cells
* **Stack**: to perform iterative DFS
* **Direction arrays**: to move in four directions

## Operations & Behavior Summary

* I visit each cell at most once.
* From each cell, I check at most 4 neighbors.
* I skip the direct parent to avoid false cycle detection.
* If I see a previously visited same-character cell that is not the parent, I return `true`.

## Complexity

* **Time Complexity:** `O(m * n)`

  * Every cell is processed once.
  * For each cell, I check up to 4 neighbors.

* **Space Complexity:** `O(m * n)`

  * For the `visited` matrix.
  * In the worst case, the DFS stack can also grow to `O(m * n)`.

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    bool containsCycle(vector<vector<char>>& grid) {
        int m = grid.size();
        int n = grid[0].size();
        vector<vector<bool>> visited(m, vector<bool>(n, false));

        int dr[4] = {1, -1, 0, 0};
        int dc[4] = {0, 0, 1, -1};

        for (int r = 0; r < m; ++r) {
            for (int c = 0; c < n; ++c) {
                if (visited[r][c]) continue;

                // stack node: {row, col, parentRow, parentCol}
                vector<array<int, 4>> st;
                st.push_back({r, c, -1, -1});
                visited[r][c] = true;

                while (!st.empty()) {
                    auto cur = st.back();
                    st.pop_back();

                    int cr = cur[0], cc = cur[1];
                    int pr = cur[2], pc = cur[3];

                    for (int k = 0; k < 4; ++k) {
                        int nr = cr + dr[k];
                        int nc = cc + dc[k];

                        if (nr < 0 || nr >= m || nc < 0 || nc >= n) continue;
                        if (grid[nr][nc] != grid[cr][cc]) continue;
                        if (nr == pr && nc == pc) continue;

                        if (visited[nr][nc]) return true;

                        visited[nr][nc] = true;
                        st.push_back({nr, nc, cr, cc});
                    }
                }
            }
        }

        return false;
    }
};
```

### Java

```java
class Solution {
    public boolean containsCycle(char[][] grid) {
        int m = grid.length;
        int n = grid[0].length;
        boolean[][] visited = new boolean[m][n];

        int[] dr = {1, -1, 0, 0};
        int[] dc = {0, 0, 1, -1};

        for (int r = 0; r < m; r++) {
            for (int c = 0; c < n; c++) {
                if (visited[r][c]) continue;

                ArrayDeque<int[]> stack = new ArrayDeque<>();
                stack.addLast(new int[]{r, c, -1, -1});
                visited[r][c] = true;

                while (!stack.isEmpty()) {
                    int[] cur = stack.removeLast();
                    int cr = cur[0], cc = cur[1], pr = cur[2], pc = cur[3];

                    for (int k = 0; k < 4; k++) {
                        int nr = cr + dr[k];
                        int nc = cc + dc[k];

                        if (nr < 0 || nr >= m || nc < 0 || nc >= n) continue;
                        if (grid[nr][nc] != grid[cr][cc]) continue;
                        if (nr == pr && nc == pc) continue;

                        if (visited[nr][nc]) return true;

                        visited[nr][nc] = true;
                        stack.addLast(new int[]{nr, nc, cr, cc});
                    }
                }
            }
        }

        return false;
    }
}
```

### JavaScript

```javascript
/**
 * @param {character[][]} grid
 * @return {boolean}
 */
var containsCycle = function(grid) {
    const m = grid.length;
    const n = grid[0].length;
    const visited = Array.from({ length: m }, () => Array(n).fill(false));

    const dr = [1, -1, 0, 0];
    const dc = [0, 0, 1, -1];

    for (let r = 0; r < m; r++) {
        for (let c = 0; c < n; c++) {
            if (visited[r][c]) continue;

            const stack = [[r, c, -1, -1]];
            visited[r][c] = true;

            while (stack.length > 0) {
                const [cr, cc, pr, pc] = stack.pop();

                for (let k = 0; k < 4; k++) {
                    const nr = cr + dr[k];
                    const nc = cc + dc[k];

                    if (nr < 0 || nr >= m || nc < 0 || nc >= n) continue;
                    if (grid[nr][nc] !== grid[cr][cc]) continue;
                    if (nr === pr && nc === pc) continue;

                    if (visited[nr][nc]) return true;

                    visited[nr][nc] = true;
                    stack.push([nr, nc, cr, cc]);
                }
            }
        }
    }

    return false;
};
```

### Python3

```python
from typing import List

class Solution:
    def containsCycle(self, grid: List[List[str]]) -> bool:
        m, n = len(grid), len(grid[0])
        visited = [[False] * n for _ in range(m)]

        dirs = [(1, 0), (-1, 0), (0, 1), (0, -1)]

        for r in range(m):
            for c in range(n):
                if visited[r][c]:
                    continue

                # stack item: (row, col, parent_row, parent_col)
                stack = [(r, c, -1, -1)]
                visited[r][c] = True

                while stack:
                    cr, cc, pr, pc = stack.pop()

                    for dr, dc in dirs:
                        nr, nc = cr + dr, cc + dc

                        if nr < 0 or nr >= m or nc < 0 or nc >= n:
                            continue
                        if grid[nr][nc] != grid[cr][cc]:
                            continue
                        if nr == pr and nc == pc:
                            continue

                        if visited[nr][nc]:
                            return True

                        visited[nr][nc] = True
                        stack.append((nr, nc, cr, cc))

        return False
```

### Go

```go
func containsCycle(grid [][]byte) bool {
 m := len(grid)
 n := len(grid[0])

 visited := make([][]bool, m)
 for i := 0; i < m; i++ {
  visited[i] = make([]bool, n)
 }

 dr := []int{1, -1, 0, 0}
 dc := []int{0, 0, 1, -1}

 type Node struct {
  r, c int
  pr, pc int
 }

 for r := 0; r < m; r++ {
  for c := 0; c < n; c++ {
   if visited[r][c] {
    continue
   }

   stack := []Node{{r, c, -1, -1}}
   visited[r][c] = true

   for len(stack) > 0 {
    cur := stack[len(stack)-1]
    stack = stack[:len(stack)-1]

    for k := 0; k < 4; k++ {
     nr := cur.r + dr[k]
     nc := cur.c + dc[k]

     if nr < 0 || nr >= m || nc < 0 || nc >= n {
      continue
     }
     if grid[nr][nc] != grid[cur.r][cur.c] {
      continue
     }
     if nr == cur.pr && nc == cur.pc {
      continue
     }

     if visited[nr][nc] {
      return true
     }

     visited[nr][nc] = true
     stack = append(stack, Node{nr, nc, cur.r, cur.c})
    }
   }
  }
 }

 return false
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

### Step 1: Mark visited cells

I create a structure to remember which cells I have already explored.
This prevents repeated work and helps me detect cycles correctly.

### Step 2: Start from every unvisited cell

The grid may have many separate regions.
So I cannot start from only one cell.
I check every cell, and when I find a new one, I launch DFS from it.

### Step 3: Store the parent cell

This is the most important part.
When I move from one cell to another, I save where I came from.
That lets me ignore the direct backward move, which is not a real cycle.

### Step 4: Check all 4 directions

From the current cell, I try moving in all four directions.
For each neighbor, I verify:

* it is inside the grid,
* it has the same character,
* it is not the parent cell.

### Step 5: Detect a cycle

If I find a same-character neighbor that is already visited and is not the parent, then I have found a cycle.
That means the current path has connected back to a previously explored node.

### Step 6: Continue DFS

If the neighbor is new, I mark it visited and push it into the stack.
Then I keep exploring from there.

### Step 7: Return the final answer

If I finish the entire grid and never detect a cycle, then the answer is `false`.
Otherwise, I return `true` as soon as I find one.

## Examples

### Example 1

Input:

```text
grid = [["a","a","a","a"],["a","b","b","a"],["a","b","b","a"],["a","a","a","a"]]
```

Output:

```text
true
```

Explanation:
The `a` cells around the border form a cycle.
The `b` cells in the center also form a cycle.

### Example 2

Input:

```text
grid = [["c","c","c","a"],["c","d","c","c"],["c","c","e","c"],["f","c","c","c"]]
```

Output:

```text
true
```

Explanation:
There is at least one valid cycle made of `c` cells.

### Example 3

Input:

```text
grid = [["a","b","b"],["b","z","b"],["b","b","a"]]
```

Output:

```text
false
```

Explanation:
There is no valid cycle of the same character.

## How to use / Run locally

### C++

* Copy the code into your LeetCode or local C++ file.
* Compile using:

```bash
g++ -std=c++17 main.cpp -o main
./main
```

### Java

* Save the file as `Solution.java`.
* Compile and run:

```bash
javac Solution.java
java Solution
```

### JavaScript

* Use Node.js to test the function.
* Run:

```bash
node main.js
```

### Python3

* Save the file as `main.py`.
* Run:

```bash
python3 main.py
```

### Go

* Save the file as `main.go`.
* Run:

```bash
go run main.go
```

## Notes & Optimizations

* I use iterative DFS to avoid recursion depth issues.
* The parent check is necessary to avoid false positives.
* The solution is efficient because each cell is processed once.
* This method works well for the maximum constraints.

## Author

* [Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
