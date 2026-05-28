# 3300. Minimum Element After Replacement With Digit Sum - LeetCode Solution

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
* [How to Use / Run Locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

## Problem Summary

In this LeetCode problem, we are given an integer array `nums`.

For every number in the array, we replace it with the sum of its digits. After all replacements are completed, we need to return the smallest value present in the modified array.

The key observation is that we do not actually need to create a new array. We only need the minimum digit sum among all numbers.

This is a simple array traversal and digit sum problem that is commonly used to test basic number manipulation skills.

## Constraints

| Constraint                | Value                 |
| ------------------------- | --------------------- |
| `1 <= nums.length <= 100` | Array size            |
| `1 <= nums[i] <= 10^4`    | Value of each element |

## Intuition

My first observation was that the final array itself is not important.

The problem only asks for the minimum element after every number has been replaced by its digit sum.

Instead of storing all transformed values, I can calculate the digit sum of each number one by one and continuously track the smallest result.

Since each number contains only a few digits, computing the digit sum is very fast.

## Approach

1. Initialize a variable to store the minimum digit sum found so far.
2. Traverse the array.
3. For each number:

   * Extract every digit using modulo (`% 10`).
   * Add those digits together.
   * Remove digits one by one using integer division.
4. Compare the calculated digit sum with the current minimum.
5. Update the answer whenever a smaller digit sum is found.
6. Return the final minimum value.

This approach is straightforward, efficient, and uses constant extra space.

## Data Structures Used

### Array

The input array `nums` is traversed once to process every number.

### Integer Variables

A few integer variables are used to:

* Store the current digit sum
* Store the current number being processed
* Track the minimum digit sum

No additional data structures are required.

## Operations & Behavior Summary

The algorithm performs the following operations:

1. Start with a very large minimum value.
2. Pick the first number from the array.
3. Calculate its digit sum.
4. Update the minimum if needed.
5. Move to the next number.
6. Repeat until every element has been processed.
7. Return the smallest digit sum encountered.

In plain English:

* Convert each number into its digit sum.
* Keep checking which digit sum is the smallest.
* Return that smallest value.

## Complexity

| Metric           | Complexity | Explanation                                                                  |
| ---------------- | ---------- | ---------------------------------------------------------------------------- |
| Time Complexity  | O(n × d)   | `n` is the number of elements and `d` is the number of digits in each number |
| Space Complexity | O(1)       | Only a few variables are used, with no extra data structures                 |

Since `nums[i] ≤ 10^4`, the number of digits is at most 5, making the solution effectively linear.

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    // Helper function to calculate digit sum of a number
    int digitSum(int num) {
        int sum = 0;

        // Process every digit
        while (num > 0) {
            sum += num % 10; // Add last digit
            num /= 10;       // Remove last digit
        }

        return sum;
    }

    int minElement(vector<int>& nums) {
        int ans = INT_MAX;

        // Calculate digit sum for every element
        for (int num : nums) {
            ans = min(ans, digitSum(num));
        }

        return ans;
    }
};
```

### Java

```java
class Solution {

    // Helper function to calculate digit sum
    private int digitSum(int num) {
        int sum = 0;

        // Process every digit
        while (num > 0) {
            sum += num % 10; // Add last digit
            num /= 10;       // Remove last digit
        }

        return sum;
    }

    public int minElement(int[] nums) {
        int ans = Integer.MAX_VALUE;

        // Check digit sum of every number
        for (int num : nums) {
            ans = Math.min(ans, digitSum(num));
        }

        return ans;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} nums
 * @return {number}
 */
var minElement = function(nums) {

    // Helper function to calculate digit sum
    const digitSum = (num) => {
        let sum = 0;

        // Process every digit
        while (num > 0) {
            sum += num % 10;          // Add last digit
            num = Math.floor(num / 10); // Remove last digit
        }

        return sum;
    };

    let ans = Infinity;

    // Check digit sum of every element
    for (const num of nums) {
        ans = Math.min(ans, digitSum(num));
    }

    return ans;
};
```

### Python3

```python
class Solution:
    def minElement(self, nums: List[int]) -> int:

        # Helper function to calculate digit sum
        def digit_sum(num):
            total = 0

            # Process every digit
            while num > 0:
                total += num % 10  # Add last digit
                num //= 10         # Remove last digit

            return total

        ans = float('inf')

        # Check digit sum of every number
        for num in nums:
            ans = min(ans, digit_sum(num))

        return ans
```

### Go

```go
func minElement(nums []int) int {

    // Helper function to calculate digit sum
    digitSum := func(num int) int {
        sum := 0

        // Process every digit
        for num > 0 {
            sum += num % 10 // Add last digit
            num /= 10       // Remove last digit
        }

        return sum
    }

    ans := int(^uint(0) >> 1) // Maximum int value

    // Check digit sum of every number
    for _, num := range nums {
        current := digitSum(num)

        if current < ans {
            ans = current
        }
    }

    return ans
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic is identical in all five languages.

### Step 1: Initialize the Answer

We begin with a very large value for the answer.

This guarantees that the first digit sum we calculate will become the current minimum.

### Step 2: Process Every Number

We iterate through the entire array.

Each number is handled independently.

### Step 3: Calculate the Digit Sum

For a number like:

```text
1234
```

We repeatedly:

```text
1234 % 10 = 4
123 % 10 = 3
12 % 10 = 2
1 % 10 = 1
```

Adding them together gives:

```text
1 + 2 + 3 + 4 = 10
```

This digit extraction process continues until the number becomes zero.

### Step 4: Update the Minimum

After computing the digit sum, we compare it with the smallest value seen so far.

If the new digit sum is smaller, we replace the current answer.

### Step 5: Return the Result

Once every element has been processed, the stored minimum digit sum is the final answer.

## Examples

### Example 1

**Input**

```text
nums = [10,12,13,14]
```

**Digit Sums**

```text
10 → 1
12 → 3
13 → 4
14 → 5
```

**Output**

```text
1
```

### Example 2

**Input**

```text
nums = [1,2,3,4]
```

**Digit Sums**

```text
1 → 1
2 → 2
3 → 3
4 → 4
```

**Output**

```text
1
```

### Example 3

**Input**

```text
nums = [999,19,199]
```

**Digit Sums**

```text
999 → 27
19 → 10
199 → 19
```

**Output**

```text
10
```

## How to Use / Run Locally

### C++

Compile:

```bash
g++ solution.cpp -o solution
```

Run:

```bash
./solution
```

### Java

Compile:

```bash
javac Solution.java
```

Run:

```bash
java Solution
```

### JavaScript

Run:

```bash
node solution.js
```

### Python3

Run:

```bash
python solution.py
```

### Go

Run:

```bash
go run solution.go
```

## Notes & Optimizations

* No extra array is needed because only the minimum digit sum matters.
* The solution performs a single pass through the input array.
* The digit sum operation is extremely small because each number contains at most five digits.
* This is already the optimal approach for the given constraints.
* An alternative approach could create a transformed array first, but that would use unnecessary extra space.
* Constant space is preferable here because the transformed values are never reused.

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
