# Problem Title

**3047. Find the Largest Area of Square Inside Two Rectangles**

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
* How to use / Run locally
* Notes & Optimizations
* Author

---

## Problem Summary

I am given multiple rectangles on a 2D plane.
Each rectangle is defined using its bottom-left and top-right coordinates.

My task is to find the **maximum possible area of a square** that can fit **inside the overlapping region of at least two rectangles**.

Important points:

* The square must be completely inside the intersection area.
* If no two rectangles overlap, the answer is `0`.
* Rectangles are aligned parallel to the x-axis and y-axis.

---

## Constraints

* Number of rectangles `n` is between 2 and 1000
* Each rectangle has:

  * bottomLeft[i] = [x1, y1]
  * topRight[i] = [x2, y2]
* 1 ≤ x1, y1, x2, y2 ≤ 10⁷
* x1 < x2 and y1 < y2

---

## Intuition

When I started thinking about this problem, I realized one simple fact.

A square can only exist inside the **intersection of rectangles**, not outside.

So instead of checking all rectangles together, I decided to:

* Take **two rectangles at a time**
* Find their **overlapping area**
* Try to fit the **largest possible square** inside that overlap

The overlapping region of two rectangles is also a rectangle.
The largest square inside a rectangle always has side length equal to the **minimum of its width and height**.

So the problem becomes very straightforward.

---

## Approach

Here is how I solved it step by step.

1. Loop through every pair of rectangles.
2. For each pair:

   * Compute the intersection rectangle using max and min coordinates.
3. If there is no overlap, skip that pair.
4. If they overlap:

   * Calculate overlap width and height.
   * The square side is the smaller of the two.
5. Square the side to get area.
6. Track the maximum area found.
7. Return the maximum area at the end.

This brute-force approach works efficiently because the limit is only 1000 rectangles.

---

## Data Structures Used

I did not use any extra data structures.

Only simple variables like:

* integers
* long long / long
* loops

So memory usage stays constant.

---

## Operations & Behavior Summary

* Pairwise rectangle comparison
* Coordinate intersection calculation
* Width and height computation
* Square area calculation
* Maximum value tracking

No recursion.
No sorting.
No additional space.

---

## Complexity

**Time Complexity:**
O(n²)
Because I compare every pair of rectangles once.

**Space Complexity:**
O(1)
Only constant extra space is used.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    long long largestSquareArea(vector<vector<int>>& bottomLeft, vector<vector<int>>& topRight) {
        int n = bottomLeft.size();
        long long ans = 0;

        for (int i = 0; i < n; i++) {
            for (int j = i + 1; j < n; j++) {

                int left   = max(bottomLeft[i][0], bottomLeft[j][0]);
                int right  = min(topRight[i][0], topRight[j][0]);
                int bottom = max(bottomLeft[i][1], bottomLeft[j][1]);
                int top    = min(topRight[i][1], topRight[j][1]);

                if (right > left && top > bottom) {
                    long long side = min(right - left, top - bottom);
                    ans = max(ans, side * side);
                }
            }
        }
        return ans;
    }
};
```

---

### Java

```java
class Solution {
    public long largestSquareArea(int[][] bottomLeft, int[][] topRight) {
        int n = bottomLeft.length;
        long ans = 0;

        for (int i = 0; i < n; i++) {
            for (int j = i + 1; j < n; j++) {

                int left   = Math.max(bottomLeft[i][0], bottomLeft[j][0]);
                int right  = Math.min(topRight[i][0], topRight[j][0]);
                int bottom = Math.max(bottomLeft[i][1], bottomLeft[j][1]);
                int top    = Math.min(topRight[i][1], topRight[j][1]);

                if (right > left && top > bottom) {
                    long side = Math.min(right - left, top - bottom);
                    ans = Math.max(ans, side * side);
                }
            }
        }
        return ans;
    }
}
```

---

### JavaScript

```javascript
var largestSquareArea = function(bottomLeft, topRight) {
    let ans = 0;
    const n = bottomLeft.length;

    for (let i = 0; i < n; i++) {
        for (let j = i + 1; j < n; j++) {

            const left   = Math.max(bottomLeft[i][0], bottomLeft[j][0]);
            const right  = Math.min(topRight[i][0], topRight[j][0]);
            const bottom = Math.max(bottomLeft[i][1], bottomLeft[j][1]);
            const top    = Math.min(topRight[i][1], topRight[j][1]);

            if (right > left && top > bottom) {
                const side = Math.min(right - left, top - bottom);
                ans = Math.max(ans, side * side);
            }
        }
    }
    return ans;
};
```

---

### Python3

```python
class Solution:
    def largestSquareArea(self, bottomLeft, topRight):
        n = len(bottomLeft)
        ans = 0

        for i in range(n):
            for j in range(i + 1, n):

                left = max(bottomLeft[i][0], bottomLeft[j][0])
                right = min(topRight[i][0], topRight[j][0])
                bottom = max(bottomLeft[i][1], bottomLeft[j][1])
                top = min(topRight[i][1], topRight[j][1])

                if right > left and top > bottom:
                    side = min(right - left, top - bottom)
                    ans = max(ans, side * side)

        return ans
```

---

### Go

```go
func largestSquareArea(bottomLeft [][]int, topRight [][]int) int64 {
    n := len(bottomLeft)
    var ans int64 = 0

    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {

            left := max(bottomLeft[i][0], bottomLeft[j][0])
            right := min(topRight[i][0], topRight[j][0])
            bottom := max(bottomLeft[i][1], bottomLeft[j][1])
            top := min(topRight[i][1], topRight[j][1])

            if right > left && top > bottom {
                side := int64(min(right-left, top-bottom))
                if side*side > ans {
                    ans = side * side
                }
            }
        }
    }
    return ans
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
```

---

## Step-by-step Detailed Explanation

* I compare rectangles pair by pair.
* For each pair, I calculate intersection coordinates.
* If width or height is zero or negative, there is no overlap.
* Otherwise, I find the largest square that can fit.
* I update the maximum square area.
* Finally, I return the best possible area.

---

## Examples

Input
`bottomLeft = [[1,1],[2,2],[3,1]]`
`topRight = [[3,3],[4,4],[6,6]]`

Output
`1`

Explanation
The largest square that fits in any overlapping region has side length 1.

---

## How to use / Run locally

1. Copy the code for your preferred language.
2. Paste it into your LeetCode editor or local environment.
3. Run with sample inputs to verify.
4. Submit directly on LeetCode.

---

## Notes & Optimizations

* Brute force is acceptable due to small constraints.
* No need for advanced geometry or sweepline.
* Clean and readable logic is preferred here.
* Works efficiently within given limits.

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
