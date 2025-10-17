# Maximize the Number of Partitions After Operations (LeetCode 3003)

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
* [Step-by-step Detailed Explanation](#step-by-step-detailed-explanation)
* [Examples](#examples)
* [How to use / Run locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

---

## Problem Summary

We are given a lowercase string `s` and an integer `k`. We may change **at most one** character in `s` to any lowercase letter. After that, we repeatedly:

1. Choose the longest *prefix* of `s` that contains at most `k` distinct characters.
2. Delete that prefix (this counts as one partition).
3. Repeat until `s` becomes empty.

Return the **maximum number of partitions** we can get after optimally choosing at most one index to change.

---

## Constraints

* `1 <= s.length <= 10^4`
* `s` consists only of lowercase English letters.
* `1 <= k <= 26`

---

## Intuition

I thought about how partitions form: scanning `s` left-to-right, whenever adding a new character would make the current segment have more than `k` distinct letters, we must cut and start a new partition. The single allowed character change can be used at the best possible moment to avoid such cuts — or to make future cuts more beneficial.

So I represented the current set of letters in the active partition as a **bitmask** (26 bits), and used **DP + memoization** to explore decisions:

* Continue adding the current char (if possible).
* Or cut (start a new partition) if necessary.
* If the change is still available, try changing the current character to any of 26 letters and take the max.

Representing sets as bitmasks makes operations (union, bitcount) fast and compact. Memoization prevents recomputing overlapping subproblems.

---

## Approach

1. Precompute mask for each character: `mask_i = 1 << (s[i] - 'a')`.
2. Define DP: `dp(i, canChange, mask)` returns maximum number of completed partitions from index `i` to end, where:

   * `mask` is the bitmask of distinct letters currently in the active partition,
   * `canChange` is `1` if we still can change one character, else `0`.
3. At position `i`:

   * `mask2 = mask | mask_i`.

     * If `bitcount(mask2) > k` → we must start a new partition: `1 + dp(i+1, canChange, mask_i)`.
     * Else → `dp(i+1, canChange, mask2)`.
   * If `canChange == 1` → try changing current char to each letter `j` in `0..25`:

     * `changeMask = mask | (1 << j)`, and evaluate similarly; change usage sets `canChange=0`.
4. Use memoization (cache keyed by `(i, canChange, mask)`) to avoid recomputation.
5. Final answer is `dp(0, 1, 0) + 1` — we add the final active partition.

---

## Data Structures Used

* **Bitmask (int / long long)** — to represent up to 26 letters present in the current partition.
* **Memoization map** — map from state `(i, canChange, mask)` to `int` result.
* Arrays to precompute character masks.

---

## Operations & Behavior Summary

* `mask | (1 << bit)` — add a letter to the current set.
* `bitcount(mask)` — count distinct letters in a set (fast in all languages via built-ins or bit tricks).
* The DP explores both: use/change or not use/change at each position.
* Once `bitcount > k` we forcedly close current partition and start one with only current char.

---

## Complexity

* **Time Complexity:** roughly `O(n * 26 * states)` where `states` denotes the number of distinct `(i, mask, canChange)` combinations that occur in practice. In effect the algorithm behaves ≈ `O(26 * n)` amortized due to pruning and the fact masks evolve in restricted ways. This is efficient for `n ≤ 10^4`.
* **Space Complexity:** `O(n * number_of_masks_seen)` for memo storage; masks are at most `2^26` theoretically, but in practice only a small fraction appear; memory fits within practical limits for given constraints.

---

## Multi-language Solutions

### C++

```c++
#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    string s;
    int k;
    unordered_map<long long,int> memo;

    int dp(int i, long long mask, bool canChange) {
        if (i == (int)s.size()) return 0;
        long long key = ( (long long)i << 33 ) ^ (mask << 1) ^ (canChange ? 1LL : 0LL);
        if (memo.count(key)) return memo[key];

        int bit = s[i] - 'a';
        long long newMask = mask | (1LL << bit);
        int best = 0;

        if (__builtin_popcountll(newMask) > k) {
            best = 1 + dp(i + 1, (1LL << bit), canChange);
        } else {
            best = dp(i + 1, newMask, canChange);
        }

        if (canChange) {
            for (int j = 0; j < 26; ++j) {
                long long cm = mask | (1LL << j);
                if (__builtin_popcountll(cm) > k) {
                    best = max(best, 1 + dp(i + 1, (1LL << j), false));
                } else {
                    best = max(best, dp(i + 1, cm, false));
                }
            }
        }
        return memo[key] = best;
    }

    int maxPartitionsAfterOperations(string s_, int k_) {
        s = move(s_);
        k = k_;
        memo.clear();
        return dp(0, 0, true) + 1;
    }
};
```

---

### Java

```java
import java.util.*;

class Solution {
    private String s;
    private int k;
    private Map<Long, Integer> memo = new HashMap<>();

    private int dp(int i, long mask, boolean canChange) {
        if (i == s.length()) return 0;
        long key = ((long)i << 33) ^ (mask << 1) ^ (canChange?1L:0L);
        if (memo.containsKey(key)) return memo.get(key);

        int ch = s.charAt(i) - 'a';
        long newMask = mask | (1L << ch);
        int best;
        if (Long.bitCount(newMask) > k) {
            best = 1 + dp(i + 1, 1L << ch, canChange);
        } else {
            best = dp(i + 1, newMask, canChange);
        }

        if (canChange) {
            for (int j = 0; j < 26; ++j) {
                long cm = mask | (1L << j);
                if (Long.bitCount(cm) > k) {
                    best = Math.max(best, 1 + dp(i + 1, 1L << j, false));
                } else {
                    best = Math.max(best, dp(i + 1, cm, false));
                }
            }
        }

        memo.put(key, best);
        return best;
    }

    public int maxPartitionsAfterOperations(String s_, int k_) {
        this.s = s_;
        this.k = k_;
        memo.clear();
        return dp(0, 0L, true) + 1;
    }
}
```

---

### JavaScript

```javascript
/**
 * @param {string} s
 * @param {number} k
 * @return {number}
 */
var maxPartitionsAfterOperations = function(s, k) {
  const n = s.length;
  const memo = new Map();

  function keyFor(i, mask, canChange) {
    return `${i},${mask},${canChange}`;
  }

  function popcount(x) {
    // x fits within 32-bit for 26 bits
    x = x >>> 0;
    let cnt = 0;
    while (x) {
      x &= (x - 1);
      ++cnt;
    }
    return cnt;
  }

  function dp(i, mask, canChange) {
    if (i === n) return 0;
    const key = keyFor(i, mask, canChange);
    if (memo.has(key)) return memo.get(key);

    const bit = s.charCodeAt(i) - 97;
    const newMask = mask | (1 << bit);
    let best = 0;

    if (popcount(newMask) > k) {
      best = 1 + dp(i + 1, (1 << bit), canChange);
    } else {
      best = dp(i + 1, newMask, canChange);
    }

    if (canChange) {
      for (let j = 0; j < 26; ++j) {
        const cm = mask | (1 << j);
        if (popcount(cm) > k) {
          best = Math.max(best, 1 + dp(i + 1, (1 << j), false));
        } else {
          best = Math.max(best, dp(i + 1, cm, false));
        }
      }
    }

    memo.set(key, best);
    return best;
  }

  return dp(0, 0, true) + 1;
};
```

---

### Python3

```python
from functools import cache

class Solution:
    def maxPartitionsAfterOperations(self, s: str, k: int) -> int:
        n = len(s)
        masks = [1 << (ord(c) - ord('a')) for c in s]

        @cache
        def dp(i: int, can_change: int, mask: int) -> int:
            if i == n:
                return 0
            m2 = mask | masks[i]
            if m2.bit_count() > k:
                ans = 1 + dp(i + 1, can_change, masks[i])
            else:
                ans = dp(i + 1, can_change, m2)

            if can_change:
                for j in range(26):
                    cm = mask | (1 << j)
                    if cm.bit_count() > k:
                        ans = max(ans, 1 + dp(i + 1, 0, 1 << j))
                    else:
                        ans = max(ans, dp(i + 1, 0, cm))
            return ans

        return dp(0, 1, 0) + 1
```

---

### Go

```go
package main

func max(a, b int) int {
    if a > b { return a }
    return b
}

func bitcount(x int) int {
    cnt := 0
    for x != 0 {
        x &= x - 1
        cnt++
    }
    return cnt
}

func maxPartitionsAfterOperations(s string, k int) int {
    n := len(s)
    type key struct{ i, can, mask int }
    memo := map[key]int{}

    var dp func(i, can, mask int) int
    dp = func(i, can, mask int) int {
        if i == n { return 0 }
        kk := key{i, can, mask}
        if v, ok := memo[kk]; ok { return v }

        bit := int(s[i] - 'a')
        newMask := mask | (1 << bit)
        var ans int
        if bitcount(newMask) > k {
            ans = 1 + dp(i+1, can, (1<<bit))
        } else {
            ans = dp(i+1, can, newMask)
        }

        if can == 1 {
            for j := 0; j < 26; j++ {
                cm := mask | (1 << j)
                if bitcount(cm) > k {
                    ans = max(ans, 1 + dp(i+1, 0, (1<<j)))
                } else {
                    ans = max(ans, dp(i+1, 0, cm))
                }
            }
        }
        memo[kk] = ans
        return ans
    }

    return dp(0, 1, 0) + 1
}
```

---

## Step-by-step Detailed Explanation

**Common mental model (across languages):**

1. I keep a mask representing distinct letters in the *current active partition*.
2. Reading `s[i]`, I try to put it into the current partition:

   * `mask2 = mask | mask(s[i])`
   * If `bitcount(mask2) <= k`: I continue with `mask2`.
   * Else: I must close the current partition (count `+1`), and start a new partition where the new mask is just `mask(s[i])`.
3. If I still have the one allowed change:

   * I also try changing `s[i]` to every letter `j` (0..25), and apply the same logic using `mask | (1 << j)`. Changing sets `canChange = 0`.
4. Use memoization keyed by `(i, canChange, mask)` to avoid recomputation.
5. Base case: once `i == n`, return 0 (no further partitions created). The top-level call adds `+1` to count the last active partition.

**Why `+1` at the end?**
`dp` counts **completed** partitions (those we cut already). At the end there's one active partition remaining (maybe empty if string empty). So we return `dp(0,1,0)+1`.

**Why bitmask?**
Bitmask operations are constant time, compact, and `bitcount` is extremely fast with processor instructions or built-in helpers.

**Why try 26 replacements when `canChange` is true?**
Because changing the current character to *any* letter could be optimal — we must probe all candidate replacements (26 choices). Memoization prevents explosion.

---

## Examples

1. Input: `s = "accca", k = 2`
   Output: `3`
   Explanation (one optimal change): change `s[2]` (0-based) `'c'` → `'b'` (or other) resulting in partitions of lengths that give 3 total partitions.

2. Input: `s = "aabaab", k = 3`
   Output: `1`
   Explanation: The whole string contains at most 3 distinct characters already.

3. Input: `s = "xxyz", k = 1`
   Output: `4`
   Explanation: Best change is to make some letter different to maximize partitions; each removal removes one char partition.

---

## How to use / Run locally

### C++

* Compile:

  ```bash
  g++ -std=c++17 -O2 -pipe solution.cpp -o solution
  ```
* Run (if you added a `main()` wrapper or test harness):

  ```bash
  ./solution
  ```

### Java

* Compile:

  ```bash
  javac Solution.java
  ```
* Run:

  ```bash
  java Solution
  ```

*(Wrap with `main` for local testing.)*

### JavaScript (Node)

* Save e.g. `solution.js` with the function and a test harness, then:

  ```bash
  node solution.js
  ```

### Python3

* Run:

  ```bash
  python3 solution.py
  ```

### Go

* Build and run:

  ```bash
  go run solution.go
  ```

---

## Notes & Optimizations

* Using a bitmask for 26 letters is the most important optimization here — all operations become O(1).
* Memoization prevents repeated state exploration; the effective state space is `O(n * small_number_of_masks_seen)`.
* In C++ / Java, using builtin popcount (`__builtin_popcountll`, `Long.bitCount`) is fast.
* Early prunings: if `k == 26`, the answer is always `1` because the entire string fits in one partition.
* The DP is top-down for clarity; bottom-up variants are possible but more complex.

---

## Author

[Md. Aarzoo Islam](https://bento.me/withaarzoo)
