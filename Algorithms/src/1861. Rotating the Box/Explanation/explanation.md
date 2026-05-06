# 1861. Rotating the Box

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

---

## Problem Summary

LeetCode 1861 - Rotating the Box is a matrix simulation problem where we are given a 2D box containing:

* Stones `#`
* Obstacles `*`
* Empty spaces `.`

The box is rotated 90 degrees clockwise. After rotation, gravity pulls all stones downward.

The main challenge is figuring out how stones move while respecting obstacles and already settled stones.

The final task is to return the rotated matrix after all stones have fallen into their correct positions.

This problem is a great example of:

* Matrix transformation
* Simulation algorithms
* Two-pointer technique
* Grid-based problem solving

---

## Constraints

| Constraint               | Value                         |
| ------------------------ | ----------------------------- |
| `m == boxGrid.length`    | Number of rows                |
| `n == boxGrid[i].length` | Number of columns             |
| `1 <= m, n <= 500`       | Matrix size                   |
| `boxGrid[i][j]`          | Either `'#'`, `'*'`, or `'.'` |

---

## Intuition

The first thing I noticed was that rotating the box changes the direction of gravity.

After rotating clockwise:

* Right side becomes downward
* So stones that move right before rotation will behave exactly like falling stones afterward

That observation simplifies the entire problem.

Instead of rotating first and then simulating gravity, I realized I could:

1. Push all stones to the right inside each row
2. Rotate the matrix once at the end

Obstacles naturally split each row into smaller independent sections.

This makes the simulation much cleaner and easier to implement.

---

## Approach

I solved the problem in two major steps.

### Step 1: Simulate Gravity in Each Row

I traversed every row from right to left.

I kept a pointer that tracked the next available empty position where a stone could fall.

While scanning:

* If I found an obstacle `*`, I reset the pointer
* If I found a stone `#`, I moved it to the correct empty position
* Empty cells were skipped

This simulates gravity horizontally before rotation.

---

### Step 2: Rotate the Matrix

Once all rows were stabilized:

* I created a new matrix of size `n x m`
* Then I applied standard clockwise matrix rotation

Formula:

```text
rotated[j][m - 1 - i] = boxGrid[i][j]
```

This generates the final rotated box.

---

## Data Structures Used

| Data Structure  | Purpose                                           |
| --------------- | ------------------------------------------------- |
| 2D Matrix       | Stores the original box                           |
| Rotated Matrix  | Stores the final rotated answer                   |
| Integer Pointer | Tracks the next valid falling position for stones |

I only used simple matrix operations and pointer movement, which keeps the solution efficient and beginner-friendly.

---

## Operations & Behavior Summary

Here is what the algorithm does step by step:

1. Start processing each row independently
2. Move from right to left
3. Track the next empty position
4. Handle obstacles by resetting movement boundaries
5. Move stones toward the right
6. Finish gravity simulation for all rows
7. Create a rotated matrix
8. Rotate every cell into its new position
9. Return the final matrix

This approach avoids unnecessary repeated falling simulations.

---

## Complexity

| Type             | Complexity |
| ---------------- | ---------- |
| Time Complexity  | `O(m × n)` |
| Space Complexity | `O(m × n)` |

### Time Complexity

I traverse the matrix:

* Once for gravity simulation
* Once for matrix rotation

So total work stays linear relative to the number of cells.

---

### Space Complexity

I use an additional rotated matrix of size `n x m`.

Other than that, only a few variables are used.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    vector<vector<char>> rotateTheBox(vector<vector<char>>& boxGrid) {
        int m = boxGrid.size();
        int n = boxGrid[0].size();

        // Process every row independently
        for (int row = 0; row < m; row++) {

            // This points to the rightmost empty position
            // where the next stone can fall
            int emptyCol = n - 1;

            // Traverse from right to left
            for (int col = n - 1; col >= 0; col--) {

                // Obstacle blocks movement
                if (boxGrid[row][col] == '*') {

                    // Stones can only fall before obstacle
                    emptyCol = col - 1;
                }

                // Found a stone
                else if (boxGrid[row][col] == '#') {

                    // Remove stone from current position
                    boxGrid[row][col] = '.';

                    // Put stone at the valid empty position
                    boxGrid[row][emptyCol] = '#';

                    // Next stone should go one step left
                    emptyCol--;
                }
            }
        }

        // Create rotated matrix
        vector<vector<char>> rotated(n, vector<char>(m));

        // Rotate clockwise
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {

                // Standard clockwise rotation formula
                rotated[j][m - 1 - i] = boxGrid[i][j];
            }
        }

        return rotated;
    }
};
```

### Java

```java
class Solution {
    public char[][] rotateTheBox(char[][] boxGrid) {

        int m = boxGrid.length;
        int n = boxGrid[0].length;

        // Process every row
        for (int row = 0; row < m; row++) {

            // Rightmost empty position
            int emptyCol = n - 1;

            // Traverse row from right to left
            for (int col = n - 1; col >= 0; col--) {

                // Obstacle found
                if (boxGrid[row][col] == '*') {

                    // Reset valid falling position
                    emptyCol = col - 1;
                }

                // Stone found
                else if (boxGrid[row][col] == '#') {

                    // Remove stone
                    boxGrid[row][col] = '.';

                    // Move stone to valid position
                    boxGrid[row][emptyCol] = '#';

                    // Update next empty spot
                    emptyCol--;
                }
            }
        }

        // Rotated matrix
        char[][] rotated = new char[n][m];

        // Rotate clockwise
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {

                rotated[j][m - 1 - i] = boxGrid[i][j];
            }
        }

        return rotated;
    }
}
```

### JavaScript

```javascript
/**
 * @param {character[][]} boxGrid
 * @return {character[][]}
 */
var rotateTheBox = function(boxGrid) {

    const m = boxGrid.length;
    const n = boxGrid[0].length;

    // Process each row
    for (let row = 0; row < m; row++) {

        // Rightmost valid empty position
        let emptyCol = n - 1;

        // Traverse from right to left
        for (let col = n - 1; col >= 0; col--) {

            // Obstacle found
            if (boxGrid[row][col] === '*') {

                // Reset empty position
                emptyCol = col - 1;
            }

            // Stone found
            else if (boxGrid[row][col] === '#') {

                // Remove current stone
                boxGrid[row][col] = '.';

                // Move stone
                boxGrid[row][emptyCol] = '#';

                // Update next empty spot
                emptyCol--;
            }
        }
    }

    // Create rotated matrix
    const rotated = Array.from(
        { length: n },
        () => Array(m).fill('.')
    );

    // Rotate clockwise
    for (let i = 0; i < m; i++) {
        for (let j = 0; j < n; j++) {

            rotated[j][m - 1 - i] = boxGrid[i][j];
        }
    }

    return rotated;
};
```

### Python3

```python
class Solution:
    def rotateTheBox(self, boxGrid: List[List[str]]) -> List[List[str]]:

        m = len(boxGrid)
        n = len(boxGrid[0])

        # Process every row
        for row in range(m):

            # Rightmost empty position
            empty_col = n - 1

            # Traverse from right to left
            for col in range(n - 1, -1, -1):

                # Obstacle blocks stones
                if boxGrid[row][col] == '*':

                    # Reset valid position
                    empty_col = col - 1

                # Stone found
                elif boxGrid[row][col] == '#':

                    # Remove stone from current position
                    boxGrid[row][col] = '.'

                    # Move stone to valid position
                    boxGrid[row][empty_col] = '#'

                    # Next stone goes left
                    empty_col -= 1

        # Create rotated matrix
        rotated = [['.'] * m for _ in range(n)]

        # Rotate clockwise
        for i in range(m):
            for j in range(n):

                rotated[j][m - 1 - i] = boxGrid[i][j]

        return rotated
```

### Go

```go
func rotateTheBox(boxGrid [][]byte) [][]byte {

    m := len(boxGrid)
    n := len(boxGrid[0])

    // Process every row
    for row := 0; row < m; row++ {

        // Rightmost empty position
        emptyCol := n - 1

        // Traverse from right to left
        for col := n - 1; col >= 0; col-- {

            // Obstacle found
            if boxGrid[row][col] == '*' {

                // Reset valid position
                emptyCol = col - 1

            // Stone found
            } else if boxGrid[row][col] == '#' {

                // Remove current stone
                boxGrid[row][col] = '.'

                // Move stone
                boxGrid[row][emptyCol] = '#'

                // Update next empty position
                emptyCol--
            }
        }
    }

    // Create rotated matrix
    rotated := make([][]byte, n)

    for i := 0; i < n; i++ {
        rotated[i] = make([]byte, m)
    }

    // Rotate clockwise
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {

            rotated[j][m-1-i] = boxGrid[i][j]
        }
    }

    return rotated
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic remains exactly the same across all five languages.

Only syntax changes.

---

### Processing Rows

The most important part of the solution is simulating gravity before rotation.

I process one row at a time because stones only interact with cells inside their own segment.

---

### Why Traverse from Right to Left?

Stones move toward the right before rotation.

So scanning backward makes movement easier.

If I scanned from left to right:

* Stones could overwrite positions
* Tracking valid movement would become messy

Right-to-left traversal keeps everything clean.

---

### Role of the Pointer

I maintain a pointer representing:

* The next available empty position
* The farthest location where a stone can fall

Whenever a stone appears:

* It gets moved to that pointer position
* Then the pointer shifts left

This simulates gravity efficiently.

---

### Handling Obstacles

Obstacles completely block stone movement.

When an obstacle appears:

* Stones cannot cross it
* The pointer resets just before the obstacle

This divides each row into independent sections automatically.

---

### Matrix Rotation

After gravity simulation:

* The box is stable
* Only rotation remains

For clockwise rotation:

```text
(i, j) → (j, m - 1 - i)
```

This transforms rows into columns correctly.

---

### Why This Solution Is Efficient

Some brute-force solutions repeatedly drop stones cell by cell.

That works, but becomes slower and harder to manage.

This optimized matrix simulation:

* Processes every cell only a constant number of times
* Avoids nested gravity loops
* Uses clean pointer logic

That is why it performs well even for large constraints.

---

## Examples

### Example 1

### Input

```text
boxGrid = [["#",".","*","."]]
```

### Output

```text
[
 [ "." ],
 [ "#" ],
 [ "*" ],
 [ "." ]
]
```

### Explanation

* The stone moves right until blocked
* The box rotates clockwise
* Final gravity-adjusted state is returned

---

### Example 2

### Input

```text
boxGrid =
[
 ["#",".","*","."],
 ["#","#","*","."]
]
```

### Output

```text
[
 [ "#","." ],
 [ "#","#" ],
 [ "*","*" ],
 [ ".","." ]
]
```

### Explanation

* Stones slide right inside each row
* Obstacles stop movement
* Matrix is rotated afterward

---

### Example 3

### Input

```text
boxGrid =
[
 ["#","#","*",".","*","."],
 ["#","#","#","*",".","."],
 ["#","#","#",".","#","."]
]
```

### Output

```text
[
 [".","#","#"],
 [".","#","#"],
 ["#","#","*"],
 ["#","*","."],
 ["#",".","*"],
 ["#",".","."]
]
```

### Explanation

Each row stabilizes independently before rotation.

This avoids complicated falling simulation after rotating.

---

## How to Use / Run Locally

### C++

Compile:

```bash
g++ main.cpp -o main
```

Run:

```bash
./main
```

---

### Java

Compile:

```bash
javac Main.java
```

Run:

```bash
java Main
```

---

### JavaScript

Run using Node.js:

```bash
node main.js
```

---

### Python3

Run:

```bash
python main.py
```

---

### Go

Run:

```bash
go run main.go
```

---

## Notes & Optimizations

* Obstacles create natural boundaries between stone segments
* Simulating gravity before rotation simplifies the implementation
* The two-pointer approach avoids repeated falling checks
* This is an optimized matrix simulation solution
* Works efficiently for the maximum constraint size of `500 x 500`

### Alternative Approach

Another method is:

1. Rotate first
2. Simulate downward gravity afterward

That also works, but usually requires more careful handling and extra loops.

The current approach is cleaner and easier to debug.

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
