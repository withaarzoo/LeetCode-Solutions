# Delete Columns to Make Sorted (LeetCode 944)

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

I am given an array of strings where **all strings have the same length**.
These strings are placed one below another to form a **grid**.

Each vertical position forms a **column**.

My task is to **delete the minimum number of columns** so that **every remaining column is sorted lexicographically from top to bottom**.

Finally, I need to return **how many columns I deleted**.

---

## Constraints

* `1 <= number of strings <= 100`
* `1 <= length of each string <= 1000`
* All strings have **equal length**
* Strings contain only **lowercase English letters**

---

## Intuition

When I saw the problem, I imagined the strings as a table.

I thought:

* Each column is independent.
* If a column is already sorted from top to bottom, I keep it.
* If even **one pair** of characters breaks the order, I must delete that column.

So instead of modifying strings,
I just **check column by column and count failures**.

---

## Approach

1. Count total rows and columns.
2. Loop through each column from left to right.
3. For each column:

   * Compare characters row by row.
   * If `upper character > lower character`, order breaks.
4. If order breaks:

   * Increment delete counter.
   * Stop checking that column.
5. Return the delete counter.

This avoids unnecessary comparisons and keeps the solution fast.

---

## Data Structures Used

* No extra data structures
* Only integer counters and loop variables

---

## Operations & Behavior Summary

* Read-only traversal of input
* No sorting
* No string modification
* Early exit for invalid columns
* Independent column validation

---

## Complexity

**Time Complexity:**
`O(n × m)`

* `n` = number of strings (rows)
* `m` = length of each string (columns)

**Space Complexity:**
`O(1)`

* Only constant extra space used

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int minDeletionSize(vector<string>& strs) {
        int rows = strs.size();
        int cols = strs[0].size();
        int deletions = 0;

        for (int c = 0; c < cols; c++) {
            for (int r = 0; r < rows - 1; r++) {
                if (strs[r][c] > strs[r + 1][c]) {
                    deletions++;
                    break;
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
        int rows = strs.length;
        int cols = strs[0].length();
        int deletions = 0;

        for (int c = 0; c < cols; c++) {
            for (int r = 0; r < rows - 1; r++) {
                if (strs[r].charAt(c) > strs[r + 1].charAt(c)) {
                    deletions++;
                    break;
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
    let rows = strs.length;
    let cols = strs[0].length;
    let deletions = 0;

    for (let c = 0; c < cols; c++) {
        for (let r = 0; r < rows - 1; r++) {
            if (strs[r][c] > strs[r + 1][c]) {
                deletions++;
                break;
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
        rows = len(strs)
        cols = len(strs[0])
        deletions = 0

        for c in range(cols):
            for r in range(rows - 1):
                if strs[r][c] > strs[r + 1][c]:
                    deletions += 1
                    break

        return deletions
```

---

### Go

```go
func minDeletionSize(strs []string) int {
    rows := len(strs)
    cols := len(strs[0])
    deletions := 0

    for c := 0; c < cols; c++ {
        for r := 0; r < rows-1; r++ {
            if strs[r][c] > strs[r+1][c] {
                deletions++
                break
            }
        }
    }
    return deletions
}
```

---

## Step-by-step Detailed Explanation

Let’s take this input:

```
["cba",
 "daf",
 "ghi"]
```

### Column 0

```
c
d
g
```

Sorted ✅ → keep it

### Column 1

```
b
a
h
```

`b > a` ❌ → delete this column

### Column 2

```
a
f
i
```

Sorted ✅ → keep it

### Final Answer

```
Deleted columns = 1
```

I repeat this exact logic for every column.

---

## Examples

**Input:**
`["zyx","wvu","tsr"]`

**Output:**
`3`

**Reason:**
All columns are in decreasing order.

---

## How to Use / Run Locally

1. Clone the repository
2. Open the file in your preferred language
3. Run using:

   * `g++` for C++
   * `javac` for Java
   * `node` for JavaScript
   * `python` for Python
   * `go run` for Go

---

## Notes & Optimizations

* Early break improves performance
* No sorting required
* Works efficiently even for maximum constraints
* Ideal for interviews and competitive programming

---

## Author

**Md Aarzoo Islam**
🔗 [https://bento.me/withaarzoo](https://bento.me/withaarzoo)
