# 960. Delete Columns to Make Sorted III

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
* How to Use / Run Locally
* Notes & Optimizations
* Author

---

## Problem Summary

I am given an array of strings where all strings have the same length.
I am allowed to delete **entire columns** (same index from every string).

After deleting some columns, **each individual string (row)** must be sorted in **non-decreasing lexicographical order** from left to right.

My task is to return the **minimum number of columns** I must delete to achieve this.

Important note:

* The array itself does NOT need to be sorted.
* Only each row must be sorted individually.

---

## Constraints

* `1 <= strs.length <= 100`
* `1 <= strs[i].length <= 100`
* All strings have the same length
* Strings contain only lowercase English letters

---

## Intuition

When I first read the problem, I realized something important.

I am not sorting rows relative to each other.
I am only making sure **each row is sorted left to right**.

So I started thinking in terms of **columns instead of rows**.

If I keep two columns `j` and `i` (where `j < i`), then:

* For **every row**, `strs[row][j] <= strs[row][i]` must be true.

This felt very similar to a **Longest Increasing Subsequence (LIS)** problem, but instead of numbers, I am comparing **columns across all rows**.

So the real goal became:

* Keep the **maximum number of valid columns**
* Delete the remaining columns

---

## Approach

I solved this using **Dynamic Programming**.

Step-by-step idea:

1. Let `m` be the number of columns.
2. Create a DP array `dp` where:

   * `dp[i]` = maximum number of columns I can keep ending at column `i`
3. Initially, every column can stand alone, so:

   * `dp[i] = 1`
4. For every pair of columns `(j < i)`:

   * I check all rows
   * If `strs[row][j] <= strs[row][i]` for every row,
     then column `i` can follow column `j`
5. Update:

   * `dp[i] = max(dp[i], dp[j] + 1)`
6. The maximum value in `dp` is the number of columns I can keep.
7. Answer = `total_columns - max_kept_columns`

---

## Data Structures Used

* **Array / Vector** for Dynamic Programming (`dp`)
* No extra complex data structures are needed

---

## Operations & Behavior Summary

* Compare column pairs
* Validate column order across all rows
* Build longest valid column sequence
* Compute minimum deletions

---

## Complexity

**Time Complexity:**
`O(m² × n)`

* `m` = number of columns
* `n` = number of strings
* Every column pair is checked across all rows

**Space Complexity:**
`O(m)`

* Only a DP array of size equal to number of columns

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int minDeletionSize(vector<string>& strs) {
        int n = strs.size();
        int m = strs[0].size();

        vector<int> dp(m, 1);

        for (int i = 0; i < m; i++) {
            for (int j = 0; j < i; j++) {
                bool valid = true;
                for (int r = 0; r < n; r++) {
                    if (strs[r][j] > strs[r][i]) {
                        valid = false;
                        break;
                    }
                }
                if (valid) {
                    dp[i] = max(dp[i], dp[j] + 1);
                }
            }
        }

        return m - *max_element(dp.begin(), dp.end());
    }
};
```

---

### Java

```java
class Solution {
    public int minDeletionSize(String[] strs) {
        int n = strs.length;
        int m = strs[0].length();

        int[] dp = new int[m];
        Arrays.fill(dp, 1);

        for (int i = 0; i < m; i++) {
            for (int j = 0; j < i; j++) {
                boolean valid = true;
                for (int r = 0; r < n; r++) {
                    if (strs[r].charAt(j) > strs[r].charAt(i)) {
                        valid = false;
                        break;
                    }
                }
                if (valid) {
                    dp[i] = Math.max(dp[i], dp[j] + 1);
                }
            }
        }

        int keep = 0;
        for (int v : dp) keep = Math.max(keep, v);
        return m - keep;
    }
}
```

---

### JavaScript

```javascript
var minDeletionSize = function(strs) {
    const n = strs.length;
    const m = strs[0].length;
    const dp = Array(m).fill(1);

    for (let i = 0; i < m; i++) {
        for (let j = 0; j < i; j++) {
            let valid = true;
            for (let r = 0; r < n; r++) {
                if (strs[r][j] > strs[r][i]) {
                    valid = false;
                    break;
                }
            }
            if (valid) {
                dp[i] = Math.max(dp[i], dp[j] + 1);
            }
        }
    }

    return m - Math.max(...dp);
};
```

---

### Python3

```python
class Solution:
    def minDeletionSize(self, strs):
        n = len(strs)
        m = len(strs[0])
        dp = [1] * m

        for i in range(m):
            for j in range(i):
                valid = True
                for r in range(n):
                    if strs[r][j] > strs[r][i]:
                        valid = False
                        break
                if valid:
                    dp[i] = max(dp[i], dp[j] + 1)

        return m - max(dp)
```

---

### Go

```go
func minDeletionSize(strs []string) int {
    n := len(strs)
    m := len(strs[0])

    dp := make([]int, m)
    for i := range dp {
        dp[i] = 1
    }

    for i := 0; i < m; i++ {
        for j := 0; j < i; j++ {
            valid := true
            for r := 0; r < n; r++ {
                if strs[r][j] > strs[r][i] {
                    valid = false
                    break
                }
            }
            if valid && dp[j]+1 > dp[i] {
                dp[i] = dp[j] + 1
            }
        }
    }

    keep := 0
    for _, v := range dp {
        if v > keep {
            keep = v
        }
    }

    return m - keep
}
```

---

## Step-by-step Detailed Explanation

1. Treat each column as a position in a sequence.
2. Try to extend a valid sequence of columns.
3. Two columns are compatible only if **all rows remain sorted**.
4. Use DP to store best result ending at each column.
5. Final answer is total columns minus the longest valid sequence.

---

## Examples

**Input:**
`["babca","bbazb"]`
**Output:**
`3`

**Input:**
`["edcba"]`
**Output:**
`4`

**Input:**
`["ghi","def","abc"]`
**Output:**
`0`

---

## How to Use / Run Locally

1. Clone the repository
2. Choose your language file
3. Compile / run using standard compiler or interpreter
4. Test with custom inputs if needed

---

## Notes & Optimizations

* This is a **classic LIS-style DP problem**
* Greedy solutions do not work here
* Optimized for constraints given
* Very common in interviews for DP + strings

---

## Author

* **Md Aarzoo Islam**
  👉 [https://bento.me/withaarzoo](https://bento.me/withaarzoo)
