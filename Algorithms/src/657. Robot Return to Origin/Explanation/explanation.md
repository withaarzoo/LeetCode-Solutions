# 657. Robot Return to Origin

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

There is a robot starting at position `(0, 0)` on a 2D plane.

A string `moves` is given where:

* `'U'` means move up
* `'D'` means move down
* `'L'` means move left
* `'R'` means move right

I need to determine whether the robot returns back to the origin `(0, 0)` after performing all moves.

---

## Constraints

* `1 <= moves.length <= 2 * 10^4`
* `moves` only contains:

  * `'U'`
  * `'D'`
  * `'L'`
  * `'R'`

---

## Intuition

I thought about tracking the robot’s final position.

* If the robot moves up, I increase the y-coordinate.
* If the robot moves down, I decrease the y-coordinate.
* If the robot moves right, I increase the x-coordinate.
* If the robot moves left, I decrease the x-coordinate.

At the end, if both `x == 0` and `y == 0`, then the robot has returned to the origin.

---

## Approach

1. Start with `x = 0` and `y = 0`.
2. Traverse the string `moves` character by character.
3. Update `x` and `y` according to the move:

   * `'U'` → `y++`
   * `'D'` → `y--`
   * `'R'` → `x++`
   * `'L'` → `x--`
4. After processing all moves, check:

```text
x == 0 && y == 0
```

1. If true, return `true`.
2. Otherwise, return `false`.

---

## Data Structures Used

* Integer variables:

  * `x` for horizontal position
  * `y` for vertical position

No extra arrays, maps, sets, or additional data structures are required.

---

## Operations & Behavior Summary

| Move | Operation | Effect     |
| ---- | --------- | ---------- |
| `U`  | `y++`     | Move up    |
| `D`  | `y--`     | Move down  |
| `R`  | `x++`     | Move right |
| `L`  | `x--`     | Move left  |

At the end:

| Condition          | Result                   |
| ------------------ | ------------------------ |
| `x == 0 && y == 0` | Robot returned to origin |
| Otherwise          | Robot did not return     |

---

## Complexity

* Time Complexity: `O(n)`

  * `n` is the length of the `moves` string.
  * I process each move exactly once.

* Space Complexity: `O(1)`

  * I only use two integer variables.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    bool judgeCircle(string moves) {
        int x = 0, y = 0;

        for (char move : moves) {
            if (move == 'U') {
                y++;
            } else if (move == 'D') {
                y--;
            } else if (move == 'R') {
                x++;
            } else if (move == 'L') {
                x--;
            }
        }

        return x == 0 && y == 0;
    }
};
```

### Java

```java
class Solution {
    public boolean judgeCircle(String moves) {
        int x = 0, y = 0;

        for (char move : moves.toCharArray()) {
            if (move == 'U') {
                y++;
            } else if (move == 'D') {
                y--;
            } else if (move == 'R') {
                x++;
            } else if (move == 'L') {
                x--;
            }
        }

        return x == 0 && y == 0;
    }
}
```

### JavaScript

```javascript
/**
 * @param {string} moves
 * @return {boolean}
 */
var judgeCircle = function(moves) {
    let x = 0, y = 0;

    for (let move of moves) {
        if (move === 'U') {
            y++;
        } else if (move === 'D') {
            y--;
        } else if (move === 'R') {
            x++;
        } else if (move === 'L') {
            x--;
        }
    }

    return x === 0 && y === 0;
};
```

### Python3

```python
class Solution:
    def judgeCircle(self, moves: str) -> bool:
        x, y = 0, 0

        for move in moves:
            if move == 'U':
                y += 1
            elif move == 'D':
                y -= 1
            elif move == 'R':
                x += 1
            elif move == 'L':
                x -= 1

        return x == 0 and y == 0
```

### Go

```go
func judgeCircle(moves string) bool {
    x, y := 0, 0

    for _, move := range moves {
        if move == 'U' {
            y++
        } else if move == 'D' {
            y--
        } else if move == 'R' {
            x++
        } else if move == 'L' {
            x--
        }
    }

    return x == 0 && y == 0
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

### Step 1: Initialize Coordinates

```cpp
int x = 0, y = 0;
```

* `x` tracks left and right movement.
* `y` tracks up and down movement.
* Initially, the robot is at `(0, 0)`.

### Step 2: Traverse Every Move

```cpp
for (char move : moves)
```

I check every move one by one.

Example:

```text
moves = "UDLR"
```

The loop processes:

* `U`
* `D`
* `L`
* `R`

### Step 3: Update Position

```cpp
if (move == 'U') y++;
```

If the move is up, increase `y`.

```cpp
else if (move == 'D') y--;
```

If the move is down, decrease `y`.

```cpp
else if (move == 'R') x++;
```

If the move is right, increase `x`.

```cpp
else if (move == 'L') x--;
```

If the move is left, decrease `x`.

### Step 4: Check Final Position

```cpp
return x == 0 && y == 0;
```

If both `x` and `y` become `0`, then the robot came back to the origin.

---

## Examples

### Example 1

```text
Input: moves = "UD"
Output: true
```

Explanation:

* `U` → `(0, 1)`
* `D` → `(0, 0)`

The robot returns to the origin.

### Example 2

```text
Input: moves = "LL"
Output: false
```

Explanation:

* First `L` → `(-1, 0)`
* Second `L` → `(-2, 0)`

The robot does not return to the origin.

---

## How to use / Run locally

### C++

```bash
g++ main.cpp -o main
./main
```

### Java

```bash
javac Solution.java
java Solution
```

### JavaScript

```bash
node solution.js
```

### Python3

```bash
python solution.py
```

### Go

```bash
go run solution.go
```

---

## Notes & Optimizations

* I do not need extra memory because only two variables are enough.
* I process the string in a single pass.
* This is the most optimized solution possible for this problem.
* Since every move must be checked once, `O(n)` time complexity is optimal.

---

## Author

* [Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
