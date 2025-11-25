* # 1015. Smallest Integer Divisible by K

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

Given a positive integer `K`, I need to find the **length** of the smallest positive integer `N` such that:

1. `N` is divisible by `K`, and
2. `N` contains **only the digit `1`** in its decimal representation.
   (Numbers like `1`, `11`, `111`, `1111`, ...)

If no such `N` exists, I must return `-1`.

---

## Constraints

* `1 <= K <= 10^5`
* `N` might not fit into a 64-bit integer, so I **cannot** store `N` directly.

---

## Intuition

I thought:

* I need the smallest number made of only `1`s that is divisible by `K`.
* These numbers grow extremely fast: `1`, `11`, `111`, `1111`, ...
* If I try to store them as normal integers, they will overflow very quickly.

So instead of storing the whole number, I just store its **remainder when divided by `K`**.

Key observations I used:

1. Any number ending in digit `1` can **never** be divisible by `2` or `5`.

   * But all our numbers (`1`, `11`, `111`, ...) end with `1`.
   * So if `K` is divisible by `2` or `5`, the answer is immediately `-1`.

2. If I know the remainder of the current repunit number (like `111`),
   I can compute the remainder of the next one (`1111`) with a simple formula,
   without ever building the big number.

3. There are only `K` possible remainders: `0, 1, 2, ..., K-1`.

   * If I never get remainder `0` within the first `K` tries,
     then some remainder repeats → the remainders enter a cycle → we will **never** reach `0`.
   * So I only need to loop at most `K` times.

---

## Approach

1. **Check impossible cases**

   * If `K % 2 == 0` or `K % 5 == 0`, return `-1`.

2. **Use modulo arithmetic to avoid large numbers**

   * Let `rem` be the current remainder of the repunit number modulo `K`.
   * I start with length `1`, representing the number `1`.
   * For each step:

     * I’m basically appending a `1` at the end:
       If old number is `N`, new number is `N' = N * 10 + 1`.
     * Remainder update:
       `rem = (rem * 10 + 1) % K`.

3. **Check for divisibility**

   * After each update, if `rem == 0`, then the current number (made of all 1s) is divisible by `K`.
   * I return the current length.

4. **Bound the loop**

   * I repeat this at most `K` times (because of the pigeonhole principle on remainders).
   * If I never see `rem == 0` within `K` iterations, I return `-1`.

This way, I never store the huge integer; I only work with remainders.

---

## Data Structures Used

* Only **basic variables**:

  * `int` / `long` for:

    * `rem` – current remainder.
    * `length` – how many `1`s are in the current candidate number.

No arrays, no lists, no extra data structures. Just constant space.

---

## Operations & Behavior Summary

* **Modulo check for impossible case**

  * One check for `k % 2` and `k % 5`.

* **Loop from 1 to K**

  * Each step:

    * Multiply remainder by 10.
    * Add 1.
    * Take modulo `K`.
    * Check if `rem == 0`.

* **Return**

  * If `rem` hits `0` → return current length.
  * If loop finishes → return `-1`.

---

## Complexity

* **Time Complexity:** `O(K)`

  * At most `K` iterations, each doing O(1) arithmetic operations.

* **Space Complexity:** `O(1)`

  * Only using a few integer variables, no extra memory that grows with `K`.

---

## Multi-language Solutions

### C++

```c++
class Solution {
public:
    int smallestRepunitDivByK(int k) {
        // If k is divisible by 2 or 5,
        // no number made only of '1's can be divisible by k.
        if (k % 2 == 0 || k % 5 == 0) return -1;

        int rem = 0;  // current remainder

        // We try at most k times (k different remainders).
        for (int length = 1; length <= k; ++length) {
            // Append '1': N' = N * 10 + 1
            // Remainder update:
            rem = (rem * 10 + 1) % k;

            // If remainder becomes 0, we found our answer.
            if (rem == 0) return length;
        }

        // If no remainder 0 found within k steps, it's impossible.
        return -1;
    }
};
```

---

### Java

```java
class Solution {
    public int smallestRepunitDivByK(int k) {
        // If k has factor 2 or 5, answer is impossible.
        if (k % 2 == 0 || k % 5 == 0) return -1;

        int rem = 0; // current remainder

        // We only need to iterate at most k times.
        for (int length = 1; length <= k; length++) {
            // Build the next remainder when appending '1'
            rem = (rem * 10 + 1) % k;

            if (rem == 0) {
                // Current length is the answer
                return length;
            }
        }

        return -1;
    }
}
```

---

### JavaScript

```javascript
/**
 * @param {number} k
 * @return {number}
 */
var smallestRepunitDivByK = function(k) {
    // If k is multiple of 2 or 5, impossible.
    if (k % 2 === 0 || k % 5 === 0) return -1;

    let rem = 0; // current remainder

    for (let length = 1; length <= k; length++) {
        // Update remainder after appending '1'
        rem = (rem * 10 + 1) % k;

        if (rem === 0) {
            return length;
        }
    }

    return -1;
};
```

---

### Python3

```python
class Solution:
    def smallestRepunitDivByK(self, k: int) -> int:
        # If k has factor 2 or 5, we can never make a multiple of k using only '1's
        if k % 2 == 0 or k % 5 == 0:
            return -1

        rem = 0  # current remainder

        # Try at most k different lengths
        for length in range(1, k + 1):
            # Append '1' and update remainder
            rem = (rem * 10 + 1) % k

            if rem == 0:
                return length

        return -1
```

---

### Go

```go
package main

func smallestRepunitDivByK(k int) int {
    // If k is divisible by 2 or 5, no repunit is divisible by k.
    if k%2 == 0 || k%5 == 0 {
        return -1
    }

    rem := 0 // current remainder

    // Only need to check up to k times.
    for length := 1; length <= k; length++ {
        // Append '1' to the number conceptually:
        rem = (rem*10 + 1) % k

        if rem == 0 {
            return length
        }
    }

    return -1
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

I’ll walk through the **core logic**, which is the same in every language.

Let’s take the C++ code as reference:

```c++
int smallestRepunitDivByK(int k) {
    if (k % 2 == 0 || k % 5 == 0) return -1;
```

1. First, I check if `k` is divisible by `2` or `5`.
2. Any number consisting only of `1`s will always end in `1`.
3. Numbers that are divisible by `2` must end in `0,2,4,6,8`.
4. Numbers that are divisible by `5` must end in `0` or `5`.
5. So an all-ones number can **never** be divisible by such `k`.
6. In that case, I return `-1` immediately.

```c++
    int rem = 0;   // current remainder
```

7. I create a variable `rem` to store the remainder of the current number modulo `k`.
8. Right now, I haven’t started building the number yet, so I set it to `0`.

```c++
    for (int length = 1; length <= k; ++length) {
```

9. I start a loop where `length` represents how many `1`s are in the current repunit number.
10. I go from `1` up to `k`.
11. I don’t need more than `k` steps because there are only `k` possible remainders (`0` to `k-1`).
12. If I don’t hit `0` by then, we are in a remainder cycle and will never get a divisible number.

```c++
        rem = (rem * 10 + 1) % k;
```

13. This line is the heart of the logic.
14. Suppose the current number is `N` and its remainder when divided by `k` is `rem = N % k`.
15. The next repunit is created by appending a `1` at the end: `N' = N * 10 + 1`.
16. The new remainder is:
    `N' % k = (N * 10 + 1) % k`.
17. I can compute this using the old remainder:
    `rem = (rem * 10 + 1) % k`.
18. This way, I don’t store `N` at all, only the remainder.

```c++
        if (rem == 0) {
            return length;  // found smallest length
        }
```

19. After each update, I check if the remainder is `0`.
20. If it is, it means the current repunit number is divisible by `k`.
21. Since I’m increasing `length` one by one from `1`, the first time I see `rem == 0`,
    this is the **smallest** such number.
22. So I return `length`.

```c++
    }
    return -1;
}
```

23. If I finish the loop without ever getting `rem == 0`, I return `-1`.
24. That means there is **no** integer made of only `1`s that is divisible by `k`.

The same logic is implemented in Java, JavaScript, Python, and Go, just using that language’s syntax.

---

## Examples

### Example 1

**Input:** `K = 1`

* `1` is made of only one `1` and is divisible by `1`.
  **Output:** `1`
  **Explanation:** Smallest number is `1` → length `1`.

---

### Example 2

**Input:** `K = 2`

* All-ones numbers: `1, 11, 111, ...` all end in `1`.
* No such number is divisible by `2`.
  **Output:** `-1`.

---

### Example 3

**Input:** `K = 3`

* Check:

  * `1 % 3 = 1`
  * `11 % 3 = 2`
  * `111 % 3 = 0`
* First divisible number is `111`, length = `3`.
  **Output:** `3`.

---

## How to use / Run locally

### 1. Clone this repository

```bash
git clone <your-repo-url>
cd <your-repo-folder>
```

### 2. C++

```bash
g++ -std=c++17 main.cpp -o main
./main
```

* Make sure `main.cpp` contains the C++ `Solution` class and a small driver if you want to test manually.

### 3. Java

```bash
javac Main.java
java Main
```

* `Main.java` should create an instance of `Solution` and call `smallestRepunitDivByK`.

### 4. JavaScript (Node.js)

```bash
node main.js
```

* `main.js` should require or include the `smallestRepunitDivByK` function and test it.

### 5. Python3

```bash
python3 main.py
```

* `main.py` should create an instance of `Solution` and call `smallestRepunitDivByK`.

### 6. Go

```bash
go run main.go
```

* `main.go` should call `smallestRepunitDivByK` from `main()` and print the result.

---

## Notes & Optimizations

* **Early rejection** of `k % 2 == 0 || k % 5 == 0` saves time.
* Using **only remainders** keeps the numbers small and avoids overflow.
* Looping only up to `K` uses the pigeonhole principle to guarantee correctness and termination.
* This is already optimal in both time (`O(K)`) and space (`O(1)`) for this problem setting.

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
