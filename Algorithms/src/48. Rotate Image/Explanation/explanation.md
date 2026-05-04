# Rotate Image (LeetCode 48) – In-Place Matrix Rotation Solution

## Table of Contents

* [Problem Summary](#problem-summary)
* [Constraints](#constraints)
* [Intuition](#intuition)
* [Approach](#approach)
* [Data Structures Used](#data-structures-used)
* [Operations & Behavior Summary](#operations--behavior-summary)
* [Complexity](#complexity)
* [Multi-language Solutions](#multi-language-solutions)
* [Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)](#step-by-step-detailed-explanation-c-java-javascript-python3-go)
* [Examples](#examples)
* [How to Use / Run Locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

## Problem Summary

The problem asks me to rotate a square matrix (n x n) by 90 degrees clockwise.

The key condition is that the rotation must be done in-place. That means I cannot create a new matrix to store the result. I have to modify the given matrix directly.

Input:

* A 2D square matrix

Output:

* The same matrix, rotated 90 degrees clockwise

This is a classic matrix manipulation problem often asked in coding interviews and DSA practice.

## Constraints

* n == matrix.length == matrix[i].length
* 1 <= n <= 20
* -1000 <= matrix[i][j] <= 1000

## Intuition

When I first looked at this problem, I tried to visualize how elements move after rotation.

I noticed that:

* The first row becomes the last column
* The second row becomes the second last column

At first, I thought about creating a new matrix and placing elements correctly, but that violates the in-place requirement.

Then I realized something important:
If I transpose the matrix and then reverse each row, I get exactly the rotated result.

That insight simplifies everything.

## Approach

I break the solution into two clear steps.

Step 1: Transpose the matrix
Swap elements across the diagonal.
This converts rows into columns.

Step 2: Reverse each row
After transpose, each row is reversed to complete the 90-degree clockwise rotation.

These two operations together give the final rotated matrix without using extra space.

## Data Structures Used

* 2D Array (Matrix)
  I directly modify the given matrix. No extra data structure is used because the problem requires an in-place solution.

## Operations & Behavior Summary

Here’s what the algorithm does:

1. Loop through the matrix
2. Swap matrix[i][j] with matrix[j][i] for i < j (transpose step)
3. After transpose, iterate through each row
4. Reverse every row
5. The matrix is now rotated 90 degrees clockwise

## Complexity

| Type             | Complexity | Explanation                                                                  |
| ---------------- | ---------- | ---------------------------------------------------------------------------- |
| Time Complexity  | O(n²)      | I traverse the matrix twice — once for transpose and once for reversing rows |
| Space Complexity | O(1)       | No extra memory is used; everything is done in-place                         |

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    void rotate(vector<vector<int>>& matrix) {
        int n = matrix.size();

        // Step 1: Transpose the matrix
        // Swap elements across diagonal
        for(int i = 0; i < n; i++) {
            for(int j = i + 1; j < n; j++) {
                swap(matrix[i][j], matrix[j][i]);
            }
        }

        // Step 2: Reverse each row
        for(int i = 0; i < n; i++) {
            reverse(matrix[i].begin(), matrix[i].end());
        }
    }
};
```

### Java

```java
class Solution {
    public void rotate(int[][] matrix) {
        int n = matrix.length;

        // Step 1: Transpose
        for(int i = 0; i < n; i++) {
            for(int j = i + 1; j < n; j++) {
                int temp = matrix[i][j];
                matrix[i][j] = matrix[j][i];
                matrix[j][i] = temp;
            }
        }

        // Step 2: Reverse each row
        for(int i = 0; i < n; i++) {
            int left = 0, right = n - 1;
            while(left < right) {
                int temp = matrix[i][left];
                matrix[i][left] = matrix[i][right];
                matrix[i][right] = temp;
                left++;
                right--;
            }
        }
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[][]} matrix
 * @return {void} Do not return anything, modify matrix in-place instead.
 */
var rotate = function(matrix) {
    let n = matrix.length;

    // Step 1: Transpose
    for(let i = 0; i < n; i++) {
        for(let j = i + 1; j < n; j++) {
            let temp = matrix[i][j];
            matrix[i][j] = matrix[j][i];
            matrix[j][i] = temp;
        }
    }

    // Step 2: Reverse each row
    for(let i = 0; i < n; i++) {
        matrix[i].reverse();
    }
};
```

### Python3

```python
class Solution:
    def rotate(self, matrix: List[List[int]]) -> None:
        n = len(matrix)

        # Step 1: Transpose
        for i in range(n):
            for j in range(i + 1, n):
                # Swap across diagonal
                matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]

        # Step 2: Reverse each row
        for i in range(n):
            matrix[i].reverse()
```

### Go

```go
func rotate(matrix [][]int) {
    n := len(matrix)

    // Step 1: Transpose
    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            // Swap elements across diagonal
            matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
        }
    }

    // Step 2: Reverse each row
    for i := 0; i < n; i++ {
        left, right := 0, n-1
        for left < right {
            matrix[i][left], matrix[i][right] = matrix[i][right], matrix[i][left]
            left++
            right--
        }
    }
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

I start with the transpose step.

For every element above the diagonal (where column index is greater than row index), I swap it with its mirrored position. This ensures I don’t swap the same pair twice.

If I didn’t restrict to one half of the matrix, I would end up undoing my swaps.

After transpose, rows become columns.

Next, I reverse each row.

This step is important because transpose alone only flips the matrix over its diagonal. The order is still not correct for rotation.

Reversing each row aligns the elements in the correct rotated positions.

This logic works exactly the same across C++, Java, JavaScript, Python, and Go. The only difference is syntax, not behavior.

## Examples

Example 1:

Input:
matrix =
[
[1,2,3],
[4,5,6],
[7,8,9]
]

Output:
[
[7,4,1],
[8,5,2],
[9,6,3]
]

Explanation:

* Transpose →
  [ [1,4,7], [2,5,8], [3,6,9] ]
* Reverse rows →
  [ [7,4,1], [8,5,2], [9,6,3] ]

---

Example 2:

Input:
matrix =
[
[5,1,9,11],
[2,4,8,10],
[13,3,6,7],
[15,14,12,16]
]

Output:
[
[15,13,2,5],
[14,3,4,1],
[12,6,8,9],
[16,7,10,11]
]

---

Example 3:

Input:
matrix =
[
[1]
]

Output:
[
[1]
]

Explanation:
Single element remains the same after rotation.

## How to Use / Run Locally

C++:

1. Save code in a file like solution.cpp
2. Compile using: g++ solution.cpp -o output
3. Run: ./output

Java:

1. Save as Solution.java
2. Compile: javac Solution.java
3. Run: java Solution

JavaScript:

1. Save as solution.js
2. Run using Node.js: node solution.js

Python:

1. Save as solution.py
2. Run: python solution.py

Go:

1. Save as solution.go
2. Run: go run solution.go

## Notes & Optimizations

* The in-place constraint is the most important part of this problem
* Transpose + reverse is the cleanest and most optimal approach
* Another approach is layer-by-layer rotation, but it is more complex to implement
* Works for all square matrices, including edge cases like 1x1
* Always be careful with index handling to avoid incorrect swaps

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
