# Problem Title: 3370. Smallest Number With All Set Bits

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

Given a positive integer `n`, return the smallest number `x` such that `x >= n` and the binary representation of `x` contains **only set bits** — i.e., `x` must look like `1`, `11`, `111`, `1111`, ... in binary. Those numbers are exactly `2^k - 1` for some integer `k >= 1`.

Examples:

* Input: `n = 5` → Output: `7` (binary `111`)
* Input: `n = 10` → Output: `15` (binary `1111`)
* Input: `n = 3` → Output: `3` (binary `11`)

---

## Constraints

* `1 <= n <= 1000` (as given in the problem statement).
* For general safety, solutions use standard integer arithmetic; 64-bit types are used in examples where appropriate.

---

## Intuition

I thought about which numbers have all bits set in binary. They are exactly numbers of the form `2^k - 1`. If I want the smallest `x` ≥ `n` with that property, I just need to find the smallest `k` such that:

```
2^k - 1 >= n
```

So I can increase `k` starting at `1` and compute `2^k - 1` until it reaches or exceeds `n`. Because `k` grows roughly as the number of bits in `n`, the loop is short — `O(log n)` steps.

---

## Approach

1. Initialize `k = 1`.
2. Compute `val = (1 << k) - 1` (which is `2^k - 1`).
3. If `val >= n`, return `val`.
4. Otherwise increment `k` and repeat.

This is simple, easy to read, and fast for the given constraints. I keep the arithmetic in a safe integer type (e.g., long/long long) in languages that require explicit size to avoid overflow in more general settings.

---

## Data Structures Used

* Primitive integer types only (int/long/long long).
* No arrays, lists, or extra data structures are required.

---

## Operations & Behavior Summary

* Bit-shift operation `(1 << k)` is used to compute `2^k` efficiently.
* Subtracting 1 gives a number consisting of `k` ones in binary (`111...1`).
* The program increments `k` iteratively until the condition is met.

---

## Complexity

* **Time Complexity:** `O(log n)` — `k` increases until `2^k - 1 >= n`. `k` ≈ number of bits of `n`, i.e., about `log2(n)`.
* **Space Complexity:** `O(1)` — only a constant number of integer variables are used.

---

## Multi-language Solutions

### C++

```c++
/*
 * 3370. Smallest Number With All Set Bits
 * C++ solution: iterate k until (2^k - 1) >= n
 */
class Solution {
public:
    int smallestNumber(int n) {
        // Use long long for safety (though n <= 1000 here)
        long long k = 1;
        while (true) {
            long long val = (1LL << k) - 1; // 2^k - 1
            if (val >= n) return (int)val;
            k++;
        }
    }
};
```

### Java

```java
/*
 * 3370. Smallest Number With All Set Bits
 * Java solution: iterate k until (2^k - 1) >= n
 */
class Solution {
    public int smallestNumber(int n) {
        int k = 1;
        while (true) {
            long val = (1L << k) - 1; // use long to be safe
            if (val >= n) return (int)val;
            k++;
        }
    }
}
```

### JavaScript

```javascript
/**
 * 3370. Smallest Number With All Set Bits
 * JavaScript solution: iterate k until (2^k - 1) >= n
 * @param {number} n
 * @return {number}
 */
var smallestNumber = function(n) {
    let k = 1;
    while (true) {
        let val = (1 << k) - 1; // safe for n <= 1000
        if (val >= n) return val;
        k++;
    }
};
```

### Python3

```python
# 3370. Smallest Number With All Set Bits
# Python solution: iterate k until (2^k - 1) >= n

class Solution:
    def smallestNumber(self, n: int) -> int:
        k = 1
        while True:
            val = (1 << k) - 1  # 2**k - 1
            if val >= n:
                return val
            k += 1
```

### Go

```go
// 3370. Smallest Number With All Set Bits
// Go solution: iterate k until (2^k - 1) >= n

package main

func smallestNumber(n int) int {
    k := 1
    for {
        val := (1 << k) - 1
        if val >= n {
            return val
        }
        k++
    }
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

I'll walk through the logic using the C++ implementation for line-by-line clarity; other languages follow the same steps.

C++ code:

```c++
long long k = 1;
while (true) {
    long long val = (1LL << k) - 1; // 2^k - 1
    if (val >= n) return (int)val;
    k++;
}
```

Line-by-line explanation:

1. `long long k = 1;`

   * I start searching from `k = 1` because the smallest "all ones" number is `1` (binary `1` = `2^1 - 1`).

2. `while (true) {`

   * I create an infinite loop that I'll break out of by returning when the condition is satisfied. This loop always terminates because `2^k - 1` grows without bound as `k` increases.

3. `long long val = (1LL << k) - 1;`

   * `(1LL << k)` computes `2^k` using a left shift on a `long long` value. Subtracting `1` gives `2^k - 1`, which in binary is `k` ones (for example `k=4` -> `1111` (binary) = 15 decimal).

4. `if (val >= n) return (int)val;`

   * If `val` is at least `n`, then `val` is the smallest number with `k` ones that is ≥ `n` because we're increasing `k` from smallest to larger. So we return `val`.

5. `k++;`

   * If `val < n`, we increase `k` and continue the search.

**Why this always finds the smallest?**
We check `k` in increasing order: `1, 2, 3, ...`. For each `k` the number `2^k - 1` is strictly increasing with `k`. The first `k` with `2^k - 1 >= n` must be the minimal `k`, hence the minimal `2^k - 1` that satisfies the requirement.

**Example (n = 10):**

* `k=1` → `val = 1` (1 < 10)
* `k=2` → `val = 3` (3 < 10)
* `k=3` → `val = 7` (7 < 10)
* `k=4` → `val = 15` (15 >= 10) → return `15`

---

### Language-specific notes on the steps

* **Java/Python/Go/JavaScript**: same iterative idea, use the appropriate integer types. For Java and C++ I used `long`/`long long` for safety. For the constraints given (`n <= 1000`) all int sizes are safe, but I still show a robust style.
* **JavaScript**: bitwise shifts are on 32-bit signed integers internally. For the problem constraints it's fine. If `n` were large (beyond 2^31), we'd need a different approach (e.g., BigInt).

---

## Examples

| Input | Output | Explanation          |
| ----- | -----: | -------------------- |
| `5`   |    `7` | `7` = binary `111`   |
| `10`  |   `15` | `15` = binary `1111` |
| `3`   |    `3` | `3` = binary `11`    |

---

## How to use / Run locally

### C++

1. Put the `Solution` class into your LeetCode submission area or into a `.cpp` file wrapper with a `main()` for local testing.
2. Compile with `g++ -std=c++17 file.cpp -O2 -o solution` and run `./solution` (if you added a `main` function for tests).

### Java

1. Put the `Solution` class into the LeetCode submission area or into `Solution.java` with a `main` driver.
2. Compile with `javac Solution.java` and run `java Solution` (if you added `main`).

### JavaScript (Node)

1. Save function in a file, add test harness, e.g., call `console.log(smallestNumber(10));`.
2. Run with `node file.js`.

### Python3

1. Save class into a `.py` file and add test code, e.g.:

```python
print(Solution().smallestNumber(10))
```

2. Run with `python3 file.py`.

### Go

1. Place function in `main.go`, add a `main()` function to call `smallestNumber`.
2. Build and run:

```
go run main.go
```

---

## Notes & Optimizations

* Because the search runs for roughly `log2(n)` iterations, this is already optimal in time complexity for the simple approach.
* A constant-time variant: you can compute `k = ceil(log2(n + 1))` and then `val = 2^k - 1`. However, using floating-point `log2` may need careful rounding and corner-case handling (floating inaccuracies). The iterative bit-shift method avoids floating-point issues and is trivial to reason about.
* For very large `n` (beyond typical 32-bit limits), use 64-bit types (C++ `long long`, Java `long`) or `BigInt` in JavaScript.
* For the given constraint `n <= 1000`, the chosen solution is simple and perfectly efficient.

---

## Author

* [Md. Aarzoo Islam](https://bento.me/withaarzoo)
