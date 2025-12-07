# 1523. Count Odd Numbers in an Interval Range

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

Given two non-negative integers `low` and `high`, return the **count of odd numbers** between `low` and `high` (inclusive).

* Example:

  * `low = 3`, `high = 7` → odd numbers are `[3, 5, 7]`, so the answer is `3`.

We need to compute this **without** iterating over the whole range (because the range can be huge).

---

## Constraints

* `0 <= low <= high <= 10^9`

These constraints tell us:

* The range length can be very large, up to `10^9` numbers.
* A simple loop from `low` to `high` will be too slow.
* We must use **O(1)** math-based logic.

---

## Intuition

I thought like this:

1. Between any two consecutive integers, one is **odd** and the other is **even**.
2. So, in a long sequence of integers, almost **half** of them are odd.
3. Instead of manually checking each number, I tried to find a **formula** to directly count the odd numbers.

Then I realized:

* Count of odd numbers from `1` to `x` is:

[
\text{odds up to } x = \left\lfloor \frac{x + 1}{2} \right\rfloor
]

So if I can compute:

* odds from `1` to `high`, and
* odds from `1` to `low - 1`

Then:

[
\text{odds in } [low, high] = \text{odds up to high} - \text{odds up to (low - 1)}
]

This gives a clean `O(1)` solution.

---

## Approach

1. Define a helper idea:

   * `oddsUpTo(x) = (x + 1) // 2`
   * This returns how many odd numbers exist from `1` to `x`.

2. Use inclusion-exclusion style logic:

   * Odds in `[1, high]` → `oddsUpTo(high)`
   * Odds in `[1, low-1]` → `oddsUpTo(low - 1)`
   * Odds in `[low, high]` = `oddsUpTo(high) - oddsUpTo(low - 1)`

3. Implement this in each language using integer division:

   * C++ / Java / Go: `/` automatically floors for integers.
   * Python: use `//`.
   * JavaScript: use `Math.floor(...)` because `/` is floating point.

4. Return the final difference as the answer.

No loops, no arrays, just math.

---

## Data Structures Used

* Only **primitive integers**:

  * `low`, `high`, maybe a local `x`.
* No arrays, no lists, no maps, nothing extra.

So the space usage is constant.

---

## Operations & Behavior Summary

* Input: two integers `low` and `high`.
* Operation:

  1. Compute `oddsUpTo(high)`.
  2. Compute `oddsUpTo(low - 1)`.
  3. Subtract the second from the first.
* Output: a **single integer** representing how many odd numbers exist in the interval `[low, high]`.

Behavior:

* Works correctly for:

  * `low == high`
  * `low == 0`
  * Very large ranges up to `10^9`.

---

## Complexity

* **Time Complexity:** `O(1)`
  We do a constant number of arithmetic operations, no loops, no recursion.

* **Space Complexity:** `O(1)`
  Only a few integer variables are used. No extra data structure grows with input size.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int countOdds(int low, int high) {
        // Helper lambda: count of odd numbers from 1 to x
        auto oddsUpTo = [](int x) -> int {
            // Integer division automatically floors the result
            return (x + 1) / 2;
        };

        // Odds in [low, high] = oddsUpTo(high) - oddsUpTo(low - 1)
        return oddsUpTo(high) - oddsUpTo(low - 1);
    }
};
```

---

### Java

```java
class Solution {
    public int countOdds(int low, int high) {
        // Odds in [low, high] = oddsUpTo(high) - oddsUpTo(low - 1)
        return oddsUpTo(high) - oddsUpTo(low - 1);
    }

    // Helper: count odd numbers from 1 to x
    private int oddsUpTo(int x) {
        return (x + 1) / 2; // integer division floors automatically
    }
}
```

---

### JavaScript

```javascript
/**
 * @param {number} low
 * @param {number} high
 * @return {number}
 */
var countOdds = function(low, high) {
    // Helper: count odd numbers from 1 to x
    const oddsUpTo = (x) => {
        // Use Math.floor because JS division is floating point
        return Math.floor((x + 1) / 2);
    };

    // Odds in [low, high]
    return oddsUpTo(high) - oddsUpTo(low - 1);
};
```

---

### Python3

```python
class Solution:
    def countOdds(self, low: int, high: int) -> int:
        # Helper: count odd numbers from 1 to x
        def odds_up_to(x: int) -> int:
            # // is integer (floor) division
            return (x + 1) // 2

        # Odds in [low, high]
        return odds_up_to(high) - odds_up_to(low - 1)
```

---

### Go

```go
package main

func countOdds(low int, high int) int {
    // Helper: count odd numbers from 1 to x
    oddsUpTo := func(x int) int {
        return (x + 1) / 2
    }

    // Odds in [low, high]
    return oddsUpTo(high) - oddsUpTo(low-1)
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic is the same in all languages; only syntax changes.
Let’s break it down step by step.

### 1. Define the helper: oddsUpTo(x)

**Core idea:**

```text
oddsUpTo(x) = floor( (x + 1) / 2 )
```

Why?

* If `x` is even:

  * Example: `x = 8`
  * Numbers: 1,2,3,4,5,6,7,8
  * Odds: 1,3,5,7 → 4 odds
  * `(8 + 1) / 2 = 9 / 2 = 4` (floor) ✅

* If `x` is odd:

  * Example: `x = 7`
  * Numbers: 1,2,3,4,5,6,7
  * Odds: 1,3,5,7 → 4 odds
  * `(7 + 1) / 2 = 8 / 2 = 4` ✅

So `(x + 1) // 2` always gives correct count of odds from 1 to x.

### 2. Use it to get odds in [low, high]

We use the concept of prefix counts.

* Odds from `1` to `high` → `oddsUpTo(high)`
* Odds from `1` to `low - 1` → `oddsUpTo(low - 1)`

Now, odds in `[low, high]` are exactly:

```text
oddsUpTo(high) - oddsUpTo(low - 1)
```

This works because:

* When we subtract odds up to `low - 1`, we remove all odds **before** `low`.
* We are left only with odds **from low to high**.

### 3. Language-specific tiny details

* **C++ / Java / Go**:

  ```cpp
  (x + 1) / 2
  ```

  Using integer variables means division automatically floors.

* **Python**:

  ```python
  (x + 1) // 2
  ```

  Here `//` is the integer division operator.

* **JavaScript**:

  ```javascript
  Math.floor((x + 1) / 2)
  ```

  Because JS `/` returns a floating point number.

### 4. Final function

In each language, we:

1. Implement `oddsUpTo(x)` as a small helper (or inline math).
2. Compute `oddsUpTo(high) - oddsUpTo(low - 1)`.
3. Return the result.

That’s the entire solution.

---

## Examples

### Example 1

**Input:**

```text
low = 3, high = 7
```

**Step-by-step:**

* Odds up to 7:

  * `(7 + 1) // 2 = 8 // 2 = 4`
* Odds up to 2 (low - 1 = 2):

  * `(2 + 1) // 2 = 3 // 2 = 1`
* Difference:

  * `4 - 1 = 3`

**Output:**

```text
3
```

**Explanation:** odd numbers are `[3, 5, 7]`.

---

### Example 2

**Input:**

```text
low = 8, high = 10
```

**Step-by-step:**

* Odds up to 10:

  * `(10 + 1) // 2 = 11 // 2 = 5`
* Odds up to 7 (low - 1 = 7):

  * `(7 + 1) // 2 = 8 // 2 = 4`
* Difference:

  * `5 - 4 = 1`

**Output:**

```text
1
```

**Explanation:** odd numbers are `[9]`.

---

## How to use / Run locally

You can copy any language solution into your local environment or into an online judge.

### C++

```bash
g++ -std=c++17 main.cpp -o main
./main
```

Make sure `main.cpp` includes the class `Solution` and some `main()` for testing.

---

### Java

```bash
javac Solution.java
java Solution
```

Put the class `Solution` in `Solution.java`.
In LeetCode, you don’t need `main`, just submit the class.

---

### JavaScript (Node.js)

```bash
node main.js
```

Create `main.js`, define the function `countOdds`, and call it with sample values to test.

---

### Python3

```bash
python3 main.py
```

Create `main.py` with the `Solution` class and write some test code like:

```python
print(Solution().countOdds(3, 7))
```

---

### Go

```bash
go run main.go
```

Create `main.go`, define `countOdds`, and call it inside `main()`.

---

## Notes & Optimizations

* This solution is already optimal:

  * **O(1)** time — just constant arithmetic.
  * **O(1)** space — no extra memory.
* No loops, no recursion, no data structures.
* The logic is purely mathematical and works safely under the given constraints.
* Also avoids overflow:

  * `x` is at most `10^9`, so `x + 1` is fine in all supported languages.

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
