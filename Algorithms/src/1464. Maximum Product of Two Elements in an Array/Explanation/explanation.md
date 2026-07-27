# 1464. Maximum Product of Two Elements in an Array

A simple and optimized solution for LeetCode 1464 - Maximum Product of Two Elements in an Array.

This repository explains the idea behind the solution, the algorithm, complexity analysis, and provides implementations in C++, Java, JavaScript, Python, and Go.

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
- [Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)](#step-by-step-detailed-explanation-c-java-javascript-python3-go)
- [Examples](#examples)
- [How to Use / Run Locally](#how-to-use--run-locally)
- [Notes & Optimizations](#notes--optimizations)
- [Author](#author)

---

## Problem Summary

The problem gives an integer array called `nums`.

Your task is to choose two different elements from the array so that the following value becomes as large as possible:

`(nums[i] - 1) × (nums[j] - 1)`

The goal is to return that maximum product.

Instead of checking every possible pair, we need an efficient algorithm that works well even for the largest allowed input size.

This problem is a good example of a simple array optimization question where finding the two largest values is enough to compute the answer.

---

## Constraints

| Constraint | Value |
|------------|-------|
| Array Length | 2 ≤ nums.length ≤ 500 |
| Element Value | 1 ≤ nums[i] ≤ 1000 |

---

## Intuition

The first thing I noticed was that the formula subtracts one from both numbers before multiplying them.

Since subtraction happens equally for every number, choosing the two largest values will always produce the largest possible result.

At first, sorting the array seemed like the easiest solution. After thinking a little more, I realized sorting the entire array is unnecessary because I only care about the largest and second largest numbers.

That means I can solve the problem in one pass while using only two variables.

---

## Approach

I keep track of two numbers while scanning the array.

The first variable stores the largest number seen so far.

The second variable stores the second largest number.

For every element in the array:

- If the current number is larger than the current maximum, I move the old maximum into the second maximum and update the maximum.
- Otherwise, if it is larger than the second maximum, I update only the second maximum.

After visiting every element, I simply calculate:

`(largest - 1) × (secondLargest - 1)`

This avoids sorting and gives the optimal solution.

---

## Data Structures Used

| Data Structure | Purpose |
| --------------- | --------- |
| Array | Stores the given input values |
| Integer Variables | Keep track of the largest and second largest numbers during traversal |

No additional arrays, stacks, queues, maps, or heaps are required.

---

## Operations & Behavior Summary

The algorithm works in four simple stages.

1. Initialize two variables to store the largest and second largest numbers.
2. Traverse the array exactly once.
3. Update the two variables whenever a larger value is found.
4. Compute the final product using the required formula and return the answer.

Since every element is processed only once, the algorithm stays fast and efficient.

---

## Complexity

| Complexity | Value | Explanation |
|------------|-------|-------------|
| Time Complexity | O(n) | The array is traversed exactly one time, where `n` is the number of elements in `nums`. |
| Space Complexity | O(1) | Only two extra variables are used regardless of the input size. |

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int maxProduct(vector<int>& nums) {
        // Store the largest value found so far
        int first = 0;

        // Store the second largest value found so far
        int second = 0;

        // Traverse the array once
        for (int num : nums) {

            // If current number becomes the largest
            if (num >= first) {
                // Old largest becomes second largest
                second = first;

                // Update largest
                first = num;
            }
            // Otherwise check if it becomes the second largest
            else if (num > second) {
                second = num;
            }
        }

        // Return the required product
        return (first - 1) * (second - 1);
    }
};
```

### Java

```java
class Solution {
    public int maxProduct(int[] nums) {

        // Store the largest value found so far
        int first = 0;

        // Store the second largest value found so far
        int second = 0;

        // Traverse every element once
        for (int num : nums) {

            // If current number becomes the largest
            if (num >= first) {
                // Old largest becomes second largest
                second = first;

                // Update largest
                first = num;
            }
            // Otherwise update second largest if needed
            else if (num > second) {
                second = num;
            }
        }

        // Return the required product
        return (first - 1) * (second - 1);
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} nums
 * @return {number}
 */
var maxProduct = function(nums) {

    // Store the largest value found so far
    let first = 0;

    // Store the second largest value found so far
    let second = 0;

    // Traverse the array once
    for (const num of nums) {

        // If current number becomes the largest
        if (num >= first) {
            // Old largest becomes second largest
            second = first;

            // Update largest
            first = num;
        }
        // Otherwise update second largest if needed
        else if (num > second) {
            second = num;
        }
    }

    // Return the required product
    return (first - 1) * (second - 1);
};
```

### Python3

```python
class Solution:
    def maxProduct(self, nums: List[int]) -> int:

        # Store the largest value found so far
        first = 0

        # Store the second largest value found so far
        second = 0

        # Traverse the array once
        for num in nums:

            # If current number becomes the largest
            if num >= first:
                # Old largest becomes second largest
                second = first

                # Update largest
                first = num

            # Otherwise update second largest if needed
            elif num > second:
                second = num

        # Return the required product
        return (first - 1) * (second - 1)
```

### Go

```go
func maxProduct(nums []int) int {

    // Store the largest value found so far
    first := 0

    // Store the second largest value found so far
    second := 0

    // Traverse the array once
    for _, num := range nums {

        // If current number becomes the largest
        if num >= first {

            // Old largest becomes second largest
            second = first

            // Update largest
            first = num

        } else if num > second {
            // Otherwise update second largest if needed
            second = num
        }
    }

    // Return the required product
    return (first - 1) * (second - 1)
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic is exactly the same in every programming language.

Only the syntax changes.

First, two variables are created.

One stores the largest number found so far.

The other stores the second largest number.

Next, the algorithm starts scanning the array from beginning to end.

For every value, it checks whether the current number is larger than the current maximum.

If it is, the previous maximum becomes the second maximum, and the current number becomes the new maximum.

If the current number is not the largest, another check is performed.

This second check determines whether the number is larger than the current second maximum.

If it is, only the second maximum is updated.

Nothing else changes.

Once every element has been processed, the algorithm has successfully identified the two largest values in the array.

The final answer is calculated by subtracting one from each value and multiplying them together.

This approach is efficient because no sorting is required.

Every language version follows exactly the same reasoning.

The only difference is the syntax used for loops, variable declarations, and function definitions.

---

## Examples

### Example 1

**Input**

```text
nums = [3,4,5,2]
```

**Output**

```text
12
```

**Trace**

Largest values are `5` and `4`.

Result:

```text
(5 - 1) × (4 - 1)
= 4 × 3
= 12
```

---

### Example 2

**Input**

```text
nums = [1,5,4,5]
```

**Output**

```text
16
```

**Trace**

Largest values are `5` and `5`.

Result:

```text
(5 - 1) × (5 - 1)
= 4 × 4
= 16
```

---

### Example 3

**Input**

```text
nums = [3,7]
```

**Output**

```text
12
```

**Trace**

Largest values are `7` and `3`.

Result:

```text
(7 - 1) × (3 - 1)
= 6 × 2
= 12
```

---

## How to Use / Run Locally

Clone the repository.

```bash
git clone <repository-url>
```

Move into the project directory.

```bash
cd <repository-name>
```

### C++

Compile:

```bash
g++ solution.cpp -o solution
```

Run:

```bash
./solution
```

---

### Java

Compile:

```bash
javac Solution.java
```

Run:

```bash
java Solution
```

---

### JavaScript

Run using Node.js:

```bash
node solution.js
```

---

### Python3

Run:

```bash
python solution.py
```

or

```bash
python3 solution.py
```

---

### Go

Run:

```bash
go run solution.go
```

---

## Notes & Optimizations

- Finding only the two largest values is enough to solve the problem.
- Sorting also works but increases the time complexity to `O(n log n)`.
- A single traversal gives the optimal `O(n)` solution.
- The problem guarantees at least two numbers, so no additional validation is required.
- Duplicate maximum values are handled naturally.
- No extra data structures are needed, making the solution memory efficient.
- This is a common interview problem for practicing array traversal and tracking maximum values.

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
