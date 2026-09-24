# Smallest Index With Digit Sum Equal to Index - LeetCode Solution

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
  - [TypeScript](#typescript)
  - [Python3](#python3)
  - [Go](#go)
- [Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)](#step-by-step-detailed-explanation-c-java-javascript-typescript-python3-go)
- [Examples](#examples)
- [How to Use / Run Locally](#how-to-use--run-locally)
- [Notes & Optimizations](#notes--optimizations)
- [Author](#author)

## Problem Summary

You are given an integer array called nums. The task is to find the smallest index i where the sum of the digits of the number at nums[i] is exactly equal to i itself.

If no such index exists in the array, the function should return -1.

This is a classic array traversal problem that also involves a simple digit sum calculation. It appears in LeetCode as problem 3550 and is tagged under array and math topics.

## Constraints

- 1 <= nums.length <= 100
- 0 <= nums[i] <= 1000

These limits make a straightforward linear scan completely safe and efficient.

## Intuition

The array is very small, with a maximum length of only 100. Each number is at most 1000, so it has at most four digits.  

Because we need the smallest index, the natural first idea is to start checking from index 0 and move right. The moment we find an index whose digit sum matches the index value, we can stop and return it. There is no need for sorting, hashing, or any advanced technique.

## Approach

We walk through the array from left to right using a simple loop.  

For every index i we calculate the digit sum of nums[i]. This is done by repeatedly taking the last digit with the modulo operator and then removing that digit by integer division.  

As soon as the digit sum equals i, we return i immediately.  

If the loop finishes without finding any matching index, we return -1.

The same logic works cleanly in every major programming language.

## Data Structures Used

- Input array (vector in C++, array in Java, slice in Go, list in Python, etc.): We only read from the given array. No extra arrays, maps, or sets are required.

No additional data structures are needed because a single linear pass is enough.

## Operations & Behavior Summary

1. Start at index 0.
2. For the current number, compute its digit sum by peeling off digits one by one.
3. Compare the digit sum with the current index.
4. If they match, return the index at once.
5. Otherwise move to the next index and repeat.
6. If the entire array is checked without a match, return -1.

This is essentially a linear scan combined with an O(1) digit-sum helper (since numbers are tiny).

## Complexity

| Type              | Complexity | Explanation |
|-------------------|------------|-------------|
| Time Complexity   | O(n × d)   | n is the length of the array (≤ 100). d is the number of digits in each element (≤ 4). In practice the whole solution runs in negligible time. |
| Space Complexity  | O(1)       | Only a few integer variables are used. No extra memory proportional to the input size is allocated. |

## Multi-language Solutions

### C++
```cpp
class Solution {
public:
    int smallestIndex(vector<int>& nums) {
        for (int i = 0; i < nums.size(); ++i) {
            int x = nums[i];
            int sum = 0;
            while (x > 0) {
                sum += x % 10;
                x /= 10;
            }
            if (sum == i) return i;
        }
        return -1;
    }
};
```

### Java
```java
class Solution {
    public int smallestIndex(int[] nums) {
        for (int i = 0; i < nums.length; ++i) {
            int x = nums[i];
            int sum = 0;
            while (x > 0) {
                sum += x % 10;
                x /= 10;
            }
            if (sum == i) return i;
        }
        return -1;
    }
}
```

### JavaScript
```javascript
/**
 * @param {number[]} nums
 * @return {number}
 */
var smallestIndex = function(nums) {
    for (let i = 0; i < nums.length; ++i) {
        let x = nums[i];
        let sum = 0;
        while (x > 0) {
            sum += x % 10;
            x = Math.floor(x / 10);
        }
        if (sum === i) return i;
    }
    return -1;
};
```

### TypeScript
```typescript
function smallestIndex(nums: number[]): number {
    for (let i = 0; i < nums.length; ++i) {
        let x = nums[i];
        let sum = 0;
        while (x > 0) {
            sum += x % 10;
            x = Math.floor(x / 10);
        }
        if (sum === i) return i;
    }
    return -1;
};
```

### Python3
```python
class Solution:
    def smallestIndex(self, nums: List[int]) -> int:
        for i in range(len(nums)):
            x = nums[i]
            s = 0
            while x > 0:
                s += x % 10
                x //= 10
            if s == i:
                return i
        return -1
```

### Go
```go
func smallestIndex(nums []int) int {
    for i := 0; i < len(nums); i++ {
        x := nums[i]
        sum := 0
        for x > 0 {
            sum += x % 10
            x /= 10
        }
        if sum == i {
            return i
        }
    }
    return -1
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

The core idea is identical across all six languages, so the reasoning stays the same.

We begin by iterating from the leftmost index. Using a zero-based index is important because the problem treats the first element as index 0.

Inside the loop we copy the current number into a temporary variable. This protects the original array and makes the digit extraction clean.

We then initialize a sum variable to zero. While the temporary number is still greater than zero we extract its last digit with the modulo-10 operation and add it to the sum. After that we remove the last digit by integer division by 10. When the number becomes zero, the sum holds the total of all digits.

A direct comparison follows: if the calculated sum equals the current index, we have found the answer and return it immediately. Early return is possible and desirable because the problem asks for the smallest index.

If the comparison fails, the loop simply continues to the next index. When every index has been examined without success, the function returns -1 to signal that no valid index exists.

Language-specific notes are minor:
- In JavaScript and TypeScript we use Math.floor for integer division.
- In Python we use the // operator.
- In Go, C++, and Java the integer division operator already truncates toward zero for positive numbers, which is exactly what we need.

Edge cases such as a single-element array, an array full of zeros, or numbers that are themselves zero are all handled naturally by the same logic.

## Examples

Example 1  
Input: nums = [1, 3, 2]  
Output: 2  
Trace:  
- Index 0 → digit sum of 1 is 1 ≠ 0  
- Index 1 → digit sum of 3 is 3 ≠ 1  
- Index 2 → digit sum of 2 is 2 == 2 → return 2

Example 2  
Input: nums = [1, 10, 11]  
Output: 1  
Trace:  
- Index 0 → digit sum of 1 is 1 ≠ 0  
- Index 1 → digit sum of 10 is 1 + 0 = 1 == 1 → return 1  
(The later match at index 2 is never examined because we already found a smaller index.)

Example 3  
Input: nums = [1, 2, 3]  
Output: -1  
Trace:  
- Index 0 → 1 ≠ 0  
- Index 1 → 2 ≠ 1  
- Index 2 → 3 ≠ 2  
No match exists, so the function returns -1.

## How to Use / Run Locally

C++  
1. Save the code in a file named main.cpp.  
2. Compile with: g++ -std=c++17 main.cpp -o main  
3. Run with: ./main  

Java  
1. Save the code in a file named Solution.java.  
2. Compile with: javac Solution.java  
3. Run with a small driver that creates an instance and calls the method.

JavaScript  
1. Save the code in a file named solution.js.  
2. Run with: node solution.js  
(You can add a few console.log test cases at the bottom.)

TypeScript  
1. Save the code in a file named solution.ts.  
2. Compile with: tsc solution.ts  
3. Run the generated JavaScript with node.

Python3  
1. Save the code in a file named solution.py.  
2. Run with: python3 solution.py  

Go  
1. Save the code in a file named main.go.  
2. Run with: go run main.go  

In every language you will need a small main or test harness that supplies sample arrays and prints the returned index.

## Notes & Optimizations

Because the constraints are extremely small, the straightforward linear scan is already optimal.  

An alternative approach would be to convert each number to a string and sum its character digits, but the arithmetic method is faster and uses less memory.  

One subtle edge case is when nums[i] equals 0. Its digit sum is 0, so it can only match index 0. The code handles this correctly without special handling.  

If the array length were much larger (for example 10^5) the same algorithm would still be acceptable because the digit-sum step is constant time. No further optimization is required for the given constraints.

## Author
[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)