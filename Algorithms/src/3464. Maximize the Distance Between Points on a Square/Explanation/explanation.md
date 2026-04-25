# 3464. Maximize the Distance Between Points on a Square

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
* [Step-by-step Detailed Explanation](#step-by-step-detailed-explanation-c-java-javascript-python3-go)
* [Examples](#examples)
* [How to use / Run locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

## Problem Summary

I am given:

* `side`: the edge length of a square.
* `points`: a list of unique points on the boundary of the square.
* `k`: the number of points I must choose.

My goal is to select exactly `k` points so that the **minimum Manhattan distance** between any two chosen points is as large as possible.

In simple words, I want to place `k` chosen points on the square boundary so that the closest pair among them is as far apart as possible.

## Constraints

* `1 <= side <= 10^9`
* `4 <= points.length <= min(4 * side, 15 * 10^3)`
* `points[i] == [xi, yi]`
* Every point lies on the boundary of the square.
* All points are unique.
* `4 <= k <= min(25, points.length)`

## Intuition

I first noticed that every point lies on the boundary of the same square.

That means I can walk around the square in clockwise order and treat the boundary like one circular line.

If I fix a distance `D`, I can ask:

"Starting from one point, what is the next point clockwise that is at least `D` away from it?"

If I can answer that quickly, then I can greedily keep jumping from one valid point to the next.

So my plan is:

1. Sort points in clockwise order.
2. For a fixed `D`, greedily try to build a set of `k` points.
3. Use binary search on `D` to find the largest possible answer.

## Approach

### 1. Convert boundary points to a clockwise position

I map every boundary point to a single number representing its position on the square perimeter.

Clockwise order from `(0, 0)`:

* Left side: `(0, y)`  -> `y`
* Top side: `(x, side)` -> `side + x`
* Right side: `(side, y)` -> `3 * side - y`
* Bottom side: `(x, 0)` -> `4 * side - x`

After this conversion, all points can be sorted by that perimeter position.

### 2. Check if a distance `D` is possible

For a candidate answer `D`, I need to know whether I can select `k` points such that every consecutive chosen pair is at least `D` apart.

I use a greedy idea:

* Start from one point.
* Jump to the first point clockwise that satisfies the distance condition.
* Repeat until I pick `k` points.

### 3. Make jumping fast

To make this efficient, I duplicate the sorted perimeter list multiple times.

That helps me handle wrap-around on the square without special cases.

Then for each point, I compute the earliest next point whose Manhattan distance is at least `D`.

This lets me quickly chain selections.

### 4. Binary search the answer

The minimum distance has a monotonic property:

* If distance `D` is possible, then every smaller distance is also possible.
* If distance `D` is not possible, then every larger distance is also not possible.

So I binary search the maximum valid value.

## Data Structures Used

* **Array / List**: to store transformed points.
* **Sorting**: to arrange points in clockwise order.
* **Binary Search**: to search the best answer and to find the next valid point.
* **Greedy traversal**: to test feasibility for a fixed distance.

## Operations & Behavior Summary

* Transform each boundary point into a perimeter index.
* Sort all points by that index.
* For a candidate distance `D`, compute the next valid clockwise point.
* Greedily select up to `k` points.
* Verify whether the first and last selected points also satisfy the distance condition.
* Binary search the best answer.

## Complexity

Let `n = points.length`.

* **Time Complexity:** `O(n log n + n * k * log n * log side)`

  * `O(n log n)` for sorting.
  * `O(log side)` for binary search on the answer.
  * For each test, I try greedy selection with binary searches.
  * Since `k <= 25`, this is efficient enough.

* **Space Complexity:** `O(n)`

  * I store the transformed points and helper arrays.

## Multi-language Solutions

### C++

```cpp
class Solution {
    struct P {
        long long pos;
        long long x, y;
    };

    long long getOffset(long long s, long long x, long long y, long long d) {
        // Return how far clockwise I must move from this point
        // to find the first point with Manhattan distance >= d.
        // If impossible, return -1.

        if (x == 0) { // left side
            if (d <= 2 * s - y) return d;
            if (d <= s + y) return 2 * s + d - 2 * y;
            return -1;
        } else if (y == s) { // top side
            if (d <= 2 * s - x) return d;
            if (d <= s + x) return 2 * s + d - 2 * x;
            return -1;
        } else if (x == s) { // right side
            if (d <= s + y) return d;
            if (d <= 2 * s - y) return d + 2 * y;
            return -1;
        } else { // bottom side
            if (d <= s + x) return d;
            if (d <= 2 * s - x) return d + 2 * x;
            return -1;
        }
    }

    bool can(long long side, vector<P>& pts, int k, long long d) {
        int n = (int)pts.size();

        vector<long long> pos3(3 * n);
        for (int i = 0; i < n; i++) {
            pos3[i] = pts[i].pos;
            pos3[i + n] = pts[i].pos + 4LL * side;
            pos3[i + 2 * n] = pts[i].pos + 8LL * side;
        }

        // nxt[i] = next selected point index in the 3-copy array
        // from duplicate position i (only need i in [0, 2n)).
        vector<int> nxt(2 * n, -1);

        for (int i = 0; i < 2 * n; i++) {
            const P& p = pts[i % n];
            long long off = getOffset(side, p.x, p.y, d);
            if (off < 0) continue;

            long long target = pos3[i] + off;
            int hi = min(i + n, 3 * n);
            auto it = lower_bound(pos3.begin() + i + 1, pos3.begin() + hi, target);
            if (it != pos3.begin() + hi) {
                nxt[i] = (int)(it - pos3.begin());
            }
        }

        for (int start = 0; start < n; start++) {
            int cur = start;
            int cnt = 1;

            while (cnt < k) {
                cur = nxt[cur];
                if (cur == -1 || cur >= start + n) break;
                cnt++;
            }

            if (cnt >= k) {
                long long dx = llabs(pts[start].x - pts[cur % n].x);
                long long dy = llabs(pts[start].y - pts[cur % n].y);
                if (dx + dy >= d) return true;
            }
        }

        return false;
    }

public:
    int maxDistance(int side, vector<vector<int>>& points, int k) {
        vector<P> pts;
        pts.reserve(points.size());

        for (auto &v : points) {
            long long x = v[0], y = v[1];
            long long pos;
            if (x == 0) pos = y;
            else if (y == side) pos = 1LL * side + x;
            else if (x == side) pos = 3LL * side - y;
            else pos = 4LL * side - x;
            pts.push_back({pos, x, y});
        }

        sort(pts.begin(), pts.end(), [](const P& a, const P& b) {
            return a.pos < b.pos;
        });

        long long lo = 0, hi = 2LL * side;
        while (lo < hi) {
            long long mid = (lo + hi + 1) >> 1;
            if (can(side, pts, k, mid)) lo = mid;
            else hi = mid - 1;
        }

        return (int)lo;
    }
};
```

### Java

```java
import java.util.*;

class Solution {
    static class P {
        long pos, x, y;
        P(long pos, long x, long y) {
            this.pos = pos;
            this.x = x;
            this.y = y;
        }
    }

    private long getOffset(long side, long x, long y, long d) {
        if (x == 0) { // left
            if (d <= 2L * side - y) return d;
            if (d <= side + y) return 2L * side + d - 2L * y;
            return -1;
        } else if (y == side) { // top
            if (d <= 2L * side - x) return d;
            if (d <= side + x) return 2L * side + d - 2L * x;
            return -1;
        } else if (x == side) { // right
            if (d <= side + y) return d;
            if (d <= 2L * side - y) return d + 2L * y;
            return -1;
        } else { // bottom
            if (d <= side + x) return d;
            if (d <= 2L * side - x) return d + 2L * x;
            return -1;
        }
    }

    private int lowerBound(long[] arr, int l, int r, long target) {
        while (l < r) {
            int m = (l + r) >>> 1;
            if (arr[m] < target) l = m + 1;
            else r = m;
        }
        return l;
    }

    private boolean can(long side, P[] pts, int k, long d) {
        int n = pts.length;

        long[] pos3 = new long[3 * n];
        for (int i = 0; i < n; i++) {
            pos3[i] = pts[i].pos;
            pos3[i + n] = pts[i].pos + 4L * side;
            pos3[i + 2 * n] = pts[i].pos + 8L * side;
        }

        int[] nxt = new int[2 * n];
        Arrays.fill(nxt, -1);

        for (int i = 0; i < 2 * n; i++) {
            P p = pts[i % n];
            long off = getOffset(side, p.x, p.y, d);
            if (off < 0) continue;

            long target = pos3[i] + off;
            int hi = Math.min(i + n, 3 * n);
            int j = lowerBound(pos3, i + 1, hi, target);
            if (j < hi) nxt[i] = j;
        }

        for (int start = 0; start < n; start++) {
            int cur = start;
            int cnt = 1;

            while (cnt < k) {
                cur = nxt[cur];
                if (cur == -1 || cur >= start + n) break;
                cnt++;
            }

            if (cnt >= k) {
                long dx = Math.abs(pts[start].x - pts[cur % n].x);
                long dy = Math.abs(pts[start].y - pts[cur % n].y);
                if (dx + dy >= d) return true;
            }
        }

        return false;
    }

    public int maxDistance(int side, int[][] points, int k) {
        P[] pts = new P[points.length];

        for (int i = 0; i < points.length; i++) {
            long x = points[i][0];
            long y = points[i][1];
            long pos;
            if (x == 0) pos = y;
            else if (y == side) pos = 1L * side + x;
            else if (x == side) pos = 3L * side - y;
            else pos = 4L * side - x;
            pts[i] = new P(pos, x, y);
        }

        Arrays.sort(pts, Comparator.comparingLong(a -> a.pos));

        long lo = 0, hi = 2L * side;
        while (lo < hi) {
            long mid = (lo + hi + 1) >>> 1;
            if (can(side, pts, k, mid)) lo = mid;
            else hi = mid - 1;
        }

        return (int)lo;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number} side
 * @param {number[][]} points
 * @param {number} k
 * @return {number}
 */
var maxDistance = function(side, points, k) {
    const pts = points.map(([x, y]) => {
        let pos;
        if (x === 0) pos = y;
        else if (y === side) pos = side + x;
        else if (x === side) pos = 3 * side - y;
        else pos = 4 * side - x;
        return { pos, x, y };
    });

    pts.sort((a, b) => a.pos - b.pos);

    function getOffset(x, y, d) {
        if (x === 0) {
            if (d <= 2 * side - y) return d;
            if (d <= side + y) return 2 * side + d - 2 * y;
            return -1;
        } else if (y === side) {
            if (d <= 2 * side - x) return d;
            if (d <= side + x) return 2 * side + d - 2 * x;
            return -1;
        } else if (x === side) {
            if (d <= side + y) return d;
            if (d <= 2 * side - y) return d + 2 * y;
            return -1;
        } else {
            if (d <= side + x) return d;
            if (d <= 2 * side - x) return d + 2 * x;
            return -1;
        }
    }

    function lowerBound(arr, l, r, target) {
        while (l < r) {
            const m = (l + r) >> 1;
            if (arr[m] < target) l = m + 1;
            else r = m;
        }
        return l;
    }

    function can(d) {
        const n = pts.length;
        const pos3 = new Array(3 * n);
        for (let i = 0; i < n; i++) {
            pos3[i] = pts[i].pos;
            pos3[i + n] = pts[i].pos + 4 * side;
            pos3[i + 2 * n] = pts[i].pos + 8 * side;
        }

        const nxt = new Array(2 * n).fill(-1);

        for (let i = 0; i < 2 * n; i++) {
            const p = pts[i % n];
            const off = getOffset(p.x, p.y, d);
            if (off < 0) continue;

            const target = pos3[i] + off;
            const hi = Math.min(i + n, 3 * n);
            const j = lowerBound(pos3, i + 1, hi, target);
            if (j < hi) nxt[i] = j;
        }

        for (let start = 0; start < n; start++) {
            let cur = start;
            let cnt = 1;

            while (cnt < k) {
                cur = nxt[cur];
                if (cur === -1 || cur >= start + n) break;
                cnt++;
            }

            if (cnt >= k) {
                const a = pts[start];
                const b = pts[cur % n];
                if (Math.abs(a.x - b.x) + Math.abs(a.y - b.y) >= d) {
                    return true;
                }
            }
        }

        return false;
    }

    let lo = 0, hi = 2 * side;
    while (lo < hi) {
        const mid = Math.floor((lo + hi + 1) / 2);
        if (can(mid)) lo = mid;
        else hi = mid - 1;
    }

    return lo;
};
```

### Python3

```python
from bisect import bisect_left
from typing import List

class Solution:
    def maxDistance(self, side: int, points: List[List[int]], k: int) -> int:
        pts = []
        for x, y in points:
            if x == 0:
                pos = y
            elif y == side:
                pos = side + x
            elif x == side:
                pos = 3 * side - y
            else:
                pos = 4 * side - x
            pts.append((pos, x, y))

        pts.sort()
        n = len(pts)

        def get_offset(x: int, y: int, d: int) -> int:
            if x == 0:  # left
                if d <= 2 * side - y:
                    return d
                if d <= side + y:
                    return 2 * side + d - 2 * y
                return -1
            elif y == side:  # top
                if d <= 2 * side - x:
                    return d
                if d <= side + x:
                    return 2 * side + d - 2 * x
                return -1
            elif x == side:  # right
                if d <= side + y:
                    return d
                if d <= 2 * side - y:
                    return d + 2 * y
                return -1
            else:  # bottom
                if d <= side + x:
                    return d
                if d <= 2 * side - x:
                    return d + 2 * x
                return -1

        def feasible(d: int) -> bool:
            pos = [p[0] for p in pts]
            pos3 = pos + [p + 4 * side for p in pos] + [p + 8 * side for p in pos]

            nxt = [-1] * (2 * n)

            for i in range(2 * n):
                _, x, y = pts[i % n]
                off = get_offset(x, y, d)
                if off < 0:
                    continue

                target = pos3[i] + off
                hi = min(i + n, 3 * n)
                j = bisect_left(pos3, target, i + 1, hi)
                if j < hi:
                    nxt[i] = j

            for start in range(n):
                cur = start
                cnt = 1

                while cnt < k:
                    cur = nxt[cur]
                    if cur == -1 or cur >= start + n:
                        break
                    cnt += 1

                if cnt >= k:
                    x1, y1 = pts[start][1], pts[start][2]
                    x2, y2 = pts[cur % n][1], pts[cur % n][2]
                    if abs(x1 - x2) + abs(y1 - y2) >= d:
                        return True

            return False

        lo, hi = 0, 2 * side
        while lo < hi:
            mid = (lo + hi + 1) // 2
            if feasible(mid):
                lo = mid
            else:
                hi = mid - 1

        return lo
```

### Go

```go
import (
 "sort"
)

type Point struct {
 pos int64
 x   int64
 y   int64
}

func maxDistance(side int, points [][]int, k int) int {
 s := int64(side)
 n := len(points)
 pts := make([]Point, n)

 for i, v := range points {
  x := int64(v[0])
  y := int64(v[1])

  var pos int64
  if x == 0 {
   pos = y
  } else if y == s {
   pos = s + x
  } else if x == s {
   pos = 3*s - y
  } else {
   pos = 4*s - x
  }
  pts[i] = Point{pos: pos, x: x, y: y}
 }

 sort.Slice(pts, func(i, j int) bool {
  return pts[i].pos < pts[j].pos
 })

 getOffset := func(x, y, d int64) int64 {
  if x == 0 {
   if d <= 2*s-y {
    return d
   }
   if d <= s+y {
    return 2*s + d - 2*y
   }
   return -1
  } else if y == s {
   if d <= 2*s-x {
    return d
   }
   if d <= s+x {
    return 2*s + d - 2*x
   }
   return -1
  } else if x == s {
   if d <= s+y {
    return d
   }
   if d <= 2*s-y {
    return d + 2*y
   }
   return -1
  } else {
   if d <= s+x {
    return d
   }
   if d <= 2*s-x {
    return d + 2*x
   }
   return -1
  }
 }

 lowerBound := func(arr []int64, l, r int, target int64) int {
  for l < r {
   m := (l + r) >> 1
   if arr[m] < target {
    l = m + 1
   } else {
    r = m
   }
  }
  return l
 }

 can := func(d int64) bool {
  pos3 := make([]int64, 3*n)
  for i := 0; i < n; i++ {
   pos3[i] = pts[i].pos
   pos3[i+n] = pts[i].pos + 4*s
   pos3[i+2*n] = pts[i].pos + 8*s
  }

  nxt := make([]int, 2*n)
  for i := range nxt {
   nxt[i] = -1
  }

  for i := 0; i < 2*n; i++ {
   p := pts[i%n]
   off := getOffset(p.x, p.y, d)
   if off < 0 {
    continue
   }

   target := pos3[i] + off
   hi := i + n
   if hi > 3*n {
    hi = 3*n
   }
   j := lowerBound(pos3, i+1, hi, target)
   if j < hi {
    nxt[i] = j
   }
  }

  for start := 0; start < n; start++ {
   cur := start
   cnt := 1

   for cnt < k {
    cur = nxt[cur]
    if cur == -1 || cur >= start+n {
     break
    }
    cnt++
   }

   if cnt >= k {
    a := pts[start]
    b := pts[cur%n]
    dist := a.x - b.x
    if dist < 0 {
     dist = -dist
    }
    dy := a.y - b.y
    if dy < 0 {
     dy = -dy
    }
    if dist+dy >= d {
     return true
    }
   }
  }

  return false
 }

 lo, hi := int64(0), 2*s
 for lo < hi {
  mid := (lo + hi + 1) >> 1
  if can(mid) {
   lo = mid
  } else {
   hi = mid - 1
  }
 }

 return int(lo)
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

### 1. Convert every point into a perimeter position

I map each boundary point to one number.

This gives me one straight clockwise order around the square.

### 2. Sort the points

After the conversion, I sort all points by their perimeter position.

Now moving clockwise becomes simple array traversal.

### 3. Build a repeated perimeter array

I copy the sorted positions multiple times.

This helps me handle wrap-around when I move past the end of the square boundary.

### 4. Check one candidate distance

For a candidate distance `D`:

* I try every point as a starting point.
* I greedily pick the first next point that is valid.
* I repeat this until I either pick `k` points or fail.

### 5. Why greedy works

I always choose the earliest valid next point because:

* choosing earlier leaves more room for future points,
* and I want to know whether a valid chain exists.

That makes the greedy test correct for feasibility.

### 6. Binary search the answer

The answer has a monotonic property.

So I binary search the maximum `D` that still works.

## Examples

### Example 1

Input:

```text
side = 2
points = [[0,2],[2,0],[2,2],[0,0]]
k = 4
```

Output:

```text
2
```

Explanation:

I select all four corner points.
The minimum Manhattan distance among them is `2`.

### Example 2

Input:

```text
side = 2
points = [[0,0],[1,2],[2,0],[2,2],[2,1]]
k = 4
```

Output:

```text
1
```

Explanation:

One valid selection is:
`(0,0), (2,0), (2,2), (2,1)`

The smallest pair distance in this set is `1`.

### Example 3

Input:

```text
side = 2
points = [[0,0],[0,1],[0,2],[1,2],[2,0],[2,2],[2,1]]
k = 5
```

Output:

```text
1
```

Explanation:

One valid selection is:
`(0,0), (0,1), (0,2), (1,2), (2,2)`

The best minimum distance I can guarantee is `1`.

## How to use / Run locally

### C++

```bash
g++ -std=c++17 -O2 -o main main.cpp
./main
```

### Java

```bash
javac Main.java
java Main
```

### JavaScript

```bash
node main.js
```

### Python3

```bash
python3 main.py
```

### Go

```bash
go run main.go
```

## Notes & Optimizations

* I use `long long` / `int64` where needed because `side` can be very large.
* I avoid brute force because the number of possible subsets is too big.
* Binary search reduces the problem to checking only a few candidate answers.
* Since `k <= 25`, the greedy feasibility check is practical.
* The perimeter conversion is the main trick that turns a 2D boundary problem into a 1D ordering problem.

## Author

* [Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
