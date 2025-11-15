# Problem Title

**3234. Count the Number of Substrings With Dominant Ones**

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

  * [C++](#c++)
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

I am given a binary string `s` (characters `'0'` and `'1'`). I must count the number of substrings such that the number of ones in the substring is **greater than or equal to** the **square** of the number of zeros in that substring.

Formally, a substring qualifies if:

```
ones_in_substring >= (zeros_in_substring)^2
```

Return the count of such substrings.

---

## Constraints

* `1 <= s.length <= 4 * 10^4` (typical LeetCode limits)
* `s` contains only `'0'` and `'1'`

Important optimization observation:

* If a substring has `k` zeros, it requires at least `k^2` ones. When `k^2 > n` (where `n = s.length`), it's impossible to satisfy. Hence only `k` up to `floor(sqrt(n))` need to be considered.

---

## Intuition

I thought about the constraint `ones >= zeros^2`. The square grows quickly, so:

* Any substring with **0 zeros** (i.e., all ones) always satisfies the condition because `ones >= 0`.
* For `k` zeros, `k` cannot be large — I only need to consider `k` up to `sqrt(n)`.
* For a specific group of `k` zeros (consecutive indexes from `zeroPos[i]` to `zeroPos[i+k-1]`), the substring can be formed by adding `x` ones on the left (0..leftOnes) and `y` ones on the right (0..rightOnes). Counting valid `(x,y)` pairs reduces to counting all pairs minus those where `x + y < t` (t derived from required length). That subtraction can be computed in O(1) using arithmetic sums.

This gives a fast solution that avoids enumerating all `O(n^2)` substrings.

---

## Approach

1. Count substrings that have 0 zeros (consecutive runs of ones): for a run of `t` ones add `t*(t+1)/2`.
2. Collect positions of zeros `zeroPos`.
3. For `k` from `1` to `floor(sqrt(n))`:

   * Slide a window of size `k` over `zeroPos` (take every block of `k` consecutive zeros).
   * For each block:

     * `leftOnes`: count of contiguous ones to the left we can include (stopping before previous zero or start).
     * `rightOnes`: same on the right (stopping before next zero or end).
     * `baseLen`: minimal length that covers the k zeros (positions difference + 1).
     * minimal substring length to satisfy ones >= k^2 is `needLen = k^2 + k` (since `ones = length - k`).
     * define `t = needLen - baseLen`. We need `x + y >= t`.
     * number of all possible pairs is `(leftOnes + 1) * (rightOnes + 1)`.
     * subtract the number of pairs with `x + y < t`. Compute that in closed form via arithmetic sums (O(1)).
4. Sum counts from step 1 and step 3 — that is the answer.

This approach is `O(n * sqrt(n))` time and `O(m)` extra space (where `m` is number of zeros).

---

## Data Structures Used

* `vector<int>` / `ArrayList<Integer>` / list / slice to store indices of zeros.
* Simple integer counters and arithmetic (64-bit accumulation to avoid overflow during intermediate calculations).

---

## Operations & Behavior Summary

* Run-length counting for all-ones substrings.
* Sliding windows over zero indices for small `k` values.
* Arithmetic combinatorics to count pairs `(leftExtension, rightExtension)` satisfying a linear inequality.

---

## Complexity

* **Time Complexity:** `O(n * sqrt(n))` where `n` = length of `s`.

  * We only loop `k` up to `sqrt(n)`. For each `k`, we iterate windows over zero indices. Each window is processed in O(1) arithmetic.
* **Space Complexity:** `O(m)` for storing zero positions, where `m` ≤ `n`.

---

## Multi-language Solutions

---

### C++

```c++
#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    int numberOfSubstrings(string s) {
        int n = (int)s.size();
        long long ans = 0;
        // count all-ones substrings
        long long run = 0;
        for (char c : s) {
            if (c == '1') run++;
            else { ans += run * (run + 1) / 2; run = 0; }
        }
        ans += run * (run + 1) / 2;

        // collect zero positions
        vector<int> zeroPos;
        for (int i = 0; i < n; ++i) if (s[i] == '0') zeroPos.push_back(i);
        int m = (int)zeroPos.size();
        if (m == 0) return (int)ans;

        int K = (int)floor(sqrt(n));
        for (int k = 1; k <= K && k <= m; ++k) {
            for (int i = 0; i + k - 1 < m; ++i) {
                int leftPrev = (i == 0 ? -1 : zeroPos[i-1]);
                int rightNext = (i + k - 1 == m - 1 ? n : zeroPos[i + k]);
                int leftOnes = zeroPos[i] - leftPrev - 1;
                int rightOnes = rightNext - zeroPos[i + k - 1] - 1;
                int baseLen = zeroPos[i + k - 1] - zeroPos[i] + 1;
                long long needLen = 1LL * k * k + k;
                long long t = needLen - baseLen; // need x+y >= t
                long long totalPairs = 1LL * (leftOnes + 1) * (rightOnes + 1);
                if (t <= 0) { ans += totalPairs; continue; }

                // compute pairs with x+y < t in closed form
                long long pairs_lt = 0;
                long long s0 = t - 1;
                if (s0 >= 0) {
                    long long L = leftOnes, R = rightOnes;
                    long long x_max = min(L, s0);
                    if (x_max >= 0) {
                        long long x0 = max(0LL, s0 - R);
                        if (x0 > x_max) {
                            pairs_lt = (x_max + 1) * (R + 1);
                        } else {
                            long long part1 = x0 * (R + 1);
                            long long n2 = x_max - x0 + 1;
                            long long sum_x = (x0 + x_max) * n2 / 2;
                            long long part2 = n2 * (s0 + 1) - sum_x;
                            pairs_lt = part1 + part2;
                        }
                    }
                }
                long long valid = totalPairs - pairs_lt;
                if (valid > 0) ans += valid;
            }
        }
        return (int)ans;
    }
};
```

---

### Java

```java
import java.util.*;

class Solution {
    public int numberOfSubstrings(String s) {
        int n = s.length();
        long ans = 0;
        // count all-ones substrings
        long run = 0;
        for (int i = 0; i < n; ++i) {
            if (s.charAt(i) == '1') run++;
            else { ans += run * (run + 1) / 2; run = 0; }
        }
        ans += run * (run + 1) / 2;

        // collect zero positions
        ArrayList<Integer> zeroPos = new ArrayList<>();
        for (int i = 0; i < n; ++i) if (s.charAt(i) == '0') zeroPos.add(i);
        int m = zeroPos.size();
        if (m == 0) return (int)ans;

        int K = (int)Math.floor(Math.sqrt(n));
        for (int k = 1; k <= K && k <= m; ++k) {
            for (int i = 0; i + k - 1 < m; ++i) {
                int leftPrev = (i == 0 ? -1 : zeroPos.get(i - 1));
                int rightNext = (i + k - 1 == m - 1 ? n : zeroPos.get(i + k));
                int leftOnes = zeroPos.get(i) - leftPrev - 1;
                int rightOnes = rightNext - zeroPos.get(i + k - 1) - 1;
                int baseLen = zeroPos.get(i + k - 1) - zeroPos.get(i) + 1;
                long needLen = 1L * k * k + k;
                long t = needLen - baseLen;
                long totalPairs = 1L * (leftOnes + 1) * (rightOnes + 1);
                if (t <= 0) { ans += totalPairs; continue; }

                long pairs_lt = 0;
                long s0 = t - 1;
                if (s0 >= 0) {
                    long L = leftOnes;
                    long R = rightOnes;
                    long x_max = Math.min(L, s0);
                    if (x_max >= 0) {
                        long x0 = Math.max(0L, s0 - R);
                        if (x0 > x_max) {
                            pairs_lt = (x_max + 1) * (R + 1);
                        } else {
                            long part1 = x0 * (R + 1);
                            long n2 = x_max - x0 + 1;
                            long sum_x = (x0 + x_max) * n2 / 2;
                            long part2 = n2 * (s0 + 1) - sum_x;
                            pairs_lt = part1 + part2;
                        }
                    }
                }
                long valid = totalPairs - pairs_lt;
                if (valid > 0) ans += valid;
            }
        }
        return (int)ans;
    }
}
```

---

### JavaScript

```javascript
/**
 * @param {string} s
 * @return {number}
 */
var numberOfSubstrings = function(s) {
    const n = s.length;
    let ans = 0;
    // all-ones substrings
    let run = 0;
    for (let i = 0; i < n; ++i) {
        if (s[i] === '1') run++;
        else { ans += run * (run + 1) / 2; run = 0; }
    }
    ans += run * (run + 1) / 2;

    // positions of zeros
    const zeroPos = [];
    for (let i = 0; i < n; ++i) if (s[i] === '0') zeroPos.push(i);
    const m = zeroPos.length;
    if (m === 0) return ans;

    const K = Math.floor(Math.sqrt(n));
    for (let k = 1; k <= K && k <= m; ++k) {
        for (let i = 0; i + k - 1 < m; ++i) {
            const leftPrev = (i === 0 ? -1 : zeroPos[i - 1]);
            const rightNext = (i + k - 1 === m - 1 ? n : zeroPos[i + k]);
            const leftOnes = zeroPos[i] - leftPrev - 1;
            const rightOnes = rightNext - zeroPos[i + k - 1] - 1;
            const baseLen = zeroPos[i + k - 1] - zeroPos[i] + 1;
            const needLen = k * k + k;
            const t = needLen - baseLen;
            const totalPairs = (leftOnes + 1) * (rightOnes + 1);
            if (t <= 0) { ans += totalPairs; continue; }

            let pairs_lt = 0;
            let s0 = t - 1;
            if (s0 >= 0) {
                let L = leftOnes, R = rightOnes;
                let x_max = Math.min(L, s0);
                if (x_max >= 0) {
                    let x0 = Math.max(0, s0 - R);
                    if (x0 > x_max) pairs_lt = (x_max + 1) * (R + 1);
                    else {
                        let part1 = x0 * (R + 1);
                        let n2 = x_max - x0 + 1;
                        let sum_x = (x0 + x_max) * n2 / 2;
                        let part2 = n2 * (s0 + 1) - sum_x;
                        pairs_lt = part1 + part2;
                    }
                }
            }
            const valid = totalPairs - pairs_lt;
            if (valid > 0) ans += valid;
        }
    }
    return ans;
};
```

---

### Python3

```python
class Solution:
    def numberOfSubstrings(self, s: str) -> int:
        n = len(s)
        ans = 0
        # count all-ones substrings
        run = 0
        for ch in s:
            if ch == '1':
                run += 1
            else:
                ans += run * (run + 1) // 2
                run = 0
        ans += run * (run + 1) // 2

        zeroPos = [i for i, ch in enumerate(s) if ch == '0']
        m = len(zeroPos)
        if m == 0:
            return ans

        import math
        K = int(math.isqrt(n))
        for k in range(1, K + 1):
            if k > m: break
            for i in range(0, m - k + 1):
                leftPrev = -1 if i == 0 else zeroPos[i - 1]
                rightNext = n if (i + k - 1 == m - 1) else zeroPos[i + k]
                leftOnes = zeroPos[i] - leftPrev - 1
                rightOnes = rightNext - zeroPos[i + k - 1] - 1
                baseLen = zeroPos[i + k - 1] - zeroPos[i] + 1
                needLen = k * k + k
                t = needLen - baseLen
                totalPairs = (leftOnes + 1) * (rightOnes + 1)
                if t <= 0:
                    ans += totalPairs
                    continue

                pairs_lt = 0
                s0 = t - 1
                if s0 >= 0:
                    L = leftOnes
                    R = rightOnes
                    x_max = min(L, s0)
                    if x_max >= 0:
                        x0 = max(0, s0 - R)
                        if x0 > x_max:
                            pairs_lt = (x_max + 1) * (R + 1)
                        else:
                            part1 = x0 * (R + 1)
                            n2 = x_max - x0 + 1
                            sum_x = (x0 + x_max) * n2 // 2
                            part2 = n2 * (s0 + 1) - sum_x
                            pairs_lt = part1 + part2
                valid = totalPairs - pairs_lt
                if valid > 0:
                    ans += valid
        return ans
```

---

### Go

```go
package main

import (
 "math"
)

func numberOfSubstrings(s string) int {
 n := len(s)
 var ans int64 = 0
 // all-ones substrings
 var run int64 = 0
 for i := 0; i < n; i++ {
  if s[i] == '1' { run++ }
  else { ans += run * (run + 1) / 2; run = 0 }
 }
 ans += run * (run + 1) / 2

 // zero positions
 zeroPos := make([]int, 0)
 for i := 0; i < n; i++ {
  if s[i] == '0' { zeroPos = append(zeroPos, i) }
 }
 m := len(zeroPos)
 if m == 0 {
  return int(ans)
 }

 K := int(math.Floor(math.Sqrt(float64(n))))
 for k := 1; k <= K && k <= m; k++ {
  for i := 0; i + k - 1 < m; i++ {
   leftPrev := -1
   if i != 0 { leftPrev = zeroPos[i-1] }
   rightNext := n
   if i + k - 1 != m-1 { rightNext = zeroPos[i+k] }
   leftOnes := zeroPos[i] - leftPrev - 1
   rightOnes := rightNext - zeroPos[i+k-1] - 1
   baseLen := zeroPos[i+k-1] - zeroPos[i] + 1
   needLen := k*k + k
   t := needLen - baseLen
   totalPairs := int64((leftOnes + 1) * (rightOnes + 1))
   if t <= 0 { ans += totalPairs; continue }

   var pairs_lt int64 = 0
   s0 := int64(t - 1)
   if s0 >= 0 {
    L := int64(leftOnes)
    R := int64(rightOnes)
    x_max := L
    if s0 < x_max { x_max = s0 }
    if x_max >= 0 {
     x0 := int64(0)
     if s0 - R > 0 { x0 = s0 - R }
     if x0 > x_max {
      pairs_lt = (x_max + 1) * (R + 1)
     } else {
      part1 := x0 * (R + 1)
      n2 := x_max - x0 + 1
      sum_x := (x0 + x_max) * n2 / 2
      part2 := n2 * (s0 + 1) - sum_x
      pairs_lt = part1 + part2
     }
    }
   }
   valid := totalPairs - pairs_lt
   if valid > 0 { ans += valid }
  }
 }
 return int(ans)
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

I will explain the important blocks and logic that are common across languages. I’ll use simple language and small steps so I can teach a friend.

1. **Count all-ones substrings**

   * Iterate over the string and keep a `run` counter that increases when we see `'1'`.
   * When the run ends (we see `'0'` or the string ends), add `run * (run + 1) / 2` to the answer, because a run of `t` ones contains `t*(t+1)/2` substrings (all of which have zero zeros, so they qualify).
   * Reset `run` and continue.

2. **Collect zero positions**

   * Build an array/list `zeroPos` of indices `i` where `s[i] == '0'`.
   * If `zeroPos` is empty, we already counted all valid substrings (they were all-one runs). Return the accumulated answer.

3. **Limit `k` to sqrt(n)**

   * For the number of zeros `k` in a substring, we need at least `k^2` ones.
   * If `k^2 > n`, no substring can satisfy this. So `k` can go only up to `floor(sqrt(n))`.

4. **Enumerate blocks of `k` zeros**

   * Slide a window of size `k` over the `zeroPos` indices. For each starting index `i`, the block covers zeros at `zeroPos[i]` through `zeroPos[i+k-1]`.
   * The substring that contains exactly these zeros and no extra zero is determined by:

     * Extending `x` ones to the left (0..leftOnes)
     * Extending `y` ones to the right (0..rightOnes)
   * `leftOnes` is `zeroPos[i] - previous_zero_index - 1` (or `zeroPos[i] - (-1) - 1` if none).
   * `rightOnes` is `next_zero_index - zeroPos[i+k-1] - 1` (or `n - zeroPos[i+k-1] - 1` if none).

5. **Length calculations and requirement**

   * `baseLen` = `zeroPos[i+k-1] - zeroPos[i] + 1` is the minimal length covering the k zeros (no extension).
   * Required minimal total length `>= k^2 + k` because `ones = length - k` and we need `ones >= k^2`.
   * Let `t = needLen - baseLen`. We require `x + y >= t` (if `t <= 0`, any `(x,y)` works).

6. **Counting valid (x, y) pairs**

   * Total possible pairs = `(leftOnes + 1) * (rightOnes + 1)`.
   * Pairs that *fail* the requirement are those with `x + y <= t - 1`.
   * Count failing pairs using arithmetic (O(1)):

     * Let `s0 = t - 1`.
     * For `x` in `[0..min(leftOnes, s0)]`:

       * `y` can be `0..min(rightOnes, s0 - x)`.
       * The sum over `x` can be split into two parts:

         * `x` where `s0 - x >= rightOnes` (i.e., all `y` are possible): contributes `(rightOnes+1)` each.
         * `x` where `s0 - x < rightOnes`: contributes `(s0 - x + 1)` each.
     * This sum simplifies into closed-form arithmetic sums (prefix sums of consecutive integers).

7. **Add valid counts**

   * Valid pairs = `totalPairs - pairs_lt`.
   * Add `valid` to the answer.

8. **Return final answer**

   * Final answer fits in 64-bit during counting, and is cast to the expected return type (usually `int` for many judges). Internally use 64-bit accumulators to be safe.

---

## Examples

Example 1:

```
Input: s = "00011"
Output: 5
Explanation: The substrings meeting ones >= zeros^2 are:
Indices (3,3), (4,4), (2,3), (3,4), (2,4)  (1-based in problem statement).
```

Example 2:

```
Input: s = "101101"
Output: 16
Explanation: Total 21 substrings, 5 have non-dominant ones, so 16 dominant ones.
```

You can test with these and random inputs.

---

## How to use / Run locally

### C++

Compile and run (if you put driver code):

```bash
g++ -std=c++17 -O2 solution.cpp -o solution
./solution
```

(LeetCode uses the `Solution` class — include driver code to test locally.)

### Java

Compile and run:

```bash
javac Solution.java
java Solution
```

(If you use LeetCode, submit class `Solution` with the method `public int numberOfSubstrings(String s)`.)

### JavaScript (Node)

Create a file `solution.js` that exports/uses the function, and run:

```bash
node solution.js
```

### Python3

Run:

```bash
python3 solution.py
```

(Place the `Solution` class and a driver `if __name__ == "__main__":` block to test locally.)

### Go

Build and run:

```bash
go run solution.go
```

---

## Notes & Optimizations

* I replaced the naive inner loop of iterating possible `x` values (which could cause extra work) with a closed-form arithmetic calculation for the number of `(x, y)` pairs where `x + y < t`. This reduced constant factors and avoids nested small loops.
* Using `sqrt(n)` upper bound for `k` drastically reduces work when string length is large, because the `k` dimension becomes small.
* Use 64-bit accumulators (`long long`, `long`, `int64`) to avoid overflow in intermediate multiplications (e.g., `(leftOnes+1)*(rightOnes+1)`).
* Returning `int` (when the judge expects `int`) avoids compile-time mismatch errors. Internally we still use `long`/`int64` to accumulate.
* Even with all optimizations, actual runtime depends on the judge's environment — I cannot guarantee `0 ms` or `100% percentile`. But these changes minimize work and are competitive.

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
