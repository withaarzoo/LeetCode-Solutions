# 3345. Smallest Divisible Digit Product I

A clean and beginner-friendly solution for **LeetCode 3345 - Smallest Divisible Digit Product I**. This repository explains the problem in simple English, walks through the thought process behind the solution, and provides implementations in **C++, Java, JavaScript, Python3, and Go**.

This README is designed for developers preparing for coding interviews, practicing Data Structures and Algorithms (DSA), or learning competitive programming.

---

## Table of Contents

- [Problem Summary](#problem-summary)
- [Constraints](#constraints)
- [Intuition](#intuition)
- [Approach](#approach)
- [Data Structures Used](#data-structures-used)
- [Operations & Behavior Summary](#operations--behavior-summary)
- [Complexity](#complexity)
- [Multi-language Solutions](#multi-language-solutions)
  - [C++](#c)
  - [Java](#java)
  - [JavaScript](#javascript)
  - [Python3](#python3)
  - [Go](#go)
- [Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)](#step-by-step-detailed-explanation-c-java-javascript-python3-go)
- [Examples](#examples)
- [How to Use / Run Locally](#how-to-use--run-locally)
- [Notes & Optimizations](#notes--optimizations)
- [Author](#author)

---

## Problem Summary

In this problem, we are given two integers:

- `n` — the starting number
- `t` — the required divisor

Our task is to find the **smallest number greater than or equal to `n`** whose **product of digits** is divisible by `t`.

For every number we check:

1. Calculate the product of all of its digits.
2. Check whether that product is divisible by `t`.
3. Return the first number that satisfies the condition.

Since the constraints are very small, a straightforward solution is enough.

This problem is a good beginner-level exercise for practicing:

- Brute Force
- Digit Manipulation
- Number Theory Basics
- Simulation
- Competitive Programming

---

## Constraints

| Constraint | Value |
|------------|-------|
| `1 ≤ n ≤ 100` | Starting number |
| `1 ≤ t ≤ 10` | Required divisor |

---

## Intuition

The first thing I noticed was how small the constraints are.

Since `n` is at most `100`, there is no need for complicated mathematics or optimization. Instead of trying to predict the answer, I can simply test each number one by one.

For every number, I calculate the product of its digits. If that product is divisible by `t`, I immediately return the number because I am checking numbers in increasing order.

This makes the solution both simple and reliable.

---

## Approach

I solve the problem in the following steps:

1. Start checking from `n`.
2. Copy the current number into a temporary variable.
3. Extract every digit using modulo (`% 10`).
4. Multiply all digits together.
5. Check whether the product is divisible by `t`.
6. If it is, return the current number.
7. Otherwise, move to the next number and repeat.

Since the search starts from the smallest possible value, the first valid number is guaranteed to be the answer.

---

## Data Structures Used

This solution does not require any special data structures.

| Data Structure | Purpose |
|---------------|---------|
| Integer variables | Store the current number, digit product, and temporary value while extracting digits |

The algorithm works using only a few variables, making it memory efficient.

---

## Operations & Behavior Summary

The algorithm performs the following operations:

1. Begin with the starting value `n`.
2. Compute the product of all digits.
3. Test whether the product is divisible by `t`.
4. If the condition is satisfied, return the number.
5. Otherwise, increase the number by one.
6. Repeat until a valid answer is found.

Because numbers containing the digit `0` produce a digit product of `0`, and `0` is divisible by every positive integer, an answer is always guaranteed to exist.

---

## Complexity

| Complexity | Value | Explanation |
|------------|-------|-------------|
| Time Complexity | `O(m × d)` | `m` is the number of integers checked before finding the answer, and `d` is the number of digits in each integer. |
| Space Complexity | `O(1)` | Only a few integer variables are used. No extra data structures are required. |

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int smallestNumber(int n, int t) {
        // Keep checking numbers starting from n
        while (true) {
            int product = 1;
            int x = n;

            // Calculate the product of all digits
            while (x > 0) {
                product *= (x % 10);
                x /= 10;
            }

            // If the product is divisible by t, this is the answer
            if (product % t == 0)
                return n;

            // Otherwise check the next number
            n++;
        }
    }
};
```

### Java

```java
class Solution {
    public int smallestNumber(int n, int t) {
        // Keep checking numbers starting from n
        while (true) {
            int product = 1;
            int x = n;

            // Calculate the product of all digits
            while (x > 0) {
                product *= (x % 10);
                x /= 10;
            }

            // Return the first valid number
            if (product % t == 0)
                return n;

            // Try the next number
            n++;
        }
    }
}
```

### JavaScript

```javascript
/**
 * @param {number} n
 * @param {number} t
 * @return {number}
 */
var smallestNumber = function(n, t) {
    // Keep checking numbers until a valid one is found
    while (true) {
        let product = 1;
        let x = n;

        // Calculate the product of all digits
        while (x > 0) {
            product *= x % 10;
            x = Math.floor(x / 10);
        }

        // Return the first valid number
        if (product % t === 0) {
            return n;
        }

        // Check the next number
        n++;
    }
};
```

### Python3

```python
class Solution:
    def smallestNumber(self, n: int, t: int) -> int:
        # Keep checking numbers starting from n
        while True:
            product = 1
            x = n

            # Calculate the product of all digits
            while x > 0:
                product *= x % 10
                x //= 10

            # Return the first valid number
            if product % t == 0:
                return n

            # Try the next number
            n += 1
```

### Go

```go
func smallestNumber(n int, t int) int {
    // Keep checking numbers starting from n
    for {
        product := 1
        x := n

        // Calculate the product of all digits
        for x > 0 {
            product *= x % 10
            x /= 10
        }

        // Return the first valid number
        if product%t == 0 {
            return n
        }

        // Check the next number
        n++
    }
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The overall logic is exactly the same in every programming language. Only the syntax changes.

First, the algorithm starts checking from the given value of `n`.

A temporary copy of the number is created. This allows the digits to be extracted without changing the original value that may eventually be returned.

The digit product starts from `1` because multiplication needs an identity value. Starting from `0` would make every product equal to zero.

The last digit is extracted using the modulo operator.

After processing a digit, the last digit is removed using integer division.

This continues until every digit has been processed.

Once the digit product has been calculated, the algorithm checks whether the product is divisible by `t`.

If it is divisible, the current number is returned immediately.

Otherwise, the number is increased by one, and the same process repeats.

One important edge case is numbers containing the digit `0`.

For example:

- `105`
- Product = `1 × 0 × 5 = 0`

Since `0` is divisible by every positive integer, any such number automatically satisfies the condition.

Because the constraints are very small, this simple simulation is already the optimal solution for this problem.

---

## Examples

### Example 1

**Input**

```text
n = 10
t = 2
```

**Output**

```text
10
```

**Trace**

- Product of digits = `1 × 0 = 0`
- `0 % 2 = 0`
- Return `10`

---

### Example 2

**Input**

```text
n = 15
t = 3
```

**Output**

```text
16
```

**Trace**

- `15 → 1 × 5 = 5`
- `5` is not divisible by `3`
- Check `16`
- `1 × 6 = 6`
- `6 % 3 = 0`
- Return `16`

---

### Example 3

**Input**

```text
n = 22
t = 8
```

**Output**

```text
24
```

**Trace**

- `22 → 2 × 2 = 4`
- Not divisible by `8`
- `23 → 2 × 3 = 6`
- Not divisible by `8`
- `24 → 2 × 4 = 8`
- Divisible by `8`
- Return `24`

---

## How to Use / Run Locally

Clone the repository:

```bash
git clone https://github.com/your-username/your-repository.git
```

Move into the project directory:

```bash
cd your-repository
```

### C++

Compile:

```bash
g++ main.cpp -o main
```

Run:

```bash
./main
```

---

### Java

Compile:

```bash
javac Main.java
```

Run:

```bash
java Main
```

---

### JavaScript

Run:

```bash
node main.js
```

---

### Python3

Run:

```bash
python main.py
```

or

```bash
python3 main.py
```

---

### Go

Run:

```bash
go run main.go
```

---

## Notes & Optimizations

- This is a brute force simulation problem.
- The constraints are intentionally small, so checking numbers one by one is completely acceptable.
- No advanced algorithms or mathematical tricks are necessary.
- The solution uses constant extra memory.
- Numbers containing the digit `0` immediately produce a digit product of `0`, which is divisible by every positive integer.
- For larger constraints, a different strategy would be needed because brute force would become inefficient. However, for this problem, the simple approach is both clean and fast enough.

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
