# Count Islands With Total Value Divisible by K – LeetCode 3619 Solution

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

You are given an m x n grid filled with non-negative integers and a positive integer k.  

Any group of positive numbers that are connected horizontally or vertically forms an island. Zeros act as water and separate the islands.  

The total value of an island is simply the sum of all numbers that belong to it.  

Your task is to return how many islands have a total value that is divisible by k (that is, sum % k == 0).

This is a classic “Number of Islands” style problem with an extra condition on the sum of each connected component.

## Constraints

- m == grid.length  
- n == grid[i].length  
- 1 <= m, n <= 1000  
- 1 <= m * n <= 10^5  
- 0 <= grid[i][j] <= 10^6  
- 1 <= k <= 10^6  

## Intuition

The moment I saw the problem I recognized the classic island-counting pattern. Positive cells are land, zeros are water, and connectivity is only four-directional.  

The only extra requirement is that I must keep track of the sum of every island while I explore it. Once the whole island is visited I just check whether that sum is divisible by k.  

Because the grid can be as large as 10^5 cells, any solution that visits each cell a constant number of times will be fast enough. That immediately points to a single DFS (or BFS) pass.

## Approach

I walk through every cell of the grid.  

Whenever I land on a positive number that has not been visited yet, I start a depth-first search from that cell. During the search I:  

1. Add the current cell’s value to a running sum.  
2. Mark the cell as visited by writing zero into it (so I never process it again).  
3. Recursively explore its four neighbors.  

When the DFS returns, I have the complete sum of the island. If the sum is divisible by k I increase my answer counter by one.  

I continue scanning the grid until every cell has been examined. At the end the counter holds the number of islands whose total value is divisible by k.

## Data Structures Used

- The input grid itself – I reuse it as a visited matrix by setting visited land cells to zero. This avoids allocating an extra boolean matrix and saves memory.  
- Recursion stack – the DFS uses the call stack to keep track of the current path. In the worst case the stack depth can be O(m * n).  
- A simple integer (or long) variable to accumulate the island sum and another integer for the final answer.

No queues, sets, or other heavy structures are needed.

## Operations & Behavior Summary

1. Read the dimensions of the grid.  
2. Initialize answer = 0.  
3. For every cell (i, j):  
   - If grid[i][j] > 0, call DFS.  
   - DFS returns the total sum of the connected component.  
   - If sum % k == 0, increment answer.  
4. Inside DFS:  
   - Base case: out of bounds or already zero → return 0.  
   - Otherwise add the current value, set the cell to zero, and sum the four recursive calls.  
5. Return the final answer.

## Complexity

| Type              | Complexity | Explanation |
|-------------------|------------|-------------|
| Time Complexity   | O(m * n)   | Every cell is visited at most a constant number of times. |
| Space Complexity  | O(m * n)   | Worst-case recursion depth when the whole grid is one island. No extra visited matrix is allocated. |

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int countIslands(vector<vector<int>>& grid, int k) {
        int m = grid.size(), n = grid[0].size();
        int ans = 0;
        auto dfs = [&](auto&& self, int i, int j) -> long long {
            if (i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0) return 0;
            long long sum = grid[i][j];
            grid[i][j] = 0;
            sum += self(self, i + 1, j);
            sum += self(self, i - 1, j);
            sum += self(self, i, j + 1);
            sum += self(self, i, j - 1);
            return sum;
        };
        for (int i = 0; i < m; ++i) {
            for (int j = 0; j < n; ++j) {
                if (grid[i][j] > 0) {
                    long long s = dfs(dfs, i, j);
                    if (s % k == 0) ++ans;
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
    public int countIslands(int[][] grid, int k) {
        int m = grid.length, n = grid[0].length;
        int ans = 0;
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (grid[i][j] > 0) {
                    long s = dfs(grid, i, j);
                    if (s % k == 0) ans++;
                }
            }
        }
        return ans;
    }
    private long dfs(int[][] grid, int i, int j) {
        int m = grid.length, n = grid[0].length;
        if (i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0) return 0;
        long sum = grid[i][j];
        grid[i][j] = 0;
        sum += dfs(grid, i + 1, j);
        sum += dfs(grid, i - 1, j);
        sum += dfs(grid, i, j + 1);
        sum += dfs(grid, i, j - 1);
        return sum;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[][]} grid
 * @param {number} k
 * @return {number}
 */
var countIslands = function(grid, k) {
    const m = grid.length, n = grid[0].length;
    let ans = 0;
    const dfs = (i, j) => {
        if (i < 0 || i >= m || j < 0 || j >= n || grid[i][j] === 0) return 0;
        let sum = grid[i][j];
        grid[i][j] = 0;
        sum += dfs(i + 1, j);
        sum += dfs(i - 1, j);
        sum += dfs(i, j + 1);
        sum += dfs(i, j - 1);
        return sum;
    };
    for (let i = 0; i < m; i++) {
        for (let j = 0; j < n; j++) {
            if (grid[i][j] > 0) {
                const s = dfs(i, j);
                if (s % k === 0) ans++;
            }
        }
    }
    return ans;
};
```

### TypeScript

```typescript
function countIslands(grid: number[][], k: number): number {
    const m = grid.length, n = grid[0].length;
    let ans = 0;
    const dfs = (i: number, j: number): number => {
        if (i < 0 || i >= m || j < 0 || j >= n || grid[i][j] === 0) return 0;
        let sum = grid[i][j];
        grid[i][j] = 0;
        sum += dfs(i + 1, j);
        sum += dfs(i - 1, j);
        sum += dfs(i, j + 1);
        sum += dfs(i, j - 1);
        return sum;
    };
    for (let i = 0; i < m; i++) {
        for (let j = 0; j < n; j++) {
            if (grid[i][j] > 0) {
                const s = dfs(i, j);
                if (s % k === 0) ans++;
            }
        }
    }
    return ans;
};
```

### Python3

```python
class Solution:
    def countIslands(self, grid: List[List[int]], k: int) -> int:
        m, n = len(grid), len(grid[0])
        ans = 0
        def dfs(i: int, j: int) -> int:
            if i < 0 or i >= m or j < 0 or j >= n or grid[i][j] == 0:
                return 0
            s = grid[i][j]
            grid[i][j] = 0
            s += dfs(i + 1, j)
            s += dfs(i - 1, j)
            s += dfs(i, j + 1)
            s += dfs(i, j - 1)
            return s
        for i in range(m):
            for j in range(n):
                if grid[i][j] > 0:
                    if dfs(i, j) % k == 0:
                        ans += 1
        return ans
```

### Go

```go
func countIslands(grid [][]int, k int) int {
    m, n := len(grid), len(grid[0])
    ans := 0
    var dfs func(i, j int) int64
    dfs = func(i, j int) int64 {
        if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 {
            return 0
        }
        sum := int64(grid[i][j])
        grid[i][j] = 0
        sum += dfs(i+1, j)
        sum += dfs(i-1, j)
        sum += dfs(i, j+1)
        sum += dfs(i, j-1)
        return sum
    }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] > 0 {
                if dfs(i, j)%int64(k) == 0 {
                    ans++
                }
            }
        }
    }
    return ans
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

All six implementations follow the exact same logic; only the language syntax changes.

I begin by storing the number of rows and columns so the boundary checks stay clean.  

A counter keeps track of how many good islands I have found.  

The heart of the solution is the DFS helper. It receives a row and a column. The first thing it does is decide whether the current position is still valid land. If the coordinates are outside the grid or the cell is already zero, there is nothing to add, so it returns zero immediately.  

Otherwise it records the value that sits in the cell, then immediately overwrites that cell with zero. Writing zero is both the “visited” mark and the way I prevent infinite recursion. After the cell is marked, the helper calls itself on the four neighbors (down, up, right, left) and adds whatever sums those calls return. The total is handed back to the caller.  

Back in the main double loop I simply look for any remaining positive cell. Each such cell is the start of a brand-new island, so I invoke DFS and receive the full island sum. A single modulo check decides whether that island should be counted.  

Because a cell value can be as large as 10^6 and an island can contain up to 10^5 cells, the sum may exceed a 32-bit integer. That is why the implementations use a 64-bit integer (long in Java, long long in C++, int64 in Go, number in JavaScript/TypeScript which is safe for this range, and plain int in Python which has arbitrary precision).  

Edge cases that are automatically handled:  

- An empty island never occurs because we only start DFS on positive cells.  
- A single isolated positive cell is treated as an island of size one.  
- When k equals 1 every island is counted, which is correct.  
- The grid can contain zeros only; the answer is then zero.

The same reasoning applies unchanged to every language version.

## Examples

**Example 1**  
Input:  
grid = [[0,2,1,0,0],[0,5,0,0,5],[0,0,1,0,0],[0,1,4,7,0],[0,2,0,0,8]], k = 5  

Islands and their sums:  

- {2,5,1,1,4,7,2} → sum 22 (not divisible by 5)  
- {5} → sum 5 (divisible)  
- {8} → sum 8 (not divisible)  

Two islands satisfy the condition, so the output is 2.

**Example 2**  
Input:  
grid = [[3,0,3,0],[0,3,0,3],[3,0,3,0]], k = 3  

There are six separate cells, each with value 3. Every sum is 3, which is divisible by 3. Output is 6.

**Example 3 (single cell)**  
grid = [[7]], k = 7 → output 1  
grid = [[7]], k = 3 → output 0

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

Replace the driver code with a small main that builds the sample grids and prints the result of countIslands.

## Notes & Optimizations

- Mutating the original grid is the most memory-efficient way to mark visited cells. If the problem required the grid to stay unchanged you would allocate a separate boolean matrix.  
- An iterative BFS with a queue works equally well and avoids deep recursion. The asymptotic complexity stays the same.  
- Because m * n ≤ 10^5 the recursive solution is accepted on all major online judges; stack-overflow risk is low.  
- The modulo operation is performed only once per island, not per cell, which keeps the constant factor small.  
- If k is larger than any possible island sum you can early-exit, but the extra check is rarely worth the complexity.

## Author

[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)
