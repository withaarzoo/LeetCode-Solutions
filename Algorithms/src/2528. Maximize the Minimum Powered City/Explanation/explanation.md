# Maximize the Minimum Powered City (LeetCode 2528)

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

You’re given an array `stations` of length `n`, where `stations[i]` is the number of power stations in city `i`.
Each power station supplies power to every city within distance `r` (inclusive).
You can build **at most** `k` additional stations anywhere (each new station has the same range `r`).

**Goal:** Build up to `k` stations so that the **minimum power among all cities** is as large as possible. Return that maximum achievable minimum power.

“Power of a city” = sum of stations located in cities within `r` distance of it.

---

## Constraints

(typical constraints from the problem statement)

* `n == stations.length`
* `1 ≤ n ≤ 1e5`
* `0 ≤ stations[i] ≤ 1e5`
* `0 ≤ r ≤ n-1`
* `0 ≤ k ≤ 1e9`

---

## Intuition

If I guess a target value `T` for the **minimum power** every city must have, I can try to verify whether it’s possible by placing at most `k` new stations.
This verification is **monotonic**: if I can reach `T`, then I can also reach any `T' ≤ T`.
So I can **binary search** the answer.

To check feasibility for a given `T`, I sweep left to right. Whenever the current city’s power is below `T`, I add enough new stations as far right as possible (so their effect lasts longer for upcoming cities). With range `r`, that effectively means placing them so their coverage window ends at `i + 2r`. I apply these additions using a **difference array** so updates are O(1) and the pass is O(n).

---

## Approach

1. **Pre-compute base power per city** (existing stations only).

   * Use a “difference array” to add `+stations[i]` to `[i-r, i+r]` in O(1), then prefix-sum to get `base[i]`.
2. **Binary search** the answer `T`.

   * `lo = 0`, `hi = sum(stations) + k`. (You can’t exceed total stations after adding `k`.)
3. **Feasibility check** for a fixed `T`:

   * Maintain a second difference array `add[]` representing the contributions from newly placed stations.
   * Sweep `i = 0..n-1`:

     * `extra += add[i]`, `curr = base[i] + extra`.
     * If `curr < T`, we **need** `need = T - curr` extra stations.

       * Spend them: `used += need`. If `used > k`, **fail**.
       * Apply their effect now: `extra += need`.
       * Schedule their removal at `end = min(n, i + 2*r + 1)` via `add[end] -= need`.
4. If check succeeds, move `lo` up; else move `hi` down. Return the best `T`.

---

## Data Structures Used

* **Difference Array**: for both the base coverage and the temporary added coverage during feasibility check.
* **Prefix Sum**: to recover the actual power from difference arrays.
* **Binary Search**: on the answer.

---

## Operations & Behavior Summary

* Build base power: O(n)
* Feasibility check for a target `T`: O(n)
* Binary search range: about `log2(sum(stations)+k)` iterations
* Total: `O(n log U)`, where `U = sum(stations) + k`

---

## Complexity

* **Time Complexity:** `O(n log U)`
  `n` = number of cities.
  `U` = upper bound of answer (`sum(stations) + k`).
  Each check is O(n); binary search needs ~log U checks.
* **Space Complexity:** `O(n)`
  Base power array + temporary difference array used in checks.

---

## Multi-language Solutions

### C++

```cpp
#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    long long maxPower(vector<int>& stations, int r, int k) {
        int n = stations.size();

        // 1) Base power using difference array.
        vector<long long> diff(n + 1), base(n);
        for (int i = 0; i < n; ++i) {
            int L = max(0, i - r);
            int R = min(n, i + r + 1);
            diff[L] += stations[i];
            diff[R] -= stations[i];
        }
        long long run = 0;
        for (int i = 0; i < n; ++i) {
            run += diff[i];
            base[i] = run;
        }

        long long lo = 0, hi = accumulate(stations.begin(), stations.end(), 0LL) + k, ans = 0;

        auto feasible = [&](long long T) -> bool {
            vector<long long> add(n + 1, 0);
            long long extra = 0, used = 0;
            for (int i = 0; i < n; ++i) {
                extra += add[i];
                long long curr = base[i] + extra;
                if (curr < T) {
                    long long need = T - curr;
                    used += need;
                    if (used > k) return false;
                    extra += need;
                    int end = min(n, i + 2 * r + 1);
                    add[end] -= need; // effect ends after end-1
                }
            }
            return true;
        };

        while (lo <= hi) {
            long long mid = (lo + hi) >> 1;
            if (feasible(mid)) { ans = mid; lo = mid + 1; }
            else hi = mid - 1;
        }
        return ans;
    }
};
```

### Java

```java
import java.util.*;

class Solution {
    public long maxPower(int[] stations, int r, int k) {
        int n = stations.length;

        // 1) Base power with difference array.
        long[] diff = new long[n + 1];
        for (int i = 0; i < n; i++) {
            int L = Math.max(0, i - r);
            int R = Math.min(n, i + r + 1);
            diff[L] += stations[i];
            diff[R] -= stations[i];
        }
        long[] base = new long[n];
        long run = 0;
        for (int i = 0; i < n; i++) {
            run += diff[i];
            base[i] = run;
        }

        long lo = 0, hi = 0, ans = 0;
        for (int v : stations) hi += v;
        hi += k;

        // 2) Feasibility check.
        java.util.function.LongPredicate ok = T -> {
            long[] add = new long[n + 1];
            long extra = 0, used = 0;
            for (int i = 0; i < n; i++) {
                extra += add[i];
                long curr = base[i] + extra;
                if (curr < T) {
                    long need = T - curr;
                    used += need;
                    if (used > k) return false;
                    extra += need;
                    int end = Math.min(n, i + 2 * r + 1);
                    add[end] -= need;
                }
            }
            return true;
        };

        while (lo <= hi) {
            long mid = lo + ((hi - lo) >> 1);
            if (ok.test(mid)) { ans = mid; lo = mid + 1; }
            else hi = mid - 1;
        }
        return ans;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} stations
 * @param {number} r
 * @param {number} k
 * @return {number}
 */
var maxPower = function (stations, r, k) {
  const n = stations.length;

  // 1) Base power via difference array
  const diff = Array(n + 1).fill(0);
  for (let i = 0; i < n; i++) {
    const L = Math.max(0, i - r);
    const R = Math.min(n, i + r + 1);
    diff[L] += stations[i];
    diff[R] -= stations[i];
  }
  const base = Array(n).fill(0);
  let run = 0;
  for (let i = 0; i < n; i++) {
    run += diff[i];
    base[i] = run;
  }

  let lo = 0, hi = stations.reduce((s, v) => s + v, 0) + k, ans = 0;

  const ok = (T) => {
    const add = Array(n + 1).fill(0);
    let extra = 0, used = 0;
    for (let i = 0; i < n; i++) {
      extra += add[i];
      let curr = base[i] + extra;
      if (curr < T) {
        const need = T - curr;
        used += need;
        if (used > k) return false;
        extra += need;
        const end = Math.min(n, i + 2 * r + 1);
        add[end] -= need;
      }
    }
    return true;
  };

  while (lo <= hi) {
    const mid = Math.floor((lo + hi) / 2);
    if (ok(mid)) { ans = mid; lo = mid + 1; }
    else hi = mid - 1;
  }
  return ans;
};
```

### Python3

```python
from typing import List

class Solution:
    def maxPower(self, stations: List[int], r: int, k: int) -> int:
        n = len(stations)

        # 1) Base power with difference array
        diff = [0] * (n + 1)
        for i, v in enumerate(stations):
            L = max(0, i - r)
            R = min(n, i + r + 1)
            diff[L] += v
            diff[R] -= v

        base = [0] * n
        run = 0
        for i in range(n):
            run += diff[i]
            base[i] = run

        lo, hi = 0, sum(stations) + k
        ans = 0

        def ok(T: int) -> bool:
            add = [0] * (n + 1)
            extra = 0
            used = 0
            for i in range(n):
                extra += add[i]
                curr = base[i] + extra
                if curr < T:
                    need = T - curr
                    used += need
                    if used > k:
                        return False
                    extra += need
                    end = min(n, i + 2 * r + 1)
                    add[end] -= need
            return True

        while lo <= hi:
            mid = (lo + hi) // 2
            if ok(mid):
                ans = mid
                lo = mid + 1
            else:
                hi = mid - 1
        return ans
```

### Go

```go
package main

func maxPower(stations []int, r int, k int) int64 {
	n := len(stations)

	// 1) Base power via difference array.
	diff := make([]int64, n+1)
	for i, v := range stations {
		L := i - r
		if L < 0 {
			L = 0
		}
		R := i + r + 1
		if R > n {
			R = n
		}
		diff[L] += int64(v)
		diff[R] -= int64(v)
	}
	base := make([]int64, n)
	var run int64
	for i := 0; i < n; i++ {
		run += diff[i]
		base[i] = run
	}

	// 2) Binary search on T.
	var sum int64
	for _, v := range stations {
		sum += int64(v)
	}
	lo, hi := int64(0), sum+int64(k)
	var ans int64

	can := func(T int64) bool {
		add := make([]int64, n+1)
		var extra, used int64
		for i := 0; i < n; i++ {
			extra += add[i]
			curr := base[i] + extra
			if curr < T {
				need := T - curr
				used += need
				if used > int64(k) {
					return false
				}
				extra += need
				end := i + 2*r + 1
				if end > n {
					end = n
				}
				add[end] -= need
			}
		}
		return true
	}

	for lo <= hi {
		mid := (lo + hi) / 2
		if can(mid) {
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

1. **Build base coverage (`base[i]`)**

   * For each city `i`, add its `stations[i]` to coverage interval `[i-r, i+r]` using a difference array `diff`.
   * Example (C++):

     ```cpp
     diff[L] += stations[i];
     diff[R] -= stations[i];   // R is exclusive
     ```
   * Prefix-sum `diff` to get `base`.

2. **Binary search**

   * `lo = 0`, `hi = sum(stations) + k`.
   * While `lo ≤ hi`: test `mid`. If feasible → move `lo` up; else move `hi` down.

3. **Feasibility check for a target `T`**

   * Maintain `add[]` (difference array for newly added stations) and `extra` (running sum of `add`).
   * At index `i`, compute `curr = base[i] + extra`.
   * If `curr < T`:

     * Need `need = T - curr` new stations.
     * Spend them (`used += need`) and fail if `used > k`.
     * Start their effect now: `extra += need`.
     * End their effect at `end = min(n, i + 2*r + 1)` by `add[end] -= need`.
   * Continue to the end. If you never exceed `k`, `T` is feasible.

This greedy placement is optimal because placing as far right as allowed gives the **longest forward coverage**, helping future cities the most in a left-to-right sweep.

---

## Examples

**Example 1**
Input: `stations = [1,2,4,5,0], r = 1, k = 2`
One optimal plan is to add both new stations near the second city.
Answer: `5`

**Example 2**
Input: `stations = [4,4,4,4], r = 0, k = 3`
You can only strengthen individual cities.
Answer: `4` (already uniform; adding stations can increase, but min remains limited by available additions and binary search will find the correct max min).

> (Use LeetCode examples for exact values; these are illustrative.)

---

## How to use / Run locally

* **C++**

  ```bash
  g++ -std=gnu++17 -O2 solution.cpp -o sol && ./sol
  ```
* **Java**

  ```bash
  javac Solution.java && java Solution
  ```
* **JavaScript (Node)**

  ```bash
  node solution.js
  ```
* **Python**

  ```bash
  python3 solution.py
  ```
* **Go**

  ```bash
  go run solution.go
  ```

(Embed I/O wrapper if you want to test locally; on LeetCode these methods are called directly by the judge.)

---

## Notes & Optimizations

* Using a **difference array** avoids O(n·r) window recomputation.
* Greedy placement at `i + r` (implemented by ending effect at `i + 2r + 1`) is key to stay linear per check.
* Upper bound `sum(stations) + k` is safe and tight enough; using `min(base)` as `lo` is also fine, but `0` works.
* Use `long long / long / int64` where necessary to prevent overflow when accumulating.

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
