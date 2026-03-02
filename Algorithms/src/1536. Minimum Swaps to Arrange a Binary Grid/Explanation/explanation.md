# 1536. Minimum Swaps to Arrange a Binary Grid

## Table of Contents

* Problem Summary
* Constraints
* Intuition
* Approach
* Data Structures Used
* Operations & Behavior Summary
* Complexity
* Multi-language Solutions

  * C++
  * Java
  * JavaScript
  * Python3
  * Go
* Step-by-step Detailed Explanation
* Examples
* How to use / Run locally
* Notes & Optimizations
* Author

---

## Problem Summary

You are given an `n x n` binary grid. In one operation, you can swap two adjacent rows.

A grid is considered valid if all the cells above the main diagonal are zeros.

Return the minimum number of swaps required to make the grid valid. If it is not possible, return -1.

---

## Constraints

* n == grid.length == grid[i].length
* 1 <= n <= 200
* grid[i][j] is either 0 or 1

---

## Intuition

When I carefully looked at the condition, I realized something important.

For row `i`, all cells where column `j > i` must be zero. That means row 0 must have `n - 1` trailing zeros, row 1 must have `n - 2` trailing zeros, and so on.

So instead of checking the entire matrix repeatedly, I simplified the problem.

I only need to count how many consecutive zeros appear at the end of each row.

If I can arrange rows such that for each position `i`, the row placed there has at least `n - 1 - i` trailing zeros, then the grid becomes valid.

This naturally leads to a greedy approach.

---

## Approach

1. Count trailing zeros for every row.
2. For each row index `i`, calculate required zeros = `n - 1 - i`.
3. Search from row `i` downward to find a row that satisfies the requirement.
4. If found at index `j`, move it upward using adjacent swaps.
5. Add `j - i` to swap count.
6. If no such row exists, return -1.
7. Continue until all rows are placed correctly.

This works because choosing the nearest valid row minimizes swaps.

---

## Data Structures Used

* Array or vector to store trailing zero counts
* Integer variable to track swap count

No extra complex data structures are required.

---

## Operations & Behavior Summary

* Traverse grid to count trailing zeros
* Greedy search for suitable row
* Simulate adjacent swaps by shifting elements
* Count total swaps

---

## Complexity

Time Complexity: O(n^2)

* Counting trailing zeros takes O(n^2)
* For each row, we may search downward up to O(n)

Space Complexity: O(n)

* We store trailing zero counts in an array

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int minSwaps(vector<vector<int>>& grid) {
        int n = grid.size();
        vector<int> trailing(n, 0);

        for (int i = 0; i < n; i++) {
            int count = 0;
            for (int j = n - 1; j >= 0; j--) {
                if (grid[i][j] == 0) count++;
                else break;
            }
            trailing[i] = count;
        }

        int swaps = 0;

        for (int i = 0; i < n; i++) {
            int required = n - 1 - i;
            int j = i;

            while (j < n && trailing[j] < required) j++;
            if (j == n) return -1;

            while (j > i) {
                swap(trailing[j], trailing[j - 1]);
                swaps++;
                j--;
            }
        }
        return swaps;
    }
};
```

### Java

```java
class Solution {
    public int minSwaps(int[][] grid) {
        int n = grid.length;
        int[] trailing = new int[n];

        for (int i = 0; i < n; i++) {
            int count = 0;
            for (int j = n - 1; j >= 0; j--) {
                if (grid[i][j] == 0) count++;
                else break;
            }
            trailing[i] = count;
        }

        int swaps = 0;

        for (int i = 0; i < n; i++) {
            int required = n - 1 - i;
            int j = i;

            while (j < n && trailing[j] < required) j++;
            if (j == n) return -1;

            while (j > i) {
                int temp = trailing[j];
                trailing[j] = trailing[j - 1];
                trailing[j - 1] = temp;
                swaps++;
                j--;
            }
        }
        return swaps;
    }
}
```

### JavaScript

```javascript
var minSwaps = function(grid) {
    const n = grid.length;
    const trailing = new Array(n).fill(0);

    for (let i = 0; i < n; i++) {
        let count = 0;
        for (let j = n - 1; j >= 0; j--) {
            if (grid[i][j] === 0) count++;
            else break;
        }
        trailing[i] = count;
    }

    let swaps = 0;

    for (let i = 0; i < n; i++) {
        let required = n - 1 - i;
        let j = i;

        while (j < n && trailing[j] < required) j++;
        if (j === n) return -1;

        while (j > i) {
            [trailing[j], trailing[j - 1]] = [trailing[j - 1], trailing[j]];
            swaps++;
            j--;
        }
    }

    return swaps;
};
```

### Python3

```python
class Solution:
    def minSwaps(self, grid: List[List[int]]) -> int:
        n = len(grid)
        trailing = []

        for i in range(n):
            count = 0
            for j in range(n - 1, -1, -1):
                if grid[i][j] == 0:
                    count += 1
                else:
                    break
            trailing.append(count)

        swaps = 0

        for i in range(n):
            required = n - 1 - i
            j = i

            while j < n and trailing[j] < required:
                j += 1

            if j == n:
                return -1

            while j > i:
                trailing[j], trailing[j - 1] = trailing[j - 1], trailing[j]
                swaps += 1
                j -= 1

        return swaps
```

### Go

```go
func minSwaps(grid [][]int) int {
    n := len(grid)
    trailing := make([]int, n)

    for i := 0; i < n; i++ {
        count := 0
        for j := n - 1; j >= 0; j-- {
            if grid[i][j] == 0 {
                count++
            } else {
                break
            }
        }
        trailing[i] = count
    }

    swaps := 0

    for i := 0; i < n; i++ {
        required := n - 1 - i
        j := i

        for j < n && trailing[j] < required {
            j++
        }

        if j == n {
            return -1
        }

        for j > i {
            trailing[j], trailing[j-1] = trailing[j-1], trailing[j]
            swaps++
            j--
        }
    }

    return swaps
}
```

---

## Step-by-step Detailed Explanation

1. We count trailing zeros for each row.
2. For position i, we compute required zeros.
3. We search for the first row below that satisfies the requirement.
4. If found, we move it upward using adjacent swaps.
5. We increase swap count accordingly.
6. If not found, we return -1.

This ensures minimum swaps because we always choose the closest valid row.

---

## Examples

Example 1:

Input:
[[0,0,1],[1,1,0],[1,0,0]]

Output:
3

Example 2:

Input:
[[0,1,1,0],[0,1,1,0],[0,1,1,0],[0,1,1,0]]

Output:
-1

---

## How to use / Run locally

C++:
Compile using g++ and run the executable.

Java:
Compile using javac and run using java.

Python:
Run using python3 filename.py.

JavaScript:
Run using node filename.js.

Go:
Run using go run filename.go.

---

## Notes & Optimizations

* We only store trailing zero counts instead of modifying grid.
* We simulate swaps logically without modifying full matrix.
* Greedy approach guarantees minimum adjacent swaps.

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
