# 120. Triangle

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
* [Step-by-step Detailed Explanation (all languages)](#step-by-step-detailed-explanation-all-languages)
* [Examples](#examples)
* [How to use / Run locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

---

## Problem Summary

Given a triangle (a 2D list where row `i` has `i+1` integers), return the **minimum path sum** from top to bottom.
At each step, I can move to adjacent numbers on the row below: from index `j` in row `i` I may move to index `j` or `j+1` in row `i+1`.

**Example:**

```bash
triangle = [
  [2],
  [3,4],
  [6,5,7],
  [4,1,8,3]
]
```

Minimum path 2 → 3 → 5 → 1 = **11**.

---

## Constraints

* `1 <= triangle.length <= 200`
* `triangle[0].length == 1`
* `triangle[i].length == triangle[i - 1].length + 1`
* `-10^4 <= triangle[i][j] <= 10^4`

Follow-up: Can it be done using `O(n)` extra space where `n` is number of rows?

---

## Intuition

I thought that if I already know the minimum path sums for the row below, then for any element in the current row I only need to add that element to the smaller of the two reachable sums beneath it. So I can start from the bottom and build up. That way I only keep one row of results at a time.

---

## Approach

1. Initialize a 1D array `dp` with the values of the last row of the triangle.
2. For each row `i` from `n-2` down to `0` (second-last to the top):

   * For each index `j` in row `i` (0 to `i`):

     * Update `dp[j] = triangle[i][j] + min(dp[j], dp[j+1])`.
3. After finishing, `dp[0]` contains the minimum path sum from top to bottom.

This is bottom-up dynamic programming and satisfies the `O(n)` extra space follow-up.

---

## Data Structures Used

* `dp`: a 1D array (vector/slice/list) of length `n` (number of rows). It stores intermediate minimum path sums for the currently processed row to the bottom.

---

## Operations & Behavior Summary

* **Initialization**: Copy last row into `dp`.
* **Processing**: For each row from bottom to top, update `dp[j]` by adding the current element and the smaller of the two `dp` values beneath.
* **Termination**: `dp[0]` holds the final answer.

---

## Complexity

* **Time Complexity:** `O(n^2)` where `n` is number of rows. The triangle has `n(n+1)/2` elements and we process each once.
* **Space Complexity:** `O(n)` extra space for the `dp` array (where `n` is number of rows). We could reduce to `O(1)` extra space by modifying the triangle in-place, but that mutates input.

---

## Multi-language Solutions

### C++

```c++
#include <vector>
#include <algorithm>
using namespace std;

class Solution {
public:
    int minimumTotal(vector<vector<int>>& triangle) {
        int n = triangle.size();
        // dp initialized to last row
        vector<int> dp(triangle.back());

        // bottom-up DP
        for (int i = n - 2; i >= 0; --i) {
            // row i has i+1 elements
            for (int j = 0; j <= i; ++j) {
                dp[j] = triangle[i][j] + min(dp[j], dp[j + 1]);
            }
        }
        return dp[0];
    }
};
```

### Java

```java
import java.util.List;

class Solution {
    public int minimumTotal(List<List<Integer>> triangle) {
        int n = triangle.size();
        int[] dp = new int[n];
        List<Integer> last = triangle.get(n - 1);
        for (int i = 0; i < n; i++) dp[i] = last.get(i);

        for (int i = n - 2; i >= 0; --i) {
            for (int j = 0; j <= i; ++j) {
                dp[j] = triangle.get(i).get(j) + Math.min(dp[j], dp[j + 1]);
            }
        }

        return dp[0];
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[][]} triangle
 * @return {number}
 */
var minimumTotal = function(triangle) {
    const n = triangle.length;
    // copy last row
    let dp = triangle[n - 1].slice();

    for (let i = n - 2; i >= 0; --i) {
        for (let j = 0; j <= i; ++j) {
            dp[j] = triangle[i][j] + Math.min(dp[j], dp[j + 1]);
        }
    }
    return dp[0];
};
```

### Python3

```python3
from typing import List

class Solution:
    def minimumTotal(self, triangle: List[List[int]]) -> int:
        n = len(triangle)
        dp = triangle[-1].copy()  # copy last row

        for i in range(n - 2, -1, -1):
            for j in range(i + 1):
                dp[j] = triangle[i][j] + min(dp[j], dp[j + 1])

        return dp[0]
```

### Go

```go
package main

func minimumTotal(triangle [][]int) int {
    n := len(triangle)
    dp := make([]int, n)
    copy(dp, triangle[n-1])

    for i := n-2; i >= 0; i-- {
        for j := 0; j <= i; j++ {
            if dp[j] < dp[j+1] {
                dp[j] = triangle[i][j] + dp[j]
            } else {
                dp[j] = triangle[i][j] + dp[j+1]
            }
        }
    }
    return dp[0]
}
```

---

## Step-by-step Detailed Explanation (all languages)

I'll explain the algorithm in small steps. The code above follows these exact steps.

1. **What `dp[j]` means**

   * After initialization and while processing row `i`, `dp[j]` represents the minimum path sum starting from position `j` in row `i+1` down to the bottom.

2. **Initialization**

   * Copy the last row of `triangle` into `dp`. This is because the minimum path from the last-row position to the bottom is the value itself.

3. **Bottom-up propagation**

   * For row index `i = n-2` down to `0`:

     * For each `j` from `0` to `i` (inclusive):

       * We can go from `triangle[i][j]` to `triangle[i+1][j]` or `triangle[i+1][j+1]`.
       * We already have the minimum sums for the row below stored in `dp[j]` and `dp[j+1]`.
       * So the minimum sum starting at `triangle[i][j]` is `triangle[i][j] + min(dp[j], dp[j+1])`.
       * Overwrite `dp[j]` with that new value. This effectively moves the `dp` meaning up one row.

4. **Finish**

   * After we process the top row (`i = 0`), `dp[0]` becomes the minimum path sum from top to bottom.

5. **Why it is correct**

   * This is dynamic programming. Each state depends only on the two adjacent states in the row below. By processing from bottom to top we ensure that when we compute `dp[j]` for row `i`, the needed states from the row `i+1` are already computed.

6. **In-place alternative**

   * If allowed to modify `triangle`, we can write the same logic overwriting the triangle rows from bottom to top to save the `dp` array. That achieves `O(1)` extra space beyond input storage but mutates input.

---

## Examples

**Example 1**

```
Input: triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
Output: 11
Explanation: Path 2 -> 3 -> 5 -> 1 = 11
```

**Example 2**

```
Input: triangle = [[-10]]
Output: -10
```

---

## How to use / Run locally

### General

* These solutions are written as functions/classes like on LeetCode. To run locally, you can either:

  1. Paste the function into the LeetCode editor and run tests.
  2. Add a small `main`/runner that constructs a sample `triangle` and prints the function result.

### Quick run examples

* **C++**: Save as `solution.cpp` with a driver `main()` that sets up `triangle` and calls `Solution().minimumTotal(triangle)`. Then compile: `g++ -std=c++17 -O2 solution.cpp -o solution && ./solution`.

* **Java**: Save class `Solution` and add a `public static void main(String[] args)` to construct triangle and call `new Solution().minimumTotal(triangle)`. Compile: `javac Solution.java` then `java Solution`.

* **JavaScript**: Save as `solution.js`, add code to create `triangle` and `console.log(minimumTotal(triangle))`. Run: `node solution.js`.

* **Python3**: Save as `solution.py`, add a test block:

```python
if __name__ == "__main__":
    triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
    print(Solution().minimumTotal(triangle))
```

Run: `python3 solution.py`.

* **Go**: Save as `solution.go` and add `func main()` to create triangle and `fmt.Println(minimumTotal(triangle))`. Run: `go run solution.go`.

---

## Notes & Optimizations

* The provided solution uses `O(n)` extra space (1D `dp`) — this satisfies the typical follow-up.
* If mutating the input is allowed, we can update `triangle` itself from bottom-up and use `O(1)` extra space (beyond input storage).
* Time complexity is optimal for this problem because we must at least visit every element.
* Edge cases: single-row triangle, negative numbers (works fine because `min` handles negatives just the same).

---

## Author

[Aarzoo Islam](https://bento.me/withaarzoo)
