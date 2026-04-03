# 3661. Maximum Walls Destroyed by Robots

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
* [Step-by-step Detailed Explanation](#step-by-step-detailed-explanation)
* [Examples](#examples)
* [How to use / Run locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

---

## Problem Summary

There is an infinite straight line with robots and walls placed at different positions.

Each robot has:

* A position
* A maximum shooting distance
* One bullet

A robot can shoot either left or right.

The bullet destroys every wall in its path.

However, if the bullet reaches another robot before reaching a wall, it immediately stops.

The goal is to find the maximum number of unique walls that can be destroyed.

---

## Constraints

* `1 <= robots.length == distance.length <= 10^5`
* `1 <= walls.length <= 10^5`
* `1 <= robots[i], walls[j] <= 10^9`
* `1 <= distance[i] <= 10^5`
* All robot positions are unique
* All wall positions are unique

---

## Intuition

I first thought about each robot independently.

For every robot, I can calculate:

* How many walls it can destroy if it shoots left
* How many walls it can destroy if it shoots right

But there is one important issue.

If two nearby robots shoot toward each other, some walls may be counted twice.

So I need Dynamic Programming to track the best answer while avoiding double-counting.

I also need Binary Search because I need to quickly count how many walls lie inside a shooting range.

---

## Approach

1. Store robots as `(position, distance)` pairs.
2. Sort robots by position.
3. Sort walls.
4. For every robot:

   * Compute the maximum valid left shooting range.
   * Compute the maximum valid right shooting range.
5. Use binary search to count how many walls exist inside any interval.
6. Use DP where:

   * `dp[i][0]` means robot `i` shoots left.
   * `dp[i][1]` means robot `i` shoots right.
7. Handle overlap carefully when:

   * Previous robot shoots right
   * Current robot shoots left

---

## Data Structures Used

* Array of pairs for storing `(robot position, distance)`
* Sorted array for walls
* DP table of size `n x 2`
* Binary Search using:

  * `lower_bound`
  * `upper_bound`

---

## Operations & Behavior Summary

| Operation       | Purpose                                 |
| --------------- | --------------------------------------- |
| Sort robots     | Helps identify previous and next robot  |
| Sort walls      | Enables binary search                   |
| Binary Search   | Counts walls inside a range efficiently |
| DP              | Stores best answer till robot `i`       |
| Overlap Removal | Prevents counting the same wall twice   |

---

## Complexity

* Time Complexity: `O((n + m) log m)`

  * `n` = number of robots
  * `m` = number of walls

* Space Complexity: `O(n)`

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int maxWalls(vector<int>& robots, vector<int>& distance, vector<int>& walls) {
        int n = robots.size();

        vector<pair<int, int>> arr;
        for (int i = 0; i < n; i++) {
            arr.push_back({robots[i], distance[i]});
        }

        sort(arr.begin(), arr.end());
        sort(walls.begin(), walls.end());

        arr.push_back({(int)1e9, 0});

        auto countWalls = [&](int left, int right) {
            if (left > right) return 0;

            int r = upper_bound(walls.begin(), walls.end(), right) - walls.begin();
            int l = lower_bound(walls.begin(), walls.end(), left) - walls.begin();

            return r - l;
        };

        vector<vector<int>> dp(n, vector<int>(2, 0));

        dp[0][0] = countWalls(arr[0].first - arr[0].second, arr[0].first);

        int rightLimit = (n == 1)
            ? arr[0].first + arr[0].second
            : min(arr[0].first + arr[0].second, arr[1].first - 1);

        dp[0][1] = countWalls(arr[0].first, rightLimit);

        for (int i = 1; i < n; i++) {
            int pos = arr[i].first;
            int dist = arr[i].second;

            int reachRight = min(pos + dist, arr[i + 1].first - 1);
            int rightWalls = countWalls(pos, reachRight);

            dp[i][1] = max(dp[i - 1][0], dp[i - 1][1]) + rightWalls;

            int leftStart = max(pos - dist, arr[i - 1].first + 1);
            int leftWalls = countWalls(leftStart, pos);

            dp[i][0] = dp[i - 1][0] + leftWalls;

            int prevRightEnd = min(arr[i - 1].first + arr[i - 1].second, pos - 1);

            int overlapLeft = leftStart;
            int overlapRight = min(prevRightEnd, pos - 1);

            int overlapWalls = countWalls(overlapLeft, overlapRight);

            dp[i][0] = max(dp[i][0], dp[i - 1][1] + leftWalls - overlapWalls);
        }

        return max(dp[n - 1][0], dp[n - 1][1]);
    }
};
```

### Java

```java
class Solution {
    public int maxWalls(int[] robots, int[] distance, int[] walls) {
        int n = robots.length;

        int[][] arr = new int[n + 1][2];

        for (int i = 0; i < n; i++) {
            arr[i][0] = robots[i];
            arr[i][1] = distance[i];
        }

        Arrays.sort(arr, 0, n, (a, b) -> Integer.compare(a[0], b[0]));
        Arrays.sort(walls);

        arr[n][0] = (int) 1e9;
        arr[n][1] = 0;

        int[][] dp = new int[n][2];

        dp[0][0] = countWalls(walls, arr[0][0] - arr[0][1], arr[0][0]);

        int firstRightEnd = (n == 1)
                ? arr[0][0] + arr[0][1]
                : Math.min(arr[0][0] + arr[0][1], arr[1][0] - 1);

        dp[0][1] = countWalls(walls, arr[0][0], firstRightEnd);

        for (int i = 1; i < n; i++) {
            int pos = arr[i][0];
            int dist = arr[i][1];

            int rightEnd = Math.min(pos + dist, arr[i + 1][0] - 1);
            int rightWalls = countWalls(walls, pos, rightEnd);

            dp[i][1] = Math.max(dp[i - 1][0], dp[i - 1][1]) + rightWalls;

            int leftStart = Math.max(pos - dist, arr[i - 1][0] + 1);
            int leftWalls = countWalls(walls, leftStart, pos);

            dp[i][0] = dp[i - 1][0] + leftWalls;

            int prevRightEnd = Math.min(arr[i - 1][0] + arr[i - 1][1], pos - 1);

            int overlapStart = leftStart;
            int overlapEnd = Math.min(prevRightEnd, pos - 1);

            int overlapWalls = countWalls(walls, overlapStart, overlapEnd);

            dp[i][0] = Math.max(dp[i][0], dp[i - 1][1] + leftWalls - overlapWalls);
        }

        return Math.max(dp[n - 1][0], dp[n - 1][1]);
    }

    private int countWalls(int[] walls, int left, int right) {
        if (left > right) return 0;

        int l = lowerBound(walls, left);
        int r = upperBound(walls, right);

        return r - l;
    }

    private int lowerBound(int[] arr, int target) {
        int left = 0, right = arr.length;

        while (left < right) {
            int mid = left + (right - left) / 2;

            if (arr[mid] < target) {
                left = mid + 1;
            } else {
                right = mid;
            }
        }

        return left;
    }

    private int upperBound(int[] arr, int target) {
        int left = 0, right = arr.length;

        while (left < right) {
            int mid = left + (right - left) / 2;

            if (arr[mid] <= target) {
                left = mid + 1;
            } else {
                right = mid;
            }
        }

        return left;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} robots
 * @param {number[]} distance
 * @param {number[]} walls
 * @return {number}
 */
var maxWalls = function(robots, distance, walls) {
    const n = robots.length;

    const arr = [];
    for (let i = 0; i < n; i++) {
        arr.push([robots[i], distance[i]]);
    }

    arr.sort((a, b) => a[0] - b[0]);
    walls.sort((a, b) => a - b);

    arr.push([1e9, 0]);

    const lowerBound = (arr, target) => {
        let left = 0, right = arr.length;

        while (left < right) {
            const mid = Math.floor((left + right) / 2);

            if (arr[mid] < target) left = mid + 1;
            else right = mid;
        }

        return left;
    };

    const upperBound = (arr, target) => {
        let left = 0, right = arr.length;

        while (left < right) {
            const mid = Math.floor((left + right) / 2);

            if (arr[mid] <= target) left = mid + 1;
            else right = mid;
        }

        return left;
    };

    const countWalls = (left, right) => {
        if (left > right) return 0;

        return upperBound(walls, right) - lowerBound(walls, left);
    };

    const dp = Array.from({ length: n }, () => [0, 0]);

    dp[0][0] = countWalls(arr[0][0] - arr[0][1], arr[0][0]);

    const firstRightEnd = n === 1
        ? arr[0][0] + arr[0][1]
        : Math.min(arr[0][0] + arr[0][1], arr[1][0] - 1);

    dp[0][1] = countWalls(arr[0][0], firstRightEnd);

    for (let i = 1; i < n; i++) {
        const pos = arr[i][0];
        const dist = arr[i][1];

        const rightEnd = Math.min(pos + dist, arr[i + 1][0] - 1);
        const rightWalls = countWalls(pos, rightEnd);

        dp[i][1] = Math.max(dp[i - 1][0], dp[i - 1][1]) + rightWalls;

        const leftStart = Math.max(pos - dist, arr[i - 1][0] + 1);
        const leftWalls = countWalls(leftStart, pos);

        dp[i][0] = dp[i - 1][0] + leftWalls;

        const prevRightEnd = Math.min(arr[i - 1][0] + arr[i - 1][1], pos - 1);

        const overlapStart = leftStart;
        const overlapEnd = Math.min(prevRightEnd, pos - 1);

        const overlapWalls = countWalls(overlapStart, overlapEnd);

        dp[i][0] = Math.max(dp[i][0], dp[i - 1][1] + leftWalls - overlapWalls);
    }

    return Math.max(dp[n - 1][0], dp[n - 1][1]);
};
```

### Python3

```python
class Solution:
    def maxWalls(self, robots: List[int], distance: List[int], walls: List[int]) -> int:
        n = len(robots)

        arr = sorted(zip(robots, distance))
        walls.sort()

        arr.append((10**9, 0))

        def count_walls(left, right):
            if left > right:
                return 0

            return bisect_right(walls, right) - bisect_left(walls, left)

        dp = [[0, 0] for _ in range(n)]

        dp[0][0] = count_walls(arr[0][0] - arr[0][1], arr[0][0])

        if n == 1:
            first_right_end = arr[0][0] + arr[0][1]
        else:
            first_right_end = min(arr[0][0] + arr[0][1], arr[1][0] - 1)

        dp[0][1] = count_walls(arr[0][0], first_right_end)

        for i in range(1, n):
            pos, dist = arr[i]

            right_end = min(pos + dist, arr[i + 1][0] - 1)
            right_walls = count_walls(pos, right_end)

            dp[i][1] = max(dp[i - 1][0], dp[i - 1][1]) + right_walls

            left_start = max(pos - dist, arr[i - 1][0] + 1)
            left_walls = count_walls(left_start, pos)

            dp[i][0] = dp[i - 1][0] + left_walls

            prev_right_end = min(arr[i - 1][0] + arr[i - 1][1], pos - 1)

            overlap_start = left_start
            overlap_end = min(prev_right_end, pos - 1)

            overlap_walls = count_walls(overlap_start, overlap_end)

            dp[i][0] = max(dp[i][0], dp[i - 1][1] + left_walls - overlap_walls)

        return max(dp[n - 1][0], dp[n - 1][1])
```

### Go

```go
func maxWalls(robots []int, distance []int, walls []int) int {
 n := len(robots)

 type Robot struct {
  pos  int
  dist int
 }

 arr := make([]Robot, 0, n+1)
 for i := 0; i < n; i++ {
  arr = append(arr, Robot{robots[i], distance[i]})
 }

 sort.Slice(arr, func(i, j int) bool {
  return arr[i].pos < arr[j].pos
 })

 sort.Ints(walls)

 arr = append(arr, Robot{int(1e9), 0})

 lowerBound := func(target int) int {
  left, right := 0, len(walls)

  for left < right {
   mid := left + (right-left)/2

   if walls[mid] < target {
    left = mid + 1
   } else {
    right = mid
   }
  }

  return left
 }

 upperBound := func(target int) int {
  left, right := 0, len(walls)

  for left < right {
   mid := left + (right-left)/2

   if walls[mid] <= target {
    left = mid + 1
   } else {
    right = mid
   }
  }

  return left
 }

 countWalls := func(leftRange int, rightRange int) int {
  if leftRange > rightRange {
   return 0
  }

  return upperBound(rightRange) - lowerBound(leftRange)
 }

 dp := make([][2]int, n)

 dp[0][0] = countWalls(arr[0].pos-arr[0].dist, arr[0].pos)

 firstRightEnd := arr[0].pos + arr[0].dist
 if n > 1 {
  firstRightEnd = min(firstRightEnd, arr[1].pos-1)
 }

 dp[0][1] = countWalls(arr[0].pos, firstRightEnd)

 for i := 1; i < n; i++ {
  pos := arr[i].pos
  dist := arr[i].dist

  rightEnd := min(pos+dist, arr[i+1].pos-1)
  rightWalls := countWalls(pos, rightEnd)

  dp[i][1] = max(dp[i-1][0], dp[i-1][1]) + rightWalls

  leftStart := max(pos-dist, arr[i-1].pos+1)
  leftWalls := countWalls(leftStart, pos)

  dp[i][0] = dp[i-1][0] + leftWalls

  prevRightEnd := min(arr[i-1].pos+arr[i-1].dist, pos-1)

  overlapStart := leftStart
  overlapEnd := min(prevRightEnd, pos-1)

  overlapWalls := countWalls(overlapStart, overlapEnd)

  dp[i][0] = max(dp[i][0], dp[i-1][1]+leftWalls-overlapWalls)
 }

 return max(dp[n-1][0], dp[n-1][1])
}

func min(a, b int) int {
 if a < b {
  return a
 }
 return b
}

func max(a, b int) int {
 if a > b {
  return a
 }
 return b
}
```

---

## Step-by-step Detailed Explanation

### Step 1: Sort robots and walls

I sort all robots by their positions.

This helps me know:

* Which robot is on the left
* Which robot is on the right

I also sort walls because binary search only works on sorted arrays.

---

### Step 2: Calculate valid shooting range

If a robot shoots left:

* It starts from `position - distance`
* But it cannot cross the previous robot

So:

```text
leftStart = max(position - distance, previousRobot + 1)
```

If a robot shoots right:

* It ends at `position + distance`
* But it cannot cross the next robot

So:

```text
rightEnd = min(position + distance, nextRobot - 1)
```

---

### Step 3: Count walls in a range

To count walls inside `[L, R]`:

```text
count = upper_bound(R) - lower_bound(L)
```

This works in logarithmic time.

---

### Step 4: Define DP states

```text
dp[i][0] = maximum walls destroyed till robot i if robot i shoots left

dp[i][1] = maximum walls destroyed till robot i if robot i shoots right
```

---

### Step 5: Transition for shooting right

If current robot shoots right:

```text
dp[i][1] = max(dp[i-1][0], dp[i-1][1]) + wallsDestroyedOnRight
```

---

### Step 6: Transition for shooting left

If current robot shoots left:

Case 1:

* Previous robot also shot left
* No overlap happens

```text
dp[i][0] = dp[i-1][0] + wallsDestroyedOnLeft
```

Case 2:

* Previous robot shot right
* Overlap may happen

```text
dp[i][0] = dp[i-1][1] + leftWalls - overlapWalls
```

---

### Step 7: Handle overlap

Overlap happens only when:

* Previous robot shoots right
* Current robot shoots left

So I find the intersection between:

* Previous robot's right shooting range
* Current robot's left shooting range

Then subtract overlapping walls.

---

## Examples

### Example 1

```text
Input:
robots = [4]
distance = [3]
walls = [1,10]

Output:
1
```

Explanation:

* Robot at position 4 can shoot left
* Shooting range becomes [1, 4]
* Wall at position 1 gets destroyed

---

### Example 2

```text
Input:
robots = [10,2]
distance = [5,1]
walls = [5,2,7]

Output:
3
```

Explanation:

* Robot at 10 shoots left and destroys walls at 5 and 7
* Robot at 2 shoots left and destroys wall at 2
* Total destroyed walls = 3

---

### Example 3

```text
Input:
robots = [1,2]
distance = [100,1]
walls = [10]

Output:
0
```

Explanation:

* Robot at 1 wants to shoot right
* But robot at 2 blocks the bullet
* So wall at 10 cannot be destroyed

---

## How to use / Run locally

### C++

```bash
g++ -std=c++17 solution.cpp -o solution
./solution
```

### Java

```bash
javac Solution.java
java Solution
```

### Python3

```bash
python solution.py
```

### JavaScript

```bash
node solution.js
```

### Go

```bash
go run solution.go
```

---

## Notes & Optimizations

* Sorting is necessary because robot interactions depend on neighboring robots.
* Binary search makes wall counting very fast.
* DP helps avoid recomputing states.
* Overlap handling is the most important part of the problem.
* Without subtracting overlap, the same wall may get counted multiple times.

---

## Author

* [Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
