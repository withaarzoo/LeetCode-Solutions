# 1582. Special Positions in a Binary Matrix

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

Given an m x n binary matrix `mat`, return the number of special positions in the matrix.

A position `(i, j)` is considered special if:

* `mat[i][j] == 1`
* All other elements in row `i` are `0`
* All other elements in column `j` are `0`

In simple words, the value `1` must be the only `1` present in its row and its column.

## Constraints

* `m == mat.length`
* `n == mat[i].length`
* `1 <= m, n <= 100`
* `mat[i][j]` is either `0` or `1`

## Intuition

When I first read the problem, I realized that a position is special only when the value `1` appears exactly once in its row and exactly once in its column.

So instead of repeatedly scanning rows and columns for every cell, I thought of counting how many `1`s exist in each row and each column first.

If a cell `(i, j)` contains `1`, and the count of `1`s in row `i` is `1`, and the count of `1`s in column `j` is also `1`, then that cell must be a special position.

This observation allows me to solve the problem efficiently in two passes over the matrix.

## Approach

Step 1
Determine the number of rows `m` and columns `n`.

Step 2
Create two arrays:

* `rowCount[m]` to store the number of `1`s in each row
* `colCount[n]` to store the number of `1`s in each column

Step 3
Traverse the matrix and count how many `1`s appear in every row and column.

Step 4
Traverse the matrix again.
If a cell `(i, j)` contains `1` and:

* `rowCount[i] == 1`
* `colCount[j] == 1`
  then that cell is a special position.

Step 5
Count such positions and return the result.

## Data Structures Used

* Array or list for row counts
* Array or list for column counts

These help us quickly determine if a row or column contains exactly one `1`.

## Operations & Behavior Summary

1. Scan the matrix once to count row and column `1`s.
2. Scan the matrix again to check valid special positions.
3. Return the total count.

## Complexity

Time Complexity: O(m * n)

The matrix is traversed twice, and each traversal processes every cell.

Space Complexity: O(m + n)

Two arrays are used to store row and column counts.

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int numSpecial(vector<vector<int>>& mat) {
        int m = mat.size();
        int n = mat[0].size();

        vector<int> rowCount(m, 0);
        vector<int> colCount(n, 0);

        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (mat[i][j] == 1) {
                    rowCount[i]++;
                    colCount[j]++;
                }
            }
        }

        int special = 0;

        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (mat[i][j] == 1 && rowCount[i] == 1 && colCount[j] == 1) {
                    special++;
                }
            }
        }

        return special;
    }
};
```

### Java

```java
class Solution {
    public int numSpecial(int[][] mat) {
        int m = mat.length;
        int n = mat[0].length;

        int[] rowCount = new int[m];
        int[] colCount = new int[n];

        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (mat[i][j] == 1) {
                    rowCount[i]++;
                    colCount[j]++;
                }
            }
        }

        int special = 0;

        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (mat[i][j] == 1 && rowCount[i] == 1 && colCount[j] == 1) {
                    special++;
                }
            }
        }

        return special;
    }
}
```

### JavaScript

```javascript
var numSpecial = function(mat) {
    const m = mat.length;
    const n = mat[0].length;

    const rowCount = new Array(m).fill(0);
    const colCount = new Array(n).fill(0);

    for (let i = 0; i < m; i++) {
        for (let j = 0; j < n; j++) {
            if (mat[i][j] === 1) {
                rowCount[i]++;
                colCount[j]++;
            }
        }
    }

    let special = 0;

    for (let i = 0; i < m; i++) {
        for (let j = 0; j < n; j++) {
            if (mat[i][j] === 1 && rowCount[i] === 1 && colCount[j] === 1) {
                special++;
            }
        }
    }

    return special;
};
```

### Python3

```python
class Solution:
    def numSpecial(self, mat):
        m = len(mat)
        n = len(mat[0])

        rowCount = [0] * m
        colCount = [0] * n

        for i in range(m):
            for j in range(n):
                if mat[i][j] == 1:
                    rowCount[i] += 1
                    colCount[j] += 1

        special = 0

        for i in range(m):
            for j in range(n):
                if mat[i][j] == 1 and rowCount[i] == 1 and colCount[j] == 1:
                    special += 1

        return special
```

### Go

```go
func numSpecial(mat [][]int) int {
    m := len(mat)
    n := len(mat[0])

    rowCount := make([]int, m)
    colCount := make([]int, n)

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if mat[i][j] == 1 {
                rowCount[i]++
                colCount[j]++
            }
        }
    }

    special := 0

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if mat[i][j] == 1 && rowCount[i] == 1 && colCount[j] == 1 {
                special++
            }
        }
    }

    return special
}
```

## Step-by-step Detailed Explanation

First we count how many `1`s exist in each row and each column.

For example, if the matrix is:

1 0 0
0 0 1
1 0 0

Row counts become:
1 1 1

Column counts become:
2 0 1

Now we scan the matrix again.

If we find a cell containing `1` where the row count is `1` and the column count is also `1`, then that cell is a valid special position.

This two-pass approach avoids repeatedly scanning rows and columns for every cell, which keeps the solution efficient.

## Examples

Example 1

Input
[[1,0,0],[0,0,1],[1,0,0]]

Output
1

Example 2

Input
[[1,0,0],[0,1,0],[0,0,1]]

Output
3

## How to use / Run locally

1. Copy the solution in your preferred programming language.
2. Paste it into your LeetCode editor or local IDE.
3. Compile and run the program.
4. Provide the matrix as input.

## Notes & Optimizations

* The algorithm runs in linear time relative to the matrix size.
* Only two extra arrays are used for counting rows and columns.
* This avoids unnecessary repeated scans of rows and columns.

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
