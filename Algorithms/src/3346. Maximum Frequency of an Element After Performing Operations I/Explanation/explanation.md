# Problem Title

3346. Maximum Frequency of an Element After Performing Operations I

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
* [Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3)](#step-by-step-detailed-explanation-c-java-javascript-python3)
* [Examples](#examples)
* [How to use / Run locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

---

## Problem Summary

You are given an integer array `nums` and two integers `k` and `numOperations`.
You may perform at most `numOperations` operations. In a single operation you pick an index that hasn't been changed before and add any integer in range `[-k, k]` to `nums[i]`. After at most `numOperations` changes (each change applied to a distinct index), return the **maximum possible frequency** (count) of any integer in `nums`.

Put simply: each number `v` can end up at any integer inside `[v-k, v+k]`. Find an integer `t` and a set of at most `numOperations` indices (distinct) that make the frequency of `t` as large as possible.

---

## Constraints

(LeetCode constraints for this problem — used to justify the chosen approach)

* `1 <= nums.length <= 10^5`
* `1 <= nums[i] <= 10^5`
* `0 <= k <= 10^5`
* `0 <= numOperations <= nums.length`

These constraints mean `nums[i]` values are bounded by `10^5`, so an array-based counting/prefix approach over values up to `max(nums)+k+1` is feasible.

---

## Intuition

I thought about the problem by imagining each value `v` in `nums` producing an interval of possible targets `[v - k, v + k]`. If I pick a candidate target integer `t`, any `v` whose interval contains `t` **can** be turned into `t`. Among those, some are already equal to `t` (zero operations needed) and the rest can be changed using my limited `numOperations`. So for target `t`:

```
max possible freq at t = freq[t] + min(numOperations, total_covering_t - freq[t])
```

* `freq[t]` = how many elements already equal `t`.
* `total_covering_t` = number of `v` such that `t ∈ [v-k, v+k]`.

I realized I can compute `total_covering_t` for every candidate `t` quickly using a value-frequency array and its prefix sums, then check every `t` in the small range `0..max(nums)+k+1`.

---

## Approach

1. Find `mx = max(nums)`.
2. Build a frequency array `count` sized `mx + k + 2`.
3. For each `v` in `nums`, increment `count[v]`.
4. Turn `count` into prefix sums so that `pref[i] = number of elements ≤ i`.
5. For each integer `t` in `[0, mx + k + 1]`:

   * `L = max(0, t - k)`, `R = min(size - 1, t + k)`
   * `total = pref[R] - (L > 0 ? pref[L-1] : 0)` → how many elements can become `t`
   * `freq_t = count[t] - (t > 0 ? count[t-1] : 0)` → how many are already `t`
   * `candidate = freq_t + min(numOperations, total - freq_t)`
   * update answer with `candidate`.
6. Return the maximum answer found.

Why this works: for each `t` we count every element that is able (by interval) to reach `t`. We can convert at most `numOperations` of the non-`t` elements among those — so the formula above gives the best achievable for `t`. Checking all `t` in the bounded range guarantees correctness.

---

## Data Structures Used

* Integer arrays (value counts and prefix sums).
* Simple scalar variables for looping and tracking maxima.

No heavy or nested data structures — just arrays sized proportionally to `max(nums) + k + 2`.

---

## Operations & Behavior Summary

* Counting: `O(n)` to count frequency for each value in `nums`.
* Prefix building: `O(M)` where `M = mx + k + 2` to convert counts to prefix sums.
* Checking all candidates `t`: `O(M)` to compute `total`, `freq_t`, and `candidate`.
* Final answer is the maximum candidate across all `t`.

All array accesses are O(1). Memory is `O(M)`.

---

## Complexity

* **Time Complexity:** `O(n + M)` where `n = nums.length` and `M = max(nums) + k + 2`. For LeetCode constraints (`max(nums) ≤ 10^5`), this is practical.
* **Space Complexity:** `O(M)` for the `count`/prefix arrays.

---

## Multi-language Solutions

### C++

```c++
#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    int maxFrequency(vector<int>& nums, int k, int numOperations) {
        if (nums.empty()) return 0;
        int mx = *max_element(nums.begin(), nums.end());
        int size = mx + k + 2;               // safe limit to include t+k
        vector<int> count(size, 0);

        // frequency counts
        for (int v : nums) count[v]++;

        // prefix sums: count[i] becomes number of elements <= i
        for (int i = 1; i < size; ++i) count[i] += count[i-1];

        int ans = 0;
        for (int t = 0; t < size; ++t) {
            int L = max(0, t - k);
            int R = min(size - 1, t + k);
            int total = count[R] - (L > 0 ? count[L-1] : 0);   // elements that can become t
            int freq_t = (t > 0) ? (count[t] - count[t-1]) : count[t];
            int canConvert = total - freq_t;
            int take = min(numOperations, canConvert);
            ans = max(ans, freq_t + take);
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
        if (nums.length == 0) return 0;
        int mx = Arrays.stream(nums).max().getAsInt();
        int size = mx + k + 2;
        int[] count = new int[size];

        for (int v : nums) count[v]++;

        for (int i = 1; i < size; ++i) count[i] += count[i-1];

        int ans = 0;
        for (int t = 0; t < size; ++t) {
            int L = Math.max(0, t - k);
            int R = Math.min(size - 1, t + k);
            int total = count[R] - (L > 0 ? count[L - 1] : 0);
            int freq_t = (t > 0) ? (count[t] - count[t-1]) : count[t];
            int canConvert = total - freq_t;
            int take = Math.min(numOperations, canConvert);
            ans = Math.max(ans, freq_t + take);
        }
        return ans;
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
const maxFrequency = (nums, k, numOperations) => {
    if (nums.length === 0) return 0;
    let mx = Math.max(...nums);
    let size = mx + k + 2;
    const count = new Array(size).fill(0);

    for (const v of nums) count[v]++;

    for (let i = 1; i < size; ++i) count[i] += count[i-1];

    let ans = 0;
    for (let t = 0; t < size; ++t) {
        const L = Math.max(0, t - k);
        const R = Math.min(size - 1, t + k);
        const total = count[R] - (L > 0 ? count[L - 1] : 0);
        const freq_t = (t > 0) ? (count[t] - count[t - 1]) : count[t];
        const canConvert = total - freq_t;
        const take = Math.min(numOperations, canConvert);
        ans = Math.max(ans, freq_t + take);
    }
    return ans;
};
```

### Python3

```python3
class Solution:
    def maxFrequency(self, nums: List[int], k: int, numOperations: int) -> int:
        if not nums:
            return 0
        mx = max(nums)
        size = mx + k + 2
        count = [0] * size

        for v in nums:
            count[v] += 1

        # prefix sums
        for i in range(1, size):
            count[i] += count[i-1]

        ans = 0
        for t in range(size):
            L = max(0, t - k)
            R = min(size - 1, t + k)
            total = count[R] - (count[L-1] if L > 0 else 0)
            freq_t = count[t] - (count[t-1] if t > 0 else 0)
            canConvert = total - freq_t
            take = min(numOperations, canConvert)
            ans = max(ans, freq_t + take)
        return ans
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3)

I will explain the essential blocks mapped to the shared algorithm. The same logic is implemented in all languages; variable names are consistent.

1. **Handle empty input**

   * If `nums` is empty, return `0`.

2. **Find maximum and prepare array size**

   * `mx = max(nums)`.
   * `size = mx + k + 2`.

     * `+k` to allow targets up to `mx + k`.
     * `+2` makes indexing safe for prefix operations and `t+k` boundaries.

3. **Build frequency array**

   * For each `v` in `nums`: `count[v]++`.
   * After this, `count[x]` is how many times `x` occurs in `nums`.

4. **Build prefix sums**

   * Loop `i = 1..size-1`: `count[i] += count[i-1]`.
   * Now `count[i]` equals number of elements `<= i`.

5. **Iterate every candidate t**

   * For `t` in `[0..size-1]`:

     * `L = max(0, t - k)` and `R = min(size - 1, t + k)`.
     * `total = count[R] - (L > 0 ? count[L-1] : 0)` → number of elements whose `[v-k, v+k]` includes `t`.
     * `freq_t = count[t] - (t > 0 ? count[t-1] : 0)` → how many are already `t`.
     * `canConvert = total - freq_t` → how many not-equal elements can be converted to `t`.
     * `take = min(numOperations, canConvert)` → how many we actually convert.
     * `candidate = freq_t + take`. Update `ans = max(ans, candidate)`.

6. **Return `ans`**

This block-by-block mapping corresponds to the loops and operations in each language's code above.

---

## Examples

1. Example from prompt:

   * Input: `nums = [1,4,5]`, `k = 1`, `numOperations = 2`

     * Output: `2` (we can change values to make two equal).
2. Edge example:

   * Input: `nums = [88,53]`, `k = 27`, `numOperations = 2`

     * Output: `2` — both can be changed to `71` (or any overlapping integer).
3. Provided failing test earlier:

   * Input: `nums = [5,11,20,20]`, `k = 5`, `numOperations = 1`

     * Output: `2` — we can convert one of the values into another target within allowed intervals to make frequency 2.

You can try these to verify correctness.

---

## How to use / Run locally

1. **C++**

   * Create a file `solution.cpp` with the above class (wrap with a `main` driver to test).
   * Compile: `g++ -std=c++17 solution.cpp -O2 -o sol`
   * Run: `./sol` (depending on your driver).

2. **Java**

   * Put the `Solution` class into `Solution.java` and add a `main` method or test harness.
   * Compile: `javac Solution.java`
   * Run: `java Solution` (with driver).

3. **JavaScript**

   * Put the `maxFrequency` function into a file `solution.js` and test with Node.js:
   * Run: `node solution.js` (after adding tests that call `maxFrequency` and print results).

4. **Python3**

   * Create `solution.py`, include `class Solution` and a small driver:
   * Run: `python3 solution.py` or use the class in LeetCode editor.

**Note:** For quick tests, use the provided examples by creating a small driver that constructs `nums`, `k`, `numOperations`, calls the function, and prints the result.

---

## Notes & Optimizations

* This approach uses `O(M)` memory where `M = mx + k + 2`. With the LeetCode upper bound `mx ≤ 10^5` and `k ≤ 10^5`, memory is acceptable in contest/LeetCode settings. If `mx` were much larger (e.g., up to `10^9`), we would need a different approach (sweep-line with compressed coordinates or event sorting).
* Alternative approach (if values were huge): compress coordinates or use event-sorting + sweep-line over interval endpoints. That approach is `O(n log n)` and `O(n)` memory, and works when values are large but `n` is small.
* The prefix-array method is extremely fast in practice for bounded `nums[i]` because it uses contiguous memory and linear scans.
* Be careful with indexing and negative targets. We clamp `L` at `0` because `nums[i] >= 1` in constraints; if negatives are allowed in variations, shift coordinates or use hash-based compression.

---

## Author

* [Md. Aarzoo Islam](https://bento.me/withaarzoo)
