# Magic Squares In Grid (LeetCode 840)

## Table of Contents

* [Problem Summary](#problem-summary)
* [Constraints](#constraints)
* [Intuition](#intuition)
* [Approach](#approach)
* [Data Structures Used](#data-structures-used)
* [Operations & Behavior Summary](#operations--behavior-summary)
* [Complexity](#complexity)
* [Multi-language Solutions](#multi-language-solutions)

  * [C++](#c)
  * [Java](#java)
  * [JavaScript](#javascript)
  * [Python3](#python3)
  * [Go](#go)
* [Step-by-step Detailed Explanation](#step-by-step-detailed-explanation)
* [Examples](#examples)
* [How to use / Run locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

---

## Problem Summary

A **3 × 3 magic square** is a grid that:

* Contains **distinct numbers from 1 to 9**
* Every **row, column, and both diagonals** have the **same sum**

You are given a grid of integers.
Your task is to **count how many 3 × 3 magic square subgrids** exist inside it.

Important note:

* The grid may contain numbers up to **15**
* But a valid magic square can only contain **1 to 9**

---

## Constraints

* `1 ≤ rows, cols ≤ 10`
* `0 ≤ grid[i][j] ≤ 15`
* Only **3 × 3** subgrids are checked

---

## Intuition

When I read the problem, I noticed something important.

The problem **does not ask for any size magic square**.
It only asks for **3 × 3 magic squares**.

So I thought:

* Why generate anything?
* I can simply **check every possible 3 × 3 subgrid**.

Then I remembered a key rule:

> In every valid 3 × 3 magic square, the **center value is always 5**.

That one check alone helps reject most invalid cases very fast.

So my idea became:

* Slide a **3 × 3 window**
* Validate it using strict rules
* Count only valid ones

---

## Approach

1. If the grid is smaller than `3 × 3`, return `0`
2. Move a `3 × 3` window across the grid
3. For each window:

   * Check center value is `5`
   * Ensure all numbers are between `1–9`
   * Ensure all numbers are **unique**
   * Check:

     * 3 row sums
     * 3 column sums
     * 2 diagonal sums
4. If all checks pass → count it
5. Return the total count

---

## Data Structures Used

* **Boolean array / Set**
  Used to track numbers `1–9` and ensure uniqueness

No extra memory proportional to input size is used.

---

## Operations & Behavior Summary

| Operation        | Purpose                     |
| ---------------- | --------------------------- |
| Sliding window   | Iterate over all 3×3 blocks |
| Center check     | Fast rejection              |
| Uniqueness check | Ensure numbers 1–9          |
| Sum validation   | Magic square rule           |

---

## Complexity

* **Time Complexity:** `O(m × n)`

  * `m` = rows, `n` = columns
  * Each 3 × 3 block is checked in constant time

* **Space Complexity:** `O(1)`

  * Only fixed-size arrays / sets are used

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int numMagicSquaresInside(vector<vector<int>>& grid) {
        int rows = grid.size(), cols = grid[0].size();
        int count = 0;

        for (int i = 0; i + 2 < rows; i++) {
            for (int j = 0; j + 2 < cols; j++) {
                if (isMagic(grid, i, j)) count++;
            }
        }
        return count;
    }

    bool isMagic(vector<vector<int>>& g, int r, int c) {
        if (g[r+1][c+1] != 5) return false;

        bool seen[10] = {false};
        for (int i = r; i < r+3; i++) {
            for (int j = c; j < c+3; j++) {
                int v = g[i][j];
                if (v < 1 || v > 9 || seen[v]) return false;
                seen[v] = true;
            }
        }

        for (int i = 0; i < 3; i++) {
            if (g[r+i][c] + g[r+i][c+1] + g[r+i][c+2] != 15) return false;
            if (g[r][c+i] + g[r+1][c+i] + g[r+2][c+i] != 15) return false;
        }

        if (g[r][c] + g[r+1][c+1] + g[r+2][c+2] != 15) return false;
        if (g[r][c+2] + g[r+1][c+1] + g[r+2][c] != 15) return false;

        return true;
    }
};
```

---

### Java

```java
class Solution {
    public int numMagicSquaresInside(int[][] grid) {
        int rows = grid.length, cols = grid[0].length;
        int count = 0;

        for (int i = 0; i + 2 < rows; i++) {
            for (int j = 0; j + 2 < cols; j++) {
                if (isMagic(grid, i, j)) count++;
            }
        }
        return count;
    }

    boolean isMagic(int[][] g, int r, int c) {
        if (g[r+1][c+1] != 5) return false;

        boolean[] seen = new boolean[10];
        for (int i = r; i < r+3; i++) {
            for (int j = c; j < c+3; j++) {
                int v = g[i][j];
                if (v < 1 || v > 9 || seen[v]) return false;
                seen[v] = true;
            }
        }

        for (int i = 0; i < 3; i++) {
            if (g[r+i][c] + g[r+i][c+1] + g[r+i][c+2] != 15) return false;
            if (g[r][c+i] + g[r+1][c+i] + g[r+2][c+i] != 15) return false;
        }

        return g[r][c] + g[r+1][c+1] + g[r+2][c+2] == 15 &&
               g[r][c+2] + g[r+1][c+1] + g[r+2][c] == 15;
    }
}
```

---

### JavaScript

```javascript
var numMagicSquaresInside = function(grid) {
    let count = 0;

    for (let i = 0; i + 2 < grid.length; i++) {
        for (let j = 0; j + 2 < grid[0].length; j++) {
            if (isMagic(grid, i, j)) count++;
        }
    }
    return count;
};

function isMagic(g, r, c) {
    if (g[r+1][c+1] !== 5) return false;

    let seen = Array(10).fill(false);
    for (let i = r; i < r+3; i++) {
        for (let j = c; j < c+3; j++) {
            let v = g[i][j];
            if (v < 1 || v > 9 || seen[v]) return false;
            seen[v] = true;
        }
    }

    for (let i = 0; i < 3; i++) {
        if (g[r+i][c] + g[r+i][c+1] + g[r+i][c+2] !== 15) return false;
        if (g[r][c+i] + g[r+1][c+i] + g[r+2][c+i] !== 15) return false;
    }

    return true;
}
```

---

### Python3

```python
class Solution:
    def numMagicSquaresInside(self, grid):
        count = 0
        for i in range(len(grid)-2):
            for j in range(len(grid[0])-2):
                if self.isMagic(grid, i, j):
                    count += 1
        return count

    def isMagic(self, g, r, c):
        if g[r+1][c+1] != 5:
            return False

        seen = set()
        for i in range(r, r+3):
            for j in range(c, c+3):
                v = g[i][j]
                if v < 1 or v > 9 or v in seen:
                    return False
                seen.add(v)

        for i in range(3):
            if sum(g[r+i][c:c+3]) != 15:
                return False
            if g[r][c+i] + g[r+1][c+i] + g[r+2][c+i] != 15:
                return False

        return True
```

---

### Go

```go
func numMagicSquaresInside(grid [][]int) int {
    count := 0
    for i := 0; i+2 < len(grid); i++ {
        for j := 0; j+2 < len(grid[0]); j++ {
            if isMagic(grid, i, j) {
                count++
            }
        }
    }
    return count
}

func isMagic(g [][]int, r, c int) bool {
    if g[r+1][c+1] != 5 {
        return false
    }

    seen := make([]bool, 10)
    for i := r; i < r+3; i++ {
        for j := c; j < c+3; j++ {
            v := g[i][j]
            if v < 1 || v > 9 || seen[v] {
                return false
            }
            seen[v] = true
        }
    }

    for i := 0; i < 3; i++ {
        if g[r+i][c]+g[r+i][c+1]+g[r+i][c+2] != 15 {
            return false
        }
        if g[r][c+i]+g[r+1][c+i]+g[r+2][c+i] != 15 {
            return false
        }
    }

    return true
}
```

---

## Step-by-step Detailed Explanation

1. I scan the grid using a sliding 3 × 3 window
2. I immediately reject blocks whose center is not 5
3. I ensure all values are unique and within 1–9
4. I validate row, column, and diagonal sums
5. If all checks pass, I count the square

---

## Examples

**Input**

```
[[4,3,8,4],
 [9,5,1,9],
 [2,7,6,2]]
```

**Output**

```
1
```

---

## How to use / Run locally

1. Clone the repository
2. Open the file for your language
3. Run using:

   * `g++` for C++
   * `javac` for Java
   * `node` for JavaScript
   * `python3` for Python
   * `go run` for Go

---

## Notes & Optimizations

* Center check (`5`) avoids unnecessary computation
* No extra memory based on grid size
* Interview-friendly and clean logic
* Works within all constraints

---

## Author

* **Md Aarzoo Islam**
  🔗 [https://bento.me/withaarzoo](https://bento.me/withaarzoo)
