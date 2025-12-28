# Problem Title

**1351. Count Negative Numbers in a Sorted Matrix**

---

## 📑 Table of Contents

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

## 🧩 Problem Summary

I am given a matrix of size `m x n`.
Each **row** and **column** of the matrix is sorted in **non-increasing order**.

My task is to **count how many negative numbers** are present in the matrix and return that count.

---

## 📏 Constraints

* `m == grid.length`
* `n == grid[i].length`
* `1 ≤ m, n ≤ 100`
* `-100 ≤ grid[i][j] ≤ 100`

---

## 💡 Intuition

When I first looked at the problem, I noticed something important.
The matrix is sorted both **row-wise** and **column-wise**.

This means:

* Larger numbers are at the **top-left**
* Smaller (negative) numbers are towards the **bottom-right**

So I thought:

> “Why should I check every element if the order itself can help me skip many checks?”

That’s when I decided to use a **smart traversal approach** instead of brute force.

---

## 🚀 Approach

1. I start from the **top-right corner** of the matrix.
2. If the current value is **negative**:

   * Then all values **below it in the same column** are also negative.
   * I directly add `(number of remaining rows)` to my answer.
   * Then I move **left**.
3. If the current value is **not negative**:

   * I move **down**, because negatives must be below.
4. I repeat this until I go outside the matrix.

This way, I only move **left or down**, never backward.

---

## 🧱 Data Structures Used

* 2D Array / Matrix
* Integer variables for pointers and count

No extra data structure is required.

---

## 🔄 Operations & Behavior Summary

* Start position: **Top-right**
* Move directions:

  * Left → when negative found
  * Down → when value is non-negative
* Count negatives in **bulk**, not one by one
* Stops when traversal is complete

---

## ⏱️ Complexity

**Time Complexity:** `O(m + n)`

* `m` = number of rows
* `n` = number of columns
* I move at most `m` steps down and `n` steps left.

**Space Complexity:** `O(1)`

* No extra memory used.

---

## 🌐 Multi-language Solutions

### 🔹 C++

```cpp
class Solution {
public:
    int countNegatives(vector<vector<int>>& grid) {
        int rows = grid.size();
        int cols = grid[0].size();
        int r = 0, c = cols - 1, count = 0;

        while (r < rows && c >= 0) {
            if (grid[r][c] < 0) {
                count += (rows - r);
                c--;
            } else {
                r++;
            }
        }
        return count;
    }
};
```

---

### 🔹 Java

```java
class Solution {
    public int countNegatives(int[][] grid) {
        int rows = grid.length;
        int cols = grid[0].length;
        int r = 0, c = cols - 1, count = 0;

        while (r < rows && c >= 0) {
            if (grid[r][c] < 0) {
                count += (rows - r);
                c--;
            } else {
                r++;
            }
        }
        return count;
    }
}
```

---

### 🔹 JavaScript

```javascript
var countNegatives = function(grid) {
    let rows = grid.length;
    let cols = grid[0].length;
    let r = 0, c = cols - 1, count = 0;

    while (r < rows && c >= 0) {
        if (grid[r][c] < 0) {
            count += (rows - r);
            c--;
        } else {
            r++;
        }
    }
    return count;
};
```

---

### 🔹 Python3

```python
class Solution:
    def countNegatives(self, grid):
        rows, cols = len(grid), len(grid[0])
        r, c, count = 0, cols - 1, 0

        while r < rows and c >= 0:
            if grid[r][c] < 0:
                count += rows - r
                c -= 1
            else:
                r += 1
        return count
```

---

### 🔹 Go

```go
func countNegatives(grid [][]int) int {
    rows := len(grid)
    cols := len(grid[0])
    r, c, count := 0, cols-1, 0

    for r < rows && c >= 0 {
        if grid[r][c] < 0 {
            count += rows - r
            c--
        } else {
            r++
        }
    }
    return count
}
```

---

## 🧠 Step-by-step Detailed Explanation

1. I initialize row and column pointers.
2. I start checking from the top-right cell.
3. If the number is negative:

   * I count all negatives below it in one step.
4. If the number is not negative:

   * I move down to find smaller values.
5. I repeat until traversal ends.
6. I return the final count.

---

## 🧪 Examples

### Example 1

```
Input:
[[4,3,2,-1],
 [3,2,1,-1],
 [1,1,-1,-2],
 [-1,-1,-2,-3]]

Output:
8
```

### Example 2

```
Input:
[[3,2],
 [1,0]]

Output:
0
```

---

## ▶️ How to Use / Run Locally

1. Clone the repository
2. Open the file in your preferred language
3. Compile or run using standard commands:

   * C++: `g++ file.cpp && ./a.out`
   * Java: `javac File.java && java File`
   * Python: `python file.py`
   * JavaScript: `node file.js`
   * Go: `go run file.go`

---

## 📝 Notes & Optimizations

* Brute force works but is slower.
* Binary search is another option but less optimal.
* This solution is **best for interviews** due to `O(m + n)` time.
* No extra memory usage.

---

## 👤 Author

* **Md Aarzoo Islam**
  👉 [https://bento.me/withaarzoo](https://bento.me/withaarzoo)
