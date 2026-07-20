# Shift 2D Grid - LeetCode 1260 Solution

A clean and beginner-friendly solution for **LeetCode 1260 - Shift 2D Grid**. This repository explains the idea behind the algorithm, walks through the solution step by step, analyzes the time and space complexity, and provides implementations in **C++, Java, JavaScript, Python3, and Go**.

If you're preparing for coding interviews or improving your Data Structures and Algorithms (DSA) skills, this problem is a good example of working with **2D arrays, index mapping, matrix manipulation, and simulation**.

---

# Table of Contents

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

# Problem Summary

The problem gives us a 2D grid and an integer `k`.

One shift operation moves every element one position forward in row-major order.

* Every element moves one column to the right.
* The last element of a row moves to the beginning of the next row.
* The last element of the entire grid moves back to the first position.

Our task is to perform this operation exactly `k` times and return the final shifted grid.

Instead of simulating one shift at a time, we can directly calculate where every element should end up after all shifts. This makes the solution much faster and cleaner.

This problem is commonly asked in coding interviews because it tests your understanding of **2D arrays, index conversion, simulation, modular arithmetic, and array manipulation**.

---

# Constraints

| Constraint    | Value                         |
| ------------- | ----------------------------- |
| Rows (`m`)    | `1 <= m <= 50`                |
| Columns (`n`) | `1 <= n <= 50`                |
| Grid Value    | `-1000 <= grid[i][j] <= 1000` |
| Shifts (`k`)  | `0 <= k <= 100`               |

---

# Intuition

The first thing I noticed was that every element always moves to a fixed position after shifting.

Instead of moving elements one step at a time, I imagined the entire matrix as one long array written row by row.

Once I looked at it this way, the problem became much easier.

A grid shift is simply a right rotation of that flattened array.

So rather than performing multiple shifts, I only needed to calculate the final position of every element and place it there directly.

---

# Approach

I solved the problem in these steps.

1. Find the number of rows and columns.
2. Compute the total number of elements.
3. Reduce `k` using modulo because shifting by the total number of elements does not change the grid.
4. Create a new grid with the same dimensions.
5. Visit every element in the original grid.
6. Convert its row and column into a single index.
7. Calculate the new index after shifting.
8. Convert the new index back into row and column values.
9. Store the element in its new position.
10. Return the new grid.

This approach avoids unnecessary repeated shifting and processes every element only once.

---

# Data Structures Used

| Data Structure | Purpose                       |
| -------------- | ----------------------------- |
| 2D Array       | Stores the original grid      |
| 2D Array       | Stores the final shifted grid |

No additional complex data structures such as queues, stacks, or hash maps are required.

---

# Operations & Behavior Summary

The algorithm works in the following order.

1. Read the dimensions of the grid.
2. Count the total number of elements.
3. Remove unnecessary full rotations using modulo.
4. Visit every cell exactly once.
5. Convert each 2D position into a 1D index.
6. Compute the shifted index.
7. Convert the shifted index back into a 2D position.
8. Place the value into the answer grid.
9. Return the completed grid.

Since every element is handled only once, the solution stays efficient even for the largest allowed input.

---

# Complexity

| Complexity       | Value      | Explanation                                              |
| ---------------- | ---------- | -------------------------------------------------------- |
| Time Complexity  | `O(m × n)` | Every element in the grid is processed exactly one time. |
| Space Complexity | `O(m × n)` | A new grid is created to store the shifted result.       |

Where:

* `m` = number of rows
* `n` = number of columns

---

# Multi-language Solutions

### C++

```cpp
class Solution {
public:
    vector<vector<int>> shiftGrid(vector<vector<int>>& grid, int k) {
        // Get the grid dimensions
        int m = grid.size();
        int n = grid[0].size();

        // Total number of elements
        int total = m * n;

        // Extra full rotations do not change the grid
        k %= total;

        // Create the answer grid
        vector<vector<int>> ans(m, vector<int>(n));

        // Move every element directly to its final position
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {

                // Current position in the flattened array
                int oldIndex = i * n + j;

                // Position after shifting k times
                int newIndex = (oldIndex + k) % total;

                // Convert back to row and column
                int newRow = newIndex / n;
                int newCol = newIndex % n;

                // Place the element
                ans[newRow][newCol] = grid[i][j];
            }
        }

        // Return the shifted grid
        return ans;
    }
};
```

### Java

```java
class Solution {
    public List<List<Integer>> shiftGrid(int[][] grid, int k) {

        // Get the grid dimensions
        int m = grid.length;
        int n = grid[0].length;

        // Total number of elements
        int total = m * n;

        // Ignore unnecessary complete rotations
        k %= total;

        // Create the answer grid
        List<List<Integer>> ans = new ArrayList<>();

        for (int i = 0; i < m; i++) {
            List<Integer> row = new ArrayList<>();

            // Fill with dummy values so we can use set()
            for (int j = 0; j < n; j++) {
                row.add(0);
            }

            ans.add(row);
        }

        // Move every element to its final position
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {

                // Flatten the current position
                int oldIndex = i * n + j;

                // Compute shifted position
                int newIndex = (oldIndex + k) % total;

                // Convert back to row and column
                int newRow = newIndex / n;
                int newCol = newIndex % n;

                // Store the value
                ans.get(newRow).set(newCol, grid[i][j]);
            }
        }

        return ans;
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
var shiftGrid = function(grid, k) {

    // Get grid dimensions
    const m = grid.length;
    const n = grid[0].length;

    // Total elements
    const total = m * n;

    // Remove unnecessary complete rotations
    k %= total;

    // Create the answer grid
    const ans = Array.from({ length: m }, () => Array(n));

    // Move every element directly
    for (let i = 0; i < m; i++) {
        for (let j = 0; j < n; j++) {

            // Flatten the current position
            const oldIndex = i * n + j;

            // New position after shifting
            const newIndex = (oldIndex + k) % total;

            // Convert back to row and column
            const newRow = Math.floor(newIndex / n);
            const newCol = newIndex % n;

            // Place the value
            ans[newRow][newCol] = grid[i][j];
        }
    }

    return ans;
};
```

### Python3

```python
class Solution:
    def shiftGrid(self, grid: List[List[int]], k: int) -> List[List[int]]:

        # Get the grid dimensions
        m = len(grid)
        n = len(grid[0])

        # Total number of elements
        total = m * n

        # Ignore unnecessary complete rotations
        k %= total

        # Create the answer grid
        ans = [[0] * n for _ in range(m)]

        # Move every element to its final position
        for i in range(m):
            for j in range(n):

                # Flatten the current position
                old_index = i * n + j

                # Compute the shifted position
                new_index = (old_index + k) % total

                # Convert back to row and column
                new_row = new_index // n
                new_col = new_index % n

                # Place the value
                ans[new_row][new_col] = grid[i][j]

        return ans
```

### Go

```go
func shiftGrid(grid [][]int, k int) [][]int {

 // Get the grid dimensions
 m := len(grid)
 n := len(grid[0])

 // Total number of elements
 total := m * n

 // Ignore unnecessary complete rotations
 k %= total

 // Create the answer grid
 ans := make([][]int, m)
 for i := 0; i < m; i++ {
  ans[i] = make([]int, n)
 }

 // Move every element directly
 for i := 0; i < m; i++ {
  for j := 0; j < n; j++ {

   // Flatten the current position
   oldIndex := i*n + j

   // Compute the shifted position
   newIndex := (oldIndex + k) % total

   // Convert back to row and column
   newRow := newIndex / n
   newCol := newIndex % n

   // Place the value
   ans[newRow][newCol] = grid[i][j]
  }
 }

 return ans
}
```

---

# Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic is exactly the same in all five languages. Only the syntax changes.

First, the solution reads the dimensions of the grid.

These values are needed because we repeatedly convert between two-dimensional positions and one-dimensional positions.

Next, the total number of elements is calculated.

This lets us treat the grid as one continuous array.

The value of `k` is then reduced using modulo.

This avoids doing unnecessary work because shifting by the total number of elements always returns the grid to its original state.

A new grid is created.

Writing into a separate grid is important because changing the original grid while reading from it could overwrite values before they are moved.

The algorithm then visits every cell exactly once.

For each element, its row and column are converted into a single linear index.

That index represents the position of the element inside the imaginary flattened array.

The shifted position is calculated by adding `k` and wrapping around using modulo.

After finding the new index, it is converted back into a row and column.

The value is placed into that location in the answer grid.

Once every element has been processed, the new grid already contains the final answer.

Although the syntax differs slightly between C++, Java, JavaScript, Python3, and Go, the algorithm, calculations, and complexity remain exactly the same.

---

# Examples

## Example 1

**Input**

```text
grid = [[1,2,3],[4,5,6],[7,8,9]]
k = 1
```

**Output**

```text
[[9,1,2],[3,4,5],[6,7,8]]
```

**Trace**

The last element (`9`) moves to the beginning.

Every other element shifts one position to the right.

---

## Example 2

**Input**

```text
grid = [[3,8,1,9],[19,7,2,5],[4,6,11,10],[12,0,21,13]]
k = 4
```

**Output**

```text
[[12,0,21,13],[3,8,1,9],[19,7,2,5],[4,6,11,10]]
```

**Trace**

There are sixteen elements.

Shifting four positions moves the last four elements to the front while preserving the order of the remaining elements.

---

## Example 3

**Input**

```text
grid = [[1,2,3],[4,5,6],[7,8,9]]
k = 9
```

**Output**

```text
[[1,2,3],[4,5,6],[7,8,9]]
```

**Trace**

The total number of elements is nine.

A complete rotation returns every element to its original position.

---

# How to Use / Run Locally

Clone the repository.

```bash
git clone <repository-url>
```

Move into the project folder.

```bash
cd <repository-folder>
```

Compile and run the language you want.

### C++

Compile:

```bash
g++ solution.cpp -o solution
```

Run:

```bash
./solution
```

### Java

Compile:

```bash
javac Solution.java
```

Run:

```bash
java Solution
```

### JavaScript

Run:

```bash
node solution.js
```

### Python3

Run:

```bash
python solution.py
```

or

```bash
python3 solution.py
```

### Go

Run:

```bash
go run solution.go
```

---

# Notes & Optimizations

* Reducing `k` using modulo avoids unnecessary rotations.
* Every element is processed exactly once.
* The solution does not repeatedly simulate shifts.
* Mapping between 2D coordinates and a 1D index keeps the implementation simple.
* The algorithm works correctly for single-row and single-column grids.
* It also handles the case where `k` is zero or equal to the total number of elements.
* Another possible approach is to flatten the grid into a separate array, rotate it, and rebuild the matrix. That method is also valid but requires an extra flattening step.

---

# Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
