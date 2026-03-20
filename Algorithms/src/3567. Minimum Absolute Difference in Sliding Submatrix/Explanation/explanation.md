# Problem Title

Minimum Absolute Difference in Sliding Submatrix

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

I am given an `m x n` integer grid and an integer `k`.

For every possible `k x k` submatrix, I need to compute the **minimum absolute difference between any two distinct values** inside that submatrix.

Finally, I return a 2D result matrix where each cell represents the answer for that submatrix.

---

## Constraints

* `1 <= m, n <= 30`
* `1 <= k <= min(m, n)`
* `-10^5 <= grid[i][j] <= 10^5`

---

## Intuition

I thought about solving this problem one submatrix at a time.

Since the grid size is small, I realized I can afford a brute-force approach per submatrix.

The key idea is:

* If I sort all elements of a submatrix, the minimum difference will always be between two adjacent elements.

So instead of checking all pairs, I only check neighbors after sorting.

---

## Approach

1. Iterate over all valid top-left positions of `k x k` submatrices.
2. For each submatrix:

   * Collect all elements into a list.
   * Sort the list.
   * Traverse the sorted list and compute minimum difference between consecutive distinct values.
3. If all values are equal, return `0`.
4. Store the result in the answer matrix.

---

## Data Structures Used

* Dynamic array / vector / list → to store elements of submatrix
* Sorting → to efficiently compute minimum difference

---

## Operations & Behavior Summary

* Extract `k x k` values
* Sort values
* Compare adjacent values
* Ignore equal values
* Store minimum difference

---

## Complexity

* **Time Complexity:**
  `O((m - k + 1) * (n - k + 1) * k^2 log(k^2))`

  Explanation:

  * Total submatrices: `(m-k+1)*(n-k+1)`
  * Each submatrix has `k^2` elements
  * Sorting takes `O(k^2 log(k^2))`

* **Space Complexity:**
  `O(k^2)`

  Explanation:

  * Temporary storage for submatrix values

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    vector<vector<int>> minAbsDiff(vector<vector<int>>& grid, int k) {
        int m = grid.size(), n = grid[0].size();
        vector<vector<int>> ans(m - k + 1, vector<int>(n - k + 1));

        for (int i = 0; i + k <= m; i++) {
            for (int j = 0; j + k <= n; j++) {
                vector<int> vals;

                for (int r = i; r < i + k; r++)
                    for (int c = j; c < j + k; c++)
                        vals.push_back(grid[r][c]);

                sort(vals.begin(), vals.end());

                int best = INT_MAX;
                for (int x = 1; x < vals.size(); x++) {
                    if (vals[x] != vals[x-1])
                        best = min(best, vals[x] - vals[x-1]);
                }

                ans[i][j] = (best == INT_MAX ? 0 : best);
            }
        }
        return ans;
    }
};
```

### Java

```java
class Solution {
    public int[][] minAbsDiff(int[][] grid, int k) {
        int m = grid.length, n = grid[0].length;
        int[][] ans = new int[m - k + 1][n - k + 1];

        for (int i = 0; i + k <= m; i++) {
            for (int j = 0; j + k <= n; j++) {
                List<Integer> vals = new ArrayList<>();

                for (int r = i; r < i + k; r++)
                    for (int c = j; c < j + k; c++)
                        vals.add(grid[r][c]);

                Collections.sort(vals);

                int best = Integer.MAX_VALUE;
                for (int x = 1; x < vals.size(); x++) {
                    if (!vals.get(x).equals(vals.get(x-1)))
                        best = Math.min(best, vals.get(x) - vals.get(x-1));
                }

                ans[i][j] = (best == Integer.MAX_VALUE ? 0 : best);
            }
        }
        return ans;
    }
}
```

### JavaScript

```javascript
var minAbsDiff = function(grid, k) {
    const m = grid.length, n = grid[0].length;
    const ans = Array.from({ length: m - k + 1 }, () => Array(n - k + 1).fill(0));

    for (let i = 0; i + k <= m; i++) {
        for (let j = 0; j + k <= n; j++) {
            let vals = [];

            for (let r = i; r < i + k; r++)
                for (let c = j; c < j + k; c++)
                    vals.push(grid[r][c]);

            vals.sort((a, b) => a - b);

            let best = Infinity;
            for (let x = 1; x < vals.length; x++) {
                if (vals[x] !== vals[x-1])
                    best = Math.min(best, vals[x] - vals[x-1]);
            }

            ans[i][j] = best === Infinity ? 0 : best;
        }
    }

    return ans;
};
```

### Python3

```python
class Solution:
    def minAbsDiff(self, grid, k):
        m, n = len(grid), len(grid[0])
        ans = [[0]*(n-k+1) for _ in range(m-k+1)]

        for i in range(m-k+1):
            for j in range(n-k+1):
                vals = []

                for r in range(i, i+k):
                    for c in range(j, j+k):
                        vals.append(grid[r][c])

                vals.sort()

                best = float('inf')
                for x in range(1, len(vals)):
                    if vals[x] != vals[x-1]:
                        best = min(best, vals[x] - vals[x-1])

                ans[i][j] = 0 if best == float('inf') else best

        return ans
```

### Go

```go
import "sort"

func minAbsDiff(grid [][]int, k int) [][]int {
    m, n := len(grid), len(grid[0])
    ans := make([][]int, m-k+1)
    for i := range ans {
        ans[i] = make([]int, n-k+1)
    }

    for i := 0; i+k <= m; i++ {
        for j := 0; j+k <= n; j++ {
            vals := []int{}

            for r := i; r < i+k; r++ {
                for c := j; c < j+k; c++ {
                    vals = append(vals, grid[r][c])
                }
            }

            sort.Ints(vals)

            best := int(1<<60)
            for x := 1; x < len(vals); x++ {
                if vals[x] != vals[x-1] {
                    diff := vals[x] - vals[x-1]
                    if diff < best {
                        best = diff
                    }
                }
            }

            if best == int(1<<60) {
                ans[i][j] = 0
            } else {
                ans[i][j] = best
            }
        }
    }

    return ans
}
```

---

## Step-by-step Detailed Explanation

1. I loop over all possible starting points of submatrices.
2. For each `(i, j)`:

   * I collect all `k x k` elements.
3. I sort the values.
4. I check only adjacent elements:

   * Because sorted order guarantees minimum difference is between neighbors.
5. I skip duplicates.
6. If no valid pair exists, I return `0`.

---

## Examples

### Example 1

Input:

```
grid = [[1,8],[3,-2]], k = 2
```

Output:

```
[[2]]
```

### Example 2

Input:

```
grid = [[3,-1]], k = 1
```

Output:

```
[[0,0]]
```

---

## How to use / Run locally

### C++

```bash
g++ solution.cpp -o solution
./solution
```

### Python

```bash
python3 solution.py
```

### Java

```bash
javac Solution.java
java Solution
```

---

## Notes & Optimizations

* Sorting is sufficient because closest elements are neighbors.
* Brute-force is acceptable due to small constraints.
* Can be optimized using balanced BST / multiset, but not necessary.

---

## Author

* [Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
