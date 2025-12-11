# Count Covered Buildings

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

We are given an integer `n` representing an `n x n` city and a list `buildings` of unique coordinates `[x, y]` where each coordinate denotes a building. A building is **covered** if there is **at least one building** in **all four directions** relative to it: left, right, above, and below. Return the number of covered buildings.

In other words, for building `(x, y)` to be covered:

* There exists at least one building with the same `x` and `y' < y` (left),
* There exists at least one building with the same `x` and `y' > y` (right),
* There exists at least one building with the same `y` and `x' < x` (above),
* There exists at least one building with the same `y` and `x' > x` (below).

---

## Constraints

* `2 <= n <= 10^5`
* `1 <= buildings.length <= 10^5`
* `buildings[i] = [x, y]`
* `1 <= x, y <= n`
* All coordinates in `buildings` are unique.

---

## Intuition

I thought about what "covered" means in simpler terms: a building `(x, y)` is covered iff it is **not** the leftmost or rightmost in its row `x` and it is **not** the topmost or bottommost in its column `y`. So if in the sorted set of `y`'s in row `x`, `y` is neither first nor last, it has both left and right neighbors. Similarly for the column.

Thus, I can:

1. Group buildings by their row and by their column.
2. Sort each group's coordinates.
3. For each building, check whether its index in the row group and column group is interior (not first, not last).
4. Count those that are interior in both row and column.

---

## Approach

1. Create two maps:

   * `row[x]` → list of `y` values for buildings on row `x`.
   * `col[y]` → list of `x` values for buildings on column `y`.
2. Populate these maps by iterating the `buildings` list once.
3. Sort each list in `row` and `col`.
4. For each building `(x, y)`:

   * Find index of `y` inside `row[x]` and check if it's not first or last → has left & right.
   * Find index of `x` inside `col[y]` and check if it's not first or last → has above & below.
   * If both checks are true, increment result.
5. Return the count.

This method uses sorting per group, and total sorting work across all groups is `O(m log m)` where `m = buildings.length`.

---

## Data Structures Used

* Hash map / dictionary (`unordered_map`, `Map`, `dict`) to group coordinates by row and by column.
* Arrays / lists to store grouped coordinates.
* Sorting (and binary search / lower_bound) to find positions quickly.

---

## Operations & Behavior Summary

* Build grouping maps: O(m)
* Sort each group's list: total O(m log m) in worst case
* For each building, binary search index in row/column lists: O(log k) per building (amortized) but overall O(m log m) dominated by sorting
* Return count of buildings interior in both row and column

---

## Complexity

* **Time Complexity:** `O(m log m)` where `m = number of buildings`.

  * Building maps: `O(m)`
  * Sorting lists across all rows and columns combined: `O(m log m)` worst-case
  * Checking each building with binary search: `O(m log k)` but dominated by the sorting step
* **Space Complexity:** `O(m)` to store the row and column groupings and any temporary lists.

---

## Multi-language Solutions

### C++

```c++
#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    int countCoveredBuildings(int n, vector<vector<int>>& buildings) {
        int m = buildings.size();
        unordered_map<int, vector<int>> row; // x -> list of y
        unordered_map<int, vector<int>> col; // y -> list of x
        row.reserve(m*2);
        col.reserve(m*2);

        // Build maps
        for (auto &b : buildings) {
            int x = b[0], y = b[1];
            row[x].push_back(y);
            col[y].push_back(x);
        }

        // Sort each row and column
        for (auto &kv : row) {
            auto &v = kv.second;
            sort(v.begin(), v.end());
        }
        for (auto &kv : col) {
            auto &v = kv.second;
            sort(v.begin(), v.end());
        }

        int ans = 0;
        for (auto &b : buildings) {
            int x = b[0], y = b[1];
            auto &r = row[x];
            auto &c = col[y];

            // find position of y in r
            auto itR = lower_bound(r.begin(), r.end(), y);
            bool insideRow = (itR != r.begin() && (itR+1) != r.end());
            // find position of x in c
            auto itC = lower_bound(c.begin(), c.end(), x);
            bool insideCol = (itC != c.begin() && (itC+1) != c.end());

            if (insideRow && insideCol) ++ans;
        }
        return ans;
    }
};
```

### Java

```java
import java.util.*;

class Solution {
    public int countCoveredBuildings(int n, int[][] buildings) {
        int m = buildings.length;
        Map<Integer, List<Integer>> row = new HashMap<>(); // x -> list of y
        Map<Integer, List<Integer>> col = new HashMap<>(); // y -> list of x

        // Build maps
        for (int[] b : buildings) {
            int x = b[0], y = b[1];
            row.computeIfAbsent(x, k -> new ArrayList<>()).add(y);
            col.computeIfAbsent(y, k -> new ArrayList<>()).add(x);
        }

        // Sort lists
        for (List<Integer> ys : row.values()) Collections.sort(ys);
        for (List<Integer> xs : col.values()) Collections.sort(xs);

        int ans = 0;
        for (int[] b : buildings) {
            int x = b[0], y = b[1];
            List<Integer> ys = row.get(x);
            List<Integer> xs = col.get(y);

            int posY = Collections.binarySearch(ys, y);
            int posX = Collections.binarySearch(xs, x);
            boolean insideRow = (posY > 0 && posY < ys.size() - 1);
            boolean insideCol = (posX > 0 && posX < xs.size() - 1);

            if (insideRow && insideCol) ans++;
        }
        return ans;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number} n
 * @param {number[][]} buildings
 * @return {number}
 */
var countCoveredBuildings = function(n, buildings) {
    const row = new Map();
    const col = new Map();

    for (const b of buildings) {
        const x = b[0], y = b[1];
        if (!row.has(x)) row.set(x, []);
        row.get(x).push(y);
        if (!col.has(y)) col.set(y, []);
        col.get(y).push(x);
    }

    for (const [k, arr] of row) arr.sort((a,b)=>a-b);
    for (const [k, arr] of col) arr.sort((a,b)=>a-b);

    function lowerBound(arr, val) {
        let l = 0, r = arr.length;
        while (l < r) {
            const mid = (l + r) >> 1;
            if (arr[mid] < val) l = mid + 1;
            else r = mid;
        }
        return l;
    }

    let ans = 0;
    for (const b of buildings) {
        const x = b[0], y = b[1];
        const ys = row.get(x);
        const xs = col.get(y);
        const posY = lowerBound(ys, y);
        const posX = lowerBound(xs, x);
        const insideRow = (posY > 0 && posY < ys.length - 1);
        const insideCol = (posX > 0 && posX < xs.length - 1);
        if (insideRow && insideCol) ans++;
    }
    return ans;
};
```

### Python3

```python
from typing import List
import bisect

class Solution:
    def countCoveredBuildings(self, n: int, buildings: List[List[int]]) -> int:
        # maps: row x -> list of y, col y -> list of x
        row = {}
        col = {}
        for x, y in buildings:
            row.setdefault(x, []).append(y)
            col.setdefault(y, []).append(x)

        # sort each list
        for ys in row.values():
            ys.sort()
        for xs in col.values():
            xs.sort()

        ans = 0
        for x, y in buildings:
            ys = row[x]
            xs = col[y]
            posY = bisect.bisect_left(ys, y)
            posX = bisect.bisect_left(xs, x)
            insideRow = (posY > 0 and posY < len(ys) - 1)
            insideCol = (posX > 0 and posX < len(xs) - 1)
            if insideRow and insideCol:
                ans += 1
        return ans
```

### Go

```go
package main

import (
 "sort"
)

func countCoveredBuildings(n int, buildings [][]int) int {
 row := make(map[int][]int)
 col := make(map[int][]int)

 for _, b := range buildings {
  x, y := b[0], b[1]
  row[x] = append(row[x], y)
  col[y] = append(col[y], x)
 }

 for k := range row {
  sort.Ints(row[k])
 }
 for k := range col {
  sort.Ints(col[k])
 }

 ans := 0
 for _, b := range buildings {
  x, y := b[0], b[1]
  ys := row[x]
  xs := col[y]
  posY := sort.SearchInts(ys, y)
  posX := sort.SearchInts(xs, x)
  insideRow := posY > 0 && posY < len(ys)-1
  insideCol := posX > 0 && posX < len(xs)-1
  if insideRow && insideCol {
   ans++
  }
 }
 return ans
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

Below are the important logical steps explained once and how they map to the code in each language. The mapping is close — the same algorithm and data structures are used across languages.

### 1) Build maps grouping by row and column

* Idea: quickly find all buildings that share a row or column.
* Code mapping:

  * C++: `row[x].push_back(y); col[y].push_back(x);`
  * Java: `row.computeIfAbsent(x, ...).add(y);`
  * JS: `row.get(x).push(y);` (after `if (!row.has(x)) row.set(x, []);`)
  * Python: `row.setdefault(x, []).append(y)`
  * Go: `row[x] = append(row[x], y)`

### 2) Sort each group's coordinate list

* Idea: After sorting, interior positions (not first or last) indicate existence of both directions.
* Code mapping:

  * C++: `sort(v.begin(), v.end());`
  * Java: `Collections.sort(ys);`
  * JS: `arr.sort((a,b)=>a-b);`
  * Python: `ys.sort()`
  * Go: `sort.Ints(row[k])`

### 3) Check each building's position in its sorted row and column

* Idea: Find index of `y` in `row[x]` and index of `x` in `col[y]`. If both indices are interior, building is covered.
* Position lookup:

  * C++: `lower_bound(r.begin(), r.end(), y)` returns iterator; check `it != begin` and `it+1 != end`.
  * Java: `Collections.binarySearch(ys, y)` returns index.
  * JS: `lowerBound` function implemented to get index.
  * Python: `bisect_left(ys, y)`
  * Go: `sort.SearchInts(ys, y)`
* Condition: `pos > 0 && pos < len-1` → interior.

### 4) Count and return

* Increment counter for each building that is interior in both row and column.
* Return final count.

---

## Examples

Example 1:

```
Input:
n = 3
buildings = [[1,2],[2,2],[3,2],[2,1],[2,3]]

Output: 1
Explanation: Only (2,2) has at least one building in all four directions.
```

Example 2:

```
Input:
n = 3
buildings = [[1,1],[1,2],[2,1],[2,2]]

Output: 0
Explanation: No building has at least one building in all four directions.
```

Example 3:

```
Input:
n = 5
buildings = [[1,3],[3,2],[3,3],[3,5],[5,3]]

Output: 1
Explanation: (3,3) is the only covered building.
```

---

## How to use / Run locally

1. **Python3**

   * Save `solution.py` with the `Solution` class.
   * Example runner:

     ```python
     from typing import List
     # import Solution class
     sol = Solution()
     print(sol.countCoveredBuildings(3, [[1,2],[2,2],[3,2],[2,1],[2,3]]))  # -> 1
     ```

2. **C++**

   * Create a file `solution.cpp`. Include test harness that constructs `vector<vector<int>> buildings`, calls `Solution().countCoveredBuildings(n, buildings)` and prints result.
   * Compile & run:

     ```bash
     g++ -std=c++17 solution.cpp -O2 -o solution
     ./solution
     ```

3. **Java**

   * Put `Solution` class in `Solution.java` with a `main` method to test.
   * Compile & run:

     ```bash
     javac Solution.java
     java Solution
     ```

4. **JavaScript (Node.js)**

   * Save function in `solution.js` and add a test harness calling it.
   * Run:

     ```bash
     node solution.js
     ```

5. **Go**

   * Create a `main.go` that calls `countCoveredBuildings`.
   * Build & run:

     ```bash
     go run main.go
     ```

---

## Notes & Optimizations

* Using maps grouping by rows and columns minimizes repeated scans. Sorting is the dominant cost.
* If memory is a concern and the number of distinct rows/columns is small, the same approach still works but be mindful of map sizes.
* The algorithm avoids building an `n x n` grid which would be impossible for large `n`.
* For very large inputs, reserve/initial capacity for maps and vectors (where language supports it) to reduce rehash/resize overhead.
* This implementation assumes coordinates are 1-indexed as per the problem statement. It works regardless of indexing as long as groupings are consistent.

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
