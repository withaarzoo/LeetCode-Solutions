# Set Intersection Size At Least Two (LeetCode 757)

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

I am given a list of intervals `intervals`, where `intervals[i] = [start_i, end_i]` and each interval represents all integers from `start_i` to `end_i` inclusive.
A *containing set* `S` is a set of integers such that each interval contains *at least two integers from `S`*. I need to return the **minimum possible size** of such a containing set `S`.

Example: for `[[1,3],[3,7],[8,9]]`, one minimal S is `[2,3,4,8,9]` of size `5`.

---

## Constraints

* `1 <= intervals.length <= 3000`
* `intervals[i].length == 2`
* `0 <= start_i <= end_i <= 10^8`

---

## Intuition

I thought about how to cover each interval with at least two chosen integers while keeping the total number of chosen integers minimal.
I realized picking numbers as *far right as possible* inside each interval is best. Why? Because numbers near the end of an interval are more likely to be inside later intervals (which are sorted by their ends). So a greedy strategy that processes intervals by increasing `end` and always tries to reuse previously chosen numbers works well.

I keep track of the two most recently chosen integers, `a < b`. For each interval `[l, r]`:

* If both `a` and `b` are inside `[l, r]`: do nothing.
* If only `b` is inside: add `r` (one new integer).
* If none are inside: add `r-1` and `r` (two new integers).

Sorting ties by `start` in descending order (when equal ends) avoids redundant picks for nested intervals.

---

## Approach

1. Sort intervals by `end` ascending. If two intervals share an `end`, sort by `start` descending.
2. Maintain two variables `a` and `b` representing the last two selected integers (initially very small).
3. Iterate intervals:

   * If `l > b`: neither `a` nor `b` are inside → choose `r-1` and `r`. Update `a = r-1`, `b = r`. Add 2 to answer.
   * Else if `l > a`: only `b` is inside → choose `r`. Update `a = b`, `b = r`. Add 1 to answer.
   * Else: both `a` and `b` already inside → do nothing.
4. Return the accumulated answer.

This is a greedy algorithm that uses minimal state and is optimal for the problem.

---

## Data Structures Used

* Array / vector (for intervals)
* Primitive variables `a`, `b` (two integers)
* Sorting (in-place or language-provided sort)

No heavy extra data structures are required — just sorting and a couple of variables.

---

## Operations & Behavior Summary

* Sort intervals by `(end ascending, start descending)` — O(n log n).
* Single pass over sorted intervals — O(n).
* For each interval do a constant-time check and possible updates to `a`, `b`, and the count.

Behavior:

* Always try to place the new points at the right end (`r` or `r-1`, `r`) to maximize reuse.
* Ties-handling (start descending) ensures intervals with same `end` but larger starts are processed first to avoid adding needless numbers.

---

## Complexity

* **Time Complexity:** `O(n log n)` where `n` is the number of intervals — sorting dominates (`O(n log n)`). The traversal is `O(n)`.
* **Space Complexity:** `O(1)` extra space (ignoring input sort space). I only store two integers and a counter. In-place sorting may use O(log n) stack depending on the language/runtime.

---

## Multi-language Solutions

### C++

```c++
#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    int intersectionSizeTwo(vector<vector<int>>& intervals) {
        // Sort by end increasing. If ends equal, sort start decreasing.
        sort(intervals.begin(), intervals.end(), [](const vector<int>& A, const vector<int>& B){
            if (A[1] != B[1]) return A[1] < B[1];
            return A[0] > B[0];
        });

        const int NEG = -1000000000; // sufficiently small sentinel
        int a = NEG, b = NEG; // keep last two chosen integers (a < b)
        int ans = 0;

        for (auto &iv : intervals) {
            int l = iv[0], r = iv[1];
            if (l > b) {
                // none of a,b in [l,r]: pick r-1 and r
                ans += 2;
                a = r - 1;
                b = r;
            } else if (l > a) {
                // only b is inside: pick r
                ans += 1;
                a = b;
                b = r;
            } else {
                // both a and b already inside: do nothing
            }
        }
        return ans;
    }
};
```

### Java

```java
import java.util.*;

class Solution {
    public int intersectionSizeTwo(int[][] intervals) {
        Arrays.sort(intervals, (x, y) -> {
            if (x[1] != y[1]) return Integer.compare(x[1], y[1]);
            return Integer.compare(y[0], x[0]); // start descending if end tie
        });

        int NEG = Integer.MIN_VALUE / 2;
        int a = NEG, b = NEG; // last two chosen
        int ans = 0;

        for (int[] iv : intervals) {
            int l = iv[0], r = iv[1];
            if (l > b) {
                ans += 2;
                a = r - 1;
                b = r;
            } else if (l > a) {
                ans += 1;
                a = b;
                b = r;
            } else {
                // already have two numbers inside
            }
        }
        return ans;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[][]} intervals
 * @return {number}
 */
var intersectionSizeTwo = function(intervals) {
    intervals.sort((A, B) => {
        if (A[1] !== B[1]) return A[1] - B[1];
        return B[0] - A[0];
    });

    let NEG = -1e18;
    let a = NEG, b = NEG;
    let ans = 0;

    for (const [l, r] of intervals) {
        if (l > b) {
            ans += 2;
            a = r - 1;
            b = r;
        } else if (l > a) {
            ans += 1;
            a = b;
            b = r;
        } else {
            // both already inside
        }
    }
    return ans;
};
```

### Python3

```python
from typing import List

class Solution:
    def intersectionSizeTwo(self, intervals: List[List[int]]) -> int:
        # sort by end ascending, start descending when tie
        intervals.sort(key=lambda x: (x[1], -x[0]))

        NEG = -10**18
        a, b = NEG, NEG
        ans = 0

        for l, r in intervals:
            if l > b:
                # none inside
                ans += 2
                a = r - 1
                b = r
            elif l > a:
                # only b inside
                ans += 1
                a = b
                b = r
            else:
                # both inside -> do nothing
                pass
        return ans
```

### Go

```go
package main

import (
 "sort"
)

func intersectionSizeTwo(intervals [][]int) int {
 sort.Slice(intervals, func(i, j int) bool {
  if intervals[i][1] != intervals[j][1] {
   return intervals[i][1] < intervals[j][1]
  }
  return intervals[i][0] > intervals[j][0]
 })

 const NEG = -1 << 60
 a, b := NEG, NEG
 ans := 0

 for _, iv := range intervals {
  l, r := iv[0], iv[1]
  if l > b {
   ans += 2
   a = r - 1
   b = r
  } else if l > a {
   ans += 1
   a = b
   b = r
  } else {
   // both already inside
  }
 }
 return ans
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

I will explain the C++ solution line-by-line and map the logic to other languages (they are the same algorithm).

1. **Sorting**

   ```c++
   sort(intervals.begin(), intervals.end(), [](const vector<int>& A, const vector<int>& B){
       if (A[1] != B[1]) return A[1] < B[1];
       return A[0] > B[0];
   });
   ```

   * I sort intervals by `end` in increasing order.
   * If two intervals share the same `end`, I sort by `start` in decreasing order. This avoids unnecessary picks for nested intervals with the same `end`.

2. **Initialize sentinel and answer**

   ```c++
   const int NEG = -1000000000;
   int a = NEG, b = NEG;
   int ans = 0;
   ```

   * `a` and `b` are the last two integers I chose.
   * Initialize them to a very negative value so they won't lie inside any real interval.
   * `ans` counts how many integers I've chosen.

3. **Iterate intervals**

   ```c++
   for (auto &iv : intervals) {
       int l = iv[0], r = iv[1];
   ```

   * For each interval, fetch `l` and `r`.

4. **Case: neither a nor b in interval**

   ```c++
   if (l > b) {
       ans += 2;
       a = r - 1;
       b = r;
   }
   ```

   * `l > b` means both `a` and `b` are left of the interval (not included).
   * I choose two numbers at the right end of the interval: `r-1` and `r`.
   * This maximizes reuse for subsequent intervals because future intervals have greater or equal `end`.

5. **Case: only b in interval**

   ```c++
   else if (l > a) {
       ans += 1;
       a = b;
       b = r;
   }
   ```

   * `l > a` but `l <= b` — indicates `b` is inside `[l, r]`, but `a` is not.
   * I add the single number `r`. Now the most recent picks become the previous `b` and the new `r`.

6. **Case: both in interval**

   ```c++
   else {
       // do nothing
   }
   ```

   * If we reach here, `l <= a`, so both `a` and `b` lie within `[l, r]`. No additions needed.

7. **Return**

   ```c++
   return ans;
   ```

   * The total chosen integers is returned.

Mapping to other languages:

* The sort comparator, sentinel values, and variable updates are exactly the same in logic. Only syntax differs.

---

## Examples

1. Input: `[[1,3],[3,7],[8,9]]`
   Process:

   * Sort → `[[1,3],[3,7],[8,9]]`
   * Interval [1,3] → no picks yet → pick {2,3} → ans = 2 (a=2,b=3)
   * Interval [3,7] → only b=3 inside → pick {7} → ans = 3 (a=3,b=7)
   * Interval [8,9] → none inside → pick {8,9} → ans = 5
     Output: `5`

2. Input: `[[1,3],[1,4],[2,5],[3,5]]`
   Output: `3`
   Explanation: I can choose `{2,3,4}` or similar minimal containing set of size 3.

3. Input: `[[1,2],[2,3],[2,4],[4,5]]`
   Output: `5`
   Explanation via greedy picks yields the minimal count.

---

## How to use / Run locally

**C++ (g++):**

1. Save the C++ code in `solution.cpp`.
2. Compile: `g++ -std=c++17 -O2 solution.cpp -o solution`
3. Run: `./solution` (Make a main wrapper to test custom cases or use the LeetCode environment.)

**Java:**

1. Save the Java class as `Solution.java`.
2. Compile: `javac Solution.java`
3. Run: `java Solution` (wrap with a `main` for local tests).

**JavaScript (Node.js):**

1. Save function in `solution.js` and add sample invocation at end.
2. Run: `node solution.js`

**Python3:**

1. Save class in `solution.py` and write a test harness below or use LeetCode to run.
2. Run: `python3 solution.py`

**Go:**

1. Save code in `main.go`, include a `main()` that calls `intersectionSizeTwo`.
2. Build: `go build -o solution main.go`
3. Run: `./solution`

> Note: For local testing, add a small `main()` or `if __name__ == "__main__"` test harness that constructs intervals and prints the function output. On LeetCode, the platform provides the harness.

---

## Notes & Optimizations

* The greedy strategy is optimal because placing points as far right as possible maximizes re-use for future intervals when they're processed in increasing `end` order.
* Sorting by `end` asc, `start` desc is crucial to avoid unnecessary picks for intervals with same `end`.
* The algorithm uses constant extra memory and runs fast for `n` up to 3000.
* Edge case: Very small intervals of length 1 (`[x, x]`) require 2 picks but are impossible — but per constraints `start <= end`, such intervals are allowed; when `l == r` we might pick `r-1` and `r` — `r-1` could be outside domain but problem treats integers freely, so picks are allowed as long as they satisfy intervals requirements. (Standard accepted implementations assume integer line is unbounded in negative direction due to sentinel initialization.)
* If you prefer stricter handling for very small `[r, r]` intervals, you can still pick `r-1` and `r`; it's allowed because we only require chosen numbers to be integers (no explicit non-negative restriction).

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
