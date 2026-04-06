# 874. Walking Robot Simulation

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
* [Step-by-step Detailed Explanation](#step-by-step-detailed-explanation)
* [Examples](#examples)
* [How to use / Run locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

---

## Problem Summary

A robot starts at coordinate `(0, 0)` facing north.

It receives a list of commands:

* `-2` means turn left 90 degrees
* `-1` means turn right 90 degrees
* `1 to 9` means move forward by that many steps

There are some obstacles placed on the grid.

If the robot reaches an obstacle, it stops before the obstacle and continues with the next command.

I need to return the maximum squared Euclidean distance from the origin `(0, 0)` reached at any point during the movement.

The squared distance formula is:

```text
x^2 + y^2
```

---

## Constraints

```text
1 <= commands.length <= 10^4
commands[i] is either -2, -1, or an integer in the range [1, 9]
0 <= obstacles.length <= 10^4
-3 * 10^4 <= xi, yi <= 3 * 10^4
```

---

## Intuition

I thought about simulating the robot exactly as the problem describes.

The robot can:

1. Turn left
2. Turn right
3. Move forward

The tricky part is the obstacles.

If I move the robot directly by `k` steps, I may skip an obstacle in the middle.

So I decided to move one step at a time.

To quickly check whether a coordinate contains an obstacle, I store all obstacle positions inside a hash set.

I also use direction indexing:

```text
0 -> North
1 -> East
2 -> South
3 -> West
```

This makes turning very simple using modulo operations.

---

## Approach

1. Store all obstacles in a hash set.
2. Use direction arrays to represent North, East, South, and West.
3. Start from `(0, 0)` facing north.
4. For every command:

   * If command is `-1`, turn right.
   * If command is `-2`, turn left.
   * Otherwise, move step by step.
5. Before each step, check whether the next position contains an obstacle.
6. If obstacle exists, stop moving for that command.
7. Otherwise, update the robot position.
8. Calculate the squared distance after every move.
9. Keep track of the maximum squared distance.

---

## Data Structures Used

| Data Structure | Purpose                                       |
| -------------- | --------------------------------------------- |
| Hash Set       | Stores obstacle coordinates for O(1) lookup   |
| Arrays         | Stores direction changes                      |
| Variables      | Track current position, direction, and answer |

---

## Operations & Behavior Summary

| Command  | Meaning                   |
| -------- | ------------------------- |
| `-2`     | Turn left                 |
| `-1`     | Turn right                |
| `1 to 9` | Move forward step by step |

Direction order:

```text
North -> East -> South -> West
```

Direction arrays:

```text
dx = [0, 1, 0, -1]
dy = [1, 0, -1, 0]
```

---

## Complexity

* Time Complexity: `O(n + m + totalSteps)`

  * `n` = number of obstacles
  * `m` = number of commands
  * `totalSteps` = total number of robot movements

* Space Complexity: `O(n)`

  * Extra space is used for storing obstacles inside a hash set.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int robotSim(vector<int>& commands, vector<vector<int>>& obstacles) {
        unordered_set<string> obstacleSet;

        for (auto &obs : obstacles) {
            obstacleSet.insert(to_string(obs[0]) + "," + to_string(obs[1]));
        }

        vector<int> dx = {0, 1, 0, -1};
        vector<int> dy = {1, 0, -1, 0};

        int dir = 0;
        int x = 0, y = 0;
        int maxDistance = 0;

        for (int command : commands) {
            if (command == -1) {
                dir = (dir + 1) % 4;
            }
            else if (command == -2) {
                dir = (dir + 3) % 4;
            }
            else {
                for (int step = 0; step < command; step++) {
                    int nextX = x + dx[dir];
                    int nextY = y + dy[dir];

                    string nextPos = to_string(nextX) + "," + to_string(nextY);

                    if (obstacleSet.count(nextPos)) {
                        break;
                    }

                    x = nextX;
                    y = nextY;

                    maxDistance = max(maxDistance, x * x + y * y);
                }
            }
        }

        return maxDistance;
    }
};
```

### Java

```java
class Solution {
    public int robotSim(int[] commands, int[][] obstacles) {
        Set<String> obstacleSet = new HashSet<>();

        for (int[] obs : obstacles) {
            obstacleSet.add(obs[0] + "," + obs[1]);
        }

        int[] dx = {0, 1, 0, -1};
        int[] dy = {1, 0, -1, 0};

        int dir = 0;
        int x = 0, y = 0;
        int maxDistance = 0;

        for (int command : commands) {
            if (command == -1) {
                dir = (dir + 1) % 4;
            }
            else if (command == -2) {
                dir = (dir + 3) % 4;
            }
            else {
                for (int step = 0; step < command; step++) {
                    int nextX = x + dx[dir];
                    int nextY = y + dy[dir];

                    String nextPos = nextX + "," + nextY;

                    if (obstacleSet.contains(nextPos)) {
                        break;
                    }

                    x = nextX;
                    y = nextY;

                    maxDistance = Math.max(maxDistance, x * x + y * y);
                }
            }
        }

        return maxDistance;
    }
}
```

### JavaScript

```javascript
var robotSim = function(commands, obstacles) {
    const obstacleSet = new Set();

    for (const [x, y] of obstacles) {
        obstacleSet.add(`${x},${y}`);
    }

    const dx = [0, 1, 0, -1];
    const dy = [1, 0, -1, 0];

    let dir = 0;
    let x = 0, y = 0;
    let maxDistance = 0;

    for (const command of commands) {
        if (command === -1) {
            dir = (dir + 1) % 4;
        }
        else if (command === -2) {
            dir = (dir + 3) % 4;
        }
        else {
            for (let step = 0; step < command; step++) {
                const nextX = x + dx[dir];
                const nextY = y + dy[dir];

                const nextPos = `${nextX},${nextY}`;

                if (obstacleSet.has(nextPos)) {
                    break;
                }

                x = nextX;
                y = nextY;

                maxDistance = Math.max(maxDistance, x * x + y * y);
            }
        }
    }

    return maxDistance;
};
```

### Python3

```python
class Solution:
    def robotSim(self, commands: List[int], obstacles: List[List[int]]) -> int:
        obstacle_set = set()

        for x, y in obstacles:
            obstacle_set.add((x, y))

        dx = [0, 1, 0, -1]
        dy = [1, 0, -1, 0]

        direction = 0
        x, y = 0, 0
        max_distance = 0

        for command in commands:
            if command == -1:
                direction = (direction + 1) % 4
            elif command == -2:
                direction = (direction + 3) % 4
            else:
                for _ in range(command):
                    next_x = x + dx[direction]
                    next_y = y + dy[direction]

                    if (next_x, next_y) in obstacle_set:
                        break

                    x, y = next_x, next_y
                    max_distance = max(max_distance, x * x + y * y)

        return max_distance
```

### Go

```go
func robotSim(commands []int, obstacles [][]int) int {
    obstacleSet := make(map[string]bool)

    for _, obs := range obstacles {
        key := fmt.Sprintf("%d,%d", obs[0], obs[1])
        obstacleSet[key] = true
    }

    dx := []int{0, 1, 0, -1}
    dy := []int{1, 0, -1, 0}

    dir := 0
    x, y := 0, 0
    maxDistance := 0

    for _, command := range commands {
        if command == -1 {
            dir = (dir + 1) % 4
        } else if command == -2 {
            dir = (dir + 3) % 4
        } else {
            for step := 0; step < command; step++ {
                nextX := x + dx[dir]
                nextY := y + dy[dir]

                key := fmt.Sprintf("%d,%d", nextX, nextY)

                if obstacleSet[key] {
                    break
                }

                x = nextX
                y = nextY

                distance := x*x + y*y
                if distance > maxDistance {
                    maxDistance = distance
                }
            }
        }
    }

    return maxDistance
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

### 1. Store obstacles in a hash set

I store every obstacle coordinate in a hash set.

Example:

```text
(2, 4) -> "2,4"
```

This helps me check whether a position is blocked in O(1) time.

---

### 2. Use direction arrays

```text
dx = [0, 1, 0, -1]
dy = [1, 0, -1, 0]
```

These represent:

```text
North -> (0, 1)
East  -> (1, 0)
South -> (0, -1)
West  -> (-1, 0)
```

---

### 3. Handle turning

For turning right:

```text
dir = (dir + 1) % 4
```

For turning left:

```text
dir = (dir + 3) % 4
```

This keeps the direction value between `0` and `3`.

---

### 4. Move one step at a time

If the command is `4`, I do not jump directly.

I move like this:

```text
Step 1
Step 2
Step 3
Step 4
```

This is important because an obstacle may exist between the start and end position.

---

### 5. Check for obstacle before moving

Before every move:

```text
nextX = x + dx[dir]
nextY = y + dy[dir]
```

If `(nextX, nextY)` is an obstacle, I stop moving for that command.

Otherwise, I update the position.

---

### 6. Update the answer

After every successful move:

```text
answer = max(answer, x*x + y*y)
```

The problem asks for squared Euclidean distance, not the actual distance.

---

## Examples

### Example 1

```text
Input:
commands = [4,-1,3]
obstacles = []

Output:
25
```

Explanation:

```text
Start at (0,0)
Move north to (0,4)
Turn right
Move east to (3,4)

Maximum distance = 3^2 + 4^2 = 25
```

### Example 2

```text
Input:
commands = [4,-1,4,-2,4]
obstacles = [[2,4]]

Output:
65
```

Explanation:

```text
Start at (0,0)
Move north to (0,4)
Turn right
Obstacle at (2,4)
Robot stops at (1,4)
Turn left
Move north to (1,8)

Maximum distance = 1^2 + 8^2 = 65
```

---

## How to use / Run locally

### C++

```bash
g++ filename.cpp -o output
./output
```

### Java

```bash
javac Main.java
java Main
```

### JavaScript

```bash
node filename.js
```

### Python3

```bash
python filename.py
```

### Go

```bash
go run filename.go
```

---

## Notes & Optimizations

* Using a hash set is important because obstacle lookup becomes O(1).
* Moving step by step is necessary to avoid skipping obstacles.
* Direction indexing makes turning logic simple.
* The total movement is small because each command is at most 9.
* This solution is efficient enough for all constraints.

---

## Author

* [Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
