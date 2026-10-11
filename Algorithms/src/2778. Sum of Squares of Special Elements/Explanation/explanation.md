# Sum of Squares of Special Elements | LeetCode 2778 Solution

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

You are given a 1-indexed integer array called nums. The length of this array is n.  

An element nums[i] is called special if the index i divides n completely. In other words, n % i equals zero.  

Your task is to find every special element, square each of them, and return the sum of those squares.  

This is a classic array traversal problem that tests basic understanding of 1-based indexing and the modulo operator.

## Constraints

- 1 <= nums.length == n <= 50
- 1 <= nums[i] <= 50

## Intuition

When I first read the problem I noticed two important things.  

First, the array is 1-indexed, so position 1 is the first element. Second, a position is special only when it is a divisor of the array length.  

Because n is at most 50, I realized I can simply walk through every possible index from 1 to n, check the divisibility condition, and accumulate the square of the value at that index. No fancy data structure or sorting is needed.

## Approach

I start by storing the length of the array in a variable n.  

I create a variable to hold the running sum and set it to zero.  

Then I loop over every index i starting from 1 up to n.  

For each i I ask whether n is divisible by i. If the answer is yes, I take the element at the corresponding 0-based position (nums[i-1]), square it, and add the result to the sum.  

When the loop finishes the sum already contains the answer, so I return it.

## Data Structures Used

- Array / Vector / List: The input itself. I only read values from it; I never modify it or create extra arrays.
- Integer variables: I use a few plain integers for the length, the loop counter, and the running sum.

No additional data structures are required because the problem only needs a single linear scan.

## Operations & Behavior Summary

1. Read the length of the input array and store it as n.  
2. Initialize an answer variable to 0.  
3. For every candidate index i from 1 to n:  
   - Test whether n % i == 0.  
   - If true, square the value sitting at position i (0-based index i-1) and add it to the answer.  
4. After examining all indices, return the accumulated answer.

## Complexity

| Type              | Value | Explanation                                                                 |
|-------------------|-------|-----------------------------------------------------------------------------|
| Time Complexity   | O(n)  | We examine each of the n indices exactly once and perform constant-time work. |
| Space Complexity  | O(1)  | Only a handful of integer variables are used; no extra memory grows with n.  |

## Multi-language Solutions

### C++
```cpp
class Solution {
public:
    int sumOfSquares(vector<int>& nums) {
        int n = nums.size();
        int ans = 0;
        for (int i = 1; i <= n; i++) {
            if (n % i == 0) {
                ans += nums[i - 1] * nums[i - 1];
            }
        }
        return ans;
    }
};
```

### Java
```java
class Solution {
    public int sumOfSquares(int[] nums) {
        int n = nums.length;
        int ans = 0;
        for (int i = 1; i <= n; i++) {
            if (n % i == 0) {
                ans += nums[i - 1] * nums[i - 1];
            }
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
var sumOfSquares = function(nums) {
    const n = nums.length;
    let ans = 0;
    for (let i = 1; i <= n; i++) {
        if (n % i === 0) {
            ans += nums[i - 1] * nums[i - 1];
        }
    }
    return ans;
};
```

### TypeScript
```typescript
function sumOfSquares(nums: number[]): number {
    const n = nums.length;
    let ans = 0;
    for (let i = 1; i <= n; i++) {
        if (n % i === 0) {
            ans += nums[i - 1] * nums[i - 1];
        }
    }
    return ans;
};
```

### Python3
```python
class Solution:
    def sumOfSquares(self, nums: List[int]) -> int:
        n = len(nums)
        ans = 0
        for i in range(1, n + 1):
            if n % i == 0:
                ans += nums[i - 1] * nums[i - 1]
        return ans
```

### Go
```go
func sumOfSquares(nums []int) int {
    n := len(nums)
    ans := 0
    for i := 1; i <= n; i++ {
        if n%i == 0 {
            ans += nums[i-1] * nums[i-1]
        }
    }
    return ans
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

The logic is identical in every language; only the syntax changes.

I first obtain the length of the array. In C++ and Java this is done with .size() or .length; in JavaScript, TypeScript, Python and Go it is simply the length property or the len function.

I declare an integer (or number) that will store the final result and set it to zero.

Then I open a loop that starts at 1 and runs while the counter is less than or equal to the length.  

Inside the loop I perform a single divisibility test: does the length modulo the current counter equal zero?  

If it does, the current counter is a valid 1-based index. Because every language stores the array starting at index 0, I subtract one from the counter to reach the correct element, square that element, and add the square to the result.

I continue until every possible index has been checked.  

Finally I return the result variable.

Edge cases are naturally handled:  
- When n equals 1 the only special index is 1, so the single element is squared and returned.  
- When every index divides n (for example n = 6) all corresponding elements are included.  
- Because the constraints guarantee positive numbers, I never worry about zero or negative values.

## Examples

**Example 1**  
Input: nums = [1, 2, 3, 4]  
n = 4  
Special indices: 1 (4 % 1 == 0), 2 (4 % 2 == 0), 4 (4 % 4 == 0)  
Squares: 1² + 2² + 4² = 1 + 4 + 16 = 21  
Output: 21

**Example 2**  
Input: nums = [2, 7, 1, 19, 18, 3]  
n = 6  
Special indices: 1, 2, 3, 6  
Squares: 2² + 7² + 1² + 3² = 4 + 49 + 1 + 9 = 63  
Output: 63

**Example 3**  
Input: nums = [5]  
n = 1  
Special index: 1  
Square: 5² = 25  
Output: 25

## How to Use / Run Locally

**C++**  
1. Save the code in a file named main.cpp.  
2. Compile with: g++ -std=c++17 main.cpp -o main  
3. Run with: ./main

**Java**  
1. Save the code in a file named Solution.java.  
2. Compile with: javac Solution.java  
3. Run with: java Solution

**JavaScript**  
1. Save the code in a file named solution.js.  
2. Run with: node solution.js

**TypeScript**  
1. Save the code in a file named solution.ts.  
2. Compile with: tsc solution.ts  
3. Run the generated JavaScript file with: node solution.js

**Python3**  
1. Save the code in a file named solution.py.  
2. Run with: python3 solution.py

**Go**  
1. Save the code in a file named main.go.  
2. Run with: go run main.go

In each case you will need to add a small driver that creates a sample array and prints the result of the function.

## Notes & Optimizations

Because the maximum length is only 50, the straightforward O(n) scan is already optimal and easy to understand.  

An alternative approach would be to find all divisors of n first and then look up only those positions. That method is also O(sqrt(n)) for finding divisors plus O(d) for the sum, but it adds unnecessary complexity for such a small constraint.  

The current solution never allocates extra memory and works correctly for every valid input, including the single-element case.

## Author
[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)