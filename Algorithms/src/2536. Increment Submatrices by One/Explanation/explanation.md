# Problem Title

**2536. Increment Submatrices by One**

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
* [Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)](#step-by-step-detailed-explanation-c-java-javascript-python3-go)
* [Examples](#examples)
* [How to use / Run locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

---

## Problem Summary

I am given a positive integer `n`, representing an `n x n` zero-initialized matrix (0-indexed). I am also given a list of queries, where each query is `[r1, c1, r2, c2]`. For each query, I must add `1` to every element inside the submatrix whose top-left corner is `(r1, c1)` and bottom-right corner is `(r2, c2)`. After applying all queries, I need to return the final `n x n` matrix.

The naive approach (increment every cell for every query) is slow when there are many queries. I use a 2D difference array and then convert it to the final matrix via 2D prefix sums for an efficient solution.

---

## Constraints

* `1 <= n <= 500`
* `1 <= queries.length <= 10^4`
* `0 <= r1 <= r2 < n`
* `0 <= c1 <= c2 < n`

---

## Intuition

I thought about the naive method: for each query, loop over all cells in the rectangle and increment — that’d be O(q * area) and too slow for many queries. I remembered a trick: using a **2D difference array**. With 4 corner updates per query, each query becomes O(1). After all queries, computing a 2D prefix sum reconstructs the resulting matrix in O(n²). This is fast and memory-friendly for the given constraints.

---

## Approach

1. Create a `(n+1) x (n+1)` difference matrix `diff` initialized to zero. The extra row/column makes boundary updates (`r2+1`, `c2+1`) safe without if-checks.
2. For a query `[r1, c1, r2, c2]`, perform four updates:

   * `diff[r1][c1] += 1`
   * `diff[r1][c2+1] -= 1`
   * `diff[r2+1][c1] -= 1`
   * `diff[r2+1][c2+1] += 1`
3. After all queries, convert `diff` into the final `n x n` matrix by computing a 2D prefix sum:

   * `diff[i][j] += diff[i-1][j] + diff[i][j-1] - diff[i-1][j-1]`
   * The value `diff[i][j]` (for `0 <= i,j < n`) is the answer cell.
4. Return the first `n x n` portion of `diff` (or copy values into a new `res` matrix).

---

## Data Structures Used

* 2D array `diff` of size `(n+1) x (n+1)` (integers).
* Output 2D array `res` of size `n x n` (integers).

---

## Operations & Behavior Summary

* Each query results in **4 constant-time updates** (top-left add, top-right subtract, bottom-left subtract, bottom-right add).
* Final assembly uses the **2D prefix-sum recurrence** to compute final values for each cell using previously computed neighbors.
* No per-query cell iteration is required, so performance is driven by `q` (queries) + `n²` (matrix size).

---

## Complexity

* **Time Complexity:** `O(q + n^2)`

  * `q` is the number of queries (each processed in O(1)).
  * `n^2` is the cost of computing the 2D prefix sums for the `n x n` matrix.

* **Space Complexity:** `O(n^2)`

  * `diff` uses `(n+1) x (n+1)` space (which is `O(n^2)`), and the output `res` uses `n x n` space (also `O(n^2)`). We can re-use `diff` to avoid an extra matrix, but asymptotically space is `O(n^2)`.

---

## Multi-language Solutions

### C++

```c++
#include <vector>
using namespace std;

class Solution {
public:
    vector<vector<int>> rangeAddQueries(int n, vector<vector<int>>& queries) {
        // diff is (n+1) x (n+1)
        vector<vector<int>> diff(n+1, vector<int>(n+1, 0));

        // Apply queries: 4 corner updates each
        for (const auto &q : queries) {
            int r1 = q[0], c1 = q[1], r2 = q[2], c2 = q[3];
            diff[r1][c1] += 1;
            diff[r1][c2 + 1] -= 1;
            diff[r2 + 1][c1] -= 1;
            diff[r2 + 1][c2 + 1] += 1;
        }

        // Build final matrix using 2D prefix sums
        vector<vector<int>> res(n, vector<int>(n, 0));
        for (int i = 0; i < n; ++i) {
            for (int j = 0; j < n; ++j) {
                int up = (i > 0) ? diff[i-1][j] : 0;
                int left = (j > 0) ? diff[i][j-1] : 0;
                int diag = (i > 0 && j > 0) ? diff[i-1][j-1] : 0;
                diff[i][j] = diff[i][j] + up + left - diag;
                res[i][j] = diff[i][j];
            }
        }
        return res;
    }
};
```

---

### Java

```java
class Solution {
    public int[][] rangeAddQueries(int n, int[][] queries) {
        int[][] diff = new int[n+1][n+1];

        // Apply corner updates for each query
        for (int[] q : queries) {
            int r1 = q[0], c1 = q[1], r2 = q[2], c2 = q[3];
            diff[r1][c1] += 1;
            diff[r1][c2+1] -= 1;
            diff[r2+1][c1] -= 1;
            diff[r2+1][c2+1] += 1;
        }

        // Convert diff to final matrix with 2D prefix sums
        int[][] res = new int[n][n];
        for (int i = 0; i < n; ++i) {
            for (int j = 0; j < n; ++j) {
                int up = (i > 0) ? diff[i-1][j] : 0;
                int left = (j > 0) ? diff[i][j-1] : 0;
                int diag = (i > 0 && j > 0) ? diff[i-1][j-1] : 0;
                diff[i][j] = diff[i][j] + up + left - diag;
                res[i][j] = diff[i][j];
            }
        }
        return res;
    }
}
```

---

### JavaScript

```javascript
/**
 * @param {number} n
 * @param {number[][]} queries
 * @return {number[][]}
 */
var rangeAddQueries = function(n, queries) {
    const diff = Array.from({length: n+1}, () => new Array(n+1).fill(0));

    // 4 updates per query
    for (const q of queries) {
        const [r1, c1, r2, c2] = q;
        diff[r1][c1] += 1;
        diff[r1][c2 + 1] -= 1;
        diff[r2 + 1][c1] -= 1;
        diff[r2 + 1][c2 + 1] += 1;
    }

    // Build final matrix using 2D prefix sums
    const res = Array.from({length: n}, () => new Array(n).fill(0));
    for (let i = 0; i < n; ++i) {
        for (let j = 0; j < n; ++j) {
            const up = (i > 0) ? diff[i-1][j] : 0;
            const left = (j > 0) ? diff[i][j-1] : 0;
            const diag = (i > 0 && j > 0) ? diff[i-1][j-1] : 0;
            diff[i][j] = diff[i][j] + up + left - diag;
            res[i][j] = diff[i][j];
        }
    }
    return res;
};
```

---

### Python3

```python
from typing import List

class Solution:
    def rangeAddQueries(self, n: int, queries: List[List[int]]) -> List[List[int]]:
        # diff array of size (n+1) x (n+1)
        diff = [[0] * (n+1) for _ in range(n+1)]

        # Apply corner updates
        for r1, c1, r2, c2 in queries:
            diff[r1][c1] += 1
            diff[r1][c2 + 1] -= 1
            diff[r2 + 1][c1] -= 1
            diff[r2 + 1][c2 + 1] += 1

        # Build result via 2D prefix sum
        res = [[0] * n for _ in range(n)]
        for i in range(n):
            for j in range(n):
                up = diff[i-1][j] if i > 0 else 0
                left = diff[i][j-1] if j > 0 else 0
                diag = diff[i-1][j-1] if i > 0 and j > 0 else 0
                diff[i][j] = diff[i][j] + up + left - diag
                res[i][j] = diff[i][j]
        return res
```

---

### Go

```go
package main

func rangeAddQueries(n int, queries [][]int) [][]int {
    diff := make([][]int, n+1)
    for i := range diff {
        diff[i] = make([]int, n+1)
    }

    for _, q := range queries {
        r1, c1, r2, c2 := q[0], q[1], q[2], q[3]
        diff[r1][c1] += 1
        diff[r1][c2+1] -= 1
        diff[r2+1][c1] -= 1
        diff[r2+1][c2+1] += 1
    }

    res := make([][]int, n)
    for i := 0; i < n; i++ {
        res[i] = make([]int, n)
    }

    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            up, left, diag := 0, 0, 0
            if i > 0 {
                up = diff[i-1][j]
            }
            if j > 0 {
                left = diff[i][j-1]
            }
            if i > 0 && j > 0 {
                diag = diff[i-1][j-1]
            }
            diff[i][j] = diff[i][j] + up + left - diag
            res[i][j] = diff[i][j]
        }
    }
    return res
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

I will explain the common logic step-by-step; the code in each language follows the same operations.

1. **Create difference array**

   * I initialize a 2D array `diff` with shape `(n+1) x (n+1)` filled with zeros.
   * The extra row and column let me safely update `r2+1` and `c2+1` positions (these are allowed to be `n`).

2. **Apply each query using four corner updates (O(1) per query)**
   For query `[r1, c1, r2, c2]` I do:

   * `diff[r1][c1] += 1`
     This starts adding +1 at the top-left of the rectangle.
   * `diff[r1][c2+1] -= 1`
     This will cancel the addition to the right of the rectangle once prefix sums are computed.
   * `diff[r2+1][c1] -= 1`
     This will cancel the addition below the rectangle once prefix sums are computed.
   * `diff[r2+1][c2+1] += 1`
     This readjusts the overlapped cancellation area (inclusion-exclusion).

   These 4 ops guarantee that when we accumulate via prefix sums, only cells inside `[r1..r2, c1..c2]` will have the +1 influence.

3. **Compute the 2D prefix sums row-major (O(n²))**
   Iterate `i` from `0 .. n-1`, `j` from `0 .. n-1` and apply:

   ```
   diff[i][j] = diff[i][j] + (i>0 ? diff[i-1][j] : 0) + (j>0 ? diff[i][j-1] : 0) - (i>0 && j>0 ? diff[i-1][j-1] : 0)
   ```

   * This is the standard 2D prefix-sum recurrence.
   * After this update, `diff[i][j]` equals the total increments applied to `(i,j)`.
   * I copy `diff[i][j]` into `res[i][j]` (or return the portion of `diff` that corresponds to `0..n-1`).

4. **Return `res`**
   The `res` matrix is the final `n x n` matrix containing all increments from all queries.

---

## Examples

**Example 1**

```
Input:
n = 3,
queries = [[1,1,2,2],[0,0,1,1]]

Output:
[[1,1,0],
 [1,2,1],
 [0,1,1]]
```

Explanation:

* After first query ([1,1,2,2]) the bottom-right 2x2 block is incremented by 1.
* After second query ([0,0,1,1]) the top-left 2x2 block is incremented by 1.
* Final matrix as above.

**Example 2**

```
Input:
n = 2,
queries = [[0,0,1,1]]

Output:
[[1,1],
 [1,1]]
```

Explanation:

* The only query covers the whole 2x2 matrix; each cell becomes 1.

---

## How to use / Run locally

### C++

1. Save the C++ code to `solution.cpp` inside a function wrapper or with a small `main` to test.
2. Compile:

```bash
g++ -std=c++17 solution.cpp -O2 -o solution
```

3. Run:

```bash
./solution
```

(If used on LeetCode, just paste the `Solution` class into the editor.)

### Java

1. Put the Java code into `Solution.java`. If you want to run locally, create a `main` method that constructs test cases and calls `new Solution().rangeAddQueries(...)`.
2. Compile:

```bash
javac Solution.java
```

3. Run:

```bash
java Solution
```

(For LeetCode: paste the `Solution` class directly.)

### JavaScript (Node.js)

1. Save the JS function in `solution.js`. Add a small test harness to call `rangeAddQueries`.
2. Run:

```bash
node solution.js
```

(For LeetCode: paste the function in the editor.)

### Python3

1. Save the Python code in `solution.py`. Add a `if __name__ == "__main__":` block for tests.
2. Run:

```bash
python3 solution.py
```

(For LeetCode: paste the `Solution` class code.)

### Go

1. Save the code in `main.go`. Add `package main` and a `main()` to build tests or use the function in your program.
2. Run:

```bash
go run main.go
```

(If using in a library, integrate `rangeAddQueries` accordingly.)

---

## Notes & Optimizations

* The difference array trick (also called "Imos method" / 2D difference) is the key optimization. It transforms O(area * q) to O(q + n²).
* I used `(n+1) x (n+1)` arrays to avoid boundary checks when writing `r2+1` and `c2+1`.
* I do in-place prefix-sum accumulation on `diff` to avoid extra memory allocation. If immutable input is needed, copy values into a new `res`.
* Because `n <= 500`, `n² <= 250k`, which is comfortably handled in memory and time for the given constraints.
* If memory is tight and `q` is small, a naive approach may also be acceptable — but difference array is simple and robust.
* The same logic works when the matrix is not square: just adapt `m` and `n` and size the diff array to `(m+1) x (n+1)`.

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
