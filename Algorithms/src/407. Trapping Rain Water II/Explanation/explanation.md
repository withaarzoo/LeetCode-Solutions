# 407. Trapping Rain Water II

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
  - [Python3](#python3)
  - [Go](#go)

- [Step-by-step Detailed Explanation](#step-by-step-detailed-explanation-c-java-javascript-python3-go)
- [Examples](#examples)
- [How to use / Run locally](#how-to-use--run-locally)
- [Notes & Optimizations](#notes--optimizations)
- [Author](#author)

---

## Problem Summary

We are given an `m x n` integer matrix `heightMap` representing the height of each unit cell in a 2D elevation map.
We need to calculate how much water can be trapped after raining.

This is the **2D version** of the classic "Trapping Rain Water" problem.

---

## Constraints

- `m == heightMap.length`
- `n == heightMap[i].length`
- `1 <= m, n <= 200`
- `0 <= heightMap[i][j] <= 2 * 10^4`

---

## Intuition

When I first thought about the problem, I realized it’s very similar to the 1D trapping rain water problem — but in **two dimensions**.
In 1D, the trapped water depends on the tallest bars to the left and right.
In 2D, the trapped water depends on the **minimum boundary height surrounding each cell**.

So, I thought:

- If we start from the **boundary (edges)** and move inward, we can always guarantee that the boundary sets the maximum water level.
- Using a **priority queue (min-heap)**, we can process the lowest boundary first and expand inward.
- This way, we always “trap water” correctly without missing any lower spots inside.

---

## Approach

1. Push all boundary cells into a **min-heap** (priority queue).
2. Mark them as visited.
3. Repeatedly pop the lowest boundary cell from the heap.
4. Check its unvisited neighbors:

   - If the neighbor is lower, water is trapped = `currentHeight - neighborHeight`.
   - Push the neighbor into the heap with updated height = `max(currentHeight, neighborHeight)`.

5. Continue until all cells are processed.

---

## Data Structures Used

- **Min-Heap / Priority Queue** → Always expand from the lowest boundary.
- **Visited Matrix** → Prevent reprocessing cells.
- **2D Grid Traversal** → Four directions (up, down, left, right).

---

## Operations & Behavior Summary

- Insert boundary cells into heap → O(m\*n)
- Pop from heap and process neighbors → O(m*n log(m*n))
- Each cell processed once.

---

## Complexity

- **Time Complexity:** `O(m * n log(m * n))`

  - Each cell is pushed into heap once → `m*n`
  - Each push/pop → `log(m*n)`

- **Space Complexity:** `O(m * n)`

  - For visited matrix and heap storage.

---

## Multi-language Solutions

### C++

```cpp
#include <vector>
#include <queue>
#include <tuple>
using namespace std;

class Solution {
public:
    int trapRainWater(vector<vector<int>>& heightMap) {
        int m = heightMap.size(), n = heightMap[0].size();
        if (m < 3 || n < 3) return 0;

        priority_queue<tuple<int,int,int>, vector<tuple<int,int,int>>, greater<>> pq;
        vector<vector<bool>> visited(m, vector<bool>(n, false));

        // Push boundary cells
        for (int i = 0; i < m; ++i) {
            pq.emplace(heightMap[i][0], i, 0);
            pq.emplace(heightMap[i][n-1], i, n-1);
            visited[i][0] = visited[i][n-1] = true;
        }
        for (int j = 0; j < n; ++j) {
            pq.emplace(heightMap[0][j], 0, j);
            pq.emplace(heightMap[m-1][j], m-1, j);
            visited[0][j] = visited[m-1][j] = true;
        }

        int result = 0, dirs[4][2] = {{0,1},{1,0},{0,-1},{-1,0}};
        while (!pq.empty()) {
            auto [h, x, y] = pq.top(); pq.pop();
            for (auto &d : dirs) {
                int nx = x + d[0], ny = y + d[1];
                if (nx>=0 && ny>=0 && nx<m && ny<n && !visited[nx][ny]) {
                    result += max(0, h - heightMap[nx][ny]);
                    pq.emplace(max(h, heightMap[nx][ny]), nx, ny);
                    visited[nx][ny] = true;
                }
            }
        }
        return result;
    }
};
```

### Java

```java
import java.util.*;

class Solution {
    public int trapRainWater(int[][] heightMap) {
        int m = heightMap.length, n = heightMap[0].length;
        if (m < 3 || n < 3) return 0;

        PriorityQueue<int[]> pq = new PriorityQueue<>((a,b)->a[0]-b[0]);
        boolean[][] visited = new boolean[m][n];

        for (int i=0; i<m; i++) {
            pq.offer(new int[]{heightMap[i][0], i, 0});
            pq.offer(new int[]{heightMap[i][n-1], i, n-1});
            visited[i][0] = visited[i][n-1] = true;
        }
        for (int j=0; j<n; j++) {
            pq.offer(new int[]{heightMap[0][j], 0, j});
            pq.offer(new int[]{heightMap[m-1][j], m-1, j});
            visited[0][j] = visited[m-1][j] = true;
        }

        int res = 0;
        int[][] dirs = {{0,1},{1,0},{0,-1},{-1,0}};
        while (!pq.isEmpty()) {
            int[] cell = pq.poll();
            int h = cell[0], x = cell[1], y = cell[2];
            for (int[] d : dirs) {
                int nx = x+d[0], ny = y+d[1];
                if (nx>=0 && ny>=0 && nx<m && ny<n && !visited[nx][ny]) {
                    res += Math.max(0, h - heightMap[nx][ny]);
                    pq.offer(new int[]{Math.max(h, heightMap[nx][ny]), nx, ny});
                    visited[nx][ny] = true;
                }
            }
        }
        return res;
    }
}
```

### JavaScript

```javascript
var trapRainWater = function (heightMap) {
  const m = heightMap.length,
    n = heightMap[0].length;
  if (m < 3 || n < 3) return 0;

  const pq = new MinPriorityQueue({ priority: (c) => c.height });
  const visited = Array.from({ length: m }, () => Array(n).fill(false));

  for (let i = 0; i < m; i++) {
    pq.enqueue({ height: heightMap[i][0], x: i, y: 0 });
    pq.enqueue({ height: heightMap[i][n - 1], x: i, y: n - 1 });
    visited[i][0] = visited[i][n - 1] = true;
  }
  for (let j = 0; j < n; j++) {
    pq.enqueue({ height: heightMap[0][j], x: 0, y: j });
    pq.enqueue({ height: heightMap[m - 1][j], x: m - 1, y: j });
    visited[0][j] = visited[m - 1][j] = true;
  }

  let res = 0;
  const dirs = [
    [0, 1],
    [1, 0],
    [0, -1],
    [-1, 0],
  ];
  while (!pq.isEmpty()) {
    const { height, x, y } = pq.dequeue().element;
    for (const [dx, dy] of dirs) {
      const nx = x + dx,
        ny = y + dy;
      if (nx >= 0 && ny >= 0 && nx < m && ny < n && !visited[nx][ny]) {
        res += Math.max(0, height - heightMap[nx][ny]);
        pq.enqueue({
          height: Math.max(height, heightMap[nx][ny]),
          x: nx,
          y: ny,
        });
        visited[nx][ny] = true;
      }
    }
  }
  return res;
};
```

### Python3

```python
import heapq

class Solution:
    def trapRainWater(self, heightMap: List[List[int]]) -> int:
        if not heightMap or len(heightMap) < 3 or len(heightMap[0]) < 3:
            return 0

        m, n = len(heightMap), len(heightMap[0])
        visited = [[False]*n for _ in range(m)]
        heap = []

        for i in range(m):
            heapq.heappush(heap, (heightMap[i][0], i, 0))
            heapq.heappush(heap, (heightMap[i][n-1], i, n-1))
            visited[i][0] = visited[i][n-1] = True
        for j in range(n):
            heapq.heappush(heap, (heightMap[0][j], 0, j))
            heapq.heappush(heap, (heightMap[m-1][j], m-1, j))
            visited[0][j] = visited[m-1][j] = True

        res = 0
        dirs = [(0,1),(1,0),(0,-1),(-1,0)]
        while heap:
            h, x, y = heapq.heappop(heap)
            for dx,dy in dirs:
                nx, ny = x+dx, y+dy
                if 0<=nx<m and 0<=ny<n and not visited[nx][ny]:
                    res += max(0, h - heightMap[nx][ny])
                    heapq.heappush(heap, (max(h, heightMap[nx][ny]), nx, ny))
                    visited[nx][ny] = True
        return res
```

### Go

```go
import "container/heap"

type Cell struct { height, x, y int }
type MinHeap []Cell
func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].height < h[j].height }
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(Cell)) }
func (h *MinHeap) Pop() interface{} { old := *h; n := len(old); x := old[n-1]; *h = old[:n-1]; return x }

func trapRainWater(heightMap [][]int) int {
    m, n := len(heightMap), len(heightMap[0])
    if m < 3 || n < 3 { return 0 }

    visited := make([][]bool, m)
    for i := range visited { visited[i] = make([]bool, n) }

    h := &MinHeap{}; heap.Init(h)
    for i := 0; i < m; i++ {
        heap.Push(h, Cell{heightMap[i][0], i, 0})
        heap.Push(h, Cell{heightMap[i][n-1], i, n-1})
        visited[i][0], visited[i][n-1] = true, true
    }
    for j := 0; j < n; j++ {
        heap.Push(h, Cell{heightMap[0][j], 0, j})
        heap.Push(h, Cell{heightMap[m-1][j], m-1, j})
        visited[0][j], visited[m-1][j] = true, true
    }

    res := 0
    dirs := [][2]int{{0,1},{1,0},{0,-1},{-1,0}}
    for h.Len() > 0 {
        cell := heap.Pop(h).(Cell)
        for _, d := range dirs {
            nx, ny := cell.x+d[0], cell.y+d[1]
            if nx>=0 && ny>=0 && nx<m && ny<n && !visited[nx][ny] {
                res += max(0, cell.height - heightMap[nx][ny])
                heap.Push(h, Cell{max(cell.height, heightMap[nx][ny]), nx, ny})
                visited[nx][ny] = true
            }
        }
    }
    return res
}
func max(a,b int) int { if a>b { return a }; return b }
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

1. Start with boundary → put into **min-heap**.
2. Always expand the **lowest boundary first**.
3. If neighbor is lower → trap water.
4. Update heap with new height = max(current boundary, neighbor height).
5. Repeat until all visited.

---

## Examples

**Input:**

```
[[1,4,3,1,3,2],
 [3,2,1,3,2,4],
 [2,3,3,2,3,1]]
```

**Output:** `4`

**Explanation:**
Two small ponds, total water trapped = 4.

---

## How to use / Run locally

- Clone repo
- Run corresponding file in your language (C++, Java, JS, Python, Go).
- Use test cases from `examples` section.

---

## Notes & Optimizations

- Using BFS with heap ensures we don’t miss any trapped water.
- Always expanding from the **lowest boundary** is the trick.
- Works efficiently for grid sizes up to `200x200`.

---

## Author

👤 **[Md Aarzoo Islam](https://bento.me/withaarzoo)**
