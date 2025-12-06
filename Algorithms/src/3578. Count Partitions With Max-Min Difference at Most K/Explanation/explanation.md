# 3578. Count Partitions With Max-Min Difference at Most K

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

We are given an integer array `nums` and an integer `k`.

We need to partition the array into **one or more non-empty contiguous segments** such that for **every segment**:

> `max(segment) - min(segment) <= k`

We must return the **total number of valid partitions**, modulo `1e9 + 7`.

---

## Constraints

* `2 <= nums.length <= 5 * 10^4`
* `1 <= nums[i] <= 10^9`
* `0 <= k <= 10^9`
* Answer must be given modulo `1_000_000_007`.

---

## Intuition

When I first looked at this problem, I thought:

* Partitions are always **contiguous**, so the order is fixed.
* For any partitioning, the last segment ends at the last index.
* If I fix the **end** of the last segment at index `r`, then I just need to know:

  * From which indices `j` (start positions) can I begin this last segment such that `[j..r]` is valid (`max - min <= k`)?
  * For each such `j`, the number of ways is equal to the number of ways to partition prefix `[0..j-1]`.

This naturally suggested a **DP on prefixes**, and to handle `max` / `min` quickly I thought of a **sliding window** with **monotonic deques**.

---

## Approach

1. Let `n = len(nums)`.

2. Define DP:

   * `dp[i]` = number of ways to partition the first `i` elements (`nums[0..i-1]`).
   * `dp[0] = 1` (empty prefix has exactly one “do nothing” partition).

3. For each position `r` (0-based index, last element of the last segment):

   * The prefix length is `i = r + 1`.
   * We want to consider all `j` such that `0 <= j <= r` and:

     * Segment `nums[j..r]` is valid (`max - min <= k`).
   * For each valid `j`, we add `dp[j]` to `dp[i]`.

4. To quickly know all valid `j` for each `r`, I use a **sliding window** `[l..r]`:

   * Maintain two deques:

     * `maxDeque`: indices in decreasing order of `nums[]`. Front is current max.
     * `minDeque`: indices in increasing order of `nums[]`. Front is current min.
   * For each new `r`, insert `nums[r]` into both deques while preserving monotonicity.
   * While `max - min > k`, move `l` to the right:

     * Pop from front of deques if that index leaves the window.
   * After this adjustment, `[l..r]` is the **smallest** valid window ending at `r`.
     So all valid `j` lie in range `[l..r]`.

5. The transition becomes:

   * Let `L = l` (smallest valid start for window ending at `r`).
   * `dp[i] = sum(dp[L] + dp[L+1] + ... + dp[r])`.

6. To compute this sum efficiently in O(1) per step, I keep a prefix sum array:

   * `pref[i] = dp[0] + dp[1] + ... + dp[i]` (mod `MOD`).
   * Then:

     * `dp[i] = pref[r] - (L > 0 ? pref[L-1] : 0)`.

7. Repeat for all `r` from `0` to `n-1`, and finally return `dp[n]`.

---

## Data Structures Used

* **Dynamic Programming Array**

  * `dp[i]` stores number of ways to partition first `i` elements.

* **Prefix Sum Array**

  * `pref[i]` = sum of `dp[0..i]` to answer range-sum queries instantly.

* **Two Deques (Double-ended Queues)**

  * `maxDeque`: keeps indices of current window in **decreasing** order of values → fast window maximum.
  * `minDeque`: keeps indices in **increasing** order of values → fast window minimum.

These structures make each element enter and leave the deques at most once → `O(n)` total operations.

---

## Operations & Behavior Summary

For each index `r`:

1. **Extend window to right**:

   * Insert `nums[r]` into both deques.
   * Update max and min of current window.

2. **Shrink window from left** until valid:

   * While `max(nums[l..r]) - min(nums[l..r]) > k`, increment `l`.
   * Remove outdated indices (`l`) from deques.

3. **Find valid start range**:

   * After shrinking, `[l..r]` is the smallest valid subarray.
   * All valid start indices `j` lie in `[l..r]`.

4. **DP update**:

   * Use prefix sums to compute sum of `dp[j]` for `j` from `l` to `r`.
   * Store this in `dp[r+1]` and update `pref[r+1]`.

---

## Complexity

* **Time Complexity:** `O(n)`

  * Each index is pushed and popped from both deques at most once.
  * DP and prefix operations are O(1) per index.
  * `n` = length of `nums`.

* **Space Complexity:** `O(n)`

  * `dp` and `pref` arrays of size `n+1`.
  * Deques may hold up to `n` indices in worst case.

---

## Multi-language Solutions

### C++

```c++
#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    int countPartitions(vector<int>& nums, int k) {
        const int MOD = 1'000'000'007;
        int n = (int)nums.size();

        // dp[i] = number of ways to partition prefix of length i (nums[0..i-1])
        vector<long long> dp(n + 1, 0), pref(n + 1, 0);
        dp[0] = 1;
        pref[0] = 1;

        deque<int> maxdq, mindq; // indices
        int l = 0;

        for (int r = 0; r < n; ++r) {
            int x = nums[r];

            // Maintain decreasing deque for max
            while (!maxdq.empty() && nums[maxdq.back()] <= x) {
                maxdq.pop_back();
            }
            maxdq.push_back(r);

            // Maintain increasing deque for min
            while (!mindq.empty() && nums[mindq.back()] >= x) {
                mindq.pop_back();
            }
            mindq.push_back(r);

            // Shrink left until window [l..r] is valid
            while (!maxdq.empty() && !mindq.empty() &&
                   (long long)nums[maxdq.front()] - nums[mindq.front()] > k) {
                if (maxdq.front() == l) maxdq.pop_front();
                if (mindq.front() == l) mindq.pop_front();
                ++l;
            }

            int L = l;     // minimal valid start index for this r
            int i = r + 1; // prefix length

            long long ways = pref[i - 1];  // sum dp[0..i-1]
            if (L > 0) ways -= pref[L - 1]; // subtract dp[0..L-1]

            ways %= MOD;
            if (ways < 0) ways += MOD;

            dp[i] = ways;
            pref[i] = (pref[i - 1] + dp[i]) % MOD;
        }

        return (int)dp[n];
    }
};
```

---

### Java

```java
class Solution {
    public int countPartitions(int[] nums, int k) {
        final int MOD = 1_000_000_007;
        int n = nums.length;

        long[] dp = new long[n + 1];
        long[] pref = new long[n + 1];

        dp[0] = 1;
        pref[0] = 1;

        Deque<Integer> maxdq = new ArrayDeque<>();
        Deque<Integer> mindq = new ArrayDeque<>();

        int l = 0;

        for (int r = 0; r < n; r++) {
            int x = nums[r];

            // Maintain decreasing deque for max
            while (!maxdq.isEmpty() && nums[maxdq.peekLast()] <= x) {
                maxdq.pollLast();
            }
            maxdq.offerLast(r);

            // Maintain increasing deque for min
            while (!mindq.isEmpty() && nums[mindq.peekLast()] >= x) {
                mindq.pollLast();
            }
            mindq.offerLast(r);

            // Shrink window while invalid
            while (!maxdq.isEmpty() && !mindq.isEmpty()
                    && (long)nums[maxdq.peekFirst()] - nums[mindq.peekFirst()] > k) {
                if (maxdq.peekFirst() == l) maxdq.pollFirst();
                if (mindq.peekFirst() == l) mindq.pollFirst();
                l++;
            }

            int L = l;
            int i = r + 1;

            long ways = pref[i - 1];
            if (L > 0) ways -= pref[L - 1];

            ways %= MOD;
            if (ways < 0) ways += MOD;

            dp[i] = ways;
            pref[i] = (pref[i - 1] + dp[i]) % MOD;
        }

        return (int)dp[n];
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
var countPartitions = function(nums, k) {
    const MOD = 1_000_000_007;
    const n = nums.length;

    const dp = new Array(n + 1).fill(0);
    const pref = new Array(n + 1).fill(0);

    dp[0] = 1;
    pref[0] = 1;

    const maxdq = []; // indices for max
    const mindq = []; // indices for min

    let l = 0;

    for (let r = 0; r < n; r++) {
        const x = nums[r];

        // Maintain decreasing deque for max
        while (maxdq.length > 0 && nums[maxdq[maxdq.length - 1]] <= x) {
            maxdq.pop();
        }
        maxdq.push(r);

        // Maintain increasing deque for min
        while (mindq.length > 0 && nums[mindq[mindq.length - 1]] >= x) {
            mindq.pop();
        }
        mindq.push(r);

        // Shrink window while invalid
        while (maxdq.length > 0 && mindq.length > 0 &&
               nums[maxdq[0]] - nums[mindq[0]] > k) {
            if (maxdq[0] === l) maxdq.shift();
            if (mindq[0] === l) mindq.shift();
            l++;
        }

        const L = l;
        const i = r + 1;

        let ways = pref[i - 1];
        if (L > 0) ways -= pref[L - 1];

        ways %= MOD;
        if (ways < 0) ways += MOD;

        dp[i] = ways;
        pref[i] = (pref[i - 1] + dp[i]) % MOD;
    }

    return dp[n];
};
```

---

### Python3

```python
from collections import deque
from typing import List

class Solution:
    def countPartitions(self, nums: List[int], k: int) -> int:
        MOD = 10**9 + 7
        n = len(nums)

        # dp[i]: ways to partition nums[0..i-1]
        dp = [0] * (n + 1)
        pref = [0] * (n + 1)

        dp[0] = 1
        pref[0] = 1

        maxdq = deque()  # indices with nums decreasing
        mindq = deque()  # indices with nums increasing
        l = 0

        for r in range(n):
            x = nums[r]

            # Maintain decreasing deque for max
            while maxdq and nums[maxdq[-1]] <= x:
                maxdq.pop()
            maxdq.append(r)

            # Maintain increasing deque for min
            while mindq and nums[mindq[-1]] >= x:
                mindq.pop()
            mindq.append(r)

            # Shrink window while invalid
            while maxdq and mindq and nums[maxdq[0]] - nums[mindq[0]] > k:
                if maxdq[0] == l:
                    maxdq.popleft()
                if mindq[0] == l:
                    mindq.popleft()
                l += 1

            L = l
            i = r + 1

            ways = pref[i - 1]
            if L > 0:
                ways -= pref[L - 1]
            ways %= MOD

            dp[i] = ways
            pref[i] = (pref[i - 1] + dp[i]) % MOD

        return dp[n]
```

---

### Go

```go
package main

func countPartitions(nums []int, k int) int {
 const MOD int64 = 1_000_000_007
 n := len(nums)

 dp := make([]int64, n+1)
 pref := make([]int64, n+1)

 dp[0] = 1
 pref[0] = 1

 maxdq := make([]int, 0) // indices for max
 mindq := make([]int, 0) // indices for min
 l := 0

 for r := 0; r < n; r++ {
  x := nums[r]

  // Maintain decreasing deque for max
  for len(maxdq) > 0 && nums[maxdq[len(maxdq)-1]] <= x {
   maxdq = maxdq[:len(maxdq)-1]
  }
  maxdq = append(maxdq, r)

  // Maintain increasing deque for min
  for len(mindq) > 0 && nums[mindq[len(mindq)-1]] >= x {
   mindq = mindq[:len(mindq)-1]
  }
  mindq = append(mindq, r)

  // Shrink while invalid
  for len(maxdq) > 0 && len(mindq) > 0 &&
   int64(nums[maxdq[0]]-nums[mindq[0]]) > int64(k) {
   if maxdq[0] == l {
    maxdq = maxdq[1:]
   }
   if mindq[0] == l {
    mindq = mindq[1:]
   }
   l++
  }

  L := l
  i := r + 1

  ways := pref[i-1]
  if L > 0 {
   ways -= pref[L-1]
  }

  ways %= MOD
  if ways < 0 {
   ways += MOD
  }

  dp[i] = ways
  pref[i] = (pref[i-1] + dp[i]) % MOD
 }

 return int(dp[n])
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The core logic is identical across all languages. I’ll describe it in a language-agnostic way; you can map it to whichever code you’re reading.

1. **Initialization**

   * Set `MOD = 1e9 + 7`.
   * `dp[0] = 1` because there is 1 way to partition an empty prefix.
   * `pref[0] = dp[0]`.
   * Create two empty deques: `maxdq`, `mindq`.
   * Left pointer `l = 0`.

2. **Loop over each index `r`** (0 to `n-1`):

   * Let `x = nums[r]`.

3. **Update deques with new element at `r`:**

   * For `maxdq`:

     * While back element’s value `<= x`, pop from back.
     * Push `r` at the back.
   * For `mindq`:

     * While back element’s value `>= x`, pop from back.
     * Push `r` at the back.

4. **Maintain validity of current window [l..r]:**

   * While `nums[maxdq.front()] - nums[mindq.front()] > k`:

     * If the front index of `maxdq` equals `l`, pop it.
     * If the front index of `mindq` equals `l`, pop it.
     * Increment `l` by 1.
   * After this loop:

     * `[l..r]` is the shortest valid segment ending at `r`.
     * All valid segment starts `j` must satisfy `l <= j <= r`.

5. **Compute DP transition for prefix of length `i = r + 1`:**

   * Let `L = l`.
   * We need `dp[i] = sum(dp[j]) for j = L..r`.
   * Using prefix sums:

     * `sum(dp[L..r]) = pref[r] - (L > 0 ? pref[L-1] : 0)`.
   * Apply modulo carefully and make sure result is non-negative.

6. **Update arrays:**

   * Set `dp[i]` to the computed `ways`.
   * Update `pref[i] = pref[i-1] + dp[i] (mod MOD)`.

7. **Final Answer**

   * After loop ends, return `dp[n]`.

This pattern is identical in C++, Java, JS, Python, and Go – only syntax changes.

---

## Examples

### Example 1

```text
Input:  nums = [9, 4, 1, 3, 7], k = 4
Output: 6
```

Valid partitions (one possible listing):

* [[9], [4, 1], [3], [7]]
* [[9], [4, 1, 3], [7]]
* [[9, 4], [1, 3], [7]]
* [[9, 4, 1], [3], [7]]
* [[9, 4, 1], [3, 7]]
* [[9, 4], [1, 3, 7]]

Each segment in these partitions satisfies `max(segment) - min(segment) <= 4`.

---

### Example 2

```text
Input:  nums = [3, 3, 4], k = 0
Output: 2
```

Valid partitions:

* [[3], [3], [4]]
* [[3, 3], [4]]

Here, only segments where all elements are equal are allowed, since `k = 0`.

---

## How to use / Run locally

Below are rough steps. Adjust based on your environment.

### C++

```bash
g++ -std=c++17 -O2 solution.cpp -o solution
./solution
```

Inside `solution.cpp`, include the class `Solution` and write a small `main()` that:

* Reads `n`, `k`
* Reads `nums`
* Calls `Solution().countPartitions(nums, k)`
* Prints the result.

---

### Java

```bash
javac Solution.java
java Solution
```

Make sure the file name is `Solution.java` and contains the `Solution` class.
Add a `main` method that constructs the `Solution` object and calls `countPartitions`.

---

### JavaScript (Node.js)

```bash
node solution.js
```

In `solution.js`, export or call the `countPartitions` function with some test input and print the result.

---

### Python3

```bash
python3 solution.py
```

In `solution.py`, paste the `Solution` class and at the bottom write a small test:

```python
if __name__ == "__main__":
    nums = [9, 4, 1, 3, 7]
    k = 4
    print(Solution().countPartitions(nums, k))
```

---

### Go

```bash
go run main.go
```

In `main.go`, keep the `countPartitions` function and a `main()` which reads input or uses test arrays and prints the value.

---

## Notes & Optimizations

* **Why O(n) is enough**
  A naive DP would try all possible starts for each end index → `O(n^2)`.
  Using the sliding window + deques, each index is processed a constant number of times, giving `O(n)` total.

* **Why prefix sums are needed**
  Without prefix sums, summing `dp[L..r]` each time would also be `O(n)` per index.
  Prefix sum converts range sum to O(1).

* **Memory optimization**
  `dp` and `pref` both require `O(n)` memory. They are needed for correct prefix computations.
  This is fine for the constraints (`n <= 5 * 10^4`).

* **Edge cases**

  * All elements equal and `k = 0` → every partition where segments are any contiguous blocks is valid.
  * Very large `k` (e.g., larger than `max(nums)-min(nums)`) → the whole array can be one single segment, but there are still many partitions. Our DP handles this naturally.
  * Pay attention to modulo operations, especially when subtracting (`add MOD` before `%` to keep positive).

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
