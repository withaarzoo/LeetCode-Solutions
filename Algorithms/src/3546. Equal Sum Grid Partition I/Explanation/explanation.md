# Problem Title

1. Equal Sum Grid Partition I

## Table of Contents

* Problem Summary
* Constraints
* Intuition
* Approach
* Data Structures Used
* Operations & Behavior Summary
* Complexity
* Multi-language Solutions

  * C++
  * Java
  * JavaScript
  * Python3
  * Go
* Step-by-step Detailed Explanation
* Examples
* How to use / Run locally
* Notes & Optimizations
* Author

## Problem Summary

You are given an m x n matrix of positive integers. You need to determine whether it is possible to make either one horizontal or one vertical cut such that:

* Both resulting parts are non-empty
* The sum of elements in both parts is equal

Return true if such a cut exists, otherwise return false.

## Constraints

* 1 <= m == grid.length <= 10^5
* 1 <= n == grid[i].length <= 10^5
* 2 <= m * n <= 10^5
* 1 <= grid[i][j] <= 10^5

## Intuition

I thought about reducing the problem to a simple observation.

If I want to split the grid into two parts with equal sum, then each part must have totalSum / 2.

So first I compute the total sum of the grid.

* If total sum is odd, equal partition is impossible.
* If total sum is even, I try to find a cut where one side equals totalSum / 2.

## Approach

1. Calculate the total sum of all elements.
2. If total sum is odd, return false.
3. Set target = totalSum / 2.
4. Try horizontal cuts:

   * Keep adding row sums.
   * After each row, check if accumulated sum == target.
5. Try vertical cuts:

   * Precompute column sums.
   * Accumulate column sums and check for target.
6. If any valid cut is found, return true.
7. Otherwise, return false.

## Data Structures Used

* Array / Matrix (input grid)
* Auxiliary array for column sums
* Integer / Long variables for accumulation

## Operations & Behavior Summary

* Traverse entire grid to compute total sum
* Traverse rows once for horizontal cut
* Traverse columns once for vertical cut
* Early exit when valid partition is found

## Complexity

* Time Complexity: O(m * n)

  * Full traversal for sum, row accumulation, and column accumulation
* Space Complexity: O(n)

  * Column sum array

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    bool canPartitionGrid(vector<vector<int>>& grid) {
        int m = grid.size(), n = grid[0].size();
        long long total = 0;

        for (auto &row : grid)
            for (int val : row)
                total += val;

        if (total % 2 != 0) return false;
        long long target = total / 2;

        long long rowSum = 0;
        for (int i = 0; i < m - 1; i++) {
            for (int j = 0; j < n; j++)
                rowSum += grid[i][j];
            if (rowSum == target) return true;
        }

        vector<long long> colSum(n, 0);
        for (int j = 0; j < n; j++)
            for (int i = 0; i < m; i++)
                colSum[j] += grid[i][j];

        long long curr = 0;
        for (int j = 0; j < n - 1; j++) {
            curr += colSum[j];
            if (curr == target) return true;
        }

        return false;
    }
};
```

### Java

```java
class Solution {
    public boolean canPartitionGrid(int[][] grid) {
        int m = grid.length, n = grid[0].length;
        long total = 0;

        for (int[] row : grid)
            for (int val : row)
                total += val;

        if (total % 2 != 0) return false;
        long target = total / 2;

        long rowSum = 0;
        for (int i = 0; i < m - 1; i++) {
            for (int j = 0; j < n; j++)
                rowSum += grid[i][j];
            if (rowSum == target) return true;
        }

        long[] colSum = new long[n];
        for (int j = 0; j < n; j++)
            for (int i = 0; i < m; i++)
                colSum[j] += grid[i][j];

        long curr = 0;
        for (int j = 0; j < n - 1; j++) {
            curr += colSum[j];
            if (curr == target) return true;
        }

        return false;
    }
}
```

### JavaScript

```javascript
var canPartitionGrid = function(grid) {
    let m = grid.length, n = grid[0].length;
    let total = 0;

    for (let row of grid)
        for (let val of row)
            total += val;

    if (total % 2 !== 0) return false;
    let target = total / 2;

    let rowSum = 0;
    for (let i = 0; i < m - 1; i++) {
        for (let j = 0; j < n; j++)
            rowSum += grid[i][j];
        if (rowSum === target) return true;
    }

    let colSum = new Array(n).fill(0);
    for (let j = 0; j < n; j++)
        for (let i = 0; i < m; i++)
            colSum[j] += grid[i][j];

    let curr = 0;
    for (let j = 0; j < n - 1; j++) {
        curr += colSum[j];
        if (curr === target) return true;
    }

    return false;
};
```

### Python3

```python
class Solution:
    def canPartitionGrid(self, grid):
        m, n = len(grid), len(grid[0])
        total = sum(sum(row) for row in grid)

        if total % 2 != 0:
            return False

        target = total // 2

        row_sum = 0
        for i in range(m - 1):
            row_sum += sum(grid[i])
            if row_sum == target:
                return True

        col_sum = [0] * n
        for j in range(n):
            for i in range(m):
                col_sum[j] += grid[i][j]

        curr = 0
        for j in range(n - 1):
            curr += col_sum[j]
            if curr == target:
                return True

        return False
```

### Go

```go
func canPartitionGrid(grid [][]int) bool {
    m, n := len(grid), len(grid[0])
    total := 0

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            total += grid[i][j]
        }
    }

    if total%2 != 0 {
        return false
    }

    target := total / 2

    rowSum := 0
    for i := 0; i < m-1; i++ {
        for j := 0; j < n; j++ {
            rowSum += grid[i][j]
        }
        if rowSum == target {
            return true
        }
    }

    colSum := make([]int, n)
    for j := 0; j < n; j++ {
        for i := 0; i < m; i++ {
            colSum[j] += grid[i][j]
        }
    }

    curr := 0
    for j := 0; j < n-1; j++ {
        curr += colSum[j]
        if curr == target {
            return true
        }
    }

    return false
}
```

## Step-by-step Detailed Explanation

1. Compute total sum of the grid.
2. Check if total sum is even.
3. Set target as half of total sum.
4. Accumulate row sums one by one.
5. If row sum equals target, horizontal cut works.
6. Precompute column sums.
7. Accumulate column sums.
8. If column sum equals target, vertical cut works.
9. If neither works, return false.

## Examples

Input: [[1,4],[2,3]]
Output: true

Input: [[1,3],[2,4]]
Output: false

## How to use / Run locally

1. Copy the code into your local IDE.
2. Provide input as a 2D array.
3. Call the function.
4. Print the result.

## Notes & Optimizations

* Early exit improves performance.
* Use long/64-bit integers to avoid overflow.
* Column precomputation avoids repeated work.

## Author

* [Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
