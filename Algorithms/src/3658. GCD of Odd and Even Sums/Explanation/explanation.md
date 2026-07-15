# 3658. GCD of Odd and Even Sums - LeetCode Solution

A simple mathematical solution for LeetCode 3658. GCD of Odd and Even Sums. This repository explains the intuition, approach, complexity analysis, and provides multi-language solution templates in C++, Java, JavaScript, Python, and Go. If you're preparing for coding interviews or improving your DSA and competitive programming skills, this problem is a great example of how mathematical observations can replace unnecessary computation.

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

You are given an integer `n`.

Your task is to calculate the Greatest Common Divisor (GCD) of two values:

- The sum of the first `n` positive odd numbers.
- The sum of the first `n` positive even numbers.

Instead of generating the sequences and calculating their sums one by one, the goal is to find the answer using the most efficient algorithm possible.

The expected output is a single integer representing the GCD of these two sums.

This is a mathematical problem that appears simple at first but becomes much easier after identifying the hidden pattern.

---

## Constraints

| Constraint | Value |
|------------|-------|
| `1 ≤ n ≤ 1000` |

---

## Intuition

The first thing I noticed was that calculating every odd and even number separately felt unnecessary.

I remembered two common mathematical formulas:

- The sum of the first `n` odd numbers is `n²`.
- The sum of the first `n` even numbers is `n × (n + 1)`.

Once I replaced the original sums with these formulas, the problem became a simple GCD calculation.

After factoring out the common term, I realized the answer always simplifies to `n`.

This observation removes the need for loops, recursion, or even a GCD function.

---

## Approach

I solved the problem using a direct mathematical observation.

First, I expressed both sums using their known formulas.

- Odd sum = `n²`
- Even sum = `n × (n + 1)`

Next, I rewrote the expression as:

`GCD(n², n(n + 1))`

Both values contain a common factor of `n`, so I factored it out.

The remaining expression becomes:

`GCD(n, n + 1)`

Since consecutive integers are always coprime, their GCD is always `1`.

That leaves the final answer as simply:

`n`

The implementation only needs to return the input value.

---

## Data Structures Used

No data structures are used.

This solution relies entirely on mathematical properties instead of storing or processing any data.

Memory usage stays constant throughout execution.

---

## Operations & Behavior Summary

1. Read the input integer `n`.
2. Use the mathematical formulas for odd and even sums.
3. Simplify the GCD expression.
4. Observe that consecutive integers always have a GCD of `1`.
5. Return `n` as the final answer.

No iteration, recursion, or additional computation is required.

---

## Complexity

| Complexity | Value | Explanation |
|------------|-------|-------------|
| Time Complexity | **O(1)** | The algorithm performs only constant-time operations regardless of the input value. |
| Space Complexity | **O(1)** | No extra memory or additional data structures are used. |

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int gcdOfOddEvenSums(int n) {
        // From the mathematical proof:
        // GCD(n^2, n*(n+1)) = n
        return n;
    }
};
```

### Java

```java
class Solution {
    public int gcdOfOddEvenSums(int n) {
        // From the mathematical proof:
        // GCD(n^2, n*(n+1)) = n
        return n;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number} n
 * @return {number}
 */
var gcdOfOddEvenSums = function(n) {
    // From the mathematical proof:
    // GCD(n^2, n*(n+1)) = n
    return n;
};
```

### Python3

```python
class Solution:
    def gcdOfOddEvenSums(self, n: int) -> int:
        # From the mathematical proof:
        # GCD(n^2, n*(n+1)) = n
        return n
```

### Go

```go
func gcdOfOddEvenSums(n int) int {
    // From the mathematical proof:
    // GCD(n*n, n*(n+1)) = n
    return n
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic is exactly the same in every programming language.

The only difference is the syntax used to write the function.

The algorithm begins by accepting the integer `n`.

Instead of computing the odd and even sums manually, it uses the mathematical identities:

- Sum of first `n` odd numbers = `n²`
- Sum of first `n` even numbers = `n × (n + 1)`

The expression then becomes:

`GCD(n², n(n + 1))`

Both values contain the factor `n`.

Factoring it out gives:

`n × GCD(n, n + 1)`

Since `n` and `n + 1` are consecutive numbers, they always share only one common divisor.

Therefore:

`GCD(n, n + 1) = 1`

The final answer becomes:

`n × 1 = n`

Because of this mathematical simplification, every language implementation only needs to return the input value.

No loops are needed.

No arrays are created.

No recursion is used.

No GCD calculation is required.

This makes the solution both elegant and optimal.

---

## Examples

### Example 1

**Input**

```text
n = 4
```

**Output**

```text
4
```

**Trace**

First 4 odd numbers:

`1 + 3 + 5 + 7 = 16`

First 4 even numbers:

`2 + 4 + 6 + 8 = 20`

`GCD(16, 20) = 4`

---

### Example 2

**Input**

```text
n = 5
```

**Output**

```text
5
```

**Trace**

Odd sum:

`1 + 3 + 5 + 7 + 9 = 25`

Even sum:

`2 + 4 + 6 + 8 + 10 = 30`

`GCD(25, 30) = 5`

---

### Example 3

**Input**

```text
n = 1
```

**Output**

```text
1
```

**Trace**

Odd sum:

`1`

Even sum:

`2`

`GCD(1, 2) = 1`

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

Run

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

---

## Notes & Optimizations

- This problem is completely based on mathematical reasoning.
- A brute-force approach would calculate every odd and even number individually, but that work is unnecessary.
- Even calling a standard GCD function is not required after simplifying the expression.
- The solution always runs in constant time.
- It also uses constant memory.
- This is the best possible solution for the given problem.
- Problems like this are good examples of how recognizing mathematical patterns can significantly simplify an algorithm.

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
