# 1488. Avoid Flood in The City

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

We are given an integer array `rains` where:

* `rains[i] > 0` means on day `i` it rains on lake `rains[i]` and that lake becomes **full**.
* `rains[i] == 0` means day `i` is dry and we can choose **one** lake to dry (empty) that day.

If it rains on a lake that is already full (and we didn't dry it after it was last filled), then there's a flood — this must be avoided. We need to return an array `ans` such that:

* `ans.length == rains.length`
* `ans[i] == -1` if `rains[i] > 0` (rainy days must be `-1`)
* `ans[i] == lake` on dry days to show which lake we dry (or any arbitrary positive number if not used)
* If impossible to avoid any flood, return `[]`.

---

## Constraints

* `1 <= rains.length <= 10^5`
* `0 <= rains[i] <= 10^9`

---

## Intuition

I thought about what causes a flood: it happens only if a lake that is already full receives rain again before I dry it. So I need to remember **when each lake was last filled** and I need to use the **available dry days** to dry those lakes before they get rained on again.

If a lake was filled at day `a` and will be rained again at day `b`, I must assign some dry day `d` with `a < d < b` to dry it. If multiple dry days are possible, I should use the **earliest** such dry day (greedy). This choice maximizes the chance later lakes can use later dry days.

---

## Approach

1. Create `ans` array initialized to `1` (we'll set `-1` on rainy days). `1` is arbitrary — any positive number is acceptable for unused dry days.
2. Maintain a map `last` mapping `lake -> last_day_index` (the last day lake was filled).
3. Maintain an ordered set (or sorted list) `dryDays` of indices of days where `rains[i] == 0`.
4. Iterate day-by-day:

   * If today rains on `lake`:

     * If `lake` is in `last` (i.e., it was previously filled), find the **smallest** dry day index in `dryDays` that is `> last[lake]`. If none exists, return `[]` (impossible).
     * Use that dry day to dry `lake` (set `ans[dryDay] = lake`) and remove dryDay from `dryDays`.
     * Set `ans[i] = -1` and update `last[lake] = i`.
   * If today is dry: add index `i` to `dryDays` (and leave `ans[i] = 1` until possibly assigned).
5. Return `ans`.

This is a greedy + bookkeeping solution: always use the earliest dry day after a lake's last fill.

---

## Data Structures Used

* `map` / `unordered_map` (C++), `HashMap` (Java), `Map` (JS/Python dict) — to store last fill day per lake.
* Ordered set `set` (C++), `TreeSet` (Java) — to store dry day indices and find the smallest `dryDay > prev`.
* Sorted list + binary search (Python/JS/Go versions) — `bisect` / `sort.Search` for index selection.

---

## Operations & Behavior Summary

* Rainy day (`rains[i] > 0`):

  * `ans[i] = -1`.
  * If lake was filled earlier, we must pick a dry day in the future (between last fill and now) to dry it.
* Dry day (`rains[i] == 0`):

  * Candidate day for drying; `ans[i]` keeps `1` unless assigned to a specific lake.
* Return `[]` if we ever cannot find a dry day to prevent flood.

---

## Complexity

* **Time Complexity:**

  * Using ordered set (`set`/`TreeSet`): `O(n log n)` where `n = rains.length`. Each dry-day insertion / query / removal costs `O(log n)`.
  * Using sorted list + binary search and slice deletion (Python/JS/Go code below): binary search is `O(log n)` but deletion costs `O(n)` leading to worst-case `O(n^2)`. In practice many inputs pass, but for strict guarantees use a balanced BST or sorted container supporting `O(log n)` deletion.
* **Space Complexity:** `O(n)` for `ans`, `last` map, and `dryDays` structure.

---

## Multi-language Solutions

### C++

```c++
#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    vector<int> avoidFlood(vector<int>& rains) {
        int n = rains.size();
        vector<int> ans(n, 1);
        unordered_map<int,int> last;   // lake -> last day index it rained
        set<int> dryDays;             // indices of dry days (sorted)

        for (int i = 0; i < n; ++i) {
            if (rains[i] > 0) {
                int lake = rains[i];
                ans[i] = -1;  // rainy day must be -1
                if (last.count(lake)) {
                    int prevDay = last[lake];
                    auto it = dryDays.lower_bound(prevDay + 1);
                    if (it == dryDays.end()) {
                        return {}; // impossible to avoid flood
                    }
                    ans[*it] = lake;   // dry that lake on this dry day
                    dryDays.erase(it); // remove used dry day
                }
                last[lake] = i; // update last rainy day for lake
            } else {
                dryDays.insert(i); // available dry day
                // ans[i] remains 1 until (maybe) assigned to dry a lake
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
    public int[] avoidFlood(int[] rains) {
        int n = rains.length;
        int[] ans = new int[n];
        Arrays.fill(ans, 1);
        Map<Integer,Integer> last = new HashMap<>(); // lake -> last day it rained
        TreeSet<Integer> dry = new TreeSet<>();      // available dry days (sorted)

        for (int i = 0; i < n; ++i) {
            if (rains[i] > 0) {
                int lake = rains[i];
                ans[i] = -1;
                if (last.containsKey(lake)) {
                    int prev = last.get(lake);
                    Integer dryDay = dry.higher(prev); // first dry day > prev
                    if (dryDay == null) return new int[0];
                    ans[dryDay] = lake; // dry that lake on dryDay
                    dry.remove(dryDay);
                }
                last.put(lake, i);
            } else {
                dry.add(i); // collect dry day
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
 * JavaScript (binary-search on a sorted dry-day list)
 * Note: splice removal is O(n). Worst-case O(n^2).
 * @param {number[]} rains
 * @return {number[]}
 */
var avoidFlood = function(rains) {
    const n = rains.length;
    const ans = new Array(n).fill(1);
    const last = new Map(); // lake -> last day index
    const dry = [];         // sorted list of dry day indices (kept sorted by appending)

    // upperBound: first index with arr[idx] > target
    const upperBound = (arr, target) => {
        let l = 0, r = arr.length;
        while (l < r) {
            let m = (l + r) >> 1;
            if (arr[m] <= target) l = m + 1;
            else r = m;
        }
        return l;
    };

    for (let i = 0; i < n; ++i) {
        if (rains[i] > 0) {
            const lake = rains[i];
            ans[i] = -1;
            if (last.has(lake)) {
                const prev = last.get(lake);
                const idx = upperBound(dry, prev); // first dry day > prev
                if (idx === dry.length) return []; // impossible
                const dryDay = dry[idx];
                ans[dryDay] = lake;        // use this dry day to dry the lake
                dry.splice(idx, 1);        // remove used dry day
            }
            last.set(lake, i);
        } else {
            dry.push(i); // append keeps `dry` sorted because i increases
        }
    }
    return ans;
};
```

---

### Python3

```python
from typing import List
import bisect

class Solution:
    def avoidFlood(self, rains: List[int]) -> List[int]:
        n = len(rains)
        ans = [1] * n
        last = {}  # lake -> last day index
        dry = []   # sorted list of dry day indices (increasing because we append)

        for i, lake in enumerate(rains):
            if lake > 0:
                ans[i] = -1
                if lake in last:
                    prev = last[lake]
                    # find the first dry day strictly greater than prev
                    idx = bisect.bisect_right(dry, prev)
                    if idx == len(dry):
                        return []
                    dry_day = dry[idx]
                    ans[dry_day] = lake
                    dry.pop(idx)  # remove used dry day
                last[lake] = i
            else:
                dry.append(i)  # append keeps it sorted (i increases each loop)
        return ans
```

---

### Go

```go
package main

import "sort"

// Go (sorted slice + sort.Search)
// Deleting from slice costs O(n) so worst-case O(n^2).
func avoidFlood(rains []int) []int {
    n := len(rains)
    ans := make([]int, n)
    for i := range ans { ans[i] = 1 }
    last := map[int]int{} // lake -> last day it rained
    var dry []int         // sorted list of dry day indices

    for i, lake := range rains {
        if lake > 0 {
            ans[i] = -1
            if prev, ok := last[lake]; ok {
                // find first dry day > prev
                j := sort.Search(len(dry), func(k int) bool { return dry[k] > prev })
                if j == len(dry) {
                    return []int{} // impossible
                }
                ans[dry[j]] = lake
                // remove dry[j]
                dry = append(dry[:j], dry[j+1:]...)
            }
            last[lake] = i
        } else {
            dry = append(dry, i)
        }
    }
    return ans
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

I'll explain the key logic (same idea in all languages) step-by-step using simple language.

### Common idea (core loop)

1. I create an `ans` array filled with `1`. This covers dry days by default. For rainy days I set `ans[i] = -1`.
2. I track `last[lake] = dayIndex` when the lake was last rained on.
3. I collect indices of dry days (`dryDays`) as they appear. `dryDays` must be searchable to find the smallest `dryDay` strictly greater than a given previous day.
4. When it's raining on `lake`:

   * If `lake` has been seen before (exists in `last`), I must dry it after `last[lake]` and before the current day.
   * So I search `dryDays` for the first index `d` such that `d > last[lake]`. If none, impossible -> return `[]`.
   * Otherwise set `ans[d] = lake` (I dry `lake` on that day) and remove `d` from `dryDays`.
   * Then update `last[lake] = currentDay`, and mark `ans[currentDay] = -1`.
5. If it's a dry day, just add the index to `dryDays` for future use.

### Why use the earliest dry day after last fill?

Because if I have multiple dry days available, using the earliest possible one for the lake that requires drying is greedy and safe: it leaves more later dry days for future conflicts.

---

### C++ specifics

* I use `unordered_map<int,int> last` to map lake -> last filled day.
* I use `set<int> dryDays` which is a balanced BST, so `lower_bound(prev + 1)` gives the smallest dry-day index greater than `prev` in `O(log n)`.
* Removing that dry day from `dryDays` is `O(log n)`.

Key snippet:

```c++
auto it = dryDays.lower_bound(prevDay + 1);
if (it == dryDays.end()) return {};
ans[*it] = lake;
dryDays.erase(it);
```

---

### Java specifics

* `HashMap<Integer,Integer> last` and `TreeSet<Integer> dry`.
* `dry.higher(prev)` finds first element `> prev` in `O(log n)`.

Key snippet:

```java
Integer dryDay = dry.higher(prev);
if (dryDay == null) return new int[0];
ans[dryDay] = lake;
dry.remove(dryDay);
```

---

### Python specifics

* I use `last` dictionary and `dry` list (kept sorted by appending indices in increasing order).
* I use `bisect.bisect_right(dry, prev)` to find first dry day `> prev`.
* `dry.pop(idx)` removes the used dry day (this removal is O(n) in worst-case).

Key snippet:

```python
idx = bisect.bisect_right(dry, prev)
if idx == len(dry): return []
dry_day = dry[idx]
ans[dry_day] = lake
dry.pop(idx)
```

---

### JavaScript specifics

* Same approach as Python: `last` Map and `dry` array.
* Implemented `upperBound` (binary search) to find first `dry` element > `prev`.
* Use `dry.splice(idx,1)` to remove used dry day (O(n) cost).

Key snippet:

```javascript
const idx = upperBound(dry, prev);
if (idx === dry.length) return [];
const dryDay = dry[idx];
ans[dryDay] = lake;
dry.splice(idx, 1);
```

---

### Go specifics

* `last` map, `dry` slice.
* Use `sort.Search` to find first element > `prev`. Remove slice element with `append(dry[:j], dry[j+1:]...)` (O(n) cost).

Key snippet:

```go
j := sort.Search(len(dry), func(k int) bool { return dry[k] > prev })
if j == len(dry) { return []int{} }
ans[dry[j]] = lake
dry = append(dry[:j], dry[j+1:]...)
```

---

## Examples

### Example 1

Input: `rains = [1,2,3,4]`
Output: `[-1,-1,-1,-1]`
Explanation: Every day it rains a different empty lake. No dry days required. No flood.

### Example 2

Input: `rains = [1,2,0,0,2,1]`
Possible Output: `[-1,-1,2,1,-1,-1]`
Explanation:

* Day0: lake1 filled -> ans[0] = -1.
* Day1: lake2 filled -> ans[1] = -1.
* Day2: dry -> we plan to dry lake2 on day2 (later assigned `2`).
* Day3: dry -> we plan to dry lake1 on day3 (assigned `1`).
* Day4: rains on lake2 (safe because we dried it on day2).
* Day5: rains on lake1 (safe because we dried it on day3).

### Example 3

Input: `rains = [1,2,0,1,2]`
Output: `[]`
Explanation: After day1 lakes {1,2} are full. Day2 is dry, we can dry only one lake. But whichever we choose, the other will be rained on later and cause a flood. So impossible.

---

## How to use / Run locally

### General

* These functions are ready to be used on platforms like LeetCode (paste the function/class).
* For local testing, wrap solutions in a small driver program to call the function with test arrays and print results.

### C++

1. Save the class in a file (e.g., `solution.cpp`) and add a `main()` driver to test.
2. Compile & run:

   ```bash
   g++ -std=c++17 -O2 solution.cpp -o solution
   ./solution
   ```

### Java

1. Put the `Solution` class into `Solution.java` and add a `main` to create an instance and test.
2. Compile & run:

   ```bash
   javac Solution.java
   java Solution
   ```

### Python3

1. Paste the `Solution` class in a file `solution.py`, add a test driver block:

   ```python
   if __name__ == "__main__":
       s = Solution()
       print(s.avoidFlood([1,2,0,0,2,1]))
   ```
2. Run:

   ```bash
   python3 solution.py
   ```

### JavaScript (Node)

1. Save the function in `solution.js` and add a small test driver:

   ```javascript
   console.log(avoidFlood([1,2,0,0,2,1]));
   ```
2. Run:

   ```bash
   node solution.js
   ```

### Go

1. Put the function into `main.go`, add a `main()` that calls `avoidFlood(...)` and prints result.
2. Run:

   ```bash
   go run main.go
   ```

---

## Notes & Optimizations

* C++ and Java solutions above guarantee `O(n log n)` using `set`/`TreeSet`.
* Python / JavaScript / Go versions use a sorted list + binary search and remove elements with `pop/splice`, so they can degrade to `O(n^2)` worst-case due to deletions. For strict `O(n log n)` in Python use `sortedcontainers.SortedList` (`pip install sortedcontainers`) or implement a balanced tree.
* The greedy choice (earliest possible dry day for a conflict) is provably safe: it never makes future decisions harder.
* On dry days that remain unused, any positive number is acceptable (we use `1`).

---

## Author

[Md. Aarzoo Islam](https://bento.me/withaarzoo)
