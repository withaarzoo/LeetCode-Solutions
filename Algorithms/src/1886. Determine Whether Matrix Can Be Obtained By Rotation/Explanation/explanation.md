# Problem Title

Determine Whether Matrix Can Be Obtained By Rotation

---

## Table of Contents

* Problem Summary
* Constraints
* Intuition
* Approach
* Data Structures Used
* Operations & Behavior Summary
* Complexity
* Multi-language Solutions
* Step-by-step Detailed Explanation
* Examples
* How to use / Run locally
* Notes & Optimizations
* Author

---

## Problem Summary

Given two n x n binary matrices `mat` and `target`, return `true` if it is possible to make `mat` equal to `target` by rotating `mat` in 90-degree increments, otherwise return `false`.

---

## Constraints

* n == mat.length == target.length
* n == mat[i].length == target[i].length
* 1 <= n <= 10
* mat[i][j] and target[i][j] are either 0 or 1

---

## Intuition

I thought about how many possible ways I can rotate a square matrix. A matrix can only be rotated 4 times before it comes back to its original form (0°, 90°, 180°, 270°).

So instead of trying complex transformations, I realized I can just:

* Check if the matrix matches the target
* If not, rotate it
* Repeat this process 4 times

If any rotation matches the target, the answer is true.

---

## Approach

1. Loop 4 times (for each possible rotation)
2. Compare `mat` with `target`
3. If equal, return true
4. Otherwise, rotate the matrix by 90 degrees clockwise
5. If no match after 4 rotations, return false

Rotation steps:

* Transpose the matrix
* Reverse each row

---

## Data Structures Used

* 2D Array (Matrix)
* No extra space required (in-place operations)

---

## Operations & Behavior Summary

* Matrix comparison: check all elements
* Matrix transpose: swap `mat[i][j]` with `mat[j][i]`
* Row reverse: reverse each row after transpose
* Perform above operations at most 4 times

---

## Complexity

* Time Complexity: O(n^2)

  * Each comparison and rotation takes O(n^2)
  * Done 4 times → O(4 * n^2) = O(n^2)

* Space Complexity: O(1)

  * No extra space used (in-place rotation)

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    void rotate(vector<vector<int>>& mat) {
        int n = mat.size();
        
        for(int i = 0; i < n; i++) {
            for(int j = i; j < n; j++) {
                swap(mat[i][j], mat[j][i]);
            }
        }
        
        for(int i = 0; i < n; i++) {
            reverse(mat[i].begin(), mat[i].end());
        }
    }

    bool findRotation(vector<vector<int>>& mat, vector<vector<int>>& target) {
        for(int k = 0; k < 4; k++) {
            if(mat == target) return true;
            rotate(mat);
        }
        return false;
    }
};
```

### Java

```java
class Solution {
    private void rotate(int[][] mat) {
        int n = mat.length;

        for(int i = 0; i < n; i++) {
            for(int j = i; j < n; j++) {
                int temp = mat[i][j];
                mat[i][j] = mat[j][i];
                mat[j][i] = temp;
            }
        }

        for(int i = 0; i < n; i++) {
            for(int j = 0; j < n / 2; j++) {
                int temp = mat[i][j];
                mat[i][j] = mat[i][n - j - 1];
                mat[i][n - j - 1] = temp;
            }
        }
    }

    public boolean findRotation(int[][] mat, int[][] target) {
        for(int k = 0; k < 4; k++) {
            if(equal(mat, target)) return true;
            rotate(mat);
        }
        return false;
    }

    private boolean equal(int[][] a, int[][] b) {
        int n = a.length;
        for(int i = 0; i < n; i++) {
            for(int j = 0; j < n; j++) {
                if(a[i][j] != b[i][j]) return false;
            }
        }
        return true;
    }
}
```

### JavaScript

```javascript
var findRotation = function(mat, target) {

    const rotate = (mat) => {
        let n = mat.length;

        for (let i = 0; i < n; i++) {
            for (let j = i; j < n; j++) {
                [mat[i][j], mat[j][i]] = [mat[j][i], mat[i][j]];
            }
        }

        for (let i = 0; i < n; i++) {
            mat[i].reverse();
        }
    };

    const isEqual = (a, b) => {
        let n = a.length;
        for (let i = 0; i < n; i++) {
            for (let j = 0; j < n; j++) {
                if (a[i][j] !== b[i][j]) return false;
            }
        }
        return true;
    };

    for (let k = 0; k < 4; k++) {
        if (isEqual(mat, target)) return true;
        rotate(mat);
    }

    return false;
};
```

### Python3

```python
class Solution:
    def findRotation(self, mat, target):
        def rotate(mat):
            n = len(mat)
            for i in range(n):
                for j in range(i, n):
                    mat[i][j], mat[j][i] = mat[j][i], mat[i][j]
            for row in mat:
                row.reverse()

        for _ in range(4):
            if mat == target:
                return True
            rotate(mat)

        return False
```

### Go

```go
func findRotation(mat [][]int, target [][]int) bool {

    rotate := func(mat [][]int) {
        n := len(mat)

        for i := 0; i < n; i++ {
            for j := i; j < n; j++ {
                mat[i][j], mat[j][i] = mat[j][i], mat[i][j]
            }
        }

        for i := 0; i < n; i++ {
            for j := 0; j < n/2; j++ {
                mat[i][j], mat[i][n-j-1] = mat[i][n-j-1], mat[i][j]
            }
        }
    }

    isEqual := func(a, b [][]int) bool {
        n := len(a)
        for i := 0; i < n; i++ {
            for j := 0; j < n; j++ {
                if a[i][j] != b[i][j] {
                    return false
                }
            }
        }
        return true
    }

    for k := 0; k < 4; k++ {
        if isEqual(mat, target) {
            return true
        }
        rotate(mat)
    }

    return false
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

1. Start a loop that runs 4 times (each representing a rotation)
2. Compare current matrix with target
3. If all elements match, return true
4. Otherwise, rotate the matrix:

   * Swap elements across diagonal (transpose)
   * Reverse each row
5. Repeat
6. If no match found, return false

---

## Examples

Input:
mat = [[0,1],[1,0]]
target = [[1,0],[0,1]]

Output:
true

---

## How to use / Run locally

1. Copy the code into your preferred language environment
2. Provide input matrices
3. Call the function `findRotation`
4. Print the result

---

## Notes & Optimizations

* Only 4 rotations are needed
* In-place rotation avoids extra memory
* Efficient for small constraints

---

## Author

* [Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
