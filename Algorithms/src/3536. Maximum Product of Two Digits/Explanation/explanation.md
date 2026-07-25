# 3536. Maximum Product of Two Digits

A simple and efficient solution for **LeetCode 3536 - Maximum Product of Two Digits**. This repository explains the intuition, approach, complexity analysis, and provides solutions in **C++, Java, JavaScript, Python, and Go**. The algorithm runs in linear time with constant extra space, making it an optimal solution for this digit manipulation problem.

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
- [Step-by-step Detailed Explanation](#step-by-step-detailed-explanation-c-java-javascript-python3-go)
- [Examples](#examples)
- [How to Use / Run Locally](#how-to-use--run-locally)
- [Notes & Optimizations](#notes--optimizations)
- [Author](#author)

---

## Problem Summary

You are given a positive integer `n`.

Your task is to find the **maximum product** that can be formed by multiplying any two digits present in the number.

Each digit can only be used once unless it appears multiple times in the number. For example, if the number is `22`, then both `2`s can be used because they are different occurrences of the same digit.

The goal is to return the largest possible product among all valid pairs of digits.

---

## Constraints

| Constraint | Value |
|------------|-------|
| Minimum value of `n` | `10` |
| Maximum value of `n` | `10^9` |

---

## Intuition

The first thing I noticed was that the largest product will always come from the two largest digits.

Instead of storing every digit and checking every possible pair, I can simply scan the number once while keeping track of the largest digit and the second largest digit seen so far.

After processing every digit, multiplying these two values immediately gives the answer.

This avoids unnecessary work and keeps the solution both simple and efficient.

---

## Approach

I solve the problem in one pass through the digits.

1. Create two variables to store the largest and second largest digits.
2. Extract the last digit using the modulo operator.
3. Compare it with the current largest digit.
4. If it becomes the new largest digit, move the previous largest to the second position.
5. Otherwise, update only the second largest digit if needed.
6. Remove the last digit from the number.
7. Repeat until all digits have been processed.
8. Return the product of the two largest digits.

This approach visits every digit exactly once.

---

## Data Structures Used

| Data Structure | Purpose |
|---------------|---------|
| Integer Variables | Store the current largest and second largest digits while scanning the number |

No arrays, lists, stacks, queues, hash maps, or other data structures are required.

---

## Operations & Behavior Summary

The algorithm performs the following operations:

1. Start with two variables representing the largest two digits found so far.
2. Read the last digit of the number.
3. Compare the digit with the current maximum values.
4. Update the largest or second largest digit whenever necessary.
5. Remove the processed digit.
6. Continue until no digits remain.
7. Multiply the two largest digits.
8. Return the final answer.

---

## Complexity

| Metric | Complexity | Explanation |
|--------|------------|-------------|
| Time Complexity | **O(d)** | `d` is the total number of digits in the given integer. Every digit is processed exactly once. |
| Space Complexity | **O(1)** | Only two integer variables are used regardless of the input size. |

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int maxProduct(int n) {
        // Store the largest digit found so far
        int first = 0;

        // Store the second largest digit found so far
        int second = 0;

        // Process every digit of the number
        while (n > 0) {
            // Extract the last digit
            int digit = n % 10;

            // If this digit becomes the new largest
            if (digit >= first) {
                // Old largest becomes second largest
                second = first;
                first = digit;
            }
            // Otherwise check if it should become second largest
            else if (digit > second) {
                second = digit;
            }

            // Remove the last digit
            n /= 10;
        }

        // Return the maximum product
        return first * second;
    }
};
```

### Java

```java
class Solution {
    public int maxProduct(int n) {

        // Store the largest digit
        int first = 0;

        // Store the second largest digit
        int second = 0;

        // Process every digit
        while (n > 0) {

            // Extract the last digit
            int digit = n % 10;

            // Update largest and second largest if needed
            if (digit >= first) {
                second = first;
                first = digit;
            } else if (digit > second) {
                second = digit;
            }

            // Remove the last digit
            n /= 10;
        }

        // Return the answer
        return first * second;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number} n
 * @return {number}
 */
var maxProduct = function(n) {

    // Store the largest digit
    let first = 0;

    // Store the second largest digit
    let second = 0;

    // Process every digit
    while (n > 0) {

        // Extract the last digit
        const digit = n % 10;

        // Update the two largest digits
        if (digit >= first) {
            second = first;
            first = digit;
        } else if (digit > second) {
            second = digit;
        }

        // Remove the last digit
        n = Math.floor(n / 10);
    }

    // Return the maximum product
    return first * second;
};
```

### Python3

```python
class Solution:
    def maxProduct(self, n: int) -> int:

        # Store the largest digit
        first = 0

        # Store the second largest digit
        second = 0

        # Process every digit
        while n > 0:

            # Extract the last digit
            digit = n % 10

            # Update the two largest digits
            if digit >= first:
                second = first
                first = digit
            elif digit > second:
                second = digit

            # Remove the last digit
            n //= 10

        # Return the maximum product
        return first * second
```

### Go

```go
func maxProduct(n int) int {

 // Store the largest digit
 first := 0

 // Store the second largest digit
 second := 0

 // Process every digit
 for n > 0 {

  // Extract the last digit
  digit := n % 10

  // Update the two largest digits
  if digit >= first {
   second = first
   first = digit
  } else if digit > second {
   second = digit
  }

  // Remove the last digit
  n /= 10
 }

 // Return the maximum product
 return first * second
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic is identical in all five programming languages. Only the syntax changes.

First, two integer variables are created to store the largest and second largest digits found so far.

The algorithm then starts processing the number from right to left.

During every iteration, the last digit is extracted using the modulo operator (`% 10`).

This digit is compared with the current largest digit.

If it is greater than or equal to the largest digit, the current largest value moves into the second-largest position, and the new digit becomes the largest.

Otherwise, the digit is compared with the second-largest value.

If it is larger than the second-largest digit, only the second-largest value is updated.

After processing the digit, the last digit is removed by dividing the number by `10`.

The same process repeats until no digits remain.

Once the loop finishes, the two largest digits have already been identified.

The final step is simply multiplying these two digits and returning the result.

Since every language follows this exact sequence of operations, the behavior, correctness, and complexity remain identical across C++, Java, JavaScript, Python3, and Go.

---

## Examples

### Example 1

**Input**

```text
n = 31
```

**Output**

```text
3
```

**Trace**

Digits are `3` and `1`.

The maximum product is:

```text
3 × 1 = 3
```

---

### Example 2

**Input**

```text
n = 22
```

**Output**

```text
4
```

**Trace**

Digits are `2` and `2`.

Since both digits exist separately, both can be used.

```text
2 × 2 = 4
```

---

### Example 3

**Input**

```text
n = 124
```

**Output**

```text
8
```

**Trace**

Digits are:

```text
1, 2, 4
```

Largest two digits:

```text
4 and 2
```

Maximum product:

```text
4 × 2 = 8
```

---

## How to Use / Run Locally

Clone the repository.

```bash
git clone <repository-url>
```

Move into the project folder.

```bash
cd <repository-name>
```

### C++

Compile

```bash
g++ solution.cpp -o solution
```

Run

```bash
./solution
```

---

### Java

Compile

```bash
javac Solution.java
```

Run

```bash
java Solution
```

---

### JavaScript

Run using Node.js

```bash
node solution.js
```

---

### Python3

Run

```bash
python solution.py
```

or

```bash
python3 solution.py
```

---

### Go

Run

```bash
go run solution.go
```

Build

```bash
go build solution.go
```

---

## Notes & Optimizations

- This is already the optimal solution for the problem.
- Every digit is processed only once.
- No extra memory is required apart from two integer variables.
- The solution naturally handles repeated digits.
- There is no need to generate all possible digit pairs.
- An alternative approach would be to store every digit in an array and sort it, but that would require extra memory and additional processing.
- Keeping track of only the two largest digits is both simpler and more efficient.

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
