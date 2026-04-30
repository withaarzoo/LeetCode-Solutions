# 3742. Maximum Path Score in a Grid

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
* [How to Use / Run Locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

## Problem Summary

In this LeetCode grid DP problem, I start from the top-left cell and move only right or down until I reach the bottom-right cell.

Each cell has a value of `0`, `1`, or `2`.

The twist is simple but important: every cell gives score, and some cells also spend cost. A `0` gives no score and no cost. A `1` gives score `1` and cost `1`. A `2` gives score `2` and cost `1`.

The goal is to find the maximum total score I can collect without letting the total cost go above `k`.

The input is:

* a grid of size `m x n`
* an integer `k`

The output is:

* the maximum possible score, or `-1` if no valid path exists

## Constraints

| Constraint             | Meaning                                                         |
| ---------------------- | --------------------------------------------------------------- |
| `1 <= m, n <= 200`     | The grid can be fairly large, so the solution must be efficient |
| `0 <= k <= 1000`       | The allowed cost budget is limited                              |
| `grid[0][0] == 0`      | The start cell is always free                                   |
| `0 <= grid[i][j] <= 2` | Every cell is only one of three values                          |

## Intuition

At first, I noticed that this is not just a normal shortest path or maximum path problem. I cannot only remember the best score at each cell, because two paths reaching the same cell may have used different amounts of cost.

That means I need to track both things together:
the cell position and the cost used so far.

Once I saw that, the problem started looking like dynamic programming on a grid with one extra budget dimension. The movement is limited to right and down, so every cell only depends on the top cell and the left cell. That made the idea feel very natural.

## Approach

I build a DP state for every cell and every possible cost value.

For each cell, I ask this question:

What is the best score I can have when I arrive here using exactly `c` cost?

To answer that, I look at the two possible previous cells:

* from above
* from the left

If the current cell is `0`, it adds `0` score and uses `0` cost.
If it is `1` or `2`, it adds that score and uses `1` cost.

So for each valid budget `c`, I try:

* coming from the top with budget `c - need`
* coming from the left with budget `c - need`

Then I keep the better score.

I only need the previous row and the current row, because each state depends on cells from the same row and the row above. That keeps memory smaller.

At the end, I check all valid costs at the bottom-right cell and take the best one.

## Data Structures Used

I use a 2D DP array for the current row and another 2D DP array for the previous row.

* `prev`: stores the DP values for the row above
* `curr`: stores the DP values for the current row

Each row stores values for every column and every cost from `0` to `k`.

I also use a very small negative number as a sentinel for impossible states. That helps me separate real answers from states that cannot be reached.

## Operations & Behavior Summary

The algorithm works like this:

1. Start from the top-left cell.
2. Mark the initial state with score `0` and cost `0`.
3. Move cell by cell from left to right, top to bottom.
4. For each cell, decide how much score it gives and how much cost it spends.
5. Try to reach that cell from the top and from the left.
6. Keep the best score for every possible cost.
7. After filling the grid, read the best answer from the bottom-right cell.
8. If no valid state exists, return `-1`.

## Complexity

| Type             |     Complexity | Explanation                                                                  |
| ---------------- | -------------: | ---------------------------------------------------------------------------- |
| Time Complexity  | `O(m * n * k)` | I process every cell, and for each cell I may try every cost from `0` to `k` |
| Space Complexity |     `O(n * k)` | I only keep two rows of DP at a time instead of storing the whole 3D table   |

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int maxPathScore(vector<vector<int>>& grid, int k) {
        int m = grid.size();                 // Number of rows in the grid.
        int n = grid[0].size();              // Number of columns in the grid.
        const int NEG = -1e9;                // Sentinel for impossible states.

        // prev[j][c] = best score at column j in the previous row using exact cost c.
        vector<vector<int>> prev(n, vector<int>(k + 1, NEG));

        for (int i = 0; i < m; ++i) {
            // Rebuild the current row from scratch so stale values never leak in.
            vector<vector<int>> curr(n, vector<int>(k + 1, NEG));

            for (int j = 0; j < n; ++j) {
                int gain = grid[i][j];           // Score gained by stepping on this cell.
                int need = (gain > 0 ? 1 : 0);   // Cost spent by this cell: 0 for 0, 1 for 1/2.

                // A path to (i, j) cannot spend more than i + j budget points.
                int limit = min(k, i + j);

                // The start cell is fixed by the statement: it is always 0.
                if (i == 0 && j == 0) {
                    curr[0][0] = 0;              // Start with zero score and zero cost.
                    continue;
                }

                for (int c = need; c <= limit; ++c) {
                    int best = NEG;

                    // Take the path from above, then pay the current cell cost and add its score.
                    if (i > 0 && prev[j][c - need] != NEG) {
                        best = max(best, prev[j][c - need] + gain);
                    }

                    // Take the path from the left, then pay the current cell cost and add its score.
                    if (j > 0 && curr[j - 1][c - need] != NEG) {
                        best = max(best, curr[j - 1][c - need] + gain);
                    }

                    curr[j][c] = best;           // Store the best exact-cost result for this cell.
                }
            }

            prev.swap(curr);                     // Move the current row into prev for the next round.
        }

        int ans = NEG;                           // Best score among all valid costs at the end cell.
        for (int c = 0; c <= k; ++c) {
            ans = max(ans, prev[n - 1][c]);
        }

        return ans < 0 ? -1 : ans;               // If nothing is reachable, the answer is -1.
    }
};
```

### Java

```java
import java.util.Arrays;

class Solution {
    public int maxPathScore(int[][] grid, int k) {
        int m = grid.length;                  // Number of rows in the grid.
        int n = grid[0].length;               // Number of columns in the grid.
        final int NEG = -1_000_000_000;       // Sentinel for impossible states.

        // prev[j][c] = best score at column j in the previous row using exact cost c.
        int[][] prev = new int[n][k + 1];
        for (int j = 0; j < n; j++) {
            Arrays.fill(prev[j], NEG);        // Mark every state as impossible first.
        }

        for (int i = 0; i < m; i++) {
            // Rebuild the current row from scratch so old values do not interfere.
            int[][] curr = new int[n][k + 1];
            for (int j = 0; j < n; j++) {
                Arrays.fill(curr[j], NEG);    // Reset the row before filling new states.
            }

            for (int j = 0; j < n; j++) {
                int gain = grid[i][j];              // Score gained by taking this cell.
                int need = gain > 0 ? 1 : 0;        // Cost spent by this cell: 0 for 0, 1 for 1/2.

                // A path to (i, j) cannot spend more than i + j budget points.
                int limit = Math.min(k, i + j);

                // The starting cell is always 0, so it is the base state.
                if (i == 0 && j == 0) {
                    curr[0][0] = 0;                 // Zero score, zero cost at the start.
                    continue;
                }

                for (int c = need; c <= limit; c++) {
                    int best = NEG;

                    // Coming from above means using the finished previous row.
                    if (i > 0 && prev[j][c - need] != NEG) {
                        best = Math.max(best, prev[j][c - need] + gain);
                    }

                    // Coming from the left means using the current row already built.
                    if (j > 0 && curr[j - 1][c - need] != NEG) {
                        best = Math.max(best, curr[j - 1][c - need] + gain);
                    }

                    curr[j][c] = best;              // Save the best exact-cost answer here.
                }
            }

            prev = curr;                             // Move this row up so it becomes the previous row.
        }

        int ans = NEG;                                // Best score among all valid costs.
        for (int c = 0; c <= k; c++) {
            ans = Math.max(ans, prev[n - 1][c]);
        }

        return ans < 0 ? -1 : ans;                   // Return -1 when no valid path exists.
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[][]} grid
 * @param {number} k
 * @return {number}
 */
var maxPathScore = function(grid, k) {
    const m = grid.length;                  // Number of rows in the grid.
    const n = grid[0].length;               // Number of columns in the grid.
    const NEG = -1e9;                       // Sentinel for impossible states.

    // prev[j][c] = best score at column j in the previous row using exact cost c.
    let prev = Array.from({ length: n }, () => Array(k + 1).fill(NEG));

    for (let i = 0; i < m; i++) {
        // Rebuild the current row from scratch so old values do not leak into new answers.
        let curr = Array.from({ length: n }, () => Array(k + 1).fill(NEG));

        for (let j = 0; j < n; j++) {
            const gain = grid[i][j];           // Score gained by taking this cell.
            const need = gain > 0 ? 1 : 0;    // Cost spent by this cell: 0 for 0, 1 for 1/2.

            // A path to (i, j) cannot spend more than i + j budget points.
            const limit = Math.min(k, i + j);

            // The start cell is fixed and always has value 0.
            if (i === 0 && j === 0) {
                curr[0][0] = 0;               // Zero score, zero cost at the start.
                continue;
            }

            for (let c = need; c <= limit; c++) {
                let best = NEG;

                // From above: use the completed previous row.
                if (i > 0 && prev[j][c - need] !== NEG) {
                    best = Math.max(best, prev[j][c - need] + gain);
                }

                // From left: use the state already computed in this row.
                if (j > 0 && curr[j - 1][c - need] !== NEG) {
                    best = Math.max(best, curr[j - 1][c - need] + gain);
                }

                curr[j][c] = best;            // Store the best exact-cost value for this cell.
            }
        }

        prev = curr;                           // Move the current row into prev for the next iteration.
    }

    let ans = NEG;                             // Best score among all allowed costs at the end.
    for (let c = 0; c <= k; c++) {
        ans = Math.max(ans, prev[n - 1][c]);
    }

    return ans < 0 ? -1 : ans;                // If nothing is reachable, return -1.
};
```

### Python3

```python
from typing import List

class Solution:
    def maxPathScore(self, grid: List[List[int]], k: int) -> int:
        m = len(grid)                         # Number of rows in the grid.
        n = len(grid[0])                      # Number of columns in the grid.
        NEG = -10**9                          # Sentinel for impossible states.

        # prev[j][c] = best score at column j in the previous row with exact cost c.
        prev = [[NEG] * (k + 1) for _ in range(n)]

        for i in range(m):
            # Rebuild the current row from scratch so old values do not interfere.
            curr = [[NEG] * (k + 1) for _ in range(n)]

            for j in range(n):
                gain = grid[i][j]             # Score gained by taking this cell.
                need = 1 if gain > 0 else 0   # Cost spent by this cell: 0 for 0, 1 for 1/2.

                # A path to (i, j) cannot spend more than i + j budget points.
                limit = min(k, i + j)

                # The first cell is the base case.
                if i == 0 and j == 0:
                    curr[0][0] = 0            # Start with zero score and zero cost.
                    continue

                for c in range(need, limit + 1):
                    best = NEG

                    # Take the path from above, then pay for this cell.
                    if i > 0 and prev[j][c - need] != NEG:
                        best = max(best, prev[j][c - need] + gain)

                    # Take the path from the left, then pay for this cell.
                    if j > 0 and curr[j - 1][c - need] != NEG:
                        best = max(best, curr[j - 1][c - need] + gain)

                    curr[j][c] = best         # Save the best exact-cost result for this cell.

            prev = curr                       # Move the current row into prev for the next round.

        ans = max(prev[n - 1])               # Check every allowed cost at the finish cell.
        return -1 if ans < 0 else ans        # If no valid path exists, return -1.
```

### Go

```go
func maxPathScore(grid [][]int, k int) int {
 m := len(grid)                         // Number of rows in the grid.
 n := len(grid[0])                      // Number of columns in the grid.
 const NEG = -1000000000               // Sentinel for impossible states.

 // prev[j][c] = best score at column j in the previous row using exact cost c.
 prev := make([][]int, n)
 for j := 0; j < n; j++ {
  prev[j] = make([]int, k+1)
  for c := 0; c <= k; c++ {
   prev[j][c] = NEG // Mark every state as impossible before the DP starts.
  }
 }

 for i := 0; i < m; i++ {
  // Rebuild the current row from scratch so old values do not leak into new states.
  curr := make([][]int, n)
  for j := 0; j < n; j++ {
   curr[j] = make([]int, k+1)
   for c := 0; c <= k; c++ {
    curr[j][c] = NEG // Reset every budget state for this row.
   }
  }

  for j := 0; j < n; j++ {
   gain := grid[i][j] // Score gained by taking this cell.
   need := 0
   if gain > 0 {
    need = 1 // 1 and 2 both spend one budget point.
   }

   limit := k
   if i+j < limit {
    limit = i + j // A path to (i, j) cannot spend more than i + j budget points.
   }

   // The starting cell is the base case.
   if i == 0 && j == 0 {
    curr[0][0] = 0 // Zero score and zero cost at the start.
    continue
   }

   for c := need; c <= limit; c++ {
    best := NEG

    // From above: use the completed previous row.
    if i > 0 && prev[j][c-need] != NEG {
     val := prev[j][c-need] + gain
     if val > best {
      best = val
     }
    }

    // From left: use the current row already computed.
    if j > 0 && curr[j-1][c-need] != NEG {
     val := curr[j-1][c-need] + gain
     if val > best {
      best = val
     }
    }

    curr[j][c] = best // Store the best exact-cost result for this cell.
   }
  }

  prev = curr // Move the current row into prev for the next iteration.
 }

 ans := NEG                        // Best score among all valid costs at the finish cell.
 for c := 0; c <= k; c++ {
  if prev[n-1][c] > ans {
   ans = prev[n-1][c]
  }
 }

 if ans < 0 {
  return -1 // If nothing is reachable, the answer is invalid.
 }
 return ans
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

I use the same idea in every language. The logic does not change, only the syntax does.

First, I create DP storage for the previous row and the current row. I fill everything with an impossible value so I can clearly detect unreachable states.

Then I process the grid row by row.

For each cell:

* I read the cell value
* I decide its score gain
* I decide its cost usage
* I compute how much budget I am allowed to consider at that position

The base case is the top-left cell. I set it to score `0` and cost `0` because the problem says it starts as `0`.

For every other cell, I look at:

* the state from above
* the state from the left

I use the same remaining cost in both cases. If I can reach the current cell with that budget, I update the best score.

After I finish one row, I move `curr` into `prev` and continue.

At the end, I check all budget values for the last cell. I do this because the answer does not need to use all `k` cost. It only must stay within the limit.

If every state is impossible, I return `-1`.

This same reasoning applies to C++, Java, JavaScript, Python3, and Go. The only difference is how each language creates arrays, fills them, and checks values.

## Examples

### Example 1

Input:
`grid = [[0, 1], [2, 0]]`, `k = 1`

Expected Output:
`2`

Trace:

* Start at `(0,0)` with score `0`, cost `0`
* Move to `(0,1)` and collect score `1`, cost `1`
* Move down to `(1,1)` through the cell `2` on the left path is not needed here
* The best valid path collects a total score of `2` while staying within cost `1`

### Example 2

Input:
`grid = [[0, 1], [1, 2]]`, `k = 1`

Expected Output:
`-1`

Trace:

* Any path to the bottom-right cell needs more than the allowed cost
* So no valid path exists
* The answer is `-1`

### Example 3

Input:
`grid = [[0, 0, 2], [1, 2, 0]]`, `k = 2`

Expected Output:
`4`

Trace:

* The best route takes the `2` cell and one `1` cell
* Total cost stays within `2`
* Total score becomes `4`

## How to Use / Run Locally

For C++, save the solution in a file like `solution.cpp`, then compile it with `g++ -std=c++17 solution.cpp -O2`.

For Java, save it as `Solution.java`, then compile with `javac Solution.java` and run it with `java Solution`.

For JavaScript, save it as `solution.js` and run it with Node.js using `node solution.js`.

For Python3, save it as `solution.py` and run it with `python3 solution.py`.

For Go, save it as `solution.go` and run it with `go run solution.go`.

## Notes & Optimizations

One small optimization is limiting the budget loop by `min(k, i + j)`. A path cannot spend more than the number of steps it has taken, so this saves time.

Another important point is the difference between score and cost. A `2` is better than a `1` because it gives more score for the same cost. That is why the DP must keep the best score for each exact budget.

A simple greedy approach does not work here. Picking the biggest value nearby can waste budget too early. The DP solution is the safe way because it checks every valid cost state.

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
