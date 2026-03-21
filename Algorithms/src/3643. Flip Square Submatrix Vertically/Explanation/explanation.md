# Problem Title

Flip Square Submatrix Vertically

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

I am given an m x n matrix `grid` and three integers `x`, `y`, and `k`.

* `(x, y)` represents the top-left corner of a square submatrix.
* `k` represents the size of that square.

My task is to flip this `k x k` submatrix vertically.

Vertical flip means reversing the order of rows inside that square.

Finally, I return the updated matrix.

## Constraints

* m == grid.length
* n == grid[i].length
* 1 <= m, n <= 50
* 1 <= grid[i][j] <= 100
* 0 <= x < m
* 0 <= y < n
* 1 <= k <= min(m - x, n - y)

## Intuition

I thought about what vertical flipping actually means.

If I take a square and flip it vertically, the rows get reversed:

* First row becomes last
* Second row becomes second last

So instead of moving everything, I can just swap rows.

And I only need to do this inside the selected square.

## Approach

1. I iterate only over half of the rows of the square (k / 2).
2. For each row index `i`, I identify:

   * Top row: `x + i`
   * Bottom row: `x + k - 1 - i`
3. Then I swap elements column-wise inside the square:

   * Columns go from `y` to `y + k - 1`
4. I swap each corresponding element of the two rows.

This efficiently flips the submatrix in-place.

## Data Structures Used

* 2D array (matrix)
* No extra data structures required

## Operations & Behavior Summary

* In-place row swapping
* Controlled iteration over a fixed square region
* No additional memory usage

## Complexity

* Time Complexity: O(k^2)

  * I process k/2 rows and k columns
* Space Complexity: O(1)

  * No extra space is used

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    vector<vector<int>> reverseSubmatrix(vector<vector<int>>& grid, int x, int y, int k) {
        for (int i = 0; i < k / 2; i++) {
            int top = x + i;
            int bottom = x + k - 1 - i;

            for (int j = 0; j < k; j++) {
                swap(grid[top][y + j], grid[bottom][y + j]);
            }
        }
        return grid;
    }
};
```

### Java

```java
class Solution {
    public int[][] reverseSubmatrix(int[][] grid, int x, int y, int k) {
        for (int i = 0; i < k / 2; i++) {
            int top = x + i;
            int bottom = x + k - 1 - i;

            for (int j = 0; j < k; j++) {
                int temp = grid[top][y + j];
                grid[top][y + j] = grid[bottom][y + j];
                grid[bottom][y + j] = temp;
            }
        }
        return grid;
    }
}
```

### JavaScript

```javascript
var reverseSubmatrix = function(grid, x, y, k) {
    for (let i = 0; i < Math.floor(k / 2); i++) {
        let top = x + i;
        let bottom = x + k - 1 - i;

        for (let j = 0; j < k; j++) {
            let temp = grid[top][y + j];
            grid[top][y + j] = grid[bottom][y + j];
            grid[bottom][y + j] = temp;
        }
    }
    return grid;
};
```

### Python3

```python
class Solution:
    def reverseSubmatrix(self, grid, x, y, k):
        for i in range(k // 2):
            top = x + i
            bottom = x + k - 1 - i

            for j in range(k):
                grid[top][y + j], grid[bottom][y + j] = grid[bottom][y + j], grid[top][y + j]

        return grid
```

### Go

```go
func reverseSubmatrix(grid [][]int, x int, y int, k int) [][]int {
    for i := 0; i < k/2; i++ {
        top := x + i
        bottom := x + k - 1 - i

        for j := 0; j < k; j++ {
            grid[top][y+j], grid[bottom][y+j] = grid[bottom][y+j], grid[top][y+j]
        }
    }
    return grid
}
```

## Step-by-step Detailed Explanation

1. I loop from `0` to `k/2` because each swap handles two rows.
2. I calculate the top and bottom rows dynamically.
3. Then I loop through all columns inside the square.
4. I swap corresponding elements to reverse the rows.
5. This effectively flips the submatrix vertically.

## Examples

Example 1:
Input:

```
grid = [[1,2,3,4],[5,6,7,8],[9,10,11,12],[13,14,15,16]]
x = 1, y = 0, k = 3
```

Output:

```
[[1,2,3,4],[13,14,15,8],[9,10,11,12],[5,6,7,16]]
```

Example 2:
Input:

```
grid = [[3,4,2,3],[2,3,4,2]]
x = 0, y = 2, k = 2
```

Output:

```
[[3,4,4,2],[2,3,2,3]]
```

## How to use / Run locally

1. Clone the repository
2. Copy the solution into your preferred language file
3. Compile and run:

C++:

```
g++ file.cpp -o output
./output
```

Java:

```
javac Solution.java
java Solution
```

Python:

```
python3 file.py
```

JavaScript:

```
node file.js
```

Go:

```
go run file.go
```

## Notes & Optimizations

* No need for extra matrix or copying
* In-place swapping keeps memory optimal
* Only process the required square region

## Author

* [Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
