# LeetCode 2267: Check if There Is a Valid Parentheses String Path - Dynamic Programming Solution

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

You are given an m x n grid filled only with the characters '(' and ')'. Your task is to check whether there exists at least one path that starts at the top-left cell (0, 0), ends at the bottom-right cell (m-1, n-1), and moves only right or down.  

The string formed by the characters along that path must be a valid parentheses string. A valid parentheses string is one that is empty, or can be built by concatenating two valid strings, or by wrapping a valid string inside a pair of matching parentheses.  

Return true if such a valid parentheses string path exists in the grid. Otherwise return false.  

This is a classic grid path problem combined with the classic balance-checking rule for parentheses, solved efficiently using dynamic programming.

## Constraints

- m == grid.length
- n == grid[i].length
- 1 <= m, n <= 100
- grid[i][j] is either '(' or ')'

## Intuition

The first thing I noticed is that every path from the top-left corner to the bottom-right corner has the exact same length: m + n - 1. A valid parentheses string must have even length, so if m + n - 1 is odd I can immediately return false.  

Next I remembered the standard balance counter for parentheses: add 1 for an opening parenthesis and subtract 1 for a closing one. The balance must never go negative and must finish at zero.  

Because we can only move right or down, the problem becomes a reachability question on a grid with an extra state for the current balance. With m and n at most 100 the longest path is under 200 steps, so the balance stays in a small range and a three-dimensional DP table fits comfortably in memory.

## Approach

I create a three-dimensional boolean DP table. The entry dp[i][j][bal] is true if I can reach cell (i, j) with the given balance.  

I start by processing the top-left cell. If it is a closing parenthesis the balance becomes negative and the answer is immediately false. Otherwise I mark the corresponding balance as reachable.  

Then I iterate over every cell from top to bottom and left to right. For each reachable balance at the current cell I try the two possible moves (down and right). I compute the new balance after reading the character in the next cell. If the new balance stays non-negative I mark that state as reachable.  

After filling the table I simply look at whether the bottom-right cell can be reached with balance 0. That single check tells me if a valid parentheses string path exists.

## Data Structures Used

- Three-dimensional boolean array (or vector of vectors of vectors) of size m x n x (m + n).  
  This stores reachability for every cell and every possible balance. I chose a boolean array because I only need to know whether a state is reachable, not how many ways it can be reached. The third dimension is sized to the maximum possible path length so every legal balance fits.

## Operations & Behavior Summary

1. Read the grid dimensions and check whether the path length is odd. If it is, return false at once.  
2. Allocate the DP table and initialise every entry to false.  
3. Process the starting cell: compute its balance contribution and mark the corresponding DP entry true (or return false if the balance is already negative).  
4. For every cell in row-major order:  
   - Look at every balance that is already marked reachable.  
   - Try moving down (if a lower row exists).  
   - Try moving right (if a right column exists).  
   - For each move compute the new balance and mark it reachable only if it stays non-negative.  
5. After the loops finish, return the value stored in the bottom-right cell for balance 0.

## Complexity

| Complexity       | Value                  | Explanation                                                                 |
|------------------|------------------------|-----------------------------------------------------------------------------|
| Time Complexity  | O(m * n * (m + n))     | We visit every cell and every possible balance (at most m + n).             |
| Space Complexity | O(m * n * (m + n))     | The three-dimensional DP table stores one boolean for each cell-balance pair. |

## Multi-language Solutions

### C++
```cpp
class Solution {
public:
    bool hasValidPath(vector<vector<char>>& grid) {
        int m = grid.size(), n = grid[0].size();
        if ((m + n - 1) % 2 != 0) return false;
        int maxBal = m + n;
        vector<vector<vector<bool>>> dp(m, vector<vector<bool>>(n, vector<bool>(maxBal, false)));
        int start = grid[0][0] == '(' ? 1 : -1;
        if (start < 0) return false;
        dp[0][0][start] = true;
        for (int i = 0; i < m; ++i) {
            for (int j = 0; j < n; ++j) {
                for (int bal = 0; bal < maxBal; ++bal) {
                    if (!dp[i][j][bal]) continue;
                    if (i + 1 < m) {
                        int nb = bal + (grid[i + 1][j] == '(' ? 1 : -1);
                        if (nb >= 0 && nb < maxBal) dp[i + 1][j][nb] = true;
                    }
                    if (j + 1 < n) {
                        int nb = bal + (grid[i][j + 1] == '(' ? 1 : -1);
                        if (nb >= 0 && nb < maxBal) dp[i][j + 1][nb] = true;
                    }
                }
            }
        }
        return dp[m - 1][n - 1][0];
    }
};
```

### Java
```java
class Solution {
    public boolean hasValidPath(char[][] grid) {
        int m = grid.length, n = grid[0].length;
        if ((m + n - 1) % 2 != 0) return false;
        int maxBal = m + n;
        boolean[][][] dp = new boolean[m][n][maxBal];
        int start = grid[0][0] == '(' ? 1 : -1;
        if (start < 0) return false;
        dp[0][0][start] = true;
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                for (int bal = 0; bal < maxBal; bal++) {
                    if (!dp[i][j][bal]) continue;
                    if (i + 1 < m) {
                        int nb = bal + (grid[i + 1][j] == '(' ? 1 : -1);
                        if (nb >= 0 && nb < maxBal) dp[i + 1][j][nb] = true;
                    }
                    if (j + 1 < n) {
                        int nb = bal + (grid[i][j + 1] == '(' ? 1 : -1);
                        if (nb >= 0 && nb < maxBal) dp[i][j + 1][nb] = true;
                    }
                }
            }
        }
        return dp[m - 1][n - 1][0];
    }
}
```

### JavaScript
```javascript
var hasValidPath = function(grid) {
    const m = grid.length, n = grid[0].length;
    if ((m + n - 1) % 2 !== 0) return false;
    const maxBal = m + n;
    const dp = Array.from({length: m}, () => Array.from({length: n}, () => Array(maxBal).fill(false)));
    const start = grid[0][0] === '(' ? 1 : -1;
    if (start < 0) return false;
    dp[0][0][start] = true;
    for (let i = 0; i < m; i++) {
        for (let j = 0; j < n; j++) {
            for (let bal = 0; bal < maxBal; bal++) {
                if (!dp[i][j][bal]) continue;
                if (i + 1 < m) {
                    const nb = bal + (grid[i + 1][j] === '(' ? 1 : -1);
                    if (nb >= 0 && nb < maxBal) dp[i + 1][j][nb] = true;
                }
                if (j + 1 < n) {
                    const nb = bal + (grid[i][j + 1] === '(' ? 1 : -1);
                    if (nb >= 0 && nb < maxBal) dp[i][j + 1][nb] = true;
                }
            }
        }
    }
    return dp[m - 1][n - 1][0];
};
```

### TypeScript
```typescript
var hasValidPath = function(grid: string[][]): boolean {
    const m = grid.length, n = grid[0].length;
    if ((m + n - 1) % 2 !== 0) return false;
    const maxBal = m + n;
    const dp: boolean[][][] = Array.from({length: m}, () => Array.from({length: n}, () => Array(maxBal).fill(false)));
    const start = grid[0][0] === '(' ? 1 : -1;
    if (start < 0) return false;
    dp[0][0][start] = true;
    for (let i = 0; i < m; i++) {
        for (let j = 0; j < n; j++) {
            for (let bal = 0; bal < maxBal; bal++) {
                if (!dp[i][j][bal]) continue;
                if (i + 1 < m) {
                    const nb = bal + (grid[i + 1][j] === '(' ? 1 : -1);
                    if (nb >= 0 && nb < maxBal) dp[i + 1][j][nb] = true;
                }
                if (j + 1 < n) {
                    const nb = bal + (grid[i][j + 1] === '(' ? 1 : -1);
                    if (nb >= 0 && nb < maxBal) dp[i][j + 1][nb] = true;
                }
            }
        }
    }
    return dp[m - 1][n - 1][0];
};
```

### Python3
```python
class Solution:
    def hasValidPath(self, grid: list[list[str]]) -> bool:
        m, n = len(grid), len(grid[0])
        if (m + n - 1) % 2 != 0:
            return False
        maxBal = m + n
        dp = [[[False] * maxBal for _ in range(n)] for _ in range(m)]
        start = 1 if grid[0][0] == '(' else -1
        if start < 0:
            return False
        dp[0][0][start] = True
        for i in range(m):
            for j in range(n):
                for bal in range(maxBal):
                    if not dp[i][j][bal]:
                        continue
                    if i + 1 < m:
                        nb = bal + (1 if grid[i + 1][j] == '(' else -1)
                        if 0 <= nb < maxBal:
                            dp[i + 1][j][nb] = True
                    if j + 1 < n:
                        nb = bal + (1 if grid[i][j + 1] == '(' else -1)
                        if 0 <= nb < maxBal:
                            dp[i][j + 1][nb] = True
        return dp[m - 1][n - 1][0]
```

### Go
```go
func hasValidPath(grid [][]byte) bool {
    m, n := len(grid), len(grid[0])
    if (m+n-1)%2 != 0 {
        return false
    }
    maxBal := m + n
    dp := make([][][]bool, m)
    for i := range dp {
        dp[i] = make([][]bool, n)
        for j := range dp[i] {
            dp[i][j] = make([]bool, maxBal)
        }
    }
    start := 1
    if grid[0][0] == ')' {
        start = -1
    }
    if start < 0 {
        return false
    }
    dp[0][0][start] = true
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            for bal := 0; bal < maxBal; bal++ {
                if !dp[i][j][bal] {
                    continue
                }
                if i+1 < m {
                    nb := bal
                    if grid[i+1][j] == '(' {
                        nb++
                    } else {
                        nb--
                    }
                    if nb >= 0 && nb < maxBal {
                        dp[i+1][j][nb] = true
                    }
                }
                if j+1 < n {
                    nb := bal
                    if grid[i][j+1] == '(' {
                        nb++
                    } else {
                        nb--
                    }
                    if nb >= 0 && nb < maxBal {
                        dp[i][j+1][nb] = true
                    }
                }
            }
        }
    }
    return dp[m-1][n-1][0]
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

The logic is identical across all six languages; only the syntax changes.  

First I read m and n. I immediately check whether m + n - 1 is odd. An odd-length string can never be a valid parentheses sequence, so returning false early saves unnecessary work.  

I allocate a three-dimensional boolean structure whose third dimension is large enough to hold every balance that can appear on a path of length m + n - 1.  

I examine the character at (0, 0). If it is ')', the balance becomes -1 and I return false. Otherwise the balance becomes 1 and I set the corresponding DP entry to true.  

The main triple loop walks through every cell in the order that guarantees predecessors are already processed. For each reachable balance I attempt the two legal moves. After reading the character in the destination cell I compute the new balance. Only non-negative balances are written into the DP table; this is exactly the same rule used when validating a parentheses string with a stack or a counter.  

When the loops finish I look at the single entry that corresponds to the bottom-right cell and balance 0. If that entry is true, at least one valid path exists.  

Edge cases such as a 1 x 1 grid, a grid that starts with a closing parenthesis, or a grid whose total length is odd are all handled by the same logic without special extra branches.

## Examples

**Example 1**  
Input:  
grid = [["(","(","("],[")","(",")"],["(","(",")"],["(","(",")"]]  

Output: true  

Trace:  
One possible path produces the string "()(())". The balance never goes negative and ends at zero, so the DP marks the final cell with balance 0 as reachable.

**Example 2**  
Input:  
grid = [[")",")"],["(","("]]  

Output: false  

Trace:  
Both possible paths produce the strings "))( (" and ")((". In each case the balance becomes negative at some point, so the DP never marks any final state with balance 0.

**Example 3** (edge case)  
Input:  
grid = [["("]]  

Output: false  

Trace:  
Path length is 1 (odd). The early parity check returns false immediately.

## How to Use / Run Locally

1. Copy the code for the language you want into a file (for example `Solution.cpp`, `Solution.java`, etc.).  
2. Make sure the function signature matches the LeetCode template for that language.  
3. Compile and run according to the language:  
   - C++: `g++ -std=c++17 Solution.cpp -o solution && ./solution`  
   - Java: `javac Solution.java && java Solution`  
   - JavaScript: `node Solution.js`  
   - TypeScript: compile with `tsc` then run the generated JavaScript, or use `ts-node`.  
   - Python3: `python3 Solution.py`  
   - Go: `go run Solution.go`  
4. Feed the grid as a two-dimensional array of characters and print the boolean result.

## Notes & Optimizations

- The parity check is a cheap early exit; the DP itself would also return false for odd-length paths, but the check makes the code faster on those inputs.  
- Memory can be reduced by keeping only two layers of the DP (previous row and current row) because we only ever look at the cell above or to the left. The current three-dimensional version is already well within limits for the given constraints.  
- An alternative recursive DFS with memoization on (row, column, balance) works and produces the same asymptotic complexity, but the iterative DP is usually clearer and avoids recursion depth concerns.  
- If the problem asked for the number of valid paths instead of existence, the same table could store counts instead of booleans.

## Author
[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)