# 1914. Cyclically Rotating a Grid

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
* [How to Use / Run Locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

---

## Problem Summary

LeetCode 1914, "Cyclically Rotating a Grid", is a matrix simulation problem where every layer of a 2D grid must be rotated counter-clockwise exactly `k` times.

Instead of rotating the entire matrix together, each rectangular layer rotates independently.

The task is to:

* Take a 2D integer matrix
* Rotate every layer cyclically
* Return the final rotated grid

This problem is a good mix of:

* Matrix traversal
* Simulation
* Array rotation
* Layer-by-layer processing

The biggest challenge is handling rotations efficiently when `k` becomes very large.

---

## Constraints

| Constraint                | Value             |
| ------------------------- | ----------------- |
| `m == grid.length`        | Number of rows    |
| `n == grid[i].length`     | Number of columns |
| `2 <= m, n <= 50`         | Grid dimensions   |
| `m` and `n` are even      | Always even-sized |
| `1 <= grid[i][j] <= 5000` | Cell values       |
| `1 <= k <= 10^9`          | Rotation count    |

---

## Intuition

The first thing I noticed was that the grid is made of multiple layers.

The outer boundary forms one layer.
Then inside that, another smaller rectangle forms another layer.

So instead of rotating the matrix directly, I treated every layer like a circular array.

Once I converted a layer into a linear array, the problem became much easier:

1. Extract layer elements
2. Rotate the array
3. Put the values back into the matrix

I also realized that rotating one step at a time would be very slow when `k` is huge.

So I used:

```text
k % layerSize
```

This removes unnecessary full rotations and keeps the solution efficient.

---

## Approach

I solved the problem layer by layer.

### Step 1: Find total layers

The number of layers is:

```text
min(rows, cols) / 2
```

---

### Step 2: Extract one layer

For every layer, I traversed:

* Top row
* Right column
* Bottom row
* Left column

and stored all values in an array.

---

### Step 3: Rotate the extracted array

Instead of rotating repeatedly, I calculated:

```text
effectiveRotation = k % arrayLength
```

Then I performed a left rotation.

---

### Step 4: Put rotated values back

Using the same traversal order, I placed the rotated values back into the grid.

Repeating this process for all layers gives the final rotated matrix.

---

## Data Structures Used

| Data Structure       | Why I Used It                             |
| -------------------- | ----------------------------------------- |
| 2D Matrix            | Stores the original grid                  |
| Dynamic Array / List | Temporarily stores one layer for rotation |
| Index Variables      | Helps track boundaries of each layer      |

The temporary array makes rotation logic much cleaner and easier to manage.

---

## Operations & Behavior Summary

Here’s what the algorithm does internally:

1. Calculate total layers in the matrix
2. Pick one layer
3. Traverse its boundary
4. Store all values in an array
5. Compute effective rotations using modulo
6. Rotate the array
7. Reinsert rotated values into the same layer
8. Move to the next inner layer
9. Return the final rotated grid

This works because every layer behaves like an independent circular ring.

---

## Complexity

| Type             | Complexity | Explanation                                               |
| ---------------- | ---------- | --------------------------------------------------------- |
| Time Complexity  | `O(m × n)` | Every cell is visited a constant number of times          |
| Space Complexity | `O(m × n)` | Extra arrays are used to store layer elements temporarily |

Where:

* `m` = number of rows
* `n` = number of columns

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    vector<vector<int>> rotateGrid(vector<vector<int>>& grid, int k) {
        int m = grid.size();
        int n = grid[0].size();

        // Number of layers in the matrix
        int layers = min(m, n) / 2;

        // Process every layer separately
        for (int layer = 0; layer < layers; layer++) {

            vector<int> nums;

            int top = layer;
            int bottom = m - layer - 1;
            int left = layer;
            int right = n - layer - 1;

            // Store top row
            for (int j = left; j <= right; j++) {
                nums.push_back(grid[top][j]);
            }

            // Store right column
            for (int i = top + 1; i <= bottom - 1; i++) {
                nums.push_back(grid[i][right]);
            }

            // Store bottom row
            for (int j = right; j >= left; j--) {
                nums.push_back(grid[bottom][j]);
            }

            // Store left column
            for (int i = bottom - 1; i >= top + 1; i--) {
                nums.push_back(grid[i][left]);
            }

            int len = nums.size();

            // Remove extra full rotations
            int rotate = k % len;

            // Rotated version of current layer
            vector<int> rotated(len);

            // Left rotation
            for (int i = 0; i < len; i++) {
                rotated[i] = nums[(i + rotate) % len];
            }

            int idx = 0;

            // Put values back into top row
            for (int j = left; j <= right; j++) {
                grid[top][j] = rotated[idx++];
            }

            // Put values back into right column
            for (int i = top + 1; i <= bottom - 1; i++) {
                grid[i][right] = rotated[idx++];
            }

            // Put values back into bottom row
            for (int j = right; j >= left; j--) {
                grid[bottom][j] = rotated[idx++];
            }

            // Put values back into left column
            for (int i = bottom - 1; i >= top + 1; i--) {
                grid[i][left] = rotated[idx++];
            }
        }

        return grid;
    }
};
```

### Java

```java
class Solution {
    public int[][] rotateGrid(int[][] grid, int k) {

        int m = grid.length;
        int n = grid[0].length;

        // Total layers
        int layers = Math.min(m, n) / 2;

        for (int layer = 0; layer < layers; layer++) {

            ArrayList<Integer> nums = new ArrayList<>();

            int top = layer;
            int bottom = m - layer - 1;
            int left = layer;
            int right = n - layer - 1;

            // Store top row
            for (int j = left; j <= right; j++) {
                nums.add(grid[top][j]);
            }

            // Store right column
            for (int i = top + 1; i <= bottom - 1; i++) {
                nums.add(grid[i][right]);
            }

            // Store bottom row
            for (int j = right; j >= left; j--) {
                nums.add(grid[bottom][j]);
            }

            // Store left column
            for (int i = bottom - 1; i >= top + 1; i--) {
                nums.add(grid[i][left]);
            }

            int len = nums.size();

            // Effective rotations
            int rotate = k % len;

            int[] rotated = new int[len];

            // Left rotation
            for (int i = 0; i < len; i++) {
                rotated[i] = nums.get((i + rotate) % len);
            }

            int idx = 0;

            // Fill top row
            for (int j = left; j <= right; j++) {
                grid[top][j] = rotated[idx++];
            }

            // Fill right column
            for (int i = top + 1; i <= bottom - 1; i++) {
                grid[i][right] = rotated[idx++];
            }

            // Fill bottom row
            for (int j = right; j >= left; j--) {
                grid[bottom][j] = rotated[idx++];
            }

            // Fill left column
            for (int i = bottom - 1; i >= top + 1; i--) {
                grid[i][left] = rotated[idx++];
            }
        }

        return grid;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[][]} grid
 * @param {number} k
 * @return {number[][]}
 */
var rotateGrid = function(grid, k) {

    let m = grid.length;
    let n = grid[0].length;

    // Number of layers
    let layers = Math.min(m, n) / 2;

    for (let layer = 0; layer < layers; layer++) {

        let nums = [];

        let top = layer;
        let bottom = m - layer - 1;
        let left = layer;
        let right = n - layer - 1;

        // Store top row
        for (let j = left; j <= right; j++) {
            nums.push(grid[top][j]);
        }

        // Store right column
        for (let i = top + 1; i <= bottom - 1; i++) {
            nums.push(grid[i][right]);
        }

        // Store bottom row
        for (let j = right; j >= left; j--) {
            nums.push(grid[bottom][j]);
        }

        // Store left column
        for (let i = bottom - 1; i >= top + 1; i--) {
            nums.push(grid[i][left]);
        }

        let len = nums.length;

        // Effective rotations only
        let rotate = k % len;

        let rotated = new Array(len);

        // Left rotation
        for (let i = 0; i < len; i++) {
            rotated[i] = nums[(i + rotate) % len];
        }

        let idx = 0;

        // Fill top row
        for (let j = left; j <= right; j++) {
            grid[top][j] = rotated[idx++];
        }

        // Fill right column
        for (let i = top + 1; i <= bottom - 1; i++) {
            grid[i][right] = rotated[idx++];
        }

        // Fill bottom row
        for (let j = right; j >= left; j--) {
            grid[bottom][j] = rotated[idx++];
        }

        // Fill left column
        for (let i = bottom - 1; i >= top + 1; i--) {
            grid[i][left] = rotated[idx++];
        }
    }

    return grid;
};
```

### Python3

```python
class Solution:
    def rotateGrid(self, grid: List[List[int]], k: int) -> List[List[int]]:

        m = len(grid)
        n = len(grid[0])

        # Total layers inside matrix
        layers = min(m, n) // 2

        for layer in range(layers):

            nums = []

            top = layer
            bottom = m - layer - 1
            left = layer
            right = n - layer - 1

            # Store top row
            for j in range(left, right + 1):
                nums.append(grid[top][j])

            # Store right column
            for i in range(top + 1, bottom):
                nums.append(grid[i][right])

            # Store bottom row
            for j in range(right, left - 1, -1):
                nums.append(grid[bottom][j])

            # Store left column
            for i in range(bottom - 1, top, -1):
                nums.append(grid[i][left])

            length = len(nums)

            # Ignore unnecessary full rotations
            rotate = k % length

            # Left rotated array
            rotated = nums[rotate:] + nums[:rotate]

            idx = 0

            # Fill top row
            for j in range(left, right + 1):
                grid[top][j] = rotated[idx]
                idx += 1

            # Fill right column
            for i in range(top + 1, bottom):
                grid[i][right] = rotated[idx]
                idx += 1

            # Fill bottom row
            for j in range(right, left - 1, -1):
                grid[bottom][j] = rotated[idx]
                idx += 1

            # Fill left column
            for i in range(bottom - 1, top, -1):
                grid[i][left] = rotated[idx]
                idx += 1

        return grid
```

### Go

```go
func rotateGrid(grid [][]int, k int) [][]int {

    m := len(grid)
    n := len(grid[0])

    // Total layers in matrix
    layers := min(m, n) / 2

    for layer := 0; layer < layers; layer++ {

        nums := []int{}

        top := layer
        bottom := m - layer - 1
        left := layer
        right := n - layer - 1

        // Store top row
        for j := left; j <= right; j++ {
            nums = append(nums, grid[top][j])
        }

        // Store right column
        for i := top + 1; i <= bottom-1; i++ {
            nums = append(nums, grid[i][right])
        }

        // Store bottom row
        for j := right; j >= left; j-- {
            nums = append(nums, grid[bottom][j])
        }

        // Store left column
        for i := bottom - 1; i >= top+1; i-- {
            nums = append(nums, grid[i][left])
        }

        length := len(nums)

        // Effective rotations only
        rotate := k % length

        rotated := make([]int, length)

        // Left rotation
        for i := 0; i < length; i++ {
            rotated[i] = nums[(i+rotate)%length]
        }

        idx := 0

        // Fill top row
        for j := left; j <= right; j++ {
            grid[top][j] = rotated[idx]
            idx++
        }

        // Fill right column
        for i := top + 1; i <= bottom-1; i++ {
            grid[i][right] = rotated[idx]
            idx++
        }

        // Fill bottom row
        for j := right; j >= left; j-- {
            grid[bottom][j] = rotated[idx]
            idx++
        }

        // Fill left column
        for i := bottom - 1; i >= top+1; i-- {
            grid[i][left] = rotated[idx]
            idx++
        }
    }

    return grid
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic stays exactly the same in all five languages.

Only syntax changes.

### 1. Finding layer boundaries

For every layer, four boundaries are maintained:

* `top`
* `bottom`
* `left`
* `right`

These boundaries tell us which cells belong to the current ring.

Without proper boundaries, elements from different layers would mix together.

---

### 2. Extracting the layer

The boundary is traversed in a fixed order:

1. Top row
2. Right column
3. Bottom row
4. Left column

This creates a linear representation of the layer.

Corner cells are carefully handled so they are not added twice.

---

### 3. Reducing unnecessary rotations

If a layer contains 20 elements and `k = 100`, rotating 100 times is pointless.

Because:

```text
100 % 20 = 0
```

the layer returns to the same state.

So modulo helps avoid wasted work.

---

### 4. Rotating the array

The extracted layer behaves exactly like a circular array.

A left rotation shifts elements forward while wrapping around.

Example:

```text
Original: [1,2,3,4,5]
Rotate by 2
Result:   [3,4,5,1,2]
```

This is much simpler than rotating directly inside the matrix.

---

### 5. Rebuilding the matrix

After rotation, values are inserted back using the same traversal order.

This is very important.

If insertion order differs from extraction order, elements would go into wrong positions.

---

### 6. Processing inner layers

Once one layer is finished, the boundaries shrink inward.

Then the same process repeats for the next layer.

Eventually all layers become rotated.

---

## Examples

### Example 1

Input:

```text
grid = [[40,10],
        [30,20]]

k = 1
```

Output:

```text
[[10,20],
 [40,30]]
```

Explanation:

The grid has only one layer.

Layer values:

```text
[40,10,20,30]
```

After one counter-clockwise rotation:

```text
[10,20,30,40]
```

These values are placed back into the matrix.

---

### Example 2

Input:

```text
grid = [[1,2,3,4],
        [5,6,7,8],
        [9,10,11,12],
        [13,14,15,16]]

k = 2
```

Output:

```text
[[3,4,8,12],
 [2,11,10,16],
 [1,7,6,15],
 [5,9,13,14]]
```

Explanation:

The matrix contains:

* Outer layer
* Inner layer

Both layers rotate independently.

---

### Example 3

Input:

```text
grid = [[1,2],
        [3,4],
        [5,6],
        [7,8]]

k = 3
```

Output:

```text
[[7,5],
 [8,3],
 [6,1],
 [4,2]]
```

Explanation:

The single layer rotates three times counter-clockwise.

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

* Using modulo rotation is extremely important for performance.
* Rotating directly inside the matrix would make the logic much more complicated.
* Treating each layer as a circular array keeps the implementation clean.
* This approach works efficiently even for very large `k`.
* Since constraints are small (`<= 50`), memory usage is completely safe.
* Another possible approach is rotating elements in-place, but that usually becomes harder to debug and maintain.

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
