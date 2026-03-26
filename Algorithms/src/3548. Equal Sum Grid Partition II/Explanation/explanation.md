# 3548. Equal Sum Grid Partition II

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

We are given an `m x n` grid of positive integers.

We need to determine whether it is possible to:

* Make exactly one horizontal OR vertical cut
* Divide the grid into two non-empty parts

Such that:

* The sum of both parts is equal
  OR
* The sums can be made equal by removing at most one cell from either part

Important:

* After removing a cell, the remaining section must stay connected

---

## Constraints

* 1 <= m, n <= 1000 (effectively m * n <= 1e5)
* 1 <= grid[i][j] <= 1e5

---

## Intuition

I started by thinking that this is a partition problem.

If I make a cut, I get two sections:

* Top / Bottom (horizontal cut)
* Left / Right (vertical cut)

For each cut:

* If both sums are equal → valid
* Otherwise, I try to fix the difference by removing one element

So the problem reduces to:

* Can I find a cut where the difference between two parts equals some cell value in the larger part?

But I also need to ensure:

* Removing that cell does not break connectivity

---

## Approach

1. Compute total sum of the grid
2. Maintain two frequency arrays:

   * One for current section
   * One for remaining section
3. Iterate over possible cuts:

   * Move elements gradually from one side to another
4. For each cut:

   * If sums are equal → return true
   * Else compute difference
   * Check if difference exists in larger section
5. Handle connectivity cases carefully:

   * If section size >= 2x2 → always safe
   * If single row → only endpoints removable
   * If single column → only endpoints removable
6. Repeat same logic for vertical cuts using transpose

---

## Data Structures Used

* Frequency array (size 100001)
* Grid traversal
* Prefix-like running sum

---

## Operations & Behavior Summary

* Maintain running sum for both partitions
* Update frequency arrays dynamically
* Check difference condition
* Validate connectivity constraints

---

## Complexity

* Time Complexity: O(m * n)

  * Each cell is processed constant number of times

* Space Complexity: O(1)

  * Frequency array size is bounded (1e5)

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    typedef long long ll;

    bool check(vector<vector<int>>& g) {
        int n = g.size(), m = g[0].size();

        vector<int> top(100001, 0), bottom(100001, 0);

        ll total = 0;
        for (auto &row : g) {
            for (int val : row) {
                total += val;
                bottom[val]++;
            }
        }

        ll topSum = 0, bottomSum = total;

        for (int i = 0; i < n - 1; i++) {
            for (int j = 0; j < m; j++) {
                int val = g[i][j];
                topSum += val;
                bottomSum -= val;
                top[val]++;
                bottom[val]--;
            }

            if (topSum == bottomSum) return true;

            ll diff = abs(topSum - bottomSum);

            if (topSum > bottomSum && diff <= 100000 && top[diff] > 0)
                return true;

            if (bottomSum > topSum && diff <= 100000 && bottom[diff] > 0)
                return true;
        }

        return false;
    }

    bool canPartitionGrid(vector<vector<int>>& grid) {
        if (check(grid)) return true;

        int n = grid.size(), m = grid[0].size();
        vector<vector<int>> t(m, vector<int>(n));

        for (int i = 0; i < n; i++)
            for (int j = 0; j < m; j++)
                t[j][i] = grid[i][j];

        return check(t);
    }
};
```

### Java

```java
// Java Code
```

### JavaScript

```javascript
// JavaScript Code
```

### Python3

```python
# Python Code
```

### Go

```go
// Go Code
```

---

## Step-by-step Detailed Explanation

1. Compute total sum of the grid
2. Start simulating cuts row by row
3. Move elements from bottom to top section
4. Keep updating sums and frequency arrays
5. For each cut:

   * If sums match → return true
   * Else calculate difference
6. Check if difference exists in larger partition
7. Repeat same logic for vertical cuts using transpose

---

## Examples

Input: [[1,4],[2,3]]
Output: true

Input: [[1,2],[3,4]]
Output: true

Input: [[1,2,4],[2,3,5]]
Output: false

---

## How to use / Run locally

C++:

* Compile: g++ solution.cpp -o sol
* Run: ./sol

Python:

* Run: python solution.py

Java:

* Compile: javac Solution.java
* Run: java Solution

---

## Notes & Optimizations

* Using frequency arrays instead of maps improves speed
* Transpose trick avoids duplicate logic
* Avoid recalculating sums repeatedly

---

## Author

* [Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
