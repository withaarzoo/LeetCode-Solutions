# Problem Title

1. Count Submatrices With Equal Frequency of X and Y

---

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

Given a 2D grid where each cell contains 'X', 'Y', or '.', count the number of submatrices starting from (0,0) such that:

* The number of 'X' and 'Y' are equal
* The submatrix contains at least one 'X'

---

## Constraints

* 1 <= grid.length, grid[i].length <= 1000
* grid[i][j] is either 'X', 'Y', or '.'

---

## Intuition

I thought brute force is not possible because checking all submatrices would take O(n^2 * m^2).

So I converted the problem into numbers:

* 'X' → +1
* 'Y' → -1
* '.' → 0

Now the problem becomes finding submatrices where sum = 0 and at least one X exists.

Then I realized that prefix sum can help me compute submatrix sums efficiently.

---

## Approach

1. Convert characters to numeric values
2. Maintain 2D prefix sum
3. Maintain another prefix array to count 'X'
4. For every cell (i, j), compute:

   * sum from (0,0) to (i,j)
   * count of X in that region
5. If:

   * sum == 0
   * countX > 0
     then increment answer
6. Use rolling array to optimize space

---

## Data Structures Used

* 2D prefix sum (optimized to 2 rows)
* Integer arrays for:

  * sum
  * count of X

---

## Operations & Behavior Summary

* Each cell contributes +1, -1, or 0
* Prefix sum allows O(1) calculation per cell
* Rolling arrays reduce memory usage

---

## Complexity

* Time Complexity: O(n * m)
* Space Complexity: O(m)

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int numberOfSubmatrices(vector<vector<char>>& grid) {
        int n = grid.size(), m = grid[0].size();
        vector<vector<int>> sum(2, vector<int>(m + 1, 0));
        vector<vector<int>> countX(2, vector<int>(m + 1, 0));
        int ans = 0;

        for (int i = 0; i < n; i++) {
            int cur = i % 2, prev = 1 - cur;
            for (int j = 0; j < m; j++) {
                int val = (grid[i][j] == 'X') ? 1 : (grid[i][j] == 'Y' ? -1 : 0);
                int isX = (grid[i][j] == 'X');

                sum[cur][j+1] = val + sum[cur][j] + sum[prev][j+1] - sum[prev][j];
                countX[cur][j+1] = isX + countX[cur][j] + countX[prev][j+1] - countX[prev][j];

                if (sum[cur][j+1] == 0 && countX[cur][j+1] > 0)
                    ans++;
            }
        }
        return ans;
    }
};
```

### Java

```java
class Solution {
    public int numberOfSubmatrices(char[][] grid) {
        int n = grid.length, m = grid[0].length;
        int[][] sum = new int[2][m+1];
        int[][] countX = new int[2][m+1];
        int ans = 0;

        for (int i = 0; i < n; i++) {
            int cur = i % 2, prev = 1 - cur;
            for (int j = 0; j < m; j++) {
                int val = grid[i][j]=='X'?1:(grid[i][j]=='Y'?-1:0);
                int isX = grid[i][j]=='X'?1:0;

                sum[cur][j+1] = val + sum[cur][j] + sum[prev][j+1] - sum[prev][j];
                countX[cur][j+1] = isX + countX[cur][j] + countX[prev][j+1] - countX[prev][j];

                if (sum[cur][j+1]==0 && countX[cur][j+1]>0) ans++;
            }
        }
        return ans;
    }
}
```

### JavaScript

```javascript
var numberOfSubmatrices = function(grid) {
    let n = grid.length, m = grid[0].length;
    let sum = Array.from({length:2},()=>Array(m+1).fill(0));
    let countX = Array.from({length:2},()=>Array(m+1).fill(0));
    let ans = 0;

    for(let i=0;i<n;i++){
        let cur=i%2, prev=1-cur;
        for(let j=0;j<m;j++){
            let val = grid[i][j]=='X'?1:(grid[i][j]=='Y'?-1:0);
            let isX = grid[i][j]=='X'?1:0;

            sum[cur][j+1] = val + sum[cur][j] + sum[prev][j+1] - sum[prev][j];
            countX[cur][j+1] = isX + countX[cur][j] + countX[prev][j+1] - countX[prev][j];

            if(sum[cur][j+1]==0 && countX[cur][j+1]>0) ans++;
        }
    }
    return ans;
};
```

### Python3

```python
class Solution:
    def numberOfSubmatrices(self, grid):
        n, m = len(grid), len(grid[0])
        sum_arr = [[0]*(m+1) for _ in range(2)]
        countX = [[0]*(m+1) for _ in range(2)]
        ans = 0

        for i in range(n):
            cur, prev = i%2, 1-(i%2)
            for j in range(m):
                val = 1 if grid[i][j]=='X' else (-1 if grid[i][j]=='Y' else 0)
                isX = 1 if grid[i][j]=='X' else 0

                sum_arr[cur][j+1] = val + sum_arr[cur][j] + sum_arr[prev][j+1] - sum_arr[prev][j]
                countX[cur][j+1] = isX + countX[cur][j] + countX[prev][j+1] - countX[prev][j]

                if sum_arr[cur][j+1]==0 and countX[cur][j+1]>0:
                    ans += 1
        return ans
```

### Go

```go
func numberOfSubmatrices(grid [][]byte) int {
    n, m := len(grid), len(grid[0])
    sum := make([][]int,2)
    countX := make([][]int,2)
    for i:=0;i<2;i++{
        sum[i]=make([]int,m+1)
        countX[i]=make([]int,m+1)
    }
    ans:=0

    for i:=0;i<n;i++{
        cur:=i%2
        prev:=1-cur
        for j:=0;j<m;j++{
            val:=0
            if grid[i][j]=='X'{val=1}else if grid[i][j]=='Y'{val=-1}
            isX:=0
            if grid[i][j]=='X'{isX=1}

            sum[cur][j+1]=val+sum[cur][j]+sum[prev][j+1]-sum[prev][j]
            countX[cur][j+1]=isX+countX[cur][j]+countX[prev][j+1]-countX[prev][j]

            if sum[cur][j+1]==0 && countX[cur][j+1]>0{
                ans++
            }
        }
    }
    return ans
}
```

---

## Step-by-step Detailed Explanation

* Convert grid values to +1, -1, 0
* Build prefix sum using 2D formula
* Maintain count of X separately
* Check condition at each cell
* Use rolling arrays to save space

---

## Examples

Input:
[["X","Y","."],["Y",".","."]]

Output:
3

---

## How to use / Run locally

1. Copy code into your local environment
2. Compile using your language compiler
3. Run with custom test cases

---

## Notes & Optimizations

* Rolling arrays reduce space from O(n*m) to O(m)
* Prefix sum avoids recomputation
* Works efficiently for large grids

---

## Author

* [Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
