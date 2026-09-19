# Circle and Rectangle Overlapping - LeetCode 1401 Solution

## Table of Contents
- [Problem Summary](#problem-summary)
- [Constraints](#constraints)
- [Intuition](#intuition)
- [Approach](#approach)
- [Data Structures Used](#data-structures-used)
- [Operations & Behavior Summary](#operations--behavior-summary)
- [Complexity](#complexity)
- [Multi-language Solutions](#multi-language-solutions)
  - [C++](#c)
  - [Java](#java)
  - [JavaScript](#javascript)
  - [TypeScript](#typescript)
  - [Python3](#python3)
  - [Go](#go)
- [Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)](#step-by-step-detailed-explanation-c-java-javascript-typescript-python3-go)
- [Examples](#examples)
- [How to Use / Run Locally](#how-to-use--run-locally)
- [Notes & Optimizations](#notes--optimizations)
- [Author](#author)

## Problem Summary

You are given a circle defined by its radius and center coordinates (xCenter, yCenter). You are also given an axis-aligned rectangle defined by the bottom-left corner (x1, y1) and the top-right corner (x2, y2).

The task is to check whether the circle and the rectangle overlap. In simple terms, return true if there is any point that belongs to both the circle and the rectangle at the same time. Otherwise return false.

This is a classic geometry problem that appears in competitive programming and DSA interviews. It tests how well you can find the shortest distance between a point and a rectangle without checking every possible location.

## Constraints

- 1 <= radius <= 2000
- -10^4 <= xCenter, yCenter <= 10^4
- -10^4 <= x1 < x2 <= 10^4
- -10^4 <= y1 < y2 <= 10^4

These limits mean the numbers stay small enough that integer arithmetic is safe when handled carefully.

## Intuition

When I first saw the problem, I realized I did not need to check every point inside the shapes. The only thing that matters is the closest point on the rectangle to the center of the circle.

If that closest point lies inside or on the boundary of the circle, then the two shapes must overlap. Because the rectangle is axis-aligned, finding that closest point is straightforward: I just clamp the circle’s center coordinates to the rectangle’s boundaries.

This single observation turns a geometry problem into a few simple min and max operations.

## Approach

I calculate the nearest point on the rectangle to the circle’s center by clamping:

- The x-coordinate of the center is forced to stay between x1 and x2.
- The y-coordinate of the center is forced to stay between y1 and y2.

The resulting point is the closest location on the rectangle (it could be a corner, an edge, or even the center itself if the center is inside the rectangle).

I then measure the squared distance from the true center to this clamped point. If the squared distance is less than or equal to radius squared, the circle reaches that point and the shapes overlap.

Using squared distance avoids an expensive square-root call and keeps everything in integers.

## Data Structures Used

No complex data structures are needed. The solution uses only a handful of integer variables to store the clamped coordinates and the differences. This keeps both time and memory usage constant.

## Operations & Behavior Summary

1. Clamp the circle center’s x value into the range [x1, x2].
2. Clamp the circle center’s y value into the range [y1, y2].
3. Compute the horizontal and vertical differences between the original center and the clamped point.
4. Square both differences and add them together.
5. Compare the result with the square of the radius.
6. Return true if the distance is small enough, otherwise false.

These steps cover every possible case: center inside the rectangle, center outside but circle touching an edge, or circle only touching a corner.

## Complexity

| Type              | Complexity | Explanation |
|-------------------|------------|-------------|
| Time Complexity   | O(1)       | Only a fixed number of arithmetic operations and comparisons are performed. Nothing depends on the magnitude of the coordinates beyond constant-time math. |
| Space Complexity  | O(1)       | A few local integer variables are used. No arrays, maps, or extra memory are allocated. |

## Multi-language Solutions

### C++
```cpp
class Solution {
public:
    bool checkOverlap(int radius, int xCenter, int yCenter, int x1, int y1, int x2, int y2) {
        int closestX = max(x1, min(xCenter, x2));
        int closestY = max(y1, min(yCenter, y2));
        int dx = closestX - xCenter;
        int dy = closestY - yCenter;
        return (long long)dx * dx + (long long)dy * dy <= (long long)radius * radius;
    }
};
```

### Java
```java
class Solution {
    public boolean checkOverlap(int radius, int xCenter, int yCenter, int x1, int y1, int x2, int y2) {
        int closestX = Math.max(x1, Math.min(xCenter, x2));
        int closestY = Math.max(y1, Math.min(yCenter, y2));
        int dx = closestX - xCenter;
        int dy = closestY - yCenter;
        return (long)dx * dx + (long)dy * dy <= (long)radius * radius;
    }
}
```

### JavaScript
```javascript
/**
 * @param {number} radius
 * @param {number} xCenter
 * @param {number} yCenter
 * @param {number} x1
 * @param {number} y1
 * @param {number} x2
 * @param {number} y2
 * @return {boolean}
 */
var checkOverlap = function(radius, xCenter, yCenter, x1, y1, x2, y2) {
    let closestX = Math.max(x1, Math.min(xCenter, x2));
    let closestY = Math.max(y1, Math.min(yCenter, y2));
    let dx = closestX - xCenter;
    let dy = closestY - yCenter;
    return dx * dx + dy * dy <= radius * radius;
};
```

### TypeScript
```typescript
function checkOverlap(radius: number, xCenter: number, yCenter: number, x1: number, y1: number, x2: number, y2: number): boolean {
    let closestX = Math.max(x1, Math.min(xCenter, x2));
    let closestY = Math.max(y1, Math.min(yCenter, y2));
    let dx = closestX - xCenter;
    let dy = closestY - yCenter;
    return dx * dx + dy * dy <= radius * radius;
};
```

### Python3
```python
class Solution:
    def checkOverlap(self, radius: int, xCenter: int, yCenter: int, x1: int, y1: int, x2: int, y2: int) -> bool:
        closestX = max(x1, min(xCenter, x2))
        closestY = max(y1, min(yCenter, y2))
        dx = closestX - xCenter
        dy = closestY - yCenter
        return dx * dx + dy * dy <= radius * radius
```

### Go
```go
func checkOverlap(radius int, xCenter int, yCenter int, x1 int, y1 int, x2 int, y2 int) bool {
    closestX := max(x1, min(xCenter, x2))
    closestY := max(y1, min(yCenter, y2))
    dx := closestX - xCenter
    dy := closestY - yCenter
    return dx*dx+dy*dy <= radius*radius
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

The core logic is identical across all six languages, so the reasoning stays the same.

First I find the closest x-coordinate on the rectangle. I take the minimum of the center’s x and the right edge x2, then the maximum of that result and the left edge x1. This forces the value into the closed interval [x1, x2]. The same process is applied independently to the y-coordinate.

Once I have the nearest point, I calculate how far it sits from the true center in both directions. Squaring those two differences and adding them gives the squared Euclidean distance.

I compare that value with radius multiplied by itself. If the distance is smaller than or equal to the radius, the circle covers the closest point and therefore overlaps the rectangle.

When the center lies inside the rectangle, both differences become zero. The distance is zero, which is always less than or equal to any positive radius, so the function correctly returns true.

If the center is outside, the clamped point lands on the nearest edge or corner. The distance test then decides whether the circle is large enough to reach that point.

In languages that use fixed-size integers (C++ and Java), I cast the multiplications to a wider type so intermediate results cannot overflow. In JavaScript, TypeScript, Python, and Go the default numeric types already handle the required range safely.

Edge cases such as the circle touching only a corner, the circle completely containing the rectangle, or the rectangle completely containing the circle are all handled by the same clamping logic without any special branches.

## Examples

**Example 1**  
Input: radius = 1, xCenter = 0, yCenter = 0, x1 = 1, y1 = -1, x2 = 3, y2 = 1  
Output: true  

The closest point on the rectangle is (1, 0). The distance from (0, 0) to (1, 0) is exactly 1, which equals the radius, so the shapes touch.

**Example 2**  
Input: radius = 1, xCenter = 1, yCenter = 1, x1 = 1, y1 = -3, x2 = 2, y2 = -1  
Output: false  

The closest point is (1, -1). The distance from (1, 1) to (1, -1) is 2, which is greater than the radius, so there is no overlap.

**Example 3**  
Input: radius = 1, xCenter = 0, yCenter = 0, x1 = -1, y1 = 0, x2 = 0, y2 = 1  
Output: true  

The closest point is (0, 0) itself because the center lies on the boundary of the rectangle. Distance is zero, so the answer is true.

## How to Use / Run Locally

**C++**  
1. Save the code in a file named `main.cpp`.  
2. Compile with `g++ -std=c++17 main.cpp -o main`.  
3. Run with `./main`.

**Java**  
1. Save the code in a file named `Solution.java`.  
2. Compile with `javac Solution.java`.  
3. Run with `java Solution`.

**JavaScript**  
1. Save the code in a file named `solution.js`.  
2. Run with `node solution.js`.

**TypeScript**  
1. Save the code in a file named `solution.ts`.  
2. Compile with `tsc solution.ts`.  
3. Run the generated JavaScript with `node solution.js`.

**Python3**  
1. Save the code in a file named `solution.py`.  
2. Run with `python3 solution.py`.

**Go**  
1. Save the code in a file named `main.go`.  
2. Run with `go run main.go`.

You can wrap the function in a simple main method that reads sample inputs and prints the result to verify the logic.

## Notes & Optimizations

The solution already runs in constant time and constant space, so further asymptotic improvement is not possible. The only practical concern is integer overflow when squaring large coordinate differences. Casting to a wider type (long long in C++, long in Java) solves that issue completely under the given constraints.

An alternative approach could iterate over the four corners and four edges of the rectangle and compute distances, but that is both longer and slower. The clamping method is shorter, clearer, and optimal.

The algorithm also correctly handles the case where the circle completely contains the rectangle or the rectangle completely contains the circle, because the closest-point distance becomes zero in both situations.

## Author

[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)