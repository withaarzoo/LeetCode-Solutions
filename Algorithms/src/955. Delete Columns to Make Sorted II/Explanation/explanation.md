# Delete Columns to Make Sorted II (LeetCode 955)

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

I am given an array of strings.
All strings have the **same length**.

I can delete **any column** (same index in every string).
After deleting some columns, the final array must be in **lexicographic (dictionary) order**.

My task is to return the **minimum number of columns** I must delete to make this happen.

Important:

* I delete columns globally (from all strings).
* I want the final strings sorted row-wise, not column-wise.

---

## Constraints

* `1 ≤ number of strings ≤ 100`
* `1 ≤ length of each string ≤ 100`
* All strings contain only lowercase English letters
* All strings have equal length

---

## Intuition

At first, I thought this was like the easy version where I just check columns one by one.

But then I realized:

* Deleting one column affects future comparisons.
* Some string pairs may already be sorted, and I should **not disturb them**.

So I needed a way to:

* Remember which adjacent string pairs are already sorted.
* Avoid deleting columns unnecessarily.

That led me to a **greedy + state tracking** idea.

---

## Approach

This is how I solved it step by step:

1. I scan columns **from left to right**.
2. I keep a boolean array called `sorted[]`.

   * `sorted[i] = true` means `strs[i]` and `strs[i+1]` are already in correct order.
3. For each column:

   * I check all **unsorted adjacent pairs**.
   * If any pair violates lexicographic order, I must delete this column.
4. If the column is valid:

   * I update `sorted[]` wherever the order is confirmed.
5. I repeat until all columns are processed.

This guarantees:

* Minimum deletions
* No breaking of already sorted rows

---

## Data Structures Used

* Boolean array (`sorted[]`)
  Used to track which adjacent rows are already sorted.

---

## Operations & Behavior Summary

* Compare only unsorted row pairs
* Greedily delete columns only when necessary
* Lock sorted pairs permanently
* Single pass through columns

---

## Complexity

**Time Complexity:**
`O(n × m)`

* `n` = number of strings
* `m` = length of each string

**Space Complexity:**
`O(n)`

* Only one boolean array is used

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int minDeletionSize(vector<string>& strs) {
        int n = strs.size();
        int m = strs[0].size();
        vector<bool> sorted(n - 1, false);
        int deletions = 0;

        for (int col = 0; col < m; col++) {
            bool bad = false;

            for (int row = 0; row < n - 1; row++) {
                if (!sorted[row] && strs[row][col] > strs[row + 1][col]) {
                    bad = true;
                    break;
                }
            }

            if (bad) {
                deletions++;
                continue;
            }

            for (int row = 0; row < n - 1; row++) {
                if (!sorted[row] && strs[row][col] < strs[row + 1][col]) {
                    sorted[row] = true;
                }
            }
        }
        return deletions;
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
        boolean[] sorted = new boolean[n - 1];
        int deletions = 0;

        for (int col = 0; col < m; col++) {
            boolean bad = false;

            for (int row = 0; row < n - 1; row++) {
                if (!sorted[row] && strs[row].charAt(col) > strs[row + 1].charAt(col)) {
                    bad = true;
                    break;
                }
            }

            if (bad) {
                deletions++;
                continue;
            }

            for (int row = 0; row < n - 1; row++) {
                if (!sorted[row] && strs[row].charAt(col) < strs[row + 1].charAt(col)) {
                    sorted[row] = true;
                }
            }
        }
        return deletions;
    }
}
```

---

### JavaScript

```javascript
var minDeletionSize = function(strs) {
    const n = strs.length;
    const m = strs[0].length;
    const sorted = Array(n - 1).fill(false);
    let deletions = 0;

    for (let col = 0; col < m; col++) {
        let bad = false;

        for (let row = 0; row < n - 1; row++) {
            if (!sorted[row] && strs[row][col] > strs[row + 1][col]) {
                bad = true;
                break;
            }
        }

        if (bad) {
            deletions++;
            continue;
        }

        for (let row = 0; row < n - 1; row++) {
            if (!sorted[row] && strs[row][col] < strs[row + 1][col]) {
                sorted[row] = true;
            }
        }
    }
    return deletions;
};
```

---

### Python3

```python
class Solution:
    def minDeletionSize(self, strs):
        n = len(strs)
        m = len(strs[0])
        sorted_rows = [False] * (n - 1)
        deletions = 0

        for col in range(m):
            bad = False

            for row in range(n - 1):
                if not sorted_rows[row] and strs[row][col] > strs[row + 1][col]:
                    bad = True
                    break

            if bad:
                deletions += 1
                continue

            for row in range(n - 1):
                if not sorted_rows[row] and strs[row][col] < strs[row + 1][col]:
                    sorted_rows[row] = True

        return deletions
```

---

### Go

```go
func minDeletionSize(strs []string) int {
 n := len(strs)
 m := len(strs[0])
 sorted := make([]bool, n-1)
 deletions := 0

 for col := 0; col < m; col++ {
  bad := false

  for row := 0; row < n-1; row++ {
   if !sorted[row] && strs[row][col] > strs[row+1][col] {
    bad = true
    break
   }
  }

  if bad {
   deletions++
   continue
  }

  for row := 0; row < n-1; row++ {
   if !sorted[row] && strs[row][col] < strs[row+1][col] {
    sorted[row] = true
   }
  }
 }
 return deletions
}
```

---

## Step-by-step Detailed Explanation

1. I compare columns from left to right.
2. I skip row pairs that are already sorted.
3. If a column breaks order for any unsorted pair, I delete it.
4. Otherwise, I lock sorted pairs.
5. Once locked, those rows never need comparison again.
6. The count of deleted columns is my answer.

---

## Examples

**Input**

```
["ca", "bb", "ac"]
```

**Output**

```
1
```

Explanation:
Deleting the first column makes the array sorted.

---

## How to Use / Run Locally

1. Copy the solution in your preferred language.
2. Paste it into:

   * LeetCode editor, or
   * Local compiler / IDE.
3. Run with custom test cases.

---

## Notes & Optimizations

* This greedy solution is optimal.
* No sorting or extra string creation is needed.
* Works efficiently within given constraints.
* Very common **FAANG interview pattern**.

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
