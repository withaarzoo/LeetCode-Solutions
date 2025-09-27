# 812. Largest Triangle Area

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

Given an array of points on the X-Y plane `points` where `points[i] = [xi, yi]`, return *the area of the largest triangle that can be formed by any three different points.* Answers within `10^-5` of the actual answer will be accepted.

In short: choose any 3 distinct points from the list and compute the triangle area. Report the maximum area found.

## Constraints

* `3 <= points.length <= 50`
* `-50 <= xi, yi <= 50`
* All the given points are unique.

## Intuition

I thought about how to compute the triangle area quickly for three points. Instead of computing side lengths and using Heron's formula (which requires square roots and is more work), I remembered the **shoelace / cross-product formula** for the area of a polygon/triangle.

For three points `A(x1,y1)`, `B(x2,y2)`, `C(x3,y3)` the (signed) doubled area is:

```
2 * area = x1*(y2 - y3) + x2*(y3 - y1) + x3*(y1 - y2)
```

So I thought: iterate over all combinations of three points, compute this expression, take the absolute value and divide by 2. Keep the maximum. Because `n <= 50`, checking all `O(n^3)` triples is fast enough.

## Approach

1. Iterate all triples `(i, j, k)` with `i < j < k` to avoid duplicates.
2. For each triple, compute the doubled signed area using the shoelace formula.
3. Take absolute value, divide by 2 → that's the triangle area.
4. Track and update the maximum area.
5. Return the maximum area after checking all triples.

This is a simple, robust brute-force approach that is optimal for the given constraints.

## Data Structures Used

* Input is an array (list) of point pairs `[[x, y], ...]`.
* I only use primitive variables for indices and area calculations. No extra data structures required.

## Operations & Behavior Summary

* Triple nested loops to enumerate every unique triple of points (`i < j < k`).
* For each triple, compute the shoelace expression (constant-time arithmetic) to find area.
* Keep the maximum area in a single variable.

## Complexity

* **Time Complexity:** `O(n^3)` where `n = points.length` because we examine all triples of points. With `n <= 50`, `O(n^3)` (≈125k triple checks in worst case) is acceptable.
* **Space Complexity:** `O(1)` additional space — we only use counters and a variable to store the best area.

---

## Multi-language Solutions

### C++

```c++
class Solution {
public:
    double largestTriangleArea(vector<vector<int>>& points) {
        int n = points.size();
        double maxArea = 0.0;
        for (int i = 0; i < n - 2; ++i) {
            for (int j = i + 1; j < n - 1; ++j) {
                for (int k = j + 1; k < n; ++k) {
                    int x1 = points[i][0], y1 = points[i][1];
                    int x2 = points[j][0], y2 = points[j][1];
                    int x3 = points[k][0], y3 = points[k][1];
                    double doubled = fabs((double)x1*(y2 - y3)
                                        + (double)x2*(y3 - y1)
                                        + (double)x3*(y1 - y2));
                    double area = doubled * 0.5;
                    if (area > maxArea) maxArea = area;
                }
            }
        }
        return maxArea;
    }
};
```

### Java

```java
class Solution {
    public double largestTriangleArea(int[][] points) {
        int n = points.length;
        double maxArea = 0.0;
        for (int i = 0; i < n - 2; i++) {
            for (int j = i + 1; j < n - 1; j++) {
                for (int k = j + 1; k < n; k++) {
                    int x1 = points[i][0], y1 = points[i][1];
                    int x2 = points[j][0], y2 = points[j][1];
                    int x3 = points[k][0], y3 = points[k][1];
                    double doubled = Math.abs(
                        (double)x1 * (y2 - y3) +
                        (double)x2 * (y3 - y1) +
                        (double)x3 * (y1 - y2)
                    );
                    double area = doubled * 0.5;
                    if (area > maxArea) maxArea = area;
                }
            }
        }
        return maxArea;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[][]} points
 * @return {number}
 */
var largestTriangleArea = function(points) {
    const n = points.length;
    let maxArea = 0;
    for (let i = 0; i < n - 2; i++) {
        for (let j = i + 1; j < n - 1; j++) {
            for (let k = j + 1; k < n; k++) {
                const [x1, y1] = points[i];
                const [x2, y2] = points[j];
                const [x3, y3] = points[k];
                const doubled = Math.abs(
                    x1 * (y2 - y3) +
                    x2 * (y3 - y1) +
                    x3 * (y1 - y2)
                );
                const area = doubled * 0.5;
                if (area > maxArea) maxArea = area;
            }
        }
    }
    return maxArea;
};
```

### Python3

```python3
from typing import List

class Solution:
    def largestTriangleArea(self, points: List[List[int]]) -> float:
        n = len(points)
        max_area = 0.0
        for i in range(n - 2):
            for j in range(i + 1, n - 1):
                for k in range(j + 1, n):
                    x1, y1 = points[i]
                    x2, y2 = points[j]
                    x3, y3 = points[k]
                    doubled = abs(x1*(y2 - y3) + x2*(y3 - y1) + x3*(y1 - y2))
                    area = doubled * 0.5
                    if area > max_area:
                        max_area = area
        return max_area
```

### Go

```go
import "math"

func largestTriangleArea(points [][]int) float64 {
    n := len(points)
    maxArea := 0.0
    for i := 0; i < n-2; i++ {
        for j := i+1; j < n-1; j++ {
            for k := j+1; k < n; k++ {
                x1, y1 := points[i][0], points[i][1]
                x2, y2 := points[j][0], points[j][1]
                x3, y3 := points[k][0], points[k][1]
                doubled := math.Abs(float64(x1*(y2 - y3) + x2*(y3 - y1) + x3*(y1 - y2)))
                area := doubled * 0.5
                if area > maxArea {
                    maxArea = area
                }
            }
        }
    }
    return maxArea
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

I'll explain the algorithm in simple steps and then point to the important lines in each language implementation.

### Shared logic (conceptual - applies to all languages)

1. `n = points.length` — number of points.
2. Initialize `maxArea = 0.0` — this will store the best area seen.
3. Use three nested loops: first picks index `i` from `0` to `n-3`, second picks `j` from `i+1` to `n-2`, third picks `k` from `j+1` to `n-1`. This enumerates each unique triple exactly once.
4. For the triple `(i, j, k)`, read coordinates `(x1,y1)`, `(x2,y2)`, `(x3,y3)`.
5. Compute `doubled = abs(x1*(y2 - y3) + x2*(y3 - y1) + x3*(y1 - y2))`. This value is `2 * area` (signed). The absolute value gives the magnitude.
6. `area = doubled * 0.5` — actual triangle area.
7. If `area > maxArea`, update `maxArea`.
8. After the loops finish, return `maxArea`.

### C++ notes

* `fabs` is used to get the absolute value of a `double` expression.
* Cast intermediate multiplications to `double` to avoid implicit integer promotions during arithmetic with negatives (not strictly necessary for these bounds but safe).

Important lines (in the snippet above):

```c++
double doubled = fabs((double)x1*(y2 - y3) + (double)x2*(y3 - y1) + (double)x3*(y1 - y2));
double area = doubled * 0.5;
if (area > maxArea) maxArea = area;
```

### Java notes

* Use `Math.abs` for absolute value of `double`.
* Convert integer arithmetic to `double` during the cross-product expression to keep floating precision.

Key lines:

```java
double doubled = Math.abs((double)x1 * (y2 - y3) + (double)x2 * (y3 - y1) + (double)x3 * (y1 - y2));
double area = doubled * 0.5;
if (area > maxArea) maxArea = area;
```

### JavaScript notes

* JavaScript uses `Number` for numeric values; `Math.abs` computes absolute value.
* Array destructuring `const [x1, y1] = points[i];` makes reading coordinates readable.

Key lines:

```javascript
const doubled = Math.abs(x1 * (y2 - y3) + x2 * (y3 - y1) + x3 * (y1 - y2));
const area = doubled * 0.5;
if (area > maxArea) maxArea = area;
```

### Python3 notes

* Python integers have arbitrary precision, so no overflow concerns within the given constraints.
* `abs(...)` returns absolute value; use `max_area` variable to track best.

Key lines:

```python
doubled = abs(x1*(y2 - y3) + x2*(y3 - y1) + x3*(y1 - y2))
area = doubled * 0.5
if area > max_area:
    max_area = area
```

### Go notes

* Use `math.Abs` to get absolute value of `float64`.
* Convert integer arithmetic results to `float64` before calling `math.Abs`.

Key lines:

```go
doubled := math.Abs(float64(x1*(y2 - y3) + x2*(y3 - y1) + x3*(y1 - y2)))
area := doubled * 0.5
if area > maxArea {
    maxArea = area
}
```

---

## Examples

**Example 1**

```
Input: points = [[0,0],[0,1],[1,0],[0,2],[2,0]]
Output: 2.00000
Explanation: The largest triangle uses points (0,2),(0,0),(2,0) -> area = 2
```

**Example 2**

```
Input: points = [[1,0],[0,0],[0,1]]
Output: 0.5
```

## How to use / Run locally

Below are quick instructions for testing locally. For each language, I give an example wrapper to run the function as a standalone program.

### C++ (g++)

1. Create `main.cpp` and paste a `main` wrapper that builds a `points` vector and calls `largestTriangleArea` from `Solution`.
2. Compile: `g++ -std=c++17 -O2 main.cpp -o main`
3. Run: `./main`

**Example wrapper idea (for local testing):** create a `main` that reads points, calls function, prints result.

### Java (javac/java)

1. Put the `Solution` class into `Solution.java` and add a `public static void main` method that creates `int[][] points` and calls `new Solution().largestTriangleArea(points)`.
2. Compile: `javac Solution.java`
3. Run: `java Solution`

### JavaScript (Node.js)

1. Create `test.js` and paste the JS function along with a small test harness:

```javascript
// include the function then
const points = [[0,0],[0,1],[1,0],[0,2],[2,0]];
console.log(largestTriangleArea(points));
```

2. Run: `node test.js`

### Python3

1. Create `test.py` with the `Solution` class and at the bottom:

```python
points = [[0,0],[0,1],[1,0],[0,2],[2,0]]
print(Solution().largestTriangleArea(points))
```

2. Run: `python3 test.py`

### Go

1. Create `main.go` with a `package main`, `func main()` wrapper that prepares `points` and calls `largestTriangleArea` then prints the result.
2. Run: `go run main.go`

---

## Notes & Optimizations

* Because `n <= 50`, `O(n^3)` brute force is simple and fast enough. For much larger `n` we would need a different strategy.
* If `n` were large (e.g., thousands), we could first compute the **convex hull** of the points (O(n log n)), because the largest-area triangle must have all three vertices on the convex hull. Let `h` be hull size. Then search triangles only among hull points. There are algorithms (rotating calipers variants) that can find the largest-area triangle on a convex polygon faster than `O(h^3)` in practice.
* Integer overflow: given constraints `xi, yi ∈ [-50, 50]`, intermediate products fit comfortably into 32-bit signed ints. Still, we cast to `double` when computing absolute area to avoid integer division surprises and to match the return type.
* Numerical precision: final result is returned as `double`/`float` and is correct to `1e-5` as required.

## Author

[Md. Aarzoo Islam](https://bento.me/withaarzoo)
