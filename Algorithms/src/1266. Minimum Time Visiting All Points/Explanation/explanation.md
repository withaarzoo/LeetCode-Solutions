# Minimum Time Visiting All Points (LeetCode 1266)

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

I am given a list of points on a 2D plane.
Each point is represented as `[x, y]`.

I must visit **all points in the given order**.

In **1 second**, I can:

* Move 1 unit horizontally
* Move 1 unit vertically
* Move 1 unit diagonally (both x and y together)

My task is to calculate the **minimum time (in seconds)** required to visit all points.

---

## Constraints

* `1 <= points.length <= 100`
* `points[i].length == 2`
* `-1000 <= x, y <= 1000`
* Points must be visited **in order**

---

## Intuition

When I started thinking about this problem, I focused on **movement rules**.

The key observation I made was:

* A diagonal move changes both x and y in **one second**
* So diagonal moves are the **most efficient**

To go from one point to another, I should:

* Use diagonal moves as much as possible
* Then use straight moves for the remaining distance

This idea simplifies the entire problem.

---

## Approach

For every two consecutive points:

1. I calculate the distance in x direction
2. I calculate the distance in y direction
3. I take the **maximum** of those two distances
4. I add it to my total time

Why max?

* Diagonal moves cover both directions
* Remaining distance is handled by straight moves

I repeat this for all points.

---

## Data Structures Used

* No extra data structure is required
* I only use integer variables to store distances and time

---

## Operations & Behavior Summary

* Traverse points one by one
* Calculate absolute difference in x and y
* Add `max(dx, dy)` to total time
* Return total time after visiting all points

---

## Complexity

**Time Complexity:**
`O(n)`
Where `n` is the number of points.
I iterate through the list only once.

**Space Complexity:**
`O(1)`
No extra memory is used.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int minTimeToVisitAllPoints(vector<vector<int>>& points) {
        int totalTime = 0;

        for (int i = 1; i < points.size(); i++) {
            int dx = abs(points[i][0] - points[i - 1][0]);
            int dy = abs(points[i][1] - points[i - 1][1]);
            totalTime += max(dx, dy);
        }

        return totalTime;
    }
};
```

---

### Java

```java
class Solution {
    public int minTimeToVisitAllPoints(int[][] points) {
        int totalTime = 0;

        for (int i = 1; i < points.length; i++) {
            int dx = Math.abs(points[i][0] - points[i - 1][0]);
            int dy = Math.abs(points[i][1] - points[i - 1][1]);
            totalTime += Math.max(dx, dy);
        }

        return totalTime;
    }
}
```

---

### JavaScript

```javascript
var minTimeToVisitAllPoints = function(points) {
    let totalTime = 0;

    for (let i = 1; i < points.length; i++) {
        const dx = Math.abs(points[i][0] - points[i - 1][0]);
        const dy = Math.abs(points[i][1] - points[i - 1][1]);
        totalTime += Math.max(dx, dy);
    }

    return totalTime;
};
```

---

### Python3

```python
class Solution:
    def minTimeToVisitAllPoints(self, points):
        total_time = 0

        for i in range(1, len(points)):
            dx = abs(points[i][0] - points[i - 1][0])
            dy = abs(points[i][1] - points[i - 1][1])
            total_time += max(dx, dy)

        return total_time
```

---

### Go

```go
func minTimeToVisitAllPoints(points [][]int) int {
    totalTime := 0

    for i := 1; i < len(points); i++ {
        dx := abs(points[i][0] - points[i-1][0])
        dy := abs(points[i][1] - points[i-1][1])

        if dx > dy {
            totalTime += dx
        } else {
            totalTime += dy
        }
    }

    return totalTime
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
```

---

## Step-by-step Detailed Explanation (All Languages)

1. Start from the first point
2. Move to the next point in order
3. Calculate horizontal distance `dx`
4. Calculate vertical distance `dy`
5. Use diagonal movement to reduce both
6. Remaining distance is handled by straight moves
7. Total time for this move is `max(dx, dy)`
8. Add this time to the answer
9. Repeat until all points are visited

---

## Examples

**Input**

```bash
points = [[1,1],[3,4],[-1,0]]
```

**Output**

```bash
7
```

**Explanation**

* From (1,1) to (3,4) → 3 seconds
* From (3,4) to (-1,0) → 4 seconds
* Total = 7 seconds

---

## How to use / Run locally

1. Copy the code in your preferred language
2. Paste it into LeetCode editor or local file
3. Compile and run
4. Pass input as a list of points

---

## Notes & Optimizations

* No BFS or DFS required
* No dynamic programming needed
* Observation based solution
* Optimal and interview friendly
* Works efficiently within constraints

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
