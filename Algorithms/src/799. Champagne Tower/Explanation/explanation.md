# 799. Champagne Tower

---

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
* [Step-by-step Detailed Explanation](#step-by-step-detailed-explanation)
* [Examples](#examples)
* [How to use / Run locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

---

## Problem Summary

We stack glasses in a pyramid form. The first row has 1 glass, the second row has 2 glasses, and so on.

Each glass can hold only 1 cup of champagne.

If a glass gets more than 1 cup, the extra champagne overflows equally to the two glasses below it.

We are given:

* `poured` → total champagne poured into the top glass
* `query_row` → target row
* `query_glass` → target glass in that row

We must return how full that glass is.

---

## Constraints

* 0 <= poured <= 10^9
* 0 <= query_glass <= query_row < 100

Since the maximum row is less than 100, we can simulate safely.

---

## Intuition

When I saw this problem, I immediately thought about simulation.

Each glass holds maximum 1 cup.
If it gets more than 1, the extra part flows equally to the next row.

So I realized I only need to simulate how overflow moves row by row.

This is a classic dynamic programming flow simulation problem.

---

## Approach

1. I create a 2D DP array where `dp[r][c]` represents the amount of champagne in row `r` and column `c`.
2. I pour all champagne into `dp[0][0]`.
3. For each glass:

   * If it has more than 1 cup
   * Calculate overflow = (value - 1) / 2
   * Add overflow to left and right glass in next row
4. Continue until `query_row`.
5. Return minimum of 1 and the stored value.

This directly simulates the physical behavior.

---

## Data Structures Used

* 2D array (double values)

We use double because champagne distribution may create fractional values.

---

## Operations & Behavior Summary

* Each glass keeps at most 1 cup
* Overflow is split equally
* Flow continues downward
* Only rows up to `query_row` need processing

---

## Complexity

**Time Complexity:** O(n^2)

Where n = query_row (maximum 100).

We process each glass once.

**Space Complexity:** O(n^2)

We store a 2D table up to 100 x 100.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    double champagneTower(int poured, int query_row, int query_glass) {
        vector<vector<double>> dp(101, vector<double>(101, 0.0));
        dp[0][0] = poured;
        
        for (int r = 0; r <= query_row; r++) {
            for (int c = 0; c <= r; c++) {
                if (dp[r][c] > 1.0) {
                    double overflow = (dp[r][c] - 1.0) / 2.0;
                    dp[r + 1][c] += overflow;
                    dp[r + 1][c + 1] += overflow;
                    dp[r][c] = 1.0;
                }
            }
        }
        return min(1.0, dp[query_row][query_glass]);
    }
};
```

### Java

```java
class Solution {
    public double champagneTower(int poured, int query_row, int query_glass) {
        double[][] dp = new double[101][101];
        dp[0][0] = poured;
        
        for (int r = 0; r <= query_row; r++) {
            for (int c = 0; c <= r; c++) {
                if (dp[r][c] > 1.0) {
                    double overflow = (dp[r][c] - 1.0) / 2.0;
                    dp[r + 1][c] += overflow;
                    dp[r + 1][c + 1] += overflow;
                    dp[r][c] = 1.0;
                }
            }
        }
        return Math.min(1.0, dp[query_row][query_glass]);
    }
}
```

### JavaScript

```javascript
var champagneTower = function(poured, query_row, query_glass) {
    const dp = Array.from({ length: 101 }, () => Array(101).fill(0));
    dp[0][0] = poured;
    
    for (let r = 0; r <= query_row; r++) {
        for (let c = 0; c <= r; c++) {
            if (dp[r][c] > 1.0) {
                let overflow = (dp[r][c] - 1.0) / 2.0;
                dp[r + 1][c] += overflow;
                dp[r + 1][c + 1] += overflow;
                dp[r][c] = 1.0;
            }
        }
    }
    return Math.min(1.0, dp[query_row][query_glass]);
};
```

### Python3

```python
class Solution:
    def champagneTower(self, poured: int, query_row: int, query_glass: int) -> float:
        dp = [[0.0] * 101 for _ in range(101)]
        dp[0][0] = poured
        
        for r in range(query_row + 1):
            for c in range(r + 1):
                if dp[r][c] > 1.0:
                    overflow = (dp[r][c] - 1.0) / 2.0
                    dp[r + 1][c] += overflow
                    dp[r + 1][c + 1] += overflow
                    dp[r][c] = 1.0
        
        return min(1.0, dp[query_row][query_glass])
```

### Go

```go
func champagneTower(poured int, query_row int, query_glass int) float64 {
    dp := make([][]float64, 101)
    for i := range dp {
        dp[i] = make([]float64, 101)
    }
    
    dp[0][0] = float64(poured)
    
    for r := 0; r <= query_row; r++ {
        for c := 0; c <= r; c++ {
            if dp[r][c] > 1.0 {
                overflow := (dp[r][c] - 1.0) / 2.0
                dp[r+1][c] += overflow
                dp[r+1][c+1] += overflow
                dp[r][c] = 1.0
            }
        }
    }
    
    if dp[query_row][query_glass] > 1.0 {
        return 1.0
    }
    return dp[query_row][query_glass]
}
```

---

## Step-by-step Detailed Explanation

1. Initialize a 2D array with all zeros.
2. Pour all champagne into top glass.
3. Loop row by row.
4. If a glass has more than 1 cup:

   * Calculate overflow.
   * Send half to left child.
   * Send half to right child.
   * Cap current glass to 1.
5. Continue until required row.
6. Return min(1, value).

This perfectly simulates real champagne flow.

---

## Examples

Input: poured = 2, query_row = 1, query_glass = 1
Output: 0.5

Input: poured = 100000009, query_row = 33, query_glass = 17
Output: 1.0

---

## How to use / Run locally

### C++

Compile using:

```bash
g++ filename.cpp -std=c++17
```

### Java

Compile and run:

```bash
javac Solution.java
java Solution
```

### Python

```bash
python3 solution.py
```

### Go

```bash
go run solution.go
```

---

## Notes & Optimizations

* Since maximum row < 100, simulation is safe.
* We can optimize space using 1D DP array.
* No need to simulate full 100 rows.

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
