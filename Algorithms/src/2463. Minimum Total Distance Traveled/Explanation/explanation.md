# 2463. Minimum Total Distance Traveled

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

## Problem Summary

There are some robots and factories placed on the X-axis.

* `robot[i]` represents the position of the `i-th` robot.
* `factory[j] = [positionj, limitj]` means the `j-th` factory is at `positionj` and can repair at most `limitj` robots.

Every robot is broken and can move either left or right.

When a robot reaches a factory that still has remaining capacity, it gets repaired and stops moving.

The goal is to minimize the total distance traveled by all robots.

## Constraints

```text
1 <= robot.length, factory.length <= 100
factory[j].length == 2
-10^9 <= robot[i], positionj <= 10^9
0 <= limitj <= robot.length
The input is always valid such that all robots can be repaired.
```

## Intuition

I first thought about sorting both robots and factories.

After sorting, I noticed that if one factory repairs multiple robots, then those robots will always form a continuous group.

For example:

* If a factory repairs robot `i`
* And also repairs robot `k`
* Then it should also repair every robot between `i` and `k`

That means I can process robots and factories from left to right.

This becomes a Dynamic Programming problem.

I define:

```text
dp[i][j] = minimum total distance needed to repair robots starting from index i using factories starting from index j
```

At every factory, I have two choices:

1. Skip the current factory
2. Use the current factory to repair some number of robots within its limit

Then I take the minimum answer.

## Approach

1. Sort the robots array.
2. Sort the factories by position.
3. Use recursion + memoization.
4. At each factory:

   * Either skip it
   * Or assign `1` to `limit` robots to it
5. Keep adding the travel distance while assigning robots.
6. Store answers in DP table so repeated states are not recalculated.
7. Return the minimum total distance.

## Data Structures Used

* Array / List

  * To store robot positions
  * To store factory position and repair limit

* 2D DP Table

  * `dp[i][j]`
  * Stores minimum distance for robot index `i` and factory index `j`

* Recursion

  * Used to try all valid assignments

## Operations & Behavior Summary

| Operation      | Description                                       |
| -------------- | ------------------------------------------------- |
| Sort robots    | Makes left-to-right processing easier             |
| Sort factories | Allows matching nearby robots to nearby factories |
| Skip factory   | Move directly to next factory                     |
| Use factory    | Assign robots up to its limit                     |
| Memoization    | Avoid recomputing same DP state                   |
| Return minimum | Choose best option among all possibilities        |

## Complexity

* Time Complexity: `O(n * m * n)`

  * `n` = number of robots
  * `m` = number of factories

* Space Complexity: `O(n * m)`

  * Used for memoization table

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    long long dp[101][101];
    const long long INF = 1e18;

    long long solve(int i, int j, vector<int>& robot, vector<vector<int>>& factory) {
        int n = robot.size();
        int m = factory.size();

        if (i == n) return 0;
        if (j == m) return INF;

        if (dp[i][j] != -1) return dp[i][j];

        long long ans = solve(i, j + 1, robot, factory);

        long long distance = 0;
        int pos = factory[j][0];
        int limit = factory[j][1];

        for (int k = 0; k < limit && i + k < n; k++) {
            distance += abs(robot[i + k] - pos);

            long long next = solve(i + k + 1, j + 1, robot, factory);

            if (next != INF) {
                ans = min(ans, distance + next);
            }
        }

        return dp[i][j] = ans;
    }

    long long minimumTotalDistance(vector<int>& robot, vector<vector<int>>& factory) {
        sort(robot.begin(), robot.end());
        sort(factory.begin(), factory.end());

        memset(dp, -1, sizeof(dp));

        return solve(0, 0, robot, factory);
    }
};
```

### Java

```java
class Solution {
    long[][] dp;
    long INF = (long)1e18;

    private long solve(int i, int j, List<Integer> robot, int[][] factory) {
        int n = robot.size();
        int m = factory.length;

        if (i == n) return 0;
        if (j == m) return INF;

        if (dp[i][j] != -1) return dp[i][j];

        long ans = solve(i, j + 1, robot, factory);

        long distance = 0;
        int pos = factory[j][0];
        int limit = factory[j][1];

        for (int k = 0; k < limit && i + k < n; k++) {
            distance += Math.abs(robot.get(i + k) - pos);

            long next = solve(i + k + 1, j + 1, robot, factory);

            if (next != INF) {
                ans = Math.min(ans, distance + next);
            }
        }

        return dp[i][j] = ans;
    }

    public long minimumTotalDistance(List<Integer> robot, int[][] factory) {
        Collections.sort(robot);

        Arrays.sort(factory, (a, b) -> Integer.compare(a[0], b[0]));

        int n = robot.size();
        int m = factory.length;

        dp = new long[n + 1][m + 1];

        for (long[] row : dp) {
            Arrays.fill(row, -1);
        }

        return solve(0, 0, robot, factory);
    }
}
```

### JavaScript

```javascript
var minimumTotalDistance = function(robot, factory) {
    robot.sort((a, b) => a - b);
    factory.sort((a, b) => a[0] - b[0]);

    const n = robot.length;
    const m = factory.length;
    const INF = Number.MAX_SAFE_INTEGER;

    const dp = Array.from({ length: n + 1 }, () => Array(m + 1).fill(-1));

    function solve(i, j) {
        if (i === n) return 0;
        if (j === m) return INF;

        if (dp[i][j] !== -1) return dp[i][j];

        let ans = solve(i, j + 1);

        let distance = 0;
        const [pos, limit] = factory[j];

        for (let k = 0; k < limit && i + k < n; k++) {
            distance += Math.abs(robot[i + k] - pos);

            const next = solve(i + k + 1, j + 1);

            if (next !== INF) {
                ans = Math.min(ans, distance + next);
            }
        }

        return dp[i][j] = ans;
    }

    return solve(0, 0);
};
```

### Python3

```python
class Solution:
    def minimumTotalDistance(self, robot: List[int], factory: List[List[int]]) -> int:
        robot.sort()
        factory.sort()

        n = len(robot)
        m = len(factory)
        INF = float('inf')

        dp = [[-1] * (m + 1) for _ in range(n + 1)]

        def solve(i: int, j: int) -> int:
            if i == n:
                return 0

            if j == m:
                return INF

            if dp[i][j] != -1:
                return dp[i][j]

            ans = solve(i, j + 1)

            distance = 0
            pos, limit = factory[j]

            for k in range(limit):
                if i + k >= n:
                    break

                distance += abs(robot[i + k] - pos)

                next_cost = solve(i + k + 1, j + 1)

                if next_cost != INF:
                    ans = min(ans, distance + next_cost)

            dp[i][j] = ans
            return ans

        return solve(0, 0)
```

### Go

```go
func minimumTotalDistance(robot []int, factory [][]int) int64 {
    sort.Ints(robot)

    sort.Slice(factory, func(i, j int) bool {
        return factory[i][0] < factory[j][0]
    })

    n := len(robot)
    m := len(factory)
    const INF int64 = 1e18

    dp := make([][]int64, n+1)
    for i := range dp {
        dp[i] = make([]int64, m+1)
        for j := range dp[i] {
            dp[i][j] = -1
        }
    }

    var solve func(int, int) int64
    solve = func(i, j int) int64 {
        if i == n {
            return 0
        }

        if j == m {
            return INF
        }

        if dp[i][j] != -1 {
            return dp[i][j]
        }

        ans := solve(i, j+1)

        var distance int64 = 0
        pos := factory[j][0]
        limit := factory[j][1]

        for k := 0; k < limit && i+k < n; k++ {
            diff := robot[i+k] - pos
            if diff < 0 {
                diff = -diff
            }

            distance += int64(diff)

            next := solve(i+k+1, j+1)

            if next != INF {
                if distance+next < ans {
                    ans = distance + next
                }
            }
        }

        dp[i][j] = ans
        return ans
    }

    return solve(0, 0)
}
```

## Step-by-step Detailed Explanation

### 1. Sort robots and factories

```cpp
sort(robot.begin(), robot.end());
sort(factory.begin(), factory.end());
```

I sort both arrays because nearby robots should usually go to nearby factories.

This also helps me process everything from left to right.

### 2. Define DP state

```cpp
dp[i][j]
```

This means:

* `i` = current robot index
* `j` = current factory index
* `dp[i][j]` = minimum distance needed from this state onward

### 3. Base cases

```cpp
if (i == n) return 0;
```

If all robots are already repaired, I do not need any more distance.

```cpp
if (j == m) return INF;
```

If all factories are used but robots are still left, then this state is impossible.

### 4. Skip current factory

```cpp
long long ans = solve(i, j + 1, robot, factory);
```

This means I completely ignore the current factory and move to the next one.

### 5. Assign robots to current factory

```cpp
for (int k = 0; k < limit && i + k < n; k++)
```

I try all possible assignments:

* 1 robot
* 2 robots
* 3 robots
* Up to factory limit

### 6. Keep adding distance

```cpp
distance += abs(robot[i + k] - pos);
```

This stores the total distance traveled by robots assigned to the current factory.

### 7. Move to next state

```cpp
long long next = solve(i + k + 1, j + 1, robot, factory);
```

After assigning some robots to the current factory, I solve the remaining robots using the next factories.

### 8. Update answer

```cpp
ans = min(ans, distance + next);
```

I compare every valid option and keep the minimum possible answer.

## Examples

### Example 1

```text
Input:
robot = [0,4,6]
factory = [[2,2],[6,2]]

Output:
4
```

Explanation:

* Robot at 0 goes to factory at 2
* Robot at 4 goes to factory at 2
* Robot at 6 goes to factory at 6

Total distance:

```text
|2 - 0| + |2 - 4| + |6 - 6| = 2 + 2 + 0 = 4
```

### Example 2

```text
Input:
robot = [1,-1]
factory = [[-2,1],[2,1]]

Output:
2
```

Explanation:

* Robot at -1 goes to factory at -2
* Robot at 1 goes to factory at 2

Total distance:

```text
|(-2) - (-1)| + |2 - 1| = 1 + 1 = 2
```

## How to use / Run locally

### C++

```bash
g++ solution.cpp -o solution
./solution
```

### Java

```bash
javac Solution.java
java Solution
```

### JavaScript

```bash
node solution.js
```

### Python3

```bash
python solution.py
```

### Go

```bash
go run solution.go
```

## Notes & Optimizations

* Sorting is very important.
* Memoization avoids recalculating the same state many times.
* Each factory only repairs a continuous block of robots.
* Returning a very large number for impossible states helps simplify DP transitions.
* Since constraints are only up to `100`, this DP solution is efficient enough.

## Author

* [Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
