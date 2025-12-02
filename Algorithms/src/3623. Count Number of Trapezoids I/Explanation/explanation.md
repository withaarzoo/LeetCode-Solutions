# 3623. Count Number of Trapezoids I

* **Problem Title:** Count Number of Trapezoids I

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

We are given `n` distinct points on a 2D plane:

```text
points[i] = [xi, yi]
```

A **horizontal trapezoid** is a convex quadrilateral that has **at least one pair of horizontal sides**, i.e., two sides are parallel to the x-axis.

We need to count how many **unique horizontal trapezoids** can be formed by choosing any 4 distinct points from `points`.

Because the answer can be large, we return it modulo `1e9 + 7`.

---

## Constraints

* `4 <= points.length <= 1e5`
* `-1e8 <= xi, yi <= 1e8`
* All points are pairwise distinct.
* Return answer modulo `10^9 + 7`.

---

## Intuition

I thought:

* For a side to be **horizontal**, the two points must have the **same y-coordinate**.
* A horizontal trapezoid with horizontal bases basically means:

  * I choose 2 points on one horizontal line `y = y1`.
  * I choose 2 points on another horizontal line `y = y2`, with `y1 != y2`.
* Any such 4 points (2 on each line) will form a quadrilateral where opposite sides are parallel to x-axis → a horizontal trapezoid.

So the problem becomes **pure counting**:

> Group points by `y`.
> For each `y`, count how many pairs of points we can form on that line.
> Then for every pair of different `y`-lines, multiply their pair-counts.

---

## Approach

1. **Group points by y-coordinate**

   * Use a hash map `freq[y]` to store how many points lie on horizontal line `y`.
   * One pass over all points.

2. **Compute number of point-pairs on each horizontal line**

   * If a line has `c` points, number of ways to choose 2 points is:
     [
     f(y) = C(c, 2) = \frac{c(c-1)}{2}
     ]
   * If `c < 2`, `f(y) = 0` (no segment possible).
   * Maintain:

     * `S = sum of all f(y)`
     * `SQ = sum of all f(y)^2`

3. **Count trapezoids from two different lines**

   * For two distinct lines with pair counts `f(a)` and `f(b)`:

     * Number of trapezoids using these two lines:
       [
       f(a) \times f(b)
       ]
   * We need:
     [
     \sum_{a < b} f(a) f(b)
     ]
   * Use algebra to avoid `O(m^2)` double loop:
     [
     \left( \sum f(y) \right)^2
     = \sum f(y)^2 + 2 \sum_{a < b} f(a)f(b)
     ]
     So:
     [
     \sum_{a < b} f(a)f(b) = \frac{S^2 - SQ}{2}
     ]

4. **Modulo arithmetic**

   * All operations are done modulo `MOD = 1e9 + 7`.
   * Division by 2 is handled with modular inverse:
     `inv2 = (MOD + 1) / 2` (since MOD is prime).

5. **Return the final count**

   * Final answer = `((S^2 - SQ) / 2) mod MOD`.

---

## Data Structures Used

* **Hash map / dictionary**

  * Key: `y-coordinate`
  * Value: count of points having this `y`.
* All other values are stored in primitive integers / longs / BigInt.

---

## Operations & Behavior Summary

* **Counting points per y:**

  * Iterate once over all points, update frequency map.

* **Computing combinations on each line:**

  * For each `y` in map, compute `C(c, 2)` if `c >= 2`.

* **Combining all lines:**

  * Use `S` and `SQ` to compute total trapezoids with the formula.

* **Return answer modulo `1e9 + 7`.**

---

## Complexity

Let `n` = number of points, `m` = number of distinct y-coordinates.

* **Time Complexity:**

  * Building frequency map: `O(n)`
  * Looping over distinct y’s: `O(m)`
  * Overall: **O(n)** (since `m <= n`).

* **Space Complexity:**

  * Hash map storing counts for each unique y: `O(m)` extra space.

---

## Multi-language Solutions

### C++

```c++
#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    int countTrapezoids(vector<vector<int>>& points) {
        const long long MOD = 1000000007LL;
        const long long INV2 = (MOD + 1) / 2; // modular inverse of 2

        // 1. Count how many points lie on each y-coordinate
        unordered_map<long long, int> freq;
        freq.reserve(points.size() * 2);
        for (auto &p : points) {
            long long y = p[1];
            ++freq[y];
        }

        long long sumF = 0;   // S = sum of f(y)
        long long sumF2 = 0;  // SQ = sum of f(y)^2

        // 2. For each y, compute C(c,2) and update S, SQ
        for (auto &kv : freq) {
            long long c = kv.second;
            if (c >= 2) {
                long long f = c * (c - 1) / 2 % MOD; // f(y) = C(c,2)
                sumF = (sumF + f) % MOD;
                sumF2 = (sumF2 + f * f % MOD) % MOD;
            }
        }

        // 3. Answer = ((S^2 - SQ) / 2) mod MOD
        long long ans = (sumF * sumF) % MOD;         // S^2
        ans = (ans - sumF2 + MOD) % MOD;             // S^2 - SQ
        ans = ans * INV2 % MOD;                      // divide by 2

        return (int)ans;
    }
};
```

---

### Java

```java
import java.util.*;

class Solution {
    public int countTrapezoids(int[][] points) {
        final long MOD = 1_000_000_007L;
        final long INV2 = (MOD + 1) / 2; // modular inverse of 2

        // 1. Count points per y-coordinate
        HashMap<Integer, Integer> freq = new HashMap<>();
        for (int[] p : points) {
            int y = p[1];
            freq.put(y, freq.getOrDefault(y, 0) + 1);
        }

        long sumF = 0;   // S
        long sumF2 = 0;  // SQ

        // 2. For each y, compute C(c,2) and accumulate
        for (int c : freq.values()) {
            if (c >= 2) {
                long cc = c;
                long f = (cc * (cc - 1) / 2) % MOD; // C(c,2)
                sumF = (sumF + f) % MOD;
                sumF2 = (sumF2 + f * f % MOD) % MOD;
            }
        }

        // 3. Answer = ((S^2 - SQ) / 2) mod MOD
        long ans = (sumF * sumF) % MOD;
        ans = (ans - sumF2 + MOD) % MOD;
        ans = ans * INV2 % MOD;

        return (int) ans;
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
    const MOD = 1000000007n;
    const INV2 = (MOD + 1n) / 2n;   // modular inverse of 2

    // 1. Count points per y
    const freq = new Map();
    for (const [x, y] of points) {
        const key = BigInt(y);
        freq.set(key, (freq.get(key) || 0n) + 1n);
    }

    let sumF = 0n;   // S
    let sumF2 = 0n;  // SQ

    // 2. For each y, compute C(c,2)
    for (const c of freq.values()) {
        if (c >= 2n) {
            const f = (c * (c - 1n) / 2n) % MOD;  // C(c,2)
            sumF = (sumF + f) % MOD;
            sumF2 = (sumF2 + (f * f) % MOD) % MOD;
        }
    }

    // 3. Answer = ((S^2 - SQ) / 2) mod MOD
    let ans = (sumF * sumF) % MOD;           // S^2
    ans = (ans - sumF2 + MOD) % MOD;         // S^2 - SQ
    ans = (ans * INV2) % MOD;                // divide by 2

    return Number(ans);                      // LeetCode expects number
};
```

---

### Python3

```python
from typing import List
from collections import defaultdict

class Solution:
    def countTrapezoids(self, points: List[List[int]]) -> int:
        MOD = 10**9 + 7
        INV2 = (MOD + 1) // 2  # modular inverse of 2

        # 1. Count points per y-coordinate
        freq = defaultdict(int)
        for x, y in points:
            freq[y] += 1

        sumF = 0  # S
        sumF2 = 0 # SQ

        # 2. For each y, compute C(c,2) and accumulate
        for c in freq.values():
            if c >= 2:
                f = c * (c - 1) // 2   # C(c,2)
                f %= MOD
                sumF = (sumF + f) % MOD
                sumF2 = (sumF2 + f * f) % MOD

        # 3. Answer = ((S^2 - SQ) / 2) mod MOD
        ans = (sumF * sumF) % MOD
        ans = (ans - sumF2 + MOD) % MOD
        ans = ans * INV2 % MOD

        return ans
```

---

### Go

```go
package main

func countTrapezoids(points [][]int) int {
 const MOD int64 = 1000000007
 const INV2 int64 = (MOD + 1) / 2 // modular inverse of 2

 // 1. Count points per y-coordinate
 freq := make(map[int]int)
 for _, p := range points {
  y := p[1]
  freq[y]++
 }

 var sumF int64 = 0  // S
 var sumF2 int64 = 0 // SQ

 // 2. For each y, compute C(c,2) and accumulate
 for _, c := range freq {
  if c >= 2 {
   cc := int64(c)
   f := (cc * (cc - 1) / 2) % MOD // C(c,2)
   sumF = (sumF + f) % MOD
   sumF2 = (sumF2 + (f*f)%MOD) % MOD
  }
 }

 // 3. Answer = ((S^2 - SQ) / 2) mod MOD
 ans := (sumF * sumF) % MOD
 ans = (ans - sumF2 + MOD) % MOD
 ans = (ans * INV2) % MOD

 return int(ans)
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

I’ll outline the key steps; all languages follow the same logic.

### 1. Build frequency map of y-coordinates

**Idea:**
All points with the same `y` lie on the same horizontal line.

Example (C++ line):

```c++
unordered_map<long long, int> freq;
for (auto &p : points) {
    long long y = p[1];
    ++freq[y];
}
```

* Loop over all points.
* `freq[y]` stores how many points are on that line.

Same thing happens in Java (`HashMap`), Python (`defaultdict`), JavaScript (`Map`), and Go (`map[int]int`).

---

### 2. Compute number of pairs on each line

For each line with `c` points:

```text
f(y) = C(c, 2) = c * (c - 1) / 2
```

Code pattern (Python):

```python
for c in freq.values():
    if c >= 2:
        f = c * (c - 1) // 2
        f %= MOD
        sumF = (sumF + f) % MOD
        sumF2 = (sumF2 + f * f) % MOD
```

* If less than 2 points → skip.
* `sumF` accumulates sum of all `f(y)`.
* `sumF2` accumulates sum of all `f(y)^2`.

Same formula is used in all other languages with appropriate integer types.

---

### 3. Use algebra to count trapezoids

We want:

```text
Total = sum over all a < b of f(a) * f(b)
```

Using:

```text
S = sum f(y)
SQ = sum f(y)^2
S^2 = SQ + 2 * Total
=> Total = (S^2 - SQ) / 2
```

Example (Java):

```java
long ans = (sumF * sumF) % MOD;  // S^2
ans = (ans - sumF2 + MOD) % MOD; // S^2 - SQ
ans = ans * INV2 % MOD;          // divide by 2
```

* `+ MOD` before `% MOD` ensures we never go negative.
* Division by 2 is replaced by multiplying with `INV2` (modular inverse of 2).

Exactly same idea in C++/Python/JS/Go.

---

### 4. Final answer

All languages finally return `ans` as an int-type, which is already modulo `1e9 + 7`.

---

## Examples

### Example 1

**Input:**

```text
points = [[1,0],[2,0],[3,0],[2,2],[3,2]]
```

* Line `y=0` has points: `(1,0), (2,0), (3,0)` → `C(3,2) = 3` pairs.
* Line `y=2` has points: `(2,2), (3,2)`     → `C(2,2) = 1` pair.

So:

```text
S  = 3 + 1 = 4
SQ = 3^2 + 1^2 = 10
Total = (S^2 - SQ) / 2 = (16 - 10) / 2 = 3
```

**Output:**

```text
3
```

---

### Example 2

**Input:**

```text
points = [[0,0],[1,0],[0,1],[2,1]]
```

* `y=0`: 2 points → `C(2,2) = 1` pair
* `y=1`: 2 points → `C(2,2) = 1` pair

So:

```text
S = 1 + 1 = 2
SQ = 1^2 + 1^2 = 2
Total = (2^2 - 2)/2 = (4-2)/2 = 1
```

**Output:**

```text
1
```

---

## How to use / Run locally

### C++

```bash
g++ -std=c++17 -O2 main.cpp -o main
./main
```

* Put the `Solution` class in `main.cpp`.
* Inside `main()`, create `Solution` object and call `countTrapezoids(points)`.

---

### Java

```bash
javac Solution.java
java Solution
```

* Ensure `Solution` class is public and in `Solution.java`.

---

### JavaScript (Node.js)

```bash
node main.js
```

* Define the `countTrapezoids` function in `main.js`.
* Call it with sample input to test.

---

### Python3

```bash
python3 main.py
```

* Place the `Solution` class in `main.py`.
* Create object and call `Solution().countTrapezoids(points)`.

---

### Go

```bash
go run main.go
```

* Put function `countTrapezoids` in `main.go`.
* In `main()`, prepare input and print the result.

---

## Notes & Optimizations

* Using a hash map lets us handle large ranges of `y` (since `yi` can be up to `1e8` and negative).
* We avoid `O(m^2)` nested loops by using the algebraic identity, so solution easily works for `n = 1e5`.
* All arithmetic is done modulo `1e9 + 7` to avoid overflow and follow problem statement.
* In JavaScript we used `BigInt` to keep numbers precise during modulo operations.

---

## Author

* **[Md Aarzoo Islam](https://bento.me/withaarzoo)** ✨
