# 2069. Walking Robot Simulation II

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
* [Step-by-step Detailed Explanation](#step-by-step-detailed-explanation-c-java-javascript-python3-go)
* [Examples](#examples)
* [How to use / Run locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

---

## Problem Summary

A robot moves inside a rectangular grid of size `width x height`.

* The robot starts at `(0, 0)`.
* Initially, it faces `East`.
* If the robot tries to move outside the grid, it rotates counterclockwise and tries again.
* We need to support:

  * `step(num)` → move the robot `num` times
  * `getPos()` → return current position
  * `getDir()` → return current direction

The main challenge is that `num` can be very large, so simulating every move directly is too slow.

---

## Constraints

* `2 <= width, height <= 100`
* `1 <= num <= 10^5`
* At most `10^4` calls will be made to:

  * `step`
  * `getPos`
  * `getDir`

---

## Intuition

I noticed that the robot never moves randomly inside the whole grid.

It always moves only along the boundary.

So instead of simulating every single move, I can think of the boundary as one circular path.

The total number of moves required to complete one full cycle is:

```text
2 * (width + height) - 4
```

This is the perimeter of the boundary path.

So if the robot needs to move `num` steps:

* I only need `num % perimeter`
* Because every full loop brings the robot back to the same position

---

## Approach

1. Store the robot's current position `(x, y)`.
2. Store the current direction index.
3. Use arrays for direction movement.
4. Precompute the perimeter.
5. Reduce the number of moves using modulo.
6. Rotate when the next move is invalid.
7. Return the current position and direction when needed.

---

## Data Structures Used

* Integer variables:

  * `x`, `y` for position
  * `dir` for direction index
  * `width`, `height`
  * `perimeter`

* Arrays:

  * `dx[]` for x-direction movement
  * `dy[]` for y-direction movement
  * `dirs[]` for direction names

Direction mapping:

```text
0 -> East
1 -> North
2 -> West
3 -> South
```

---

## Operations & Behavior Summary

| Operation   | Description                                            |
| ----------- | ------------------------------------------------------ |
| `step(num)` | Moves the robot `num` steps                            |
| `getPos()`  | Returns `[x, y]`                                       |
| `getDir()`  | Returns current direction                              |
| Rotate Rule | Rotate counterclockwise when next move is outside grid |
| Full Cycle  | Perimeter steps return robot to same position          |

---

## Complexity

* Time Complexity: `O(perimeter)` per `step(num)` call

  * Perimeter = `2 * (width + height) - 4`
  * Since width and height are at most 100, this is very small

* Space Complexity: `O(1)`

  * Only a few variables are used

---

## Multi-language Solutions

### C++

```cpp
class Robot {
public:
    int width, height, perimeter;
    int x, y, dir;
    
    vector<int> dx = {1, 0, -1, 0};
    vector<int> dy = {0, 1, 0, -1};
    vector<string> dirs = {"East", "North", "West", "South"};

    Robot(int width, int height) {
        this->width = width;
        this->height = height;
        this->perimeter = 2 * (width + height) - 4;
        
        x = 0;
        y = 0;
        dir = 0;
    }
    
    void step(int num) {
        num %= perimeter;
        
        if (num == 0) {
            num = perimeter;
        }

        while (num > 0) {
            int nx = x + dx[dir];
            int ny = y + dy[dir];

            if (nx < 0 || nx >= width || ny < 0 || ny >= height) {
                dir = (dir + 1) % 4;
                continue;
            }

            x = nx;
            y = ny;
            num--;
        }
    }
    
    vector<int> getPos() {
        return {x, y};
    }
    
    string getDir() {
        return dirs[dir];
    }
};
```

### Java

```java
class Robot {
    private int width, height, perimeter;
    private int x, y, dir;

    private int[] dx = {1, 0, -1, 0};
    private int[] dy = {0, 1, 0, -1};
    private String[] dirs = {"East", "North", "West", "South"};

    public Robot(int width, int height) {
        this.width = width;
        this.height = height;
        this.perimeter = 2 * (width + height) - 4;

        this.x = 0;
        this.y = 0;
        this.dir = 0;
    }

    public void step(int num) {
        num %= perimeter;

        if (num == 0) {
            num = perimeter;
        }

        while (num > 0) {
            int nx = x + dx[dir];
            int ny = y + dy[dir];

            if (nx < 0 || nx >= width || ny < 0 || ny >= height) {
                dir = (dir + 1) % 4;
                continue;
            }

            x = nx;
            y = ny;
            num--;
        }
    }

    public int[] getPos() {
        return new int[]{x, y};
    }

    public String getDir() {
        return dirs[dir];
    }
}
```

### JavaScript

```javascript
var Robot = function(width, height) {
    this.width = width;
    this.height = height;
    this.perimeter = 2 * (width + height) - 4;

    this.x = 0;
    this.y = 0;
    this.dir = 0;

    this.dx = [1, 0, -1, 0];
    this.dy = [0, 1, 0, -1];
    this.dirs = ["East", "North", "West", "South"];
};

Robot.prototype.step = function(num) {
    num %= this.perimeter;

    if (num === 0) {
        num = this.perimeter;
    }

    while (num > 0) {
        let nx = this.x + this.dx[this.dir];
        let ny = this.y + this.dy[this.dir];

        if (nx < 0 || nx >= this.width || ny < 0 || ny >= this.height) {
            this.dir = (this.dir + 1) % 4;
            continue;
        }

        this.x = nx;
        this.y = ny;
        num--;
    }
};

Robot.prototype.getPos = function() {
    return [this.x, this.y];
};

Robot.prototype.getDir = function() {
    return this.dirs[this.dir];
};
```

### Python3

```python
class Robot:

    def __init__(self, width: int, height: int):
        self.width = width
        self.height = height
        self.perimeter = 2 * (width + height) - 4

        self.x = 0
        self.y = 0
        self.dir = 0

        self.dx = [1, 0, -1, 0]
        self.dy = [0, 1, 0, -1]
        self.dirs = ["East", "North", "West", "South"]

    def step(self, num: int) -> None:
        num %= self.perimeter

        if num == 0:
            num = self.perimeter

        while num > 0:
            nx = self.x + self.dx[self.dir]
            ny = self.y + self.dy[self.dir]

            if nx < 0 or nx >= self.width or ny < 0 or ny >= self.height:
                self.dir = (self.dir + 1) % 4
                continue

            self.x = nx
            self.y = ny
            num -= 1

    def getPos(self) -> List[int]:
        return [self.x, self.y]

    def getDir(self) -> str:
        return self.dirs[self.dir]
```

### Go

```go
type Robot struct {
    width, height, perimeter int
    x, y, dir int
    dx []int
    dy []int
    dirs []string
}

func Constructor(width int, height int) Robot {
    return Robot{
        width: width,
        height: height,
        perimeter: 2 * (width + height) - 4,
        x: 0,
        y: 0,
        dir: 0,
        dx: []int{1, 0, -1, 0},
        dy: []int{0, 1, 0, -1},
        dirs: []string{"East", "North", "West", "South"},
    }
}

func (this *Robot) Step(num int) {
    num %= this.perimeter

    if num == 0 {
        num = this.perimeter
    }

    for num > 0 {
        nx := this.x + this.dx[this.dir]
        ny := this.y + this.dy[this.dir]

        if nx < 0 || nx >= this.width || ny < 0 || ny >= this.height {
            this.dir = (this.dir + 1) % 4
            continue
        }

        this.x = nx
        this.y = ny
        num--
    }
}

func (this *Robot) GetPos() []int {
    return []int{this.x, this.y}
}

func (this *Robot) GetDir() string {
    return this.dirs[this.dir]
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

### 1. Store Current State

I store:

* Width and height
* Current position `(x, y)`
* Current direction `dir`
* Perimeter value

---

### 2. Use Direction Arrays

```text
East  -> (1, 0)
North -> (0, 1)
West  -> (-1, 0)
South -> (0, -1)
```

These arrays help me move quickly without writing many if-else conditions.

---

### 3. Reduce Large Moves

```text
num %= perimeter
```

This avoids unnecessary simulation.

If the robot makes one full cycle, it comes back to the same place.

---

### 4. Handle Full Cycle Carefully

If `num % perimeter == 0`, I still need to simulate one full perimeter because the robot direction changes.

Example:

* Start at `(0,0)` facing East
* After one full cycle, robot is again at `(0,0)`
* But now it faces South

---

### 5. Rotate When Move Is Invalid

If the next cell is outside the grid:

```text
dir = (dir + 1) % 4
```

This rotates the robot counterclockwise.

---

### 6. Move to Valid Cell

After finding a valid next position:

```text
x = nx
y = ny
```

Then reduce remaining moves.

---

## Examples

```text
Input:
Robot robot = new Robot(6, 3)
robot.step(2)
robot.step(2)
robot.getPos()
robot.getDir()

Output:
[4, 0]
East
```

Explanation:

* First 2 steps: `(0,0)` -> `(2,0)`
* Next 2 steps: `(2,0)` -> `(4,0)`
* Position becomes `[4,0]`
* Direction remains `East`

---

## How to use / Run locally

### C++

```bash
g++ main.cpp -o main
./main
```

### Java

```bash
javac Main.java
java Main
```

### JavaScript

```bash
node main.js
```

### Python3

```bash
python main.py
```

### Go

```bash
go run main.go
```

---

## Notes & Optimizations

* The biggest optimization is using modulo with perimeter.
* Without modulo, large values like `10^5` would be slow.
* Since width and height are small, simulating only the remaining steps is fast.
* The robot always stays on the boundary.
* Extra space usage is constant.

---

## Author

* [Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
