# Problem Title

**1970. Last Day Where You Can Still Cross**

---

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

You are given a grid with `row` rows and `col` columns.
Initially, all cells are land.

Each day, one cell becomes water based on the given `cells` array.
You can walk only on land cells and move in four directions.

You can start from **any cell in the top row** and want to reach **any cell in the bottom row**.

The task is to find the **last day** when it is still possible to cross from top to bottom using only land cells.

---

## Constraints

* 2 ≤ row, col ≤ 2 × 10⁴
* row × col ≤ 2 × 10⁴
* cells.length == row × col
* All cell positions are unique
* Cells are 1-based indexed in input

---

## Intuition

At first, I thought about simulating the flooding day by day.

But then I realized something important.

As days increase, crossing becomes harder.
Once crossing is not possible, it will never become possible again.

Instead of simulating forward, I reversed the thinking.

I imagined the grid is **fully water** at the end.
Then I started **adding land back day by day in reverse order**.

The moment a land path connects the **top row to the bottom row**, that day is the answer.

To detect connectivity efficiently, I used **Union Find (Disjoint Set Union)**.

---

## Approach

1. Treat each cell as a node
2. Process days from last to first
3. Turn water into land for the current day
4. Union the current land cell with its neighboring land cells
5. Use two virtual nodes:

   * One for the top row
   * One for the bottom row
6. If top and bottom become connected, return that day

This avoids repeated BFS and works efficiently for large inputs.

---

## Data Structures Used

* Disjoint Set Union (Union Find)
* Parent array
* Rank array
* Boolean grid to track land cells

---

## Operations & Behavior Summary

* Reverse traversal of days
* Union adjacent land cells
* Path compression for fast lookup
* Rank optimization for balanced unions
* Early stop when top and bottom connect

---

## Complexity

**Time Complexity**
O(row × col × α(n))
α(n) is inverse Ackermann function, almost constant

**Space Complexity**
O(row × col)
For parent, rank, and grid storage

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int latestDayToCross(int row, int col, vector<vector<int>>& cells) {
        int n = row * col;
        int top = n, bottom = n + 1;

        vector<int> parent(n + 2), rank(n + 2, 0);
        vector<vector<bool>> grid(row, vector<bool>(col, false));

        for (int i = 0; i < n + 2; i++) parent[i] = i;

        function<int(int)> find = [&](int x) {
            if (parent[x] != x)
                parent[x] = find(parent[x]);
            return parent[x];
        };

        auto unite = [&](int a, int b) {
            a = find(a);
            b = find(b);
            if (a == b) return;
            if (rank[a] < rank[b]) swap(a, b);
            parent[b] = a;
            if (rank[a] == rank[b]) rank[a]++;
        };

        int dr[4] = {1, -1, 0, 0};
        int dc[4] = {0, 0, 1, -1};

        for (int d = n - 1; d >= 0; d--) {
            int r = cells[d][0] - 1;
            int c = cells[d][1] - 1;
            grid[r][c] = true;
            int id = r * col + c;

            if (r == 0) unite(id, top);
            if (r == row - 1) unite(id, bottom);

            for (int k = 0; k < 4; k++) {
                int nr = r + dr[k];
                int nc = c + dc[k];
                if (nr >= 0 && nr < row && nc >= 0 && nc < col && grid[nr][nc]) {
                    unite(id, nr * col + nc);
                }
            }

            if (find(top) == find(bottom)) return d;
        }
        return 0;
    }
};
```

---

### Java

```java
class Solution {
    public int latestDayToCross(int row, int col, int[][] cells) {
        int n = row * col;
        int top = n, bottom = n + 1;

        int[] parent = new int[n + 2];
        int[] rank = new int[n + 2];
        boolean[][] grid = new boolean[row][col];

        for (int i = 0; i < n + 2; i++) parent[i] = i;

        int[] dr = {1, -1, 0, 0};
        int[] dc = {0, 0, 1, -1};

        for (int d = n - 1; d >= 0; d--) {
            int r = cells[d][0] - 1;
            int c = cells[d][1] - 1;
            grid[r][c] = true;
            int id = r * col + c;

            if (r == 0) union(id, top, parent, rank);
            if (r == row - 1) union(id, bottom, parent, rank);

            for (int k = 0; k < 4; k++) {
                int nr = r + dr[k];
                int nc = c + dc[k];
                if (nr >= 0 && nr < row && nc >= 0 && nc < col && grid[nr][nc]) {
                    union(id, nr * col + nc, parent, rank);
                }
            }

            if (find(top, parent) == find(bottom, parent)) return d;
        }
        return 0;
    }

    private int find(int x, int[] parent) {
        if (parent[x] != x)
            parent[x] = find(parent[x], parent);
        return parent[x];
    }

    private void union(int a, int b, int[] parent, int[] rank) {
        a = find(a, parent);
        b = find(b, parent);
        if (a == b) return;
        if (rank[a] < rank[b]) parent[a] = b;
        else {
            parent[b] = a;
            if (rank[a] == rank[b]) rank[a]++;
        }
    }
}
```

---

### JavaScript

```javascript
var latestDayToCross = function(row, col, cells) {
    const n = row * col;
    const top = n, bottom = n + 1;

    const parent = Array(n + 2).fill(0).map((_, i) => i);
    const rank = Array(n + 2).fill(0);
    const grid = Array.from({ length: row }, () => Array(col).fill(false));

    const find = (x) => {
        if (parent[x] !== x) parent[x] = find(parent[x]);
        return parent[x];
    };

    const union = (a, b) => {
        a = find(a);
        b = find(b);
        if (a === b) return;
        if (rank[a] < rank[b]) parent[a] = b;
        else {
            parent[b] = a;
            if (rank[a] === rank[b]) rank[a]++;
        }
    };

    const dr = [1, -1, 0, 0];
    const dc = [0, 0, 1, -1];

    for (let d = n - 1; d >= 0; d--) {
        const r = cells[d][0] - 1;
        const c = cells[d][1] - 1;
        grid[r][c] = true;
        const id = r * col + c;

        if (r === 0) union(id, top);
        if (r === row - 1) union(id, bottom);

        for (let k = 0; k < 4; k++) {
            const nr = r + dr[k];
            const nc = c + dc[k];
            if (nr >= 0 && nr < row && nc >= 0 && nc < col && grid[nr][nc]) {
                union(id, nr * col + nc);
            }
        }

        if (find(top) === find(bottom)) return d;
    }
    return 0;
};
```

---

### Python3

```python
class Solution:
    def latestDayToCross(self, row: int, col: int, cells: List[List[int]]) -> int:
        n = row * col
        top, bottom = n, n + 1

        parent = list(range(n + 2))
        rank = [0] * (n + 2)
        grid = [[False] * col for _ in range(row)]

        def find(x):
            if parent[x] != x:
                parent[x] = find(parent[x])
            return parent[x]

        def union(a, b):
            a, b = find(a), find(b)
            if a == b:
                return
            if rank[a] < rank[b]:
                parent[a] = b
            else:
                parent[b] = a
                if rank[a] == rank[b]:
                    rank[a] += 1

        dr = [1, -1, 0, 0]
        dc = [0, 0, 1, -1]

        for d in range(n - 1, -1, -1):
            r, c = cells[d][0] - 1, cells[d][1] - 1
            grid[r][c] = True
            idx = r * col + c

            if r == 0:
                union(idx, top)
            if r == row - 1:
                union(idx, bottom)

            for k in range(4):
                nr, nc = r + dr[k], c + dc[k]
                if 0 <= nr < row and 0 <= nc < col and grid[nr][nc]:
                    union(idx, nr * col + nc)

            if find(top) == find(bottom):
                return d
        return 0
```

---

### Go

```go
func latestDayToCross(row int, col int, cells [][]int) int {
 n := row * col
 top, bottom := n, n+1

 parent := make([]int, n+2)
 rank := make([]int, n+2)
 grid := make([]bool, n)

 for i := 0; i < n+2; i++ {
  parent[i] = i
 }

 var find func(int) int
 find = func(x int) int {
  if parent[x] != x {
   parent[x] = find(parent[x])
  }
  return parent[x]
 }

 union := func(a, b int) {
  a = find(a)
  b = find(b)
  if a == b {
   return
  }
  if rank[a] < rank[b] {
   parent[a] = b
  } else {
   parent[b] = a
   if rank[a] == rank[b] {
    rank[a]++
   }
  }
 }

 dr := []int{1, -1, 0, 0}
 dc := []int{0, 0, 1, -1}

 for d := n - 1; d >= 0; d-- {
  r := cells[d][0] - 1
  c := cells[d][1] - 1
  id := r*col + c
  grid[id] = true

  if r == 0 {
   union(id, top)
  }
  if r == row-1 {
   union(id, bottom)
  }

  for k := 0; k < 4; k++ {
   nr := r + dr[k]
   nc := c + dc[k]
   if nr >= 0 && nr < row && nc >= 0 && nc < col {
    nid := nr*col + nc
    if grid[nid] {
     union(id, nid)
    }
   }
  }

  if find(top) == find(bottom) {
   return d
  }
 }
 return 0
}
```

---

## Step-by-step Detailed Explanation

1. Treat each grid cell as a node
2. Reverse the flooding process
3. Add land one cell at a time
4. Union adjacent land cells
5. Connect top row and bottom row to virtual nodes
6. Stop when both virtual nodes are connected

---

## Examples

**Input**

```
row = 2, col = 2
cells = [[1,1],[2,1],[1,2],[2,2]]
```

**Output**

```
2
```

---

## How to use / Run locally

1. Clone the repository
2. Choose your preferred language folder
3. Compile or run the file using standard compiler or interpreter
4. Provide input through test cases or main function

---

## Notes & Optimizations

* Union Find is faster than BFS for repeated connectivity checks
* Reverse simulation avoids unnecessary recalculations
* Path compression improves performance significantly
* This approach is optimal for large grids

---

## Author

* **Md Aarzoo Islam**
  [https://bento.me/withaarzoo](https://bento.me/withaarzoo)
