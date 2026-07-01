# 2812. Find the Safest Path in a Grid

A clean and optimized solution for **LeetCode 2812 - Find the Safest Path in a Grid** using **Multi-Source BFS + Binary Search**. This repository explains the complete thought process, algorithm, complexity analysis, and provides implementations in **C++, Java, JavaScript, Python, and Go**.

This problem is a great example of combining graph algorithms with binary search on the answer. If you're preparing for coding interviews or improving your Data Structures and Algorithms (DSA) skills, this problem is definitely worth understanding.

---

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
- [How to Use / Run Locally](#how-to-use--run-locally)
- [Notes & Optimizations](#notes--optimizations)
- [Author](#author)

---

## Problem Summary

You are given an `n × n` grid where:

- `1` represents a thief.
- `0` represents an empty cell.

You start at the top-left corner `(0, 0)` and want to reach the bottom-right corner `(n - 1, n - 1)` by moving one step at a time in the four possible directions.

The **safeness factor** of a path is defined as the minimum Manhattan distance between any cell on that path and the nearest thief.

Your goal is to find the **maximum possible safeness factor** among all valid paths from the starting cell to the destination.

This problem combines **graph traversal**, **multi-source BFS**, and **binary search**, making it a popular interview question for companies that focus on advanced graph algorithms.

---

## Constraints

| Constraint | Value |
|------------|-------|
| Grid Size | `n × n` |
| `1 <= n <= 400` | Yes |
| `grid[i].length == n` | Yes |
| `grid[i][j]` | `0` or `1` |
| At least one thief exists | Yes |

---

## Intuition

The first thing I noticed was that every path depends on how close its cells are to the nearest thief.

Instead of checking the distance separately for every path, I first calculate the distance from every cell to its nearest thief.

The fastest way to do this is by starting a BFS from every thief at the same time. This is called a **Multi-Source BFS**.

Once every cell knows its distance from the closest thief, the original problem becomes much easier.

Now I only need to answer one question:

"Can I reach the destination if every cell on my path has a distance of at least `X`?"

Since larger values become harder to satisfy, binary search fits perfectly.

---

## Approach

I solve the problem in two major steps.

First, I calculate the minimum distance from every cell to its nearest thief.

- Push every thief into a queue.
- Perform a Multi-Source BFS.
- Store the shortest distance for every cell.

After that, I binary search the answer.

For every candidate safeness factor:

1. Check whether the starting and ending cells satisfy the minimum distance.
2. Run a BFS.
3. Only move into cells whose distance is at least the current safeness factor.
4. If I can reach the destination, the answer may be even larger.
5. Otherwise, I reduce the search range.

Eventually, binary search finds the largest possible safeness factor.

---

## Data Structures Used

| Data Structure | Purpose |
|---------------|---------|
| Queue | Used for Multi-Source BFS and path checking BFS |
| 2D Distance Array | Stores the nearest thief distance for every cell |
| 2D Visited Array | Prevents visiting the same cell multiple times |
| Binary Search Variables | Finds the maximum valid safeness factor efficiently |

---

## Operations & Behavior Summary

The algorithm works in the following order:

1. Scan the grid.
2. Insert every thief into a queue.
3. Perform Multi-Source BFS.
4. Store the shortest distance to the nearest thief for every cell.
5. Binary search the answer.
6. For every candidate answer:
   - Check whether a valid path exists.
   - Ignore cells that do not satisfy the required safeness.
7. Keep searching until the maximum possible safeness factor is found.

---

## Complexity

| Type | Complexity | Explanation |
|------|------------|-------------|
| Time Complexity | **O(n² log n)** | Multi-Source BFS takes `O(n²)` and each binary search step performs one BFS. |
| Space Complexity | **O(n²)** | Extra space is used for the distance matrix, visited matrix, and BFS queue. |

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int maximumSafenessFactor(vector<vector<int>>& grid) {
        int n = grid.size();

        // Distance of every cell from the nearest thief
        vector<vector<int>> dist(n, vector<int>(n, -1));

        queue<pair<int,int>> q;

        // Push every thief into the queue
        for(int i=0;i<n;i++){
            for(int j=0;j<n;j++){
                if(grid[i][j]==1){
                    dist[i][j]=0;
                    q.push({i,j});
                }
            }
        }

        int dir[5]={-1,0,1,0,-1};

        // Multi-source BFS
        while(!q.empty()){
            auto [x,y]=q.front();
            q.pop();

            for(int k=0;k<4;k++){
                int nx=x+dir[k];
                int ny=y+dir[k+1];

                if(nx<0||ny<0||nx>=n||ny>=n||dist[nx][ny]!=-1)
                    continue;

                // First visit always gives the shortest distance
                dist[nx][ny]=dist[x][y]+1;
                q.push({nx,ny});
            }
        }

        // Check if a path exists with safeness >= limit
        auto canReach=[&](int limit){

            if(dist[0][0]<limit || dist[n-1][n-1]<limit)
                return false;

            vector<vector<int>> vis(n, vector<int>(n,0));
            queue<pair<int,int>> bfs;

            bfs.push({0,0});
            vis[0][0]=1;

            while(!bfs.empty()){
                auto [x,y]=bfs.front();
                bfs.pop();

                if(x==n-1 && y==n-1)
                    return true;

                for(int k=0;k<4;k++){
                    int nx=x+dir[k];
                    int ny=y+dir[k+1];

                    if(nx<0||ny<0||nx>=n||ny>=n)
                        continue;

                    if(vis[nx][ny])
                        continue;

                    // Only move through safe enough cells
                    if(dist[nx][ny]<limit)
                        continue;

                    vis[nx][ny]=1;
                    bfs.push({nx,ny});
                }
            }

            return false;
        };

        int left=0;
        int right=2*n;
        int ans=0;

        // Binary search on the answer
        while(left<=right){
            int mid=left+(right-left)/2;

            if(canReach(mid)){
                ans=mid;
                left=mid+1;
            }else{
                right=mid-1;
            }
        }

        return ans;
    }
};
```

### Java

```java
class Solution {
    public int maximumSafenessFactor(List<List<Integer>> grid) {

        int n = grid.size();

        int[][] dist = new int[n][n];
        for(int[] row : dist)
            Arrays.fill(row, -1);

        Queue<int[]> queue = new LinkedList<>();

        // Start BFS from every thief
        for(int i=0;i<n;i++){
            for(int j=0;j<n;j++){
                if(grid.get(i).get(j)==1){
                    dist[i][j]=0;
                    queue.offer(new int[]{i,j});
                }
            }
        }

        int[] dx={-1,0,1,0};
        int[] dy={0,1,0,-1};

        // Multi-source BFS
        while(!queue.isEmpty()){
            int[] cur=queue.poll();

            for(int k=0;k<4;k++){
                int nx=cur[0]+dx[k];
                int ny=cur[1]+dy[k];

                if(nx<0||ny<0||nx>=n||ny>=n||dist[nx][ny]!=-1)
                    continue;

                dist[nx][ny]=dist[cur[0]][cur[1]]+1;
                queue.offer(new int[]{nx,ny});
            }
        }

        int left=0;
        int right=2*n;
        int ans=0;

        while(left<=right){

            int mid=left+(right-left)/2;

            if(canReach(dist, mid, n)){
                ans=mid;
                left=mid+1;
            }else{
                right=mid-1;
            }
        }

        return ans;
    }

    private boolean canReach(int[][] dist,int limit,int n){

        if(dist[0][0]<limit || dist[n-1][n-1]<limit)
            return false;

        Queue<int[]> queue=new LinkedList<>();
        boolean[][] vis=new boolean[n][n];

        queue.offer(new int[]{0,0});
        vis[0][0]=true;

        int[] dx={-1,0,1,0};
        int[] dy={0,1,0,-1};

        while(!queue.isEmpty()){

            int[] cur=queue.poll();

            if(cur[0]==n-1 && cur[1]==n-1)
                return true;

            for(int k=0;k<4;k++){

                int nx=cur[0]+dx[k];
                int ny=cur[1]+dy[k];

                if(nx<0||ny<0||nx>=n||ny>=n)
                    continue;

                if(vis[nx][ny] || dist[nx][ny]<limit)
                    continue;

                vis[nx][ny]=true;
                queue.offer(new int[]{nx,ny});
            }
        }

        return false;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[][]} grid
 * @return {number}
 */
var maximumSafenessFactor = function(grid) {

    const n = grid.length;

    const dist = Array.from({length:n},()=>Array(n).fill(-1));
    const queue = [];

    // Push every thief into the queue
    for(let i=0;i<n;i++){
        for(let j=0;j<n;j++){
            if(grid[i][j]===1){
                dist[i][j]=0;
                queue.push([i,j]);
            }
        }
    }

    const dir=[-1,0,1,0,-1];
    let head=0;

    // Multi-source BFS
    while(head<queue.length){

        const [x,y]=queue[head++];

        for(let k=0;k<4;k++){

            const nx=x+dir[k];
            const ny=y+dir[k+1];

            if(nx<0||ny<0||nx>=n||ny>=n)
                continue;

            if(dist[nx][ny]!==-1)
                continue;

            dist[nx][ny]=dist[x][y]+1;
            queue.push([nx,ny]);
        }
    }

    function canReach(limit){

        if(dist[0][0]<limit || dist[n-1][n-1]<limit)
            return false;

        const vis=Array.from({length:n},()=>Array(n).fill(false));
        const bfs=[[0,0]];
        let idx=0;

        vis[0][0]=true;

        while(idx<bfs.length){

            const [x,y]=bfs[idx++];

            if(x===n-1 && y===n-1)
                return true;

            for(let k=0;k<4;k++){

                const nx=x+dir[k];
                const ny=y+dir[k+1];

                if(nx<0||ny<0||nx>=n||ny>=n)
                    continue;

                if(vis[nx][ny] || dist[nx][ny]<limit)
                    continue;

                vis[nx][ny]=true;
                bfs.push([nx,ny]);
            }
        }

        return false;
    }

    let left=0;
    let right=2*n;
    let ans=0;

    while(left<=right){

        const mid=Math.floor((left+right)/2);

        if(canReach(mid)){
            ans=mid;
            left=mid+1;
        }else{
            right=mid-1;
        }
    }

    return ans;
};
```

### Python3

```python
from collections import deque

class Solution:
    def maximumSafenessFactor(self, grid: List[List[int]]) -> int:

        n = len(grid)

        # Distance from the nearest thief
        dist = [[-1] * n for _ in range(n)]

        q = deque()

        # Push every thief into the queue
        for i in range(n):
            for j in range(n):
                if grid[i][j] == 1:
                    dist[i][j] = 0
                    q.append((i, j))

        directions = [(-1,0),(1,0),(0,-1),(0,1)]

        # Multi-source BFS
        while q:
            x, y = q.popleft()

            for dx, dy in directions:
                nx, ny = x + dx, y + dy

                if 0 <= nx < n and 0 <= ny < n and dist[nx][ny] == -1:
                    dist[nx][ny] = dist[x][y] + 1
                    q.append((nx, ny))

        # Check whether a path exists
        def canReach(limit):

            if dist[0][0] < limit or dist[n-1][n-1] < limit:
                return False

            vis = [[False] * n for _ in range(n)]
            bfs = deque([(0, 0)])
            vis[0][0] = True

            while bfs:

                x, y = bfs.popleft()

                if x == n - 1 and y == n - 1:
                    return True

                for dx, dy in directions:

                    nx, ny = x + dx, y + dy

                    if 0 <= nx < n and 0 <= ny < n:
                        if not vis[nx][ny] and dist[nx][ny] >= limit:
                            vis[nx][ny] = True
                            bfs.append((nx, ny))

            return False

        left = 0
        right = 2 * n
        ans = 0

        # Binary search on the answer
        while left <= right:

            mid = (left + right) // 2

            if canReach(mid):
                ans = mid
                left = mid + 1
            else:
                right = mid - 1

        return ans
```

### Go

```go
func maximumSafenessFactor(grid [][]int) int {

 n := len(grid)

 // Distance from every cell to the nearest thief
 dist := make([][]int, n)
 for i := range dist {
  dist[i] = make([]int, n)
  for j := range dist[i] {
   dist[i][j] = -1
  }
 }

 type Pair struct {
  x, y int
 }

 queue := []Pair{}

 // Push every thief
 for i := 0; i < n; i++ {
  for j := 0; j < n; j++ {
   if grid[i][j] == 1 {
    dist[i][j] = 0
    queue = append(queue, Pair{i, j})
   }
  }
 }

 dir := []int{-1, 0, 1, 0, -1}
 head := 0

 // Multi-source BFS
 for head < len(queue) {

  cur := queue[head]
  head++

  for k := 0; k < 4; k++ {

   nx := cur.x + dir[k]
   ny := cur.y + dir[k+1]

   if nx < 0 || ny < 0 || nx >= n || ny >= n || dist[nx][ny] != -1 {
    continue
   }

   dist[nx][ny] = dist[cur.x][cur.y] + 1
   queue = append(queue, Pair{nx, ny})
  }
 }

 canReach := func(limit int) bool {

  if dist[0][0] < limit || dist[n-1][n-1] < limit {
   return false
  }

  vis := make([][]bool, n)
  for i := range vis {
   vis[i] = make([]bool, n)
  }

  bfs := []Pair{{0, 0}}
  vis[0][0] = true
  idx := 0

  for idx < len(bfs) {

   cur := bfs[idx]
   idx++

   if cur.x == n-1 && cur.y == n-1 {
    return true
   }

   for k := 0; k < 4; k++ {

    nx := cur.x + dir[k]
    ny := cur.y + dir[k+1]

    if nx < 0 || ny < 0 || nx >= n || ny >= n {
     continue
    }

    if vis[nx][ny] || dist[nx][ny] < limit {
     continue
    }

    vis[nx][ny] = true
    bfs = append(bfs, Pair{nx, ny})
   }
  }

  return false
 }

 left := 0
 right := 2 * n
 ans := 0

 // Binary search on the answer
 for left <= right {

  mid := left + (right-left)/2

  if canReach(mid) {
   ans = mid
   left = mid + 1
  } else {
   right = mid - 1
  }
 }

 return ans
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic remains exactly the same across all five languages. Only the syntax changes.

### Step 1 — Calculate the nearest thief distance

I begin by creating a distance matrix.

Initially, every distance is unknown.

Next, I add every thief cell into the queue before starting BFS.

Since all thieves start spreading at the same time, the first time a cell is visited always represents its shortest distance from the nearest thief.

This finishes in one complete traversal of the grid.

---

### Step 2 — Prepare for Binary Search

Now every cell has a safety value.

Instead of directly searching for a path, I binary search the answer.

If a path exists with safeness factor `5`, then it must also exist for `4`, `3`, `2`, and `1`.

This property makes binary search possible.

---

### Step 3 — Validate One Candidate Answer

For every middle value produced by binary search:

- Check the starting cell.
- Check the destination.
- If either one is already unsafe, stop immediately.

Otherwise, perform another BFS.

During this BFS, only cells whose distance is at least the required safeness are allowed.

Unsafe cells behave like blocked cells.

---

### Step 4 — Reach the Destination

If BFS reaches the bottom-right corner, then the current safeness factor is possible.

Binary search then tries a larger value.

If BFS cannot reach the destination, the safeness factor is too high.

Binary search then looks for a smaller answer.

---

### Step 5 — Return the Best Answer

Eventually, binary search finishes.

The stored answer is the maximum safeness factor possible for any valid path.

This combination of Multi-Source BFS and Binary Search keeps the solution fast enough even for the largest constraints.

---

## Examples

### Example 1

**Input**

```text
grid = [[1,0,0],
        [0,0,0],
        [0,0,1]]
```

**Output**

```text
0
```

**Explanation**

The start and destination both contain thieves.

No matter which path is chosen, the minimum distance becomes `0`.

---

### Example 2

**Input**

```text
grid = [[0,0,1],
        [0,0,0],
        [0,0,0]]
```

**Output**

```text
2
```

**Explanation**

Multi-Source BFS computes the nearest thief distance for every cell.

Binary search verifies that a path with minimum distance `2` exists, while no path with a larger safeness factor is possible.

---

### Example 3

**Input**

```text
grid = [[0,0,0,1],
        [0,0,0,0],
        [0,0,0,0],
        [1,0,0,0]]
```

**Output**

```text
2
```

**Explanation**

The safest path stays away from both thieves and maintains a minimum distance of `2` throughout the journey.

---

## How to Use / Run Locally

Clone the repository.

```bash
git clone <repository-url>
```

Move into the project folder.

```bash
cd <repository-name>
```

### C++

Compile:

```bash
g++ solution.cpp -o solution
```

Run:

```bash
./solution
```

---

### Java

Compile:

```bash
javac Solution.java
```

Run:

```bash
java Solution
```

---

### JavaScript

Run:

```bash
node solution.js
```

---

### Python3

Run:

```bash
python solution.py
```

or

```bash
python3 solution.py
```

---

### Go

Run:

```bash
go run solution.go
```

---

## Notes & Optimizations

- Running BFS from every cell separately would be far too slow.
- Multi-Source BFS computes all nearest thief distances in one traversal.
- Binary search reduces the number of path checks significantly.
- BFS is used twice, but each traversal still visits every cell at most once.
- This solution comfortably handles the maximum grid size of `400 × 400`.
- Another possible solution uses a priority queue together with Union Find, but the Multi-Source BFS + Binary Search approach is easier to understand and implement.
- This is a classic interview problem involving graph traversal, shortest distance computation, and binary search on the answer.

---

## Author

**Md Aarzoo Islam**

<https://www.instagram.com/code.with.aarzoo/>
