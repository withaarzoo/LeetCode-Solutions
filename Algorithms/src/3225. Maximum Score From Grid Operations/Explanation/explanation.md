# Maximum Score From Grid Operations

## Table of Contents

* [Problem Summary](#problem-summary)
* [Constraints](#constraints)
* [Intuition](#intuition)
* [Approach](#approach)
* [Data Structures Used](#data-structures-used)
* [Operations & Behavior Summary](#operations--behavior-summary)
* [Complexity](#complexity)
* [Multi-language Solutions](#multi-language-solutions)
* [Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)](#step-by-step-detailed-explanation-c-java-javascript-python3-go)
* [Examples](#examples)
* [How to Use / Run Locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

---

## Problem Summary

This problem is about maximizing a score from a grid using specific operations.

You are given an `n x n` grid. Initially, all cells are white.

In one operation:

* You pick a column
* You paint cells from the top down to a certain row black

After performing any number of such operations, you calculate a score.

A cell contributes to the score if:

* It is still white
* It has at least one horizontal neighbor (left or right) that is black

Your goal is to choose operations in such a way that the total score is maximized.

This is a classic **dynamic programming + prefix sum** problem involving grid processing and optimization.

---

## Constraints

* `1 ≤ n ≤ 100`
* `n == grid.length`
* `n == grid[i].length`
* `0 ≤ grid[i][j] ≤ 10^9`

---

## Intuition

When I first looked at the problem, I realized something important:

Each column can only be painted from the top downward. That means:

* The black cells in a column always form a prefix

So instead of thinking about each cell individually, I can represent each column using just one value:

* The number of black cells in that column

Now the problem becomes:

* How do different column heights interact to generate score?

A white cell contributes only if its neighbor column has black cells reaching that row.

This means the score depends on:

* The current column height
* The heights of its left and right neighbors

That immediately suggests using **Dynamic Programming** with states based on column heights.

---

## Approach

Let’s break it down step by step:

### Step 1: Convert grid into prefix sums

For each column, I calculate prefix sums so I can quickly get the sum of any vertical segment.

### Step 2: Represent columns by height

Each column is represented by how many top cells are painted black.

### Step 3: Define DP state

I define:

* `dp[a][b]` = maximum score when:

  * previous column has height `a`
  * current column has height `b`

### Step 4: Transition

When I move to the next column with height `c`, I compute the score contribution of the current column.

This depends on:

* max height between left and right neighbors

So I calculate:

* score = contribution based on `max(a, c)` and current `b`

### Step 5: Optimization

Instead of checking all combinations (which is slow), I optimize using:

* prefix maximum
* suffix maximum

This reduces complexity significantly.

---

## Data Structures Used

* **2D Array (DP Table)**
  Stores the best score for combinations of column heights

* **Prefix Sum Arrays**
  Helps in calculating column sums efficiently in O(1)

* **Temporary Arrays (prefix max / suffix max)**
  Used to optimize transitions and reduce time complexity

---

## Operations & Behavior Summary

Here’s what the algorithm does in simple terms:

1. Precompute prefix sums for all columns
2. Initialize DP for the first column
3. For each column:

   * Try all possible heights
   * Calculate contribution using prefix sums
   * Use prefix/suffix max to optimize transitions
4. Keep updating DP
5. Return the maximum value from DP

---

## Complexity

| Type             | Complexity |
| ---------------- | ---------- |
| Time Complexity  | O(n³)      |
| Space Complexity | O(n²)      |

### Explanation

* We process `n` columns
* For each column, we consider all height pairs → `n²`
* Optimized transitions keep it at `O(n³)`

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    long long maximumScore(vector<vector<int>>& grid) {
        int n = (int)grid.size();
        if (n == 1) return 0;

        // pref[c][k] = sum of first k cells in column c
        vector<vector<long long>> pref(n, vector<long long>(n + 1, 0));
        for (int c = 0; c < n; ++c) {
            for (int r = 0; r < n; ++r) {
                pref[c][r + 1] = pref[c][r] + grid[r][c];
            }
        }

        const long long NEG = -(1LL << 60);

        // dp[a][b] = best score after processing up to current column,
        // with previous height = a and current height = b.
        vector<vector<long long>> dp(n + 1, vector<long long>(n + 1, NEG));

        // Initialize using the first column (column 0).
        // Its left neighbor is a dummy column of height 0.
        for (int a = 0; a <= n; ++a) {
            for (int b = 0; b <= n; ++b) {
                dp[a][b] = max(0LL, pref[0][b] - pref[0][a]);
            }
        }

        // Process columns 1..n-1
        for (int col = 1; col < n; ++col) {
            vector<vector<long long>> ndp(n + 1, vector<long long>(n + 1, NEG));

            for (int mid = 0; mid <= n; ++mid) {
                // q[x] = gain of column 'col' if max(neighbor height) becomes x
                vector<long long> q(n + 1, 0);
                for (int x = 0; x <= n; ++x) {
                    q[x] = max(0LL, pref[col][x] - pref[col][mid]);
                }

                // prefixBest[c] = max dp[a][mid] for all a <= c
                vector<long long> prefixBest(n + 1, NEG);
                prefixBest[0] = dp[0][mid];
                for (int a = 1; a <= n; ++a) {
                    prefixBest[a] = max(prefixBest[a - 1], dp[a][mid]);
                }

                // suffixBest[c] = max(dp[a][mid] + q[a]) for all a >= c
                vector<long long> suffixBest(n + 2, NEG);
                suffixBest[n] = dp[n][mid] + q[n];
                for (int a = n - 1; a >= 0; --a) {
                    suffixBest[a] = max(suffixBest[a + 1], dp[a][mid] + q[a]);
                }

                // For the last real column, the next height is fixed to 0.
                int limit = (col == n - 1 ? 0 : n);

                for (int nxt = 0; nxt <= limit; ++nxt) {
                    long long best = NEG;

                    if (prefixBest[nxt] != NEG) {
                        best = max(best, prefixBest[nxt] + q[nxt]);
                    }
                    if (suffixBest[nxt + 1] != NEG) {
                        best = max(best, suffixBest[nxt + 1]);
                    }

                    ndp[mid][nxt] = max(ndp[mid][nxt], best);
                }
            }

            dp.swap(ndp);
        }

        long long ans = 0;
        for (int a = 0; a <= n; ++a) {
            for (int b = 0; b <= n; ++b) {
                ans = max(ans, dp[a][b]);
            }
        }
        return ans;
    }
};
```

### Java

```java
class Solution {
    public long maximumScore(int[][] grid) {
        int n = grid.length;
        if (n == 1) return 0L;

        // pref[c][k] = sum of first k cells in column c
        long[][] pref = new long[n][n + 1];
        for (int c = 0; c < n; c++) {
            for (int r = 0; r < n; r++) {
                pref[c][r + 1] = pref[c][r] + grid[r][c];
            }
        }

        final long NEG = -(1L << 60);

        // dp[a][b] = best score after processing up to current column,
        // with previous height = a and current height = b.
        long[][] dp = new long[n + 1][n + 1];
        for (int a = 0; a <= n; a++) {
            for (int b = 0; b <= n; b++) {
                dp[a][b] = Math.max(0L, pref[0][b] - pref[0][a]);
            }
        }

        for (int col = 1; col < n; col++) {
            long[][] ndp = new long[n + 1][n + 1];
            for (int i = 0; i <= n; i++) {
                for (int j = 0; j <= n; j++) {
                    ndp[i][j] = NEG;
                }
            }

            for (int mid = 0; mid <= n; mid++) {
                long[] q = new long[n + 1];
                for (int x = 0; x <= n; x++) {
                    q[x] = Math.max(0L, pref[col][x] - pref[col][mid]);
                }

                long[] prefixBest = new long[n + 1];
                prefixBest[0] = dp[0][mid];
                for (int a = 1; a <= n; a++) {
                    prefixBest[a] = Math.max(prefixBest[a - 1], dp[a][mid]);
                }

                long[] suffixBest = new long[n + 2];
                for (int i = 0; i <= n + 1; i++) suffixBest[i] = NEG;
                suffixBest[n] = dp[n][mid] + q[n];
                for (int a = n - 1; a >= 0; a--) {
                    suffixBest[a] = Math.max(suffixBest[a + 1], dp[a][mid] + q[a]);
                }

                int limit = (col == n - 1 ? 0 : n);
                for (int nxt = 0; nxt <= limit; nxt++) {
                    long best = NEG;

                    if (prefixBest[nxt] != NEG) {
                        best = Math.max(best, prefixBest[nxt] + q[nxt]);
                    }
                    if (suffixBest[nxt + 1] != NEG) {
                        best = Math.max(best, suffixBest[nxt + 1]);
                    }

                    ndp[mid][nxt] = Math.max(ndp[mid][nxt], best);
                }
            }

            dp = ndp;
        }

        long ans = 0;
        for (int a = 0; a <= n; a++) {
            for (int b = 0; b <= n; b++) {
                ans = Math.max(ans, dp[a][b]);
            }
        }
        return ans;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[][]} grid
 * @return {number}
 */
var maximumScore = function (grid) {
    const n = grid.length;
    if (n === 1) return 0;

    // pref[c][k] = sum of first k cells in column c
    const pref = Array.from({ length: n }, () => Array(n + 1).fill(0));
    for (let c = 0; c < n; c++) {
        for (let r = 0; r < n; r++) {
            pref[c][r + 1] = pref[c][r] + grid[r][c];
        }
    }

    const NEG = -1e30;

    // dp[a][b] = best score after processing up to current column,
    // with previous height = a and current height = b.
    let dp = Array.from({ length: n + 1 }, () => Array(n + 1).fill(NEG));

    // Initialize using the first column.
    for (let a = 0; a <= n; a++) {
        for (let b = 0; b <= n; b++) {
            dp[a][b] = Math.max(0, pref[0][b] - pref[0][a]);
        }
    }

    for (let col = 1; col < n; col++) {
        const ndp = Array.from({ length: n + 1 }, () => Array(n + 1).fill(NEG));

        for (let mid = 0; mid <= n; mid++) {
            const q = Array(n + 1).fill(0);
            for (let x = 0; x <= n; x++) {
                q[x] = Math.max(0, pref[col][x] - pref[col][mid]);
            }

            const prefixBest = Array(n + 1).fill(NEG);
            prefixBest[0] = dp[0][mid];
            for (let a = 1; a <= n; a++) {
                prefixBest[a] = Math.max(prefixBest[a - 1], dp[a][mid]);
            }

            const suffixBest = Array(n + 2).fill(NEG);
            suffixBest[n] = dp[n][mid] + q[n];
            for (let a = n - 1; a >= 0; a--) {
                suffixBest[a] = Math.max(suffixBest[a + 1], dp[a][mid] + q[a]);
            }

            const limit = (col === n - 1 ? 0 : n);
            for (let nxt = 0; nxt <= limit; nxt++) {
                let best = NEG;

                if (prefixBest[nxt] !== NEG) {
                    best = Math.max(best, prefixBest[nxt] + q[nxt]);
                }
                if (suffixBest[nxt + 1] !== NEG) {
                    best = Math.max(best, suffixBest[nxt + 1]);
                }

                ndp[mid][nxt] = Math.max(ndp[mid][nxt], best);
            }
        }

        dp = ndp;
    }

    let ans = 0;
    for (let a = 0; a <= n; a++) {
        for (let b = 0; b <= n; b++) {
            ans = Math.max(ans, dp[a][b]);
        }
    }
    return ans;
};
```

### Python3

```python
from typing import List

class Solution:
    def maximumScore(self, grid: List[List[int]]) -> int:
        n = len(grid)
        if n == 1:
            return 0

        # pref[c][k] = sum of first k cells in column c
        pref = [[0] * (n + 1) for _ in range(n)]
        for c in range(n):
            s = 0
            for r in range(n):
                s += grid[r][c]
                pref[c][r + 1] = s

        NEG = -10**30

        # dp[a][b] = best score after processing up to current column,
        # with previous height = a and current height = b.
        dp = [[NEG] * (n + 1) for _ in range(n + 1)]

        # Initialize using the first column.
        for a in range(n + 1):
            for b in range(n + 1):
                dp[a][b] = max(0, pref[0][b] - pref[0][a])

        for col in range(1, n):
            ndp = [[NEG] * (n + 1) for _ in range(n + 1)]

            for mid in range(n + 1):
                # q[x] = gain of current column if the tallest neighbor height is x
                q = [max(0, pref[col][x] - pref[col][mid]) for x in range(n + 1)]

                # prefixBest[c] = max dp[a][mid] for a <= c
                prefixBest = [NEG] * (n + 1)
                prefixBest[0] = dp[0][mid]
                for a in range(1, n + 1):
                    prefixBest[a] = max(prefixBest[a - 1], dp[a][mid])

                # suffixBest[c] = max(dp[a][mid] + q[a]) for a >= c
                suffixBest = [NEG] * (n + 2)
                suffixBest[n] = dp[n][mid] + q[n]
                for a in range(n - 1, -1, -1):
                    suffixBest[a] = max(suffixBest[a + 1], dp[a][mid] + q[a])

                # For the last real column, the next height is fixed to 0.
                limit = 0 if col == n - 1 else n

                for nxt in range(limit + 1):
                    best = NEG

                    if prefixBest[nxt] != NEG:
                        best = max(best, prefixBest[nxt] + q[nxt])
                    if suffixBest[nxt + 1] != NEG:
                        best = max(best, suffixBest[nxt + 1])

                    ndp[mid][nxt] = max(ndp[mid][nxt], best)

            dp = ndp

        ans = 0
        for row in dp:
            ans = max(ans, max(row))
        return ans
```

### Go

```go
func maximumScore(grid [][]int) int64 {
 n := len(grid)
 if n == 1 {
  return 0
 }

 // pref[c][k] = sum of first k cells in column c
 pref := make([][]int64, n)
 for c := 0; c < n; c++ {
  pref[c] = make([]int64, n+1)
  var s int64 = 0
  for r := 0; r < n; r++ {
   s += int64(grid[r][c])
   pref[c][r+1] = s
  }
 }

 const NEG int64 = -(1 << 60)

 // dp[a][b] = best score after processing up to current column,
 // with previous height = a and current height = b.
 dp := make([][]int64, n+1)
 for i := 0; i <= n; i++ {
  dp[i] = make([]int64, n+1)
  for j := 0; j <= n; j++ {
   val := pref[0][j] - pref[0][i]
   if val < 0 {
    val = 0
   }
   dp[i][j] = val
  }
 }

 for col := 1; col < n; col++ {
  ndp := make([][]int64, n+1)
  for i := 0; i <= n; i++ {
   ndp[i] = make([]int64, n+1)
   for j := 0; j <= n; j++ {
    ndp[i][j] = NEG
   }
  }

  for mid := 0; mid <= n; mid++ {
   q := make([]int64, n+1)
   for x := 0; x <= n; x++ {
    val := pref[col][x] - pref[col][mid]
    if val < 0 {
     val = 0
    }
    q[x] = val
   }

   prefixBest := make([]int64, n+1)
   prefixBest[0] = dp[0][mid]
   for a := 1; a <= n; a++ {
    if dp[a][mid] > prefixBest[a-1] {
     prefixBest[a] = dp[a][mid]
    } else {
     prefixBest[a] = prefixBest[a-1]
    }
   }

   suffixBest := make([]int64, n+2)
   for i := 0; i <= n+1; i++ {
    suffixBest[i] = NEG
   }
   suffixBest[n] = dp[n][mid] + q[n]
   for a := n - 1; a >= 0; a-- {
    cand := dp[a][mid] + q[a]
    if cand > suffixBest[a+1] {
     suffixBest[a] = cand
    } else {
     suffixBest[a] = suffixBest[a+1]
    }
   }

   limit := n
   if col == n-1 {
    limit = 0
   }

   for nxt := 0; nxt <= limit; nxt++ {
    best := NEG

    if prefixBest[nxt] != NEG {
     cand := prefixBest[nxt] + q[nxt]
     if cand > best {
      best = cand
     }
    }
    if suffixBest[nxt+1] != NEG && suffixBest[nxt+1] > best {
     best = suffixBest[nxt+1]
    }

    if best > ndp[mid][nxt] {
     ndp[mid][nxt] = best
    }
   }
  }

  dp = ndp
 }

 var ans int64 = 0
 for a := 0; a <= n; a++ {
  for b := 0; b <= n; b++ {
   if dp[a][b] > ans {
    ans = dp[a][b]
   }
  }
 }
 return ans
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic is exactly the same across all languages.

### Step 1: Prefix Sum Creation

I compute prefix sums column-wise so I can quickly get sums between rows.

### Step 2: DP Initialization

For the first column:

* I assume the left neighbor has height 0
* I compute initial values

### Step 3: Transition Logic

For each column:

* I fix current height
* Try all possible next heights
* Compute score using prefix sums

### Step 4: Optimization

Instead of checking every previous height:

* I use prefix max for `a <= c`
* I use suffix max for `a > c`

This avoids unnecessary nested loops.

### Step 5: Final Answer

After processing all columns:

* I scan DP table
* Return the maximum value

---

## Examples

### Example 1

Input:

```
grid = [
 [0,0,0,0,0],
 [0,0,3,0,0],
 [0,1,0,0,0],
 [5,0,0,3,0],
 [0,0,0,0,2]
]
```

Output:

```
11
```

Explanation:

* Select optimal column heights
* Count only valid white cells with black neighbors

---

### Example 2

Input:

```
grid = [
 [10,9,0,0,15],
 [7,1,0,8,0],
 [5,20,0,11,0],
 [0,0,0,1,2],
 [8,12,1,10,3]
]
```

Output:

```
94
```

---

## How to Use / Run Locally

### C++

```
g++ -std=c++17 solution.cpp -o solution
./solution
```

### Java

```
javac Solution.java
java Solution
```

### JavaScript (Node.js)

```
node solution.js
```

### Python

```
python3 solution.py
```

### Go

```
go run solution.go
```

---

## Notes & Optimizations

* Always use prefix sums for fast range queries
* DP state should be carefully designed — brute force will TLE
* Optimization using prefix/suffix max is key here
* Edge case: when `n = 1`, answer is always 0
* Values can be large, so use `long long` or equivalent

Alternative approach:

* A naive 4D DP exists but is too slow
* This optimized 3-state DP is the intended solution

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
