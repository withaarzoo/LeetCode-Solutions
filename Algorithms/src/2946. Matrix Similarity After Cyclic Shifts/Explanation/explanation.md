# 2946. Matrix Similarity After Cyclic Shifts

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
* [Step-by-step Detailed Explanation](#step-by-step-detailed-explanation-c-java-javascript-python3-go)
* [Examples](#examples)
* [How to use / Run locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

## Problem Summary

I am given an `m x n` matrix `mat` and an integer `k`.

The matrix rows are 0-indexed.

For each of the `k` steps:

* Every **even-indexed row** (`0, 2, 4, ...`) is shifted **left** by 1.
* Every **odd-indexed row** (`1, 3, 5, ...`) is shifted **right** by 1.

My task is to return `true` if the matrix becomes exactly the same as the original matrix after `k` steps, otherwise return `false`.

## Constraints

* `1 <= mat.length <= 25`
* `1 <= mat[i].length <= 25`
* `1 <= mat[i][j] <= 25`
* `1 <= k <= 50`

## Intuition

I first thought about simulating the shifting process step by step.

But then I noticed something important:

* A row with `n` columns comes back to its original state after `n` cyclic shifts.
* So instead of doing `k` full operations, I only need to care about `k % n` shifts.

That means I do not need to actually rebuild the matrix.
I can simply compare each element with the position where it should land after the shifts.

For even rows, the shift is to the left.
For odd rows, the shift is to the right.

So the main idea is to check whether every value still matches its expected value after applying the circular movement.

## Approach

1. Let `n` be the number of columns.
2. Reduce the number of effective shifts using:

   ```text
   k = k % n
   ```

3. Traverse every cell `(i, j)` in the matrix.
4. For an **even row**:

   * The element at column `j` should move to `(j + k) % n`.
5. For an **odd row**:

   * The element at column `j` should move to `(j - k + n) % n`.
6. Compare the original value with the value at the expected shifted position.
7. If any comparison fails, return `false`.
8. If all comparisons pass, return `true`.

This works because a cyclic shift only changes positions, not values.

## Data Structures Used

I do not need any extra data structure.

* **Input matrix**: used directly for comparison
* **Variables**: only a few integer counters and indices

So the solution stays simple and memory efficient.

## Operations & Behavior Summary

* Even-indexed rows shift left by `k % n`
* Odd-indexed rows shift right by `k % n`
* The matrix is not modified
* I only verify whether the shifted arrangement matches the original one

## Complexity

* **Time Complexity:** `O(m * n)`

  * I visit each cell once.
  * Here, `m` is the number of rows and `n` is the number of columns.

* **Space Complexity:** `O(1)`

  * I use only a constant amount of extra space.

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    bool areSimilar(vector<vector<int>>& mat, int k) {
        int m = mat.size();
        int n = mat[0].size();
        k %= n;

        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                int newCol;

                if (i % 2 == 0) {
                    // Even-indexed row: left shift
                    newCol = (j + k) % n;
                } else {
                    // Odd-indexed row: right shift
                    newCol = (j - k + n) % n;
                }

                if (mat[i][j] != mat[i][newCol]) {
                    return false;
                }
            }
        }

        return true;
    }
};
```

### Java

```java
class Solution {
    public boolean areSimilar(int[][] mat, int k) {
        int m = mat.length;
        int n = mat[0].length;
        k %= n;

        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                int newCol;

                if (i % 2 == 0) {
                    // Even-indexed row: left shift
                    newCol = (j + k) % n;
                } else {
                    // Odd-indexed row: right shift
                    newCol = (j - k + n) % n;
                }

                if (mat[i][j] != mat[i][newCol]) {
                    return false;
                }
            }
        }

        return true;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[][]} mat
 * @param {number} k
 * @return {boolean}
 */
var areSimilar = function(mat, k) {
    const m = mat.length;
    const n = mat[0].length;
    k = k % n;

    for (let i = 0; i < m; i++) {
        for (let j = 0; j < n; j++) {
            let newCol;

            if (i % 2 === 0) {
                // Even-indexed row: left shift
                newCol = (j + k) % n;
            } else {
                // Odd-indexed row: right shift
                newCol = (j - k + n) % n;
            }

            if (mat[i][j] !== mat[i][newCol]) {
                return false;
            }
        }
    }

    return true;
};
```

### Python3

```python
from typing import List

class Solution:
    def areSimilar(self, mat: List[List[int]], k: int) -> bool:
        m, n = len(mat), len(mat[0])
        k %= n

        for i in range(m):
            for j in range(n):
                if i % 2 == 0:
                    # Even-indexed row: left shift
                    new_col = (j + k) % n
                else:
                    # Odd-indexed row: right shift
                    new_col = (j - k + n) % n

                if mat[i][j] != mat[i][new_col]:
                    return False

        return True
```

### Go

```go
func areSimilar(mat [][]int, k int) bool {
    m := len(mat)
    n := len(mat[0])
    k %= n

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            var newCol int

            if i%2 == 0 {
                // Even-indexed row: left shift
                newCol = (j + k) % n
            } else {
                // Odd-indexed row: right shift
                newCol = (j - k + n) % n
            }

            if mat[i][j] != mat[i][newCol] {
                return false
            }
        }
    }

    return true
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

### 1. Read the matrix size

I first find the number of rows and columns.

* `m` = number of rows
* `n` = number of columns

This helps me know how many elements I need to check.

### 2. Reduce the number of shifts

I compute:

```text
k = k % n
```

This is important because a row returns to its original state after `n` cyclic shifts.

So if `k` is larger than `n`, the extra shifts do not matter.

### 3. Visit every cell

I go through every row `i` and every column `j`.

This lets me check the whole matrix without creating a new one.

### 4. Handle even rows

If the row index is even, the row shifts left.

So the value originally at column `j` is expected at:

```text
(j + k) % n
```

The `% n` part keeps the index inside the row.

### 5. Handle odd rows

If the row index is odd, the row shifts right.

So the value originally at column `j` is expected at:

```text
(j - k + n) % n
```

I add `n` before taking modulo to avoid negative indices.

### 6. Compare values

I compare:

* the current value `mat[i][j]`
* the value at the expected shifted position `mat[i][newCol]`

If they are different, the matrix is not similar after the shifts.

### 7. Return the answer

* If one mismatch is found, I return `false` immediately.
* If I finish checking everything, I return `true`.

## Examples

### Example 1

**Input:**

```text
mat = [[1,2,3],[4,5,6],[7,8,9]], k = 4
```

**Output:**

```text
false
```

**Explanation:**
After the shifts, the matrix does not match the original matrix.

### Example 2

**Input:**

```text
mat = [[1,2,1,2],[5,5,5,5],[6,3,6,3]], k = 2
```

**Output:**

```text
true
```

**Explanation:**
After applying the cyclic shifts twice, the matrix becomes the same as the original one.

### Example 3

**Input:**

```text
mat = [[2,2],[2,2]], k = 3
```

**Output:**

```text
true
```

**Explanation:**
All values are the same, so any cyclic shift keeps the matrix unchanged.

## How to use / Run locally

### C++

```bash
g++ -std=c++17 solution.cpp -o solution
./solution
```

### Java

```bash
javac Solution.java
java Solution
```

### JavaScript

```bash
node solution.js
```

### Python3

```bash
python3 solution.py
```

### Go

```bash
go run solution.go
```

## Notes & Optimizations

* I do not simulate all `k` operations.
* I use `k % n` to remove unnecessary repeated work.
* I do not create a temporary matrix.
* This keeps the code clean, fast, and memory efficient.
* The solution is easy to extend to other languages because the logic is based on simple index math.

## Author

* [Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
