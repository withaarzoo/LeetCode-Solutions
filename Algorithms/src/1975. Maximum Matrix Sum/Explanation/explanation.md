# Problem Title

**1975. Maximum Matrix Sum (LeetCode)**

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

You are given an `n x n` integer matrix. You can perform the following operation any number of times:

* Choose **any two adjacent cells** (sharing a border) and multiply both values by `-1`.

Your task is to **maximize the total sum of the matrix** after performing optimal operations.

---

## Constraints

* `n == matrix.length == matrix[i].length`
* `2 <= n <= 250`
* `-10^5 <= matrix[i][j] <= 10^5`

---

## Intuition

When I read the problem, I focused on the operation. Every operation flips **two numbers together**.

That gave me an important observation:

* The **count of negative numbers always changes by 2**.
* So the **parity (even or odd)** number of negative elements **never changes**.

From this, I understood:

* If the total number of negative values is **even**, I can make **all values positive**.
* If the total number of negative values is **odd**, then **one value must stay negative**.

To maximize the sum, that remaining negative value should have the **smallest absolute value**.

---

## Approach

I solved the problem using a greedy + math approach:

1. Traverse the entire matrix.
2. Add the **absolute value** of every element to the total sum.
3. Count how many negative numbers exist.
4. Track the **minimum absolute value** in the matrix.
5. If the count of negatives is odd, subtract `2 × minAbs` from the sum.

This gives the maximum possible matrix sum.

---

## Data Structures Used

* No extra data structure is required.
* Only basic variables are used to track:

  * Total sum
  * Count of negative numbers
  * Minimum absolute value

---

## Operations & Behavior Summary

* Flipping two adjacent cells preserves the parity of negative numbers.
* Absolute values give the maximum possible contribution.
* Only one smallest value needs adjustment when negatives are odd.

---

## Complexity

* **Time Complexity:** `O(n²)`

  * Every cell in the matrix is visited once.

* **Space Complexity:** `O(1)`

  * No extra space is used apart from variables.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    long long maxMatrixSum(vector<vector<int>>& matrix) {
        long long sum = 0;
        int negativeCount = 0;
        int minAbs = INT_MAX;

        for (auto &row : matrix) {
            for (int val : row) {
                sum += llabs(val);
                if (val < 0) negativeCount++;
                minAbs = min(minAbs, abs(val));
            }
        }

        if (negativeCount % 2 == 1) {
            sum -= 2LL * minAbs;
        }

        return sum;
    }
};
```

---

### Java

```java
class Solution {
    public long maxMatrixSum(int[][] matrix) {
        long sum = 0;
        int negativeCount = 0;
        int minAbs = Integer.MAX_VALUE;

        for (int[] row : matrix) {
            for (int val : row) {
                sum += Math.abs(val);
                if (val < 0) negativeCount++;
                minAbs = Math.min(minAbs, Math.abs(val));
            }
        }

        if (negativeCount % 2 == 1) {
            sum -= 2L * minAbs;
        }

        return sum;
    }
}
```

---

### JavaScript

```javascript
var maxMatrixSum = function(matrix) {
    let sum = 0;
    let negativeCount = 0;
    let minAbs = Infinity;

    for (let row of matrix) {
        for (let val of row) {
            sum += Math.abs(val);
            if (val < 0) negativeCount++;
            minAbs = Math.min(minAbs, Math.abs(val));
        }
    }

    if (negativeCount % 2 === 1) {
        sum -= 2 * minAbs;
    }

    return sum;
};
```

---

### Python3

```python
class Solution:
    def maxMatrixSum(self, matrix):
        total_sum = 0
        negative_count = 0
        min_abs = float('inf')

        for row in matrix:
            for val in row:
                total_sum += abs(val)
                if val < 0:
                    negative_count += 1
                min_abs = min(min_abs, abs(val))

        if negative_count % 2 == 1:
            total_sum -= 2 * min_abs

        return total_sum
```

---

### Go

```go
func maxMatrixSum(matrix [][]int) int64 {
    var sum int64 = 0
    negativeCount := 0
    minAbs := int(1e9)

    for _, row := range matrix {
        for _, val := range row {
            if val < 0 {
                negativeCount++
                val = -val
            }
            sum += int64(val)
            if val < minAbs {
                minAbs = val
            }
        }
    }

    if negativeCount%2 == 1 {
        sum -= int64(2 * minAbs)
    }

    return sum
}
```

---

## Step-by-step Detailed Explanation

1. Traverse every element of the matrix.
2. Convert each value to absolute and add it to the sum.
3. Count how many elements were negative.
4. Track the smallest absolute value.
5. If negatives are odd, subtract twice the smallest value.
6. Return the final sum.

---

## Examples

**Input:**

```bash
[[1, -1], [-1, 1]]
```

**Output:**

```bash
4
```

---

**Input:**

```bash
[[1,2,3],[-1,-2,-3],[1,2,3]]
```

**Output:**

```bash
16
```

---

## How to use / Run locally

1. Copy the solution code in your preferred language.
2. Paste it into your LeetCode editor or local IDE.
3. Run with sample inputs to verify.
4. Submit on LeetCode.

---

## Notes & Optimizations

* No matrix modification is required.
* Greedy + math approach avoids simulation.
* Works efficiently even for large `n`.

---

## Author

* **Md Aarzoo Islam**
* Portfolio: [https://bento.me/withaarzoo](https://bento.me/withaarzoo)
