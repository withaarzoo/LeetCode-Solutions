# Maximum Square Area by Removing Fences From a Field

LeetCode Problem 2975

---

## Table of Contents

* Problem Summary
* Constraints
* Intuition
* Approach
* Data Structures Used
* Operations & Behavior Summary
* Complexity
* Multi-language Solutions

  * C++
  * Java
  * JavaScript
  * Python3
  * Go
* Step-by-step Detailed Explanation
* Examples
* How to Use / Run Locally
* Notes & Optimizations
* Author

---

## Problem Summary

I am given a large rectangular field of size `(m - 1) × (n - 1)`.

Some **horizontal fences** and **vertical fences** are placed inside the field.
The **outer boundary fences cannot be removed**, but all other fences can be removed.

My task is to **remove some fences (or none)** and create the **largest possible square field**.

If it is possible, I return the **maximum square area** modulo `10^9 + 7`.
If it is **not possible**, I return `-1`.

---

## Constraints

* `3 ≤ m, n ≤ 10^9`
* `1 ≤ hFences.length, vFences.length ≤ 600`
* `1 < hFences[i] < m`
* `1 < vFences[i] < n`
* All fence positions are unique

---

## Intuition

When I first read the problem, I understood one key thing.

A square needs:

* equal height
* equal width

Since I can remove fences, the only thing that matters is the **distance between two remaining fences**.

So I thought:

* Any two horizontal fences can give me a possible height
* Any two vertical fences can give me a possible width

If the **same distance exists in both directions**, I can form a square.

My goal became very simple:

> Find the **largest distance** that exists in both horizontal and vertical directions.

---

## Approach

1. I first add the **boundary fences**

   * Horizontal → `1` and `m`
   * Vertical → `1` and `n`

2. I sort both fence arrays.

3. I calculate **all possible distances** between every pair of horizontal fences.

4. I do the same for vertical fences.

5. I store these distances in sets for fast lookup.

6. I find the **largest common distance** in both sets.

7. If no common distance exists, I return `-1`.

8. Otherwise, I return `distance × distance % (10^9 + 7)`.

---

## Data Structures Used

* Arrays / Lists
* Hash Set / Set
* Sorting

---

## Operations & Behavior Summary

* Fence removal is simulated by selecting distances between remaining fences
* Distance calculation is done using pairwise differences
* Sets are used to avoid duplicate distances
* Final result is calculated using modular arithmetic

---

## Complexity

### Time Complexity

`O(H² + V²)`

* `H` = number of horizontal fences (≤ 602 including boundaries)
* `V` = number of vertical fences (≤ 602 including boundaries)

This works because the limits are small.

### Space Complexity

`O(H² + V²)`

Used to store all possible distances.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int maximizeSquareArea(int m, int n, vector<int>& hFences, vector<int>& vFences) {
        const long long MOD = 1e9 + 7;

        hFences.push_back(1);
        hFences.push_back(m);
        vFences.push_back(1);
        vFences.push_back(n);

        sort(hFences.begin(), hFences.end());
        sort(vFences.begin(), vFences.end());

        unordered_set<long long> hSet, vSet;

        for (int i = 0; i < hFences.size(); i++)
            for (int j = i + 1; j < hFences.size(); j++)
                hSet.insert(hFences[j] - hFences[i]);

        for (int i = 0; i < vFences.size(); i++)
            for (int j = i + 1; j < vFences.size(); j++)
                vSet.insert(vFences[j] - vFences[i]);

        long long maxSide = 0;
        for (auto d : hSet)
            if (vSet.count(d))
                maxSide = max(maxSide, d);

        if (maxSide == 0) return -1;
        return (maxSide * maxSide) % MOD;
    }
};
```

---

### Java

```java
class Solution {
    public int maximizeSquareArea(int m, int n, int[] hFences, int[] vFences) {
        long MOD = 1000000007L;

        List<Integer> h = new ArrayList<>();
        List<Integer> v = new ArrayList<>();

        for (int x : hFences) h.add(x);
        for (int x : vFences) v.add(x);

        h.add(1); h.add(m);
        v.add(1); v.add(n);

        Collections.sort(h);
        Collections.sort(v);

        Set<Long> hs = new HashSet<>();
        Set<Long> vs = new HashSet<>();

        for (int i = 0; i < h.size(); i++)
            for (int j = i + 1; j < h.size(); j++)
                hs.add((long) h.get(j) - h.get(i));

        for (int i = 0; i < v.size(); i++)
            for (int j = i + 1; j < v.size(); j++)
                vs.add((long) v.get(j) - v.get(i));

        long maxSide = 0;
        for (long d : hs)
            if (vs.contains(d))
                maxSide = Math.max(maxSide, d);

        if (maxSide == 0) return -1;
        return (int) ((maxSide * maxSide) % MOD);
    }
}
```

---

### JavaScript (BigInt Safe)

```javascript
var maximizeSquareArea = function (m, n, hFences, vFences) {
    const MOD = 1000000007n;

    hFences.push(1, m);
    vFences.push(1, n);

    hFences.sort((a, b) => a - b);
    vFences.sort((a, b) => a - b);

    const hSet = new Set();
    const vSet = new Set();

    for (let i = 0; i < hFences.length; i++)
        for (let j = i + 1; j < hFences.length; j++)
            hSet.add(hFences[j] - hFences[i]);

    for (let i = 0; i < vFences.length; i++)
        for (let j = i + 1; j < vFences.length; j++)
            vSet.add(vFences[j] - vFences[i]);

    let maxSide = 0;
    for (let d of hSet)
        if (vSet.has(d))
            maxSide = Math.max(maxSide, d);

    if (maxSide === 0) return -1;

    const side = BigInt(maxSide);
    return Number((side * side) % MOD);
};
```

---

### Python3

```python
class Solution:
    def maximizeSquareArea(self, m: int, n: int, hFences: list[int], vFences: list[int]) -> int:
        MOD = 10**9 + 7

        hFences += [1, m]
        vFences += [1, n]

        hFences.sort()
        vFences.sort()

        hs, vs = set(), set()

        for i in range(len(hFences)):
            for j in range(i + 1, len(hFences)):
                hs.add(hFences[j] - hFences[i])

        for i in range(len(vFences)):
            for j in range(i + 1, len(vFences)):
                vs.add(vFences[j] - vFences[i])

        maxSide = 0
        for d in hs:
            if d in vs:
                maxSide = max(maxSide, d)

        if maxSide == 0:
            return -1

        return (maxSide * maxSide) % MOD
```

---

### Go

```go
func maximizeSquareArea(m int, n int, hFences []int, vFences []int) int {
    const MOD int64 = 1000000007

    hFences = append(hFences, 1, m)
    vFences = append(vFences, 1, n)

    sort.Ints(hFences)
    sort.Ints(vFences)

    hs := make(map[int]bool)
    vs := make(map[int]bool)

    for i := 0; i < len(hFences); i++ {
        for j := i + 1; j < len(hFences); j++ {
            hs[hFences[j]-hFences[i]] = true
        }
    }

    for i := 0; i < len(vFences); i++ {
        for j := i + 1; j < len(vFences); j++ {
            vs[vFences[j]-vFences[i]] = true
        }
    }

    maxSide := 0
    for d := range hs {
        if vs[d] && d > maxSide {
            maxSide = d
        }
    }

    if maxSide == 0 {
        return -1
    }

    return int((int64(maxSide) * int64(maxSide)) % MOD)
}
```

---

## Step-by-step Detailed Explanation

* I include boundary fences so no valid distance is missed
* I sort fences to simplify difference calculation
* I calculate every possible distance
* I use sets to avoid duplicates
* I search for the largest common distance
* That distance becomes the square side
* Area is calculated safely using modulo

---

## Examples

Input
`m = 4, n = 3`
`hFences = [2,3]`
`vFences = [2]`

Output
`4`

---

Input
`m = 6, n = 7`
`hFences = [2]`
`vFences = [4]`

Output
`-1`

---

## How to Use / Run Locally

1. Clone the repository
2. Open the file for your language
3. Compile or run using standard tools

Example (C++)

```bash
g++ solution.cpp
./a.out
```

Example (Python)

```bash
python3 solution.py
```

---

## Notes & Optimizations

* JavaScript requires **BigInt** due to precision limits
* Other languages handle large integers safely
* This approach is optimal for given constraints
* Works for all edge cases

---

## Author

**Md Aarzoo Islam**
[https://bento.me/withaarzoo](https://bento.me/withaarzoo)
