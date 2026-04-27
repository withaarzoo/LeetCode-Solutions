# 1391. Check if There is a Valid Path in a Grid

## Table of Contents

* Problem Summary
* Constraints
* Intuition
* Approach
* Data Structures Used
* Operations & Behavior Summary
* Complexity
* Multi-language Solutions
* Step-by-step Detailed Explanation
* Examples
* How to use / Run locally
* Notes & Optimizations
* Author

---

## Problem Summary

I am given an `m x n` grid where each cell represents a type of street. Each street connects certain directions (left, right, up, down).

My goal is to check whether I can travel from the top-left cell `(0,0)` to the bottom-right cell `(m-1,n-1)` by only following valid street connections.

---

## Constraints

* `m == grid.length`
* `n == grid[i].length`
* `1 <= m, n <= 300`
* `1 <= grid[i][j] <= 6`

---

## Intuition

I thought of this problem like a graph traversal.

Each cell is a node, and I can move only if:

1. The current street allows movement in that direction
2. The next street allows entry from the opposite direction

So I decided to explore the grid using BFS.

---

## Approach

1. Define all 4 directions (left, right, up, down)
2. Map each street type to valid directions
3. Start BFS from `(0,0)`
4. For every move:

   * Check boundaries
   * Check if next cell is already visited
   * Check if next street connects back
5. If I reach `(m-1,n-1)`, return true
6. Otherwise, return false

---

## Data Structures Used

* Queue (for BFS)
* Visited matrix
* Direction arrays
* Mapping for street connections

---

## Operations & Behavior Summary

* Traverse grid like graph
* Validate bidirectional street connection
* Avoid revisiting nodes
* Stop early if destination reached

---

## Complexity

* Time Complexity: O(m * n)

  * Each cell is visited once

* Space Complexity: O(m * n)

  * For visited array and queue

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    bool hasValidPath(vector<vector<int>>& grid) {
        int m = grid.size(), n = grid[0].size();

        vector<pair<int,int>> dirs = {{0,-1},{0,1},{-1,0},{1,0}};

        vector<vector<int>> streetDirs = {
            {}, {0,1}, {2,3}, {0,3}, {1,3}, {0,2}, {1,2}
        };

        vector<int> opposite = {1,0,3,2};

        vector<vector<bool>> vis(m, vector<bool>(n,false));
        queue<pair<int,int>> q;
        q.push({0,0});
        vis[0][0] = true;

        while(!q.empty()){
            auto [r,c] = q.front(); q.pop();

            if(r==m-1 && c==n-1) return true;

            for(int d: streetDirs[grid[r][c]]){
                int nr = r + dirs[d].first;
                int nc = c + dirs[d].second;

                if(nr<0||nc<0||nr>=m||nc>=n||vis[nr][nc]) continue;

                for(int nd: streetDirs[grid[nr][nc]]){
                    if(nd == opposite[d]){
                        vis[nr][nc]=true;
                        q.push({nr,nc});
                    }
                }
            }
        }
        return false;
    }
};
```

### Java

```java
class Solution {
    public boolean hasValidPath(int[][] grid) {
        int m = grid.length, n = grid[0].length;

        int[][] dirs = {{0,-1},{0,1},{-1,0},{1,0}};

        int[][] streetDirs = {
            {}, {0,1}, {2,3}, {0,3}, {1,3}, {0,2}, {1,2}
        };

        int[] opposite = {1,0,3,2};

        boolean[][] vis = new boolean[m][n];
        Queue<int[]> q = new LinkedList<>();
        q.offer(new int[]{0,0});
        vis[0][0] = true;

        while(!q.isEmpty()){
            int[] cur = q.poll();
            int r = cur[0], c = cur[1];

            if(r==m-1 && c==n-1) return true;

            for(int d: streetDirs[grid[r][c]]){
                int nr = r + dirs[d][0];
                int nc = c + dirs[d][1];

                if(nr<0||nc<0||nr>=m||nc>=n||vis[nr][nc]) continue;

                for(int nd: streetDirs[grid[nr][nc]]){
                    if(nd == opposite[d]){
                        vis[nr][nc]=true;
                        q.offer(new int[]{nr,nc});
                    }
                }
            }
        }
        return false;
    }
}
```

### JavaScript

```javascript
var hasValidPath = function(grid) {
    const m = grid.length, n = grid[0].length;

    const dirs = [[0,-1],[0,1],[-1,0],[1,0]];

    const streetDirs = [
        [], [0,1], [2,3], [0,3], [1,3], [0,2], [1,2]
    ];

    const opposite = [1,0,3,2];

    const vis = Array.from({length:m},()=>Array(n).fill(false));
    const q = [[0,0]];
    vis[0][0] = true;

    while(q.length){
        const [r,c] = q.shift();

        if(r===m-1 && c===n-1) return true;

        for(const d of streetDirs[grid[r][c]]){
            const nr = r + dirs[d][0];
            const nc = c + dirs[d][1];

            if(nr<0||nc<0||nr>=m||nc>=n||vis[nr][nc]) continue;

            for(const nd of streetDirs[grid[nr][nc]]){
                if(nd===opposite[d]){
                    vis[nr][nc]=true;
                    q.push([nr,nc]);
                }
            }
        }
    }
    return false;
};
```

### Python3

```python
from collections import deque

class Solution:
    def hasValidPath(self, grid):
        m, n = len(grid), len(grid[0])

        dirs = [(0,-1),(0,1),(-1,0),(1,0)]

        street_dirs = {
            1:[0,1], 2:[2,3], 3:[0,3],
            4:[1,3], 5:[0,2], 6:[1,2]
        }

        opposite = {0:1,1:0,2:3,3:2}

        vis = [[False]*n for _ in range(m)]
        q = deque([(0,0)])
        vis[0][0] = True

        while q:
            r,c = q.popleft()

            if r==m-1 and c==n-1:
                return True

            for d in street_dirs[grid[r][c]]:
                nr, nc = r+dirs[d][0], c+dirs[d][1]

                if nr<0 or nc<0 or nr>=m or nc>=n or vis[nr][nc]:
                    continue

                if opposite[d] in street_dirs[grid[nr][nc]]:
                    vis[nr][nc]=True
                    q.append((nr,nc))

        return False
```

### Go

```go
func hasValidPath(grid [][]int) bool {
    m, n := len(grid), len(grid[0])

    dirs := [4][2]int{{0,-1},{0,1},{-1,0},{1,0}}

    streetDirs := map[int][]int{
        1:{0,1},2:{2,3},3:{0,3},
        4:{1,3},5:{0,2},6:{1,2},
    }

    opposite := [4]int{1,0,3,2}

    vis := make([][]bool,m)
    for i:=range vis{
        vis[i]=make([]bool,n)
    }

    queue := [][2]int{{0,0}}
    vis[0][0]=true

    for len(queue)>0{
        r,c := queue[0][0], queue[0][1]
        queue = queue[1:]

        if r==m-1 && c==n-1 {return true}

        for _,d := range streetDirs[grid[r][c]]{
            nr, nc := r+dirs[d][0], c+dirs[d][1]

            if nr<0||nc<0||nr>=m||nc>=n||vis[nr][nc]{continue}

            for _,nd := range streetDirs[grid[nr][nc]]{
                if nd==opposite[d]{
                    vis[nr][nc]=true
                    queue = append(queue,[2]int{nr,nc})
                }
            }
        }
    }
    return false
}
```

---

## Step-by-step Detailed Explanation

I start BFS from (0,0).

For each cell:

* I check allowed directions from current street
* Then I calculate next cell
* I verify bounds and visited condition
* I check reverse connection validity
* If valid, I push it into queue

I repeat until:

* I reach destination → return true
* Queue becomes empty → return false

---

## Examples

Input: grid = [[2,4,3],[6,5,2]]
Output: true

Input: grid = [[1,2,1],[1,2,1]]
Output: false

---

## How to use / Run locally

1. Copy code into your local IDE
2. Choose language (C++ / Java / Python etc.)
3. Provide input grid
4. Run the function

---

## Notes & Optimizations

* BFS is better than DFS for early stopping
* Avoid repeated visits using visited array
* Constant direction checks → efficient

---

## Author

* [Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
