# 2435. Paths in Matrix Whose Sum Is Divisible by K

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

We are given an `m x n` grid of **non-negative integers** and an integer `k`.

* We start at the **top-left** cell `(0, 0)`.
* We want to reach the **bottom-right** cell `(m-1, n-1)`.
* At each step we can move **only right** or **only down**.
* Every time we move onto a cell, we add the cell’s value to the path sum.

**Goal:**
Count how many different paths have a **total sum divisible by `k`**, and return the count **modulo `1e9 + 7`**.

---

## Constraints

From the original LeetCode problem:

* `m == grid.length`
* `n == grid[i].length`
* `1 <= m, n <= 5 * 10^4` and `m * n <= 5 * 10^4`
* `0 <= grid[i][j] <= 100`
* `1 <= k <= 50`
* Answer must be returned modulo `10^9 + 7`.

---

## Intuition

I first thought:
“I must count all paths and check if each path sum is divisible by `k`.”

But brute force is impossible:

* There can be an exponential number of paths.

Then I realized:

* To know if the sum is divisible by `k`, I do **not** need the full sum.
* I only care about the **remainder** of the sum when divided by `k`.

So:

* At each cell `(i, j)` I can have multiple paths, each with its own remainder `r` (`0 … k-1`).
* If I know **how many paths** reach `(i, j)` with each remainder, I can build the answer cell by cell using dynamic programming.

This leads to a DP where:

* The state records **position** + **remainder of sum mod `k`**.
* The transitions come from the **top** and **left** neighbors.

To save memory, I don’t need the whole 3D array.
I only need the **previous row** and the **current row**.

---

## Approach

1. **Define DP state**

   Let:

   ```text
   dp[i][j][r] = number of paths from (0,0) to (i,j)
                 with (sum of values along that path) % k == r
   ```

   But we compress it by rows:

   * `prev[j][r]` → `dp` for row `i-1`
   * `cur[j][r]`  → `dp` for row `i`

2. **Base case**

   At cell `(0,0)`:

   * Sum = `grid[0][0]`
   * Remainder = `grid[0][0] % k`
   * Exactly 1 path here.

   So:

   ```text
   cur[0][ grid[0][0] % k ] = 1
   ```

3. **Transition**

   For each cell `(i, j)` with value `val = grid[i][j]`:

   We consider all previous remainders `r_prev` from:

   * **Top cell** `(i-1, j)` → from `prev[j]`
   * **Left cell** `(i, j-1)` → from `cur[j-1]`

   When we add `val` to a path with remainder `r_prev`, new remainder is:

   ```text
   r_new = (r_prev + val) % k
   ```

   So:

   ```text
   cur[j][r_new] += paths_from_top_with_r_prev
   cur[j][r_new] += paths_from_left_with_r_prev
   cur[j][r_new] %= MOD
   ```

4. **Row iteration**

   * For each row `i`:

     * Reset `cur` to all zeros.
     * For each column `j`:

       * If `(i, j)` is the start cell, set the base once.
       * Otherwise, gather paths from **top** and **left** using the transition above.
   * After processing row `i`, swap `prev` and `cur`.

5. **Result**

   After finishing all rows:

   * `prev[n-1]` contains the DP values for the last row.
   * We want the number of paths whose sum is divisible by `k`, i.e., remainder `0`.

   So the answer is:

   ```text
   prev[n-1][0]
   ```

---

## Data Structures Used

* **2D array of size `n x k`** for `prev` row.
* **2D array of size `n x k`** for `cur` row.
* Both contain integers (or longs in JS) representing counts modulo `1e9+7`.

No extra complex structures are needed.

---

## Operations & Behavior Summary

For each cell `(i, j)`:

1. We compute `val = grid[i][j] % k`.
2. If `(i, j)` is `(0, 0)`:

   * Initialize `cur[0][val] = 1`.
3. Else:

   * For all remainders `r` (0 to `k-1`):

     * If from top exists (i > 0): update using `prev[j][r]`.
     * If from left exists (j > 0): update using `cur[j - 1][r]`.
4. Always keep counts modulo `1e9+7`.
5. After each row, we swap `prev` and `cur`.

This ensures we correctly count all paths with specific remainders, moving row by row.

---

## Complexity

Let:

* `m` = number of rows
* `n` = number of columns
* `k` = divisor

**Time Complexity:**
`O(m * n * k)`
For every cell `(i, j)`, we iterate over all `k` remainders.

**Space Complexity:**
`O(n * k)`
We only keep two rows of DP: `prev[n][k]` and `cur[n][k]`, not the whole `m x n x k`.

---

## Multi-language Solutions

### C++

```c++
#include <vector>
using namespace std;

class Solution {
public:
    int numberOfPaths(vector<vector<int>>& grid, int k) {
        const int MOD = 1'000'000'007;
        int m = grid.size();
        int n = grid[0].size();

        // prev[j][r] = number of paths to (i-1, j) with sum % k == r
        // cur[j][r]  = number of paths to (i,   j) with sum % k == r
        vector<vector<int>> prev(n, vector<int>(k, 0));
        vector<vector<int>> cur(n, vector<int>(k, 0));

        for (int i = 0; i < m; ++i) {
            // reset current row
            for (int j = 0; j < n; ++j)
                fill(cur[j].begin(), cur[j].end(), 0);

            for (int j = 0; j < n; ++j) {
                int val = grid[i][j] % k;

                // starting cell initialization
                if (i == 0 && j == 0) {
                    cur[0][val] = 1;
                    continue;
                }

                // paths coming from top (i-1, j)
                if (i > 0) {
                    for (int r = 0; r < k; ++r) {
                        if (prev[j][r] == 0) continue;
                        int nr = (r + val) % k;
                        cur[j][nr] = (cur[j][nr] + prev[j][r]) % MOD;
                    }
                }

                // paths coming from left (i, j-1)
                if (j > 0) {
                    for (int r = 0; r < k; ++r) {
                        if (cur[j - 1][r] == 0) continue;
                        int nr = (r + val) % k;
                        cur[j][nr] = (cur[j][nr] + cur[j - 1][r]) % MOD;
                    }
                }
            }

            // prepare for next row
            prev.swap(cur);
        }

        // bottom-right cell with remainder 0
        return prev[n - 1][0];
    }
};
```

---

### Java

```java
class Solution {
    private static final int MOD = 1_000_000_007;

    public int numberOfPaths(int[][] grid, int k) {
        int m = grid.length;
        int n = grid[0].length;

        // prev[j][r] = paths to (i-1, j) with sum % k == r
        // cur[j][r]  = paths to (i,   j) with sum % k == r
        int[][] prev = new int[n][k];
        int[][] cur  = new int[n][k];

        for (int i = 0; i < m; i++) {
            // clear current row
            for (int j = 0; j < n; j++) {
                java.util.Arrays.fill(cur[j], 0);
            }

            for (int j = 0; j < n; j++) {
                int val = grid[i][j] % k;

                // starting cell
                if (i == 0 && j == 0) {
                    cur[0][val] = 1;
                    continue;
                }

                // from top (i-1, j)
                if (i > 0) {
                    for (int r = 0; r < k; r++) {
                        if (prev[j][r] == 0) continue;
                        int nr = (r + val) % k;
                        cur[j][nr] = (cur[j][nr] + prev[j][r]) % MOD;
                    }
                }

                // from left (i, j-1)
                if (j > 0) {
                    for (int r = 0; r < k; r++) {
                        if (cur[j - 1][r] == 0) continue;
                        int nr = (r + val) % k;
                        cur[j][nr] = (cur[j][nr] + cur[j - 1][r]) % MOD;
                    }
                }
            }

            // move current row into prev
            int[][] tmp = prev;
            prev = cur;
            cur = tmp;
        }

        return prev[n - 1][0];
    }
}
```

---

### JavaScript

```javascript
/**
 * @param {number[][]} grid
 * @param {number} k
 * @return {number}
 */
var numberOfPaths = function(grid, k) {
    const MOD = 1_000_000_007;
    const m = grid.length;
    const n = grid[0].length;

    const makeRow = () =>
        Array.from({ length: n }, () => Array(k).fill(0));

    let prev = makeRow();
    let cur  = makeRow();

    for (let i = 0; i < m; i++) {
        // reset current row
        for (let j = 0; j < n; j++) cur[j].fill(0);

        for (let j = 0; j < n; j++) {
            const val = grid[i][j] % k;

            // starting cell
            if (i === 0 && j === 0) {
                cur[0][val] = 1;
                continue;
            }

            // from top (i-1, j)
            if (i > 0) {
                for (let r = 0; r < k; r++) {
                    if (prev[j][r] === 0) continue;
                    const nr = (r + val) % k;
                    cur[j][nr] = (cur[j][nr] + prev[j][r]) % MOD;
                }
            }

            // from left (i, j-1)
            if (j > 0) {
                for (let r = 0; r < k; r++) {
                    if (cur[j - 1][r] === 0) continue;
                    const nr = (r + val) % k;
                    cur[j][nr] = (cur[j][nr] + cur[j - 1][r]) % MOD;
                }
            }
        }

        // swap
        [prev, cur] = [cur, prev];
    }

    return prev[n - 1][0];
};
```

---

### Python3

```python
from typing import List

class Solution:
    def numberOfPaths(self, grid: List[List[int]], k: int) -> int:
        MOD = 10**9 + 7
        m, n = len(grid), len(grid[0])

        # prev[j][r] = paths to (i-1, j) with sum % k == r
        # cur[j][r]  = paths to (i,   j) with sum % k == r
        prev = [[0] * k for _ in range(n)]
        cur  = [[0] * k for _ in range(n)]

        for i in range(m):
            # reset current row
            for j in range(n):
                for r in range(k):
                    cur[j][r] = 0

            for j in range(n):
                val = grid[i][j] % k

                # starting cell
                if i == 0 and j == 0:
                    cur[0][val] = 1
                    continue

                # from top (i-1, j)
                if i > 0:
                    for r in range(k):
                        if prev[j][r] == 0:
                            continue
                        nr = (r + val) % k
                        cur[j][nr] = (cur[j][nr] + prev[j][r]) % MOD

                # from left (i, j-1)
                if j > 0:
                    for r in range(k):
                        if cur[j - 1][r] == 0:
                            continue
                        nr = (r + val) % k
                        cur[j][nr] = (cur[j][nr] + cur[j - 1][r]) % MOD

            prev, cur = cur, prev

        return prev[n - 1][0]
```

---

### Go

```go
package main

func numberOfPaths(grid [][]int, k int) int {
 const MOD int = 1_000_000_007
 m, n := len(grid), len(grid[0])

 // prev[j][r] = paths to (i-1, j) with sum % k == r
 // cur[j][r]  = paths to (i,   j) with sum % k == r
 prev := make([][]int, n)
 cur := make([][]int, n)
 for j := 0; j < n; j++ {
  prev[j] = make([]int, k)
  cur[j] = make([]int, k)
 }

 for i := 0; i < m; i++ {
  // reset current row
  for j := 0; j < n; j++ {
   for r := 0; r < k; r++ {
    cur[j][r] = 0
   }
  }

  for j := 0; j < n; j++ {
   val := grid[i][j] % k

   // starting cell
   if i == 0 && j == 0 {
    cur[0][val] = 1
    continue
   }

   // from top (i-1, j)
   if i > 0 {
    for r := 0; r < k; r++ {
     if prev[j][r] == 0 {
      continue
     }
     nr := (r + val) % k
     cur[j][nr] += prev[j][r]
     if cur[j][nr] >= MOD {
      cur[j][nr] -= MOD
     }
    }
   }

   // from left (i, j-1)
   if j > 0 {
    for r := 0; r < k; r++ {
     if cur[j-1][r] == 0 {
      continue
     }
     nr := (r + val) % k
     cur[j][nr] += cur[j-1][r]
     if cur[j][nr] >= MOD {
      cur[j][nr] -= MOD
     }
    }
   }
  }

  prev, cur = cur, prev
 }

 return prev[n-1][0]
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

I’ll describe the logic in a language-agnostic way; each implementation follows the same steps.

1. **Initialize constants and sizes**

   * Store `MOD = 1e9 + 7`.
   * Get `m = grid.length` and `n = grid[0].length`.

2. **Create DP arrays**

   * `prev[n][k]` and `cur[n][k]` filled with zeros.
   * `prev[j][r]` means: after finishing row `i-1`, number of paths to cell `(i-1, j)` with remainder `r`.
   * `cur[j][r]` means: during row `i`, paths to `(i, j)` with remainder `r`.

3. **Loop over rows**

   ```text
   for each row i from 0 to m-1:
       reset cur to all zeros
       for each column j from 0 to n-1:
           process cell (i, j)
       swap(prev, cur)
   ```

4. **Processing a cell `(i, j)`**

   * Compute `val = grid[i][j] % k`.
   * If `(i, j)` is `(0, 0)`:

     * Set `cur[0][val] = 1` and move on.
   * Else:

     * **From top**: if `i > 0`:

       * For each remainder `r` from `0` to `k-1`:

         * If `prev[j][r]` is not zero:

           * `nr = (r + val) % k`
           * Add `prev[j][r]` to `cur[j][nr]`.
     * **From left**: if `j > 0`:

       * For each remainder `r` from `0` to `k-1`:

         * If `cur[j-1][r]` is not zero:

           * `nr = (r + val) % k`
           * Add `cur[j-1][r]` to `cur[j][nr]`.
     * In both adds, we take modulo `MOD`.

5. **Swap rows**

   * After finishing row `i`, `cur` contains dp for row `i`.
   * Swap `prev` and `cur`.
   * Now `prev` represents row `i`, and `cur` will be reused for row `i+1`.

6. **Final answer**

   * After the last row (`i = m-1`), `prev[n-1]` is the dp for cell `(m-1, n-1)`.
   * We need remainder `0` ⇒ sum divisible by `k`.
   * Return `prev[n-1][0]`.

---

## Examples

### Example 1

```text
Input:
grid = [[5,2,4],
        [3,0,5],
        [0,7,2]],
k = 3

Output:
2
```

Explanation:

* There are exactly 2 paths from `(0,0)` to `(2,2)` where the path sum is divisible by 3.
* Our DP counts all paths grouped by remainder modulo `k`, and the count for remainder `0` at bottom-right is `2`.

---

### Example 2

```text
Input:
grid = [[0,0]],
k = 5

Output:
1
```

Explanation:

* Only one path (go right once).
* Sum = 0 which is divisible by 5.

---

### Example 3

```text
Input:
grid = [[7,3,4,9],
        [2,3,6,2],
        [2,3,7,0]],
k = 1

Output:
10
```

Explanation:

* Every integer is divisible by 1.
* So every possible path is valid, and the DP returns the total number of paths.

---

## How to use / Run locally

### C++

```bash
g++ -std=c++17 -O2 main.cpp -o main
./main
```

Make sure `main.cpp` contains the `Solution` class and a small `main()` to test:

```c++
int main() {
    vector<vector<int>> grid = {{5,2,4},{3,0,5},{0,7,2}};
    int k = 3;
    Solution sol;
    cout << sol.numberOfPaths(grid, k) << endl; // should print 2
}
```

---

### Java

```bash
javac Solution.java
java Solution
```

Your `Solution.java` can wrap a `main`:

```java
public class Solution {
    // ... class Solution implementation above ...
    public static void main(String[] args) {
        int[][] grid = {{5,2,4},{3,0,5},{0,7,2}};
        int k = 3;
        Solution sol = new Solution();
        System.out.println(sol.numberOfPaths(grid, k)); // 2
    }
}
```

---

### JavaScript (Node.js)

```bash
node main.js
```

`main.js`:

```javascript
// paste numberOfPaths implementation here

const grid = [[5,2,4],[3,0,5],[0,7,2]];
const k = 3;
console.log(numberOfPaths(grid, k)); // 2
```

---

### Python3

```bash
python3 main.py
```

`main.py`:

```python
# paste Solution class here

if __name__ == "__main__":
    grid = [[5,2,4],[3,0,5],[0,7,2]]
    k = 3
    print(Solution().numberOfPaths(grid, k))  # 2
```

---

### Go

```bash
go run main.go
```

`main.go`:

```go
package main

import "fmt"

// paste numberOfPaths function here

func main() {
    grid := [][]int{{5,2,4},{3,0,5},{0,7,2}}
    k := 3
    fmt.Println(numberOfPaths(grid, k)) // 2
}
```

---

## Notes & Optimizations

* The key optimization is **space reduction**:

  * Naive DP would use `O(m * n * k)` space.
  * We only store two rows (`prev` and `cur`), so we use `O(n * k)` space.
* We skip updates when counts are zero to save some constant time.
* Because `k <= 50`, the `O(m * n * k)` time is acceptable where `m * n <= 5 * 10^4`.
* Using modulo at every addition keeps values within integer range.

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
