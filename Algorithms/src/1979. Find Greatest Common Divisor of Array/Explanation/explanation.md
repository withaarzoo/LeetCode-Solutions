# 1979. Find Greatest Common Divisor of Array

A clean and optimized solution for **LeetCode 1979 - Find Greatest Common Divisor of Array** using the **Euclidean Algorithm**. This repository explains the intuition, approach, complexity analysis, and provides implementations in **C++, Java, JavaScript, Python, and Go**.

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

The goal is to find the **greatest common divisor (GCD)** of the **smallest** and **largest** numbers in a given integer array.

Instead of finding the GCD of every element, the problem only asks for the GCD of the minimum and maximum values. This small detail makes the solution much simpler and more efficient.

### Input

- An integer array `nums`.

### Output

- Return the greatest common divisor of the smallest and largest elements in the array.

This problem is a good introduction to the **Euclidean Algorithm**, array traversal, and basic mathematical reasoning often used in competitive programming and coding interviews.

---

## Constraints

| Constraint | Value |
|------------|-------|
| Array Length | `2 <= nums.length <= 1000` |
| Element Value | `1 <= nums[i] <= 1000` |

---

## Intuition

The first thing I noticed was that the problem never asks for the GCD of the entire array.

It only cares about two numbers:

- the smallest value
- the largest value

That means there is no need to process every element after finding these two values.

Once I have both numbers, I can use the Euclidean Algorithm, which is one of the fastest ways to calculate the greatest common divisor.

This keeps the solution simple, clean, and efficient.

---

## Approach

I solve the problem in three simple steps.

1. Traverse the array once to find the smallest number.
2. During the same traversal, also find the largest number.
3. Apply the Euclidean Algorithm to calculate their GCD.

The Euclidean Algorithm repeatedly replaces the larger number with the remainder after division until the remainder becomes zero.

When that happens, the remaining number is the greatest common divisor.

---

## Data Structures Used

| Data Structure | Purpose |
|---------------|---------|
| Array | Stores the input numbers |
| Integer Variables | Store the current minimum and maximum values |

No additional data structures are required, which keeps the memory usage constant.

---

## Operations & Behavior Summary

The algorithm performs the following operations:

1. Read the input array.
2. Find the minimum element.
3. Find the maximum element.
4. Pass both numbers into the Euclidean Algorithm.
5. Repeatedly calculate the remainder.
6. Stop when the remainder becomes zero.
7. Return the final value as the answer.

This approach avoids unnecessary calculations and only processes the array once.

---

## Complexity

| Metric | Complexity | Explanation |
|--------|------------|-------------|
| Time Complexity | **O(n)** | The array is scanned once to find the minimum and maximum values. The GCD computation takes logarithmic time and is much smaller than the array traversal. |
| Space Complexity | **O(1)** | Only a few variables are used. No extra arrays, lists, or hash maps are created. |

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    // Function to find GCD using the Euclidean Algorithm
    int gcd(int a, int b) {
        // Keep reducing until the remainder becomes 0
        while (b != 0) {
            int temp = b;   // Store current value of b
            b = a % b;      // Update b with the remainder
            a = temp;       // Move previous b into a
        }

        // a now stores the GCD
        return a;
    }

    int findGCD(vector<int>& nums) {
        // Find the smallest element
        int mn = *min_element(nums.begin(), nums.end());

        // Find the largest element
        int mx = *max_element(nums.begin(), nums.end());

        // Return the GCD of the smallest and largest numbers
        return gcd(mn, mx);
    }
};
```

### Java

```java
class Solution {

    // Function to find GCD using the Euclidean Algorithm
    private int gcd(int a, int b) {
        // Continue until b becomes 0
        while (b != 0) {
            int temp = b;   // Store current b
            b = a % b;      // Update b with remainder
            a = temp;       // Move previous b into a
        }

        // a now contains the GCD
        return a;
    }

    public int findGCD(int[] nums) {

        // Initialize minimum and maximum with the first element
        int min = nums[0];
        int max = nums[0];

        // Find the minimum and maximum values
        for (int num : nums) {
            min = Math.min(min, num);
            max = Math.max(max, num);
        }

        // Return their GCD
        return gcd(min, max);
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} nums
 * @return {number}
 */

// Function to find GCD using the Euclidean Algorithm
function gcd(a, b) {
    // Repeat until b becomes 0
    while (b !== 0) {
        let temp = b;   // Store current b
        b = a % b;      // Update b with remainder
        a = temp;       // Move previous b into a
    }

    // a is the GCD
    return a;
}

var findGCD = function(nums) {

    // Initialize minimum and maximum
    let min = nums[0];
    let max = nums[0];

    // Find smallest and largest values
    for (const num of nums) {
        min = Math.min(min, num);
        max = Math.max(max, num);
    }

    // Return their GCD
    return gcd(min, max);
};
```

### Python3

```python
class Solution:
    # Function to find GCD using the Euclidean Algorithm
    def gcd(self, a, b):
        # Continue until b becomes 0
        while b:
            a, b = b, a % b

        # a now stores the GCD
        return a

    def findGCD(self, nums: List[int]) -> int:

        # Find the smallest element
        minimum = min(nums)

        # Find the largest element
        maximum = max(nums)

        # Return the GCD of the smallest and largest values
        return self.gcd(minimum, maximum)
```

### Go

```go
// Function to find GCD using the Euclidean Algorithm
func gcd(a, b int) int {
 // Continue until b becomes 0
 for b != 0 {
  a, b = b, a%b
 }

 // a now stores the GCD
 return a
}

func findGCD(nums []int) int {

 // Initialize minimum and maximum with the first element
 minimum := nums[0]
 maximum := nums[0]

 // Find the smallest and largest values
 for _, num := range nums {
  if num < minimum {
   minimum = num
  }
  if num > maximum {
   maximum = num
  }
 }

 // Return the GCD of the smallest and largest numbers
 return gcd(minimum, maximum)
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The overall logic remains exactly the same in every language. Only the syntax changes.

### C++

The C++ solution can use built-in functions such as `min_element()` and `max_element()` to quickly find the smallest and largest values in the array. After that, a helper function applies the Euclidean Algorithm until the remainder becomes zero.

### Java

In Java, I manually iterate through the array to track the minimum and maximum values because arrays do not have built-in methods for this. After the scan finishes, the same Euclidean Algorithm computes the answer.

### JavaScript

JavaScript also scans the array once to determine the smallest and largest numbers. A helper function repeatedly updates the two values using the remainder operation until the second value becomes zero.

### Python3

Python makes the first part very simple because the built-in `min()` and `max()` functions return the required values directly. The GCD logic remains identical to every other language.

### Go

Go follows the same approach as Java. I iterate through the slice once to find the minimum and maximum values, then repeatedly apply the Euclidean Algorithm until the answer is found.

Although the syntax changes from language to language, the algorithm and complexity remain exactly the same.

---

## Examples

### Example 1

**Input**

```
nums = [2,5,6,9,10]
```

**Output**

```
2
```

**Trace**

- Smallest number = 2
- Largest number = 10
- GCD(2,10) = 2

---

### Example 2

**Input**

```
nums = [7,5,6,8,3]
```

**Output**

```
1
```

**Trace**

- Smallest number = 3
- Largest number = 8
- GCD(3,8) = 1

---

### Example 3

**Input**

```
nums = [3,3]
```

**Output**

```
3
```

**Trace**

- Smallest number = 3
- Largest number = 3
- GCD(3,3) = 3

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

- The key observation is that only the smallest and largest numbers matter.
- There is no benefit in calculating the GCD of every element.
- The Euclidean Algorithm is much faster than checking every possible divisor.
- The solution uses constant extra memory.
- The algorithm works correctly even when every element in the array is the same.
- Since the constraints are small, this solution easily satisfies all limits while remaining the optimal approach.

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
