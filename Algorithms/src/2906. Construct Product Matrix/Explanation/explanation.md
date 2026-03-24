# Problem Title

1. Construct Product Matrix

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

Given a 2D grid, I need to construct a new matrix where each cell contains the product of all elements in the grid except itself, modulo 12345.

---

## Constraints

* 1 <= n, m <= 10^5
* 2 <= n * m <= 10^5
* 1 <= grid[i][j] <= 10^9

---

## Intuition

I thought of converting this 2D problem into a 1D product-except-self problem.

Instead of calculating product for each cell separately (which would be too slow), I realized I can use prefix and suffix products.

---

## Approach

1. Initialize answer matrix with 1.
2. Traverse grid forward and store prefix product.
3. Traverse grid backward and multiply suffix product.
4. Take modulo at each step.

---

## Data Structures Used

* 2D vector / array for result
* Variables for prefix and suffix products

---

## Operations & Behavior Summary

* Forward pass builds prefix products
* Backward pass builds suffix products
* Combine both to get final result

---

## Complexity

* Time Complexity: O(n * m)
* Space Complexity: O(1) extra space

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    vector<vector<int>> constructProductMatrix(vector<vector<int>>& grid) {
        const int MOD = 12345;
        int n = grid.size(), m = grid[0].size();
        vector<vector<int>> ans(n, vector<int>(m, 1));

        long long prefix = 1;
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < m; j++) {
                ans[i][j] = prefix;
                prefix = (prefix * grid[i][j]) % MOD;
            }
        }

        long long suffix = 1;
        for (int i = n - 1; i >= 0; i--) {
            for (int j = m - 1; j >= 0; j--) {
                ans[i][j] = (ans[i][j] * suffix) % MOD;
                suffix = (suffix * grid[i][j]) % MOD;
            }
        }

        return ans;
    }
};
```

### Java

```java
class Solution {
    public int[][] constructProductMatrix(int[][] grid) {
        final int MOD = 12345;
        int n = grid.length, m = grid[0].length;
        int[][] ans = new int[n][m];

        long prefix = 1;
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < m; j++) {
                ans[i][j] = (int) prefix;
                prefix = (prefix * grid[i][j]) % MOD;
            }
        }

        long suffix = 1;
        for (int i = n - 1; i >= 0; i--) {
            for (int j = m - 1; j >= 0; j--) {
                ans[i][j] = (int)((ans[i][j] * suffix) % MOD);
                suffix = (suffix * grid[i][j]) % MOD;
            }
        }

        return ans;
    }
}
```

### JavaScript

```javascript
var constructProductMatrix = function(grid) {
    const MOD = 12345;
    const n = grid.length, m = grid[0].length;
    const ans = Array.from({ length: n }, () => Array(m).fill(1));

    let prefix = 1;
    for (let i = 0; i < n; i++) {
        for (let j = 0; j < m; j++) {
            ans[i][j] = prefix;
            prefix = (prefix * grid[i][j]) % MOD;
        }
    }

    let suffix = 1;
    for (let i = n - 1; i >= 0; i--) {
        for (let j = m - 1; j >= 0; j--) {
            ans[i][j] = (ans[i][j] * suffix) % MOD;
            suffix = (suffix * grid[i][j]) % MOD;
        }
    }

    return ans;
};
```

### Python3

```python
class Solution:
    def constructProductMatrix(self, grid):
        MOD = 12345
        n, m = len(grid), len(grid[0])
        ans = [[1]*m for _ in range(n)]

        prefix = 1
        for i in range(n):
            for j in range(m):
                ans[i][j] = prefix
                prefix = (prefix * grid[i][j]) % MOD

        suffix = 1
        for i in range(n-1, -1, -1):
            for j in range(m-1, -1, -1):
                ans[i][j] = (ans[i][j] * suffix) % MOD
                suffix = (suffix * grid[i][j]) % MOD

        return ans
```

### Go

```go
func constructProductMatrix(grid [][]int) [][]int {
    const MOD int64 = 12345
    n, m := len(grid), len(grid[0])

    ans := make([][]int, n)
    for i := range ans {
        ans[i] = make([]int, m)
    }

    var prefix int64 = 1
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            ans[i][j] = int(prefix)
            prefix = (prefix * int64(grid[i][j])) % MOD
        }
    }

    var suffix int64 = 1
    for i := n-1; i >= 0; i-- {
        for j := m-1; j >= 0; j-- {
            ans[i][j] = int((int64(ans[i][j]) * suffix) % MOD)
            suffix = (suffix * int64(grid[i][j])) % MOD
        }
    }

    return ans
}
```

---

## Step-by-step Detailed Explanation

1. Initialize result matrix with 1.
2. Traverse grid and store prefix product before updating it.
3. Traverse from reverse and multiply suffix product.
4. Final matrix gives required output.

---

## Examples

Input: [[1,2],[3,4]]
Output: [[24,12],[8,6]]

---

## How to use / Run locally

* Copy code into your preferred language compiler.
* Provide input grid.
* Run and verify output.

---

## Notes & Optimizations

* Avoid division because modulo is not prime.
* Works with zeros safely.
* Optimal O(n*m) solution.

---

## Author

* [Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
