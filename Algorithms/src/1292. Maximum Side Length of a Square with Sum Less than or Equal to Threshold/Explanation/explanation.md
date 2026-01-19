# Maximum Side Length of a Square with Sum ≤ Threshold

## Problem Title

**1292. Maximum Side Length of a Square with Sum Less than or Equal to Threshold**

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

I am given a 2D matrix `mat` and an integer `threshold`.

My task is to find the **largest possible square sub-matrix** such that:

* The square has equal height and width
* The **sum of all elements inside the square is ≤ threshold**

If no such square exists, I return **0**.

---

## Constraints

* `1 ≤ m, n ≤ 300`
* `0 ≤ mat[i][j] ≤ 10⁴`
* `0 ≤ threshold ≤ 10⁵`
* Matrix size can be large, so brute force is not efficient

---

## Intuition

When I saw this problem, I understood two things immediately:

1. Calculating the sum of every square again and again will be **too slow**
2. If a square of size `k` works, then **all smaller squares also work**

So I thought:

* I need a way to calculate square sums **fast**
* And I should search for the answer **smartly**

That’s why I decided to use:

* **2D Prefix Sum** to get square sum in O(1)
* **Binary Search** to find the maximum valid square size

---

## Approach

I solved the problem step by step:

1. Build a **prefix sum matrix**
2. Binary search the square size from `0` to `min(rows, cols)`
3. For each size `mid`:

   * Slide a square of that size over the matrix
   * Calculate its sum using prefix sum
4. If any square has sum ≤ threshold:

   * Try bigger size
5. Else:

   * Reduce the size
6. The largest valid size is the answer

---

## Data Structures Used

* 2D Array → Prefix Sum Matrix
* Binary Search variables (`low`, `high`, `mid`)

---

## Operations & Behavior Summary

* Prefix sum allows constant time square sum calculation
* Binary search reduces unnecessary checks
* Efficient even for large matrices (300 × 300)

---

## Complexity

### Time Complexity

**O(m × n × log(min(m, n)))**

* `m × n` → checking all possible square positions
* `log(min(m, n))` → binary search on side length

### Space Complexity

**O(m × n)**

* Extra space used for prefix sum matrix

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int maxSideLength(vector<vector<int>>& mat, int threshold) {
        int m = mat.size(), n = mat[0].size();
        vector<vector<int>> pre(m + 1, vector<int>(n + 1, 0));

        for (int i = 1; i <= m; i++) {
            for (int j = 1; j <= n; j++) {
                pre[i][j] = mat[i - 1][j - 1]
                          + pre[i - 1][j]
                          + pre[i][j - 1]
                          - pre[i - 1][j - 1];
            }
        }

        int left = 0, right = min(m, n), ans = 0;

        while (left <= right) {
            int mid = (left + right) / 2;
            bool found = false;

            for (int i = mid; i <= m && !found; i++) {
                for (int j = mid; j <= n; j++) {
                    int sum = pre[i][j]
                            - pre[i - mid][j]
                            - pre[i][j - mid]
                            + pre[i - mid][j - mid];

                    if (sum <= threshold) {
                        found = true;
                        break;
                    }
                }
            }

            if (found) {
                ans = mid;
                left = mid + 1;
            } else {
                right = mid - 1;
            }
        }
        return ans;
    }
};
```

---

### Java

```java
class Solution {
    public int maxSideLength(int[][] mat, int threshold) {
        int m = mat.length, n = mat[0].length;
        int[][] pre = new int[m + 1][n + 1];

        for (int i = 1; i <= m; i++) {
            for (int j = 1; j <= n; j++) {
                pre[i][j] = mat[i - 1][j - 1]
                          + pre[i - 1][j]
                          + pre[i][j - 1]
                          - pre[i - 1][j - 1];
            }
        }

        int left = 0, right = Math.min(m, n), ans = 0;

        while (left <= right) {
            int mid = (left + right) / 2;
            boolean found = false;

            for (int i = mid; i <= m && !found; i++) {
                for (int j = mid; j <= n; j++) {
                    int sum = pre[i][j]
                            - pre[i - mid][j]
                            - pre[i][j - mid]
                            + pre[i - mid][j - mid];
                    if (sum <= threshold) {
                        found = true;
                        break;
                    }
                }
            }

            if (found) {
                ans = mid;
                left = mid + 1;
            } else {
                right = mid - 1;
            }
        }
        return ans;
    }
}
```

---

### JavaScript

```javascript
var maxSideLength = function(mat, threshold) {
    const m = mat.length, n = mat[0].length;
    const pre = Array.from({ length: m + 1 }, () => Array(n + 1).fill(0));

    for (let i = 1; i <= m; i++) {
        for (let j = 1; j <= n; j++) {
            pre[i][j] = mat[i - 1][j - 1]
                      + pre[i - 1][j]
                      + pre[i][j - 1]
                      - pre[i - 1][j - 1];
        }
    }

    let left = 0, right = Math.min(m, n), ans = 0;

    while (left <= right) {
        const mid = Math.floor((left + right) / 2);
        let found = false;

        for (let i = mid; i <= m && !found; i++) {
            for (let j = mid; j <= n; j++) {
                const sum = pre[i][j]
                          - pre[i - mid][j]
                          - pre[i][j - mid]
                          + pre[i - mid][j - mid];
                if (sum <= threshold) {
                    found = true;
                    break;
                }
            }
        }

        if (found) {
            ans = mid;
            left = mid + 1;
        } else {
            right = mid - 1;
        }
    }
    return ans;
};
```

---

### Python3

```python
class Solution:
    def maxSideLength(self, mat, threshold):
        m, n = len(mat), len(mat[0])
        pre = [[0] * (n + 1) for _ in range(m + 1)]

        for i in range(1, m + 1):
            for j in range(1, n + 1):
                pre[i][j] = mat[i - 1][j - 1] \
                            + pre[i - 1][j] \
                            + pre[i][j - 1] \
                            - pre[i - 1][j - 1]

        left, right, ans = 0, min(m, n), 0

        while left <= right:
            mid = (left + right) // 2
            found = False

            for i in range(mid, m + 1):
                for j in range(mid, n + 1):
                    if pre[i][j] - pre[i-mid][j] - pre[i][j-mid] + pre[i-mid][j-mid] <= threshold:
                        found = True
                        break
                if found:
                    break

            if found:
                ans = mid
                left = mid + 1
            else:
                right = mid - 1

        return ans
```

---

### Go

```go
func maxSideLength(mat [][]int, threshold int) int {
    m, n := len(mat), len(mat[0])
    pre := make([][]int, m+1)

    for i := range pre {
        pre[i] = make([]int, n+1)
    }

    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            pre[i][j] = mat[i-1][j-1] +
                pre[i-1][j] +
                pre[i][j-1] -
                pre[i-1][j-1]
        }
    }

    left, right, ans := 0, min(m, n), 0

    for left <= right {
        mid := (left + right) / 2
        found := false

        for i := mid; i <= m && !found; i++ {
            for j := mid; j <= n; j++ {
                sum := pre[i][j] - pre[i-mid][j] - pre[i][j-mid] + pre[i-mid][j-mid]
                if sum <= threshold {
                    found = true
                    break
                }
            }
        }

        if found {
            ans = mid
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return ans
}
```

---

## Step-by-step Detailed Explanation

1. Prefix sum stores cumulative sum up to each cell
2. Square sum is calculated using 4 prefix values
3. Binary search reduces time from brute force
4. We always try to expand square size if valid
5. Final result is maximum valid side length

---

## Examples

**Input**

```bash
mat = [[1,1,3,2,4,3,2],
       [1,1,3,2,4,3,2],
       [1,1,3,2,4,3,2]]
threshold = 4
```

**Output**

```bash
2
```

---

## How to use / Run locally

1. Copy the solution for your language
2. Paste into LeetCode / IDE
3. Run with custom inputs
4. No external libraries required

---

## Notes & Optimizations

* Prefix sum is mandatory for performance
* Binary search avoids checking all sizes
* Works efficiently within constraints

---

## Author

**Md Aarzoo Islam**
🔗 [https://bento.me/withaarzoo](https://bento.me/withaarzoo)
