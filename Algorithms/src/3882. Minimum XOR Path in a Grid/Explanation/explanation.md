# Minimum XOR Path in a Grid – LeetCode 3882 Solution

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

You are given a 2D grid of integers. You start at the top-left cell and must reach the bottom-right cell. From any cell you can only move right or down.

The cost of a path is the bitwise XOR of every value you visit, including the starting cell and the ending cell. Your job is to find the smallest possible XOR value among all valid paths.

This is a classic grid path problem mixed with bit manipulation. Because XOR does not behave like normal addition, keeping only the minimum value at each cell fails. You need to track every possible XOR that can reach a cell.

## Constraints

- 1 ≤ m = grid.length ≤ 1000
- 1 ≤ n = grid[i].length ≤ 1000
- m × n ≤ 1000
- 0 ≤ grid[i][j] ≤ 1023

The product of rows and columns is at most 1000, and every cell value fits in 10 bits. These two facts make a full dynamic-programming solution practical.

## Intuition

The first thing I noticed is that every grid value is at most 1023. That means any XOR of those values also stays inside the range 0 to 1023. There are only 1024 possible XOR results.

Because we can only move right or down, the order of cells is fixed. Once I know every XOR that can reach the cell above me and the cell to my left, I can generate every XOR that can reach the current cell by simply XORing those previous values with the current cell value.

I therefore decided to keep a set of reachable XOR values for every cell. With only 1024 possible values the set is tiny, and the total work stays well under a million operations.

## Approach

I create a three-dimensional structure that stores a boolean flag for every cell and every possible XOR value from 0 to 1023. A true flag means “there exists at least one path that reaches this cell with this exact XOR”.

I mark the starting cell with the single value that sits in grid[0][0]. Then I process the grid from top to bottom and left to right. For each cell I look at the cell above (if it exists) and the cell on the left (if it exists). For every XOR that is already reachable in those neighbors I compute the new XOR by combining it with the current cell value and I set the corresponding flag in the current cell.

When the whole grid has been processed I simply scan the flags of the bottom-right cell from 0 upward and return the first true flag. That is the minimum reachable XOR.

## Data Structures Used

- A three-dimensional boolean array (or an equivalent bitset / Uint8Array) of size m × n × 1024.  
  This is the core of the solution. Each slice of 1024 bits records every XOR that can reach that particular cell.

- No extra queues, heaps or maps are required. The fixed-size boolean container is both fast and memory-efficient because m × n never exceeds 1000.

## Operations & Behavior Summary

1. Allocate a boolean container of size 1024 for every cell.
2. Set the flag that matches grid[0][0] in the top-left cell.
3. For every cell in row-major order:
   - If a cell exists above, copy every reachable XOR from that cell, XOR it with the current value, and mark the result in the current cell.
   - If a cell exists to the left, do the same with the left neighbor.
4. After the loops finish, walk the 1024 flags of the bottom-right cell from the smallest index to the largest and return the first true flag.

## Complexity

| Type              | Value                  | Explanation |
|-------------------|------------------------|-------------|
| Time Complexity   | O(m × n × 1024)        | For each of the ≤ 1000 cells we examine up to 1024 possible previous XOR values. |
| Space Complexity  | O(m × n × 1024)        | We store a 1024-bit container for every cell. Because m × n ≤ 1000 the total memory stays under 2 MB. |

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int minCost(vector<vector<int>>& grid) {
        int m = grid.size(), n = grid[0].size();
        vector<vector<bitset<1024>>> dp(m, vector<bitset<1024>>(n));
        dp[0][0].set(grid[0][0]);
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (i > 0) {
                    for (int x = 0; x < 1024; x++) {
                        if (dp[i - 1][j][x]) dp[i][j].set(x ^ grid[i][j]);
                    }
                }
                if (j > 0) {
                    for (int x = 0; x < 1024; x++) {
                        if (dp[i][j - 1][x]) dp[i][j].set(x ^ grid[i][j]);
                    }
                }
            }
        }
        for (int x = 0; x < 1024; x++) {
            if (dp[m - 1][n - 1][x]) return x;
        }
        return 0;
    }
};
```

### Java

```java
class Solution {
    public int minCost(int[][] grid) {
        int m = grid.length, n = grid[0].length;
        boolean[][][] dp = new boolean[m][n][1024];
        dp[0][0][grid[0][0]] = true;
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (i > 0) {
                    for (int x = 0; x < 1024; x++) {
                        if (dp[i - 1][j][x]) dp[i][j][x ^ grid[i][j]] = true;
                    }
                }
                if (j > 0) {
                    for (int x = 0; x < 1024; x++) {
                        if (dp[i][j - 1][x]) dp[i][j][x ^ grid[i][j]] = true;
                    }
                }
            }
        }
        for (int x = 0; x < 1024; x++) {
            if (dp[m - 1][n - 1][x]) return x;
        }
        return 0;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[][]} grid
 * @return {number}
 */
var minCost = function(grid) {
    const m = grid.length, n = grid[0].length;
    const dp = Array.from({length: m}, () => Array.from({length: n}, () => new Uint8Array(1024)));
    dp[0][0][grid[0][0]] = 1;
    for (let i = 0; i < m; i++) {
        for (let j = 0; j < n; j++) {
            if (i > 0) {
                for (let x = 0; x < 1024; x++) {
                    if (dp[i - 1][j][x]) dp[i][j][x ^ grid[i][j]] = 1;
                }
            }
            if (j > 0) {
                for (let x = 0; x < 1024; x++) {
                    if (dp[i][j - 1][x]) dp[i][j][x ^ grid[i][j]] = 1;
                }
            }
        }
    }
    for (let x = 0; x < 1024; x++) {
        if (dp[m - 1][n - 1][x]) return x;
    }
    return 0;
};
```

### TypeScript

```typescript
function minCost(grid: number[][]): number {
    const m = grid.length, n = grid[0].length;
    const dp: Uint8Array[][] = Array.from({length: m}, () => Array.from({length: n}, () => new Uint8Array(1024)));
    dp[0][0][grid[0][0]] = 1;
    for (let i = 0; i < m; i++) {
        for (let j = 0; j < n; j++) {
            if (i > 0) {
                for (let x = 0; x < 1024; x++) {
                    if (dp[i - 1][j][x]) dp[i][j][x ^ grid[i][j]] = 1;
                }
            }
            if (j > 0) {
                for (let x = 0; x < 1024; x++) {
                    if (dp[i][j - 1][x]) dp[i][j][x ^ grid[i][j]] = 1;
                }
            }
        }
    }
    for (let x = 0; x < 1024; x++) {
        if (dp[m - 1][n - 1][x]) return x;
    }
    return 0;
};
```

### Python3

```python
class Solution:
    def minCost(self, grid: list[list[int]]) -> int:
        m, n = len(grid), len(grid[0])
        dp = [[[False] * 1024 for _ in range(n)] for _ in range(m)]
        dp[0][0][grid[0][0]] = True
        for i in range(m):
            for j in range(n):
                if i > 0:
                    for x in range(1024):
                        if dp[i - 1][j][x]:
                            dp[i][j][x ^ grid[i][j]] = True
                if j > 0:
                    for x in range(1024):
                        if dp[i][j - 1][x]:
                            dp[i][j][x ^ grid[i][j]] = True
        for x in range(1024):
            if dp[m - 1][n - 1][x]:
                return x
        return 0
```

### Go

```go
func minCost(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    dp := make([][][1024]bool, m)
    for i := range dp {
        dp[i] = make([][1024]bool, n)
    }
    dp[0][0][grid[0][0]] = true
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if i > 0 {
                for x := 0; x < 1024; x++ {
                    if dp[i-1][j][x] {
                        dp[i][j][x^grid[i][j]] = true
                    }
                }
            }
            if j > 0 {
                for x := 0; x < 1024; x++ {
                    if dp[i][j-1][x] {
                        dp[i][j][x^grid[i][j]] = true
                    }
                }
            }
        }
    }
    for x := 0; x < 1024; x++ {
        if dp[m-1][n-1][x] {
            return x
        }
    }
    return 0
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

The logic is identical in every language; only the syntax for creating the three-dimensional boolean container changes.

First I read the dimensions m and n. Then I allocate a structure that can hold a true/false value for every combination of row, column and XOR result. In C++ I use vector of bitset, in Java a three-dimensional boolean array, in JavaScript and TypeScript a nested array of Uint8Array, in Python a list of lists of lists, and in Go a slice of slices of fixed-size boolean arrays.

I immediately mark the starting cell: the only reachable XOR is the value sitting in grid[0][0].

I then walk every cell. Because movement is only right or down, when I reach cell (i, j) the cells (i-1, j) and (i, j-1) have already been fully computed. For each of those two neighbors I loop over the 1024 possible XOR values. Whenever a flag is true I compute newXor = oldXor XOR grid[i][j] and set the corresponding flag in the current cell.

After the double loop finishes I look at the bottom-right cell and return the smallest index whose flag is true. If somehow no path existed the function would return 0, but the problem guarantees at least one path.

The same reasoning applies to every language. The only differences are how boolean arrays are declared and how the XOR operator is written (^ in most languages, the same symbol in all of them).

## Examples

Example 1  
Input: grid = [[1,2],[3,4]]  
Output: 6  

- Path 1 → 2 → 4 gives 1 ⊕ 2 ⊕ 4 = 7  
- Path 1 → 3 → 4 gives 1 ⊕ 3 ⊕ 4 = 6  
The algorithm records both 7 and 6 at the bottom-right cell and returns the smaller one.

Example 2  
Input: grid = [[6,7],[5,8]]  
Output: 9  

- 6 ⊕ 7 ⊕ 8 = 9  
- 6 ⊕ 5 ⊕ 8 = 11  
Minimum is 9.

Example 3  
Input: grid = [[2,7,5]]  
Output: 0  

Only one path exists: 2 ⊕ 7 ⊕ 5 = 0. The single flag for value 0 is set and returned.

## How to Use / Run Locally

1. Copy the code for the language you want into a file named Solution.cpp, Solution.java, solution.js, solution.ts, solution.py or solution.go.
2. C++: g++ -std=c++17 Solution.cpp -o sol && ./sol  
3. Java: javac Solution.java && java Solution  
4. JavaScript: node solution.js  
5. TypeScript: tsc solution.ts && node solution.js  
6. Python: python3 solution.py  
7. Go: go run solution.go  

Inside the main function (or at the bottom of the script) create a small test grid, call the function and print the result so you can verify it matches the examples.

## Notes & Optimizations

- Because every value is ≤ 1023 we never need more than 1024 bits. Using a fixed-size container instead of a hash set removes all hashing overhead and produces the fastest possible solution.
- Space can be reduced further with a rolling-array technique (keep only the previous row), but the current memory usage is already tiny.
- A pure recursive DFS with memoization also works, yet the bottom-up version is simpler and has better constant factors.
- Edge case of a 1×1 grid is handled automatically: the answer is simply the single cell value.

## Author

[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)
