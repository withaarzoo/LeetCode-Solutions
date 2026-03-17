# 1727. Largest Submatrix With Rearrangements

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

---

## Problem Summary

Given a binary matrix of size m x n, we are allowed to rearrange the columns of the matrix in any order.

The goal is to find the area of the largest submatrix consisting entirely of 1s after optimally rearranging the columns.

---

## Constraints

* m == matrix.length
* n == matrix[i].length
* 1 <= m * n <= 10^5
* matrix[i][j] is either 0 or 1

---

## Intuition

I thought about how to maximize a rectangle of 1s. Since I can rearrange columns, I am not restricted by column positions.

So instead of worrying about original positions, I focused on heights of consecutive 1s.

If I treat each row as a histogram, then I can rearrange columns such that taller heights come first. This helps me form the largest possible rectangle.

---

## Approach

1. Traverse the matrix row by row.
2. Convert each cell into height of consecutive 1s above it.
3. For each row:

   * Sort heights in descending order.
   * For each column index j:

     * width = j + 1
     * area = height[j] * width
4. Track maximum area.

---

## Data Structures Used

* Matrix (2D array) for storing heights
* Temporary array for sorting each row

---

## Operations & Behavior Summary

* Build heights using previous row
* Sort each row (column rearrangement simulation)
* Compute area using width and minimum height

---

## Complexity

* Time Complexity: O(m * n log n)

  * Sorting each row takes O(n log n)
* Space Complexity: O(1)

  * In-place modification (excluding sorting copy)

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int largestSubmatrix(vector<vector<int>>& matrix) {
        int m = matrix.size(), n = matrix[0].size();

        for (int i = 1; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (matrix[i][j] == 1) {
                    matrix[i][j] += matrix[i - 1][j];
                }
            }
        }

        int maxArea = 0;

        for (int i = 0; i < m; i++) {
            vector<int> row = matrix[i];
            sort(row.begin(), row.end(), greater<int>());

            for (int j = 0; j < n; j++) {
                maxArea = max(maxArea, row[j] * (j + 1));
            }
        }

        return maxArea;
    }
};
```

### Java

```java
class Solution {
    public int largestSubmatrix(int[][] matrix) {
        int m = matrix.length, n = matrix[0].length;

        for (int i = 1; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (matrix[i][j] == 1) {
                    matrix[i][j] += matrix[i - 1][j];
                }
            }
        }

        int maxArea = 0;

        for (int i = 0; i < m; i++) {
            int[] row = matrix[i].clone();
            Arrays.sort(row);

            for (int j = 0; j < n; j++) {
                int height = row[n - 1 - j];
                maxArea = Math.max(maxArea, height * (j + 1));
            }
        }

        return maxArea;
    }
}
```

### JavaScript

```javascript
var largestSubmatrix = function(matrix) {
    let m = matrix.length, n = matrix[0].length;

    for (let i = 1; i < m; i++) {
        for (let j = 0; j < n; j++) {
            if (matrix[i][j] === 1) {
                matrix[i][j] += matrix[i - 1][j];
            }
        }
    }

    let maxArea = 0;

    for (let i = 0; i < m; i++) {
        let row = [...matrix[i]].sort((a, b) => b - a);

        for (let j = 0; j < n; j++) {
            maxArea = Math.max(maxArea, row[j] * (j + 1));
        }
    }

    return maxArea;
};
```

### Python3

```python
class Solution:
    def largestSubmatrix(self, matrix):
        m, n = len(matrix), len(matrix[0])

        for i in range(1, m):
            for j in range(n):
                if matrix[i][j] == 1:
                    matrix[i][j] += matrix[i-1][j]

        max_area = 0

        for i in range(m):
            row = sorted(matrix[i], reverse=True)

            for j in range(n):
                max_area = max(max_area, row[j] * (j + 1))

        return max_area
```

### Go

```go
import "sort"

func largestSubmatrix(matrix [][]int) int {
    m, n := len(matrix), len(matrix[0])

    for i := 1; i < m; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j] == 1 {
                matrix[i][j] += matrix[i-1][j]
            }
        }
    }

    maxArea := 0

    for i := 0; i < m; i++ {
        row := make([]int, n)
        copy(row, matrix[i])

        sort.Sort(sort.Reverse(sort.IntSlice(row)))

        for j := 0; j < n; j++ {
            if row[j]*(j+1) > maxArea {
                maxArea = row[j] * (j + 1)
            }
        }
    }

    return maxArea
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

1. Start from second row and build heights.
2. If current cell is 1, add value from above.
3. Now each row represents histogram heights.
4. Copy row and sort in descending order.
5. Iterate through sorted row.
6. Calculate width as (index + 1).
7. Compute area = height * width.
8. Track maximum area.

---

## Examples

Input:
matrix = [[0,0,1],[1,1,1],[1,0,1]]
Output:
4

Input:
matrix = [[1,0,1,0,1]]
Output:
3

---

## How to use / Run locally

1. Copy code into your preferred language file.
2. Compile and run.
3. Provide matrix input.

---

## Notes & Optimizations

* Sorting is required because of column rearrangement.
* Greedy approach works since we always try tallest columns first.
* Matrix can be reused to store heights (space optimization).

---

## Author

* [Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
