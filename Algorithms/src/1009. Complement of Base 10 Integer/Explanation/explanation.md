# 1009. Complement of Base 10 Integer

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

The complement of a base-10 integer is defined by flipping every bit in its binary representation.
This means every `0` becomes `1`, and every `1` becomes `0`.

However, we only flip bits **up to the most significant bit of the number**.
We do not consider leading zeros.

Example:

Input: `5`
Binary: `101`
Flip bits → `010`
Output: `2`

The task is to return the **complement of the given integer `n`**.

---

## Constraints

```
0 <= n < 10^9
```

---

## Intuition

When I first read the problem, I immediately thought about flipping bits using the **bitwise NOT (`~`) operator**.

But that does not work directly.

The reason is that computers store integers using a fixed number of bits (like 32 bits).
So applying `~n` flips **all bits**, including the leading zeros.

Example:

```
n = 5
binary (32-bit) = 00000000000000000000000000000101
~n              = 11111111111111111111111111111010
```

That produces a negative number, which is not what we want.

So I realized that I must **only flip the bits that exist in the number's binary representation**.

To achieve this, I create a **mask** consisting only of `1`s that has the same number of bits as `n`.

Example:

```
n    = 101
mask = 111
```

Then I simply XOR the number with the mask.

```
101 XOR 111 = 010
```

This correctly flips the bits.

---

## Approach

Step-by-step process:

1. Handle the edge case:

   * If `n == 0`, the complement should be `1`.

2. Build a mask:

   * Start with `mask = 0`.
   * Keep shifting the mask left and adding `1` until it becomes greater than or equal to `n`.

3. Apply XOR:

   * XOR the mask with the number to flip the bits.

4. Return the result.

Key idea:

```
result = mask ^ n
```

---

## Data Structures Used

No complex data structures are needed.

Only a few integer variables:

* `n` → input number
* `mask` → bitmask used to flip bits

---

## Operations & Behavior Summary

| Operation         | Purpose                      |                             |
| ----------------- | ---------------------------- | --------------------------- |
| Left Shift (`<<`) | Expands the mask to the left |                             |
| Bitwise OR (`     | `)                           | Sets the rightmost bit to 1 |
| Bitwise XOR (`^`) | Flips the bits               |                             |
| Comparison (`<`)  | Controls mask growth         |                             |

---

## Complexity

### Time Complexity

```
O(log n)
```

The loop runs once for each bit of the number.

If `n` has `k` bits, then:

```
k ≈ log2(n)
```

---

### Space Complexity

```
O(1)
```

Only a few variables are used.

No extra memory or data structures are required.

---

# Multi-language Solutions

## C++

```cpp
class Solution {
public:
    int bitwiseComplement(int n) {

        // Edge case: complement of 0 is 1
        if (n == 0) return 1;

        int mask = 0;

        // Build a mask of all 1s with same length as n
        while (mask < n) {
            mask = (mask << 1) | 1;
        }

        // XOR flips the bits
        return mask ^ n;
    }
};
```

---

## Java

```java
class Solution {
    public int bitwiseComplement(int n) {

        // Edge case
        if (n == 0) return 1;

        int mask = 0;

        // Create mask of all 1s
        while (mask < n) {
            mask = (mask << 1) | 1;
        }

        // XOR flips bits
        return mask ^ n;
    }
}
```

---

## JavaScript

```javascript
/**
 * @param {number} n
 * @return {number}
 */
var bitwiseComplement = function(n) {

    // Edge case
    if (n === 0) return 1;

    let mask = 0;

    // Build mask of all 1s
    while (mask < n) {
        mask = (mask << 1) | 1;
    }

    // XOR flips bits
    return mask ^ n;
};
```

---

## Python3

```python
class Solution:
    def bitwiseComplement(self, n: int) -> int:

        # Edge case
        if n == 0:
            return 1

        mask = 0

        # Build mask of all 1s
        while mask < n:
            mask = (mask << 1) | 1

        # XOR flips bits
        return mask ^ n
```

---

## Go

```go
func bitwiseComplement(n int) int {

    // Edge case
    if n == 0 {
        return 1
    }

    mask := 0

    // Build mask of all 1s
    for mask < n {
        mask = (mask << 1) | 1
    }

    // XOR flips bits
    return mask ^ n
}
```

---

# Step-by-step Detailed Explanation

### Step 1 — Handle Edge Case

```
if (n == 0) return 1;
```

Binary representation of `0`:

```
0
```

Complement should be:

```
1
```

So we directly return `1`.

---

### Step 2 — Create a Mask

```
mask = (mask << 1) | 1
```

This operation grows the mask like this:

| Iteration | Binary Mask | Decimal |
| --------- | ----------- | ------- |
| 1         | 1           | 1       |
| 2         | 11          | 3       |
| 3         | 111         | 7       |
| 4         | 1111        | 15      |

The loop continues until the mask becomes greater than or equal to `n`.

Example for `n = 10`:

```
n = 1010
mask = 1111
```

---

### Step 3 — Flip Bits Using XOR

```
result = mask ^ n
```

Example:

```
mask = 1111
n    = 1010
--------------
XOR  = 0101
```

Binary `0101` equals decimal `5`.

So the complement of `10` is `5`.

---

# Examples

### Example 1

Input

```
n = 5
```

Binary

```
101
```

Mask

```
111
```

Result

```
101 XOR 111 = 010
```

Output

```
2
```

---

### Example 2

Input

```
n = 7
```

Binary

```
111
```

Mask

```
111
```

Result

```
111 XOR 111 = 000
```

Output

```
0
```

---

### Example 3

Input

```
n = 10
```

Binary

```
1010
```

Mask

```
1111
```

Result

```
1010 XOR 1111 = 0101
```

Output

```
5
```

---

# How to use / Run locally

### C++

```
g++ solution.cpp -o solution
./solution
```

---

### Java

```
javac Solution.java
java Solution
```

---

### Python

```
python solution.py
```

---

### JavaScript (Node.js)

```
node solution.js
```

---

### Go

```
go run solution.go
```

---

# Notes & Optimizations

1. The algorithm uses **bit manipulation**, which is very efficient.
2. The mask construction ensures we only flip the **relevant bits**.
3. No string conversion is used, which keeps the solution fast.
4. This solution works efficiently even for values close to `10^9`.

Possible micro-optimization:

Some programmers compute the mask directly using:

```
mask = (1 << bitLength) - 1
```

But the loop method is easier to understand and works reliably.

---

# Author

* [Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
