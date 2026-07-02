# 3286. Find a Safe Walk Through a Grid

A clean and optimized solution for **LeetCode 3286 - Find a Safe Walk Through a Grid** using **0-1 BFS (Breadth-First Search)**. This repository explains the intuition, approach, complexity analysis, and provides solutions in **C++, Java, JavaScript, Python, and Go**.

This guide is written for beginners who want to understand not only how the algorithm works but also why it works.

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

In this problem, you are given a binary grid and an integer called **health**.

- A cell with value **0** is safe.
- A cell with value **1** is unsafe and decreases your health by **1** when you enter it.

You start from the top-left corner and want to reach the bottom-right corner by moving one cell at a time in four directions:

- Up
- Down
- Left
- Right

The goal is to determine whether it is possible to reach the destination while keeping your health **greater than or equal to 1**.

At first, this looks like a normal grid traversal problem. After looking more carefully, the real objective becomes finding a path that passes through the fewest unsafe cells. That observation naturally leads to a shortest path problem where every move has a cost of either **0** or **1**, making **0-1 BFS** the perfect algorithm.

This problem is a great example of combining **graph traversal**, **shortest path algorithms**, and **deque-based BFS optimization**.

---

## Constraints

| Constraint | Value |
|------------|-------|
| Rows | `m == grid.length` |
| Columns | `n == grid[i].length` |
| Grid Size | `1 <= m, n <= 50` |
| Total Cells | `2 <= m × n` |
| Health | `1 <= health <= m + n` |
| Grid Values | `0` or `1` |

---

## Intuition

The first thing I noticed was that every unsafe cell always costs exactly one unit of health.

That means I don't really care about how many steps I take. I only care about how many unsafe cells I walk through.

So instead of thinking about remaining health after every move, I changed the problem into finding the path with the minimum total cost.

The cost of entering a cell becomes:

- Safe cell = 0
- Unsafe cell = 1

Once I know the minimum health required to reach the destination, checking the answer becomes very simple.

If the minimum health loss is smaller than the given health, I can safely reach the destination.

Since every edge weight is either 0 or 1, I don't need Dijkstra's algorithm. A much faster and cleaner solution is **0-1 BFS**, which is designed specifically for this type of graph.

---

## Approach

I solve the problem using the following steps.

1. Create a distance matrix that stores the minimum health lost to reach every cell.

2. Initialize every cell with infinity because initially no path is known.

3. The starting position already counts, so its initial cost becomes the value of the first grid cell.

4. Use a deque instead of a normal queue.

5. Explore all four neighboring cells.

6. If moving into the next cell costs 0, insert it at the front of the deque.

7. If moving into the next cell costs 1, insert it at the back.

8. Continue updating the minimum cost whenever a better path is found.

9. After the traversal finishes, compare the minimum health loss with the given health.

If the remaining health is at least one, return `true`. Otherwise, return `false`.

---

## Data Structures Used

### Grid

The original matrix stores whether each cell is safe or unsafe.

### Distance Matrix

This keeps track of the minimum health lost before reaching every cell.

Without this matrix, the same cell could be visited many unnecessary times.

### Deque

A deque is the key data structure behind **0-1 BFS**.

- Cost 0 transitions are processed immediately.
- Cost 1 transitions wait behind them.

This guarantees an optimal traversal without using a priority queue.

---

## Operations & Behavior Summary

The algorithm performs the following operations.

- Read the grid dimensions.
- Create a distance matrix.
- Initialize the starting cost.
- Push the starting cell into the deque.
- Repeatedly remove the front cell.
- Explore all four neighboring cells.
- Ignore positions outside the grid.
- Compute the new health loss.
- Update the distance if the new path is better.
- Push zero-cost moves to the front.
- Push one-cost moves to the back.
- Continue until every reachable state has been processed.
- Compare the final minimum cost with the available health.

---

## Complexity

| Complexity | Value | Explanation |
|------------|-------|-------------|
| Time Complexity | **O(m × n)** | Every cell is processed only a constant number of times because edge weights are only 0 or 1. |
| Space Complexity | **O(m × n)** | Extra space is used for the distance matrix and deque. |

Where:

- **m** = number of rows
- **n** = number of columns

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    bool findSafeWalk(vector<vector<int>>& grid, int health) {
        int m = grid.size(), n = grid[0].size();

        // Store the minimum health lost to reach every cell
        vector<vector<int>> dist(m, vector<int>(n, INT_MAX));

        // Deque used by 0-1 BFS
        deque<pair<int, int>> dq;

        // Starting cost includes the starting cell itself
        dist[0][0] = grid[0][0];
        dq.push_front({0, 0});

        // Four possible movement directions
        int dir[5] = {-1, 0, 1, 0, -1};

        while (!dq.empty()) {
            auto [x, y] = dq.front();
            dq.pop_front();

            // Try all four neighboring cells
            for (int k = 0; k < 4; k++) {
                int nx = x + dir[k];
                int ny = y + dir[k + 1];

                // Ignore cells outside the grid
                if (nx < 0 || ny < 0 || nx >= m || ny >= n)
                    continue;

                // Entering the next cell adds either 0 or 1 cost
                int newCost = dist[x][y] + grid[nx][ny];

                // Update only if this path is better
                if (newCost < dist[nx][ny]) {
                    dist[nx][ny] = newCost;

                    // Cost 0 goes to the front, cost 1 goes to the back
                    if (grid[nx][ny] == 0)
                        dq.push_front({nx, ny});
                    else
                        dq.push_back({nx, ny});
                }
            }
        }

        // Health must remain at least 1
        return dist[m - 1][n - 1] < health;
    }
};
```

### Java

```java
class Solution {
    public boolean findSafeWalk(List<List<Integer>> grid, int health) {
        int m = grid.size();
        int n = grid.get(0).size();

        // Store minimum health lost for every cell
        int[][] dist = new int[m][n];
        for (int i = 0; i < m; i++) {
            java.util.Arrays.fill(dist[i], Integer.MAX_VALUE);
        }

        // Deque for 0-1 BFS
        java.util.ArrayDeque<int[]> dq = new java.util.ArrayDeque<>();

        // Starting cost includes the starting cell
        dist[0][0] = grid.get(0).get(0);
        dq.offerFirst(new int[]{0, 0});

        int[] dir = {-1, 0, 1, 0, -1};

        while (!dq.isEmpty()) {
            int[] cur = dq.pollFirst();
            int x = cur[0];
            int y = cur[1];

            // Visit all four neighbors
            for (int k = 0; k < 4; k++) {
                int nx = x + dir[k];
                int ny = y + dir[k + 1];

                // Skip invalid positions
                if (nx < 0 || ny < 0 || nx >= m || ny >= n)
                    continue;

                // Add the cost of entering the next cell
                int newCost = dist[x][y] + grid.get(nx).get(ny);

                // Keep only the best cost
                if (newCost < dist[nx][ny]) {
                    dist[nx][ny] = newCost;

                    // Weight 0 goes to the front, weight 1 goes to the back
                    if (grid.get(nx).get(ny) == 0)
                        dq.offerFirst(new int[]{nx, ny});
                    else
                        dq.offerLast(new int[]{nx, ny});
                }
            }
        }

        // Final health must be positive
        return dist[m - 1][n - 1] < health;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[][]} grid
 * @param {number} health
 * @return {boolean}
 */
var findSafeWalk = function(grid, health) {
    const m = grid.length;
    const n = grid[0].length;

    // Minimum health lost for each cell
    const dist = Array.from({ length: m }, () => Array(n).fill(Infinity));

    // Deque implementation using an array and head pointer
    const deque = [];
    let head = 0;

    // Starting cost includes the starting cell
    dist[0][0] = grid[0][0];
    deque.push([0, 0]);

    const dx = [-1, 0, 1, 0];
    const dy = [0, 1, 0, -1];

    while (head < deque.length) {
        const [x, y] = deque[head++];

        // Explore all four directions
        for (let k = 0; k < 4; k++) {
            const nx = x + dx[k];
            const ny = y + dy[k];

            // Skip invalid cells
            if (nx < 0 || ny < 0 || nx >= m || ny >= n) continue;

            // Cost after entering the next cell
            const newCost = dist[x][y] + grid[nx][ny];

            // Update only if this path is better
            if (newCost < dist[nx][ny]) {
                dist[nx][ny] = newCost;

                // Weight 0 should be processed earlier
                if (grid[nx][ny] === 0) {
                    deque.splice(head, 0, [nx, ny]);
                } else {
                    deque.push([nx, ny]);
                }
            }
        }
    }

    // Health must stay at least 1
    return dist[m - 1][n - 1] < health;
};
```

### Python3

```python
from collections import deque
from typing import List

class Solution:
    def findSafeWalk(self, grid: List[List[int]], health: int) -> bool:
        m, n = len(grid), len(grid[0])

        # Store the minimum health lost to reach every cell
        dist = [[float("inf")] * n for _ in range(m)]

        # Deque used by 0-1 BFS
        dq = deque()

        # Starting cost includes the starting cell
        dist[0][0] = grid[0][0]
        dq.appendleft((0, 0))

        directions = [(-1, 0), (1, 0), (0, -1), (0, 1)]

        while dq:
            x, y = dq.popleft()

            # Visit all neighboring cells
            for dx, dy in directions:
                nx, ny = x + dx, y + dy

                # Ignore invalid positions
                if not (0 <= nx < m and 0 <= ny < n):
                    continue

                # Cost after entering the next cell
                new_cost = dist[x][y] + grid[nx][ny]

                # Update if this path is better
                if new_cost < dist[nx][ny]:
                    dist[nx][ny] = new_cost

                    # Weight 0 goes to the front
                    if grid[nx][ny] == 0:
                        dq.appendleft((nx, ny))
                    else:
                        dq.append((nx, ny))

        # Final health must remain positive
        return dist[m - 1][n - 1] < health
```

### Go

```go
func findSafeWalk(grid [][]int, health int) bool {
 m := len(grid)
 n := len(grid[0])

 const INF = int(1e9)

 // Store the minimum health lost for every cell
 dist := make([][]int, m)
 for i := 0; i < m; i++ {
  dist[i] = make([]int, n)
  for j := 0; j < n; j++ {
   dist[i][j] = INF
  }
 }

 // Simple deque implementation
 type Pair struct {
  x, y int
 }

 deque := make([]Pair, 0)
 head := 0

 // Starting cost includes the starting cell
 dist[0][0] = grid[0][0]
 deque = append(deque, Pair{0, 0})

 dx := []int{-1, 0, 1, 0}
 dy := []int{0, 1, 0, -1}

 for head < len(deque) {
  cur := deque[head]
  head++

  // Explore all four directions
  for k := 0; k < 4; k++ {
   nx := cur.x + dx[k]
   ny := cur.y + dy[k]

   // Skip cells outside the grid
   if nx < 0 || ny < 0 || nx >= m || ny >= n {
    continue
   }

   // Cost after entering the next cell
   newCost := dist[cur.x][cur.y] + grid[nx][ny]

   // Keep only the best cost
   if newCost < dist[nx][ny] {
    dist[nx][ny] = newCost

    // Weight 0 should be processed first
    if grid[nx][ny] == 0 {
     deque = append([]Pair{{nx, ny}}, deque[head:]...)
     head = 0
    } else {
     deque = append(deque, Pair{nx, ny})
    }
   }
  }
 }

 // Health must remain at least 1
 return dist[m-1][n-1] < health
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic is exactly the same in every language. Only the syntax changes.

The algorithm begins by creating a distance matrix. This matrix stores the minimum amount of health that has been lost before reaching each cell.

Initially, every value is set to infinity because no path has been discovered yet.

The starting position is handled separately. Since the player starts on that cell, its value immediately contributes to the total health loss.

Next, a deque is created.

Unlike a normal queue, a deque allows insertion from both the front and the back.

This is what makes **0-1 BFS** so efficient.

The algorithm repeatedly removes one cell from the front of the deque.

For every removed cell, all four neighboring positions are checked.

If a neighboring position lies outside the grid, it is skipped immediately.

Otherwise, the new health loss is calculated by adding the value of the destination cell.

If this new value is smaller than the previously stored distance, the algorithm updates the distance matrix.

Now comes the most important step.

If entering the new cell costs zero health, that cell is inserted at the front of the deque.

This allows free moves to be processed immediately.

If entering the new cell costs one health, the cell is inserted at the back.

This naturally processes cheaper paths before more expensive ones without requiring a priority queue.

Eventually, the deque becomes empty.

At this point, every reachable cell already contains its minimum possible health loss.

The destination cell now stores the smallest amount of health needed to reach it.

Finally, the algorithm compares that value with the available health.

If the remaining health is at least one, the answer is `true`.

Otherwise, the answer is `false`.

Since the algorithm depends only on the order of processing and not on language-specific features, the behavior is identical in C++, Java, JavaScript, Python3, and Go.

---

## Examples

### Example 1

**Input**

```text
grid = [[0,1,0,0,0],
        [0,1,0,1,0],
        [0,0,0,1,0]]

health = 1
```

**Output**

```text
true
```

**Explanation**

The algorithm finds a path that never loses more than zero health.

The remaining health is still positive after reaching the destination.

---

### Example 2

**Input**

```text
grid = [[0,1,1,0,0,0],
        [1,0,1,0,0,0],
        [0,1,1,1,0,1],
        [0,0,1,0,1,0]]

health = 3
```

**Output**

```text
false
```

**Explanation**

Even the safest possible path passes through too many unsafe cells.

The remaining health becomes zero before reaching the destination.

---

### Example 3

**Input**

```text
grid = [[1,1,1],
        [1,0,1],
        [1,1,1]]

health = 5
```

**Output**

```text
true
```

**Explanation**

The algorithm correctly chooses the path with the minimum health loss.

Although many cells are unsafe, the available health is still enough to survive.

---

## How to Use / Run Locally

Clone the repository.

```bash
git clone https://github.com/your-username/your-repository.git
```

Move into the project directory.

```bash
cd your-repository
```

### Compile and Run C++

```bash
g++ main.cpp -o main
./main
```

### Compile and Run Java

```bash
javac Solution.java
java Solution
```

### Run JavaScript

```bash
node solution.js
```

### Run Python3

```bash
python solution.py
```

or

```bash
python3 solution.py
```

### Run Go

```bash
go run solution.go
```

---

## Notes & Optimizations

- This problem is one of the best use cases for **0-1 BFS**.
- A normal BFS does not work because every move does not have the same cost.
- Dijkstra's algorithm also works but is slower because it requires a priority queue.
- The distance matrix prevents revisiting cells with worse paths.
- The starting cell contributes to the total health loss, so it must be included in the initial cost.
- The final comparison should check whether the minimum health loss is strictly smaller than the available health, since the remaining health must stay positive.
- This solution easily handles the maximum grid size within the given constraints.

---

## Author

**Md Aarzoo Islam**

Instagram: <https://www.instagram.com/code.with.aarzoo/>
