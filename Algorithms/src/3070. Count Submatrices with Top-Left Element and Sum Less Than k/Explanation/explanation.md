# Problem Title

1. Count Submatrices with Top-Left Element and Sum Less Than k

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
* [Step-by-step Detailed Explanation](#step-by-step-detailed-explanation-c-java-javascript-python3-go)
* [Examples](#examples)
* [How to use / Run locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

## Problem Summary

Given a 0-indexed integer matrix `grid` and an integer `k`, return the number of submatrices that:

* Include the top-left element `(0,0)`
* Have a sum less than or equal to `k`

## Constraints

* `m == grid.length`
* `n == grid[i].length`
* `1 <= m, n <= 1000`
* `0 <= grid[i][j] <= 1000`
* `1 <= k <= 1e9`

## Intuition

I noticed that every valid submatrix must start from `(0,0)`. So instead of checking all possible submatrices, I only need to consider submatrices that end at `(i, j)`.

So the problem reduces to:
Count how many positions `(i, j)` have sum of submatrix `(0,0) → (i,j)` less than or equal to `k`.

To compute sums efficiently, I used a 2D prefix sum.

## Approach

1. Traverse the grid.
2. Convert the grid into a prefix sum matrix.
3. For each cell `(i, j)`, compute cumulative sum from `(0,0)`.
4. If prefix sum is `<= k`, increment the count.
5. Return the final count.

## Data Structures Used

* 2D array (grid reused as prefix sum matrix)
* Integer counter variable

## Operations & Behavior Summary

* Build prefix sum in-place
* Each cell stores sum from `(0,0)` to `(i,j)`
* Compare each prefix sum with `k`
* Count valid submatrices

## Complexity

* Time Complexity: O(m * n)
* Space Complexity: O(1) (in-place computation)

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int countSubmatrices(vector<vector<int>>& grid, int k) {
        int m = grid.size(), n = grid[0].size();
        int count = 0;

        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (i > 0) grid[i][j] += grid[i - 1][j];
                if (j > 0) grid[i][j] += grid[i][j - 1];
                if (i > 0 && j > 0) grid[i][j] -= grid[i - 1][j - 1];

                if (grid[i][j] <= k) count++;
            }
        }
        return count;
    }
};
```

### Java

```java
class Solution {
    public int countSubmatrices(int[][] grid, int k) {
        int m = grid.length, n = grid[0].length;
        int count = 0;

        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (i > 0) grid[i][j] += grid[i - 1][j];
                if (j > 0) grid[i][j] += grid[i][j - 1];
                if (i > 0 && j > 0) grid[i][j] -= grid[i - 1][j - 1];

                if (grid[i][j] <= k) count++;
            }
        }
        return count;
    }
}
```

### JavaScript

```javascript
var countSubmatrices = function(grid, k) {
    let m = grid.length, n = grid[0].length;
    let count = 0;

    for (let i = 0; i < m; i++) {
        for (let j = 0; j < n; j++) {
            if (i > 0) grid[i][j] += grid[i - 1][j];
            if (j > 0) grid[i][j] += grid[i][j - 1];
            if (i > 0 && j > 0) grid[i][j] -= grid[i - 1][j - 1];

            if (grid[i][j] <= k) count++;
        }
    }
    return count;
};
```

### Python3

```python
class Solution:
    def countSubmatrices(self, grid, k):
        m, n = len(grid), len(grid[0])
        count = 0

        for i in range(m):
            for j in range(n):
                if i > 0:
                    grid[i][j] += grid[i - 1][j]
                if j > 0:
                    grid[i][j] += grid[i][j - 1]
                if i > 0 and j > 0:
                    grid[i][j] -= grid[i - 1][j - 1]

                if grid[i][j] <= k:
                    count += 1

        return count
```

### Go

```go
func countSubmatrices(grid [][]int, k int) int {
    m, n := len(grid), len(grid[0])
    count := 0

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if i > 0 {
                grid[i][j] += grid[i-1][j]
            }
            if j > 0 {
                grid[i][j] += grid[i][j-1]
            }
            if i > 0 && j > 0 {
                grid[i][j] -= grid[i-1][j-1]
            }

            if grid[i][j] <= k {
                count++
            }
        }
    }

    return count
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

1. Traverse each cell `(i, j)`.
2. Add value from top cell.
3. Add value from left cell.
4. Subtract overlapping top-left value.
5. Now current cell holds prefix sum.
6. If prefix sum `<= k`, increment count.
7. Continue until all cells are processed.

## Examples

Input:

```
grid = [[7,6,3],[6,6,1]], k = 18
```

Output:

```
4
```

Input:

```
grid = [[7,2,9],[1,5,0],[2,6,6]], k = 20
```

Output:

```
6
```

## How to use / Run locally

1. Copy the code into your local IDE.
2. Use any compiler (g++, javac, node, python, go).
3. Provide input inside main function.
4. Run and verify output.

## Notes & Optimizations

* Avoid brute force O(n^4)
* Prefix sum reduces it to O(n^2)
* In-place modification saves space
* Works efficiently for large constraints

## Author

* [Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
