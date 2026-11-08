# Count Servers that Communicate - LeetCode 1267 Solution

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

You are given a 2D grid that represents a server center. Each cell is either 1 (there is a server) or 0 (empty). Two servers can communicate if they share the same row or the same column.

The task is to count how many servers can talk to at least one other server. Isolated servers that sit alone in both their row and column are ignored.

Input is an m by n integer matrix called grid. Output is a single integer — the number of servers that are not isolated.

This is a classic matrix counting problem that appears often in coding interviews and on LeetCode as problem 1267.

## Constraints

- m == grid.length
- n == grid[i].length
- 1 <= m <= 250
- 1 <= n <= 250
- grid[i][j] is either 0 or 1

## Intuition

The first thing I noticed is that a server is useful only when it is not alone. If a row has two or more servers, every server in that row can communicate. The same rule applies to columns.

Instead of checking every pair of servers (which would be slow), I can simply count how many servers live in each row and each column. Once I have those two numbers, I can walk the grid again and keep only the servers whose row count or column count is greater than 1.

That observation turns a potentially expensive pairwise check into two clean linear passes over the matrix.

## Approach

I create two simple arrays: one that stores the number of servers in every row, and another that stores the number of servers in every column.

First I scan the entire grid once. Whenever I see a 1, I increase the matching row counter and the matching column counter.

After the first scan I know the exact population of every row and every column.

I scan the grid a second time. For every cell that contains a server I look at its row population and its column population. If either number is bigger than 1, the server is not isolated, so I add one to the answer.

At the end the answer is exactly the number of servers that can communicate with at least one other server.

This approach works for the full range of constraints (up to 250 by 250) and stays easy to implement in any language.

## Data Structures Used

- Two integer arrays (or vectors/slices/lists) of size m and n.  
  These hold the server counts for each row and each column. They are the only extra space needed and make the second pass O(1) per cell.

No maps, sets, or graphs are required. The plain count arrays are enough because communication only depends on sharing a row or a column.

## Operations & Behavior Summary

1. Read the dimensions m and n of the grid.  
2. Create rowCount of length m and colCount of length n, both starting at zero.  
3. Walk every cell. When a server is found, increment rowCount[i] and colCount[j].  
4. Initialize answer = 0.  
5. Walk every cell again. If the cell has a server and (rowCount[i] > 1 or colCount[j] > 1), increase the answer by one.  
6. Return the answer.

That is the complete logic. Everything else is just language-specific syntax.

## Complexity

| Type              | Complexity | Explanation |
|-------------------|------------|-------------|
| Time Complexity   | O(m * n)   | Two full passes over the grid. Each cell is visited a constant number of times. |
| Space Complexity  | O(m + n)   | Only the two count arrays are allocated. No other extra memory is used. |

## Multi-language Solutions

### C++
```cpp
class Solution {
public:
    int countServers(vector<vector<int>>& grid) {
        int m = grid.size();
        int n = grid[0].size();
        vector<int> row(m, 0);
        vector<int> col(n, 0);
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (grid[i][j] == 1) {
                    row[i]++;
                    col[j]++;
                }
            }
        }
        int ans = 0;
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (grid[i][j] == 1 && (row[i] > 1 || col[j] > 1)) {
                    ans++;
                }
            }
        }
        return ans;
    }
};
```

### Java
```java
class Solution {
    public int countServers(int[][] grid) {
        int m = grid.length;
        int n = grid[0].length;
        int[] row = new int[m];
        int[] col = new int[n];
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (grid[i][j] == 1) {
                    row[i]++;
                    col[j]++;
                }
            }
        }
        int ans = 0;
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (grid[i][j] == 1 && (row[i] > 1 || col[j] > 1)) {
                    ans++;
                }
            }
        }
        return ans;
    }
}
```

### JavaScript
```javascript
/**
 * @param {number[][]} grid
 * @return {number}
 */
var countServers = function(grid) {
    let m = grid.length;
    let n = grid[0].length;
    let row = new Array(m).fill(0);
    let col = new Array(n).fill(0);
    for (let i = 0; i < m; i++) {
        for (let j = 0; j < n; j++) {
            if (grid[i][j] === 1) {
                row[i]++;
                col[j]++;
            }
        }
    }
    let ans = 0;
    for (let i = 0; i < m; i++) {
        for (let j = 0; j < n; j++) {
            if (grid[i][j] === 1 && (row[i] > 1 || col[j] > 1)) {
                ans++;
            }
        }
    }
    return ans;
};
```

### TypeScript
```typescript
function countServers(grid: number[][]): number {
    let m = grid.length;
    let n = grid[0].length;
    let row = new Array(m).fill(0);
    let col = new Array(n).fill(0);
    for (let i = 0; i < m; i++) {
        for (let j = 0; j < n; j++) {
            if (grid[i][j] === 1) {
                row[i]++;
                col[j]++;
            }
        }
    }
    let ans = 0;
    for (let i = 0; i < m; i++) {
        for (let j = 0; j < n; j++) {
            if (grid[i][j] === 1 && (row[i] > 1 || col[j] > 1)) {
                ans++;
            }
        }
    }
    return ans;
};
```

### Python3
```python
class Solution:
    def countServers(self, grid: list[list[int]]) -> int:
        m = len(grid)
        n = len(grid[0])
        row = [0] * m
        col = [0] * n
        for i in range(m):
            for j in range(n):
                if grid[i][j] == 1:
                    row[i] += 1
                    col[j] += 1
        ans = 0
        for i in range(m):
            for j in range(n):
                if grid[i][j] == 1 and (row[i] > 1 or col[j] > 1):
                    ans += 1
        return ans
```

### Go
```go
func countServers(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    row := make([]int, m)
    col := make([]int, n)
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                row[i]++
                col[j]++
            }
        }
    }
    ans := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 && (row[i] > 1 || col[j] > 1) {
                ans++
            }
        }
    }
    return ans
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

The logic is identical in every language, so the reasoning stays the same.

I start by reading the number of rows and columns. This gives me the exact sizes I need for the two count arrays.

I allocate the row count array of size m and the column count array of size n. Both are filled with zeros at the beginning.

In the first double loop I look at every cell. The moment I see a 1 I increase the corresponding entry in the row array and the corresponding entry in the column array. After this pass both arrays contain the true population of every row and every column.

I set the answer variable to zero. This variable will hold the final result.

In the second double loop I again look at every cell. When the cell contains a server I check two simple conditions: is the row population greater than 1, or is the column population greater than 1? If either is true the server can communicate, so I increase the answer.

I never need to look at other cells during the second pass because the counts already tell me whether company exists.

Edge cases are handled naturally. A completely empty grid returns 0. A grid with only one server returns 0. A full row or full column of servers correctly counts every server in that line.

Because the two languages differ only in syntax (vector vs array, list vs slice, etc.), the same reasoning applies to C++, Java, JavaScript, TypeScript, Python, and Go.

## Examples

**Example 1**  
Input: grid = [[1,0],[0,1]]  
Output: 0  

The first server sits alone in row 0 and column 0. The second server sits alone in row 1 and column 1. Both are isolated, so the answer is 0.

**Example 2**  
Input: grid = [[1,0],[1,1]]  
Output: 3  

Row 0 has one server, row 1 has two servers, column 0 has two servers, column 1 has one server. Every server shares either its row or its column with another server, so all three are counted.

**Example 3**  
Input: grid = [[1,1,0,0],[0,0,1,0],[0,0,1,0],[0,0,0,1]]  
Output: 4  

The two servers in the first row can talk to each other. The two servers in the third column can talk to each other. The single server at the bottom-right corner is alone in both its row and its column, so it is ignored. Total = 4.

## How to Use / Run Locally

**C++**  
1. Copy the solution into a file named main.cpp.  
2. Compile with: g++ -std=c++17 main.cpp -o main  
3. Run with: ./main  
(You can add a simple main function that builds a test grid and prints the result.)

**Java**  
1. Put the class inside a file named Solution.java.  
2. Compile with: javac Solution.java  
3. Run a small driver class that creates a test grid and calls countServers.

**JavaScript**  
1. Save the function in a file named solution.js.  
2. Run with Node: node solution.js  
(Add a few console.log calls with sample grids to see the output.)

**TypeScript**  
1. Save the function in a file named solution.ts.  
2. Compile with: tsc solution.ts  
3. Run the generated JavaScript file with Node.

**Python3**  
1. Save the class in a file named solution.py.  
2. Run with: python3 solution.py  
(Add a few print statements that call the method on sample grids.)

**Go**  
1. Save the function in a file named main.go.  
2. Run with: go run main.go  
(Wrap the function in a main package and print the results of a few test cases.)

## Notes & Optimizations

The solution already runs in the best possible time for this problem because every cell must be examined at least once.

An alternative approach would be to treat the servers as nodes in a graph and run DFS or Union-Find, but that uses more memory and is slower for no gain.

The only real edge case is a grid that contains zero servers or exactly one server; both correctly return 0.

Because the constraints are small (250 by 250), even a slightly less optimal solution would pass, but the two-pass counting method is clean, fast, and easy to remember for interviews.

## Author
[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)
