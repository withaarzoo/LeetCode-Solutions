# 3513. Number of Unique XOR Triplets I

A simple and optimized solution for LeetCode 3513 - Number of Unique XOR Triplets I. This repository explains the intuition, approach, complexity analysis, and provides multi-language solutions in C++, Java, JavaScript, Python, and Go. The solution uses bit manipulation and mathematical observation to achieve constant time complexity.

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

You are given an integer array `nums` that is a permutation of numbers from `1` to `n`.

A XOR triplet is created by choosing three indices `(i, j, k)` such that `i <= j <= k` and computing:

`nums[i] XOR nums[j] XOR nums[k]`

The goal is to count how many **unique XOR values** can be produced from every valid triplet.

The challenge is that `n` can be as large as `100,000`, so checking every possible triplet is impossible. Instead, the solution depends on understanding the mathematical properties of XOR and the structure of the permutation.

---

## Constraints

| Constraint | Value |
|------------|-------|
| `1 <= n == nums.length <= 10^5` | Yes |
| `1 <= nums[i] <= n` | Yes |
| `nums` is a permutation of integers from `1` to `n` | Yes |

---

## Intuition

My first idea was to generate every possible triplet and store all XOR values inside a set. That works for very small inputs, but the number of triplets grows far too quickly.

Then I noticed something important.

The array is always a permutation of numbers from `1` to `n`. Since every value appears exactly once, the order of the elements does not matter.

After trying several small examples, I realized that once `n` becomes at least `3`, every XOR value inside a complete binary range becomes possible. That means I no longer need to examine triplets individually.

The problem becomes a simple bit manipulation question.

---

## Approach

I solve the problem in the following steps.

1. Find the size of the array.
2. Handle the small cases separately.
3. If `n` is `1` or `2`, return `n`.
4. Otherwise, calculate how many bits are needed to represent `n`.
5. Every value from `0` to `(2^bits - 1)` can be generated.
6. Return `2^bits`.

Instead of searching through millions of triplets, the solution uses one mathematical observation to produce the answer immediately.

---

## Data Structures Used

This solution does not require any complex data structures.

| Data Structure | Purpose |
|---------------|---------|
| Integer Variables | Store the array size and bit count |

No arrays, sets, maps, stacks, queues, or hash tables are needed.

---

## Operations & Behavior Summary

The algorithm performs only a few simple operations.

- Read the length of the array.
- Check whether the array size is `1` or `2`.
- Count the number of bits needed to represent `n`.
- Compute `2^bits`.
- Return the final answer.

No XOR triplets are actually generated during execution.

---

## Complexity

| Type | Complexity | Explanation |
|------|------------|-------------|
| Time Complexity | **O(1)** | Only a few arithmetic and bit operations are performed. |
| Space Complexity | **O(1)** | No extra data structures are used. |

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int uniqueXorTriplets(vector<int>& nums) {
        // Get the size of the permutation
        int n = nums.size();

        // Handle the small cases separately
        if (n <= 2) return n;

        // Find how many bits are needed to represent n
        int bits = 0;
        int x = n;
        while (x) {
            bits++;
            x >>= 1;
        }

        // Total values in the range [0, 2^bits - 1]
        return 1 << bits;
    }
};
```

### Java

```java
class Solution {
    public int uniqueXorTriplets(int[] nums) {
        // Get the size of the permutation
        int n = nums.length;

        // Small cases
        if (n <= 2) return n;

        // Count the number of bits in n
        int bits = 0;
        int x = n;
        while (x > 0) {
            bits++;
            x >>= 1;
        }

        // Every value in [0, 2^bits - 1] is possible
        return 1 << bits;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} nums
 * @return {number}
 */
var uniqueXorTriplets = function(nums) {
    // Length of the permutation
    const n = nums.length;

    // Handle small cases
    if (n <= 2) return n;

    // Count the number of bits in n
    let bits = 0;
    let x = n;
    while (x > 0) {
        bits++;
        x >>= 1;
    }

    // Number of possible XOR values
    return 1 << bits;
};
```

### Python3

```python
class Solution:
    def uniqueXorTriplets(self, nums: List[int]) -> int:
        # Length of the permutation
        n = len(nums)

        # Small cases
        if n <= 2:
            return n

        # bit_length() gives the number of bits needed to represent n
        return 1 << n.bit_length()
```

### Go

```go
func uniqueXorTriplets(nums []int) int {
 // Length of the permutation
 n := len(nums)

 // Handle the small cases
 if n <= 2 {
  return n
 }

 // Count how many bits are needed to represent n
 bits := 0
 x := n
 for x > 0 {
  bits++
  x >>= 1
 }

 // Every value in [0, 2^bits - 1] can be generated
 return 1 << bits
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic is exactly the same in every programming language.

The first step is finding the length of the input array. Since the array is guaranteed to be a permutation of numbers from `1` to `n`, the actual order of the elements does not matter.

Next, the algorithm handles the smallest inputs separately.

When `n` equals `1`, only one XOR value can exist.

When `n` equals `2`, there are only two unique XOR values.

These two cases do not follow the general mathematical pattern, so handling them separately keeps the solution simple.

For every larger value of `n`, the algorithm counts how many binary bits are needed to represent `n`.

For example:

- `3` needs `2` bits.
- `5` needs `3` bits.
- `8` needs `4` bits.

Once the bit length is known, the answer becomes straightforward.

Every integer from `0` to `(2^bits - 1)` can be produced by some valid XOR triplet.

Since there are exactly `2^bits` numbers in that range, the final answer is simply:

`2^bits`

Although the syntax changes slightly between C++, Java, JavaScript, Python, and Go, the algorithm and complexity remain exactly the same.

---

## Examples

### Example 1

**Input**

```text
nums = [1,2]
```

**Output**

```text
2
```

**Trace**

- n = 2
- Small case
- Return 2

---

### Example 2

**Input**

```text
nums = [3,1,2]
```

**Output**

```text
4
```

**Trace**

- n = 3
- Binary representation of 3 is `11`
- Bit length = 2
- Answer = 2² = 4

---

### Example 3

**Input**

```text
nums = [1,2,3,4,5]
```

**Output**

```text
8
```

**Trace**

- n = 5
- Binary representation of 5 is `101`
- Bit length = 3
- Answer = 2³ = 8

---

## How to Use / Run Locally

Clone the repository.

```bash
git clone https://github.com/your-username/your-repository.git
```

Move into the project directory.

```bash
cd your-repository
```

### Compile and Run C++

Compile:

```bash
g++ solution.cpp -o solution
```

Run:

```bash
./solution
```

---

### Compile and Run Java

Compile:

```bash
javac Solution.java
```

Run:

```bash
java Solution
```

---

### Run JavaScript

```bash
node solution.js
```

---

### Run Python3

```bash
python solution.py
```

or

```bash
python3 solution.py
```

---

### Run Go

```bash
go run solution.go
```

---

## Notes & Optimizations

- The solution depends on a mathematical observation instead of brute force.
- No XOR triplets are generated during execution.
- The array values never need to be inspected because every input is guaranteed to be a permutation.
- This makes the solution significantly faster than any simulation-based approach.
- Both time complexity and space complexity are constant.
- This is one of those problems where understanding the XOR property is much more important than implementing complicated logic.

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
