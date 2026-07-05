# 1301. Number of Paths with Max Score

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

In LeetCode 1301, Number of Paths with Max Score, we are given a square board containing digits, obstacles, a starting point `S`, and an ending point `E`.

The starting point `S` is at the bottom-right corner, while `E` is at the top-left corner.

From any cell, we can move in only three directions:

* Up
* Left
* Diagonally up-left

A cell containing `X` is an obstacle, so we cannot move through it.

Every digit we visit adds to our score. The main goal is not just to find a valid path. We need to find the maximum score possible among all valid paths from `S` to `E`.

We also need to count how many different paths achieve that exact maximum score.

The final answer contains two values:

```text
[maximum score, number of maximum-score paths]
```

The number of paths must be returned modulo `10^9 + 7`.

If there is no valid path from `S` to `E`, the answer is:

```text
[0, 0]
```

This problem is a good example of dynamic programming on a matrix because every cell needs to track both the best score and the number of ways to achieve it.

## Constraints

| Constraint                                                 | Meaning                                            |
| ---------------------------------------------------------- | -------------------------------------------------- |
| `2 <= board.length <= 100`                                 | The board has at least 2 rows and at most 100 rows |
| `board.length == board[i].length`                          | The board is always square                         |
| `board[i][j]` is `E`, `S`, `X`, or a digit from `1` to `9` | Every cell has a valid character                   |
| `E` is at the top-left corner                              | The destination is fixed                           |
| `S` is at the bottom-right corner                          | The starting point is fixed                        |
| Path count is returned modulo `10^9 + 7`                   | This prevents very large path counts               |

## Intuition

My first thought was to start from `S` and explore every possible path to `E`.

That approach is easy to understand, but it becomes too slow because the same cells can be reached through many different paths. A recursive brute-force solution would repeat the same work again and again.

I then noticed that every cell only needs two answers:

1. What is the maximum score possible from this cell to `S`?
2. How many paths produce that maximum score?

This immediately suggests dynamic programming.

Instead of moving from `S` toward `E`, I process the board in reverse, from the bottom-right corner toward the top-left corner.

For any cell, I only need information from three nearby cells:

* The cell below
* The cell to the right
* The bottom-right diagonal cell

These are the reverse directions of the moves allowed in the original problem.

For each cell, I choose the largest score among these three possible next positions. If multiple positions have the same largest score, I add their path counts together.

This gives both the maximum path score and the number of paths with that score.

## Approach

I use dynamic programming with space optimization.

The board is processed from bottom to top. Inside each row, cells are processed from right to left.

For every cell, I store two values:

* The best score possible from that cell to `S`
* The number of paths that achieve that best score

I use two rows instead of a full two-dimensional DP table.

The previous row stores information about the row below the current cell. The current row stores information that I am calculating now.

For a cell at position `(i, j)`, I check:

1. `(i + 1, j)` for the cell below
2. `(i, j + 1)` for the cell to the right
3. `(i + 1, j + 1)` for the diagonal cell

I find the maximum reachable score among these three cells.

If none of them can reach `S`, the current cell is also unreachable.

If one or more cells have the maximum score, I add the number of paths from all of them.

Then I add the digit value of the current cell to the score.

The `E` and `S` cells add `0`.

An obstacle is skipped completely.

After finishing a row, the current row becomes the previous row for the next iteration.

Once the top-left cell is processed, it contains the final maximum score and the number of maximum-score paths.

## Data Structures Used

### String Array or List

The board is stored as an array or list of strings.

Each character represents one cell:

* `S` for the starting point
* `E` for the ending point
* `X` for an obstacle
* `1` to `9` for score values

This format makes it easy to access any cell using its row and column position.

### Score Arrays

Two one-dimensional arrays store the maximum score values:

* One for the row below
* One for the current row

A score of `-1` represents an unreachable cell.

Using `-1` is safe because every valid path score is `0` or greater.

### Path Count Arrays

Two more one-dimensional arrays store the number of maximum-score paths:

* One for the row below
* One for the current row

The path count is always taken modulo `10^9 + 7`.

Using rolling arrays keeps the space complexity at `O(n)` instead of `O(n²)`.

## Operations & Behavior Summary

The algorithm follows these main stages:

1. Create score and path-count arrays for the row below.
2. Mark all initial score positions as unreachable.
3. Start processing from the bottom row.
4. Move through each row from right to left.
5. Skip the cell if it contains an obstacle.
6. If the cell is `S`, set its score to `0` and path count to `1`.
7. For every other cell, check the three possible next positions.
8. Find the highest score among the reachable positions.
9. If no next position is reachable, leave the current cell unreachable.
10. Add the path counts from every position that has the highest score.
11. Add the current digit value to the score.
12. Store the result for the current cell.
13. After completing the row, reuse it as the row below.
14. Check the result stored for `E`.
15. Return `[0, 0]` if `E` is unreachable.
16. Otherwise, return the maximum score and the number of paths.

## Complexity

| Complexity       | Value   | Explanation                                                                                          |
| ---------------- | ------- | ---------------------------------------------------------------------------------------------------- |
| Time Complexity  | `O(n²)` | The algorithm visits every cell of the `n × n` board exactly once and checks only three nearby cells |
| Space Complexity | `O(n)`  | Only two rows of scores and two rows of path counts are stored                                       |

Here, `n` is the number of rows and columns in the square board.

A full dynamic programming table would use `O(n²)` extra space. The rolling-array optimization reduces this to `O(n)` because only the current row and the row below are needed.

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    vector<int> pathsWithMaxScore(vector<string>& board) {
        const int MOD = 1000000007;
        int n = board.size();

        // These arrays store the DP values for the row below.
        // A score of -1 means the cell cannot reach S.
        vector<int> nextScore(n + 1, -1);
        vector<int> nextWays(n + 1, 0);

        // I process rows from bottom to top.
        for (int i = n - 1; i >= 0; --i) {
            // Fresh arrays are needed for the current row.
            vector<int> currScore(n + 1, -1);
            vector<int> currWays(n + 1, 0);

            // I process right to left so the right cell is already solved.
            for (int j = n - 1; j >= 0; --j) {
                char cell = board[i][j];

                // An obstacle can never be part of a valid path.
                if (cell == 'X') {
                    continue;
                }

                // S is the starting point of the reversed DP.
                if (cell == 'S') {
                    currScore[j] = 0;
                    currWays[j] = 1;
                    continue;
                }

                // Find the best score among down, right, and diagonal.
                int best = max({
                    nextScore[j],
                    currScore[j + 1],
                    nextScore[j + 1]
                });

                // If all three cells are unreachable, this cell is unreachable too.
                if (best == -1) {
                    continue;
                }

                long long ways = 0;

                // Add counts only from cells that achieve the best score.
                if (nextScore[j] == best) {
                    ways += nextWays[j];
                }
                if (currScore[j + 1] == best) {
                    ways += currWays[j + 1];
                }
                if (nextScore[j + 1] == best) {
                    ways += nextWays[j + 1];
                }

                // E adds no points; a digit adds its numeric value.
                int value = (cell == 'E') ? 0 : cell - '0';

                currScore[j] = best + value;
                currWays[j] = ways % MOD;
            }

            // The current row becomes the row below for the next iteration.
            nextScore = move(currScore);
            nextWays = move(currWays);
        }

        // If E cannot reach S, the required answer is [0, 0].
        if (nextScore[0] == -1) {
            return {0, 0};
        }

        return {nextScore[0], nextWays[0]};
    }
};
```

### Java

```java
class Solution {
    public int[] pathsWithMaxScore(List<String> board) {
        final int MOD = 1_000_000_007;
        int n = board.size();

        // These arrays store DP values for the row below.
        int[] nextScore = new int[n + 1];
        int[] nextWays = new int[n + 1];

        // -1 marks every score as unreachable at the beginning.
        Arrays.fill(nextScore, -1);

        // I process the board from bottom to top.
        for (int i = n - 1; i >= 0; i--) {
            // These arrays store values for the current row.
            int[] currScore = new int[n + 1];
            int[] currWays = new int[n + 1];

            // Every current-row cell starts as unreachable.
            Arrays.fill(currScore, -1);

            // I move right to left so the right cell is already solved.
            for (int j = n - 1; j >= 0; j--) {
                char cell = board.get(i).charAt(j);

                // Obstacles cannot be used.
                if (cell == 'X') {
                    continue;
                }

                // S starts with score 0 and exactly one path.
                if (cell == 'S') {
                    currScore[j] = 0;
                    currWays[j] = 1;
                    continue;
                }

                // Check down, right, and bottom-right diagonal.
                int best = Math.max(
                    nextScore[j],
                    Math.max(currScore[j + 1], nextScore[j + 1])
                );

                // No reachable next cell means this cell is also unreachable.
                if (best == -1) {
                    continue;
                }

                long ways = 0;

                // Count every next cell that gives the maximum score.
                if (nextScore[j] == best) {
                    ways += nextWays[j];
                }
                if (currScore[j + 1] == best) {
                    ways += currWays[j + 1];
                }
                if (nextScore[j + 1] == best) {
                    ways += nextWays[j + 1];
                }

                // E contributes 0; digit cells contribute their digit.
                int value = (cell == 'E') ? 0 : cell - '0';

                currScore[j] = best + value;
                currWays[j] = (int) (ways % MOD);
            }

            // Reuse the completed current row as the next row.
            nextScore = currScore;
            nextWays = currWays;
        }

        // An unreachable E means no valid path exists.
        if (nextScore[0] == -1) {
            return new int[]{0, 0};
        }

        return new int[]{nextScore[0], nextWays[0]};
    }
}
```

### JavaScript

```javascript
/**
 * @param {string[]} board
 * @return {number[]}
 */
var pathsWithMaxScore = function(board) {
    const MOD = 1000000007;
    const n = board.length;

    // These arrays store DP values for the row below.
    // -1 means that a cell cannot reach S.
    let nextScore = new Array(n + 1).fill(-1);
    let nextWays = new Array(n + 1).fill(0);

    // I process rows from bottom to top.
    for (let i = n - 1; i >= 0; i--) {
        // Fresh arrays hold the current row.
        const currScore = new Array(n + 1).fill(-1);
        const currWays = new Array(n + 1).fill(0);

        // I process right to left so the right cell is ready.
        for (let j = n - 1; j >= 0; j--) {
            const cell = board[i][j];

            // Obstacles are never reachable.
            if (cell === 'X') {
                continue;
            }

            // S is the base case of the reversed DP.
            if (cell === 'S') {
                currScore[j] = 0;
                currWays[j] = 1;
                continue;
            }

            // Find the best score from down, right, or diagonal.
            const best = Math.max(
                nextScore[j],
                currScore[j + 1],
                nextScore[j + 1]
            );

            // This cell cannot reach S if all next cells are unreachable.
            if (best === -1) {
                continue;
            }

            let ways = 0;

            // Add only paths that continue with the maximum score.
            if (nextScore[j] === best) {
                ways = (ways + nextWays[j]) % MOD;
            }
            if (currScore[j + 1] === best) {
                ways = (ways + currWays[j + 1]) % MOD;
            }
            if (nextScore[j + 1] === best) {
                ways = (ways + nextWays[j + 1]) % MOD;
            }

            // E adds 0; digit cells add their numeric value.
            const value = cell === 'E' ? 0 : Number(cell);

            currScore[j] = best + value;
            currWays[j] = ways;
        }

        // The current row becomes the row below.
        nextScore = currScore;
        nextWays = currWays;
    }

    // No reachable value at E means no valid path exists.
    if (nextScore[0] === -1) {
        return [0, 0];
    }

    return [nextScore[0], nextWays[0]];
};
```

### Python3

```python
class Solution:
    def pathsWithMaxScore(self, board: List[str]) -> List[int]:
        MOD = 10**9 + 7
        n = len(board)

        # These arrays store DP values for the row below.
        # A score of -1 means the cell cannot reach S.
        next_score = [-1] * (n + 1)
        next_ways = [0] * (n + 1)

        # I process rows from bottom to top.
        for i in range(n - 1, -1, -1):
            # Fresh arrays store the current row.
            curr_score = [-1] * (n + 1)
            curr_ways = [0] * (n + 1)

            # I process right to left so the right cell is already solved.
            for j in range(n - 1, -1, -1):
                cell = board[i][j]

                # Obstacles cannot be part of any valid path.
                if cell == 'X':
                    continue

                # S is the base case for the reversed DP.
                if cell == 'S':
                    curr_score[j] = 0
                    curr_ways[j] = 1
                    continue

                # Check down, right, and bottom-right diagonal.
                best = max(
                    next_score[j],
                    curr_score[j + 1],
                    next_score[j + 1]
                )

                # If no next cell can reach S, this cell cannot either.
                if best == -1:
                    continue

                ways = 0

                # Add counts only from next cells with the best score.
                if next_score[j] == best:
                    ways += next_ways[j]
                if curr_score[j + 1] == best:
                    ways += curr_ways[j + 1]
                if next_score[j + 1] == best:
                    ways += next_ways[j + 1]

                # E contributes 0; a digit contributes its integer value.
                value = 0 if cell == 'E' else int(cell)

                curr_score[j] = best + value
                curr_ways[j] = ways % MOD

            # The completed row becomes the row below.
            next_score = curr_score
            next_ways = curr_ways

        # If E cannot reach S, there is no valid path.
        if next_score[0] == -1:
            return [0, 0]

        return [next_score[0], next_ways[0]]
```

### Go

```go
func pathsWithMaxScore(board []string) []int {
    const MOD int = 1000000007
    n := len(board)

    // These arrays store DP values for the row below.
    nextScore := make([]int, n+1)
    nextWays := make([]int, n+1)

    // -1 means that a cell cannot reach S.
    for j := 0; j <= n; j++ {
        nextScore[j] = -1
    }

    // I process rows from bottom to top.
    for i := n - 1; i >= 0; i-- {
        // Fresh arrays store the current row.
        currScore := make([]int, n+1)
        currWays := make([]int, n+1)

        // Every score starts as unreachable.
        for j := 0; j <= n; j++ {
            currScore[j] = -1
        }

        // I process right to left so the right cell is already solved.
        for j := n - 1; j >= 0; j-- {
            cell := board[i][j]

            // Obstacles cannot be used.
            if cell == 'X' {
                continue
            }

            // S is the base case of the reversed DP.
            if cell == 'S' {
                currScore[j] = 0
                currWays[j] = 1
                continue
            }

            // Find the best score among down, right, and diagonal.
            best := nextScore[j]
            if currScore[j+1] > best {
                best = currScore[j+1]
            }
            if nextScore[j+1] > best {
                best = nextScore[j+1]
            }

            // No reachable next cell means this cell is unreachable.
            if best == -1 {
                continue
            }

            ways := 0

            // Count every next cell that gives the best score.
            if nextScore[j] == best {
                ways = (ways + nextWays[j]) % MOD
            }
            if currScore[j+1] == best {
                ways = (ways + currWays[j+1]) % MOD
            }
            if nextScore[j+1] == best {
                ways = (ways + nextWays[j+1]) % MOD
            }

            // E adds no points; digit cells add their digit value.
            value := 0
            if cell != 'E' {
                value = int(cell - '0')
            }

            currScore[j] = best + value
            currWays[j] = ways
        }

        // The current row becomes the row below.
        nextScore = currScore
        nextWays = currWays
    }

    // An unreachable E means no valid path exists.
    if nextScore[0] == -1 {
        return []int{0, 0}
    }

    return []int{nextScore[0], nextWays[0]}
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The core dynamic programming logic is the same in C++, Java, JavaScript, Python3, and Go. Only the syntax and array handling are different.

### Step 1: Store the Board Size

The first step is to get the size of the square board.

If the board has `n` rows, it also has `n` columns.

This value controls both loops and the size of the dynamic programming arrays.

### Step 2: Create the Rolling DP Arrays

I create two arrays for scores and two arrays for path counts.

The first pair represents the row below the current row.

The second pair represents the row currently being processed.

The score arrays start with `-1`.

This value means that the cell is unreachable.

The path-count arrays start with `0` because an unreachable cell has no valid paths.

### Step 3: Process Rows from Bottom to Top

The original path moves from `S` toward `E`.

However, the dynamic programming solution works more naturally in reverse.

I begin at the bottom row and move upward.

This ensures that when I process a cell, the row below has already been solved.

### Step 4: Process Columns from Right to Left

Inside each row, I move from the last column toward the first column.

This order is necessary because the current cell depends on the cell to its right.

By moving right to left, that right-side value is already available.

If I moved from left to right, the right cell would not be calculated yet.

### Step 5: Handle Obstacles

If the current cell contains `X`, I skip it.

The score stays at `-1`, and the path count stays at `0`.

This automatically prevents other cells from using the obstacle as part of a valid path.

### Step 6: Handle the Starting Cell

The `S` cell is the base case.

Its score is `0` because it does not add any points.

Its path count is `1` because being at `S` represents one completed path in the reversed dynamic programming process.

Without this base case, no other cell could become reachable.

### Step 7: Check the Three Possible Next Cells

For every normal cell, I check:

* The cell below
* The cell to the right
* The bottom-right diagonal cell

These three cells represent the reversed version of the original movement rules.

I compare their scores and find the largest one.

### Step 8: Detect an Unreachable Cell

If all three possible next cells have a score of `-1`, the current cell cannot reach `S`.

In that case, I do nothing.

The current cell remains unreachable.

This also handles situations where obstacles completely block part of the board.

### Step 9: Count Only Maximum-Score Paths

After finding the best score, I check all three next cells again.

If a cell has that exact best score, I add its path count.

For example, suppose the next scores are:

```text
7, 7, 5
```

The maximum is `7`.

The first two cells both contribute to the answer.

The third cell is ignored because its paths cannot produce the maximum total score.

This is the most important part of the counting logic.

### Step 10: Add the Current Cell Value

If the current cell contains a digit, I convert it into its numeric value and add it to the best score.

If the current cell is `E`, I add `0`.

The `S` cell was already handled separately.

### Step 11: Apply the Modulo

The number of valid maximum-score paths can become very large.

For that reason, I store the path count modulo:

```text
1,000,000,007
```

This keeps the values within the required limit.

### Step 12: Reuse the Current Row

After finishing one row, the current row becomes the row below for the next iteration.

The older row is no longer needed.

This rolling-array technique is what reduces the extra space from `O(n²)` to `O(n)`.

### Step 13: Build the Final Answer

After all rows are processed, the top-left position contains the result for `E`.

If its score is still `-1`, there is no valid path.

The answer is:

```text
[0, 0]
```

Otherwise, the algorithm returns:

```text
[maximum score, number of maximum-score paths]
```

### Language-Specific Behavior

The algorithm itself does not change between languages.

In C++, `vector<int>` is used for the rolling arrays.

In Java, integer arrays are used, and unreachable score arrays are filled with `-1`.

In JavaScript, standard arrays are created with `.fill()`.

In Python3, lists are used for scores and path counts.

In Go, slices are used, and the score slices are manually initialized with `-1`.

These differences only affect syntax. The dynamic programming state, traversal order, maximum-score comparison, and path-counting logic remain exactly the same.

## Examples

### Example 1

Input:

```text
board = ["E23", "2X2", "12S"]
```

Expected Output:

```text
[7, 1]
```

The algorithm starts from `S` with score `0` and one valid path.

It processes the board from bottom-right to top-left.

The obstacle in the center blocks some routes.

Among all valid paths, the maximum collected score is `7`.

Only one path achieves that score.

So the final answer is:

```text
[7, 1]
```

### Example 2

Input:

```text
board = ["E12", "1X1", "21S"]
```

Expected Output:

```text
[4, 2]
```

The center cell is blocked.

There are two different valid paths that collect the maximum score of `4`.

When the dynamic programming process finds equal best scores from multiple directions, it adds their path counts together.

So the final answer is:

```text
[4, 2]
```

### Example 3

Input:

```text
board = ["E11", "XXX", "11S"]
```

Expected Output:

```text
[0, 0]
```

The middle row contains only obstacles.

There is no way to travel between `S` and `E`.

The top-left cell remains unreachable.

So the algorithm returns:

```text
[0, 0]
```

## How to Use / Run Locally

Before running any version, place the solution code inside a file and add a small test driver if needed. LeetCode provides the input and calls the solution method automatically, but local execution requires a `main` function or equivalent test code.

### C++

Save the code in a file named:

```text
solution.cpp
```

Compile it:

```bash
g++ -std=c++17 solution.cpp -o solution
```

Run it:

```bash
./solution
```

On Windows, run:

```bash
solution.exe
```

### Java

Save the code in:

```text
Solution.java
```

Compile it:

```bash
javac Solution.java
```

Run it:

```bash
java Solution
```

Make sure the file name matches the public class name.

### JavaScript

Save the code in:

```text
solution.js
```

Run it with Node.js:

```bash
node solution.js
```

Node.js must be installed on the system.

### Python3

Save the code in:

```text
solution.py
```

Run it:

```bash
python3 solution.py
```

On some systems, the command may be:

```bash
python solution.py
```

### Go

Save the code in:

```text
main.go
```

Run it directly:

```bash
go run main.go
```

Or build an executable:

```bash
go build main.go
```

Then run the generated executable.

## Notes & Optimizations

The most important edge case is an unreachable destination.

Using `-1` as the unreachable score makes this easy to detect because every real path score is non-negative.

Another important case is when multiple next cells have the same maximum score. Their path counts must be added together. Counting only one of them would produce the correct score but the wrong number of paths.

A full two-dimensional dynamic programming solution is also possible. It is often easier to visualize because every cell stores its own result.

That version uses:

```text
O(n²) time
O(n²) space
```

The rolling-array version keeps the same `O(n²)` time complexity but reduces the extra space to `O(n)`.

A recursive solution with memoization can also solve the problem, but the iterative bottom-up approach avoids recursion overhead and makes the space optimization easier.

The processing order must remain bottom to top and right to left. Changing this order without changing the DP dependencies can cause the algorithm to read values that have not been calculated yet.

This solution is useful for learning several common DSA patterns:

* Dynamic programming on a grid
* Maximum path score calculation
* Counting optimal paths
* Handling obstacles in a matrix
* Rolling-array space optimization
* Combining optimization and path counting in one DP state

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
