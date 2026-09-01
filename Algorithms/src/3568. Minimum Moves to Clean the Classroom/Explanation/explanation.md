# 3568. Minimum Moves to Clean the Classroom

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

In this problem, I have a classroom represented as an `m x n` grid. A student starts from the cell marked `S` and needs to collect every litter cell marked `L`.

The classroom can also contain:

* `S` — starting position
* `L` — litter that needs to be collected
* `R` — reset area that restores energy to its maximum value
* `X` — obstacle that cannot be crossed
* `.` — empty space

The student starts with the given maximum `energy`.

Every move to an adjacent cell costs exactly `1` unit of energy. The student can move up, down, left, or right.

If the student enters an `R` cell, their energy is immediately restored to the original maximum capacity.

The goal is to find the minimum number of moves needed to collect all litter.

If it is impossible to collect all litter, I return `-1`.

This problem combines **Breadth-First Search (BFS)**, **bitmasking**, and **state-space optimization**. The important part is that simply knowing the current cell is not enough. I also need to know which litter has already been collected and how much energy is left.

## Constraints

| Constraint             | Description                          |
| ---------------------- | ------------------------------------ |
| `1 <= m <= 20`         | Number of rows in the classroom      |
| `1 <= n <= 20`         | Number of columns in the classroom   |
| `classroom[i][j]`      | Can be `S`, `L`, `R`, `X`, or `.`    |
| `1 <= energy <= 50`    | Maximum energy capacity              |
| Exactly one `S`        | There is only one starting position  |
| At most `10` `L` cells | There can be at most 10 litter cells |

The small limit of at most `10` litter cells is especially important because it makes a bitmask solution practical.

## Intuition

My first thought was to use BFS because every movement costs exactly one move, and BFS is a natural choice when I need the shortest path.

But there is a problem with using only `(row, column)` as the BFS state.

Suppose I reach the same cell twice. The first time I may have collected some litter and have `5` energy left. The second time I may have collected the same litter but have only `1` energy left.

These two states are not equivalent.

So I realized that the complete state needs to contain:

```text
row + column + collected litter + remaining energy
```

The number of litter cells is at most `10`, so I can use a bitmask to store which litter cells have already been collected.

For example, with three litter cells:

```text
001
```

can mean the first litter is collected, while:

```text
111
```

means all three litter cells are collected.

Then I use BFS over these states.

There is one more optimization. If I have already reached the same position with the same litter mask and more energy, then a state with less energy can never be better.

So I keep the maximum energy seen for every `(row, column, mask)` state and discard weaker states.

## Approach

I solve the problem in these steps:

1. I scan the classroom to find the starting position `S`.

2. During the same scan, I assign every `L` cell a unique bit index from `0` to `k - 1`.

3. I create a target mask where all `k` litter bits are set.

4. I start BFS from the starting position with:

   * no litter collected,
   * full energy,
   * zero moves.

5. For every BFS state, I try moving in all four directions.

6. I skip the move if:

   * the new position is outside the grid,
   * the new cell is an obstacle `X`,
   * or the student's energy becomes negative.

7. Every valid move decreases the energy by `1`.

8. If the new cell is `R`, I reset the energy back to the original maximum.

9. If the new cell is `L`, I set its corresponding bit in the mask.

10. If the new mask contains all litter, I return the current number of moves plus one.

11. Before adding a new state to BFS, I check whether I have already reached the same `(row, column, mask)` with greater or equal energy.

12. If I have, I discard the new state. Otherwise, I save the new energy and add the state to the queue.

If BFS finishes without collecting all litter, I return `-1`.

## Data Structures Used

### Bitmask

I use an integer as a bitmask to represent collected litter.

Since there are at most `10` litter cells, I need at most `10` bits.

This gives at most:

```text
2^10 = 1024
```

different litter combinations.

### BFS Queue

I use a queue to process states in increasing order of the number of moves.

This guarantees that when I first reach a state where all litter has been collected, the answer is minimal.

### `bestEnergy`

I use a 3D array:

```text
bestEnergy[row][column][mask]
```

It stores the maximum energy with which I have reached a particular position after collecting a particular set of litter.

This prevents unnecessary BFS states.

### State

Each BFS state contains:

```text
row
column
mask
remaining energy
moves
```

The `moves` value tells me how far this state is from the starting position.

## Operations & Behavior Summary

The algorithm can be viewed as the following simple process:

```text
Find S.
Find every L and assign it a bit.

Create the mask containing all litter bits.

Start BFS from S with full energy and mask = 0.

While the queue is not empty:
    Take the next state.

    Try up, down, left, and right.

    Ignore invalid cells and obstacles.

    Spend 1 energy for the move.

    If energy becomes negative:
        Ignore the move.

    If the new cell is R:
        Restore full energy.

    If the new cell is L:
        Mark that litter as collected.

    If all litter is collected:
        Return the number of moves.

    If this state has already been reached
    with greater or equal energy:
        Ignore it.

    Otherwise:
        Save the new energy.
        Add the state to the BFS queue.

Return -1 if BFS cannot collect everything.
```

## Complexity

Let:

* `m` = number of rows
* `n` = number of columns
* `k` = number of litter cells
* `E` = maximum energy capacity

| Type  | Complexity                      | Explanation                                                                                                                   |
| ----- | ------------------------------- | ----------------------------------------------------------------------------------------------------------------------------- |
| Time  | `O(m × n × 2^k × E)`            | There are `m × n` positions, `2^k` possible litter masks, and up to `E` useful energy levels. Each state checks 4 directions. |
| Space | `O(m × n × 2^k × E)` worst case | The BFS can contain different energy states, while `bestEnergy` uses `O(m × n × 2^k)` memory.                                 |

Here `k <= 10` and `E <= 50`, so the state space is small enough for this approach.

The `bestEnergy` pruning is important because it removes states that are strictly worse than a state already visited at the same position with the same collected litter.

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int minMoves(vector<string>& classroom, int energy) {
        int m = classroom.size();
        int n = classroom[0].size();

        // Give every litter cell a unique bit position.
        vector<vector<int>> id(m, vector<int>(n, -1));

        int k = 0;
        int sr = 0, sc = 0;

        // Find the starting position and assign IDs to all litter cells.
        for (int r = 0; r < m; r++) {
            for (int c = 0; c < n; c++) {
                if (classroom[r][c] == 'S') {
                    sr = r;
                    sc = c;
                } else if (classroom[r][c] == 'L') {
                    id[r][c] = k++;
                }
            }
        }

        // If there is no litter, the student is already done.
        if (k == 0) return 0;

        int totalMask = (1 << k) - 1;

        // bestEnergy[r][c][mask] stores the maximum energy seen
        // at this position after collecting exactly this mask.
        vector<vector<vector<int>>> best(
            m, vector<vector<int>>(n, vector<int>(1 << k, -1))
        );

        // Each BFS state contains position, collected litter mask,
        // remaining energy, and number of moves used.
        struct State {
            int r, c, mask, e, moves;
        };

        queue<State> q;

        // Initially, no litter is collected and full energy is available.
        best[sr][sc][0] = energy;
        q.push({sr, sc, 0, energy, 0});

        // Four possible movement directions.
        int dr[] = {-1, 1, 0, 0};
        int dc[] = {0, 0, -1, 1};

        while (!q.empty()) {
            State cur = q.front();
            q.pop();

            // Try moving in all four directions.
            for (int d = 0; d < 4; d++) {
                int nr = cur.r + dr[d];
                int nc = cur.c + dc[d];

                // Ignore positions outside the classroom.
                if (nr < 0 || nr >= m || nc < 0 || nc >= n)
                    continue;

                // The student cannot move through obstacles.
                if (classroom[nr][nc] == 'X')
                    continue;

                // Every movement consumes one unit of energy.
                int ne = cur.e - 1;

                // The student cannot make a move without energy.
                if (ne < 0)
                    continue;

                int nmask = cur.mask;

                // Reset the energy immediately after entering an R cell.
                if (classroom[nr][nc] == 'R') {
                    ne = energy;
                }

                // If this cell contains litter, mark its bit as collected.
                if (classroom[nr][nc] == 'L') {
                    nmask |= (1 << id[nr][nc]);
                }

                // All litter has been collected, so this is the answer.
                if (nmask == totalMask) {
                    return cur.moves + 1;
                }

                // If we already reached this position with the same
                // mask and at least this much energy, this state is useless.
                if (ne <= best[nr][nc][nmask])
                    continue;

                // Keep only the strongest energy value for this state.
                best[nr][nc][nmask] = ne;

                // Add the improved state to BFS.
                q.push({nr, nc, nmask, ne, cur.moves + 1});
            }
        }

        // BFS finished without collecting all litter.
        return -1;
    }
};
```

### Java

```java
class Solution {
    public int minMoves(String[] classroom, int energy) {
        int m = classroom.length;
        int n = classroom[0].length();

        // id[r][c] stores the bit number assigned to a litter cell.
        int[][] id = new int[m][n];

        // Initialize all IDs to -1 because non-litter cells have no bit.
        for (int r = 0; r < m; r++) {
            java.util.Arrays.fill(id[r], -1);
        }

        int k = 0;
        int sr = 0, sc = 0;

        // Find S and assign a unique bit to every L cell.
        for (int r = 0; r < m; r++) {
            for (int c = 0; c < n; c++) {
                if (classroom[r].charAt(c) == 'S') {
                    sr = r;
                    sc = c;
                } else if (classroom[r].charAt(c) == 'L') {
                    id[r][c] = k++;
                }
            }
        }

        // No litter means zero moves are needed.
        if (k == 0) return 0;

        int totalMask = (1 << k) - 1;

        // bestEnergy[r][c][mask] keeps the maximum energy
        // already seen for the same position and collected litter.
        int[][][] best = new int[m][n][1 << k];

        // Initialize every state as unseen.
        for (int r = 0; r < m; r++) {
            for (int c = 0; c < n; c++) {
                java.util.Arrays.fill(best[r][c], -1);
            }
        }

        // A state stores position, mask, remaining energy, and moves.
        class State {
            int r, c, mask, e, moves;

            State(int r, int c, int mask, int e, int moves) {
                this.r = r;
                this.c = c;
                this.mask = mask;
                this.e = e;
                this.moves = moves;
            }
        }

        // ArrayDeque gives efficient FIFO operations for BFS.
        java.util.ArrayDeque<State> queue = new java.util.ArrayDeque<>();

        // Start from S with no collected litter and full energy.
        best[sr][sc][0] = energy;
        queue.offer(new State(sr, sc, 0, energy, 0));

        // Four possible movement directions.
        int[] dr = {-1, 1, 0, 0};
        int[] dc = {0, 0, -1, 1};

        while (!queue.isEmpty()) {
            State cur = queue.poll();

            // Try all four neighboring cells.
            for (int d = 0; d < 4; d++) {
                int nr = cur.r + dr[d];
                int nc = cur.c + dc[d];

                // Ignore cells outside the classroom.
                if (nr < 0 || nr >= m || nc < 0 || nc >= n)
                    continue;

                // Obstacles cannot be entered.
                if (classroom[nr].charAt(nc) == 'X')
                    continue;

                // Every move costs one energy.
                int ne = cur.e - 1;

                // Without enough energy, this move is impossible.
                if (ne < 0)
                    continue;

                int nmask = cur.mask;

                // R restores the student's energy to full capacity.
                if (classroom[nr].charAt(nc) == 'R') {
                    ne = energy;
                }

                // Collect the litter by setting its corresponding bit.
                if (classroom[nr].charAt(nc) == 'L') {
                    nmask |= (1 << id[nr][nc]);
                }

                // Return immediately when every litter item is collected.
                if (nmask == totalMask) {
                    return cur.moves + 1;
                }

                // A state with no better energy than an existing state
                // cannot lead to a better answer.
                if (ne <= best[nr][nc][nmask])
                    continue;

                // Store the strongest energy seen for this state.
                best[nr][nc][nmask] = ne;

                // Add the improved state to the BFS queue.
                queue.offer(new State(nr, nc, nmask, ne, cur.moves + 1));
            }
        }

        // No valid path can collect all litter.
        return -1;
    }
}
```

### JavaScript

```javascript
/**
 * @param {string[]} classroom
 * @param {number} energy
 * @return {number}
 */
var minMoves = function(classroom, energy) {
    const m = classroom.length;
    const n = classroom[0].length;

    // id[r][c] stores the bit assigned to a litter cell.
    const id = Array.from({ length: m }, () => Array(n).fill(-1));

    let k = 0;
    let sr = 0;
    let sc = 0;

    // Find S and assign a unique bit to every L cell.
    for (let r = 0; r < m; r++) {
        for (let c = 0; c < n; c++) {
            if (classroom[r][c] === 'S') {
                sr = r;
                sc = c;
            } else if (classroom[r][c] === 'L') {
                id[r][c] = k++;
            }
        }
    }

    // If there is no litter, nothing needs to be moved.
    if (k === 0) return 0;

    // When all k bits are 1, every litter cell has been collected.
    const totalMask = (1 << k) - 1;

    // best[r][c][mask] stores the maximum energy seen for this state.
    const best = Array.from(
        { length: m },
        () => Array.from(
            { length: n },
            () => new Int16Array(1 << k).fill(-1)
        )
    );

    // Each queue item is [row, column, mask, energy, moves].
    const queue = [[sr, sc, 0, energy, 0]];

    // The head index avoids O(n) array shifting during BFS.
    let head = 0;

    // Start with full energy and no collected litter.
    best[sr][sc][0] = energy;

    // Four possible movement directions.
    const dr = [-1, 1, 0, 0];
    const dc = [0, 0, -1, 1];

    while (head < queue.length) {
        // Read the next BFS state without removing earlier elements.
        const [r, c, mask, e, moves] = queue[head++];

        // Try all four neighboring cells.
        for (let d = 0; d < 4; d++) {
            const nr = r + dr[d];
            const nc = c + dc[d];

            // Ignore cells outside the classroom.
            if (nr < 0 || nr >= m || nc < 0 || nc >= n)
                continue;

            // Obstacles cannot be entered.
            if (classroom[nr][nc] === 'X')
                continue;

            // Every move consumes one unit of energy.
            let ne = e - 1;

            // The move is impossible if energy becomes negative.
            if (ne < 0)
                continue;

            let nmask = mask;

            // Entering R restores energy to its full capacity.
            if (classroom[nr][nc] === 'R') {
                ne = energy;
            }

            // Set the bit corresponding to the collected litter.
            if (classroom[nr][nc] === 'L') {
                nmask |= (1 << id[nr][nc]);
            }

            // All litter is collected, so this is the shortest answer.
            if (nmask === totalMask) {
                return moves + 1;
            }

            // Reaching the same state with less or equal energy is dominated.
            if (ne <= best[nr][nc][nmask])
                continue;

            // Store the improved energy value.
            best[nr][nc][nmask] = ne;

            // Add the improved state to BFS.
            queue.push([nr, nc, nmask, ne, moves + 1]);
        }
    }

    // No valid path was able to collect every litter item.
    return -1;
};
```

### Python3

```python3
from typing import List
from collections import deque

class Solution:
    def minMoves(self, classroom: List[str], energy: int) -> int:
        m = len(classroom)
        n = len(classroom[0])

        # id[r][c] stores the bit assigned to a litter cell.
        id = [[-1] * n for _ in range(m)]

        k = 0
        sr = 0
        sc = 0

        # Find the starting position and assign bits to all litter cells.
        for r in range(m):
            for c in range(n):
                if classroom[r][c] == 'S':
                    sr = r
                    sc = c
                elif classroom[r][c] == 'L':
                    id[r][c] = k
                    k += 1

        # If there is no litter, the task is already complete.
        if k == 0:
            return 0

        # This mask has all k litter bits turned on.
        total_mask = (1 << k) - 1

        # best[r][c][mask] stores the maximum energy seen
        # for this position and collected-litter mask.
        best = [
            [
                [-1] * (1 << k)
                for _ in range(n)
            ]
            for _ in range(m)
        ]

        # BFS state is (row, column, mask, remaining energy, moves).
        queue = deque()

        # Start at S with no collected litter and full energy.
        best[sr][sc][0] = energy
        queue.append((sr, sc, 0, energy, 0))

        # Four possible movement directions.
        directions = [(-1, 0), (1, 0), (0, -1), (0, 1)]

        while queue:
            # Get the next state in BFS order.
            r, c, mask, e, moves = queue.popleft()

            # Try moving to all four neighboring cells.
            for dr, dc in directions:
                nr = r + dr
                nc = c + dc

                # Ignore positions outside the classroom.
                if nr < 0 or nr >= m or nc < 0 or nc >= n:
                    continue

                # Obstacles cannot be entered.
                if classroom[nr][nc] == 'X':
                    continue

                # Every movement uses one unit of energy.
                ne = e - 1

                # We cannot move if energy becomes negative.
                if ne < 0:
                    continue

                nmask = mask

                # Reset energy when entering an R cell.
                if classroom[nr][nc] == 'R':
                    ne = energy

                # Mark the litter as collected by setting its bit.
                if classroom[nr][nc] == 'L':
                    nmask |= 1 << id[nr][nc]

                # All litter has been collected, so return the move count.
                if nmask == total_mask:
                    return moves + 1

                # If an equal or stronger state already exists,
                # this new state cannot give us a better path.
                if ne <= best[nr][nc][nmask]:
                    continue

                # Keep the maximum energy for this position and mask.
                best[nr][nc][nmask] = ne

                # Add the improved state to the BFS queue.
                queue.append((nr, nc, nmask, ne, moves + 1))

        # BFS could not find a path that collects all litter.
        return -1
```

### Go

```go
func minMoves(classroom []string, energy int) int {
 m := len(classroom)
 n := len(classroom[0])

 // id[r][c] stores the bit assigned to a litter cell.
 id := make([][]int, m)
 for r := 0; r < m; r++ {
  id[r] = make([]int, n)

  // -1 means this cell does not contain litter.
  for c := 0; c < n; c++ {
   id[r][c] = -1
  }
 }

 k := 0
 sr, sc := 0, 0

 // Find S and assign a unique bit to every L cell.
 for r := 0; r < m; r++ {
  for c := 0; c < n; c++ {
   if classroom[r][c] == 'S' {
    sr, sc = r, c
   } else if classroom[r][c] == 'L' {
    id[r][c] = k
    k++
   }
  }
 }

 // If there is no litter, zero moves are required.
 if k == 0 {
  return 0
 }

 // This mask has every litter bit turned on.
 totalMask := (1 << k) - 1

 // best[r][c][mask] stores the maximum energy seen
 // for this position and collected-litter mask.
 best := make([][][]int, m)

 for r := 0; r < m; r++ {
  best[r] = make([][]int, n)

  for c := 0; c < n; c++ {
   best[r][c] = make([]int, 1<<k)

   // -1 means this position and mask has not been reached yet.
   for mask := 0; mask < (1 << k); mask++ {
    best[r][c][mask] = -1
   }
  }
 }

 // A state stores position, collected mask, remaining energy, and moves.
 type State struct {
  r, c   int
  mask   int
  energy int
  moves  int
 }

 // A slice with a head pointer works as an efficient BFS queue.
 queue := make([]State, 0)
 head := 0

 // Start from S with full energy and no collected litter.
 best[sr][sc][0] = energy
 queue = append(queue, State{sr, sc, 0, energy, 0})

 // Four possible movement directions.
 dr := [4]int{-1, 1, 0, 0}
 dc := [4]int{0, 0, -1, 1}

 for head < len(queue) {
  // Get the next state in BFS order.
  cur := queue[head]
  head++

  // Try moving in all four directions.
  for d := 0; d < 4; d++ {
   nr := cur.r + dr[d]
   nc := cur.c + dc[d]

   // Ignore positions outside the classroom.
   if nr < 0 || nr >= m || nc < 0 || nc >= n {
    continue
   }

   // Obstacles cannot be entered.
   if classroom[nr][nc] == 'X' {
    continue
   }

   // Every move consumes one unit of energy.
   ne := cur.energy - 1

   // The student cannot make a move without energy.
   if ne < 0 {
    continue
   }

   nmask := cur.mask

   // R restores the energy to its maximum capacity.
   if classroom[nr][nc] == 'R' {
    ne = energy
   }

   // Mark the litter as collected using its assigned bit.
   if classroom[nr][nc] == 'L' {
    nmask |= 1 << id[nr][nc]
   }

   // All litter has been collected, so return the shortest distance.
   if nmask == totalMask {
    return cur.moves + 1
   }

   // A state with less or equal energy is dominated by an existing one.
   if ne <= best[nr][nc][nmask] {
    continue
   }

   // Save the strongest energy value seen for this state.
   best[nr][nc][nmask] = ne

   // Add the improved state to the BFS queue.
   queue = append(queue, State{
    nr, nc, nmask, ne, cur.moves + 1,
   })
  }
 }

 // No valid path can collect every litter cell.
 return -1
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

### 1. Reading the grid

The first thing I do in every language is scan the complete classroom.

I look for two things:

* the starting cell `S`
* every litter cell `L`

I save the coordinates of `S` because BFS needs to start there.

For every litter cell, I assign a unique number.

For example:

```text
L . L
. L .
```

can become:

```text
0 . 1
. 2 .
```

The actual grid does not change. I only keep these IDs separately so that I know which bit belongs to each litter.

### 2. Creating the litter mask

If there are `k` litter cells, I need `k` bits.

For example, when `k = 3`:

```text
000
```

means nothing has been collected.

```text
101
```

means litter `0` and litter `2` have been collected.

```text
111
```

means everything has been collected.

The target mask is:

```text
(1 << k) - 1
```

For `k = 3`:

```text
1 << 3 = 1000
1000 - 1 = 0111
```

So `0111` represents all three litter cells.

### 3. Starting BFS

I start with:

```text
row = start row
column = start column
mask = 0
energy = maximum energy
moves = 0
```

At this point, no litter has been collected.

The student also has the full energy capacity.

### 4. Why BFS works

Every movement has the same cost: exactly one move.

That makes BFS a good fit for finding the minimum number of moves.

BFS explores states like this:

```text
0 moves
    ↓
1 move
    ↓
2 moves
    ↓
3 moves
    ↓
...
```

Therefore, when I first reach a state where all litter is collected, I know that no solution with fewer moves exists.

### 5. Trying the four directions

From every state, I try:

```text
up
down
left
right
```

I calculate the new row and column.

If the new position is outside the grid, I skip it.

I also skip `X` because obstacles cannot be crossed.

### 6. Spending energy

Every movement costs one energy.

So I calculate:

```text
newEnergy = currentEnergy - 1
```

If this becomes negative, the student cannot make the move.

An important detail is that an `R` cell can be entered only if the move itself is affordable.

After entering the `R` cell, the energy is restored to its maximum value.

### 7. Handling reset cells

Suppose the maximum energy is `4` and the student currently has `1` energy.

Moving into `R` costs one energy:

```text
1 -> 0
```

The student has successfully entered the cell.

Then the reset happens:

```text
0 -> 4
```

So the student can continue with full energy.

This is why the reset must happen after checking that the movement itself is valid.

### 8. Collecting litter

When the student enters an `L` cell, I set its bit.

Suppose the current mask is:

```text
001
```

and the new litter has ID `2`.

Its bit is:

```text
100
```

I combine them using bitwise OR:

```text
001
OR 100
---
101
```

Now the mask records that both litter `0` and litter `2` have been collected.

If the same litter is visited again, its bit is already set, so nothing changes.

### 9. Checking the answer

After updating the mask, I compare it with the target mask.

If:

```text
newMask == totalMask
```

then every litter cell has been collected.

Since BFS processes states in increasing move count, I can immediately return:

```text
current moves + 1
```

There is no need to continue searching.

### 10. Why I need `bestEnergy`

This is the most important optimization in the solution.

Imagine these two states:

```text
Position: (4, 5)
Mask:     101

State A: energy = 8
State B: energy = 3
```

Both states have the same position and have collected exactly the same litter.

State A is strictly better because it has more energy.

Anything State B can do, State A can also do. State A may even be able to travel farther before needing a reset.

So I do not need State B.

I store:

```text
bestEnergy[4][5][101] = 8
```

If another state arrives with energy `3`, I discard it.

If another state arrives with energy `10`, I replace `8` with `10` and process the new state.

This reduces the number of states that BFS needs to explore.

### 11. C++ implementation details

In C++, I use `queue` for BFS.

A small `State` structure keeps all information needed for one BFS state.

The 3D `best` vector stores the strongest energy value for each position and mask.

C++ is efficient with these arrays and queues, so the implementation is straightforward.

### 12. Java implementation details

In Java, I use `ArrayDeque` instead of the older `LinkedList` style for the BFS queue.

Each state is represented by a small `State` object containing the row, column, mask, energy, and number of moves.

The `best` 3D array plays the same role as in C++.

Using `ArrayDeque` keeps queue operations efficient.

### 13. JavaScript implementation details

For JavaScript, I use an array as the BFS queue.

I do not use `shift()` because removing the first element repeatedly can be expensive for a large queue.

Instead, I keep an integer called `head`.

The next state is read using:

```text
queue[head++]
```

This gives queue-like behavior without repeatedly moving all remaining elements.

I also use typed arrays for the `best` structure, which keeps the numeric state storage compact.

### 14. Python3 implementation details

In Python, I use `collections.deque`.

A `deque` is designed for efficient insertion and removal from both ends, making it a natural choice for BFS.

The state is stored as a tuple containing:

```text
(row, column, mask, energy, moves)
```

The same pruning rule is used before adding a new state.

### 15. Go implementation details

In Go, I use a slice as the BFS queue together with a `head` index.

This follows the same idea as the JavaScript implementation.

Instead of removing the first element from the slice every time, I simply move the `head` forward.

The BFS state is represented using a Go `struct`.

### 16. Edge cases

There are a few cases I need to handle carefully.

If there is no `L` cell, the answer is immediately `0` because there is nothing to collect.

If the student runs out of energy before reaching a useful reset area, that path cannot continue.

If obstacles completely block access to some litter, BFS eventually becomes empty and I return `-1`.

The student can visit an `R` cell multiple times. I do not mark reset cells as consumed because their effect can be used repeatedly.

The same litter can also be visited multiple times, but its bit remains set after the first collection.

## Examples

### Example 1

Input:

```text
classroom = ["S.", "XL"]
energy = 2
```

Output:

```text
2
```

The student starts at `(0, 0)` with `2` energy.

Moving down is impossible because `(1, 0)` is an obstacle.

So the student moves right:

```text
(0, 0) -> (0, 1)
```

Energy becomes `1`.

Then the student moves down:

```text
(0, 1) -> (1, 1)
```

This cell contains litter, so the litter is collected.

Total moves:

```text
2
```

### Example 2

Input:

```text
classroom = ["LS", "RL"]
energy = 4
```

Output:

```text
3
```

The student starts at `(0, 1)` with `4` energy.

First, the student moves left:

```text
(0, 1) -> (0, 0)
```

This collects the first litter.

Energy becomes `3`.

Then the student moves down:

```text
(0, 0) -> (1, 0)
```

The cell is `R`, so the energy is restored to `4`.

Finally, the student moves right:

```text
(1, 0) -> (1, 1)
```

This collects the second litter.

Total moves:

```text
3
```

### Example 3

Input:

```text
classroom = ["L.S", "RXL"]
energy = 3
```

Output:

```text
-1
```

There is no valid path that allows the student to collect every litter cell.

BFS explores every reachable state while respecting obstacles and energy limits.

Eventually, the queue becomes empty without reaching the mask containing all litter.

Therefore, the answer is:

```text
-1
```

## How to Use / Run Locally

The solution is written in five languages. Since LeetCode provides the `Solution` class or function wrapper, I can paste the corresponding implementation directly into LeetCode.

For local testing, I can also create a small driver program around the solution.

### C++

Save the solution as:

```text
solution.cpp
```

Compile it with:

```bash
g++ -std=c++17 -O2 solution.cpp -o solution
```

Run it with:

```bash
./solution
```

On Windows, the generated executable can be run with:

```bash
solution.exe
```

### Java

Save the solution as:

```text
Solution.java
```

Compile it with:

```bash
javac Solution.java
```

Run it with:

```bash
java Solution
```

The class name should match the file name when running a standalone Java program.

### JavaScript

Save the solution as:

```text
solution.js
```

Run it using Node.js:

```bash
node solution.js
```

Make sure Node.js is installed on your system.

### Python3

Save the solution as:

```text
solution.py
```

Run it with:

```bash
python3 solution.py
```

On some Windows installations, the command may be:

```bash
python solution.py
```

### Go

Save the solution as:

```text
solution.go
```

Run it directly with:

```bash
go run solution.go
```

Or build an executable with:

```bash
go build solution.go
```

## Notes & Optimizations

The biggest optimization in this solution is the `bestEnergy` pruning.

A normal BFS state could be thought of as:

```text
(row, column, mask, energy)
```

That can create many states.

Instead, for every `(row, column, mask)`, I only care about the highest energy value reached so far.

If I already have more energy at the same position with the same collected litter, a weaker state cannot improve the answer.

The bitmask is also important. Since there are at most `10` litter cells, `2^10` possible masks are manageable.

Another important point is that I cannot use a normal `visited[row][column]` array. Two visits to the same cell may have completely different collected-litter masks or energy levels.

I also cannot use only:

```text
visited[row][column][mask]
```

because reaching the same state with different energy values can matter.

The `R` cells do not need their own special visited tracking. Their only job is to restore energy whenever they are entered.

Finally, BFS is preferable to DFS here because the problem asks for the minimum number of moves. DFS could find a valid solution, but it would need additional work to guarantee that the first answer is the shortest.

## Author

[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)
