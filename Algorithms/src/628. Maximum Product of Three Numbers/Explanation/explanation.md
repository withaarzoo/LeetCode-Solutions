# 628. Maximum Product of Three Numbers

A beginner-friendly solution for **LeetCode 628 - Maximum Product of Three Numbers** using an **optimal O(n) algorithm**. This repository explains the intuition, approach, complexity analysis, and provides solutions in **C++, Java, JavaScript, Python3, and Go**. If you're preparing for coding interviews or improving your Data Structures and Algorithms (DSA) skills, this problem is a great example of handling arrays, negative numbers, and greedy observations efficiently.

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

The goal is to find the **maximum possible product of any three numbers** from the given integer array.

At first, it may seem like choosing the three largest numbers is always the best option. But this problem becomes interesting because the array can also contain negative numbers.

Since multiplying two negative numbers produces a positive number, the best answer can sometimes come from the **two smallest (most negative) numbers together with the largest positive number**.

The input is an integer array, and the output is a single integer representing the maximum product that can be formed using exactly three numbers.

---

## Constraints

| Constraint | Value |
|------------|-------|
| Array Length | 3 ≤ nums.length ≤ 10⁴ |
| Element Value | -1000 ≤ nums[i] ≤ 1000 |

---

## Intuition

The first thing I noticed was that this problem isn't just about finding the three largest values.

Negative numbers completely change the situation.

For example, if the array contains two very small negative numbers, multiplying them creates a large positive number. If that result is multiplied by the largest positive value, the final product can easily become larger than the product of the three biggest numbers.

That means there are only two combinations worth checking:

- The product of the three largest numbers.
- The product of the two smallest numbers and the largest number.

Instead of sorting the entire array, I realized I could find these five important numbers while scanning the array only once.

---

## Approach

I go through the array exactly one time.

During traversal, I keep track of:

- The largest number.
- The second largest number.
- The third largest number.
- The smallest number.
- The second smallest number.

Whenever I read a new value, I update these variables if needed.

After the traversal finishes, I calculate two products.

- Product of the three largest numbers.
- Product of the two smallest numbers and the largest number.

Finally, I return whichever product is larger.

This avoids sorting and gives the best possible time complexity.

---

## Data Structures Used

| Data Structure | Purpose |
|---------------|---------|
| Integer Variables | Store the three largest and two smallest values while traversing the array |

No extra arrays, lists, stacks, queues, or hash maps are required.

---

## Operations & Behavior Summary

The algorithm works in the following order:

1. Start with variables that represent the largest and smallest values.
2. Traverse the array once.
3. Continuously update the three largest numbers.
4. Continuously update the two smallest numbers.
5. Compute the product of the three largest numbers.
6. Compute the product of the two smallest numbers and the largest number.
7. Return whichever product is greater.

The algorithm never sorts the array and never performs multiple passes.

---

## Complexity

| Complexity | Value | Explanation |
|------------|-------|-------------|
| Time Complexity | O(n) | Every element is visited exactly once, where **n** is the size of the input array. |
| Space Complexity | O(1) | Only a fixed number of variables are used. No extra data structures are allocated. |

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int maximumProduct(vector<int>& nums) {
        // Store the three largest numbers
        int max1 = INT_MIN, max2 = INT_MIN, max3 = INT_MIN;

        // Store the two smallest numbers
        int min1 = INT_MAX, min2 = INT_MAX;

        // Traverse the array once
        for (int num : nums) {

            // Update the three largest numbers
            if (num >= max1) {
                max3 = max2;
                max2 = max1;
                max1 = num;
            } else if (num >= max2) {
                max3 = max2;
                max2 = num;
            } else if (num >= max3) {
                max3 = num;
            }

            // Update the two smallest numbers
            if (num <= min1) {
                min2 = min1;
                min1 = num;
            } else if (num <= min2) {
                min2 = num;
            }
        }

        // Product of the three largest numbers
        int product1 = max1 * max2 * max3;

        // Product of two smallest and the largest number
        int product2 = min1 * min2 * max1;

        // Return the larger product
        return max(product1, product2);
    }
};
```

### Java

```java
class Solution {
    public int maximumProduct(int[] nums) {

        // Store the three largest numbers
        int max1 = Integer.MIN_VALUE;
        int max2 = Integer.MIN_VALUE;
        int max3 = Integer.MIN_VALUE;

        // Store the two smallest numbers
        int min1 = Integer.MAX_VALUE;
        int min2 = Integer.MAX_VALUE;

        // Traverse the array once
        for (int num : nums) {

            // Update the three largest numbers
            if (num >= max1) {
                max3 = max2;
                max2 = max1;
                max1 = num;
            } else if (num >= max2) {
                max3 = max2;
                max2 = num;
            } else if (num >= max3) {
                max3 = num;
            }

            // Update the two smallest numbers
            if (num <= min1) {
                min2 = min1;
                min1 = num;
            } else if (num <= min2) {
                min2 = num;
            }
        }

        // Product of the three largest numbers
        int product1 = max1 * max2 * max3;

        // Product of the two smallest numbers and the largest number
        int product2 = min1 * min2 * max1;

        // Return the maximum product
        return Math.max(product1, product2);
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} nums
 * @return {number}
 */
var maximumProduct = function(nums) {

    // Store the three largest numbers
    let max1 = -Infinity;
    let max2 = -Infinity;
    let max3 = -Infinity;

    // Store the two smallest numbers
    let min1 = Infinity;
    let min2 = Infinity;

    // Traverse the array once
    for (const num of nums) {

        // Update the three largest numbers
        if (num >= max1) {
            max3 = max2;
            max2 = max1;
            max1 = num;
        } else if (num >= max2) {
            max3 = max2;
            max2 = num;
        } else if (num >= max3) {
            max3 = num;
        }

        // Update the two smallest numbers
        if (num <= min1) {
            min2 = min1;
            min1 = num;
        } else if (num <= min2) {
            min2 = num;
        }
    }

    // Product of the three largest numbers
    const product1 = max1 * max2 * max3;

    // Product of the two smallest numbers and the largest number
    const product2 = min1 * min2 * max1;

    // Return the larger product
    return Math.max(product1, product2);
};
```

### Python3

```python
class Solution:
    def maximumProduct(self, nums: List[int]) -> int:

        # Store the three largest numbers
        max1 = float("-inf")
        max2 = float("-inf")
        max3 = float("-inf")

        # Store the two smallest numbers
        min1 = float("inf")
        min2 = float("inf")

        # Traverse the array once
        for num in nums:

            # Update the three largest numbers
            if num >= max1:
                max3 = max2
                max2 = max1
                max1 = num
            elif num >= max2:
                max3 = max2
                max2 = num
            elif num >= max3:
                max3 = num

            # Update the two smallest numbers
            if num <= min1:
                min2 = min1
                min1 = num
            elif num <= min2:
                min2 = num

        # Product of the three largest numbers
        product1 = max1 * max2 * max3

        # Product of the two smallest numbers and the largest number
        product2 = min1 * min2 * max1

        # Return the larger product
        return max(product1, product2)
```

### Go

```go
func maximumProduct(nums []int) int {

 // Store the three largest numbers
 max1, max2, max3 := -1001, -1001, -1001

 // Store the two smallest numbers
 min1, min2 := 1001, 1001

 // Traverse the array once
 for _, num := range nums {

  // Update the three largest numbers
  if num >= max1 {
   max3 = max2
   max2 = max1
   max1 = num
  } else if num >= max2 {
   max3 = max2
   max2 = num
  } else if num >= max3 {
   max3 = num
  }

  // Update the two smallest numbers
  if num <= min1 {
   min2 = min1
   min1 = num
  } else if num <= min2 {
   min2 = num
  }
 }

 // Product of the three largest numbers
 product1 := max1 * max2 * max3

 // Product of the two smallest numbers and the largest number
 product2 := min1 * min2 * max1

 // Return the larger product
 if product1 > product2 {
  return product1
 }
 return product2
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The overall logic is exactly the same in every language. Only the syntax changes.

First, I create five variables.

Three variables store the largest values found so far.

Two variables store the smallest values found so far.

These variables are initialized using the smallest and largest possible values available in each language so that the first few numbers naturally replace them.

Next, I iterate through every element in the array.

For each number, I first check whether it belongs among the three largest numbers.

If it becomes the largest number, the previous largest values shift one position to make room.

If it is only the second or third largest, I update only those positions.

After updating the largest values, I repeat the same idea for the two smallest numbers.

Once the traversal finishes, I already have everything needed to compute the answer.

I calculate two possible products.

The first product comes from the three largest values.

The second product comes from the two smallest values and the largest value.

Finally, I compare both products and return the larger one.

This solution works for positive numbers, negative numbers, mixed arrays, and arrays containing zeros without requiring any special cases.

Since every language follows exactly the same algorithm, the behavior remains identical in C++, Java, JavaScript, Python3, and Go.

---

## Examples

### Example 1

**Input**

```text
nums = [1,2,3]
```

**Output**

```text
6
```

**Trace**

- Three largest numbers are 1, 2, and 3.
- Their product is 6.
- There are no useful negative numbers.
- Final answer is 6.

---

### Example 2

**Input**

```text
nums = [-10,-10,5,2]
```

**Output**

```text
500
```

**Trace**

- Product of three largest numbers:
  - 5 × 2 × -10 = -100
- Product of two smallest numbers and largest number:
  - -10 × -10 × 5 = 500
- 500 is larger.

---

### Example 3

**Input**

```text
nums = [-4,-3,-2,-1,60]
```

**Output**

```text
720
```

**Trace**

- Three largest numbers:
  - 60, -1, -2
  - Product = 120
- Two smallest numbers:
  - -4 and -3
- Largest number:
  - 60
- Product = 720
- Final answer is 720.

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

Run using Node.js.

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

- Sorting the array also works, but it requires O(n log n) time.
- A single traversal reduces the time complexity to O(n).
- The algorithm works correctly even when the array contains only negative numbers.
- Arrays containing zeros are naturally handled without additional conditions.
- Only five variables are required, making the space complexity constant.
- This is one of the most common interview problems involving greedy observation and array traversal.
- Understanding why two negative numbers can produce the maximum product is the key insight behind this problem.

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
