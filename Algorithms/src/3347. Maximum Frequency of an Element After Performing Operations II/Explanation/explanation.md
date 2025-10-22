# Problem Title

**3347. Maximum Frequency of an Element After Performing Operations II**

---

## Table of Contents

* Problem Summary
* Constraints
* Intuition
* Approach
* Data Structures Used
* Operations & Behavior Summary
* Complexity
* Multi-language Solutions

  * C++
  * Java
  * JavaScript
  * Python3
  * Go
* Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)
* Examples
* How to use / Run locally
* Notes & Optimizations
* Author

---

## Problem Summary

You are given an integer array `nums` and two integers `k` and `numOperations`. You may perform up to `numOperations` operations. In one operation you select an index that hasn't been selected before and add any integer from `[-k, k]` to `nums[i]`. Return the maximum possible frequency of any element in `nums` after performing the operations.

Put simply: each element can be moved anywhere inside `[value - k, value + k]` once (if selected). You can choose up to `numOperations` distinct indices to adjust. What's the maximum number of elements that can become equal?

---

## Constraints

* `1 <= nums.length <= 1e5`
* `1 <= nums[i] <= 1e9`
* `0 <= k <= 1e9`
* `0 <= numOperations <= nums.length`

---

## Intuition

I thought: each number `a` can move anywhere in `[a-k, a+k]`.

* If I pick an existing value `v` as the final value, any number inside `[v-k, v+k]` can be converted to `v` with one operation per index (numbers already equal to `v` are free). So I can compute how many elements fall in `[v-k, v+k]`, count how many are already `v`, and then I can convert up to `numOperations` of the remaining ones.
* If I pick a **new** meeting value (not originally in array), elements `a` and `b` can meet if their reachable intervals overlap. Two intervals `[a-k, a+k]` and `[b-k, b+k]` overlap if `|a-b| <= 2*k`. That means any group inside a window of length `2*k` can be made equal to some common value (maybe new), limited by `numOperations`.

So I check both:

1. each existing value `v` as target (binary search the range `[v-k, v+k]`), and
2. sliding windows of width `2*k` to allow meeting at a non-existing value.

Take the maximum result.

---

## Approach

1. Sort `nums`. (Sorting helps us binary search ranges and use two pointers.)
2. Build frequency map for `nums` to know how many of each value are already present.
3. **Case A (existing target):** For each distinct value `v`:

   * Find `L = lower_bound(nums, v-k)` and `R = upper_bound(nums, v+k)`.
   * `totalInRange = R - L` (elements that can be made equal to `v`).
   * `already = freq[v]`.
   * `need = totalInRange - already` (indexes requiring an operation).
   * We can fix up to `min(need, numOperations)` of those. Candidate = `already + min(need, numOperations)`.
4. **Case B (meeting at new value):**

   * Two pointers `l, r` to find maximal windows where `nums[r] - nums[l] <= 2*k`.
   * Window size `w = r - l + 1`. They can all meet to some value (maybe new). Candidate = `min(w, numOperations)` (covers cases with no duplicates).
5. Answer = max over candidates from both cases.

This gives `O(n log n)` time (sorting + binary searches) and `O(n)` auxiliary for frequency (or `O(1)` if we only iterate distinct values via the sorted array).

---

## Data Structures Used

* Sorted array (in-place after sorting).
* Frequency map (hash map / Counter) to count existing duplicates.
* Binary search (`lower_bound`, `upper_bound`) or equivalent.
* Two-pointer sliding window.

---

## Operations & Behavior Summary

* `sort(nums)` — to allow ranged queries and sliding windows.
* For each distinct `v`: `lower_bound(v-k)` and `upper_bound(v+k)` to get the range convertible to `v`.
* Two-pointer scan to find lengths of subarrays where `nums[r]-nums[l] <= 2*k`.
* Use `min(numOperations, need)` to limit conversions by available operations.
* Maximize across all candidates.

---

## Complexity

* **Time Complexity:** `O(n log n)` where `n = nums.length`. Sorting is `O(n log n)`. For Case A we do a constant number of binary searches per distinct value (overall `O(n log n)`), and Case B is `O(n)`.
* **Space Complexity:** `O(n)` worst-case for the frequency map; otherwise `O(1)` extra beyond the sorted array.

---

## Multi-language Solutions

### C++

```c++
#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    int maxFrequency(vector<int>& nums, int k, int numOperations) {
        int n = nums.size();
        if (n == 0) return 0;
        sort(nums.begin(), nums.end());

        // freq map of existing values
        unordered_map<long long,int> freq;
        freq.reserve(n*2);
        for (int x : nums) freq[x]++;

        int ans = 1;

        // Case A: existing values as target
        for (auto &p : freq) {
            long long v = p.first;
            int already = p.second;
            long long lowVal = v - k;
            long long highVal = v + k;
            auto L = lower_bound(nums.begin(), nums.end(), (int)lowVal);
            auto R = upper_bound(nums.begin(), nums.end(), (int)highVal);
            int totalInRange = int(R - L);
            int need = totalInRange - already;
            int canFix = min(need, numOperations);
            ans = max(ans, already + canFix);
        }

        // Case B: non-existing meeting point using 2*k window
        int l = 0;
        for (int r = 0; r < n; ++r) {
            while (l <= r && (long long)nums[r] - nums[l] > 2LL * k) ++l;
            int w = r - l + 1;
            ans = max(ans, min(w, numOperations));
        }
        return ans;
    }
};
```

### Java

```java
import java.util.*;

class Solution {
    public int maxFrequency(int[] nums, int k, int numOperations) {
        int n = nums.length;
        if (n == 0) return 0;
        Arrays.sort(nums);

        Map<Long, Integer> freq = new HashMap<>();
        for (int x : nums) freq.put((long)x, freq.getOrDefault((long)x, 0) + 1);

        int ans = 1;

        // Case A
        for (Map.Entry<Long, Integer> entry : freq.entrySet()) {
            long v = entry.getKey();
            int already = entry.getValue();
            long lowVal = v - k;
            long highVal = v + k;
            int L = lowerBound(nums, lowVal);
            int R = upperBound(nums, highVal);
            int totalInRange = R - L;
            int need = totalInRange - already;
            int canFix = Math.min(need, numOperations);
            ans = Math.max(ans, already + canFix);
        }

        // Case B
        int l = 0;
        for (int r = 0; r < n; ++r) {
            while (l <= r && (long)nums[r] - nums[l] > 2L * k) l++;
            int w = r - l + 1;
            ans = Math.max(ans, Math.min(w, numOperations));
        }
        return ans;
    }

    private int lowerBound(int[] arr, long target) {
        int l = 0, r = arr.length;
        while (l < r) {
            int mid = (l + r) >>> 1;
            if ((long)arr[mid] < target) l = mid + 1;
            else r = mid;
        }
        return l;
    }

    private int upperBound(int[] arr, long target) {
        int l = 0, r = arr.length;
        while (l < r) {
            int mid = (l + r) >>> 1;
            if ((long)arr[mid] <= target) l = mid + 1;
            else r = mid;
        }
        return l;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} nums
 * @param {number} k
 * @param {number} numOperations
 * @return {number}
 */
var maxFrequency = function(nums, k, numOperations) {
  const n = nums.length;
  if (n === 0) return 0;
  nums.sort((a,b) => a - b);

  const freq = new Map();
  for (const x of nums) freq.set(x, (freq.get(x) || 0) + 1);

  const lowerBound = (arr, target) => {
    let l = 0, r = arr.length;
    while (l < r) {
      const mid = (l + r) >> 1;
      if (arr[mid] < target) l = mid + 1; else r = mid;
    }
    return l;
  };
  const upperBound = (arr, target) => {
    let l = 0, r = arr.length;
    while (l < r) {
      const mid = (l + r) >> 1;
      if (arr[mid] <= target) l = mid + 1; else r = mid;
    }
    return l;
  };

  let ans = 1;

  // Case A
  for (const [v, already] of freq.entries()) {
    const L = lowerBound(nums, v - k);
    const R = upperBound(nums, v + k);
    const totalInRange = R - L;
    const need = totalInRange - already;
    const canFix = Math.min(need, numOperations);
    ans = Math.max(ans, already + canFix);
  }

  // Case B
  let l = 0;
  for (let r = 0; r < n; ++r) {
    while (l <= r && nums[r] - nums[l] > 2 * k) l++;
    const w = r - l + 1;
    ans = Math.max(ans, Math.min(w, numOperations));
  }

  return ans;
};
```

### Python3

```python
from typing import List
import bisect
from collections import Counter

class Solution:
    def maxFrequency(self, nums: List[int], k: int, numOperations: int) -> int:
        if not nums:
            return 0
        nums.sort()
        n = len(nums)
        freq = Counter(nums)
        ans = 1

        # Case A
        for v, already in freq.items():
            L = bisect.bisect_left(nums, v - k)
            R = bisect.bisect_right(nums, v + k)
            totalInRange = R - L
            need = totalInRange - already
            canFix = min(need, numOperations)
            ans = max(ans, already + canFix)

        # Case B
        l = 0
        for r in range(n):
            while l <= r and nums[r] - nums[l] > 2 * k:
                l += 1
            w = r - l + 1
            ans = max(ans, min(w, numOperations))
        return ans
```

### Go

```go
package main

import (
 "sort"
)

func maxFrequency(nums []int, k int, numOperations int) int {
 n := len(nums)
 if n == 0 {
  return 0
 }
 sort.Ints(nums)

 // frequency map
 freq := make(map[int]int, n)
 for _, x := range nums {
  freq[x]++
 }
 ans := 1

 // Case A
 for v, already := range freq {
  lowVal := v - k
  highVal := v + k
  L := lowerBound(nums, lowVal)
  R := upperBound(nums, highVal)
  totalInRange := R - L
  need := totalInRange - already
  canFix := need
  if canFix > numOperations {
   canFix = numOperations
  }
  if already+canFix > ans {
   ans = already + canFix
  }
 }

 // Case B: sliding window with 2*k
 l := 0
 for r := 0; r < n; r++ {
  for l <= r && nums[r]-nums[l] > 2*k {
   l++
  }
  w := r - l + 1
  cand := w
  if cand > numOperations {
   cand = numOperations
  }
  if cand > ans {
   ans = cand
  }
 }
 return ans
}

func lowerBound(a []int, target int) int {
 l, r := 0, len(a)
 for l < r {
  m := (l + r) / 2
  if a[m] < target {
   l = m + 1
  } else {
   r = m
  }
 }
 return l
}

func upperBound(a []int, target int) int {
 l, r := 0, len(a)
 for l < r {
  m := (l + r) / 2
  if a[m] <= target {
   l = m + 1
  } else {
   r = m
  }
 }
 return l
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

I'll explain the logic in simple steps, mapping to the code blocks above.

1. **Sort `nums`.**

   * This groups equal numbers and makes intervals contiguous. In code: `sort(nums)` (or `Arrays.sort`, `.sort()` etc).

2. **Count existing frequencies.**

   * I count how many times each value appears (`freq`). These are free contributors to any target equal to that value.

3. **Case A — existing value as target**

   * For a chosen `v` (each distinct value in `freq`):

     * `L = lower_bound(nums, v - k)` gives the first index `>= v-k`.
     * `R = upper_bound(nums, v + k)` gives first index `> v+k`.
     * `totalInRange = R - L` is how many numbers lie in `[v-k, v+k]`.
     * `already = freq[v]` are already equal to `v`.
     * `need = totalInRange - already` is how many need an operation.
     * `canFix = min(need, numOperations)` is how many we can actually convert.
     * Candidate frequency: `already + canFix`. Update `ans`.

   * This step is implemented across languages with binary search (`lower_bound`/`upper_bound`), or `bisect` in Python.

4. **Case B — meeting at new value (2*k window)**

   * Two pointers `l` and `r` maintain window with `nums[r] - nums[l] <= 2*k`.
   * Window size `w = r - l + 1` means the elements inside can meet at some value (not necessarily existing).
   * Candidate = `min(w, numOperations)` (you cannot convert more elements than `numOperations` if none are already equal; case A handles existing duplicates). Update `ans`.

5. **Return `ans`** after checking all candidates.

**Why both cases?**

* Some optimal solutions rely on an existing value as the final value (Case A).
* Others require creating a new meeting value (Case B), for example when no duplicates exist but elements are pairwise within `2*k` so they can meet halfway — a case that breaks naive solutions (e.g., `nums = [5, 64], k = 42, numOperations = 2`).

---

## Examples

1. Example from prompt-like tests:

   * Input: `nums = [1,4,5], k = 1, numOperations = 2`

     * Output: `2`
     * Explanation: we can make two `4`s (convert `1` -> `4` by adding 3 not allowed because of k=1? this is a simplified example; main logic uses intervals).

2. Edge case that failed naive approach:

   * Input: `nums = [5,64], k = 42, numOperations = 2`

     * Output: `2` (we can move both to some common value, because `64 - 5 = 59 <= 84 = 2*k`).

3. If array has duplicates:

   * Input: `nums = [2,2,4,4,9], k = 2, numOperations = 1`

     * We can pick existing target `4` and convert one `2` to `4` (if within `k`) or pick other possibilities — algorithm checks both cases.

---

## How to use / Run locally

**C++**

1. Put the `Solution` class in a `.cpp` file. Add main and parse input if you want to run locally.
2. Compile: `g++ -std=c++17 solution.cpp -O2 -o solution`
3. Run: `./solution`

**Java**

1. Put the `Solution` class in `Solution.java`. Add a `main` wrapper to call `maxFrequency`.
2. Compile: `javac Solution.java`
3. Run: `java Solution`

**JavaScript (Node)**

1. Put code into `solution.js` and export a helper or add a driver.
2. Run: `node solution.js`

**Python3**

1. Put `class Solution` code into `solution.py`, add a driver or run via an online judge (LeetCode).
2. Run: `python3 solution.py`

**Go**

1. Put code in `main.go` and call `maxFrequency` from `main`.
2. Run: `go run main.go`

(For LeetCode, just paste the language-specific `Solution` class/function into the editor and run tests.)

---

## Notes & Optimizations

* Using the distinct sorted values (iterating over distinct `v`) prevents duplicate work in Case A.
* We used binary search for the `[v-k, v+k]` interval queries — efficient and simple.
* You can optimize further by:

  * Using prefix counts for faster `totalInRange` queries (prefix sums over counts) so you avoid repeated binary searches.
  * Iterating distinct values by scanning the sorted array and performing range queries with two pointers combined with frequency counts — this may reduce constant factors and be linear after sorting.
* The combined approach (Case A + Case B) ensures we don't miss solutions that require meeting at a new value.

---

## Author

* [Md. Aarzoo Islam](https://bento.me/withaarzoo)
