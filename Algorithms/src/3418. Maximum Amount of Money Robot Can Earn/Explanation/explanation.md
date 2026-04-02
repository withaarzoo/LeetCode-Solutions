# Maximum Amount of Money Robot Can Earn

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

You are given an `m x n` grid called `coins`.

A robot starts from the top-left corner `(0, 0)` and wants to reach the bottom-right corner `(m - 1, n - 1)`.

The robot can only move:

* Right
* Down

Each cell contains either:

* A positive value → robot gains that many coins
* A negative value → robber steals that many coins

The robot has a special power:

* It can neutralize robbers in at most 2 cells
* If a robber is neutralized, no coins are lost in that cell

The goal is to return the maximum money the robot can collect.

## Constraints

* `m == coins.length`
* `n == coins[i].length`
* `1 <= m, n <= 500`
* `-1000 <= coins[i][j] <= 1000`

## Intuition

I thought of this problem as a Dynamic Programming problem on a grid.

Normally, in a grid DP problem, I only need row and column.

But here, I also need to know how many robber neutralizations I have already used.

So my DP state becomes:

```text
dp[i][j][k]
```

Where:

* `i` = current row
* `j` = current column
* `k` = number of neutralizations already used

Since I can use at most 2 neutralizations, `k` can be:

* `0`
* `1`
* `2`

## Approach

1. Create a 3D DP array.
2. `dp[i][j][k]` stores the maximum money possible at cell `(i, j)` after using `k` neutralizations.
3. From every cell, try moving:

   * Down
   * Right
4. If next cell is positive:

   * Add its value normally
5. If next cell is negative:

   * Either take the loss
   * Or neutralize it if we still have neutralizations left
6. At the end, return the maximum value among all states at the destination cell.

## Data Structures Used

* 3D Dynamic Programming array
* Integer variables for transitions

```text
dp[row][col][usedNeutralizations]
```

## Operations & Behavior Summary

| Operation      | Description                                |
| -------------- | ------------------------------------------ |
| Move Right     | Move from `(i, j)` to `(i, j + 1)`         |
| Move Down      | Move from `(i, j)` to `(i + 1, j)`         |
| Positive Cell  | Add value to current coins                 |
| Negative Cell  | Lose coins                                 |
| Neutralization | Ignore loss in a negative cell             |
| Final Answer   | Maximum among all DP states at destination |

## Complexity

* Time Complexity: `O(m * n * 3)`

  * We process every cell for all 3 neutralization states

* Space Complexity: `O(m * n * 3)`

  * We store DP states for every cell and every neutralization count

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int maximumAmount(vector<vector<int>>& coins) {
        int m = coins.size();
        int n = coins[0].size();

        const int NEG = -1e9;

        vector<vector<vector<int>>> dp(
            m, vector<vector<int>>(n, vector<int>(3, NEG))
        );

        if (coins[0][0] >= 0) {
            dp[0][0][0] = coins[0][0];
        } else {
            dp[0][0][0] = coins[0][0];
            dp[0][0][1] = 0;
        }

        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                for (int k = 0; k <= 2; k++) {
                    if (dp[i][j][k] == NEG) continue;

                    if (i + 1 < m) {
                        int val = coins[i + 1][j];

                        dp[i + 1][j][k] = max(dp[i + 1][j][k], dp[i][j][k] + val);

                        if (val < 0 && k < 2) {
                            dp[i + 1][j][k + 1] = max(dp[i + 1][j][k + 1], dp[i][j][k]);
                        }
                    }

                    if (j + 1 < n) {
                        int val = coins[i][j + 1];

                        dp[i][j + 1][k] = max(dp[i][j + 1][k], dp[i][j][k] + val);

                        if (val < 0 && k < 2) {
                            dp[i][j + 1][k + 1] = max(dp[i][j + 1][k + 1], dp[i][j][k]);
                        }
                    }
                }
            }
        }

        return max({dp[m - 1][n - 1][0], dp[m - 1][n - 1][1], dp[m - 1][n - 1][2]});
    }
};
```

### Java

```java
class Solution {
    public int maximumAmount(int[][] coins) {
        int m = coins.length;
        int n = coins[0].length;
        int NEG = -(int)1e9;

        int[][][] dp = new int[m][n][3];

        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                for (int k = 0; k < 3; k++) {
                    dp[i][j][k] = NEG;
                }
            }
        }

        if (coins[0][0] >= 0) {
            dp[0][0][0] = coins[0][0];
        } else {
            dp[0][0][0] = coins[0][0];
            dp[0][0][1] = 0;
        }

        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                for (int k = 0; k <= 2; k++) {
                    if (dp[i][j][k] == NEG) continue;

                    if (i + 1 < m) {
                        int val = coins[i + 1][j];

                        dp[i + 1][j][k] = Math.max(dp[i + 1][j][k], dp[i][j][k] + val);

                        if (val < 0 && k < 2) {
                            dp[i + 1][j][k + 1] = Math.max(dp[i + 1][j][k + 1], dp[i][j][k]);
                        }
                    }

                    if (j + 1 < n) {
                        int val = coins[i][j + 1];

                        dp[i][j + 1][k] = Math.max(dp[i][j + 1][k], dp[i][j][k] + val);

                        if (val < 0 && k < 2) {
                            dp[i][j + 1][k + 1] = Math.max(dp[i][j + 1][k + 1], dp[i][j][k]);
                        }
                    }
                }
            }
        }

        return Math.max(dp[m - 1][n - 1][0],
               Math.max(dp[m - 1][n - 1][1], dp[m - 1][n - 1][2]));
    }
}
```

### JavaScript

```javascript
var maximumAmount = function(coins) {
    const m = coins.length;
    const n = coins[0].length;
    const NEG = -1e9;

    const dp = Array.from({ length: m }, () =>
        Array.from({ length: n }, () => Array(3).fill(NEG))
    );

    if (coins[0][0] >= 0) {
        dp[0][0][0] = coins[0][0];
    } else {
        dp[0][0][0] = coins[0][0];
        dp[0][0][1] = 0;
    }

    for (let i = 0; i < m; i++) {
        for (let j = 0; j < n; j++) {
            for (let k = 0; k <= 2; k++) {
                if (dp[i][j][k] === NEG) continue;

                if (i + 1 < m) {
                    const val = coins[i + 1][j];

                    dp[i + 1][j][k] = Math.max(dp[i + 1][j][k], dp[i][j][k] + val);

                    if (val < 0 && k < 2) {
                        dp[i + 1][j][k + 1] = Math.max(dp[i + 1][j][k + 1], dp[i][j][k]);
                    }
                }

                if (j + 1 < n) {
                    const val = coins[i][j + 1];

                    dp[i][j + 1][k] = Math.max(dp[i][j + 1][k], dp[i][j][k] + val);

                    if (val < 0 && k < 2) {
                        dp[i][j + 1][k + 1] = Math.max(dp[i][j + 1][k + 1], dp[i][j][k]);
                    }
                }
            }
        }
    }

    return Math.max(
        dp[m - 1][n - 1][0],
        dp[m - 1][n - 1][1],
        dp[m - 1][n - 1][2]
    );
};
```

### Python3

```python
class Solution:
    def maximumAmount(self, coins: List[List[int]]) -> int:
        m, n = len(coins), len(coins[0])
        NEG = -10**9

        dp = [[[NEG] * 3 for _ in range(n)] for _ in range(m)]

        if coins[0][0] >= 0:
            dp[0][0][0] = coins[0][0]
        else:
            dp[0][0][0] = coins[0][0]
            dp[0][0][1] = 0

        for i in range(m):
            for j in range(n):
                for k in range(3):
                    if dp[i][j][k] == NEG:
                        continue

                    if i + 1 < m:
                        val = coins[i + 1][j]

                        dp[i + 1][j][k] = max(dp[i + 1][j][k], dp[i][j][k] + val)

                        if val < 0 and k < 2:
                            dp[i + 1][j][k + 1] = max(dp[i + 1][j][k + 1], dp[i][j][k])

                    if j + 1 < n:
                        val = coins[i][j + 1]

                        dp[i][j + 1][k] = max(dp[i][j + 1][k], dp[i][j][k] + val)

                        if val < 0 and k < 2:
                            dp[i][j + 1][k + 1] = max(dp[i][j + 1][k + 1], dp[i][j][k])

        return max(dp[m - 1][n - 1])
```

### Go

```go
func maximumAmount(coins [][]int) int {
    m := len(coins)
    n := len(coins[0])
    NEG := -1000000000

    dp := make([][][]int, m)
    for i := 0; i < m; i++ {
        dp[i] = make([][]int, n)
        for j := 0; j < n; j++ {
            dp[i][j] = make([]int, 3)
            for k := 0; k < 3; k++ {
                dp[i][j][k] = NEG
            }
        }
    }

    if coins[0][0] >= 0 {
        dp[0][0][0] = coins[0][0]
    } else {
        dp[0][0][0] = coins[0][0]
        dp[0][0][1] = 0
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            for k := 0; k <= 2; k++ {
                if dp[i][j][k] == NEG {
                    continue
                }

                if i+1 < m {
                    val := coins[i+1][j]

                    dp[i+1][j][k] = max(dp[i+1][j][k], dp[i][j][k]+val)

                    if val < 0 && k < 2 {
                        dp[i+1][j][k+1] = max(dp[i+1][j][k+1], dp[i][j][k])
                    }
                }

                if j+1 < n {
                    val := coins[i][j+1]

                    dp[i][j+1][k] = max(dp[i][j+1][k], dp[i][j][k]+val)

                    if val < 0 && k < 2 {
                        dp[i][j+1][k+1] = max(dp[i][j+1][k+1], dp[i][j][k])
                    }
                }
            }
        }
    }

    ans := dp[m-1][n-1][0]
    ans = max(ans, dp[m-1][n-1][1])
    ans = max(ans, dp[m-1][n-1][2])

    return ans
}
```

## Step-by-step Detailed Explanation

### C++, Java, JavaScript, Python3, Go

1. Create a 3D DP array.
2. Each state stores:

```text
dp[i][j][k]
```

1. Here:

   * `i` = row
   * `j` = column
   * `k` = neutralizations used

2. Initialize all states with a very small negative value.

3. Handle the starting cell separately.

4. From every cell:

   * Move Down
   * Move Right

5. If next cell is positive:

   * Add its value

6. If next cell is negative:

   * Option 1: Take the loss
   * Option 2: Neutralize it

7. Keep updating the best possible value.

8. Return the maximum value at the destination cell.

## Examples

### Example 1

```text
Input:
coins = [[0,1,-1],[1,-2,3],[2,-3,4]]

Output:
8
```

Explanation:

* Start at `(0, 0)` with `0`
* Move right to `(0, 1)` and gain `1`
* Move down to `(1, 1)` which is `-2`
* Neutralize robber there
* Move right to `(1, 2)` and gain `3`
* Move down to `(2, 2)` and gain `4`
* Final answer = `8`

### Example 2

```text
Input:
coins = [[10,10,10],[10,10,10]]

Output:
40
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

* A normal 2D DP is not enough because we must also track neutralizations.
* Using 3 states for neutralizations keeps the solution efficient.
* Since only 3 states are used, the extra memory is manageable.
* This solution works well within the given constraints.

## Author

* [Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
