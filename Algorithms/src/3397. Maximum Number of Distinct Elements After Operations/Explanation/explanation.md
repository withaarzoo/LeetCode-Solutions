# Maximum Number of Distinct Elements After Operations (LeetCode 3397)

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

You are given an integer array `nums` and an integer `k`. For each element, you may perform the following operation **at most once**:

* Add any integer in the range `[-k, k]` to the element.

Return the **maximum possible number of distinct** elements in `nums` after performing these operations.

Equivalently: Each element `x` can become any integer in interval `[x - k, x + k]`. Choose a value inside each interval (or leave it unassigned) so that chosen values are distinct; maximize how many elements get a distinct value.

---

## Constraints

* `1 <= nums.length <= 10^5`
* `1 <= nums[i] <= 10^9`
* `0 <= k <= 10^9`

---

## Intuition

I thought of each element `x` as an interval `[x - k, x + k]` of integers it can become. The problem becomes: *Given `n` intervals, assign a distinct integer (point) inside as many intervals as possible.*
I recognized this as a standard greedy interval-to-point assignment problem: sort intervals by their right endpoint and always pick the smallest available integer inside an interval. This strategy leaves the most flexibility for later intervals and maximizes the number of assignments.

---

## Approach

1. Convert each `nums[i]` into interval `[nums[i] - k, nums[i] + k]` (use 64-bit to avoid overflow).
2. Sort intervals by their right endpoint (`r`) primarily, and by left endpoint (`l`) secondarily.
3. Maintain `last_assigned` as the last integer I placed. Initialize it to a very small value.
4. For each interval `[l, r]` (in sorted order):

   * Choose `assigned = max(l, last_assigned + 1)` (the earliest integer inside the interval that’s unused).
   * If `assigned <= r`, this interval gets `assigned`; increment count and update `last_assigned`.
   * Otherwise skip this interval (no available integer remains inside).
5. Return the count.

This greedy is optimal because choosing the earliest feasible integer for the interval with smallest right endpoint leaves maximum room for later intervals.

---

## Data Structures Used

* Array / vector to store intervals (pairs or two-element arrays).
* Sorting (O(n log n)).
* A few primitive variables (counters, `last_assigned`).

---

## Operations & Behavior Summary

* Convert each number into its reachable integer interval.
* Sort intervals by right boundary for greedy selection.
* Attempt to assign the next unused integer to each interval.
* Count successful unique assignments.

---

## Complexity

* **Time Complexity:** `O(n log n)` — `n` is `nums.length`. Sorting dominates time.
* **Space Complexity:** `O(n)` — for storing intervals. Extra O(1) beyond intervals.

---

## Multi-language Solutions

### C++

```c++
#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    int maxDistinctElements(vector<int>& nums, int k) {
        int n = nums.size();
        vector<pair<long long,long long>> intervals;
        intervals.reserve(n);
        for (int x : nums) {
            long long l = (long long)x - k;
            long long r = (long long)x + k;
            intervals.emplace_back(l, r);
        }
        // Sort by right endpoint, then left
        sort(intervals.begin(), intervals.end(), [](auto &a, auto &b){
            if (a.second != b.second) return a.second < b.second;
            return a.first < b.first;
        });

        long long last_assigned = LLONG_MIN / 4; // safe very small value
        int ans = 0;
        for (auto &it : intervals) {
            long long l = it.first, r = it.second;
            long long assigned = max(l, last_assigned + 1);
            if (assigned <= r) {
                ans++;
                last_assigned = assigned;
            }
        }
        return ans;
    }
};
```

---

### Java

```java
import java.util.*;

class Solution {
    public int maxDistinctElements(int[] nums, int k) {
        int n = nums.length;
        long[][] intervals = new long[n][2];
        for (int i = 0; i < n; ++i) {
            intervals[i][0] = (long)nums[i] - k;
            intervals[i][1] = (long)nums[i] + k;
        }
        Arrays.sort(intervals, (a,b) -> {
            if (a[1] != b[1]) return Long.compare(a[1], b[1]);
            return Long.compare(a[0], b[0]);
        });

        long lastAssigned = Long.MIN_VALUE / 4;
        int ans = 0;
        for (int i = 0; i < n; ++i) {
            long l = intervals[i][0], r = intervals[i][1];
            long assigned = Math.max(l, lastAssigned + 1);
            if (assigned <= r) {
                ans++;
                lastAssigned = assigned;
            }
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
 * @param {number} k
 * @return {number}
 */
var maxDistinctElements = function(nums, k) {
    const n = nums.length;
    const intervals = new Array(n);
    for (let i = 0; i < n; ++i) {
        intervals[i] = [nums[i] - k, nums[i] + k];
    }
    intervals.sort((a,b) => {
        if (a[1] !== b[1]) return a[1] - b[1];
        return a[0] - b[0];
    });

    let lastAssigned = Number.MIN_SAFE_INTEGER / 4;
    let ans = 0;
    for (const [l, r] of intervals) {
        const assigned = Math.max(l, lastAssigned + 1);
        if (assigned <= r) {
            ans++;
            lastAssigned = assigned;
        }
    }
    return ans;
};
```

---

### Python3

```python
from typing import List

class Solution:
    def maxDistinctElements(self, nums: List[int], k: int) -> int:
        intervals = [(x - k, x + k) for x in nums]
        intervals.sort(key=lambda t: (t[1], t[0]))
        last_assigned = -10**30  # very small sentinel
        ans = 0
        for l, r in intervals:
            assigned = max(l, last_assigned + 1)
            if assigned <= r:
                ans += 1
                last_assigned = assigned
        return ans
```

---

### Go

```go
package main

import (
 "sort"
 "math"
)

func maxDistinctElements(nums []int, k int) int {
 n := len(nums)
 intervals := make([][2]int64, n)
 for i := 0; i < n; i++ {
  intervals[i][0] = int64(nums[i]) - int64(k)
  intervals[i][1] = int64(nums[i]) + int64(k)
 }
 sort.Slice(intervals, func(i, j int) bool {
  if intervals[i][1] != intervals[j][1] {
   return intervals[i][1] < intervals[j][1]
  }
  return intervals[i][0] < intervals[j][0]
 })
 lastAssigned := int64(math.MinInt64 / 4)
 ans := 0
 for _, iv := range intervals {
  l, r := iv[0], iv[1]
  assigned := l
  if lastAssigned+1 > assigned {
   assigned = lastAssigned + 1
  }
  if assigned <= r {
   ans++
   lastAssigned = assigned
  }
 }
 return ans
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

I will explain the logic line-by-line in a language-agnostic manner, tying it to key lines in the provided implementations.

1. **Create intervals**

   * Code (Python example): `intervals = [(x - k, x + k) for x in nums]`
   * Explanation: For each element `x`, compute left `l = x - k` and right `r = x + k`. This interval contains all possible integers `x` can become. Use 64-bit integers in languages where overflow is possible (C++, Java, Go).

2. **Sort intervals by right endpoint**

   * Code (C++ example): `sort(intervals.begin(), intervals.end(), cmp_by_r_then_l);`
   * Explanation: Sort intervals ascending by `r`. If two `r` are equal, tie-break by `l` (not required, but stable). Sorting by `r` ensures we handle the interval that 'expires' earliest first — the classic greedy choice.

3. **Initialize `last_assigned` and `ans`**

   * Code (Java): `long lastAssigned = Long.MIN_VALUE / 4; int ans = 0;`
   * Explanation: `last_assigned` stores the last integer value I assigned to some interval. I place a very small sentinel initially so the first assignment is not constrained. `ans` counts successful distinct assignments.

4. **For each interval compute earliest feasible assignment**

   * Code (JavaScript): `assigned = Math.max(l, lastAssigned + 1);`
   * Explanation:

     * The smallest unused integer after previous assignments is `last_assigned + 1`.
     * But it must be at least `l` to lie in the current interval.
     * Hence `assigned = max(l, last_assigned + 1)` is the earliest integer that satisfies both conditions.

5. **Check feasibility and update**

   * Code (C++): `if (assigned <= r) { ++ans; last_assigned = assigned; }`
   * Explanation: If this `assigned` value is `<= r`, it lies inside the interval so I can use it. Increment the answer and update `last_assigned`. If `assigned > r`, this interval cannot get a fresh distinct integer (there is no available integer inside it) — skip it.

6. **Return the final answer**

   * Code (Python): `return ans`
   * Explanation: `ans` is the maximum number of elements that can be converted to distinct integers with at most one operation each.

---

## Examples

1. Example 1:

   * Input: `nums = [1,2,2,3,3,4], k = 2`
   * Explanation: Intervals -> `[ -1,3 ], [0,4], [0,4], [1,5], [1,5], [2,6]`. Greedy assignment yields `-1,0,1,2,3,4` (6 distinct).
   * Output: `6`

2. Example 2:

   * Input: `nums = [4,4,4,4], k = 1`
   * Explanation: Intervals -> `[3,5]` (four times). Greedy assignment gives `3,4,5` to three intervals; the fourth cannot get a new integer.
   * Output: `3`

---

## How to use / Run locally

### C++

* Put the `Solution` class into your local LeetCode submission or embed into a main driver for local testing.
* Compile with: `g++ -std=c++17 -O2 your_file.cpp -o run`
* Run: `./run`

### Java

* Place the `Solution` class in your Java project or submit to LeetCode.
* For local testing, wrap with `public static void main` and call the method.

### JavaScript (Node)

* Export the function or call it inside a script.
* Run: `node your_script.js`

### Python3

* Place the `Solution` class and call:

  ```py
  sol = Solution()
  print(sol.maxDistinctElements([1,2,2,3,3,4], 2))
  ```

### Go

* Put the function in `main` and call it.
* Build: `go build -o run main.go`
* Run: `./run`

---

## Notes & Optimizations

* Use 64-bit integers to compute `x ± k` to avoid overflow (`long long` in C++, `long` in Java, `int64` in Go).
* Sorting is the bottleneck — hence overall `O(n log n)` time.
* Memory usage is `O(n)` for intervals; we could sort an array of pairs in place.
* This greedy is optimal and memory/time efficient given constraints.
* Edge cases:

  * `k = 0`: intervals are single points. The algorithm reduces to counting unique elements.
  * Large `k`: intervals can overlap widely; still works.

---

## Author

* [Md. Aarzoo Islam](https://bento.me/withaarzoo)
