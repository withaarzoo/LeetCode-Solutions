# Max Dot Product of Two Subsequences (LeetCode 1458)

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

I am given two integer arrays `nums1` and `nums2`.

My task is to select **non-empty subsequences** from both arrays such that:

* Both subsequences have the **same length**
* The **dot product** of the two subsequences is **maximum**

The dot product is calculated as:
`a1*b1 + a2*b2 + ... + ak*bk`

The order of elements must be preserved, but I am allowed to skip elements.

---

## Constraints

* `1 <= nums1.length, nums2.length <= 500`
* `-1000 <= nums1[i], nums2[i] <= 1000`
* Subsequence **must be non-empty**

---

## Intuition

When I first saw this problem, I understood that brute force is not possible because the number of subsequences grows exponentially.

Since:

* Order matters
* I can skip elements
* I must compare two arrays

This problem strongly matches the pattern of **Dynamic Programming**, similar to Longest Common Subsequence.

So I decided to use DP where each state represents
the **best dot product I can form starting from certain indices** in both arrays.

---

## Approach

1. I define a DP table `dp[i][j]`

   * It represents the **maximum dot product** using subsequences from
     `nums1[i…]` and `nums2[j…]`

2. At each position `(i, j)`, I have three choices:

   * Take both elements `nums1[i] * nums2[j]`
   * Skip element from `nums1`
   * Skip element from `nums2`

3. Since subsequence must be non-empty, I never allow an empty result.

   * I initialize DP values with very small negative numbers.

4. I compute the DP table from bottom-right to top-left.

5. The final answer is stored in `dp[0][0]`.

---

## Data Structures Used

* 2D Array (Dynamic Programming Table)

---

## Operations & Behavior Summary

* Multiplication is done only when both elements are chosen
* DP ensures the subsequence is non-empty
* Skipping elements preserves order
* Handles negative values correctly

---

## Complexity

**Time Complexity:**
`O(n × m)`
Where `n` is length of `nums1` and `m` is length of `nums2`

**Space Complexity:**
`O(n × m)`
Because a 2D DP table is used

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int maxDotProduct(vector<int>& nums1, vector<int>& nums2) {
        int n = nums1.size(), m = nums2.size();
        vector<vector<int>> dp(n + 1, vector<int>(m + 1, INT_MIN));

        for (int i = n - 1; i >= 0; i--) {
            for (int j = m - 1; j >= 0; j--) {
                int product = nums1[i] * nums2[j];
                int take = product;
                if (dp[i + 1][j + 1] != INT_MIN)
                    take = max(take, product + dp[i + 1][j + 1]);

                dp[i][j] = max({take, dp[i + 1][j], dp[i][j + 1]});
            }
        }
        return dp[0][0];
    }
};
```

---

### Java

```java
class Solution {
    public int maxDotProduct(int[] nums1, int[] nums2) {
        int n = nums1.length, m = nums2.length;
        int[][] dp = new int[n + 1][m + 1];

        for (int[] row : dp)
            Arrays.fill(row, Integer.MIN_VALUE);

        for (int i = n - 1; i >= 0; i--) {
            for (int j = m - 1; j >= 0; j--) {
                int product = nums1[i] * nums2[j];
                int take = product;
                if (dp[i + 1][j + 1] != Integer.MIN_VALUE)
                    take = Math.max(take, product + dp[i + 1][j + 1]);

                dp[i][j] = Math.max(take, Math.max(dp[i + 1][j], dp[i][j + 1]));
            }
        }
        return dp[0][0];
    }
}
```

---

### JavaScript

```javascript
var maxDotProduct = function(nums1, nums2) {
    const n = nums1.length, m = nums2.length;
    const dp = Array.from({ length: n + 1 }, () => Array(m + 1).fill(-Infinity));

    for (let i = n - 1; i >= 0; i--) {
        for (let j = m - 1; j >= 0; j--) {
            const product = nums1[i] * nums2[j];
            let take = product;
            if (dp[i + 1][j + 1] !== -Infinity)
                take = Math.max(take, product + dp[i + 1][j + 1]);

            dp[i][j] = Math.max(take, dp[i + 1][j], dp[i][j + 1]);
        }
    }
    return dp[0][0];
};
```

---

### Python3

```python
class Solution:
    def maxDotProduct(self, nums1, nums2):
        n, m = len(nums1), len(nums2)
        dp = [[float('-inf')] * (m + 1) for _ in range(n + 1)]

        for i in range(n - 1, -1, -1):
            for j in range(m - 1, -1, -1):
                product = nums1[i] * nums2[j]
                take = product
                if dp[i + 1][j + 1] != float('-inf'):
                    take = max(take, product + dp[i + 1][j + 1])

                dp[i][j] = max(take, dp[i + 1][j], dp[i][j + 1])

        return dp[0][0]
```

---

### Go

```go
func maxDotProduct(nums1 []int, nums2 []int) int {
    n, m := len(nums1), len(nums2)
    negInf := -1 << 60

    dp := make([][]int, n+1)
    for i := range dp {
        dp[i] = make([]int, m+1)
        for j := range dp[i] {
            dp[i][j] = negInf
        }
    }

    for i := n - 1; i >= 0; i-- {
        for j := m - 1; j >= 0; j-- {
            product := nums1[i] * nums2[j]
            take := product
            if dp[i+1][j+1] != negInf {
                take = max(take, product+dp[i+1][j+1])
            }
            dp[i][j] = max(take, max(dp[i+1][j], dp[i][j+1]))
        }
    }
    return dp[0][0]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

---

## Step-by-step Detailed Explanation

1. I start from the end of both arrays.
2. At every position, I try forming a pair.
3. I also consider skipping elements.
4. I store the best result in DP.
5. This guarantees optimal result without breaking order.

---

## Examples

**Input:**
`nums1 = [2,1,-2,5]`
`nums2 = [3,0,-6]`

**Output:**
`18`

**Explanation:**
Subsequences `[2, -2]` and `[3, -6]`
Dot product = `2×3 + (-2)×(-6) = 18`

---

## How to use / Run locally

1. Clone the repository
2. Open the file in your preferred language
3. Compile or run using standard compiler
4. Test with custom inputs

---

## Notes & Optimizations

* This solution safely handles negative numbers
* Guarantees non-empty subsequence
* Space can be optimized to 2 rows if needed
* Very common interview DP pattern

---

## Author

* **Md Aarzoo Islam**
  [https://bento.me/withaarzoo](https://bento.me/withaarzoo)
