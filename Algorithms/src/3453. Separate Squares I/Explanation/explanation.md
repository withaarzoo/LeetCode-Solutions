# Separate Squares I (LeetCode 3453)

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
* How to use / Run locally
* Notes & Optimizations
* Author

---

## Problem Summary

I am given multiple squares on a 2D plane.
Each square is defined by:

* bottom-left y-coordinate `yi`
* side length `li`

All squares are parallel to the x-axis.

My task is to find the **minimum y-coordinate** of a horizontal line such that:

* total area **above** the line
* equals total area **below** the line

Important points:

* Squares can overlap
* Overlapping area is counted multiple times
* Answer should be accurate within `1e-5`

---

## Constraints

* 1 ≤ number of squares ≤ 5 × 10⁴
* 0 ≤ xi, yi ≤ 10⁹
* 1 ≤ li ≤ 10⁹
* Total area of all squares ≤ 10¹²

---

## Intuition

When I read the problem, I realized one key thing.

If I move a horizontal line **upward**, then:

* area below the line **increases**
* area above the line **decreases**

This change is smooth and always in one direction.

So instead of checking every possible y-value,
I can **binary search the answer**.

My goal becomes:
Find a y where
area below = total area / 2

That makes this problem perfect for **binary search on answer**.

---

## Approach

1. First, I calculate the **total area** of all squares.
2. I decide the binary search range:

   * lowest possible y = minimum bottom of all squares
   * highest possible y = maximum top of all squares
3. For a guessed height `mid`, I calculate how much area lies **below** it.
4. For each square:

   * If the line is below the square → add 0
   * If the line is above the square → add full area
   * If the line cuts the square → add partial area
5. If area below is less than half, I move the line up.
6. Otherwise, I move the line down.
7. I repeat this until precision is good enough.
8. Finally, I return the smallest valid y.

---

## Data Structures Used

No complex data structures are required.

* Only basic variables
* Simple loops
* Constant extra memory

---

## Operations & Behavior Summary

* Binary search on y-coordinate
* Area calculation using geometry
* Floating-point precision handling
* Monotonic behavior guarantees correctness

---

## Complexity

**Time Complexity:**
O(n log R)

* n = number of squares
* R = range of y-coordinates

**Space Complexity:**
O(1)

* No extra data structures used

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    double separateSquares(vector<vector<int>>& squares) {
        double totalArea = 0;
        double low = 1e18, high = -1e18;

        for (auto &s : squares) {
            double y = s[1], l = s[2];
            totalArea += l * l;
            low = min(low, y);
            high = max(high, y + l);
        }

        for (int i = 0; i < 80; i++) {
            double mid = (low + high) / 2.0;
            double areaBelow = 0;

            for (auto &s : squares) {
                double y = s[1], l = s[2];
                if (mid <= y) continue;
                if (mid >= y + l) areaBelow += l * l;
                else areaBelow += l * (mid - y);
            }

            if (areaBelow * 2 < totalArea) low = mid;
            else high = mid;
        }
        return low;
    }
};
```

---

### Java

```java
class Solution {
    public double separateSquares(int[][] squares) {
        double totalArea = 0;
        double low = 1e18, high = -1e18;

        for (int[] s : squares) {
            double y = s[1], l = s[2];
            totalArea += l * l;
            low = Math.min(low, y);
            high = Math.max(high, y + l);
        }

        for (int i = 0; i < 80; i++) {
            double mid = (low + high) / 2.0;
            double areaBelow = 0;

            for (int[] s : squares) {
                double y = s[1], l = s[2];
                if (mid <= y) continue;
                if (mid >= y + l) areaBelow += l * l;
                else areaBelow += l * (mid - y);
            }

            if (areaBelow * 2 < totalArea) low = mid;
            else high = mid;
        }
        return low;
    }
}
```

---

### JavaScript

```javascript
var separateSquares = function(squares) {
    let totalArea = 0;
    let low = 1e18, high = -1e18;

    for (let [x, y, l] of squares) {
        totalArea += l * l;
        low = Math.min(low, y);
        high = Math.max(high, y + l);
    }

    for (let i = 0; i < 80; i++) {
        let mid = (low + high) / 2;
        let areaBelow = 0;

        for (let [x, y, l] of squares) {
            if (mid <= y) continue;
            if (mid >= y + l) areaBelow += l * l;
            else areaBelow += l * (mid - y);
        }

        if (areaBelow * 2 < totalArea) low = mid;
        else high = mid;
    }
    return low;
};
```

---

### Python3

```python
class Solution:
    def separateSquares(self, squares):
        totalArea = 0.0
        low, high = 1e18, -1e18

        for x, y, l in squares:
            totalArea += l * l
            low = min(low, y)
            high = max(high, y + l)

        for _ in range(80):
            mid = (low + high) / 2
            areaBelow = 0.0

            for x, y, l in squares:
                if mid <= y:
                    continue
                elif mid >= y + l:
                    areaBelow += l * l
                else:
                    areaBelow += l * (mid - y)

            if areaBelow * 2 < totalArea:
                low = mid
            else:
                high = mid

        return low
```

---

### Go

```go
func separateSquares(squares [][]int) float64 {
    totalArea := 0.0
    low, high := 1e18, -1e18

    for _, s := range squares {
        y := float64(s[1])
        l := float64(s[2])
        totalArea += l * l
        if y < low {
            low = y
        }
        if y+l > high {
            high = y + l
        }
    }

    for i := 0; i < 80; i++ {
        mid := (low + high) / 2
        areaBelow := 0.0

        for _, s := range squares {
            y := float64(s[1])
            l := float64(s[2])

            if mid <= y {
                continue
            } else if mid >= y+l {
                areaBelow += l * l
            } else {
                areaBelow += l * (mid - y)
            }
        }

        if areaBelow*2 < totalArea {
            low = mid
        } else {
            high = mid
        }
    }
    return low
}
```

---

## Step-by-step Detailed Explanation

1. I calculate the total area of all squares.
2. I define the binary search limits using square bottoms and tops.
3. I guess a height using mid.
4. I compute area below that height.
5. I compare it with half of total area.
6. Based on the result, I move the search space.
7. I repeat until precision is achieved.
8. I return the smallest valid y.

---

## Examples

**Input:**
`[[0,0,1],[2,2,1]]`

**Output:**
`1.00000`

---

**Input:**
`[[0,0,2],[1,1,1]]`

**Output:**
`1.16667`

---

## How to use / Run locally

1. Copy the solution code in your preferred language
2. Paste it into LeetCode editor or local compiler
3. Run with provided test cases
4. Compare output with expected values

---

## Notes & Optimizations

* Binary search guarantees accuracy
* Works for very large coordinates
* Overlapping areas handled naturally
* Precision controlled using iteration count

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
