# Problem Title

Maximum Non-Negative Product in a Matrix

## Table of Contents

* Problem Summary
* Constraints
* Intuition
* Approach
* Data Structures Used
* Operations & Behavior Summary
* Complexity
* Multi-language Solutions
* Step-by-step Detailed Explanation
* Examples
* How to use / Run locally
* Notes & Optimizations
* Author

## Problem Summary

Given an m x n grid, I start from the top-left cell (0,0) and can only move right or down. Each path forms a product of visited cells. I need to find the maximum non-negative product possible. If all paths result in a negative product, I return -1.

## Constraints

* 1 <= m, n <= 15
* -4 <= grid[i][j] <= 4

## Intuition

I noticed that negative numbers can flip the result. A negative product can become positive if multiplied again by another negative number. So tracking only the maximum is not enough. I must also track the minimum (most negative) value.

## Approach

I use dynamic programming with two states at each cell:

* maxDp[i][j]: maximum product up to that cell
* minDp[i][j]: minimum product up to that cell

From top and left, I calculate all possibilities and update both values.

## Data Structures Used

* 2D arrays for max and min product tracking

## Operations & Behavior Summary

* Multiply current cell value with previous max and min
* Choose best and worst outcomes
* Continue till bottom-right

## Complexity

* Time Complexity: O(m * n)
* Space Complexity: O(m * n)

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int maxProductPath(vector<vector<int>>& grid) {
        int m = grid.size(), n = grid[0].size();
        const long long MOD = 1e9 + 7;

        vector<vector<long long>> maxDp(m, vector<long long>(n));
        vector<vector<long long>> minDp(m, vector<long long>(n));

        maxDp[0][0] = minDp[0][0] = grid[0][0];

        for (int i = 1; i < m; i++)
            maxDp[i][0] = minDp[i][0] = maxDp[i-1][0] * grid[i][0];

        for (int j = 1; j < n; j++)
            maxDp[0][j] = minDp[0][j] = maxDp[0][j-1] * grid[0][j];

        for (int i = 1; i < m; i++) {
            for (int j = 1; j < n; j++) {
                long long val = grid[i][j];
                long long a = maxDp[i-1][j] * val;
                long long b = minDp[i-1][j] * val;
                long long c = maxDp[i][j-1] * val;
                long long d = minDp[i][j-1] * val;

                maxDp[i][j] = max({a, b, c, d});
                minDp[i][j] = min({a, b, c, d});
            }
        }

        long long res = maxDp[m-1][n-1];
        return res < 0 ? -1 : res % MOD;
    }
};
```

### Java

```java
class Solution {
    public int maxProductPath(int[][] grid) {
        int m = grid.length, n = grid[0].length;
        long MOD = (long)1e9 + 7;

        long[][] maxDp = new long[m][n];
        long[][] minDp = new long[m][n];

        maxDp[0][0] = minDp[0][0] = grid[0][0];

        for (int i = 1; i < m; i++)
            maxDp[i][0] = minDp[i][0] = maxDp[i-1][0] * grid[i][0];

        for (int j = 1; j < n; j++)
            maxDp[0][j] = minDp[0][j] = maxDp[0][j-1] * grid[0][j];

        for (int i = 1; i < m; i++) {
            for (int j = 1; j < n; j++) {
                long val = grid[i][j];

                long a = maxDp[i-1][j] * val;
                long b = minDp[i-1][j] * val;
                long c = maxDp[i][j-1] * val;
                long d = minDp[i][j-1] * val;

                maxDp[i][j] = Math.max(Math.max(a, b), Math.max(c, d));
                minDp[i][j] = Math.min(Math.min(a, b), Math.min(c, d));
            }
        }

        long res = maxDp[m-1][n-1];
        return res < 0 ? -1 : (int)(res % MOD);
    }
}
```

### JavaScript

```javascript
var maxProductPath = function(grid) {
    const m = grid.length, n = grid[0].length;
    const MOD = 1e9 + 7;

    const maxDp = Array.from({length: m}, () => Array(n).fill(0));
    const minDp = Array.from({length: m}, () => Array(n).fill(0));

    maxDp[0][0] = minDp[0][0] = grid[0][0];

    for (let i = 1; i < m; i++)
        maxDp[i][0] = minDp[i][0] = maxDp[i-1][0] * grid[i][0];

    for (let j = 1; j < n; j++)
        maxDp[0][j] = minDp[0][j] = maxDp[0][j-1] * grid[0][j];

    for (let i = 1; i < m; i++) {
        for (let j = 1; j < n; j++) {
            let val = grid[i][j];

            let a = maxDp[i-1][j] * val;
            let b = minDp[i-1][j] * val;
            let c = maxDp[i][j-1] * val;
            let d = minDp[i][j-1] * val;

            maxDp[i][j] = Math.max(a, b, c, d);
            minDp[i][j] = Math.min(a, b, c, d);
        }
    }

    let res = maxDp[m-1][n-1];
    return res < 0 ? -1 : res % MOD;
};
```

### Python3

```python
class Solution:
    def maxProductPath(self, grid: List[List[int]]) -> int:
        m, n = len(grid), len(grid[0])
        MOD = 10**9 + 7

        maxDp = [[0]*n for _ in range(m)]
        minDp = [[0]*n for _ in range(m)]

        maxDp[0][0] = minDp[0][0] = grid[0][0]

        for i in range(1, m):
            maxDp[i][0] = minDp[i][0] = maxDp[i-1][0] * grid[i][0]

        for j in range(1, n):
            maxDp[0][j] = minDp[0][j] = maxDp[0][j-1] * grid[0][j]

        for i in range(1, m):
            for j in range(1, n):
                val = grid[i][j]

                candidates = [
                    maxDp[i-1][j] * val,
                    minDp[i-1][j] * val,
                    maxDp[i][j-1] * val,
                    minDp[i][j-1] * val
                ]

                maxDp[i][j] = max(candidates)
                minDp[i][j] = min(candidates)

        res = maxDp[m-1][n-1]
        return -1 if res < 0 else res % MOD
```

### Go

```go
func maxProductPath(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    MOD := int64(1e9 + 7)

    maxDp := make([][]int64, m)
    minDp := make([][]int64, m)

    for i := range maxDp {
        maxDp[i] = make([]int64, n)
        minDp[i] = make([]int64, n)
    }

    maxDp[0][0] = int64(grid[0][0])
    minDp[0][0] = int64(grid[0][0])

    for i := 1; i < m; i++ {
        maxDp[i][0] = maxDp[i-1][0] * int64(grid[i][0])
        minDp[i][0] = maxDp[i][0]
    }

    for j := 1; j < n; j++ {
        maxDp[0][j] = maxDp[0][j-1] * int64(grid[0][j])
        minDp[0][j] = maxDp[0][j]
    }

    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            val := int64(grid[i][j])

            a := maxDp[i-1][j] * val
            b := minDp[i-1][j] * val
            c := maxDp[i][j-1] * val
            d := minDp[i][j-1] * val

            maxDp[i][j] = max4(a, b, c, d)
            minDp[i][j] = min4(a, b, c, d)
        }
    }

    res := maxDp[m-1][n-1]
    if res < 0 {
        return -1
    }
    return int(res % MOD)
}
```

## Step-by-step Detailed Explanation

At each cell, I compute four values using top and left cells. I multiply current value with both maximum and minimum values from previous states. Then I update current cell's max and min accordingly.

## Examples

Input: [[1,-2,1],[1,-2,1],[3,-4,1]]
Output: 8

## How to use / Run locally

* Copy code into your preferred language
* Compile and run using standard tools

## Notes & Optimizations

* Can be optimized to O(n) space
* Always track both min and max

## Author

* [Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
