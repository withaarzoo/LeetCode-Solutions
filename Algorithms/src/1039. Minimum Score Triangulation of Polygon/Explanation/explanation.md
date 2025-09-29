# 1039. Minimum Score Triangulation of Polygon — README

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

We are given a convex polygon with `n` vertices. Each vertex `i` has an integer value `values[i]`. We must triangulate the polygon (split into `n-2` triangles using original vertices only). The **weight** of a triangle formed by vertices `i, j, k` is `values[i] * values[j] * values[k]`. The **score** of a triangulation is the sum of weights of all triangles. Return the **minimum possible score** of any valid triangulation.

---

## Constraints

* `n == values.length`
* `3 <= n <= 50`
* `1 <= values[i] <= 100` (integers)

---

## Intuition

I thought about how triangulations break the polygon into smaller sub-polygons. If I look at an interval of vertices from `i` to `j`, any triangulation of that interval must pick some vertex `k` between `i` and `j` to form triangle `(i, k, j)` and then triangulate the left and right parts. That means the problem has overlapping subproblems and I can use dynamic programming (interval DP) to build answers for larger intervals from smaller ones.

---

## Approach

1. Let `dp[i][j]` be the minimum triangulation score for the sub-polygon with vertices from index `i` to index `j` (inclusive).
2. Base: If `j - i + 1 < 3` (less than 3 vertices), score = 0 because no triangle.
3. For each interval length `len = 3..n`:

   * For every starting `i`, set `j = i + len - 1`.
   * For each possible `k` in `(i, j)` (i.e., `i+1..j-1`):

     * Consider triangle `(i, k, j)` with cost `values[i] * values[k] * values[j]`.
     * Combine with `dp[i][k]` and `dp[k][j]`.
     * `dp[i][j] = min(dp[i][j], dp[i][k] + dp[k][j] + values[i]*values[k]*values[j])`.
4. The answer is `dp[0][n-1]`.

This is classical interval DP: compute smaller intervals first so larger ones can reuse results.

---

## Data Structures Used

* 2D array `dp` of size `n x n` (integer). `dp[i][j]` stores minimum score for sub-polygon `[i..j]`.

---

## Operations & Behavior Summary

* We iterate intervals by increasing length.
* For each interval `[i..j]` we try every possible middle vertex `k` to split the interval into two subproblems and add the triangle cost `(i,k,j)`.
* We take the minimum across all `k`.

---

## Complexity

* **Time Complexity:** `O(n^3)`.

  * There are `O(n^2)` intervals `(i, j)` and for each interval we loop `O(n)` possible `k` choices.
* **Space Complexity:** `O(n^2)` for the `dp` table.

With `n <= 50` this is easily acceptable.

---

## Multi-language Solutions

### C++

```c++
#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    int minScoreTriangulation(vector<int>& values) {
        int n = values.size();
        if (n < 3) return 0;
        // dp[i][j] = minimum score to triangulate vertices i..j
        vector<vector<int>> dp(n, vector<int>(n, 0));
        // length is the number of vertices in sub-polygon
        for (int len = 3; len <= n; ++len) {
            for (int i = 0; i + len - 1 < n; ++i) {
                int j = i + len - 1;
                int best = INT_MAX;
                for (int k = i + 1; k < j; ++k) {
                    int cost = dp[i][k] + dp[k][j] + values[i] * values[k] * values[j];
                    best = min(best, cost);
                }
                dp[i][j] = best;
            }
        }
        return dp[0][n-1];
    }
};
```

### Java

```java
class Solution {
    public int minScoreTriangulation(int[] values) {
        int n = values.length;
        if (n < 3) return 0;
        int[][] dp = new int[n][n];
        for (int len = 3; len <= n; len++) {
            for (int i = 0; i + len - 1 < n; i++) {
                int j = i + len - 1;
                int best = Integer.MAX_VALUE;
                for (int k = i + 1; k < j; k++) {
                    int cost = dp[i][k] + dp[k][j] + values[i] * values[k] * values[j];
                    best = Math.min(best, cost);
                }
                dp[i][j] = best;
            }
        }
        return dp[0][n-1];
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} values
 * @return {number}
 */
var minScoreTriangulation = function(values) {
    const n = values.length;
    if (n < 3) return 0;
    const dp = Array.from({length: n}, () => Array(n).fill(0));
    for (let len = 3; len <= n; len++) {
        for (let i = 0; i + len - 1 < n; i++) {
            const j = i + len - 1;
            let best = Number.MAX_SAFE_INTEGER;
            for (let k = i + 1; k < j; k++) {
                const cost = dp[i][k] + dp[k][j] + values[i] * values[k] * values[j];
                if (cost < best) best = cost;
            }
            dp[i][j] = best;
        }
    }
    return dp[0][n-1];
};
```

### Python3

```python
from typing import List

class Solution:
    def minScoreTriangulation(self, values: List[int]) -> int:
        n = len(values)
        if n < 3:
            return 0
        dp = [[0] * n for _ in range(n)]
        for length in range(3, n + 1):
            for i in range(0, n - length + 1):
                j = i + length - 1
                best = float('inf')
                for k in range(i + 1, j):
                    cost = dp[i][k] + dp[k][j] + values[i] * values[k] * values[j]
                    if cost < best:
                        best = cost
                dp[i][j] = best
        return dp[0][n-1]
```

### Go

```go
package main

func minScoreTriangulation(values []int) int {
    n := len(values)
    if n < 3 {
        return 0
    }
    dp := make([][]int, n)
    for i := 0; i < n; i++ {
        dp[i] = make([]int, n)
    }
    const INF = int(^uint(0) >> 1)
    for length := 3; length <= n; length++ {
        for i := 0; i+length-1 < n; i++ {
            j := i + length - 1
            best := INF
            for k := i + 1; k < j; k++ {
                cost := dp[i][k] + dp[k][j] + values[i]*values[k]*values[j]
                if cost < best {
                    best = cost
                }
            }
            dp[i][j] = best
        }
    }
    return dp[0][n-1]
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

Below I explain the key logic in slow, simple steps. The same logic applies to all languages; only syntax changes.

### 1. Problem decomposition — subproblems

* I saw that triangulating a polygon from `i` to `j` always involves choosing a `k` where `i < k < j` to form the triangle `(i, k, j)`.
* After choosing `k`, the problem splits into two independent subproblems:

  * triangulate `i..k`
  * triangulate `k..j`
* Combined cost = `dp[i][k] + dp[k][j] + values[i]*values[k]*values[j]`.

### 2. Initialization

* `dp[i][j] = 0` whenever `j - i + 1 < 3` (less than 3 vertices). There is no triangle.
* We fill `dp` for intervals of increasing length starting at `3` up to `n`.

### 3. Main loops (common pattern)

Pseudocode of core loops:

```text
for len from 3 to n:
  for i from 0 to n - len:
    j = i + len - 1
    dp[i][j] = +inf
    for k from i+1 to j-1:
      dp[i][j] = min(dp[i][j], dp[i][k] + dp[k][j] + values[i]*values[k]*values[j])
```

* Outer loop: interval size.
* Middle loop: start index `i`.
* Inner loop: split vertex `k`.

This builds from small intervals to larger intervals, ensuring `dp[i][k]` and `dp[k][j]` are already computed.

### 4. Language-specific notes

* **C++**: `vector<vector<int>> dp(n, vector<int>(n, 0));` and `INT_MAX` as infinity sentinel. For safety with product sums you can use `long long` if you prefer.
* **Java**: Use `int[][] dp = new int[n][n];` and `Integer.MAX_VALUE` as sentinel.
* **JavaScript**: Use `Array.from` to create 2D array; `Number.MAX_SAFE_INTEGER` as sentinel.
* **Python3**: `dp = [[0] * n for _ in range(n)]` and `float('inf')` as sentinel.
* **Go**: Initialize `dp` as slice of slices. Use `^uint(0) >> 1` trick to get max int.

---

## Examples

1. Input: `values = [1,2,3]`
   Output: `6`
   Explanation: Only one triangle `(0,1,2)`: cost = `1*2*3 = 6`.

2. Input: `values = [3,7,4,5]`
   Output: `144`
   Explanation: Optimal triangulation yields cost `144`.

3. Input: `values = [1,3,1,4,1,5]`
   Output: `13`
   Explanation: One optimal triangulation gives cost `13`.

---

## How to use / Run locally

Below are simple ways to test the solutions locally. For each language, either paste the solution into a LeetCode custom test or use the small driver shown.

### C++

Save as `solution.cpp`:

```c++
#include <bits/stdc++.h>
using namespace std;

int minScoreTriangulation(vector<int>& values) {
    int n = values.size();
    if (n < 3) return 0;
    vector<vector<int>> dp(n, vector<int>(n, 0));
    for (int len = 3; len <= n; ++len) {
        for (int i = 0; i + len - 1 < n; ++i) {
            int j = i + len - 1;
            int best = INT_MAX;
            for (int k = i + 1; k < j; ++k) {
                int cost = dp[i][k] + dp[k][j] + values[i] * values[k] * values[j];
                best = min(best, cost);
            }
            dp[i][j] = best;
        }
    }
    return dp[0][n-1];
}

int main() {
    vector<int> values = {3,7,4,5};
    cout << minScoreTriangulation(values) << "\n";
    return 0;
}
```

Compile & run:

```bash
g++ -std=c++17 solution.cpp -O2 -o sol && ./sol
```

### Java

Save as `Solution.java`:

```java
public class Solution {
    public static int minScoreTriangulation(int[] values) {
        int n = values.length;
        if (n < 3) return 0;
        int[][] dp = new int[n][n];
        for (int len = 3; len <= n; len++) {
            for (int i = 0; i + len - 1 < n; i++) {
                int j = i + len - 1;
                int best = Integer.MAX_VALUE;
                for (int k = i + 1; k < j; k++) {
                    int cost = dp[i][k] + dp[k][j] + values[i] * values[k] * values[j];
                    best = Math.min(best, cost);
                }
                dp[i][j] = best;
            }
        }
        return dp[0][n-1];
    }
    public static void main(String[] args) {
        int[] values = {3,7,4,5};
        System.out.println(minScoreTriangulation(values));
    }
}
```

Compile & run:

```bash
javac Solution.java && java Solution
```

### JavaScript (Node)

Save as `solution.js`:

```javascript
function minScoreTriangulation(values) {
    const n = values.length;
    if (n < 3) return 0;
    const dp = Array.from({length: n}, () => Array(n).fill(0));
    for (let len = 3; len <= n; len++) {
        for (let i = 0; i + len - 1 < n; i++) {
            const j = i + len - 1;
            let best = Number.MAX_SAFE_INTEGER;
            for (let k = i + 1; k < j; k++) {
                const cost = dp[i][k] + dp[k][j] + values[i] * values[k] * values[j];
                if (cost < best) best = cost;
            }
            dp[i][j] = best;
        }
    }
    return dp[0][n-1];
}

console.log(minScoreTriangulation([3,7,4,5]));
```

Run:

```bash
node solution.js
```

### Python3

Save as `solution.py`:

```python
def minScoreTriangulation(values):
    n = len(values)
    if n < 3:
        return 0
    dp = [[0] * n for _ in range(n)]
    for length in range(3, n + 1):
        for i in range(0, n - length + 1):
            j = i + length - 1
            best = float('inf')
            for k in range(i + 1, j):
                cost = dp[i][k] + dp[k][j] + values[i] * values[k] * values[j]
                if cost < best:
                    best = cost
            dp[i][j] = best
    return dp[0][n-1]

print(minScoreTriangulation([3,7,4,5]))
```

Run:

```bash
python3 solution.py
```

### Go

Save as `main.go`:

```go
package main

import "fmt"

func minScoreTriangulation(values []int) int {
    n := len(values)
    if n < 3 {
        return 0
    }
    dp := make([][]int, n)
    for i := range dp {
        dp[i] = make([]int, n)
    }
    const INF = int(^uint(0) >> 1)
    for length := 3; length <= n; length++ {
        for i := 0; i+length-1 < n; i++ {
            j := i + length - 1
            best := INF
            for k := i + 1; k < j; k++ {
                cost := dp[i][k] + dp[k][j] + values[i]*values[k]*values[j]
                if cost < best {
                    best = cost
                }
            }
            dp[i][j] = best
        }
    }
    return dp[0][n-1]
}

func main() {
    fmt.Println(minScoreTriangulation([]int{3,7,4,5}))
}
```

Run:

```bash
go run main.go
```

---

## Notes & Optimizations

* For given constraints (`n <= 50`), `O(n^3)` is perfectly fine.
* `values[i] <= 100`, maximum triangle product is `100^3 = 1,000,000`. Summing up to `n-2` triangles still fits in 32-bit signed integer for `n <= 50`. If you prefer safety, use 64-bit (`long long` in C++, `long` in Java) for accumulation.
* Alternative: top-down memoization (recursive DP + memo) gives the same complexity but some find bottom-up easier to reason about.
* Micro-optimizations: iterate `len` increasing and break loops early if possible. But main bottleneck is cubic behavior and unavoidable for this DP formulation.

---

## Author

[Md. Aarzoo Islam](https://bento.me/withaarzoo)
