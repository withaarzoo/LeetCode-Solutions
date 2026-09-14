---

# LeetCode 836 Rectangle Overlap Solution – Check if Two Axis-Aligned Rectangles Overlap

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

The Rectangle Overlap problem asks you to decide whether two axis-aligned rectangles share any positive area.  

Each rectangle is given as a list of four integers `[x1, y1, x2, y2]`.  
- `(x1, y1)` is the bottom-left corner  
- `(x2, y2)` is the top-right corner  

The sides are parallel to the axes. Two rectangles overlap only when the area of their intersection is strictly greater than zero. If they just touch at a corner or along an edge, that does not count as overlap.  

You are given two such rectangles `rec1` and `rec2` and must return `true` if they overlap, otherwise `false`.

## Constraints

- Both input arrays always contain exactly four integers  
- Every coordinate satisfies `-10^9 <= value <= 10^9`  
- Each rectangle has a non-zero area (so `x1 < x2` and `y1 < y2`)

## Intuition

When I first looked at the problem I realized that two rectangles can only overlap if their projections on both the x-axis and the y-axis also overlap.  

If one rectangle sits completely to the left of the other, or completely above it, their intersection area is zero. Checking those four “miss” cases is enough. If none of the four miss conditions is true, the rectangles must share positive area.

## Approach

I simply test the four ways the rectangles can completely miss each other:

1. The right edge of the first rectangle is at or left of the left edge of the second.  
2. The left edge of the first rectangle is at or right of the right edge of the second.  
3. The top edge of the first rectangle is at or below the bottom edge of the second.  
4. The bottom edge of the first rectangle is at or above the top edge of the second.  

If any of those four statements is true, the rectangles do not overlap. Otherwise they do.  

This check works for every valid pair of axis-aligned rectangles and needs only a handful of comparisons.

## Data Structures Used

No extra data structures are required.  
The solution works directly on the two input arrays that already hold the four corner coordinates of each rectangle. Using only the given arrays keeps both time and memory constant.

## Operations & Behavior Summary

- Read the four coordinates of the first rectangle.  
- Read the four coordinates of the second rectangle.  
- Test whether the first rectangle ends before the second starts on the x-axis.  
- Test whether the first rectangle starts after the second ends on the x-axis.  
- Test the same two conditions on the y-axis.  
- If any of the four tests is true, return false.  
- Otherwise return true.

## Complexity

| Complexity       | Value | Explanation                                      |
|------------------|-------|--------------------------------------------------|
| Time Complexity  | O(1)  | Only a fixed number of comparisons are performed |
| Space Complexity | O(1)  | No additional memory is allocated                |

## Multi-language Solutions

### C++
```cpp
class Solution {
public:
    bool isRectangleOverlap(vector<int>& rec1, vector<int>& rec2) {
        // return true only when the rectangles share positive area on both axes
        // the four conditions below cover every way they can miss each other
        return !(rec1[2] <= rec2[0] ||   // rec1 completely left of rec2
                 rec1[0] >= rec2[2] ||   // rec1 completely right of rec2
                 rec1[3] <= rec2[1] ||   // rec1 completely below rec2
                 rec1[1] >= rec2[3]);    // rec1 completely above rec2
    }
};
```

### Java
```java
class Solution {
    public boolean isRectangleOverlap(int[] rec1, int[] rec2) {
        // return true only when the rectangles share positive area on both axes
        // the four conditions below cover every way they can miss each other
        return !(rec1[2] <= rec2[0] ||   // rec1 completely left of rec2
                 rec1[0] >= rec2[2] ||   // rec1 completely right of rec2
                 rec1[3] <= rec2[1] ||   // rec1 completely below rec2
                 rec1[1] >= rec2[3]);    // rec1 completely above rec2
    }
}
```

### JavaScript
```javascript
/**
 * @param {number[]} rec1
 * @param {number[]} rec2
 * @return {boolean}
 */
var isRectangleOverlap = function(rec1, rec2) {
    // return true only when the rectangles share positive area on both axes
    // the four conditions below cover every way they can miss each other
    return !(rec1[2] <= rec2[0] ||   // rec1 completely left of rec2
             rec1[0] >= rec2[2] ||   // rec1 completely right of rec2
             rec1[3] <= rec2[1] ||   // rec1 completely below rec2
             rec1[1] >= rec2[3]);    // rec1 completely above rec2
};
```

### TypeScript
```typescript
function isRectangleOverlap(rec1: number[], rec2: number[]): boolean {
    // return true only when the rectangles share positive area on both axes
    // the four conditions below cover every way they can miss each other
    return !(rec1[2] <= rec2[0] ||   // rec1 completely left of rec2
             rec1[0] >= rec2[2] ||   // rec1 completely right of rec2
             rec1[3] <= rec2[1] ||   // rec1 completely below rec2
             rec1[1] >= rec2[3]);    // rec1 completely above rec2
};
```

### Python3
```python
class Solution:
    def isRectangleOverlap(self, rec1: List[int], rec2: List[int]) -> bool:
        # return true only when the rectangles share positive area on both axes
        # the four conditions below cover every way they can miss each other
        return not (rec1[2] <= rec2[0] or   # rec1 completely left of rec2
                    rec1[0] >= rec2[2] or   # rec1 completely right of rec2
                    rec1[3] <= rec2[1] or   # rec1 completely below rec2
                    rec1[1] >= rec2[3])     # rec1 completely above rec2
```

### Go
```go
func isRectangleOverlap(rec1 []int, rec2 []int) bool {
    // return true only when the rectangles share positive area on both axes
    // the four conditions below cover every way they can miss each other
    return !(rec1[2] <= rec2[0] || // rec1 completely left of rec2
             rec1[0] >= rec2[2] || // rec1 completely right of rec2
             rec1[3] <= rec2[1] || // rec1 completely below rec2
             rec1[1] >= rec2[3])   // rec1 completely above rec2
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

The logic is identical in every language, so the reasoning stays the same.

I grab the four numbers that describe each rectangle.  
Index 0 and 1 are the bottom-left corner; index 2 and 3 are the top-right corner.  

On the x-axis the first rectangle occupies the interval from `rec1[0]` to `rec1[2]`.  
The second occupies the interval from `rec2[0]` to `rec2[2]`.  
These two intervals share a positive length only when the right end of one is strictly greater than the left end of the other.  

I write the opposite test: the intervals do not overlap when one ends at or before the other starts. The same idea is applied to the y-coordinates.  

That produces the four Boolean expressions:

- `rec1[2] <= rec2[0]`  
- `rec1[0] >= rec2[2]`  
- `rec1[3] <= rec2[1]`  
- `rec1[1] >= rec2[3]`  

If any of them is true the shared area is zero, so I return false.  
Otherwise I return true.  

Using non-strict inequalities (`<=` and `>=`) correctly treats touching edges as non-overlapping, which matches the problem statement.  

Because every access is to a constant index, the whole function runs in constant time and uses constant extra memory in all six languages.

## Examples

**Example 1**  
Input: `rec1 = [0,0,2,2]`, `rec2 = [1,1,3,3]`  
Output: `true`  

Both rectangles cover the region from x = 1 to x = 2 and from y = 1 to y = 2.  
The intersection area is 1, so they overlap.

**Example 2**  
Input: `rec1 = [0,0,1,1]`, `rec2 = [1,0,2,1]`  
Output: `false`  

The first rectangle ends at x = 1 and the second starts at x = 1.  
They share an edge but no positive area.

**Example 3**  
Input: `rec1 = [0,0,1,1]`, `rec2 = [2,2,3,3]`  
Output: `false`  

The rectangles are completely separated on both axes, so the intersection is empty.

## How to Use / Run Locally

**C++**  
1. Save the code in a file named `solution.cpp`.  
2. Compile with `g++ -std=c++17 solution.cpp -o solution`.  
3. Run with `./solution`.  

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
1. Save the code in a file named `solution.go`.  
2. Run with `go run solution.go`.  

Add a small `main` function in each language if you want to test the examples yourself.

## Notes & Optimizations

The solution already runs in constant time and constant space, so further optimization is unnecessary for the given constraints.  

One common alternative is to compute the actual intersection rectangle and check whether its width and height are both positive. That approach uses a few more arithmetic operations and is slightly longer, but produces the same result.  

Edge cases to keep in mind:  
- Rectangles that touch only at a single point  
- Rectangles that share an entire side  
- Negative coordinates (the comparisons still work correctly)  

All of these are handled correctly by the four simple inequality checks.

## Author
[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)