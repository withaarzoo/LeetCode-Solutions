# Problem Title

3350. Adjacent Increasing Subarrays Detection II

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

Given an integer array `nums` of length `n`, find the **maximum integer `k`** such that there exist two **adjacent** (consecutive) subarrays of length `k` each that are **strictly increasing**. Concretely, find the largest `k` for which there exists an index `a` with:

* `nums[a..a+k-1]` strictly increasing, and
* `nums[a+k..a+2k-1]` strictly increasing.

Return that maximum `k`. If no such positive `k` exists, return `0`.

---

## Constraints

* `2 <= nums.length <= 2 * 10^5`
* `-10^9 <= nums[i] <= 10^9`
* Two adjacent subarrays of length `k` require `2*k <= n`.

---

## Intuition

I thought: to check the property for a candidate `k` I need to know whether each required block of length `k` is strictly increasing. If I precompute, for every index `i`, how long a strictly increasing run starts at `i` (`inc[i]`), then testing any `k` becomes fast: for start `a`, the first block is valid if `inc[a] >= k` and second block (starting at `a+k`) is valid if `inc[a+k] >= k`. Since if some `k` works then all smaller `k` work, I can binary-search `k` to find the maximum.

---

## Approach

1. Precompute `inc[i]` = length of strictly increasing subarray starting at index `i`.

   * Traverse from right to left: if `nums[i] < nums[i+1]` then `inc[i] = inc[i+1] + 1`, else `inc[i] = 1`.
2. Binary search `k` in range `0 .. n/2`:

   * For candidate `k`, check if there exists `a` with `inc[a] >= k` and `inc[a+k] >= k` where `a + 2*k <= n`.
3. Return the largest feasible `k`.

This yields a simple, reliable O(n log n) solution.

---

## Data Structures Used

* Array / vector `inc[]` (length `n`) to store increasing-run lengths.
* No extra complex data structures. Binary search uses constant extra variables.

---

## Operations & Behavior Summary

* Precompute runs in O(n).
* For each `k` checked during binary search, perform an O(n) scan verifying adjacent runs.
* Binary search reduces the number of `k` candidates to O(log n).
* Behavior: returns maximum `k` such that two adjacent strictly increasing subarrays of length `k` exist; otherwise 0.

---

## Complexity

* **Time Complexity:** `O(n log n)`

  * `n` — length of `nums`. Precomputing `inc[]` is `O(n)`. Each binary-search iteration checks feasibility in `O(n)`, and there are `O(log n)` iterations.
* **Space Complexity:** `O(n)`

  * For the `inc[]` array. Other extra memory is `O(1)`.

---

## Multi-language Solutions

### C++

```c++
class Solution {
public:
    int maxIncreasingSubarrays(vector<int>& nums) {
        int n = nums.size();
        if (n < 2) return 0;

        // inc[i] = length of strictly increasing run starting at i
        vector<int> inc(n, 1);
        for (int i = n - 2; i >= 0; --i) {
            inc[i] = (nums[i] < nums[i+1]) ? inc[i+1] + 1 : 1;
        }

        auto feasible = [&](int k) -> bool {
            if (k == 0) return true;
            for (int a = 0; a + 2*k <= n; ++a) {
                if (inc[a] >= k && inc[a + k] >= k) return true;
            }
            return false;
        };

        int lo = 0, hi = n / 2, ans = 0;
        while (lo <= hi) {
            int mid = lo + (hi - lo) / 2;
            if (feasible(mid)) { ans = mid; lo = mid + 1; }
            else { hi = mid - 1; }
        }
        return ans;
    }
};
```

---

### Java

> **Note:** Many judges (LeetCode/Geeks) sometimes pass `List<Integer>` instead of `int[]`. This Java signature uses `List<Integer>` for compatibility.

```java
import java.util.List;
import java.util.function.IntPredicate;

class Solution {
    public int maxIncreasingSubarrays(List<Integer> nums) {
        int n = nums.size();
        if (n < 2) return 0;

        int[] inc = new int[n];
        inc[n-1] = 1;
        for (int i = n - 2; i >= 0; --i) {
            if (nums.get(i) < nums.get(i + 1)) inc[i] = inc[i + 1] + 1;
            else inc[i] = 1;
        }

        IntPredicate feasible = (k) -> {
            if (k == 0) return true;
            for (int a = 0; a + 2 * k <= n; ++a) {
                if (inc[a] >= k && inc[a + k] >= k) return true;
            }
            return false;
        };

        int lo = 0, hi = n / 2, ans = 0;
        while (lo <= hi) {
            int mid = lo + (hi - lo) / 2;
            if (feasible.test(mid)) { ans = mid; lo = mid + 1; }
            else { hi = mid - 1; }
        }
        return ans;
    }
}
```

---

### JavaScript

```javascript
/**
 * @param {number[]} nums
 * @return {number}
 */
var maxIncreasingSubarrays = function(nums) {
    const n = nums.length;
    if (n < 2) return 0;

    const inc = new Array(n).fill(1);
    for (let i = n - 2; i >= 0; --i) {
        inc[i] = nums[i] < nums[i+1] ? inc[i+1] + 1 : 1;
    }

    const feasible = (k) => {
        if (k === 0) return true;
        for (let a = 0; a + 2*k <= n; ++a) {
            if (inc[a] >= k && inc[a + k] >= k) return true;
        }
        return false;
    };

    let lo = 0, hi = Math.floor(n / 2), ans = 0;
    while (lo <= hi) {
        const mid = Math.floor((lo + hi) / 2);
        if (feasible(mid)) { ans = mid; lo = mid + 1; }
        else { hi = mid - 1; }
    }
    return ans;
};
```

---

### Python3

```python
from typing import List

class Solution:
    def maxIncreasingSubarrays(self, nums: List[int]) -> int:
        n = len(nums)
        if n < 2:
            return 0

        inc = [1] * n
        for i in range(n-2, -1, -1):
            inc[i] = inc[i+1] + 1 if nums[i] < nums[i+1] else 1

        def feasible(k: int) -> bool:
            if k == 0:
                return True
            for a in range(0, n - 2*k + 1):
                if inc[a] >= k and inc[a + k] >= k:
                    return True
            return False

        lo, hi, ans = 0, n // 2, 0
        while lo <= hi:
            mid = (lo + hi) // 2
            if feasible(mid):
                ans = mid
                lo = mid + 1
            else:
                hi = mid - 1
        return ans
```

---

### Go

```go
package main

func maxIncreasingSubarrays(nums []int) int {
    n := len(nums)
    if n < 2 { return 0 }

    inc := make([]int, n)
    inc[n-1] = 1
    for i := n - 2; i >= 0; i-- {
        if nums[i] < nums[i+1] {
            inc[i] = inc[i+1] + 1
        } else {
            inc[i] = 1
        }
    }

    feasible := func(k int) bool {
        if k == 0 { return true }
        for a := 0; a + 2*k <= n; a++ {
            if inc[a] >= k && inc[a+k] >= k { return true }
        }
        return false
    }

    lo, hi, ans := 0, n/2, 0
    for lo <= hi {
        mid := (lo + hi) / 2
        if feasible(mid) {
            ans = mid
            lo = mid + 1
        } else {
            hi = mid - 1
        }
    }
    return ans
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

I will explain the algorithm line-by-line in a way that applies to all implementations (the logic is the same in every language).

### 1. Edge case:

```text
if n < 2: return 0
```

* If the array length is less than 2, it's impossible to have two non-empty adjacent subarrays — answer is 0.

### 2. Build `inc[]` (increasing-run lengths)

```text
inc[n-1] = 1
for i = n-2 down to 0:
    if nums[i] < nums[i+1]:
        inc[i] = inc[i+1] + 1
    else:
        inc[i] = 1
```

* `inc[i]` stores how many consecutive positions starting from `i` form a strictly increasing sequence.
* We initialize every position as at least `1` (itself).
* By scanning right-to-left we can reuse `inc[i+1]` to compute `inc[i]` in O(1) time.

### 3. Feasibility check for a candidate `k`

```text
for a from 0 to n - 2*k:
    if inc[a] >= k and inc[a+k] >= k:
        return true
return false
```

* For adjacent blocks `[a .. a+k-1]` and `[a+k .. a+2k-1]`, the first block is strictly increasing if its start `a` has `inc[a] >= k`. The second block is strictly increasing if `inc[a+k] >= k`.
* If any `a` satisfies both, `k` is feasible.

### 4. Binary search for maximum `k`

```text
lo = 0
hi = n // 2
ans = 0
while lo <= hi:
    mid = (lo + hi) // 2
    if feasible(mid):
        ans = mid
        lo = mid + 1
    else:
        hi = mid - 1
return ans
```

* We binary search because the feasibility predicate is monotonic: if `k` is feasible, all `k' < k` are feasible too.
* Searching 0..n//2 ensures `2*k <= n`.

---

## Examples

1. Example 1:

```
Input: nums = [2,5,7,8,9,2,3,4,3,1]
Output: 3
Explanation:
 - Subarray [7,8,9] (start index 2) is strictly increasing, length 3.
 - Adjacent subarray [2,3,4] (start index 5) is strictly increasing, length 3.
```

2. Example 2:

```
Input: nums = [1,2,3,4,4,4,4,5,6,7]
Output: 2
Explanation:
 - Possible adjacent pairs of length 2: [1,2] and [3,4] (start 0 and 2).
 - No larger k works since duplicates break strict increase.
```

3. Edge case:

```
Input: nums = [5]
Output: 0
```

---

## How to use / Run locally

### Python

1. Create a file `solution.py` and paste the `Solution` class code.
2. Use a simple test harness:

```python
from typing import List
# paste Solution class here

s = Solution()
print(s.maxIncreasingSubarrays([2,5,7,8,9,2,3,4,3,1]))  # expects 3
```

### Java

1. Create a file `Solution.java` and paste the Java class (ensure the method signature matches your testing harness — e.g., convert from `List<Integer>` to `int[]` if your harness passes arrays).
2. Compile and run:

```bash
javac Solution.java
java Solution
```

### C++

1. Create `solution.cpp` and include the `Solution` class function inside `main()` or adapt to your judge harness.
2. Compile and run:

```bash
g++ solution.cpp -std=c++17 -O2 -o solution
./solution
```

### JavaScript (Node)

1. Create `solution.js` and export / call the function with test arrays.
2. Run:

```bash
node solution.js
```

### Go

1. Create `main.go`, paste the `maxIncreasingSubarrays` function, and call it inside `main`.
2. Run:

```bash
go run main.go
```

---

## Notes & Optimizations

* The `inc[]` precomputation is essential — it lets each feasibility check run in linear time without redundant comparisons.
* We pick binary search because the predicate "k is feasible" is monotonic. This reduces checks from O(n^2) to O(n log n).
* If you want an `O(n)` solution: it's possible with two-pointer or sliding-window-like techniques relying on runs; but the binary-search + `inc[]` approach is simple, clear, and fast enough for constraints up to `2*10^5`.
* Memory can be reduced if in-place modifications or compressed storage are applicable, but `O(n)` extra memory is acceptable for these constraints.

---

## Author

[Md. Aarzoo Islam](https://bento.me/withaarzoo)
