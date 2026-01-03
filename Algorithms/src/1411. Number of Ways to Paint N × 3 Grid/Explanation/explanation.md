# Problem Title

**1411. Number of Ways to Paint N × 3 Grid**

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
* How to Use / Run Locally
* Notes & Optimizations
* Author

---

## Problem Summary

I am given an `n × 3` grid.
I need to paint each cell using **exactly one of three colors**: Red, Yellow, or Green.

Rules:

* No two **adjacent cells** can have the same color
* Adjacent means **horizontal or vertical**

My task is to return the **total number of valid ways** to paint the grid.
Since the number can be large, I must return the answer **modulo 10⁹ + 7**.

---

## Constraints

* `1 ≤ n ≤ 5000`
* Grid size is always `n × 3`
* Colors available = 3
* Modulo = `1000000007`

---

## Intuition

When I first saw the problem, I realized that trying all color combinations would be impossible because `n` can go up to 5000.

So I started thinking **row by row**.

Then I noticed something important:

Every valid row of 3 cells falls into **only two patterns**.

This observation helped me reduce a very complex problem into a simple dynamic programming solution with constant space.

---

## Approach

I classify each row into **two types**:

### Type 1: ABA pattern

* First and third cells have the same color
* Middle cell has a different color
* Example: Red Yellow Red

Number of such patterns for one row = **6**

---

### Type 2: ABC pattern

* All three cells have different colors
* Example: Red Yellow Green

Number of such patterns for one row = **6**

---

### Dynamic Programming Idea

I keep track of:

* `same` → number of ways where the current row is ABA type
* `diff` → number of ways where the current row is ABC type

Initial values (for first row):

```bash
same = 6
diff = 6
```

For each new row:

```bash
newSame = same * 3 + diff * 2
newDiff = same * 2 + diff * 2
```

I repeat this process for `n` rows.

Final Answer:

```bash
(same + diff) % MOD
```

---

## Data Structures Used

* Simple integer variables
* No arrays or matrices
* Constant space dynamic programming

---

## Operations & Behavior Summary

* Process grid row by row
* Maintain only two values for DP
* Apply modulo at every step
* Avoid unnecessary memory usage
* Optimized for large `n`

---

## Complexity

**Time Complexity:** `O(n)`
I iterate once for each row.

**Space Complexity:** `O(1)`
I only use a few variables, no extra data structures.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int numOfWays(int n) {
        const int MOD = 1e9 + 7;
        long same = 6, diff = 6;

        for (int i = 2; i <= n; i++) {
            long newSame = (same * 3 + diff * 2) % MOD;
            long newDiff = (same * 2 + diff * 2) % MOD;
            same = newSame;
            diff = newDiff;
        }
        return (same + diff) % MOD;
    }
};
```

---

### Java

```java
class Solution {
    public int numOfWays(int n) {
        final int MOD = 1_000_000_007;
        long same = 6, diff = 6;

        for (int i = 2; i <= n; i++) {
            long newSame = (same * 3 + diff * 2) % MOD;
            long newDiff = (same * 2 + diff * 2) % MOD;
            same = newSame;
            diff = newDiff;
        }
        return (int)((same + diff) % MOD);
    }
}
```

---

### JavaScript

```javascript
var numOfWays = function(n) {
    const MOD = 1e9 + 7;
    let same = 6, diff = 6;

    for (let i = 2; i <= n; i++) {
        const newSame = (same * 3 + diff * 2) % MOD;
        const newDiff = (same * 2 + diff * 2) % MOD;
        same = newSame;
        diff = newDiff;
    }
    return (same + diff) % MOD;
};
```

---

### Python3

```python
class Solution:
    def numOfWays(self, n: int) -> int:
        MOD = 10**9 + 7
        same, diff = 6, 6

        for _ in range(2, n + 1):
            new_same = (same * 3 + diff * 2) % MOD
            new_diff = (same * 2 + diff * 2) % MOD
            same, diff = new_same, new_diff

        return (same + diff) % MOD
```

---

### Go

```go
func numOfWays(n int) int {
    const MOD = 1e9 + 7
    same, diff := 6, 6

    for i := 2; i <= n; i++ {
        newSame := (same*3 + diff*2) % MOD
        newDiff := (same*2 + diff*2) % MOD
        same, diff = newSame, newDiff
    }
    return (same + diff) % MOD
}
```

---

## Step-by-step Detailed Explanation

1. I observed that each row only depends on the previous row.
2. I reduced all row patterns into two types.
3. I calculated transitions between these two types.
4. I used dynamic programming with constant space.
5. I applied modulo to avoid overflow.
6. I returned the final count after `n` rows.

---

## Examples

### Example 1

Input:

```bash
n = 1
```

Output:

```bash
12
```

### Example 2

Input:

```bash
n = 5000
```

Output:

```bash
30228214
```

---

## How to Use / Run Locally

1. Clone the repository
2. Choose your language file
3. Run using the respective compiler or interpreter

Example (C++):

```bash
g++ solution.cpp
./a.out
```

Example (Python):

```bash
python solution.py
```

---

## Notes & Optimizations

* Brute force is impossible due to large constraints
* Matrix DP is unnecessary
* Constant space DP is optimal
* This approach is interview-friendly and scalable

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
