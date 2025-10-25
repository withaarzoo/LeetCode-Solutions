# 1716. Calculate Money in Leetcode Bank — README

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
* [Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)](#step-by-step-detailed-explanation-c-java-javascript-python3-go)
* [Examples](#examples)
* [How to use / Run locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

---

## Problem Summary

Hercy saves money every day following a pattern:

* On the first Monday he puts $1.
* Each next day of that week he puts $1 more than the previous day (so Monday=1, Tuesday=2, ..., Sunday=7 in the first week).
* On the next Monday, he starts with $2 (one more than previous Monday), and then proceeds +1 each day again for that week.
  Given `n` (number of days), return the total amount of money saved after `n` days.

---

## Constraints

* `1 <= n <= 1000`

---

## Intuition

I thought about the pattern week-by-week. Each week has 7 days with consecutive numbers (base, base+1, ..., base+6). Every new week the base increases by 1. So I can separate `n` into:

* `w` = number of full weeks (`n // 7`)
* `r` = remaining days (`n % 7`)
  Then compute sum of full weeks with a closed formula and add the sum for the remaining days. This avoids looping day-by-day and yields O(1) time.

---

## Approach

1. Compute `w = n // 7` and `r = n % 7`.
2. Sum of one week starting with base `b` is `b + (b+1) + ... + (b+6) = 7*b + (0+1+...+6) = 7*b + 21`. For base = 1, the sum is `28` (since `7*1 + 21 = 28`).
3. For full weeks 0..w-1, base of week `i` is `1 + i`. Sum over these weeks:

   * baseline `28 * w` (if every week started at 1),
   * plus incremental increase `7 * sum(0..w-1) = 7 * w*(w-1)/2`.
     So full weeks sum = `28*w + 7*(w*(w-1)/2)`.
4. For the leftover `r` days (starting base `1 + w`), the sum = arithmetic series:

   * `r*(1 + w) + r*(r-1)/2`.
5. Final answer = `full_weeks_sum + leftover_sum`.

---

## Data Structures Used

* Only primitive integer variables.
* No arrays, lists, or extra containers are needed.

---

## Operations & Behavior Summary

* Integer division and modulo to split days into weeks and leftover days.
* Use of arithmetic sum formulas (sum of consecutive integers) to compute totals in O(1).

---

## Complexity

* **Time Complexity:** `O(1)` — only a constant number of arithmetic operations regardless of `n`.

  * `n` is the number of days input.
* **Space Complexity:** `O(1)` — constant extra space used for variables `w`, `r`, and intermediate results.

---

## Multi-language Solutions

### C++

```c++
class Solution {
public:
    int totalMoney(int n) {
        int w = n / 7;         // full weeks
        int r = n % 7;         // remaining days
        // full weeks sum = w*28 + 7 * (0 + 1 + ... + (w-1))
        int fullWeeksSum = w * 28 + 7 * (w * (w - 1) / 2);
        // remaining days sum = r*(1 + w) + r*(r-1)/2
        int remSum = r * (1 + w) + (r * (r - 1) / 2);
        return fullWeeksSum + remSum;
    }
};
```

### Java

```java
class Solution {
    public int totalMoney(int n) {
        int w = n / 7;  // full weeks
        int r = n % 7;  // remaining days
        int fullWeeksSum = w * 28 + 7 * (w * (w - 1) / 2);
        int remSum = r * (1 + w) + (r * (r - 1) / 2);
        return fullWeeksSum + remSum;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number} n
 * @return {number}
 */
var totalMoney = function(n) {
    const w = Math.floor(n / 7);
    const r = n % 7;
    const fullWeeksSum = w * 28 + 7 * (w * (w - 1) / 2);
    const remSum = r * (1 + w) + (r * (r - 1) / 2);
    return fullWeeksSum + remSum;
};
```

### Python3

```python
class Solution:
    def totalMoney(self, n: int) -> int:
        w = n // 7
        r = n % 7
        full_weeks_sum = w * 28 + 7 * (w * (w - 1) // 2)
        rem_sum = r * (1 + w) + (r * (r - 1) // 2)
        return full_weeks_sum + rem_sum
```

### Go

```go
func totalMoney(n int) int {
    w := n / 7
    r := n % 7
    fullWeeksSum := w*28 + 7*(w*(w-1)/2)
    remSum := r*(1+w) + (r*(r-1)/2)
    return fullWeeksSum + remSum
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

I'll explain the key parts (same logic in all languages). I explain as if I'm teaching a friend.

1. **Split `n` into weeks and days**

   * `w = n // 7` (or `n / 7`) — this gives how many full 7-day weeks I have.
   * `r = n % 7` — remaining days after counting full weeks.

2. **Compute sum for full weeks**

   * Each week `i` (0-indexed) starts with `1 + i` on Monday and increases by 1 each day.
   * Sum of one week starting at base `b` is `b + (b+1) + ... + (b+6) = 7*b + 21`.
   * If every week started at 1, `w` weeks would sum to `28 * w` (since a week starting at 1 sums to 28).
   * But weeks start at `1, 2, 3, ..., 1+(w-1)`. The extra contribution from these increasing bases equals `7 * (0 + 1 + ... + (w-1))`.
   * Use formula for sum of first (w-1) natural numbers: `0 + 1 + ... + (w-1) = w*(w-1)/2`.
   * So: `fullWeeksSum = 28*w + 7*(w*(w-1)/2)`.

3. **Compute sum for remaining `r` days**

   * After `w` full weeks Monday base becomes `1 + w`.
   * Remaining `r` days form an arithmetic sequence starting at `1 + w` with difference `1`.
   * Sum of `r` terms where first term is `a = 1 + w`: `r*a + sum(0..r-1)` = `r*(1+w) + r*(r-1)/2`.

4. **Return total**

   * `total = fullWeeksSum + remSum`.

5. **Arithmetic edge cases**

   * Works for `n < 7` because `w = 0` and `r = n`, so the formula reduces to arithmetic sum from `1` to `n`.
   * Works for `n` multiple of 7 as `r = 0` so leftover sum is 0.

---

## Examples

* **Example 1**

  * Input: `n = 4`
  * Calculation: `1 + 2 + 3 + 4 = 10`
  * Output: `10`

* **Example 2**

  * Input: `n = 10`
  * Explanation: Week1 = `1+2+3+4+5+6+7 = 28`, leftover 3 days of week2 = `2+3+4=9`, total = `37`.
  * Output: `37`

* **Example 3**

  * Input: `n = 20`
  * Output: `96` (matches arithmetic breakdown in problem).

---

## How to use / Run locally

### C++

1. Put the `Solution` class into a `.cpp` file and call the method from `main`.
2. Compile & run:

```bash
g++ -std=c++17 solution.cpp -O2 -o solution
./solution
```

### Java

1. Place `Solution` in `Solution.java` and add a `main` to test.
2. Compile & run:

```bash
javac Solution.java
java Solution
```

### JavaScript (Node.js)

1. Save to `solution.js`, export or call `totalMoney` from a small driver.
2. Run:

```bash
node solution.js
```

### Python3

1. Save class in `solution.py` and call `Solution().totalMoney(n)` in a `main` block.
2. Run:

```bash
python3 solution.py
```

### Go

1. Put `totalMoney` into `main.go` with a `main` function that prints results.
2. Build & run:

```bash
go run main.go
```

---

## Notes & Optimizations

* The formula-based approach is optimal for time and space; no loops (other than small constant-time arithmetic) are required.
* For languages with integer division, ensure integer arithmetic is used to avoid float rounding (`//` in Python, integer ops in others).
* All operations fit comfortably within 32-bit integers for `n <= 1000` (maximum sum ≈ a few thousand). If `n` were much larger, consider using 64-bit integers.
* This solution is deterministic and exact (no iterative simulation).

---

## Author

* [Md. Aarzoo Islam](https://bento.me/withaarzoo)
