# 2169. Count Operations to Obtain Zero

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

You are given two non-negative integers `num1` and `num2`.
In one operation:

* If `num1 >= num2` and `num2 > 0`, replace `num1` with `num1 - num2`.
* Otherwise if `num2 >= num1` and `num1 > 0`, replace `num2` with `num2 - num1`.

Return the **number of operations** needed until either `num1 == 0` or `num2 == 0`.

---

## Constraints

* `0 <= num1, num2 <= 10^5`
* Answer fits in a 32-bit integer.

---

## Intuition

I first thought I should simulate the process: keep subtracting the smaller number from the bigger one until one becomes zero.
But if the numbers are large, subtracting one-by-one is slow. Then I realized this is the same idea as the **Euclidean algorithm** for GCD.
Instead of doing many single subtractions, I can count them in **batches** using division:
If `a >= b`, the number of times I would subtract `b` from `a` is `a / b`. After those, the remainder is `a % b`. That jumps forward fast while counting the exact number of operations.

---

## Approach

1. Let `a = num1`, `b = num2`. If either is `0`, return `0`.
2. While both `a` and `b` are positive:

   * Ensure `a >= b` (swap if not).
   * Add `a / b` to the operation count (this equals repeated subtractions).
   * Set `a = a % b` (what’s left after removing as many `b`s as possible).
3. When one becomes `0`, stop and return the count.

This is exactly the problem’s process but **batched**, so it runs in logarithmic time.

---

## Data Structures Used

* None beyond a few integers. The solution is **in-place** and uses **O(1)** extra space.

---

## Operations & Behavior Summary

* **Swap (optional):** Keep `a >= b` so logic is simple.
* **Batch count:** `ops += a / b` counts how many single-step subtractions happen at once.
* **Reduce:** `a %= b` applies all those subtractions in one step.
* **Stop condition:** When `a == 0` or `b == 0`.

---

## Complexity

* **Time Complexity:** `O(log(max(num1, num2)))` — same as the Euclidean algorithm. Each iteration shrinks one value quickly using modulo.
* **Space Complexity:** `O(1)` — only constant extra variables are used.

---

## Multi-language Solutions

### C++

```c++
class Solution {
public:
    int countOperations(int num1, int num2) {
        long long a = num1, b = num2; // int is fine; long long is extra safe
        int ops = 0;
        while (a > 0 && b > 0) {
            if (a < b) std::swap(a, b); // ensure a >= b
            ops += a / b;               // batch of subtractions
            a %= b;                     // remainder after removing (a/b) copies of b
        }
        return ops;
    }
};
```

### Java

```java
class Solution {
    public int countOperations(int num1, int num2) {
        long a = num1, b = num2; // widen just to be safe
        int ops = 0;
        while (a > 0 && b > 0) {
            if (a < b) { long t = a; a = b; b = t; } // ensure a >= b
            ops += (int)(a / b);                     // number of subtractions at once
            a %= b;                                  // leftover after those subtractions
        }
        return ops;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number} num1
 * @param {number} num2
 * @return {number}
 */
var countOperations = function(num1, num2) {
    let a = num1, b = num2;
    let ops = 0;
    while (a > 0 && b > 0) {
        if (a < b) { const t = a; a = b; b = t; } // ensure a >= b
        ops += Math.floor(a / b);                 // batched subtractions
        a = a % b;                                // remainder
    }
    return ops;
};
```

### Python3

```python
class Solution:
    def countOperations(self, num1: int, num2: int) -> int:
        a, b = num1, num2
        ops = 0
        while a > 0 and b > 0:
            if a < b:
                a, b = b, a          # ensure a >= b
            ops += a // b             # count how many subtractions at once
            a %= b                    # remainder left in a
        return ops
```

### Go

```go
package main

func countOperations(num1 int, num2 int) int {
    a, b := num1, num2
    ops := 0
    for a > 0 && b > 0 {
        if a < b {
            a, b = b, a // ensure a >= b
        }
        ops += a / b    // how many subtractions we do at once
        a = a % b       // what's left in a
    }
    return ops
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

1. **Keep the larger number first**
   If `a < b`, I swap them so `a >= b`. This way, I always “subtract from the bigger one”.

2. **Batch the subtractions**
   `ops += a / b` means: if I were doing it slowly, I’d subtract `b` from `a` exactly `a / b` times.
   Example: `a = 23`, `b = 5` → I’d subtract 5 four times. So I add `4` to `ops`.

3. **Apply them instantly**
   `a %= b` leaves only the remainder after those subtractions.
   From the example, `23 % 5 = 3`. So now `a` becomes `3`.

4. **Repeat**
   Loop while both numbers are positive. Swap again if needed and continue.

5. **Stop**
   When one number becomes `0`, the process ends. Return the total count I accumulated.

This mirrors the problem’s rules exactly, just faster.

---

## Examples

1. **Input:** `num1 = 2, num2 = 3`

   * Step 1: `a=3, b=2` → `ops += 3/2 = 1` → `a = 3 % 2 = 1`
   * Step 2: swap → `a=2, b=1` → `ops += 2/1 = 2` → `a = 0`
   * **Answer:** `ops = 3`

2. **Input:** `num1 = 10, num2 = 10`

   * `ops += 10/10 = 1` → `a = 0`
   * **Answer:** `1`

3. **Input:** `num1 = 0, num2 = 7`

   * Already one is zero → **Answer:** `0`

---

## How to use / Run locally

**C++**

```bash
g++ -std=c++17 main.cpp -O2 && ./a.out
```

**Java**

```bash
javac Solution.java && java Solution
```

**JavaScript (Node.js)**

```bash
node main.js
```

**Python3**

```bash
python3 main.py
```

**Go**

```bash
go run main.go
```

*Tip:* Wrap the function into a small driver that reads two integers and prints the result to test locally.

---

## Notes & Optimizations

* The naive loop that subtracts once per operation can take up to `O(max(num1, num2))` steps.
* Using division + modulo reduces it to `O(log(max(num1, num2)))`.
* This is the same optimization idea behind the Euclidean algorithm for computing GCD.

---

## Author

* [Md. Aarzoo Islam](https://bento.me/withaarzoo)
