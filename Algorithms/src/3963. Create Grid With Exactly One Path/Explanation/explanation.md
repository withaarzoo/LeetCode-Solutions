# Create Grid With Exactly One Path - LeetCode 3963 Solution

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

You are given two integers `m` and `n`. These represent the number of rows and columns in a grid.

Your task is to build any `m x n` grid using only two characters:

- `.` for a free cell
- `#` for an obstacle cell

A valid path is a sequence of free cells that starts at the top-left cell `(0, 0)` and ends at the bottom-right cell `(m-1, n-1)`. You can only move right or down.

The goal is to return any grid that contains **exactly one** valid path from the top-left corner to the bottom-right corner.

This is a classic constructive problem often seen in competitive programming and LeetCode path construction challenges.

## Constraints

- `1 <= m, n <= 25`

## Intuition

When I first read the problem, I realized I do not need to count paths or search for them. I just need to force the grid to have only one possible route.

The simplest idea that came to mind was to choose one clear path and block everything else. Going all the way right on the first row and then all the way down the last column felt like the most straightforward route. Once that path is open and every other cell is blocked, no other sequence of right and down moves can reach the end.

This approach works cleanly even when the grid is just a single row or single column.

## Approach

I start by creating a full grid filled with obstacles (`#`).

Then I open exactly one path:

1. Make every cell in the first row free (`.`).
2. Make every cell in the last column free (`.`).

Everything else stays blocked.

From `(0, 0)` the only free moves are right along the top row until the last column. After that, the only free moves are down the last column until the bottom-right cell. Any attempt to go down earlier hits an obstacle, so the path is unique.

This construction satisfies the condition of having exactly one valid path while staying within the given constraints.

## Data Structures Used

- A 2D grid (vector of strings in C++, array of strings in Java, list of lists or 2D arrays in other languages) to store the final answer.
- Simple loops to fill and update the grid cells.

No advanced data structures are required because the problem only asks us to construct a valid grid.

## Operations & Behavior Summary

1. Allocate an `m` by `n` grid and fill every cell with `#`.
2. Set all cells in row `0` to `.`.
3. Set all cells in column `n-1` to `.`.
4. Return the completed grid.

These three steps guarantee that only one sequence of right and down moves exists from start to finish.

## Complexity

| Type              | Complexity | Explanation                                      |
|-------------------|------------|--------------------------------------------------|
| Time Complexity   | O(m * n)   | We visit every cell a constant number of times while building the grid. |
| Space Complexity  | O(m * n)   | The output grid itself requires this much space. No significant extra space is used. |

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    vector<string> createGrid(int m, int n) {
        vector<string> grid(m, string(n, '#'));
        for (int j = 0; j < n; ++j) {
            grid[0][j] = '.';
        }
        for (int i = 0; i < m; ++i) {
            grid[i][n - 1] = '.';
        }
        return grid;
    }
};
```

### Java

```java
class Solution {
    public String[] createGrid(int m, int n) {
        String[] grid = new String[m];
        char[] row = new char[n];
        for (int j = 0; j < n; ++j) {
            row[j] = '#';
        }
        for (int i = 0; i < m; ++i) {
            grid[i] = new String(row);
        }
        char[] first = grid[0].toCharArray();
        for (int j = 0; j < n; ++j) {
            first[j] = '.';
        }
        grid[0] = new String(first);
        for (int i = 0; i < m; ++i) {
            char[] r = grid[i].toCharArray();
            r[n - 1] = '.';
            grid[i] = new String(r);
        }
        return grid;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number} m
 * @param {number} n
 * @return {string[]}
 */
var createGrid = function(m, n) {
    let grid = Array.from({length: m}, () => Array(n).fill('#'));
    for (let j = 0; j < n; ++j) {
        grid[0][j] = '.';
    }
    for (let i = 0; i < m; ++i) {
        grid[i][n - 1] = '.';
    }
    return grid.map(row => row.join(''));
};
```

### TypeScript

```typescript
function createGrid(m: number, n: number): string[] {
    let grid: string[][] = Array.from({length: m}, () => Array(n).fill('#'));
    for (let j = 0; j < n; ++j) {
        grid[0][j] = '.';
    }
    for (let i = 0; i < m; ++i) {
        grid[i][n - 1] = '.';
    }
    return grid.map(row => row.join(''));
};
```

### Python3

```python
class Solution:
    def createGrid(self, m: int, n: int) -> list[str]:
        grid = [['#'] * n for _ in range(m)]
        for j in range(n):
            grid[0][j] = '.'
        for i in range(m):
            grid[i][n - 1] = '.'
        return [''.join(row) for row in grid]
```

### Go

```go
func createGrid(m int, n int) []string {
    grid := make([][]byte, m)
    for i := 0; i < m; i++ {
        grid[i] = make([]byte, n)
        for j := 0; j < n; j++ {
            grid[i][j] = '#'
        }
    }
    for j := 0; j < n; j++ {
        grid[0][j] = '.'
    }
    for i := 0; i < m; i++ {
        grid[i][n-1] = '.'
    }
    res := make([]string, m)
    for i := 0; i < m; i++ {
        res[i] = string(grid[i])
    }
    return res
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

Across all languages the logic stays the same. I first create a grid full of obstacles. This makes sure that any cell I forget to open will remain blocked and cannot create an extra path.

Next I open the top row completely. After this step the only free cells on the top row are the ones I just set. From any position on the top row (except the last column) the cell below is still an obstacle, so I cannot leave the top row early.

Then I open the entire last column. Now the top-right corner connects downward all the way to the bottom-right corner. Because the last column is the only free vertical path, there is no way to branch sideways.

When `m = 1` the grid is a single row. Setting the first row free and the last column free simply leaves the whole row open, which is correct because the only possible moves are rightward. The same idea works when `n = 1`: the single column becomes completely free and the only moves are downward.

In every language I convert the internal grid representation into the required return type (vector of strings, array of strings, list of strings, etc.) before returning it. No extra searching or path counting is performed.

## Examples

**Example 1**

Input: `m = 2`, `n = 3`

One possible output:

```
..#
#..
```

Trace:

- Top row is free.
- Last column is free.
- The only possible sequence is `(0,0) -> (0,1) -> (0,2) -> (1,2)`.

**Example 2**

Input: `m = 3`, `n = 3`

One possible output:

```
..#
#.#
##.
```

(The exact placement of obstacles can vary as long as only one path remains. The construction used in the solutions produces a clean top-row + last-column path.)

**Example 3**

Input: `m = 1`, `n = 4`

Output:

```
....
```

Trace:

- Single row is completely free.
- The only possible path is moving right four times.

## How to Use / Run Locally

**C++**

```bash
g++ -std=c++17 solution.cpp -o solution
./solution
```

**Java**

```bash
javac Solution.java
java Solution
```

**JavaScript**

```bash
node solution.js
```

**TypeScript**

```bash
tsc solution.ts
node solution.js
```

**Python3**

```bash
python3 solution.py
```

**Go**

```bash
go run solution.go
```

Replace the empty function bodies with the code from the Multi-language Solutions section before compiling or running.

## Notes & Optimizations

- The construction is optimal for the given constraints (`m` and `n` up to 25). There is no need for more complex path-generation techniques.
- Any single simple path can be used. Going right then down is just one of the easiest to implement and verify.
- Edge cases (`m = 1` or `n = 1`) are handled naturally by the same logic.
- Alternative approaches exist (for example, a staircase path), but they require more code and do not improve time or space complexity.

## Author

[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)
