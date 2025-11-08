# 1611. Minimum One Bit Operations to Make Integers Zero

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

Given an integer `n`, we can repeatedly apply either of the following bit operations:

1. **Operation A:** Flip the rightmost (least significant) bit.
2. **Operation B:** Choose the rightmost `1` at position `i` and transform bits `i..0` into: set bit `i+1` to `1`, and set bits `i..0` to `0`. (Equivalently, change binary `...x1000..0` to `...x+1 000..0`.)

Return the **minimum number** of operations required to transform `n` into `0`.

This is LeetCode **1611. Minimum One Bit Operations to Make Integers Zero**.

---

## Constraints

* `0 ≤ n ≤ 10^9` (fits in 32-bit signed integer)
* Answer fits in 32-bit signed integer

---

## Intuition

I looked at the two operations and noticed they behave like steps in **Gray code**—where only one bit changes between consecutive numbers.
That led me to a key fact:

> The minimum number of operations to reach `0` equals the **inverse Gray code** of `n`.

For a Gray code `g`, its binary value `b` can be recovered by:

```
b = g ^ (g >> 1) ^ (g >> 2) ^ ...
```

So I just need to XOR-accumulate `n` with its right shifts until it becomes `0`.

---

## Approach

1. Treat `n` as a Gray code value.
2. Compute its inverse Gray code:

   * Initialize `ans = 0`.
   * While `n > 0`:

     * `ans ^= n`
     * `n >>= 1`
3. Return `ans`.
   This `ans` is the minimum number of operations.

---

## Data Structures Used

* Only a few integers. **No extra data structures** are required.

---

## Operations & Behavior Summary

* **Flip LSB:** toggles only bit 0.
* **Carry-like shift (Operation B):** when encountering the lowest `1` at position `i`, it clears all bits `i..0` and sets `i+1`.
* The optimal strategy effectively removes set bits from **left to right**, matching how inverse Gray code accumulates higher-bit influence into lower bits.

---

## Complexity

* **Time Complexity:** `O(log n)` — one loop per bit of `n`.
* **Space Complexity:** `O(1)` — only constant extra variables.

---

## Multi-language Solutions

### C++

```c++
class Solution {
public:
    int minimumOneBitOperations(int n) {
        // Inverse Gray code:
        // result = n ^ (n >> 1) ^ (n >> 2) ^ ...
        int ans = 0;
        while (n) {
            ans ^= n;   // accumulate toggles
            n >>= 1;    // move influence down
        }
        return ans;
    }
};
```

### Java

```java
class Solution {
    public int minimumOneBitOperations(int n) {
        int ans = 0;
        while (n != 0) {
            ans ^= n;
            n >>= 1;
        }
        return ans;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number} n
 * @return {number}
 */
var minimumOneBitOperations = function(n) {
    // Bitwise ops in JS operate on 32-bit signed ints; constraints are safe.
    let ans = 0;
    while (n !== 0) {
        ans ^= n;
        n >>= 1;
    }
    return ans;
};
```

### Python3

```python
class Solution:
    def minimumOneBitOperations(self, n: int) -> int:
        ans = 0
        while n:
            ans ^= n
            n >>= 1
        return ans
```

### Go

```go
package main

func minimumOneBitOperations(n int) int {
    ans := 0
    for n > 0 {
        ans ^= n
        n >>= 1
    }
    return ans
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

**1) `ans = 0`**
I’ll build the answer using XOR accumulation. This will become the inverse Gray code.

**2) Loop while `n > 0`**
We process the number bit by bit from higher significance to lower by right-shifting.

**3) `ans ^= n`**
XOR records whether each bit needs an odd (1) or even (0) number of toggles when turning `n` to zero using the allowed operations.
This mirrors Gray→binary: each shift spreads the effect of higher bits downwards, and XOR totals those effects.

**4) `n >>= 1`**
Shift to propagate the next higher influence to the right.

**5) Return `ans`**
The final `ans` equals the minimal number of operations.

---

## Examples

| n (dec) | n (bin) | Minimum Ops (answer) | Reason (inverse Gray) |
| ------- | ------- | -------------------- | --------------------- |
| 0       | 0       | 0                    | already zero          |
| 1       | 1       | 1                    | flip LSB once         |
| 2       | 10      | 3                    | `2 ^ 1 = 3`           |
| 3       | 11      | 2                    | `3 ^ 1 = 2`           |
| 6       | 110     | 4                    | `6 ^ 3 ^ 1 = 4`       |
| 9       | 1001    | 14                   | `9 ^ 4 ^ 2 ^ 1 = 14`  |

You can verify these with the provided functions.

---

## How to use / Run locally

**C++ (g++)**

```bash
g++ -std=c++17 main.cpp -O2 && ./a.out
```

**Java**

```bash
javac Solution.java
java Solution
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

Each `main.*` can read a number `n`, call `minimumOneBitOperations(n)`, and print the result.

---

## Notes & Optimizations

* This solution is essentially **constant space** and **bit-length time**.
* The key trick is recognizing the equivalence between the problem’s optimal sequence and **inverse Gray code**.
* No recursion, DP, or precomputation is needed.
* Works safely within 32-bit integer limits of all listed languages.

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
