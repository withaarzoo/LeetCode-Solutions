# Problem Title

**3349. Adjacent Increasing Subarrays Detection I**

---

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

Given an array `nums` of `n` integers and an integer `k`, determine whether there exist **two adjacent subarrays** of length `k` such that both subarrays are **strictly increasing**. The two subarrays must be adjacent (the second starts immediately after the first ends). Return `true` if such two adjacent strictly increasing subarrays exist, otherwise return `false`.

Concretely: find `a` such that subarray `nums[a..a+k-1]` and `nums[a+k..a+2k-1]` are both strictly increasing.

---

## Constraints

* `2 <= nums.length <= 100`
* `1 <= 2 * k <= nums.length`
* `-1000 <= nums[i] <= 1000`

---

## Intuition

I thought about what makes a length-`k` subarray strictly increasing: it needs `k-1` adjacent pairs where each pair satisfies `nums[j] < nums[j+1]`. If I can quickly know, for each index `i`, how many consecutive increasing adjacent pairs start at `i`, then I can check whether a length-`k` subarray starting at `i` is strictly increasing by testing whether that count is at least `k-1`. So I precompute an array `nextInc` where `nextInc[i]` = number of consecutive `nums[j] < nums[j+1]` pairs starting at `i`. Then I just check pairs of adjacent starts `i` and `i+k`.

---

## Approach

1. If `2*k > n`, return `false` because two adjacent length-`k` subarrays don't fit.
2. Build `nextInc[]` of length `n` by scanning from right to left:

   * If `nums[i] < nums[i+1]`, then `nextInc[i] = nextInc[i+1] + 1`.
   * Else `nextInc[i] = 0`.
     `nextInc[i]` counts consecutive increasing adjacent pairs starting at `i`.
3. For each start `i` such that both subarrays fit (`i + 2*k <= n`), check:

   * `nextInc[i] >= k-1` and `nextInc[i+k] >= k-1`.
   * If both true, return `true`.
4. If none found, return `false`.

This uses one backward pass to compute `nextInc[]` and one forward pass to check adjacency — total linear time.

---

## Data Structures Used

* `vector<int>` / `int[]` / `List<Integer>` / `Array` depending on language for the `nextInc` helper array.
* No extra complex data structures; only arrays and simple counters.

---

## Operations & Behavior Summary

* Preprocessing pass: compute consecutive increasing runs.
* Query pass: test each candidate `i` whether both the subarray at `i` and at `i+k` are strictly increasing.
* Return boolean result.

---

## Complexity

* **Time Complexity:** `O(n)` where `n = nums.length`. We do two linear scans: one to build `nextInc` and one to test candidates.
* **Space Complexity:** `O(n)` for the `nextInc` array. (We could reduce to `O(1)` space by using two sliding-window counters, but `O(n)` is straightforward and fine for the constraints.)

---

## Multi-language Solutions

### C++

```c++
#include <vector>
using namespace std;

class Solution {
public:
    bool hasIncreasingSubarrays(vector<int>& nums, int k) {
        int n = nums.size();
        if (2 * k > n) return false; // cannot place two adjacent subarrays

        // nextInc[i] = number of consecutive increasing adjacent pairs starting at i
        vector<int> nextInc(n, 0);
        for (int i = n - 2; i >= 0; --i) {
            if (nums[i] < nums[i + 1]) nextInc[i] = nextInc[i + 1] + 1;
            else nextInc[i] = 0;
        }

        int need = k - 1;
        for (int i = 0; i + 2 * k <= n; ++i) {
            if (nextInc[i] >= need && nextInc[i + k] >= need) return true;
        }
        return false;
    }
};
```

---

### Java

```java
import java.util.List;

class Solution {
    public boolean hasIncreasingSubarrays(List<Integer> nums, int k) {
        int n = nums.size();
        if (2 * k > n) return false;

        int[] nextInc = new int[n];
        for (int i = n - 2; i >= 0; --i) {
            if (nums.get(i) < nums.get(i + 1)) nextInc[i] = nextInc[i + 1] + 1;
            else nextInc[i] = 0;
        }

        int need = k - 1;
        for (int i = 0; i + 2 * k <= n; ++i) {
            if (nextInc[i] >= need && nextInc[i + k] >= need) return true;
        }
        return false;
    }
}
```

---

### JavaScript

```javascript
/**
 * @param {number[]} nums
 * @param {number} k
 * @return {boolean}
 */
var hasIncreasingSubarrays = function(nums, k) {
    const n = nums.length;
    if (2 * k > n) return false;

    const nextInc = new Array(n).fill(0);
    for (let i = n - 2; i >= 0; --i) {
        if (nums[i] < nums[i + 1]) nextInc[i] = nextInc[i + 1] + 1;
        else nextInc[i] = 0;
    }

    const need = k - 1;
    for (let i = 0; i + 2 * k <= n; ++i) {
        if (nextInc[i] >= need && nextInc[i + k] >= need) return true;
    }
    return false;
};
```

---

### Python3

```python
from typing import List

class Solution:
    def hasIncreasingSubarrays(self, nums: List[int], k: int) -> bool:
        n = len(nums)
        if 2 * k > n:
            return False

        next_inc = [0] * n
        for i in range(n - 2, -1, -1):
            if nums[i] < nums[i + 1]:
                next_inc[i] = next_inc[i + 1] + 1
            else:
                next_inc[i] = 0

        need = k - 1
        for i in range(0, n - 2 * k + 1):
            if next_inc[i] >= need and next_inc[i + k] >= need:
                return True
        return False
```

---

### Go

```go
package main

func hasIncreasingSubarrays(nums []int, k int) bool {
    n := len(nums)
    if 2*k > n {
        return false
    }

    nextInc := make([]int, n)
    for i := n-2; i >= 0; i-- {
        if nums[i] < nums[i+1] {
            nextInc[i] = nextInc[i+1] + 1
        } else {
            nextInc[i] = 0
        }
    }

    need := k - 1
    for i := 0; i + 2*k <= n; i++ {
        if nextInc[i] >= need && nextInc[i+k] >= need {
            return true
        }
    }
    return false
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

I'll explain the logic in a slow, friendly way and tie it to important blocks of code:

### 1. Quick feasibility check

* Why: If the array doesn't have room for two adjacent subarrays of length `k`, return `false`.
* Code (any language):

  ```cpp
  if (2 * k > n) return false;
  ```

  This checks that `n` can hold `2*k` elements.

### 2. Build `nextInc` array

* What `nextInc[i]` means: number of consecutive increasing *adjacent pairs* starting at index `i`.
* Example: `nums = [1, 2, 4, 3, 5]`

  * `nextInc[0] = 2` because `1<2` and `2<4` (two adjacent increasing pairs starting at 0).
  * A subarray of length `k` is strictly increasing iff `nextInc[start] >= k-1`.
* How to compute (right-to-left):

  ```python
  for i in range(n-2, -1, -1):
      if nums[i] < nums[i+1]:
          next_inc[i] = next_inc[i+1] + 1
      else:
          next_inc[i] = 0
  ```

  I scan from `n-2` down to `0`. If `nums[i] < nums[i+1]`, I extend the run from `i+1`.

### 3. Check adjacent subarrays

* For each possible start `i` where both subarrays fit (`i + 2*k <= n`):

  * Check if `nextInc[i] >= k-1` (first subarray strictly increasing).
  * Check if `nextInc[i+k] >= k-1` (second subarray strictly increasing).
  * If both true, return `true`.
* Code snippet:

  ```js
  const need = k - 1;
  for (let i = 0; i + 2*k <= n; ++i) {
      if (nextInc[i] >= need && nextInc[i + k] >= need) return true;
  }
  return false;
  ```

### 4. Why this is correct

* A strictly increasing subarray of length `k` requires `k-1` increasing adjacent pairs.
* `nextInc[start]` counts exactly those adjacent pairs; so comparing against `k-1` is sufficient.
* We only need adjacent subarrays, so the second start is exactly `start + k`.

---

## Examples

1. Example 1:

   * Input: `nums = [2,5,7,8,9,2,3,4,3,1]`, `k = 3`
   * Explanation:

     * Subarray starting at index `2`: `[7,8,9]` is strictly increasing.
     * Subarray starting at index `5`: `[2,3,4]` is strictly increasing.
     * These are adjacent (start at 2 and 5). Output: `true`.

2. Example 2:

   * Input: `nums = [1,2,3,4,4,4,4,5,6,7]`, `k = 5`
   * Explanation:

     * You cannot find two adjacent length-5 strictly increasing subarrays because runs break at equal values. Output: `false`.

3. Small edge:

   * Input: `nums = [1,2]`, `k = 1` → `true`?

     * Two adjacent subarrays of length 1 always trivially "increasing" because a single element is vacuously strictly increasing; both subarrays exist if `n >= 2*k`. Our solution handles such cases.

---

## How to use / Run locally

### C++ (g++)

1. Put the class into a `.cpp` file.
2. Wrap with a `main()` to test:

   ```cpp
   int main() {
       Solution s;
       vector<int> nums = {2,5,7,8,9,2,3,4,3,1};
       int k = 3;
       cout << (s.hasIncreasingSubarrays(nums, k) ? "true" : "false") << endl;
       return 0;
   }
   ```
3. Compile and run:

   ```bash
   g++ -std=c++17 -O2 solution.cpp -o solution
   ./solution
   ```

### Java

1. Save `Solution` class in `Solution.java`. Add `main()` to create a `List<Integer>` and test.
2. Compile and run:

   ```bash
   javac Solution.java
   java Solution
   ```

### JavaScript (Node.js)

1. Put the function in `solution.js`, add tests and `console.log`.
2. Run:

   ```bash
   node solution.js
   ```

### Python3

1. Save the class in `solution.py`, add a small test driver:

   ```bash
   python3 solution.py
   ```

### Go

1. Put function in `main.go` with sample call in `main()`.
2. Run:

   ```bash
   go run main.go
   ```

---

## Notes & Optimizations

* The solution is already `O(n)` time. Space is `O(n)` due to `nextInc`.
* If space is a premium, we can convert to `O(1)` extra space by maintaining counts of consecutive increasing pairs while scanning forward with a sliding window, but that makes code slightly more tricky. For clarity and maintainability, `nextInc` is preferred.
* Because `n <= 100` in problem constraints, both memory and time are trivial in practice; the main goal is correctness and clarity.

---

## Author

* [Md. Aarzoo Islam](https://bento.me/withaarzoo)
