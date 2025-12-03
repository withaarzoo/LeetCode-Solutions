* Problem Title
  **LeetCode 3625 – Count Number of Trapezoids II**

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

I am given **n distinct points** on the 2D plane.

I have to **count how many unique quadrilaterals (4 distinct points)** form a **trapezoid**.

* A trapezoid is a convex quadrilateral that has **at least one pair of parallel sides**.
* Parallelograms are also counted (they have two pairs of parallel sides).
* I must consider **every set of 4 points**, but I cannot brute force all `C(n, 4)` because `n` can be up to 500.

The answer must be the **number of unique sets of 4 points** that can form such a quadrilateral.

---

## Constraints

* `4 <= points.length <= 500`
* `-1000 <= xi, yi <= 1000`
* All points are pairwise distinct.

---

## Intuition

First, I tried to imagine the shapes:

* Any quadrilateral (4 points) has 6 segments between them.
* For it to be a trapezoid, at least **one opposite pair of sides is parallel**.

Brute force over all 4-point sets is around `C(500,4) ≈ 2.6e8` – too slow.

So I thought:

1. **Work with segments instead of quadrilaterals.**
   If I know all segments that have the **same slope**, then any two segments that:

   * are **on two different parallel lines**, and
   * are used as opposite sides
     can form one quadrilateral.

2. For each slope:

   * There are many **parallel lines** with that slope.
   * On each line I can have several segments.
   * If a slope has line segment counts `c1, c2, …` on its parallel lines,
     the number of ways to pick two segments on **different** lines is:
     `Σ_{i<j} (ci * cj)`.

3. Summing this over all slopes gives the number of quadrilaterals that have **at least one pair of parallel sides**.

4. But **parallelograms** are special:

   * They have **two pairs** of parallel sides.
   * In the counting above, each parallelogram will appear **twice** (one for each direction).

5. To fix this, I also count using the **full vector** `(dx, dy)` of the segment:

   * Same vector + different lines = two opposite sides of a parallelogram.
   * If I repeat the same formula `Σ_{i<j} ci * cj` but grouping by full vector, I actually count:
     **2 × number of parallelograms**.

6. Final formula:

```text
Answer = (#quadrilaterals with ≥1 parallel pair)
        − (#parallelograms)

        = countBySlope − countByVector / 2
```

7. Also, if all four points lie on a **single line**, they never form a quadrilateral:

   * For that slope, all segments fall on just one `lineId`.
   * So `Σ_{i<j} ci * cj = 0`.
   * No invalid shapes are counted – handled automatically.

---

## Approach

Step-by-step:

1. **Enumerate all point pairs `(i, j)` with `i < j`.**

2. For each pair:

   * Compute:

     ```text
     dx = xj - xi
     dy = yj - yi
     ```

   * Normalize direction (to avoid duplicates with opposite sign):

     * If `dx < 0` or (`dx == 0` and `dy < 0`), multiply both by `-1`.
   * Compute `g = gcd(|dx|, |dy|)` and reduce:

     ```text
     ux = dx / g
     uy = dy / g     (this represents the slope)
     ```

   * Compute a **line identifier**:

     * For direction `(ux, uy)` and point `(x, y)` on the line, value:

       ```text
       lineId = ux * y - uy * x
       ```

       is constant for the whole line.
   * Create two keys:

     * `slopeKey  = encode(ux, uy)`   → slope
     * `vectorKey = encode(dx, dy)`   → full vector

3. **Store counts**:

   * `bySlope[slopeKey][lineId]++`
   * `byVector[vectorKey][lineId]++`

4. **Write a helper `countPairs(map)`**:

   For each outer group (either slope or vector):

   * Take all inner counts `c1, c2, …`.

   * Let `S = Σ ci`, `sumSq = Σ ci²`.

   * The number of **pairs on different lines** is:

     ```text
     Σ_{i<j} ci * cj = (S² - sumSq) / 2
     ```

   * Add this to the global result.

5. Compute:

   ```text
   withParallel      = countPairs(bySlope)
   parallelogramTwo  = countPairs(byVector)  // each parallelogram counted twice
   answer            = withParallel - parallelogramTwo / 2
   ```

6. Use 64-bit integer for counting (C++: `long long`, Java: `long`, etc.).

---

## Data Structures Used

* `unordered_map` / `HashMap` / `Map`:

  * Outer key: encoded pair `(ux, uy)` or `(dx, dy)`.
  * Inner key: `lineId`.
  * Value: count of segments for that slope/vector on that line.

* Basic helper functions:

  * GCD function.
  * Pair encoder: pack `(a, b)` into one integer key.

No heavy advanced structures needed – main cost is just iterating over all `O(n²)` pairs.

---

## Operations & Behavior Summary

For each pair of points:

1. Compute direction `(dx, dy)`.
2. Normalize orientation.
3. Reduce to `(ux, uy)` by gcd.
4. Compute `lineId = ux*y - uy*x`.
5. Update:

   * `bySlope[(ux,uy)][lineId] += 1`
   * `byVector[(dx,dy)][lineId] += 1`

After the loop:

1. For each slope:

   * Combine counts of its lines to get the number of valid **parallel-side pairs**.
2. For each vector:

   * Combine counts to get **2 × parallelograms**.
3. Subtract half of the vector-based count from the slope-based count.

Result is the required number of trapezoids (including parallelograms).

---

## Complexity

Let `n` be the number of points.

* We generate all segments: `m = C(n,2) = n*(n-1)/2` → `O(n²)`.

**Time Complexity**

* Pair enumeration: `O(n²)`.
* All insertions to maps: `O(n²)` average.
* `countPairs` over all groups: sums over all segments again → `O(n²)`.

So overall: **`O(n²)` time**.

**Space Complexity**

* We only store counts of segments grouped in maps.
* Total entries across all maps are proportional to the number of segments, `O(n²)`.

So: **`O(n²)` space**.

---

## Multi-language Solutions

### C++

```c++
#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    using ll = long long;
    // shift used to pack signed integers into single int
    static constexpr int SHIFT = 3000;

    // encode (a,b) into one 32-bit key
    int encodePair(int a, int b) {
        // each value after +SHIFT fits safely in 13 bits (since |a|,|b| <= 2000)
        return ((a + SHIFT) << 13) ^ (b + SHIFT);
    }

    int myGcd(int a, int b) {
        if (a < 0) a = -a;
        if (b < 0) b = -b;
        while (b != 0) {
            int t = a % b;
            a = b;
            b = t;
        }
        return a;
    }

    // count Σ_{i<j} ci*cj for each outer key
    long long countPairs(const unordered_map<int, unordered_map<int,int>>& mp) {
        long long res = 0;
        for (const auto &outer : mp) {
            const auto &inner = outer.second;
            long long sum = 0, sumSq = 0;
            for (auto &kv : inner) {
                long long c = kv.second;
                sum   += c;
                sumSq += c * c;
            }
            res += (sum * sum - sumSq) / 2;
        }
        return res;
    }

    int countTrapezoids(vector<vector<int>>& points) {
        int n = (int)points.size();

        // slopeKey -> (lineId -> count)
        unordered_map<int, unordered_map<int,int>> bySlope;
        // vectorKey -> (lineId -> count)
        unordered_map<int, unordered_map<int,int>> byVector;

        bySlope.reserve(n * n);
        byVector.reserve(n * n);

        for (int i = 0; i < n; ++i) {
            int x1 = points[i][0];
            int y1 = points[i][1];
            for (int j = i + 1; j < n; ++j) {
                int x2 = points[j][0];
                int y2 = points[j][1];

                int dx = x2 - x1;
                int dy = y2 - y1;

                // normalize direction sign
                if (dx < 0 || (dx == 0 && dy < 0)) {
                    dx = -dx;
                    dy = -dy;
                }

                int g = myGcd(dx, dy);
                int ux = dx / g;
                int uy = dy / g;

                // unique id for the supporting line
                int lineId = ux * y1 - uy * x1;

                int slopeKey  = encodePair(ux, uy);
                int vectorKey = encodePair(dx, dy);

                bySlope[slopeKey][lineId]  += 1;
                byVector[vectorKey][lineId] += 1;
            }
        }

        long long withParallel     = countPairs(bySlope);
        long long parallelogramTwo = countPairs(byVector);
        long long ans = withParallel - parallelogramTwo / 2;

        return (int)ans;
    }
};
```

---

### Java

```java
import java.util.*;

class Solution {
    private static final int SHIFT = 3000;

    // encode (a,b) into a single int key
    private int encodePair(int a, int b) {
        return ((a + SHIFT) << 13) ^ (b + SHIFT);
    }

    private int myGcd(int a, int b) {
        a = Math.abs(a);
        b = Math.abs(b);
        while (b != 0) {
            int t = a % b;
            a = b;
            b = t;
        }
        return a;
    }

    private long countPairs(Map<Integer, Map<Integer, Integer>> mp) {
        long res = 0;
        for (Map<Integer, Integer> inner : mp.values()) {
            long sum = 0;
            long sumSq = 0;
            for (int c : inner.values()) {
                long cc = c;
                sum += cc;
                sumSq += cc * cc;
            }
            res += (sum * sum - sumSq) / 2;
        }
        return res;
    }

    public int countTrapezoids(int[][] points) {
        int n = points.length;

        Map<Integer, Map<Integer, Integer>> bySlope  = new HashMap<>();
        Map<Integer, Map<Integer, Integer>> byVector = new HashMap<>();

        for (int i = 0; i < n; ++i) {
            int x1 = points[i][0];
            int y1 = points[i][1];
            for (int j = i + 1; j < n; ++j) {
                int x2 = points[j][0];
                int y2 = points[j][1];

                int dx = x2 - x1;
                int dy = y2 - y1;

                if (dx < 0 || (dx == 0 && dy < 0)) {
                    dx = -dx;
                    dy = -dy;
                }

                int g = myGcd(dx, dy);
                int ux = dx / g;
                int uy = dy / g;

                int lineId = ux * y1 - uy * x1;

                int slopeKey  = encodePair(ux, uy);
                int vectorKey = encodePair(dx, dy);

                bySlope
                    .computeIfAbsent(slopeKey, k -> new HashMap<>())
                    .merge(lineId, 1, Integer::sum);

                byVector
                    .computeIfAbsent(vectorKey, k -> new HashMap<>())
                    .merge(lineId, 1, Integer::sum);
            }
        }

        long withParallel     = countPairs(bySlope);
        long parallelogramTwo = countPairs(byVector);
        long ans = withParallel - parallelogramTwo / 2;

        return (int)ans;
    }
}
```

---

### JavaScript

```javascript
/**
 * @param {number[][]} points
 * @return {number}
 */
var countTrapezoids = function(points) {
    const n = points.length;
    const SHIFT = 3000;

    const encodePair = (a, b) => ((a + SHIFT) << 13) ^ (b + SHIFT);

    const gcd = (a, b) => {
        a = Math.abs(a);
        b = Math.abs(b);
        while (b !== 0) {
            const t = a % b;
            a = b;
            b = t;
        }
        return a;
    };

    // slopeKey -> (lineId -> count)
    const bySlope = new Map();
    // vectorKey -> (lineId -> count)
    const byVector = new Map();

    const addTo = (outer, key, lineId) => {
        if (!outer.has(key)) outer.set(key, new Map());
        const inner = outer.get(key);
        inner.set(lineId, (inner.get(lineId) || 0) + 1);
    };

    for (let i = 0; i < n; ++i) {
        const [x1, y1] = points[i];
        for (let j = i + 1; j < n; ++j) {
            const [x2, y2] = points[j];

            let dx = x2 - x1;
            let dy = y2 - y1;

            if (dx < 0 || (dx === 0 && dy < 0)) {
                dx = -dx;
                dy = -dy;
            }

            const g = gcd(dx, dy);
            const ux = dx / g;
            const uy = dy / g;

            const lineId = ux * y1 - uy * x1;

            const slopeKey  = encodePair(ux, uy);
            const vectorKey = encodePair(dx, dy);

            addTo(bySlope, slopeKey, lineId);
            addTo(byVector, vectorKey, lineId);
        }
    }

    function countPairs(map) {
        let res = 0;
        for (const inner of map.values()) {
            let sum = 0;
            let sumSq = 0;
            for (const c of inner.values()) {
                sum += c;
                sumSq += c * c;
            }
            res += (sum * sum - sumSq) / 2;
        }
        return res;
    }

    const withParallel = countPairs(bySlope);
    const parallelogramTwo = countPairs(byVector);

    return withParallel - Math.floor(parallelogramTwo / 2);
};
```

---

### Python3

```python
from typing import List
from collections import defaultdict
from math import gcd

class Solution:
    def countTrapezoids(self, points: List[List[int]]) -> int:
        n = len(points)
        SHIFT = 3000

        def encode_pair(a: int, b: int) -> int:
            # pack (a,b) into one integer
            return ((a + SHIFT) << 13) ^ (b + SHIFT)

        # slope_key  -> { line_id: count }
        by_slope = defaultdict(lambda: defaultdict(int))
        # vector_key -> { line_id: count }
        by_vector = defaultdict(lambda: defaultdict(int))

        for i in range(n):
            x1, y1 = points[i]
            for j in range(i + 1, n):
                x2, y2 = points[j]
                dx = x2 - x1
                dy = y2 - y1

                if dx < 0 or (dx == 0 and dy < 0):
                    dx = -dx
                    dy = -dy

                g = gcd(dx, dy)
                ux = dx // g
                uy = dy // g

                line_id = ux * y1 - uy * x1

                slope_key = encode_pair(ux, uy)
                vector_key = encode_pair(dx, dy)

                by_slope[slope_key][line_id] += 1
                by_vector[vector_key][line_id] += 1

        def count_pairs(mp) -> int:
            ans = 0
            for inner in mp.values():
                vals = inner.values()
                s = sum(vals)
                sum_sq = sum(c * c for c in vals)
                ans += (s * s - sum_sq) // 2
            return ans

        with_parallel = count_pairs(by_slope)
        parallelogram_twice = count_pairs(by_vector)

        return with_parallel - parallelogram_twice // 2
```

---

### Go

```go
package main

func countTrapezoids(points [][]int) int {
 n := len(points)
 const SHIFT = 3000

 encodePair := func(a, b int) int {
  return ((a + SHIFT) << 13) ^ (b + SHIFT)
 }

 gcd := func(a, b int) int {
  if a < 0 {
   a = -a
  }
  if b < 0 {
   b = -b
  }
  for b != 0 {
   a, b = b, a%b
  }
  return a
 }

 // outer key -> inner map(lineId -> count)
 bySlope := make(map[int]map[int]int)
 byVector := make(map[int]map[int]int)

 addTo := func(mp map[int]map[int]int, key int, lineId int) {
  if mp[key] == nil {
   mp[key] = make(map[int]int)
  }
  mp[key][lineId]++
 }

 for i := 0; i < n; i++ {
  x1, y1 := points[i][0], points[i][1]
  for j := i + 1; j < n; j++ {
   x2, y2 := points[j][0], points[j][1]
   dx := x2 - x1
   dy := y2 - y1

   if dx < 0 || (dx == 0 && dy < 0) {
    dx = -dx
    dy = -dy
   }

   g := gcd(dx, dy)
   ux := dx / g
   uy := dy / g

   lineId := ux*y1 - uy*x1

   slopeKey := encodePair(ux, uy)
   vectorKey := encodePair(dx, dy)

   addTo(bySlope, slopeKey, lineId)
   addTo(byVector, vectorKey, lineId)
  }
 }

 countPairs := func(mp map[int]map[int]int) int64 {
  var res int64
  for _, inner := range mp {
   var sum, sumSq int64
   for _, c := range inner {
    cc := int64(c)
    sum += cc
    sumSq += cc * cc
   }
   res += (sum*sum - sumSq) / 2
  }
  return res
 }

 withParallel := countPairs(bySlope)
 parallelogramTwice := countPairs(byVector)

 ans := withParallel - parallelogramTwice/2
 return int(ans)
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

I’ll explain using the C++ version, but the logic is identical in all languages.

1. **Encoding a pair `(a, b)`**

   ```cpp
   static constexpr int SHIFT = 3000;

   int encodePair(int a, int b) {
       return ((a + SHIFT) << 13) ^ (b + SHIFT);
   }
   ```

   * I shift both numbers by a constant so they become non-negative.
   * Then I pack them into one integer: high bits store `a`, low bits store `b`.
   * This gives me a simple unique key for maps.

2. **GCD helper**

   ```cpp
   int myGcd(int a, int b) {
       if (a < 0) a = -a;
       if (b < 0) b = -b;
       while (b != 0) {
           int t = a % b;
           a = b;
           b = t;
       }
       return a;
   }
   ```

   * Standard Euclid’s algorithm.
   * Used to reduce direction vectors.

3. **Counting pairs across different lines**

   ```cpp
   long long countPairs(const unordered_map<int, unordered_map<int,int>>& mp) {
       long long res = 0;
       for (const auto &outer : mp) {
           const auto &inner = outer.second;
           long long sum = 0, sumSq = 0;
           for (auto &kv : inner) {
               long long c = kv.second;
               sum   += c;
               sumSq += c * c;
           }
           res += (sum * sum - sumSq) / 2;
       }
       return res;
   }
   ```

   * For each slope (or vector) we have an inner map from lineId → count.
   * `sum` is total segments for this slope.
   * `sumSq` is sum of squares of counts per line.
   * `(sum² - sumSq) / 2` equals `Σ_{i<j} ci * cj` (pairs in different lines).

4. **Filling the maps**

   ```cpp
   for (int i = 0; i < n; ++i) {
       int x1 = points[i][0];
       int y1 = points[i][1];
       for (int j = i + 1; j < n; ++j) {
           int x2 = points[j][0];
           int y2 = points[j][1];

           int dx = x2 - x1;
           int dy = y2 - y1;

           if (dx < 0 || (dx == 0 && dy < 0)) {
               dx = -dx;
               dy = -dy;
           }

           int g = myGcd(dx, dy);
           int ux = dx / g;
           int uy = dy / g;

           int lineId = ux * y1 - uy * x1;

           int slopeKey  = encodePair(ux, uy);
           int vectorKey = encodePair(dx, dy);

           bySlope[slopeKey][lineId]  += 1;
           byVector[vectorKey][lineId] += 1;
       }
   }
   ```

   * For each segment:

     * Normalize the vector.
     * Get reduced direction `(ux, uy)` and lineId for that slope.
     * Update both maps.

5. **Final answer**

   ```cpp
   long long withParallel     = countPairs(bySlope);
   long long parallelogramTwo = countPairs(byVector);
   long long ans = withParallel - parallelogramTwo / 2;
   return (int)ans;
   ```

   * `withParallel` = all quadrilaterals with at least one parallel side pair (parallelograms counted twice).
   * `parallelogramTwo / 2` = exact number of parallelograms (each counted twice).
   * Subtract to get each quadrilateral exactly once.

In Java / JS / Python / Go, the structure is the same:

* Same two maps,
* Same formula,
* Same GCD and direction normalization.

---

## Examples

### Example 1

Input:

```text
points = [[-3,2],[3,0],[2,3],[3,2],[2,-3]]
```

Output:

```text
2
```

Explanation (high-level):

* There are exactly two different sets of 4 points among these that form trapezoids.

---

### Example 2

Input:

```text
points = [[0,0],[1,0],[0,1],[2,1]]
```

Output:

```text
1
```

Explanation:

* Only one trapezoid can be formed from these four points.

---

### Example 3 (All collinear – should give 0)

Input:

```text
points = [[82,7],[82,-9],[82,-52],[82,78]]
```

Output:

```text
0
```

Explanation:

* All points lie on a single vertical line.
* No quadrilateral can be formed, so answer is 0.
* Our algorithm handles this naturally, because for that slope we only have one `lineId` so we never count any pair of segments across different lines.

---

## How to use / Run locally

### C++

```bash
g++ -std=c++17 -O2 main.cpp -o main
./main
```

Make sure `main.cpp` contains the `Solution` class and a tiny driver if you want to test locally.

### Java

```bash
javac Solution.java
java Solution
```

Place the class `Solution` in `Solution.java`.

### JavaScript (Node.js)

```bash
node main.js
```

`main.js` should export/use the `countTrapezoids` function and run some test cases.

### Python3

```bash
python3 main.py
```

`main.py` can import the `Solution` class and call `Solution().countTrapezoids(points)`.

### Go

```bash
go run main.go
```

Make sure the `countTrapezoids` function is in `main` package and you call it from `main()`.

---

## Notes & Optimizations

* I use `O(n²)` time and space.
  Since `n <= 500`, `n² = 250000` which is very manageable.
* Using encoded integer keys instead of tuples helps maps stay fast.
* I do **not** recompute slopes or lengths multiple times; each pair is processed only once.
* All parallelogram overcount is handled mathematically via the second pass (`byVector`), keeping the code clean.
* This approach safely avoids:

  * Collinear 4-tuples,
  * Double counting,
  * Numeric issues with floating point (I only use integers and gcd).

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
